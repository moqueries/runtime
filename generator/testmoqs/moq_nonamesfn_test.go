// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmoqs_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmoqs"
	"github.com/myshkin5/moqueries/moq"
)

// moqNoNamesFn holds the state of a moq of the NoNamesFn type
type moqNoNamesFn struct {
	scene           *moq.Scene
	config          moq.Config
	resultsByParams []moqNoNamesFn_resultsByParams
}

// moqNoNamesFn_mock isolates the mock interface of the NoNamesFn type
type moqNoNamesFn_mock struct {
	moq *moqNoNamesFn
}

// moqNoNamesFn_params holds the params of the NoNamesFn type
type moqNoNamesFn_params struct {
	param1 string
	param2 bool
}

// moqNoNamesFn_paramsKey holds the map key params of the NoNamesFn type
type moqNoNamesFn_paramsKey struct {
	param1 string
	param2 bool
}

// moqNoNamesFn_resultsByParams contains the results for a given set of parameters for the NoNamesFn type
type moqNoNamesFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqNoNamesFn_paramsKey]*moqNoNamesFn_results
}

// moqNoNamesFn_doFn defines the type of function needed when calling andDo for the NoNamesFn type
type moqNoNamesFn_doFn func(string, bool)

// moqNoNamesFn_doReturnFn defines the type of function needed when calling doReturnResults for the NoNamesFn type
type moqNoNamesFn_doReturnFn func(string, bool) (string, error)

// moqNoNamesFn_results holds the results of the NoNamesFn type
type moqNoNamesFn_results struct {
	params  moqNoNamesFn_params
	results []struct {
		values *struct {
			result1 string
			result2 error
		}
		sequence   uint32
		doFn       moqNoNamesFn_doFn
		doReturnFn moqNoNamesFn_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// moqNoNamesFn_fnRecorder routes recorded function calls to the moqNoNamesFn moq
type moqNoNamesFn_fnRecorder struct {
	params    moqNoNamesFn_params
	paramsKey moqNoNamesFn_paramsKey
	anyParams uint64
	sequence  bool
	results   *moqNoNamesFn_results
	moq       *moqNoNamesFn
}

// newMoqNoNamesFn creates a new moq of the NoNamesFn type
func newMoqNoNamesFn(scene *moq.Scene, config *moq.Config) *moqNoNamesFn {
	if config == nil {
		config = &moq.Config{}
	}
	m := &moqNoNamesFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMoq(m)
	return m
}

// mock returns the moq implementation of the NoNamesFn type
func (m *moqNoNamesFn) mock() testmoqs.NoNamesFn {
	return func(param1 string, param2 bool) (string, error) {
		moq := &moqNoNamesFn_mock{moq: m}
		return moq.fn(param1, param2)
	}
}

func (m *moqNoNamesFn_mock) fn(param1 string, param2 bool) (result1 string, result2 error) {
	params := moqNoNamesFn_params{
		param1: param1,
		param2: param2,
	}
	var results *moqNoNamesFn_results
	for _, resultsByParams := range m.moq.resultsByParams {
		var param1Used string
		if resultsByParams.anyParams&(1<<0) == 0 {
			param1Used = param1
		}
		var param2Used bool
		if resultsByParams.anyParams&(1<<1) == 0 {
			param2Used = param2
		}
		paramsKey := moqNoNamesFn_paramsKey{
			param1: param1Used,
			param2: param2Used,
		}
		var ok bool
		results, ok = resultsByParams.results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.moq.config.Expectation == moq.Strict {
			m.moq.scene.T.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.index, 1)) - 1
	if i >= len(results.results) {
		if !results.anyTimes {
			if m.moq.config.Expectation == moq.Strict {
				m.moq.scene.T.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.results) - 1
	}

	result := results.results[i]
	if result.sequence != 0 {
		sequence := m.moq.scene.NextMockSequence()
		if (!results.anyTimes && result.sequence != sequence) || result.sequence > sequence {
			m.moq.scene.T.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.doFn != nil {
		result.doFn(param1, param2)
	}

	if result.values != nil {
		result1 = result.values.result1
		result2 = result.values.result2
	}
	if result.doReturnFn != nil {
		result1, result2 = result.doReturnFn(param1, param2)
	}
	return
}

func (m *moqNoNamesFn) onCall(param1 string, param2 bool) *moqNoNamesFn_fnRecorder {
	return &moqNoNamesFn_fnRecorder{
		params: moqNoNamesFn_params{
			param1: param1,
			param2: param2,
		},
		paramsKey: moqNoNamesFn_paramsKey{
			param1: param1,
			param2: param2,
		},
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		moq:      m,
	}
}

func (r *moqNoNamesFn_fnRecorder) anyParam1() *moqNoNamesFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *moqNoNamesFn_fnRecorder) anyParam2() *moqNoNamesFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *moqNoNamesFn_fnRecorder) seq() *moqNoNamesFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqNoNamesFn_fnRecorder) noSeq() *moqNoNamesFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqNoNamesFn_fnRecorder) returnResults(result1 string, result2 error) *moqNoNamesFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 string
			result2 error
		}
		sequence   uint32
		doFn       moqNoNamesFn_doFn
		doReturnFn moqNoNamesFn_doReturnFn
	}{
		values: &struct {
			result1 string
			result2 error
		}{
			result1: result1,
			result2: result2,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqNoNamesFn_fnRecorder) andDo(fn moqNoNamesFn_doFn) *moqNoNamesFn_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqNoNamesFn_fnRecorder) doReturnResults(fn moqNoNamesFn_doReturnFn) *moqNoNamesFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 string
			result2 error
		}
		sequence   uint32
		doFn       moqNoNamesFn_doFn
		doReturnFn moqNoNamesFn_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqNoNamesFn_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *moqNoNamesFn_resultsByParams
		for n, res := range r.moq.resultsByParams {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &moqNoNamesFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[moqNoNamesFn_paramsKey]*moqNoNamesFn_results{},
			}
			r.moq.resultsByParams = append(r.moq.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams) {
				copy(r.moq.resultsByParams[insertAt+1:], r.moq.resultsByParams[insertAt:0])
				r.moq.resultsByParams[insertAt] = *results
			}
		}

		var param1Used string
		if r.anyParams&(1<<0) == 0 {
			param1Used = r.paramsKey.param1
		}
		var param2Used bool
		if r.anyParams&(1<<1) == 0 {
			param2Used = r.paramsKey.param2
		}
		paramsKey := moqNoNamesFn_paramsKey{
			param1: param1Used,
			param2: param2Used,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &moqNoNamesFn_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *moqNoNamesFn_fnRecorder) times(count int) *moqNoNamesFn_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					result1 string
					result2 error
				}
				sequence   uint32
				doFn       moqNoNamesFn_doFn
				doReturnFn moqNoNamesFn_doReturnFn
			}{
				values: &struct {
					result1 string
					result2 error
				}{
					result1: last.values.result1,
					result2: last.values.result2,
				},
				sequence: r.moq.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *moqNoNamesFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the moq
func (m *moqNoNamesFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *moqNoNamesFn) AssertExpectationsMet() {
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := len(results.results) - int(atomic.LoadUint32(&results.index))
			if missing == 1 && results.anyTimes == true {
				continue
			}
			if missing > 0 {
				m.scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}
