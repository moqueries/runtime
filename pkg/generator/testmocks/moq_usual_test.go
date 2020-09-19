// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/hash"
	"github.com/myshkin5/moqueries/pkg/testing"
)

// mockUsual holds the state of a mock of the Usual type
type mockUsual struct {
	t                           testing.MoqT
	resultsByParams_Usual       map[mockUsual_Usual_params]*mockUsual_Usual_resultMgr
	params_Usual                chan mockUsual_Usual_params
	resultsByParams_NoNames     map[mockUsual_NoNames_params]*mockUsual_NoNames_resultMgr
	params_NoNames              chan mockUsual_NoNames_params
	resultsByParams_NoResults   map[mockUsual_NoResults_params]*mockUsual_NoResults_resultMgr
	params_NoResults            chan mockUsual_NoResults_params
	resultsByParams_NoParams    map[mockUsual_NoParams_params]*mockUsual_NoParams_resultMgr
	params_NoParams             chan mockUsual_NoParams_params
	resultsByParams_Nothing     map[mockUsual_Nothing_params]*mockUsual_Nothing_resultMgr
	params_Nothing              chan mockUsual_Nothing_params
	resultsByParams_Variadic    map[mockUsual_Variadic_params]*mockUsual_Variadic_resultMgr
	params_Variadic             chan mockUsual_Variadic_params
	resultsByParams_RepeatedIds map[mockUsual_RepeatedIds_params]*mockUsual_RepeatedIds_resultMgr
	params_RepeatedIds          chan mockUsual_RepeatedIds_params
}

// mockUsual_mock isolates the mock interface of the Usual type
type mockUsual_mock struct {
	mock *mockUsual
}

// mockUsual_recorder isolates the recorder interface of the Usual type
type mockUsual_recorder struct {
	mock *mockUsual
}

// mockUsual_Usual_params holds the params of the Usual type
type mockUsual_Usual_params struct {
	sParam string
	bParam bool
}

// mockUsual_Usual_resultMgr manages multiple results and the state of the Usual type
type mockUsual_Usual_resultMgr struct {
	results  []*mockUsual_Usual_results
	index    uint32
	anyTimes bool
}

// mockUsual_Usual_results holds the results of the Usual type
type mockUsual_Usual_results struct {
	sResult string
	err     error
}

// mockUsual_Usual_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_Usual_fnRecorder struct {
	params  mockUsual_Usual_params
	results *mockUsual_Usual_resultMgr
	mock    *mockUsual
}

// mockUsual_NoNames_params holds the params of the Usual type
type mockUsual_NoNames_params struct {
	param1 string
	param2 bool
}

// mockUsual_NoNames_resultMgr manages multiple results and the state of the Usual type
type mockUsual_NoNames_resultMgr struct {
	results  []*mockUsual_NoNames_results
	index    uint32
	anyTimes bool
}

// mockUsual_NoNames_results holds the results of the Usual type
type mockUsual_NoNames_results struct {
	result1 string
	result2 error
}

// mockUsual_NoNames_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_NoNames_fnRecorder struct {
	params  mockUsual_NoNames_params
	results *mockUsual_NoNames_resultMgr
	mock    *mockUsual
}

// mockUsual_NoResults_params holds the params of the Usual type
type mockUsual_NoResults_params struct {
	sParam string
	bParam bool
}

// mockUsual_NoResults_resultMgr manages multiple results and the state of the Usual type
type mockUsual_NoResults_resultMgr struct {
	results  []*mockUsual_NoResults_results
	index    uint32
	anyTimes bool
}

// mockUsual_NoResults_results holds the results of the Usual type
type mockUsual_NoResults_results struct {
}

// mockUsual_NoResults_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_NoResults_fnRecorder struct {
	params  mockUsual_NoResults_params
	results *mockUsual_NoResults_resultMgr
	mock    *mockUsual
}

// mockUsual_NoParams_params holds the params of the Usual type
type mockUsual_NoParams_params struct{}

