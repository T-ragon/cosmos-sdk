//go:build rocksdb
// +build rocksdb

package rocksdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"slices"

	"github.com/linxGnu/grocksdb"

	corestore "cosmossdk.io/core/store"
	"cosmossdk.io/store/v2"
	"cosmossdk.io/store/v2/errors"
	"cosmossdk.io/store/v2/storage"
	"cosmossdk.io/store/v2/storage/util"
)

const (
	TimestampSize = 8

	batchBufferCount = 1000

	StorePrefixTpl   = "s/k:%s/"
	latestVersionKey = "s/latest"
)

var (
	_ storage.Database         = (*Database)(nil)
	_ store.UpgradableDatabase = (*Database)(nil)

	defaultWriteOpts = grocksdb.NewDefaultWriteOptions()
	defaultReadOpts  = grocksdb.NewDefaultReadOptions()
)

type Database struct {
	storage  *grocksdb.DB
	cfHandle *grocksdb.ColumnFamilyHandle

	// tsLow reflects the full_history_ts_low CF value, which is earliest version
	// supported
	tsLow uint64
}

func New(dataDir string) (*Database, error) {
	storage, cfHandle, err := OpenRocksDB(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open RocksDB: %w", err)
	}

	slice, err := storage.GetFullHistoryTsLow(cfHandle)
	if err != nil {
		return nil, fmt.Errorf("failed to get full_history_ts_low: %w", err)
	}

	var tsLow uint64
	tsLowBz := copyAndFreeSlice(slice)
	if len(tsLowBz) > 0 {
		tsLow = binary.LittleEndian.Uint64(tsLowBz)
	}

	return &Database{
		storage:  storage,
		cfHandle: cfHandle,
		tsLow:    tsLow,
	}, nil
}

func NewWithDB(storage *grocksdb.DB, cfHandle *grocksdb.ColumnFamilyHandle) (*Database, error) {
	slice, err := storage.GetFullHistoryTsLow(cfHandle)
	if err != nil {
		return nil, fmt.Errorf("failed to get full_history_ts_low: %w", err)
	}

	var tsLow uint64
	tsLowBz := copyAndFreeSlice(slice)
	if len(tsLowBz) > 0 {
		tsLow = binary.LittleEndian.Uint64(tsLowBz)
	}

	return &Database{
		storage:  storage,
		cfHandle: cfHandle,
		tsLow:    tsLow,
	}, nil
}

func (db *Database) Close() error {
	db.storage.Close()

	db.storage = nil
	db.cfHandle = nil

	return nil
}

func (db *Database) NewBatch(version uint64) (store.Batch, error) {
	return NewBatch(db, version), nil
}

func (db *Database) getSlice(storeKey []byte, version uint64, key []byte) (*grocksdb.Slice, error) {
	if version < db.tsLow {
		return nil, errors.ErrVersionPruned{EarliestVersion: db.tsLow, RequestedVersion: version}
	}

	return db.storage.GetCF(
		newTSReadOptions(version),
		db.cfHandle,
		prependStoreKey(storeKey, key),
	)
}

func (db *Database) SetLatestVersion(version uint64) error {
	var ts [TimestampSize]byte
	binary.LittleEndian.PutUint64(ts[:], version)

	return db.storage.Put(defaultWriteOpts, []byte(latestVersionKey), ts[:])
}

func (db *Database) GetLatestVersion() (uint64, error) {
	bz, err := db.storage.GetBytes(defaultReadOpts, []byte(latestVersionKey))
	if err != nil {
		return 0, err
	}

	if len(bz) == 0 {
		// in case of a fresh database
		return 0, nil
	}

	return binary.LittleEndian.Uint64(bz), nil
}

func (db *Database) Has(storeKey []byte, version uint64, key []byte) (bool, error) {
	slice, err := db.getSlice(storeKey, version, key)
	if err != nil {
		return false, err
	}

	return slice.Exists(), nil
}

func (db *Database) Get(storeKey []byte, version uint64, key []byte) ([]byte, error) {
	slice, err := db.getSlice(storeKey, version, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get RocksDB slice: %w", err)
	}

	return copyAndFreeSlice(slice), nil
}

