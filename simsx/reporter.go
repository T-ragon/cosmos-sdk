package simsx

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/cosmos/gogoproto/proto"
	"golang.org/x/exp/maps"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

type ReportResult struct {
	Status     string
	Error      error
	MsgProtoBz []byte
}

func (r ReportResult) String() string {
	return fmt.Sprintf("error: %q, status: %q", r.Error, r.Status)
}

type SimulationReporter interface {
	WithScope(msg sdk.Msg, optionalSkipHook ...SkipHook) SimulationReporter
	Skip(comment string)
	Skipf(comment string, args ...any)
	// IsSkipped returns true when skipped or completed
	IsSkipped() bool
	ToLegacyOperationMsg() simtypes.OperationMsg
	// Fail complete with failure
	Fail(err error, comments ...string)
	// Success complete with success
	Success(msg sdk.Msg, comments ...string)
	// Close returns error captured on fail
	Close() error
	Comment() string
}

var _ SimulationReporter = &BasicSimulationReporter{}

type ReporterStatus uint8

const (
	undefined ReporterStatus = iota
	skipped   ReporterStatus = iota
	completed ReporterStatus = iota
)

func (s ReporterStatus) String() string {
	switch s {
	case skipped:
		return "skipped"
	case completed:
		return "completed"
	default:
		return "undefined"
	}
}

type SkipHook interface {
	Skip(args ...any)
}

var _ SkipHook = SkipHookFn(nil)

type SkipHookFn func(args ...any)

func (s SkipHookFn) Skip(args ...any) {
	s(args...)
}

type BasicSimulationReporter struct {
	skipCallbacks     []SkipHook
	completedCallback func(reporter *BasicSimulationReporter)
	module            string
	msgTypeURL        string

	status atomic.Uint32

	cMX        sync.RWMutex
	comments   []string
	error      error
	msgProtoBz []byte

	summary *ExecutionSummary
}

// NewBasicSimulationReporter constructor that accepts an optional callback hook that is called on state transition to skipped status
// A typical implementation for this hook is testing.T
func NewBasicSimulationReporter(optionalSkipHook ...SkipHook) *BasicSimulationReporter {
	r := &BasicSimulationReporter{
		skipCallbacks: optionalSkipHook,
		summary:       NewExecutionSummary(),
	}
	r.completedCallback = func(child *BasicSimulationReporter) {
		r.summary.Add(child.module, child.msgTypeURL, ReporterStatus(child.status.Load()), child.Comment())
	}
	return r
}

func (x *BasicSimulationReporter) WithScope(msg sdk.Msg, optionalSkipHook ...SkipHook) SimulationReporter {
	typeURL := sdk.MsgTypeURL(msg)
	r := &BasicSimulationReporter{
		skipCallbacks:     append(x.skipCallbacks, optionalSkipHook...),
		completedCallback: x.completedCallback,
		error:             x.error,
		msgProtoBz:        x.msgProtoBz,
		msgTypeURL:        typeURL,
		module:            sdk.GetModuleNameFromTypeURL(typeURL),
		comments:          slices.Clone(x.comments),
	}
	r.status.Store(x.status.Load())
	return r
}

func (x *BasicSimulationReporter) Skip(comment string) {
	x.toStatus(skipped, comment)
}

func (x *BasicSimulationReporter) Skipf(comment string, args ...any) {
	x.Skip(fmt.Sprintf(comment, args...))
}

func (x *BasicSimulationReporter) IsSkipped() bool {
	return ReporterStatus(x.status.Load()) > undefined
}

func (x *BasicSimulationReporter) ToLegacyOperationMsg() simtypes.OperationMsg {
	switch ReporterStatus(x.status.Load()) {
	case skipped:
		return simtypes.NoOpMsg(x.module, x.msgTypeURL, x.Comment())
	case completed:
		x.cMX.RLock()
		err := x.error
		x.cMX.RUnlock()
		if err == nil {
			return simtypes.NoOpMsg(x.module, x.msgTypeURL, x.Comment())
		} else {
			return simtypes.NewOperationMsgBasic(x.module, x.msgTypeURL, x.Comment(), true, x.msgProtoBz)
		}
	default:
		x.Fail(errors.New("operation aborted before msg was executed"))
		return x.ToLegacyOperationMsg()
	}
}

func (x *BasicSimulationReporter) Fail(err error, comments ...string) {
	if !x.toStatus(completed, comments...) {
		return
	}
	x.cMX.Lock()
	defer x.cMX.Unlock()
	x.error = err
}

func (x *BasicSimulationReporter) Success(msg sdk.Msg, comments ...string) {
	if !x.toStatus(completed, comments...) {
		return
	}
	protoBz, err := proto.Marshal(msg) // todo: not great to capture the proto bytes here again but legacy test are using it.
	if err != nil {
		panic(err)
	}
	x.cMX.Lock()
	defer x.cMX.Unlock()
	x.msgProtoBz = protoBz
}

func (x *BasicSimulationReporter) Close() error {
	x.completedCallback(x)
	x.cMX.RLock()
	defer x.cMX.RUnlock()
	return x.error
}

func (x *BasicSimulationReporter) toStatus(next ReporterStatus, comments ...string) bool {
	oldStatus := ReporterStatus(x.status.Load())
	if oldStatus > next {
		panic(fmt.Sprintf("can not switch from status %s to %s", oldStatus, next))
	}
	if !x.status.CompareAndSwap(uint32(oldStatus), uint32(next)) {
		return false
	}
	x.cMX.Lock()
	newComments := append(x.comments, comments...)
	x.comments = newComments
	x.cMX.Unlock()

	if oldStatus != skipped && next == skipped {
		prettyComments := strings.Join(newComments, ", ")
		for _, hook := range x.skipCallbacks {
			hook.Skip(prettyComments)
		}
	}
	return true
}

func (x *BasicSimulationReporter) Comment() string {
	x.cMX.RLock()
	defer x.cMX.RUnlock()
	return strings.Join(x.comments, ", ")
}

func (x *BasicSimulationReporter) Summary() *ExecutionSummary {
	return x.summary
}

type ExecutionSummary struct {
	mx      sync.RWMutex
	counts  map[string]int
	reasons map[string]map[string]int
}

func NewExecutionSummary() *ExecutionSummary {
	return &ExecutionSummary{counts: make(map[string]int), reasons: make(map[string]map[string]int)}
}

func (s *ExecutionSummary) Add(module, url string, status ReporterStatus, comment string) {
	s.mx.Lock()
	defer s.mx.Unlock()
	combinedKey := fmt.Sprintf("%s_%s", module, status.String())
	s.counts[combinedKey] += 1
	if status == completed {
		return
	}
	r, ok := s.reasons[url]
	if !ok {
		r = make(map[string]int)
		s.reasons[url] = r
	}
	r[comment] += 1
}

func (s *ExecutionSummary) String() string {
	s.mx.RLock()
	defer s.mx.RUnlock()
	keys := maps.Keys(s.counts)
	sort.Strings(keys)
	var sb strings.Builder
	for _, key := range keys {
		sb.WriteString(fmt.Sprintf("%s: %d\n", key, s.counts[key]))
	}
	for m, c := range s.reasons {
		sb.WriteString(fmt.Sprintf("%d\t%s: %q\n", sum(maps.Values(c)), m, maps.Keys(c)))
	}
	return sb.String()
}

func sum(values []int) int {
	var r int
	for _, v := range values {
		r += v
	}
	return r
}