// mockUsual_NoParams_resultMgr manages multiple results and the state of the Usual type
type mockUsual_NoParams_resultMgr struct {
	results  []*mockUsual_NoParams_results
	index    uint32
	anyTimes bool
}

// mockUsual_NoParams_results holds the results of the Usual type
type mockUsual_NoParams_results struct {
	sResult string
	err     error
}

// mockUsual_NoParams_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_NoParams_fnRecorder struct {
	params  mockUsual_NoParams_params
	results *mockUsual_NoParams_resultMgr
	mock    *mockUsual
}

// mockUsual_Nothing_params holds the params of the Usual type
type mockUsual_Nothing_params struct{}

// mockUsual_Nothing_resultMgr manages multiple results and the state of the Usual type
type mockUsual_Nothing_resultMgr struct {
	results  []*mockUsual_Nothing_results
	index    uint32
	anyTimes bool
}

// mockUsual_Nothing_results holds the results of the Usual type
type mockUsual_Nothing_results struct {
}

// mockUsual_Nothing_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_Nothing_fnRecorder struct {
	params  mockUsual_Nothing_params
	results *mockUsual_Nothing_resultMgr
	mock    *mockUsual
}

// mockUsual_Variadic_params holds the params of the Usual type
type mockUsual_Variadic_params struct {
	other bool
	args  hash.Hash
}

// mockUsual_Variadic_resultMgr manages multiple results and the state of the Usual type
type mockUsual_Variadic_resultMgr struct {
	results  []*mockUsual_Variadic_results
	index    uint32
	anyTimes bool
}

// mockUsual_Variadic_results holds the results of the Usual type
type mockUsual_Variadic_results struct {
	sResult string
	err     error
}

// mockUsual_Variadic_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_Variadic_fnRecorder struct {
	params  mockUsual_Variadic_params
	results *mockUsual_Variadic_resultMgr
	mock    *mockUsual
}

// mockUsual_RepeatedIds_params holds the params of the Usual type
type mockUsual_RepeatedIds_params struct {
	sParam1, sParam2 string
	bParam           bool
}

// mockUsual_RepeatedIds_resultMgr manages multiple results and the state of the Usual type
type mockUsual_RepeatedIds_resultMgr struct {
	results  []*mockUsual_RepeatedIds_results
	index    uint32
	anyTimes bool
}

// mockUsual_RepeatedIds_results holds the results of the Usual type
type mockUsual_RepeatedIds_results struct {
	sResult1, sResult2 string
	err                error
}

// mockUsual_RepeatedIds_fnRecorder routes recorded function calls to the mockUsual mock
type mockUsual_RepeatedIds_fnRecorder struct {
	params  mockUsual_RepeatedIds_params
	results *mockUsual_RepeatedIds_resultMgr
	mock    *mockUsual
}

// newMockUsual creates a new mock of the Usual type
func newMockUsual(t testing.MoqT) *mockUsual {
	return &mockUsual{
		t:                           t,
		resultsByParams_Usual:       map[mockUsual_Usual_params]*mockUsual_Usual_resultMgr{},
		params_Usual:                make(chan mockUsual_Usual_params, 100),
		resultsByParams_NoNames:     map[mockUsual_NoNames_params]*mockUsual_NoNames_resultMgr{},
		params_NoNames:              make(chan mockUsual_NoNames_params, 100),
		resultsByParams_NoResults:   map[mockUsual_NoResults_params]*mockUsual_NoResults_resultMgr{},
		params_NoResults:            make(chan mockUsual_NoResults_params, 100),
		resultsByParams_NoParams:    map[mockUsual_NoParams_params]*mockUsual_NoParams_resultMgr{},
		params_NoParams:             make(chan mockUsual_NoParams_params, 100),
		resultsByParams_Nothing:     map[mockUsual_Nothing_params]*mockUsual_Nothing_resultMgr{},
		params_Nothing:              make(chan mockUsual_Nothing_params, 100),
		resultsByParams_Variadic:    map[mockUsual_Variadic_params]*mockUsual_Variadic_resultMgr{},
		params_Variadic:             make(chan mockUsual_Variadic_params, 100),
		resultsByParams_RepeatedIds: map[mockUsual_RepeatedIds_params]*mockUsual_RepeatedIds_resultMgr{},
		params_RepeatedIds:          make(chan mockUsual_RepeatedIds_params, 100),
	}
}

