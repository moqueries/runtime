// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmocks"
	"github.com/myshkin5/moqueries/moq"
)

// mockNoParamsFn holds the state of a mock of the NoParamsFn type
type mockNoParamsFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockNoParamsFn_resultsByParams
}

// mockNoParamsFn_mock isolates the mock interface of the NoParamsFn type
type mockNoParamsFn_mock struct {
	mock *mockNoParamsFn
}

// mockNoParamsFn_params holds the params of the NoParamsFn type
type mockNoParamsFn_params struct{}

// mockNoParamsFn_paramsKey holds the map key params of the NoParamsFn type
type mockNoParamsFn_paramsKey struct{}

// mockNoParamsFn_resultsByParams contains the results for a given set of parameters for the NoParamsFn type
type mockNoParamsFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockNoParamsFn_paramsKey]*mockNoParamsFn_results
}

// mockNoParamsFn_doFn defines the type of function needed when calling andDo for the NoParamsFn type
type mockNoParamsFn_doFn func()

// mockNoParamsFn_doReturnFn defines the type of function needed when calling doReturnResults for the NoParamsFn type
type mockNoParamsFn_doReturnFn func() (sResult string, err error)

// mockNoParamsFn_results holds the results of the NoParamsFn type
type mockNoParamsFn_results struct {
	params  mockNoParamsFn_params
	results []struct {
		values *struct {
			sResult string
			err     error
		}
		sequence   uint32
		doFn       mockNoParamsFn_doFn
		doReturnFn mockNoParamsFn_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// mockNoParamsFn_fnRecorder routes recorded function calls to the mockNoParamsFn mock
type mockNoParamsFn_fnRecorder struct {
	params    mockNoParamsFn_params
	paramsKey mockNoParamsFn_paramsKey
	anyParams uint64
	sequence  bool
	results   *mockNoParamsFn_results
	mock      *mockNoParamsFn
}

// newMockNoParamsFn creates a new mock of the NoParamsFn type
func newMockNoParamsFn(scene *moq.Scene, config *moq.MockConfig) *mockNoParamsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockNoParamsFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the NoParamsFn type
func (m *mockNoParamsFn) mock() testmocks.NoParamsFn {
	return func() (sResult string, err error) { mock := &mockNoParamsFn_mock{mock: m}; return mock.fn() }
}

func (m *mockNoParamsFn_mock) fn() (sResult string, err error) {
	params := mockNoParamsFn_params{}
	var results *mockNoParamsFn_results
	for _, resultsByParams := range m.mock.resultsByParams {
		paramsKey := mockNoParamsFn_paramsKey{}
		var ok bool
		results, ok = resultsByParams.results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.mock.config.Expectation == moq.Strict {
			m.mock.scene.MoqT.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.index, 1)) - 1
	if i >= len(results.results) {
		if !results.anyTimes {
			if m.mock.config.Expectation == moq.Strict {
				m.mock.scene.MoqT.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.results) - 1
	}

	result := results.results[i]
	if result.sequence != 0 {
		sequence := m.mock.scene.NextMockSequence()
		if (!results.anyTimes && result.sequence != sequence) || result.sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.doFn != nil {
		result.doFn()
	}

	if result.values != nil {
		sResult = result.values.sResult
		err = result.values.err
	}
	if result.doReturnFn != nil {
		sResult, err = result.doReturnFn()
	}
	return
}

func (m *mockNoParamsFn) onCall() *mockNoParamsFn_fnRecorder {
	return &mockNoParamsFn_fnRecorder{
		params:    mockNoParamsFn_params{},
		paramsKey: mockNoParamsFn_paramsKey{},
		sequence:  m.config.Sequence == moq.SeqDefaultOn,
		mock:      m,
	}
}

func (r *mockNoParamsFn_fnRecorder) seq() *mockNoParamsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockNoParamsFn_fnRecorder) noSeq() *mockNoParamsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockNoParamsFn_fnRecorder) returnResults(sResult string, err error) *mockNoParamsFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			sResult string
			err     error
		}
		sequence   uint32
		doFn       mockNoParamsFn_doFn
		doReturnFn mockNoParamsFn_doReturnFn
	}{
		values: &struct {
			sResult string
			err     error
		}{
			sResult: sResult,
			err:     err,
		},
		sequence: sequence,
	})
	return r
}

func (r *mockNoParamsFn_fnRecorder) andDo(fn mockNoParamsFn_doFn) *mockNoParamsFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *mockNoParamsFn_fnRecorder) doReturnResults(fn mockNoParamsFn_doReturnFn) *mockNoParamsFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			sResult string
			err     error
		}
		sequence   uint32
		doFn       mockNoParamsFn_doFn
		doReturnFn mockNoParamsFn_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *mockNoParamsFn_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockNoParamsFn_resultsByParams
		for n, res := range r.mock.resultsByParams {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockNoParamsFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockNoParamsFn_paramsKey]*mockNoParamsFn_results{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		paramsKey := mockNoParamsFn_paramsKey{}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &mockNoParamsFn_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *mockNoParamsFn_fnRecorder) times(count int) *mockNoParamsFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					sResult string
					err     error
				}
				sequence   uint32
				doFn       mockNoParamsFn_doFn
				doReturnFn mockNoParamsFn_doReturnFn
			}{
				values: &struct {
					sResult string
					err     error
				}{
					sResult: last.values.sResult,
					err:     last.values.err,
				},
				sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockNoParamsFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockNoParamsFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockNoParamsFn) AssertExpectationsMet() {
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := len(results.results) - int(atomic.LoadUint32(&results.index))
			if missing == 1 && results.anyTimes == true {
				continue
			}
			if missing > 0 {
				m.scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}