// Prune prunes all versions up to and including the provided version argument.
// Internally, this performs a manual compaction, the data with older timestamp
// will be GCed by compaction.
func (db *Database) Prune(version uint64) error {
	tsLow := version + 1 // we increment by 1 to include the provided version

	var ts [TimestampSize]byte
	binary.LittleEndian.PutUint64(ts[:], tsLow)
	compactOpts := grocksdb.NewCompactRangeOptions()
	compactOpts.SetFullHistoryTsLow(ts[:])
	db.storage.CompactRangeCFOpt(db.cfHandle, grocksdb.Range{}, compactOpts)

	db.tsLow = tsLow
	return nil
}

func (db *Database) Iterator(storeKey []byte, version uint64, start, end []byte) (corestore.Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, errors.ErrKeyEmpty
	}

	if start != nil && end != nil && bytes.Compare(start, end) > 0 {
		return nil, errors.ErrStartAfterEnd
	}

	prefix := storePrefix(storeKey)
	start, end = util.IterateWithPrefix(prefix, start, end)

	itr := db.storage.NewIteratorCF(newTSReadOptions(version), db.cfHandle)
	return newRocksDBIterator(itr, prefix, start, end, false), nil
}

func (db *Database) ReverseIterator(storeKey []byte, version uint64, start, end []byte) (corestore.Iterator, error) {
	if (start != nil && len(start) == 0) || (end != nil && len(end) == 0) {
		return nil, errors.ErrKeyEmpty
	}

	if start != nil && end != nil && bytes.Compare(start, end) > 0 {
		return nil, errors.ErrStartAfterEnd
	}

	prefix := storePrefix(storeKey)
	start, end = util.IterateWithPrefix(prefix, start, end)

	itr := db.storage.NewIteratorCF(newTSReadOptions(version), db.cfHandle)
	return newRocksDBIterator(itr, prefix, start, end, true), nil
}

// PruneStoreKey will do nothing for RocksDB, it will be pruned by compaction
// when the version is pruned
func (db *Database) PruneStoreKey(storeKey []byte) error {
	return nil
}

func (db *Database) MigrateStoreKey(fromStoreKey, toStoreKey []byte) error {
	latestVersion, err := db.GetLatestVersion()
	if err != nil {
		return fmt.Errorf("failed to get latest version: %w", err)
	}

	// need to copy all keys with the given fromStoreKey prefix to toStoreKey
	readOpts := newTSReadOptions(latestVersion)
	readOpts.SetIterStartTimestamp([]byte{0, 0, 0, 0, 0, 0, 0, 0})
	itr := db.storage.NewIteratorCF(readOpts, db.cfHandle)
	prefix := storePrefix(fromStoreKey)
	ritr := newRocksDBIterator(itr, prefix, nil, nil, false)
	defer ritr.Close()

	batch := grocksdb.NewWriteBatch()
	defer batch.Destroy()
	for ; ritr.Valid(); ritr.Next() {
		// replace the prefix
		key := ritr.Key()
		key = key[:len(key)-16] // remove the timestamp
		prefixedKey := append(storePrefix(toStoreKey), key...)
		batch.PutCFWithTS(db.cfHandle, prefixedKey, ritr.Timestamp(), ritr.Value())
		if batch.Count() >= batchBufferCount {
			if err := db.storage.Write(defaultWriteOpts, batch); err != nil {
				return err
			}
			batch.Clear()
		}
	}

	return db.storage.Write(defaultWriteOpts, batch)
}

// newTSReadOptions returns ReadOptions used in the RocksDB column family read.
func newTSReadOptions(version uint64) *grocksdb.ReadOptions {
	var ts [TimestampSize]byte
	binary.LittleEndian.PutUint64(ts[:], version)

	readOpts := grocksdb.NewDefaultReadOptions()
	readOpts.SetTimestamp(ts[:])

	return readOpts
}

func storePrefix(storeKey []byte) []byte {
	return append([]byte(StorePrefixTpl), storeKey...)
}

func prependStoreKey(storeKey, key []byte) []byte {
	return append(storePrefix(storeKey), key...)
}

// copyAndFreeSlice will copy a given RocksDB slice and free it. If the slice does
// not exist, <nil> will be returned.
func copyAndFreeSlice(s *grocksdb.Slice) []byte {
	defer s.Free()
	if !s.Exists() {
		return nil
	}

	return slices.Clone(s.Data())
}

func readOnlySlice(s *grocksdb.Slice) []byte {
	if !s.Exists() {
		return nil
	}

	return s.Data()
}
