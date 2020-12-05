// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package generator_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/dave/dst"
	"github.com/myshkin5/moqueries/generator"
	"github.com/myshkin5/moqueries/moq"
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
	results   map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_results
}

// mockLoadTypesFn_doFn defines the type of function needed when calling andDo for the LoadTypesFn type
type mockLoadTypesFn_doFn func(pkg string, loadTestTypes bool)

// mockLoadTypesFn_doReturnFn defines the type of function needed when calling doReturnResults for the LoadTypesFn type
type mockLoadTypesFn_doReturnFn func(pkg string, loadTestTypes bool) (
	typeSpecs []*dst.TypeSpec, pkgPath string, err error)

// mockLoadTypesFn_results holds the results of the LoadTypesFn type
type mockLoadTypesFn_results struct {
	params  mockLoadTypesFn_params
	results []struct {
		values *struct {
			typeSpecs []*dst.TypeSpec
			pkgPath   string
			err       error
		}
		sequence   uint32
		doFn       mockLoadTypesFn_doFn
		doReturnFn mockLoadTypesFn_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// mockLoadTypesFn_fnRecorder routes recorded function calls to the mockLoadTypesFn mock
type mockLoadTypesFn_fnRecorder struct {
	params    mockLoadTypesFn_params
	paramsKey mockLoadTypesFn_paramsKey
	anyParams uint64
	sequence  bool
	results   *mockLoadTypesFn_results
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
	var results *mockLoadTypesFn_results
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
	if result.sequence != 0 {
		sequence := m.mock.scene.NextMockSequence()
		if (!results.anyTimes && result.sequence != sequence) || result.sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.doFn != nil {
		result.doFn(pkg, loadTestTypes)
	}

	if result.values != nil {
		typeSpecs = result.values.typeSpecs
		pkgPath = result.values.pkgPath
		err = result.values.err
	}
	if result.doReturnFn != nil {
		typeSpecs, pkgPath, err = result.doReturnFn(pkg, loadTestTypes)
	}
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
		r.mock.scene.MoqT.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockLoadTypesFn_fnRecorder) anyLoadTestTypes() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockLoadTypesFn_fnRecorder) seq() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockLoadTypesFn_fnRecorder) noSeq() *mockLoadTypesFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockLoadTypesFn_fnRecorder) returnResults(
	typeSpecs []*dst.TypeSpec, pkgPath string, err error) *mockLoadTypesFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			typeSpecs []*dst.TypeSpec
			pkgPath   string
			err       error
		}
		sequence   uint32
		doFn       mockLoadTypesFn_doFn
		doReturnFn mockLoadTypesFn_doReturnFn
	}{
		values: &struct {
			typeSpecs []*dst.TypeSpec
			pkgPath   string
			err       error
		}{
			typeSpecs: typeSpecs,
			pkgPath:   pkgPath,
			err:       err,
		},
		sequence: sequence,
	})
	return r
}

func (r *mockLoadTypesFn_fnRecorder) andDo(fn mockLoadTypesFn_doFn) *mockLoadTypesFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *mockLoadTypesFn_fnRecorder) doReturnResults(fn mockLoadTypesFn_doReturnFn) *mockLoadTypesFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			typeSpecs []*dst.TypeSpec
			pkgPath   string
			err       error
		}
		sequence   uint32
		doFn       mockLoadTypesFn_doFn
		doReturnFn mockLoadTypesFn_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *mockLoadTypesFn_fnRecorder) findResults() {
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
				results:   map[mockLoadTypesFn_paramsKey]*mockLoadTypesFn_results{},
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

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &mockLoadTypesFn_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *mockLoadTypesFn_fnRecorder) times(count int) *mockLoadTypesFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					typeSpecs []*dst.TypeSpec
					pkgPath   string
					err       error
				}
				sequence   uint32
				doFn       mockLoadTypesFn_doFn
				doReturnFn mockLoadTypesFn_doReturnFn
			}{
				values: &struct {
					typeSpecs []*dst.TypeSpec
					pkgPath   string
					err       error
				}{
					typeSpecs: last.values.typeSpecs,
					pkgPath:   last.values.pkgPath,
					err:       last.values.err,
				},
				sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockLoadTypesFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
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
