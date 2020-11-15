// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/hash"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockVariadicFn holds the state of a mock of the VariadicFn type
type mockVariadicFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockVariadicFn_resultsByParams
}

// mockVariadicFn_mock isolates the mock interface of the VariadicFn type
type mockVariadicFn_mock struct {
	mock *mockVariadicFn
}

// mockVariadicFn_recorder isolates the recorder interface of the VariadicFn type
type mockVariadicFn_recorder struct {
	mock *mockVariadicFn
}

// mockVariadicFn_params holds the params of the VariadicFn type
type mockVariadicFn_params struct {
	other bool
	args  []string
}

// mockVariadicFn_paramsKey holds the map key params of the VariadicFn type
type mockVariadicFn_paramsKey struct {
	other bool
	args  hash.Hash
}

// mockVariadicFn_resultsByParams contains the results for a given set of parameters for the VariadicFn type
type mockVariadicFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockVariadicFn_paramsKey]*mockVariadicFn_resultMgr
}

// mockVariadicFn_resultMgr manages multiple results and the state of the VariadicFn type
type mockVariadicFn_resultMgr struct {
	params   mockVariadicFn_params
	results  []*mockVariadicFn_results
	index    uint32
	anyTimes bool
}

// mockVariadicFn_results holds the results of the VariadicFn type
type mockVariadicFn_results struct {
	sResult string
	err     error
}

// mockVariadicFn_fnRecorder routes recorded function calls to the mockVariadicFn mock
type mockVariadicFn_fnRecorder struct {
	params    mockVariadicFn_params
	paramsKey mockVariadicFn_paramsKey
	anyParams uint64
	results   *mockVariadicFn_resultMgr
	mock      *mockVariadicFn
}

// newMockVariadicFn creates a new mock of the VariadicFn type
func newMockVariadicFn(scene *moq.Scene, config *moq.MockConfig) *mockVariadicFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockVariadicFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the VariadicFn type
func (m *mockVariadicFn) mock() testmocks.VariadicFn {
	return func(other bool, args ...string) (sResult string, err error) {
		mock := &mockVariadicFn_mock{mock: m}
		return mock.fn(other, args...)
	}
}

func (m *mockVariadicFn_mock) fn(other bool, args ...string) (sResult string, err error) {
	params := mockVariadicFn_params{
		other: other,
		args:  args,
	}
	var results *mockVariadicFn_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		var otherUsed bool
		if resultsByParams.anyParams&(1<<0) == 0 {
			otherUsed = other
		}
		var argsUsed hash.Hash
		if resultsByParams.anyParams&(1<<1) == 0 {
			argsUsed = hash.DeepHash(args)
		}
		paramsKey := mockVariadicFn_paramsKey{
			other: otherUsed,
			args:  argsUsed,
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
	sResult = result.sResult
	err = result.err
	return
}

func (m *mockVariadicFn) onCall(other bool, args ...string) *mockVariadicFn_fnRecorder {
	return &mockVariadicFn_fnRecorder{
		params: mockVariadicFn_params{
			other: other,
			args:  args,
		},
		paramsKey: mockVariadicFn_paramsKey{
			other: other,
			args:  hash.DeepHash(args),
		},
		mock: m,
	}
}

func (r *mockVariadicFn_fnRecorder) anyOther() *mockVariadicFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockVariadicFn_fnRecorder) anyArgs() *mockVariadicFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockVariadicFn_fnRecorder) returnResults(sResult string, err error) *mockVariadicFn_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockVariadicFn_resultsByParams
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
			results = &mockVariadicFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockVariadicFn_paramsKey]*mockVariadicFn_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		var otherUsed bool
		if r.anyParams&(1<<0) == 0 {
			otherUsed = r.paramsKey.other
		}
		var argsUsed hash.Hash
		if r.anyParams&(1<<1) == 0 {
			argsUsed = r.paramsKey.args
		}
		paramsKey := mockVariadicFn_paramsKey{
			other: otherUsed,
			args:  argsUsed,
		}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockVariadicFn_resultMgr{
			params:   r.params,
			results:  []*mockVariadicFn_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockVariadicFn_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockVariadicFn_fnRecorder) times(count int) *mockVariadicFn_fnRecorder {
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

func (r *mockVariadicFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockVariadicFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockVariadicFn) AssertExpectationsMet() {
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
