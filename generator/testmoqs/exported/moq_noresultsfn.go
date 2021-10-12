// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmoqs"
	"github.com/myshkin5/moqueries/moq"
)

// MoqNoResultsFn holds the state of a moq of the NoResultsFn type
type MoqNoResultsFn struct {
	Scene           *moq.Scene
	Config          moq.Config
	ResultsByParams []MoqNoResultsFn_resultsByParams
}

// MoqNoResultsFn_mock isolates the mock interface of the NoResultsFn type
type MoqNoResultsFn_mock struct {
	Moq *MoqNoResultsFn
}

// MoqNoResultsFn_params holds the params of the NoResultsFn type
type MoqNoResultsFn_params struct {
	SParam string
	BParam bool
}

// MoqNoResultsFn_paramsKey holds the map key params of the NoResultsFn type
type MoqNoResultsFn_paramsKey struct {
	SParam string
	BParam bool
}

// MoqNoResultsFn_resultsByParams contains the results for a given set of parameters for the NoResultsFn type
type MoqNoResultsFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNoResultsFn_paramsKey]*MoqNoResultsFn_results
}

// MoqNoResultsFn_doFn defines the type of function needed when calling AndDo for the NoResultsFn type
type MoqNoResultsFn_doFn func(sParam string, bParam bool)

// MoqNoResultsFn_doReturnFn defines the type of function needed when calling DoReturnResults for the NoResultsFn type
type MoqNoResultsFn_doReturnFn func(sParam string, bParam bool)

// MoqNoResultsFn_results holds the results of the NoResultsFn type
type MoqNoResultsFn_results struct {
	Params  MoqNoResultsFn_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqNoResultsFn_doFn
		DoReturnFn MoqNoResultsFn_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNoResultsFn_fnRecorder routes recorded function calls to the MoqNoResultsFn moq
type MoqNoResultsFn_fnRecorder struct {
	Params    MoqNoResultsFn_params
	ParamsKey MoqNoResultsFn_paramsKey
	AnyParams uint64
	Sequence  bool
	Results   *MoqNoResultsFn_results
	Moq       *MoqNoResultsFn
}

// MoqNoResultsFn_anyParams isolates the any params functions of the NoResultsFn type
type MoqNoResultsFn_anyParams struct {
	Recorder *MoqNoResultsFn_fnRecorder
}

// NewMoqNoResultsFn creates a new moq of the NoResultsFn type
func NewMoqNoResultsFn(scene *moq.Scene, config *moq.Config) *MoqNoResultsFn {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNoResultsFn{
		Scene:  scene,
		Config: *config,
	}
	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NoResultsFn type
func (m *MoqNoResultsFn) Mock() testmoqs.NoResultsFn {
	return func(sParam string, bParam bool) { moq := &MoqNoResultsFn_mock{Moq: m}; moq.Fn(sParam, bParam) }
}

func (m *MoqNoResultsFn_mock) Fn(sParam string, bParam bool) {
	params := MoqNoResultsFn_params{
		SParam: sParam,
		BParam: bParam,
	}
	var results *MoqNoResultsFn_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		var sParamUsed string
		if resultsByParams.AnyParams&(1<<0) == 0 {
			sParamUsed = sParam
		}
		var bParamUsed bool
		if resultsByParams.AnyParams&(1<<1) == 0 {
			bParamUsed = bParam
		}
		paramsKey := MoqNoResultsFn_paramsKey{
			SParam: sParamUsed,
			BParam: bParamUsed,
		}
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
		result.DoFn(sParam, bParam)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(sParam, bParam)
	}
	return
}

func (m *MoqNoResultsFn) OnCall(sParam string, bParam bool) *MoqNoResultsFn_fnRecorder {
	return &MoqNoResultsFn_fnRecorder{
		Params: MoqNoResultsFn_params{
			SParam: sParam,
			BParam: bParam,
		},
		ParamsKey: MoqNoResultsFn_paramsKey{
			SParam: sParam,
			BParam: bParam,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNoResultsFn_fnRecorder) Any() *MoqNoResultsFn_anyParams {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	return &MoqNoResultsFn_anyParams{Recorder: r}
}

func (a *MoqNoResultsFn_anyParams) SParam() *MoqNoResultsFn_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNoResultsFn_anyParams) BParam() *MoqNoResultsFn_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNoResultsFn_fnRecorder) Seq() *MoqNoResultsFn_fnRecorder {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNoResultsFn_fnRecorder) NoSeq() *MoqNoResultsFn_fnRecorder {
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNoResultsFn_fnRecorder) ReturnResults() *MoqNoResultsFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqNoResultsFn_doFn
		DoReturnFn MoqNoResultsFn_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNoResultsFn_fnRecorder) AndDo(fn MoqNoResultsFn_doFn) *MoqNoResultsFn_fnRecorder {
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNoResultsFn_fnRecorder) DoReturnResults(fn MoqNoResultsFn_doReturnFn) *MoqNoResultsFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqNoResultsFn_doFn
		DoReturnFn MoqNoResultsFn_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNoResultsFn_fnRecorder) FindResults() {
	if r.Results == nil {
		anyCount := bits.OnesCount64(r.AnyParams)
		insertAt := -1
		var results *MoqNoResultsFn_resultsByParams
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
			results = &MoqNoResultsFn_resultsByParams{
				AnyCount:  anyCount,
				AnyParams: r.AnyParams,
				Results:   map[MoqNoResultsFn_paramsKey]*MoqNoResultsFn_results{},
			}
			r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
				copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
				r.Moq.ResultsByParams[insertAt] = *results
			}
		}

		var sParamUsed string
		if r.AnyParams&(1<<0) == 0 {
			sParamUsed = r.ParamsKey.SParam
		}
		var bParamUsed bool
		if r.AnyParams&(1<<1) == 0 {
			bParamUsed = r.ParamsKey.BParam
		}
		paramsKey := MoqNoResultsFn_paramsKey{
			SParam: sParamUsed,
			BParam: bParamUsed,
		}

		var ok bool
		r.Results, ok = results.Results[paramsKey]
		if !ok {
			r.Results = &MoqNoResultsFn_results{
				Params:  r.Params,
				Results: nil,
				Index:   0,
				Repeat:  &moq.RepeatVal{},
			}
			results.Results[paramsKey] = r.Results
		}
	}
	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNoResultsFn_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNoResultsFn_fnRecorder {
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqNoResultsFn_doFn
				DoReturnFn MoqNoResultsFn_doReturnFn
			}{
				Values:   &struct{}{},
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

// Reset resets the state of the moq
func (m *MoqNoResultsFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNoResultsFn) AssertExpectationsMet() {
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.Params)
			}
		}
	}
}
