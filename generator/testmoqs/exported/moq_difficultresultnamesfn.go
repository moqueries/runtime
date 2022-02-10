// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmoqs"
	"github.com/myshkin5/moqueries/moq"
)

// MoqDifficultResultNamesFn holds the state of a moq of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDifficultResultNamesFn_mock

	ResultsByParams []MoqDifficultResultNamesFn_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqDifficultResultNamesFn_mock isolates the mock interface of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_mock struct {
	Moq *MoqDifficultResultNamesFn
}

// MoqDifficultResultNamesFn_params holds the params of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_params struct{}

// MoqDifficultResultNamesFn_paramsKey holds the map key params of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDifficultResultNamesFn_resultsByParams contains the results for a given set of parameters for the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDifficultResultNamesFn_paramsKey]*MoqDifficultResultNamesFn_results
}

// MoqDifficultResultNamesFn_doFn defines the type of function needed when calling AndDo for the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_doFn func()

// MoqDifficultResultNamesFn_doReturnFn defines the type of function needed when calling DoReturnResults for the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_doReturnFn func() (m, r string, sequence error, param, params int, result, results float32)

// MoqDifficultResultNamesFn_results holds the results of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_results struct {
	Params  MoqDifficultResultNamesFn_params
	Results []struct {
		Values *struct {
			Result1, Result2 string
			Result3          error
			Param, Result5   int
			Result6, Result7 float32
		}
		Sequence   uint32
		DoFn       MoqDifficultResultNamesFn_doFn
		DoReturnFn MoqDifficultResultNamesFn_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDifficultResultNamesFn_fnRecorder routes recorded function calls to the MoqDifficultResultNamesFn moq
type MoqDifficultResultNamesFn_fnRecorder struct {
	Params    MoqDifficultResultNamesFn_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDifficultResultNamesFn_results
	Moq       *MoqDifficultResultNamesFn
}

// MoqDifficultResultNamesFn_anyParams isolates the any params functions of the DifficultResultNamesFn type
type MoqDifficultResultNamesFn_anyParams struct {
	Recorder *MoqDifficultResultNamesFn_fnRecorder
}

// NewMoqDifficultResultNamesFn creates a new moq of the DifficultResultNamesFn type
func NewMoqDifficultResultNamesFn(scene *moq.Scene, config *moq.Config) *MoqDifficultResultNamesFn {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDifficultResultNamesFn{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDifficultResultNamesFn_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DifficultResultNamesFn type
func (m *MoqDifficultResultNamesFn) Mock() testmoqs.DifficultResultNamesFn {
	return func() (_, _ string, _ error, _, _ int, _, _ float32) {
		moq := &MoqDifficultResultNamesFn_mock{Moq: m}
		return moq.Fn()
	}
}

func (m *MoqDifficultResultNamesFn_mock) Fn() (result1, result2 string, result3 error, param, result5 int, result6, result7 float32) {
	params := MoqDifficultResultNamesFn_params{}
	var results *MoqDifficultResultNamesFn_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
		result3 = result.Values.Result3
		param = result.Values.Param
		result5 = result.Values.Result5
		result6 = result.Values.Result6
		result7 = result.Values.Result7
	}
	if result.DoReturnFn != nil {
		result1, result2, result3, param, result5, result6, result7 = result.DoReturnFn()
	}
	return
}

func (m *MoqDifficultResultNamesFn) OnCall() *MoqDifficultResultNamesFn_fnRecorder {
	return &MoqDifficultResultNamesFn_fnRecorder{
		Params:   MoqDifficultResultNamesFn_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDifficultResultNamesFn_fnRecorder) Any() *MoqDifficultResultNamesFn_anyParams {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	return &MoqDifficultResultNamesFn_anyParams{Recorder: r}
}

func (r *MoqDifficultResultNamesFn_fnRecorder) Seq() *MoqDifficultResultNamesFn_fnRecorder {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDifficultResultNamesFn_fnRecorder) NoSeq() *MoqDifficultResultNamesFn_fnRecorder {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDifficultResultNamesFn_fnRecorder) ReturnResults(result1, result2 string, result3 error, param, result5 int, result6, result7 float32) *MoqDifficultResultNamesFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1, Result2 string
			Result3          error
			Param, Result5   int
			Result6, Result7 float32
		}
		Sequence   uint32
		DoFn       MoqDifficultResultNamesFn_doFn
		DoReturnFn MoqDifficultResultNamesFn_doReturnFn
	}{
		Values: &struct {
			Result1, Result2 string
			Result3          error
			Param, Result5   int
			Result6, Result7 float32
		}{
			Result1: result1,
			Result2: result2,
			Result3: result3,
			Param:   param,
			Result5: result5,
			Result6: result6,
			Result7: result7,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDifficultResultNamesFn_fnRecorder) AndDo(fn MoqDifficultResultNamesFn_doFn) *MoqDifficultResultNamesFn_fnRecorder {
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDifficultResultNamesFn_fnRecorder) DoReturnResults(fn MoqDifficultResultNamesFn_doReturnFn) *MoqDifficultResultNamesFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1, Result2 string
			Result3          error
			Param, Result5   int
			Result6, Result7 float32
		}
		Sequence   uint32
		DoFn       MoqDifficultResultNamesFn_doFn
		DoReturnFn MoqDifficultResultNamesFn_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDifficultResultNamesFn_fnRecorder) FindResults() {
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDifficultResultNamesFn_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDifficultResultNamesFn_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDifficultResultNamesFn_paramsKey]*MoqDifficultResultNamesFn_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDifficultResultNamesFn_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDifficultResultNamesFn_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDifficultResultNamesFn_fnRecorder {
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Result1, Result2 string
					Result3          error
					Param, Result5   int
					Result6, Result7 float32
				}
				Sequence   uint32
				DoFn       MoqDifficultResultNamesFn_doFn
				DoReturnFn MoqDifficultResultNamesFn_doReturnFn
			}{
				Values: &struct {
					Result1, Result2 string
					Result3          error
					Param, Result5   int
					Result6, Result7 float32
				}{
					Result1: last.Values.Result1,
					Result2: last.Values.Result2,
					Result3: last.Values.Result3,
					Param:   last.Values.Param,
					Result5: last.Values.Result5,
					Result6: last.Values.Result6,
					Result7: last.Values.Result7,
				},
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDifficultResultNamesFn) ParamsKey(params MoqDifficultResultNamesFn_params, anyParams uint64) MoqDifficultResultNamesFn_paramsKey {
	return MoqDifficultResultNamesFn_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqDifficultResultNamesFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDifficultResultNamesFn) AssertExpectationsMet() {
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.Params)
			}
		}
	}
}
