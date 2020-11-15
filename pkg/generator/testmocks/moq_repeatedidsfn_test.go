// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockRepeatedIdsFn holds the state of a mock of the RepeatedIdsFn type
type mockRepeatedIdsFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockRepeatedIdsFn_resultsByParams
}

// mockRepeatedIdsFn_mock isolates the mock interface of the RepeatedIdsFn type
type mockRepeatedIdsFn_mock struct {
	mock *mockRepeatedIdsFn
}

// mockRepeatedIdsFn_recorder isolates the recorder interface of the RepeatedIdsFn type
type mockRepeatedIdsFn_recorder struct {
	mock *mockRepeatedIdsFn
}

// mockRepeatedIdsFn_params holds the params of the RepeatedIdsFn type
type mockRepeatedIdsFn_params struct {
	sParam1, sParam2 string
	bParam           bool
}

// mockRepeatedIdsFn_paramsKey holds the map key params of the RepeatedIdsFn type
type mockRepeatedIdsFn_paramsKey struct {
	sParam1, sParam2 string
	bParam           bool
}

// mockRepeatedIdsFn_resultsByParams contains the results for a given set of parameters for the RepeatedIdsFn type
type mockRepeatedIdsFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockRepeatedIdsFn_paramsKey]*mockRepeatedIdsFn_resultMgr
}

// mockRepeatedIdsFn_resultMgr manages multiple results and the state of the RepeatedIdsFn type
type mockRepeatedIdsFn_resultMgr struct {
	params   mockRepeatedIdsFn_params
	results  []*mockRepeatedIdsFn_results
	index    uint32
	anyTimes bool
}

// mockRepeatedIdsFn_results holds the results of the RepeatedIdsFn type
type mockRepeatedIdsFn_results struct {
	sResult1, sResult2 string
	err                error
}

// mockRepeatedIdsFn_fnRecorder routes recorded function calls to the mockRepeatedIdsFn mock
type mockRepeatedIdsFn_fnRecorder struct {
	params    mockRepeatedIdsFn_params
	paramsKey mockRepeatedIdsFn_paramsKey
	anyParams uint64
	results   *mockRepeatedIdsFn_resultMgr
	mock      *mockRepeatedIdsFn
}

// newMockRepeatedIdsFn creates a new mock of the RepeatedIdsFn type
func newMockRepeatedIdsFn(scene *moq.Scene, config *moq.MockConfig) *mockRepeatedIdsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockRepeatedIdsFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the RepeatedIdsFn type
func (m *mockRepeatedIdsFn) mock() testmocks.RepeatedIdsFn {
	return func(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
		mock := &mockRepeatedIdsFn_mock{mock: m}
		return mock.fn(sParam1, sParam2, bParam)
	}
}

func (m *mockRepeatedIdsFn_mock) fn(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
	params := mockRepeatedIdsFn_params{
		sParam1: sParam1,
		sParam2: sParam2,
		bParam:  bParam,
	}
	var results *mockRepeatedIdsFn_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		var sParam1Used string
		if resultsByParams.anyParams&(1<<0) == 0 {
			sParam1Used = sParam1
		}
		var sParam2Used string
		if resultsByParams.anyParams&(1<<1) == 0 {
			sParam2Used = sParam2
		}
		var bParamUsed bool
		if resultsByParams.anyParams&(1<<2) == 0 {
			bParamUsed = bParam
		}
		paramsKey := mockRepeatedIdsFn_paramsKey{
			sParam1: sParam1Used,
			sParam2: sParam2Used,
			bParam:  bParamUsed,
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
	sResult1 = result.sResult1
	sResult2 = result.sResult2
	err = result.err
	return
}

func (m *mockRepeatedIdsFn) onCall(sParam1, sParam2 string, bParam bool) *mockRepeatedIdsFn_fnRecorder {
	return &mockRepeatedIdsFn_fnRecorder{
		params: mockRepeatedIdsFn_params{
			sParam1: sParam1,
			sParam2: sParam2,
			bParam:  bParam,
		},
		paramsKey: mockRepeatedIdsFn_paramsKey{
			sParam1: sParam1,
			sParam2: sParam2,
			bParam:  bParam,
		},
		mock: m,
	}
}

func (r *mockRepeatedIdsFn_fnRecorder) anySParam1() *mockRepeatedIdsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockRepeatedIdsFn_fnRecorder) anySParam2() *mockRepeatedIdsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockRepeatedIdsFn_fnRecorder) anyBParam() *mockRepeatedIdsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 2
	return r
}

func (r *mockRepeatedIdsFn_fnRecorder) returnResults(sResult1, sResult2 string, err error) *mockRepeatedIdsFn_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockRepeatedIdsFn_resultsByParams
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
			results = &mockRepeatedIdsFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockRepeatedIdsFn_paramsKey]*mockRepeatedIdsFn_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		var sParam1Used string
		if r.anyParams&(1<<0) == 0 {
			sParam1Used = r.paramsKey.sParam1
		}
		var sParam2Used string
		if r.anyParams&(1<<1) == 0 {
			sParam2Used = r.paramsKey.sParam2
		}
		var bParamUsed bool
		if r.anyParams&(1<<2) == 0 {
			bParamUsed = r.paramsKey.bParam
		}
		paramsKey := mockRepeatedIdsFn_paramsKey{
			sParam1: sParam1Used,
			sParam2: sParam2Used,
			bParam:  bParamUsed,
		}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockRepeatedIdsFn_resultMgr{
			params:   r.params,
			results:  []*mockRepeatedIdsFn_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockRepeatedIdsFn_results{
		sResult1: sResult1,
		sResult2: sResult2,
		err:      err,
	})
	return r
}

func (r *mockRepeatedIdsFn_fnRecorder) times(count int) *mockRepeatedIdsFn_fnRecorder {
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

func (r *mockRepeatedIdsFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockRepeatedIdsFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockRepeatedIdsFn) AssertExpectationsMet() {
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
