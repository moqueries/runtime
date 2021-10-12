// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package demo_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/demo"
	"github.com/myshkin5/moqueries/moq"
)

// moqIsFavorite holds the state of a moq of the IsFavorite type
type moqIsFavorite struct {
	scene           *moq.Scene
	config          moq.Config
	resultsByParams []moqIsFavorite_resultsByParams
}

// moqIsFavorite_mock isolates the mock interface of the IsFavorite type
type moqIsFavorite_mock struct {
	moq *moqIsFavorite
}

// moqIsFavorite_params holds the params of the IsFavorite type
type moqIsFavorite_params struct{ n int }

// moqIsFavorite_paramsKey holds the map key params of the IsFavorite type
type moqIsFavorite_paramsKey struct{ n int }

// moqIsFavorite_resultsByParams contains the results for a given set of parameters for the IsFavorite type
type moqIsFavorite_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqIsFavorite_paramsKey]*moqIsFavorite_results
}

// moqIsFavorite_doFn defines the type of function needed when calling andDo for the IsFavorite type
type moqIsFavorite_doFn func(n int)

// moqIsFavorite_doReturnFn defines the type of function needed when calling doReturnResults for the IsFavorite type
type moqIsFavorite_doReturnFn func(n int) bool

// moqIsFavorite_results holds the results of the IsFavorite type
type moqIsFavorite_results struct {
	params  moqIsFavorite_params
	results []struct {
		values *struct {
			result1 bool
		}
		sequence   uint32
		doFn       moqIsFavorite_doFn
		doReturnFn moqIsFavorite_doReturnFn
	}
	index  uint32
	repeat *moq.RepeatVal
}

// moqIsFavorite_fnRecorder routes recorded function calls to the moqIsFavorite moq
type moqIsFavorite_fnRecorder struct {
	params    moqIsFavorite_params
	paramsKey moqIsFavorite_paramsKey
	anyParams uint64
	sequence  bool
	results   *moqIsFavorite_results
	moq       *moqIsFavorite
}

// moqIsFavorite_anyParams isolates the any params functions of the IsFavorite type
type moqIsFavorite_anyParams struct {
	recorder *moqIsFavorite_fnRecorder
}

// newMoqIsFavorite creates a new moq of the IsFavorite type
func newMoqIsFavorite(scene *moq.Scene, config *moq.Config) *moqIsFavorite {
	if config == nil {
		config = &moq.Config{}
	}
	m := &moqIsFavorite{
		scene:  scene,
		config: *config,
	}
	scene.AddMoq(m)
	return m
}

// mock returns the moq implementation of the IsFavorite type
func (m *moqIsFavorite) mock() demo.IsFavorite {
	return func(n int) bool { moq := &moqIsFavorite_mock{moq: m}; return moq.fn(n) }
}

func (m *moqIsFavorite_mock) fn(n int) (result1 bool) {
	params := moqIsFavorite_params{
		n: n,
	}
	var results *moqIsFavorite_results
	for _, resultsByParams := range m.moq.resultsByParams {
		var nUsed int
		if resultsByParams.anyParams&(1<<0) == 0 {
			nUsed = n
		}
		paramsKey := moqIsFavorite_paramsKey{
			n: nUsed,
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
	if i >= results.repeat.ResultCount {
		if !results.repeat.AnyTimes {
			if m.moq.config.Expectation == moq.Strict {
				m.moq.scene.T.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = results.repeat.ResultCount - 1
	}

	result := results.results[i]
	if result.sequence != 0 {
		sequence := m.moq.scene.NextMockSequence()
		if (!results.repeat.AnyTimes && result.sequence != sequence) || result.sequence > sequence {
			m.moq.scene.T.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.doFn != nil {
		result.doFn(n)
	}

	if result.values != nil {
		result1 = result.values.result1
	}
	if result.doReturnFn != nil {
		result1 = result.doReturnFn(n)
	}
	return
}

func (m *moqIsFavorite) onCall(n int) *moqIsFavorite_fnRecorder {
	return &moqIsFavorite_fnRecorder{
		params: moqIsFavorite_params{
			n: n,
		},
		paramsKey: moqIsFavorite_paramsKey{
			n: n,
		},
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		moq:      m,
	}
}

func (r *moqIsFavorite_fnRecorder) any() *moqIsFavorite_anyParams {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	return &moqIsFavorite_anyParams{recorder: r}
}

func (a *moqIsFavorite_anyParams) n() *moqIsFavorite_fnRecorder {
	a.recorder.anyParams |= 1 << 0
	return a.recorder
}

func (r *moqIsFavorite_fnRecorder) seq() *moqIsFavorite_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqIsFavorite_fnRecorder) noSeq() *moqIsFavorite_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqIsFavorite_fnRecorder) returnResults(result1 bool) *moqIsFavorite_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 bool
		}
		sequence   uint32
		doFn       moqIsFavorite_doFn
		doReturnFn moqIsFavorite_doReturnFn
	}{
		values: &struct {
			result1 bool
		}{
			result1: result1,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqIsFavorite_fnRecorder) andDo(fn moqIsFavorite_doFn) *moqIsFavorite_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqIsFavorite_fnRecorder) doReturnResults(fn moqIsFavorite_doReturnFn) *moqIsFavorite_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 bool
		}
		sequence   uint32
		doFn       moqIsFavorite_doFn
		doReturnFn moqIsFavorite_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqIsFavorite_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *moqIsFavorite_resultsByParams
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
			results = &moqIsFavorite_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[moqIsFavorite_paramsKey]*moqIsFavorite_results{},
			}
			r.moq.resultsByParams = append(r.moq.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams) {
				copy(r.moq.resultsByParams[insertAt+1:], r.moq.resultsByParams[insertAt:0])
				r.moq.resultsByParams[insertAt] = *results
			}
		}

		var nUsed int
		if r.anyParams&(1<<0) == 0 {
			nUsed = r.paramsKey.n
		}
		paramsKey := moqIsFavorite_paramsKey{
			n: nUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &moqIsFavorite_results{
				params:  r.params,
				results: nil,
				index:   0,
				repeat:  &moq.RepeatVal{},
			}
			results.results[paramsKey] = r.results
		}
	}
	r.results.repeat.Increment(r.moq.scene.T)
}

func (r *moqIsFavorite_fnRecorder) repeat(repeaters ...moq.Repeater) *moqIsFavorite_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling repeat")
		return nil
	}
	r.results.repeat.Repeat(r.moq.scene.T, repeaters)
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < r.results.repeat.ResultCount-1; n++ {
		if r.sequence {
			last = struct {
				values *struct {
					result1 bool
				}
				sequence   uint32
				doFn       moqIsFavorite_doFn
				doReturnFn moqIsFavorite_doReturnFn
			}{
				values: &struct {
					result1 bool
				}{
					result1: last.values.result1,
				},
				sequence: r.moq.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

// Reset resets the state of the moq
func (m *moqIsFavorite) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *moqIsFavorite) AssertExpectationsMet() {
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := results.repeat.MinTimes - int(atomic.LoadUint32(&results.index))
			if missing > 0 {
				m.scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}
