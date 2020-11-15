// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package demo_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/demo"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockIsFavorite holds the state of a mock of the IsFavorite type
type mockIsFavorite struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockIsFavorite_resultsByParams
}

// mockIsFavorite_mock isolates the mock interface of the IsFavorite type
type mockIsFavorite_mock struct {
	mock *mockIsFavorite
}

// mockIsFavorite_recorder isolates the recorder interface of the IsFavorite type
type mockIsFavorite_recorder struct {
	mock *mockIsFavorite
}

// mockIsFavorite_params holds the params of the IsFavorite type
type mockIsFavorite_params struct{ n int }

// mockIsFavorite_paramsKey holds the map key params of the IsFavorite type
type mockIsFavorite_paramsKey struct{ n int }

// mockIsFavorite_resultsByParams contains the results for a given set of parameters for the IsFavorite type
type mockIsFavorite_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockIsFavorite_paramsKey]*mockIsFavorite_resultMgr
}

// mockIsFavorite_resultMgr manages multiple results and the state of the IsFavorite type
type mockIsFavorite_resultMgr struct {
	params   mockIsFavorite_params
	results  []*mockIsFavorite_results
	index    uint32
	anyTimes bool
}

// mockIsFavorite_results holds the results of the IsFavorite type
type mockIsFavorite_results struct {
	result1 bool
}

// mockIsFavorite_fnRecorder routes recorded function calls to the mockIsFavorite mock
type mockIsFavorite_fnRecorder struct {
	params    mockIsFavorite_params
	paramsKey mockIsFavorite_paramsKey
	anyParams uint64
	results   *mockIsFavorite_resultMgr
	mock      *mockIsFavorite
}

// newMockIsFavorite creates a new mock of the IsFavorite type
func newMockIsFavorite(scene *moq.Scene, config *moq.MockConfig) *mockIsFavorite {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockIsFavorite{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the IsFavorite type
func (m *mockIsFavorite) mock() demo.IsFavorite {
	return func(n int) bool { mock := &mockIsFavorite_mock{mock: m}; return mock.fn(n) }
}

func (m *mockIsFavorite_mock) fn(n int) (result1 bool) {
	params := mockIsFavorite_params{
		n: n,
	}
	var results *mockIsFavorite_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		var nUsed int
		if resultsByParams.anyParams&(1<<0) == 0 {
			nUsed = n
		}
		paramsKey := mockIsFavorite_paramsKey{
			n: nUsed,
		}
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
	result1 = result.result1
	return
}

func (m *mockIsFavorite) onCall(n int) *mockIsFavorite_fnRecorder {
	return &mockIsFavorite_fnRecorder{
		params: mockIsFavorite_params{
			n: n,
		},
		paramsKey: mockIsFavorite_paramsKey{
			n: n,
		},
		mock: m,
	}
}

func (r *mockIsFavorite_fnRecorder) anyN() *mockIsFavorite_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockIsFavorite_fnRecorder) returnResults(result1 bool) *mockIsFavorite_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockIsFavorite_resultsByParams
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
			results = &mockIsFavorite_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockIsFavorite_paramsKey]*mockIsFavorite_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		var nUsed int
		if r.anyParams&(1<<0) == 0 {
			nUsed = r.paramsKey.n
		}
		paramsKey := mockIsFavorite_paramsKey{
			n: nUsed,
		}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockIsFavorite_resultMgr{
			params:   r.params,
			results:  []*mockIsFavorite_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockIsFavorite_results{
		result1: result1,
	})
	return r
}

func (r *mockIsFavorite_fnRecorder) times(count int) *mockIsFavorite_fnRecorder {
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

func (r *mockIsFavorite_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockIsFavorite) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockIsFavorite) AssertExpectationsMet() {
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
