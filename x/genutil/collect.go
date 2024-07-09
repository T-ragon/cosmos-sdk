package genutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	cfg "github.com/cometbft/cometbft/config"

	"cosmossdk.io/core/address"
	stakingtypes "cosmossdk.io/x/staking/types"

	"github.com/T-ragon/cosmos-sdk/v3/client"
	"github.com/T-ragon/cosmos-sdk/v3/codec"
	sdk "github.com/T-ragon/cosmos-sdk/v3/types"
	"github.com/T-ragon/cosmos-sdk/v3/x/genutil/types"
)

// GenAppStateFromConfig gets the genesis app state from the config
func GenAppStateFromConfig(cdc codec.JSONCodec, txEncodingConfig client.TxEncodingConfig,
	config *cfg.Config, initCfg types.InitConfig, genesis *types.AppGenesis,
	validator types.MessageValidator, valAddrCodec address.ValidatorAddressCodec, addressCodec address.Codec,
) (appState json.RawMessage, err error) {
	// process genesis transactions, else create default genesis.json
	appGenTxs, persistentPeers, err := CollectTxs(
		txEncodingConfig.TxJSONDecoder(), config.Moniker, initCfg.GenTxsDir, genesis, validator, valAddrCodec, addressCodec)
	if err != nil {
		return appState, err
	}

	config.P2P.PersistentPeers = persistentPeers
	cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return appState, errors.New("there must be at least one genesis tx")
	}

	// create the app state
	appGenesisState, err := types.GenesisStateFromAppGenesis(genesis)
	if err != nil {
		return appState, err
	}

	appGenesisState, err = SetGenTxsInAppGenesisState(cdc, txEncodingConfig.TxJSONEncoder(), appGenesisState, appGenTxs)
	if err != nil {
		return appState, err
	}

	appState, err = json.MarshalIndent(appGenesisState, "", "  ")
	if err != nil {
		return appState, err
	}

	genesis.AppState = appState
	err = ExportGenesisFile(genesis, config.GenesisFile())

	return appState, err
}

// CollectTxs processes and validates application's genesis Txs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
func CollectTxs(txJSONDecoder sdk.TxDecoder, moniker, genTxsDir string,
	genesis *types.AppGenesis,
	validator types.MessageValidator, valAddrCodec address.ValidatorAddressCodec, addressCodec address.Codec,
) (appGenTxs []sdk.Tx, persistentPeers string, err error) {
	// prepare a map of all balances in genesis state to then validate
	// against the validators addresses
	var appState map[string]json.RawMessage
	if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	fos, err := os.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	// TODO (https://github.com/T-ragon/cosmos-sdk/v3/issues/17815):
	// Examine CPU and RAM profiles to see if we can parsing
	// and ValidateAndGetGenTx concurrent.
	for _, fo := range fos {
		if fo.IsDir() {
			continue
		}
		if !strings.HasSuffix(fo.Name(), ".json") {
			continue
		}

		// get the genTx
		jsonRawTx, err := os.ReadFile(filepath.Join(genTxsDir, fo.Name()))
		if err != nil {
			return appGenTxs, persistentPeers, err
		}

		genTx, err := types.ValidateAndGetGenTx(jsonRawTx, txJSONDecoder, validator)
		if err != nil {
			return appGenTxs, persistentPeers, err
		}

		appGenTxs = append(appGenTxs, genTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"

		memoTx, ok := genTx.(sdk.TxWithMemo)
		if !ok {
			return appGenTxs, persistentPeers, fmt.Errorf("expected TxWithMemo, got %T", genTx)
		}
		nodeAddrIP := memoTx.GetMemo()

		// genesis transactions must be single-message
		msgs := genTx.GetMsgs()

		// TODO abstract out staking message validation back to staking
		msg := msgs[0].(*stakingtypes.MsgCreateValidator)

		// exclude itself from persistent peers
		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}
