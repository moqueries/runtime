package impl

import (
	"math/bits"
	"sync/atomic"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

var titler = cases.Title(language.Und, cases.NoLower)

// Adaptor defines the interface all mocks need to implement for this
// implementation to interact properly
type Adaptor[P any, K comparable] interface {
	PrettyParams(params P) string
	ParamsKey(params P, anyParams uint64) K
}

// Moq holds the state of a mocked function
type Moq[A Adaptor[P, K], P any, K comparable, R any] struct {
	Scene  *moq.Scene
	Config moq.Config

	Adaptor         A
	ResultsByParams []ResultsByParams[P, K, R]
}

// ResultsByParams contains the results for a given set of parameters
type ResultsByParams[P any, K comparable, R any] struct {
	AnyCount  int
	AnyParams uint64
	Results   map[K]*Results[P, R]
}

// DoFn defines the type of function needed when calling AndDo
type DoFn[P any] func(params P)

// DoReturnFn defines the type of function needed when calling DoReturnResults
type DoReturnFn[P any, R any] func(params P) (results *R)

// Results holds the inner results which determine how a coll is validated
// and responded to
type Results[P any, R any] struct {
	Params  P
	Results []struct {
		Values     *R
		Sequence   uint32
		DoFn       DoFn[P]
		DoReturnFn DoReturnFn[P, R]
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// Recorder routes recorded function calls to the mock
type Recorder[A Adaptor[P, K], P any, K comparable, R any] struct {
	Scene *moq.Scene
	Moq   *Moq[A, P, K, R]

	Adaptor   A
	Params    P
	AnyParams uint64
	Sequence  bool
	Results   *Results[P, R]
}

// NewMoq creates a new mock
func NewMoq[A Adaptor[P, K], P any, K comparable, R any](
	scene *moq.Scene,
	adaptor A,
	config *moq.Config,
) *Moq[A, P, K, R] {
	if config == nil {
		config = &moq.Config{}
	}
	return &Moq[A, P, K, R]{
		Scene:  scene,
		Config: *config,

		Adaptor: adaptor,
	}
}

// Function implements a function call
func (m *Moq[A, P, K, R]) Function(params P) *R {
	m.Scene.T.Helper()
	var results *Results[P, R]
	for _, resultsByParams := range m.ResultsByParams {
		paramsKey := m.Adaptor.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Config.Expectation == moq.Strict {
			m.Scene.T.Fatalf("Unexpected call to %s", m.Adaptor.PrettyParams(params))
		}
		return nil
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Config.Expectation == moq.Strict {
				m.Scene.T.Fatalf("Too many calls to %s", m.Adaptor.PrettyParams(params))
			}
			return nil
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Scene.T.Fatalf("Call sequence does not match call to %s",
				m.Adaptor.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(params)
	}
	retVals := result.Values
	if result.DoReturnFn != nil {
		retVals = result.DoReturnFn(params)
	}

	return retVals
}

// OnCall prepares a mock to record a new expected function call
func (m *Moq[A, P, K, R]) OnCall(params P) *Recorder[A, P, K, R] {
	return &Recorder[A, P, K, R]{
		Scene: m.Scene,
		Moq:   m,

		Adaptor:  m.Adaptor,
		Params:   params,
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
	}
}

// IsAnyPermitted checks if "any params" can currently be set by checking the
// state of the recorder. It should be called by a mock before setting an any
// param.
func (r *Recorder[A, P, K, R]) IsAnyPermitted(exported bool) bool {
	r.Scene.T.Helper()
	if r.Results != nil {
		r.Scene.T.Fatalf("Any functions must be called before %s or %s calls, recording %s",
			export("returnResults", exported),
			export("doReturnResults", exported),
			r.Adaptor.PrettyParams(r.Params))
		return false
	}
	return true
}

// AnyParam records that a parameter should not be used when recording an
// expectation and that the parameter should also not be used to find results
func (r *Recorder[A, P, K, R]) AnyParam(n int) {
	r.AnyParams |= 1 << n
}

// Seq is called by Seq and NoSeq in a mock to specify whether a sequence
// value should be reserved and then checked by the mock
func (r *Recorder[A, P, K, R]) Seq(seq bool, fnName string, exported bool) bool {
	r.Scene.T.Helper()
	if r.Results != nil {
		r.Scene.T.Fatalf("%s must be called before %s or %s calls, recording %s",
			fnName,
			export("returnResults", exported),
			export("doReturnResults", exported),
			r.Adaptor.PrettyParams(r.Params))
		return false
	}
	r.Sequence = seq
	return true
}

// ReturnResults records the results to be returned by a mock
func (r *Recorder[A, P, K, R]) ReturnResults(results R) {
	r.findResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *R
		Sequence   uint32
		DoFn       DoFn[P]
		DoReturnFn DoReturnFn[P, R]
	}{
		Values:   &results,
		Sequence: sequence,
	})
}

// AndDo records a "do function" that is called in addition to returning the
// results specified by ReturnResults
func (r *Recorder[A, P, K, R]) AndDo(fn DoFn[P], exported bool) bool {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("%s must be called before calling %s",
			export("returnResults", exported),
			export("andDo", exported))
		return false
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return true
}

// DoReturnResults records a "do return function" that is called to determine
// the results to return. Each call to ReturnResults or DoReturnResults
// indicates an additional mock call.
func (r *Recorder[A, P, K, R]) DoReturnResults(fn DoReturnFn[P, R]) {
	r.findResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *R
		Sequence   uint32
		DoFn       DoFn[P]
		DoReturnFn DoReturnFn[P, R]
	}{Sequence: sequence, DoReturnFn: fn})
}

