// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// MockNoResultsFn holds the state of a mock of the NoResultsFn type
type MockNoResultsFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams map[MockNoResultsFn_paramsKey]*MockNoResultsFn_resultMgr
}

// MockNoResultsFn_mock isolates the mock interface of the NoResultsFn type
type MockNoResultsFn_mock struct {
	Mock *MockNoResultsFn
}

// MockNoResultsFn_recorder isolates the recorder interface of the NoResultsFn type
type MockNoResultsFn_recorder struct {
	Mock *MockNoResultsFn
}

// MockNoResultsFn_params holds the params of the NoResultsFn type
type MockNoResultsFn_params struct {
	SParam string
	BParam bool
}

// MockNoResultsFn_paramsKey holds the map key params of the NoResultsFn type
type MockNoResultsFn_paramsKey struct {
	SParam string
	BParam bool
}

// MockNoResultsFn_resultMgr manages multiple results and the state of the NoResultsFn type
type MockNoResultsFn_resultMgr struct {
	Results  []*MockNoResultsFn_results
	Index    uint32
	AnyTimes bool
}

// MockNoResultsFn_results holds the results of the NoResultsFn type
type MockNoResultsFn_results struct {
}

// MockNoResultsFn_fnRecorder routes recorded function calls to the MockNoResultsFn mock
type MockNoResultsFn_fnRecorder struct {
	Params    MockNoResultsFn_params
	ParamsKey MockNoResultsFn_paramsKey
	Results   *MockNoResultsFn_resultMgr
	Mock      *MockNoResultsFn
}

// NewMockNoResultsFn creates a new mock of the NoResultsFn type
func NewMockNoResultsFn(scene *moq.Scene, config *moq.MockConfig) *MockNoResultsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockNoResultsFn{
		Scene:  scene,
		Config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the NoResultsFn type
func (m *MockNoResultsFn) Mock() testmocks.NoResultsFn {
	return func(sParam string, bParam bool) { mock := &MockNoResultsFn_mock{Mock: m}; mock.Fn(sParam, bParam) }
}

func (m *MockNoResultsFn_mock) Fn(sParam string, bParam bool) {
	params := MockNoResultsFn_paramsKey{
		SParam: sParam,
		BParam: bParam,
	}
	results, ok := m.Mock.ResultsByParams[params]
	if !ok {
		if m.Mock.Config.Expectation == moq.Strict {
			m.Mock.Scene.MoqT.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= len(results.Results) {
		if !results.AnyTimes {
			if m.Mock.Config.Expectation == moq.Strict {
				m.Mock.Scene.MoqT.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.Results) - 1
	}
	return
}

func (m *MockNoResultsFn) OnCall(sParam string, bParam bool) *MockNoResultsFn_fnRecorder {
	return &MockNoResultsFn_fnRecorder{
		Params: MockNoResultsFn_params{
			SParam: sParam,
			BParam: bParam,
		},
		ParamsKey: MockNoResultsFn_paramsKey{
			SParam: sParam,
			BParam: bParam,
		},
		Mock: m,
	}
}

func (r *MockNoResultsFn_fnRecorder) ReturnResults() *MockNoResultsFn_fnRecorder {
	if r.Results == nil {
		if _, ok := r.Mock.ResultsByParams[r.ParamsKey]; ok {
			r.Mock.Scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.ParamsKey)
			return nil
		}

		r.Results = &MockNoResultsFn_resultMgr{Results: []*MockNoResultsFn_results{}, Index: 0, AnyTimes: false}
		r.Mock.ResultsByParams[r.ParamsKey] = r.Results
	}
	r.Results.Results = append(r.Results.Results, &MockNoResultsFn_results{})
	return r
}

func (r *MockNoResultsFn_fnRecorder) Times(count int) *MockNoResultsFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < count-1; n++ {
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (r *MockNoResultsFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockNoResultsFn) Reset() {
	m.ResultsByParams = map[MockNoResultsFn_paramsKey]*MockNoResultsFn_resultMgr{}
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockNoResultsFn) AssertExpectationsMet() {
	for params, results := range m.ResultsByParams {
		missing := len(results.Results) - int(atomic.LoadUint32(&results.Index))
		if missing == 1 && results.AnyTimes == true {
			continue
		}
		if missing > 0 {
			m.Scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, params)
		}
	}
}
