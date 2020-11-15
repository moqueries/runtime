// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package moq_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockMock holds the state of a mock of the Mock type
type mockMock struct {
	scene                                 *moq.Scene
	config                                moq.MockConfig
	resultsByParams_Reset                 []mockMock_Reset_resultsByParams
	resultsByParams_AssertExpectationsMet []mockMock_AssertExpectationsMet_resultsByParams
}

// mockMock_mock isolates the mock interface of the Mock type
type mockMock_mock struct {
	mock *mockMock
}

// mockMock_recorder isolates the recorder interface of the Mock type
type mockMock_recorder struct {
	mock *mockMock
}

// mockMock_Reset_params holds the params of the Mock type
type mockMock_Reset_params struct{}

// mockMock_Reset_paramsKey holds the map key params of the Mock type
type mockMock_Reset_paramsKey struct{}

// mockMock_Reset_resultsByParams contains the results for a given set of parameters for the Mock type
type mockMock_Reset_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockMock_Reset_paramsKey]*mockMock_Reset_resultMgr
}

// mockMock_Reset_resultMgr manages multiple results and the state of the Mock type
type mockMock_Reset_resultMgr struct {
	params   mockMock_Reset_params
	results  []*mockMock_Reset_results
	index    uint32
	anyTimes bool
}

// mockMock_Reset_results holds the results of the Mock type
type mockMock_Reset_results struct {
}

// mockMock_Reset_fnRecorder routes recorded function calls to the mockMock mock
type mockMock_Reset_fnRecorder struct {
	params    mockMock_Reset_params
	paramsKey mockMock_Reset_paramsKey
	anyParams uint64
	results   *mockMock_Reset_resultMgr
	mock      *mockMock
}

// mockMock_AssertExpectationsMet_params holds the params of the Mock type
type mockMock_AssertExpectationsMet_params struct{}

// mockMock_AssertExpectationsMet_paramsKey holds the map key params of the Mock type
type mockMock_AssertExpectationsMet_paramsKey struct{}

// mockMock_AssertExpectationsMet_resultsByParams contains the results for a given set of parameters for the Mock type
type mockMock_AssertExpectationsMet_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockMock_AssertExpectationsMet_paramsKey]*mockMock_AssertExpectationsMet_resultMgr
}

// mockMock_AssertExpectationsMet_resultMgr manages multiple results and the state of the Mock type
type mockMock_AssertExpectationsMet_resultMgr struct {
	params   mockMock_AssertExpectationsMet_params
	results  []*mockMock_AssertExpectationsMet_results
	index    uint32
	anyTimes bool
}

// mockMock_AssertExpectationsMet_results holds the results of the Mock type
type mockMock_AssertExpectationsMet_results struct {
}

// mockMock_AssertExpectationsMet_fnRecorder routes recorded function calls to the mockMock mock
type mockMock_AssertExpectationsMet_fnRecorder struct {
	params    mockMock_AssertExpectationsMet_params
	paramsKey mockMock_AssertExpectationsMet_paramsKey
	anyParams uint64
	results   *mockMock_AssertExpectationsMet_resultMgr
	mock      *mockMock
}

// newMockMock creates a new mock of the Mock type
func newMockMock(scene *moq.Scene, config *moq.MockConfig) *mockMock {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockMock{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the Mock type
func (m *mockMock) mock() *mockMock_mock {
	return &mockMock_mock{
		mock: m,
	}
}

func (m *mockMock_mock) Reset() {
	params := mockMock_Reset_params{}
	var results *mockMock_Reset_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams_Reset {
		paramsKey := mockMock_Reset_paramsKey{}
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

func (m *mockMock_mock) AssertExpectationsMet() {
	params := mockMock_AssertExpectationsMet_params{}
	var results *mockMock_AssertExpectationsMet_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams_AssertExpectationsMet {
		paramsKey := mockMock_AssertExpectationsMet_paramsKey{}
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

// onCall returns the recorder implementation of the Mock type
func (m *mockMock) onCall() *mockMock_recorder {
	return &mockMock_recorder{
		mock: m,
	}
}

func (m *mockMock_recorder) Reset() *mockMock_Reset_fnRecorder {
	return &mockMock_Reset_fnRecorder{
		params:    mockMock_Reset_params{},
		paramsKey: mockMock_Reset_paramsKey{},
		mock:      m.mock,
	}
}

func (r *mockMock_Reset_fnRecorder) returnResults() *mockMock_Reset_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockMock_Reset_resultsByParams
		for n, res := range r.mock.resultsByParams_Reset {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockMock_Reset_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockMock_Reset_paramsKey]*mockMock_Reset_resultMgr{},
			}
			r.mock.resultsByParams_Reset = append(r.mock.resultsByParams_Reset, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams_Reset) {
				copy(r.mock.resultsByParams_Reset[insertAt+1:], r.mock.resultsByParams_Reset[insertAt:0])
				r.mock.resultsByParams_Reset[insertAt] = *results
			}
		}

		paramsKey := mockMock_Reset_paramsKey{}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockMock_Reset_resultMgr{
			params:   r.params,
			results:  []*mockMock_Reset_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockMock_Reset_results{})
	return r
}

func (r *mockMock_Reset_fnRecorder) times(count int) *mockMock_Reset_fnRecorder {
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

func (r *mockMock_Reset_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockMock_recorder) AssertExpectationsMet() *mockMock_AssertExpectationsMet_fnRecorder {
	return &mockMock_AssertExpectationsMet_fnRecorder{
		params:    mockMock_AssertExpectationsMet_params{},
		paramsKey: mockMock_AssertExpectationsMet_paramsKey{},
		mock:      m.mock,
	}
}

func (r *mockMock_AssertExpectationsMet_fnRecorder) returnResults() *mockMock_AssertExpectationsMet_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockMock_AssertExpectationsMet_resultsByParams
		for n, res := range r.mock.resultsByParams_AssertExpectationsMet {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockMock_AssertExpectationsMet_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockMock_AssertExpectationsMet_paramsKey]*mockMock_AssertExpectationsMet_resultMgr{},
			}
			r.mock.resultsByParams_AssertExpectationsMet = append(r.mock.resultsByParams_AssertExpectationsMet, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams_AssertExpectationsMet) {
				copy(r.mock.resultsByParams_AssertExpectationsMet[insertAt+1:], r.mock.resultsByParams_AssertExpectationsMet[insertAt:0])
				r.mock.resultsByParams_AssertExpectationsMet[insertAt] = *results
			}
		}

		paramsKey := mockMock_AssertExpectationsMet_paramsKey{}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockMock_AssertExpectationsMet_resultMgr{
			params:   r.params,
			results:  []*mockMock_AssertExpectationsMet_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockMock_AssertExpectationsMet_results{})
	return r
}

func (r *mockMock_AssertExpectationsMet_fnRecorder) times(count int) *mockMock_AssertExpectationsMet_fnRecorder {
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

func (r *mockMock_AssertExpectationsMet_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockMock) Reset() {
	m.resultsByParams_Reset = nil
	m.resultsByParams_AssertExpectationsMet = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockMock) AssertExpectationsMet() {
	for _, res := range m.resultsByParams_Reset {
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
	for _, res := range m.resultsByParams_AssertExpectationsMet {
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