// mock returns the mock implementation of the Usual type
func (m *mockUsual) mock() *mockUsual_mock {
	return &mockUsual_mock{
		mock: m,
	}
}

func (m *mockUsual_mock) Usual(sParam string, bParam bool) (sResult string, err error) {
	params := mockUsual_Usual_params{
		sParam: sParam,
		bParam: bParam,
	}
	m.mock.params_Usual <- params
	results, ok := m.mock.resultsByParams_Usual[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
		result := results.results[i]
		sResult = result.sResult
		err = result.err
	}
	return sResult, err
}

func (m *mockUsual_mock) NoNames(param1 string, param2 bool) (result1 string, result2 error) {
	params := mockUsual_NoNames_params{
		param1: param1,
		param2: param2,
	}
	m.mock.params_NoNames <- params
	results, ok := m.mock.resultsByParams_NoNames[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
		result := results.results[i]
		result1 = result.result1
		result2 = result.result2
	}
	return result1, result2
}

func (m *mockUsual_mock) NoResults(sParam string, bParam bool) {
	params := mockUsual_NoResults_params{
		sParam: sParam,
		bParam: bParam,
	}
	m.mock.params_NoResults <- params
	results, ok := m.mock.resultsByParams_NoResults[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
	}
	return
}

func (m *mockUsual_mock) NoParams() (sResult string, err error) {
	params := mockUsual_NoParams_params{}
	m.mock.params_NoParams <- params
	results, ok := m.mock.resultsByParams_NoParams[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
		result := results.results[i]
		sResult = result.sResult
		err = result.err
	}
	return sResult, err
}

func (m *mockUsual_mock) Nothing() {
	params := mockUsual_Nothing_params{}
	m.mock.params_Nothing <- params
	results, ok := m.mock.resultsByParams_Nothing[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
	}
	return
}

func (m *mockUsual_mock) Variadic(other bool, args ...string) (sResult string, err error) {
	params := mockUsual_Variadic_params{
		other: other,
		args:  hash.DeepHash(args),
	}
	m.mock.params_Variadic <- params
	results, ok := m.mock.resultsByParams_Variadic[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
		result := results.results[i]
		sResult = result.sResult
		err = result.err
	}
	return sResult, err
}

func (m *mockUsual_mock) RepeatedIds(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
	params := mockUsual_RepeatedIds_params{
		sParam1: sParam1,
		sParam2: sParam2,
		bParam:  bParam,
	}
	m.mock.params_RepeatedIds <- params
	results, ok := m.mock.resultsByParams_RepeatedIds[params]
	if ok {
		i := int(atomic.AddUint32(&results.index, 1)) - 1
		if i >= len(results.results) {
			if !results.anyTimes {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
				return
			}
			i = len(results.results) - 1
		}
		result := results.results[i]
		sResult1 = result.sResult1
		sResult2 = result.sResult2
		err = result.err
	}
	return sResult1, sResult2, err
}

// onCall returns the recorder implementation of the Usual type
func (m *mockUsual) onCall() *mockUsual_recorder {
	return &mockUsual_recorder{
		mock: m,
	}
}