// Repeat records how many repeated calls are expected
func (r *Recorder[A, P, K, R]) Repeat(repeaters []moq.Repeater, exported bool) bool {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("%s or %s must be called before calling %s",
			export("returnResults", exported),
			export("doReturnResults", exported),
			export("repeat", exported))
		return false
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values     *R
				Sequence   uint32
				DoFn       DoFn[P]
				DoReturnFn DoReturnFn[P, R]
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}

	return true
}

// Reset resets the state of the mock
func (m *Moq[A, P, K, R]) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *Moq[A, P, K, R]) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s",
					missing, m.Adaptor.PrettyParams(results.Params))
			}
		}
	}
}

// ParamKey returns a parameter's values if anyParams determines the parameter
// is not ignored (zero values are returned if the parameter is ignored).
// iType is used to determine if a parameter is indexed by value or by hash.
// Parameters indexed by value have a value returned and a zero-value hash.
// Parameters indexed by hash have a hash value returned and a zero-value
// value.
func ParamKey[T comparable](val T, n int, iType moq.ParamIndexing, anyParams uint64) (T, hash.Hash) {
	var used T
	var usedHash hash.Hash
	if anyParams&(1<<n) == 0 {
		if iType == moq.ParamIndexByValue {
			used = val
		} else {
			usedHash = hash.DeepHash(val)
		}
	}
	return used, usedHash
}

// HashOnlyParamKey returns a parameter's hash if anyParams determines the
// parameter is not ignored (a zero value is returned if the parameter is
// ignored). If iType indicates that a parameter should be indexed by value,
// a fatal test failure is recorded and the test failed. Mocks should call
// HashOnlyParamKey for parameters that are incompatible with a parameter key.
func HashOnlyParamKey[T any](t moq.T, val T, name string, n int, iType moq.ParamIndexing, anyParams uint64) hash.Hash {
	t.Helper()
	var usedHash hash.Hash
	if anyParams&(1<<n) == 0 {
		if iType == moq.ParamIndexByValue {
			t.Fatalf("The %s parameter can't be indexed by value", name)
		}
		usedHash = hash.DeepHash(val)
	}
	return usedHash
}

func (r *Recorder[A, P, K, R]) findResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *ResultsByParams[P, K, R]
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &r.Moq.ResultsByParams[n]
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &ResultsByParams[P, K, R]{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[K]*Results[P, R]{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.Adaptor.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &Results[P, R]{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func export(fnName string, exported bool) string {
	if exported {
		return titler.String(fnName)
	}
	return fnName
}
