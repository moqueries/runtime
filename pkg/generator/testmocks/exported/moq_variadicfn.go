// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/hash"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// MockVariadicFn holds the state of a mock of the VariadicFn type
type MockVariadicFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams []MockVariadicFn_resultsByParams
}

// MockVariadicFn_mock isolates the mock interface of the VariadicFn type
type MockVariadicFn_mock struct {
	Mock *MockVariadicFn
}

// MockVariadicFn_recorder isolates the recorder interface of the VariadicFn type
type MockVariadicFn_recorder struct {
	Mock *MockVariadicFn
}

// MockVariadicFn_params holds the params of the VariadicFn type
type MockVariadicFn_params struct {
	Other bool
	Args  []string
}

// MockVariadicFn_paramsKey holds the map key params of the VariadicFn type
type MockVariadicFn_paramsKey struct {
	Other bool
	Args  hash.Hash
}

// MockVariadicFn_resultsByParams contains the results for a given set of parameters for the VariadicFn type
type MockVariadicFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MockVariadicFn_paramsKey]*MockVariadicFn_resultMgr
}

// MockVariadicFn_resultMgr manages multiple results and the state of the VariadicFn type
type MockVariadicFn_resultMgr struct {
	Params   MockVariadicFn_params
	Results  []*MockVariadicFn_results
	Index    uint32
	AnyTimes bool
}

// MockVariadicFn_results holds the results of the VariadicFn type
type MockVariadicFn_results struct {
	SResult string
	Err     error
}

// MockVariadicFn_fnRecorder routes recorded function calls to the MockVariadicFn mock
type MockVariadicFn_fnRecorder struct {
	Params    MockVariadicFn_params
	ParamsKey MockVariadicFn_paramsKey
	AnyParams uint64
	Results   *MockVariadicFn_resultMgr
	Mock      *MockVariadicFn
}

// NewMockVariadicFn creates a new mock of the VariadicFn type
func NewMockVariadicFn(scene *moq.Scene, config *moq.MockConfig) *MockVariadicFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockVariadicFn{
		Scene:  scene,
		Config: *config,
	}
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the VariadicFn type
func (m *MockVariadicFn) Mock() testmocks.VariadicFn {
	return func(other bool, args ...string) (sResult string, err error) {
		mock := &MockVariadicFn_mock{Mock: m}
		return mock.Fn(other, args...)
	}
}

func (m *MockVariadicFn_mock) Fn(other bool, args ...string) (sResult string, err error) {
	params := MockVariadicFn_params{
		Other: other,
		Args:  args,
	}
	var results *MockVariadicFn_resultMgr
	for _, resultsByParams := range m.Mock.ResultsByParams {
		var otherUsed bool
		if resultsByParams.AnyParams&(1<<0) == 0 {
			otherUsed = other
		}
		var argsUsed hash.Hash
		if resultsByParams.AnyParams&(1<<1) == 0 {
			argsUsed = hash.DeepHash(args)
		}
		paramsKey := MockVariadicFn_paramsKey{
			Other: otherUsed,
			Args:  argsUsed,
		}
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
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
	result := results.Results[i]
	sResult = result.SResult
	err = result.Err
	return
}

func (m *MockVariadicFn) OnCall(other bool, args ...string) *MockVariadicFn_fnRecorder {
	return &MockVariadicFn_fnRecorder{
		Params: MockVariadicFn_params{
			Other: other,
			Args:  args,
		},
		ParamsKey: MockVariadicFn_paramsKey{
			Other: other,
			Args:  hash.DeepHash(args),
		},
		Mock: m,
	}
}

func (r *MockVariadicFn_fnRecorder) AnyOther() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 0
	return r
}

func (r *MockVariadicFn_fnRecorder) AnyArgs() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 1
	return r
}

func (r *MockVariadicFn_fnRecorder) ReturnResults(sResult string, err error) *MockVariadicFn_fnRecorder {
	if r.Results == nil {
		anyCount := bits.OnesCount64(r.AnyParams)
		insertAt := -1
		var results *MockVariadicFn_resultsByParams
		for n, res := range r.Mock.ResultsByParams {
			if res.AnyParams == r.AnyParams {
				results = &res
				break
			}
			if res.AnyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &MockVariadicFn_resultsByParams{
				AnyCount:  anyCount,
				AnyParams: r.AnyParams,
				Results:   map[MockVariadicFn_paramsKey]*MockVariadicFn_resultMgr{},
			}
			r.Mock.ResultsByParams = append(r.Mock.ResultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.Mock.ResultsByParams) {
				copy(r.Mock.ResultsByParams[insertAt+1:], r.Mock.ResultsByParams[insertAt:0])
				r.Mock.ResultsByParams[insertAt] = *results
			}
		}

		var otherUsed bool
		if r.AnyParams&(1<<0) == 0 {
			otherUsed = r.ParamsKey.Other
		}
		var argsUsed hash.Hash
		if r.AnyParams&(1<<1) == 0 {
			argsUsed = r.ParamsKey.Args
		}
		paramsKey := MockVariadicFn_paramsKey{
			Other: otherUsed,
			Args:  argsUsed,
		}

		if _, ok := results.Results[paramsKey]; ok {
			r.Mock.Scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.Params)
			return nil
		}

		r.Results = &MockVariadicFn_resultMgr{
			Params:   r.Params,
			Results:  []*MockVariadicFn_results{},
			Index:    0,
			AnyTimes: false,
		}
		results.Results[paramsKey] = r.Results
	}
	r.Results.Results = append(r.Results.Results, &MockVariadicFn_results{
		SResult: sResult,
		Err:     err,
	})
	return r
}

func (r *MockVariadicFn_fnRecorder) Times(count int) *MockVariadicFn_fnRecorder {
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

func (r *MockVariadicFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockVariadicFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockVariadicFn) AssertExpectationsMet() {
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := len(results.Results) - int(atomic.LoadUint32(&results.Index))
			if missing == 1 && results.AnyTimes == true {
				continue
			}
			if missing > 0 {
				m.Scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.Params)
			}
		}
	}
}