func (m *mockUsual_recorder) Usual(sParam string, bParam bool) *mockUsual_Usual_fnRecorder {
	return &mockUsual_Usual_fnRecorder{
		params: mockUsual_Usual_params{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m.mock,
	}
}

func (r *mockUsual_Usual_fnRecorder) returnResults(sResult string, err error) *mockUsual_Usual_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_Usual[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_Usual_resultMgr{results: []*mockUsual_Usual_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_Usual[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_Usual_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockUsual_Usual_fnRecorder) times(count int) *mockUsual_Usual_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_Usual_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) NoNames(param1 string, param2 bool) *mockUsual_NoNames_fnRecorder {
	return &mockUsual_NoNames_fnRecorder{
		params: mockUsual_NoNames_params{
			param1: param1,
			param2: param2,
		},
		mock: m.mock,
	}
}

func (r *mockUsual_NoNames_fnRecorder) returnResults(result1 string, result2 error) *mockUsual_NoNames_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_NoNames[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_NoNames_resultMgr{results: []*mockUsual_NoNames_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_NoNames[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_NoNames_results{
		result1: result1,
		result2: result2,
	})
	return r
}

func (r *mockUsual_NoNames_fnRecorder) times(count int) *mockUsual_NoNames_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_NoNames_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) NoResults(sParam string, bParam bool) *mockUsual_NoResults_fnRecorder {
	return &mockUsual_NoResults_fnRecorder{
		params: mockUsual_NoResults_params{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m.mock,
	}
}

func (r *mockUsual_NoResults_fnRecorder) returnResults() *mockUsual_NoResults_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_NoResults[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_NoResults_resultMgr{results: []*mockUsual_NoResults_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_NoResults[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_NoResults_results{})
	return r
}

func (r *mockUsual_NoResults_fnRecorder) times(count int) *mockUsual_NoResults_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_NoResults_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) NoParams() *mockUsual_NoParams_fnRecorder {
	return &mockUsual_NoParams_fnRecorder{
		params: mockUsual_NoParams_params{},
		mock:   m.mock,
	}
}

func (r *mockUsual_NoParams_fnRecorder) returnResults(sResult string, err error) *mockUsual_NoParams_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_NoParams[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_NoParams_resultMgr{results: []*mockUsual_NoParams_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_NoParams[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_NoParams_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockUsual_NoParams_fnRecorder) times(count int) *mockUsual_NoParams_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_NoParams_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) Nothing() *mockUsual_Nothing_fnRecorder {
	return &mockUsual_Nothing_fnRecorder{
		params: mockUsual_Nothing_params{},
		mock:   m.mock,
	}
}

func (r *mockUsual_Nothing_fnRecorder) returnResults() *mockUsual_Nothing_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_Nothing[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_Nothing_resultMgr{results: []*mockUsual_Nothing_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_Nothing[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_Nothing_results{})
	return r
}

func (r *mockUsual_Nothing_fnRecorder) times(count int) *mockUsual_Nothing_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_Nothing_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) Variadic(other bool, args ...string) *mockUsual_Variadic_fnRecorder {
	return &mockUsual_Variadic_fnRecorder{
		params: mockUsual_Variadic_params{
			other: other,
			args:  hash.DeepHash(args),
		},
		mock: m.mock,
	}
}

func (r *mockUsual_Variadic_fnRecorder) returnResults(sResult string, err error) *mockUsual_Variadic_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_Variadic[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_Variadic_resultMgr{results: []*mockUsual_Variadic_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_Variadic[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_Variadic_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockUsual_Variadic_fnRecorder) times(count int) *mockUsual_Variadic_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_Variadic_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockUsual_recorder) RepeatedIds(sParam1, sParam2 string, bParam bool) *mockUsual_RepeatedIds_fnRecorder {
	return &mockUsual_RepeatedIds_fnRecorder{
		params: mockUsual_RepeatedIds_params{
			sParam1: sParam1,
			sParam2: sParam2,
			bParam:  bParam,
		},
		mock: m.mock,
	}
}

func (r *mockUsual_RepeatedIds_fnRecorder) returnResults(sResult1, sResult2 string, err error) *mockUsual_RepeatedIds_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams_RepeatedIds[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsual_RepeatedIds_resultMgr{results: []*mockUsual_RepeatedIds_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams_RepeatedIds[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsual_RepeatedIds_results{
		sResult1: sResult1,
		sResult2: sResult2,
		err:      err,
	})
	return r
}

func (r *mockUsual_RepeatedIds_fnRecorder) times(count int) *mockUsual_RepeatedIds_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsual_RepeatedIds_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}