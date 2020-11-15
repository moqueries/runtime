// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockNothingFn holds the state of a mock of the NothingFn type
type mockNothingFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockNothingFn_resultsByParams
}

// mockNothingFn_mock isolates the mock interface of the NothingFn type
type mockNothingFn_mock struct {
	mock *mockNothingFn
}

// mockNothingFn_recorder isolates the recorder interface of the NothingFn type
type mockNothingFn_recorder struct {
	mock *mockNothingFn
}

// mockNothingFn_params holds the params of the NothingFn type
type mockNothingFn_params struct{}

// mockNothingFn_paramsKey holds the map key params of the NothingFn type
type mockNothingFn_paramsKey struct{}

// mockNothingFn_resultsByParams contains the results for a given set of parameters for the NothingFn type
type mockNothingFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockNothingFn_paramsKey]*mockNothingFn_resultMgr
}

// mockNothingFn_resultMgr manages multiple results and the state of the NothingFn type
type mockNothingFn_resultMgr struct {
	params   mockNothingFn_params
	results  []*mockNothingFn_results
	index    uint32
	anyTimes bool
}

// mockNothingFn_results holds the results of the NothingFn type
type mockNothingFn_results struct {
}

// mockNothingFn_fnRecorder routes recorded function calls to the mockNothingFn mock
type mockNothingFn_fnRecorder struct {
	params    mockNothingFn_params
	paramsKey mockNothingFn_paramsKey
	anyParams uint64
	results   *mockNothingFn_resultMgr
	mock      *mockNothingFn
}

// newMockNothingFn creates a new mock of the NothingFn type
func newMockNothingFn(scene *moq.Scene, config *moq.MockConfig) *mockNothingFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockNothingFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the NothingFn type
func (m *mockNothingFn) mock() testmocks.NothingFn {
	return func() { mock := &mockNothingFn_mock{mock: m}; mock.fn() }
}

func (m *mockNothingFn_mock) fn() {
	params := mockNothingFn_params{}
	var results *mockNothingFn_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		paramsKey := mockNothingFn_paramsKey{}
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
	return
}

func (m *mockNothingFn) onCall() *mockNothingFn_fnRecorder {
	return &mockNothingFn_fnRecorder{
		params:    mockNothingFn_params{},
		paramsKey: mockNothingFn_paramsKey{},
		mock:      m,
	}
}

func (r *mockNothingFn_fnRecorder) returnResults() *mockNothingFn_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockNothingFn_resultsByParams
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
			results = &mockNothingFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockNothingFn_paramsKey]*mockNothingFn_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		paramsKey := mockNothingFn_paramsKey{}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockNothingFn_resultMgr{
			params:   r.params,
			results:  []*mockNothingFn_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockNothingFn_results{})
	return r
}

func (r *mockNothingFn_fnRecorder) times(count int) *mockNothingFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockNothingFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockNothingFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockNothingFn) AssertExpectationsMet() {
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
