// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package generator_test

import (
	"sync/atomic"

	"github.com/dave/dst"
	"github.com/myshkin5/moqueries/pkg/generator"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockLoadTypesFn holds the state of a mock of the LoadTypesFn type
type mockLoadTypesFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_resultMgr
}

// mockLoadTypesFn_mock isolates the mock interface of the LoadTypesFn type
type mockLoadTypesFn_mock struct {
	mock *mockLoadTypesFn
}

// mockLoadTypesFn_recorder isolates the recorder interface of the LoadTypesFn type
type mockLoadTypesFn_recorder struct {
	mock *mockLoadTypesFn
}

// mockLoadTypesFn_params holds the params of the LoadTypesFn type
type mockLoadTypesFn_params struct {
	pkg           string
	loadTestTypes bool
}

// mockLoadTypesFn_paramsKey holds the map key params of the LoadTypesFn type
type mockLoadTypesFn_paramsKey struct {
	pkg           string
	loadTestTypes bool
}

// mockLoadTypesFn_resultMgr manages multiple results and the state of the LoadTypesFn type
type mockLoadTypesFn_resultMgr struct {
	results  []*mockLoadTypesFn_results
	index    uint32
	anyTimes bool
}

// mockLoadTypesFn_results holds the results of the LoadTypesFn type
type mockLoadTypesFn_results struct {
	typeSpecs []*dst.TypeSpec
	pkgPath   string
	err       error
}

// mockLoadTypesFn_fnRecorder routes recorded function calls to the mockLoadTypesFn mock
type mockLoadTypesFn_fnRecorder struct {
	params    mockLoadTypesFn_params
	paramsKey mockLoadTypesFn_paramsKey
	results   *mockLoadTypesFn_resultMgr
	mock      *mockLoadTypesFn
}

// newMockLoadTypesFn creates a new mock of the LoadTypesFn type
func newMockLoadTypesFn(scene *moq.Scene, config *moq.MockConfig) *mockLoadTypesFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockLoadTypesFn{
		scene:  scene,
		config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the LoadTypesFn type
func (m *mockLoadTypesFn) mock() generator.LoadTypesFn {
	return func(pkg string, loadTestTypes bool) (
		typeSpecs []*dst.TypeSpec, pkgPath string, err error) {
		mock := &mockLoadTypesFn_mock{mock: m}
		return mock.fn(pkg, loadTestTypes)
	}
}

func (m *mockLoadTypesFn_mock) fn(pkg string, loadTestTypes bool) (
	typeSpecs []*dst.TypeSpec, pkgPath string, err error) {
	params := mockLoadTypesFn_paramsKey{
		pkg:           pkg,
		loadTestTypes: loadTestTypes,
	}
	results, ok := m.mock.resultsByParams[params]
	if !ok {
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
	typeSpecs = result.typeSpecs
	pkgPath = result.pkgPath
	err = result.err
	return
}

func (m *mockLoadTypesFn) onCall(pkg string, loadTestTypes bool) *mockLoadTypesFn_fnRecorder {
	return &mockLoadTypesFn_fnRecorder{
		params: mockLoadTypesFn_params{
			pkg:           pkg,
			loadTestTypes: loadTestTypes,
		},
		paramsKey: mockLoadTypesFn_paramsKey{
			pkg:           pkg,
			loadTestTypes: loadTestTypes,
		},
		mock: m,
	}
}

func (r *mockLoadTypesFn_fnRecorder) returnResults(
	typeSpecs []*dst.TypeSpec, pkgPath string, err error) *mockLoadTypesFn_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams[r.paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.paramsKey)
			return nil
		}

		r.results = &mockLoadTypesFn_resultMgr{results: []*mockLoadTypesFn_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams[r.paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockLoadTypesFn_results{
		typeSpecs: typeSpecs,
		pkgPath:   pkgPath,
		err:       err,
	})
	return r
}

func (r *mockLoadTypesFn_fnRecorder) times(count int) *mockLoadTypesFn_fnRecorder {
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

func (r *mockLoadTypesFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockLoadTypesFn) Reset() {
	m.resultsByParams = map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_resultMgr{}
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockLoadTypesFn) AssertExpectationsMet() {
	for params, results := range m.resultsByParams {
		missing := len(results.results) - int(atomic.LoadUint32(&results.index))
		if missing == 1 && results.anyTimes == true {
			continue
		}
		if missing > 0 {
			m.scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, params)
		}
	}
}
