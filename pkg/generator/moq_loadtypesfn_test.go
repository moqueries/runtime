// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package generator_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/dave/dst"
	"github.com/myshkin5/moqueries/pkg/generator"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockLoadTypesFn holds the state of a mock of the LoadTypesFn type
type mockLoadTypesFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockLoadTypesFn_resultsByParams
}

// mockLoadTypesFn_mock isolates the mock interface of the LoadTypesFn type
type mockLoadTypesFn_mock struct {
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

// mockLoadTypesFn_resultsByParams contains the results for a given set of parameters for the LoadTypesFn type
type mockLoadTypesFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_resultMgr
}

// mockLoadTypesFn_resultMgr manages multiple results and the state of the LoadTypesFn type
type mockLoadTypesFn_resultMgr struct {
	params   mockLoadTypesFn_params
	results  []*mockLoadTypesFn_results
	index    uint32
	anyTimes bool
}

// mockLoadTypesFn_results holds the results of the LoadTypesFn type
type mockLoadTypesFn_results struct {
	typeSpecs    []*dst.TypeSpec
	pkgPath      string
	err          error
	moq_sequence uint32
}

// mockLoadTypesFn_fnRecorder routes recorded function calls to the mockLoadTypesFn mock
type mockLoadTypesFn_fnRecorder struct {
	params    mockLoadTypesFn_params
	paramsKey mockLoadTypesFn_paramsKey
	anyParams uint64
	sequence  bool
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
	params := mockLoadTypesFn_params{
		pkg:           pkg,
		loadTestTypes: loadTestTypes,
	}
	var results *mockLoadTypesFn_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		var pkgUsed string
		if resultsByParams.anyParams&(1<<0) == 0 {
			pkgUsed = pkg
		}
		var loadTestTypesUsed bool
		if resultsByParams.anyParams&(1<<1) == 0 {
			loadTestTypesUsed = loadTestTypes
		}
		paramsKey := mockLoadTypesFn_paramsKey{
			pkg:           pkgUsed,
			loadTestTypes: loadTestTypesUsed,
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
	if result.moq_sequence != 0 {
		sequence := m.mock.scene.NextMockSequence()
		if (!results.anyTimes && result.moq_sequence != sequence) || result.moq_sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

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
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		mock:     m,
	}
}

func (r *mockLoadTypesFn_fnRecorder) anyPkg() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockLoadTypesFn_fnRecorder) anyLoadTestTypes() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockLoadTypesFn_fnRecorder) seq() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockLoadTypesFn_fnRecorder) noSeq() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockLoadTypesFn_fnRecorder) returnResults(
	typeSpecs []*dst.TypeSpec, pkgPath string, err error) *mockLoadTypesFn_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockLoadTypesFn_resultsByParams
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
			results = &mockLoadTypesFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		var pkgUsed string
		if r.anyParams&(1<<0) == 0 {
			pkgUsed = r.paramsKey.pkg
		}
		var loadTestTypesUsed bool
		if r.anyParams&(1<<1) == 0 {
			loadTestTypesUsed = r.paramsKey.loadTestTypes
		}
		paramsKey := mockLoadTypesFn_paramsKey{
			pkg:           pkgUsed,
			loadTestTypes: loadTestTypesUsed,
		}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockLoadTypesFn_resultMgr{
			params:   r.params,
			results:  []*mockLoadTypesFn_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, &mockLoadTypesFn_results{
		typeSpecs:    typeSpecs,
		pkgPath:      pkgPath,
		err:          err,
		moq_sequence: sequence,
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
		if last.moq_sequence != 0 {
			last = &mockLoadTypesFn_results{
				typeSpecs:    last.typeSpecs,
				pkgPath:      last.pkgPath,
				err:          last.err,
				moq_sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
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
func (m *mockLoadTypesFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockLoadTypesFn) AssertExpectationsMet() {
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
