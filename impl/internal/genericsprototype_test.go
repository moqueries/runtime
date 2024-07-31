package internal_test

import (
	"fmt"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/impl"
	"moqueries.org/runtime/moq"
)

type Generics[S, B any, R ~string, E error] interface {
	Usual(sParam S, bParam B) (sResult R, err E)
}

// The following type assertion assures that testmoqs.Generics is mocked
// completely
var _ Generics[any, any, string, error] = (*MoqGenerics_mock[any, any, string, error])(nil)

// MoqGenerics holds the state of a moq of the Generics type
type MoqGenerics[S, B any, R ~string, E error] struct {
	Moq *MoqGenerics_mock[S, B, R, E]

	Moq_Usual *impl.Moq[
		*MoqGenerics_Usual_adaptor[S, B, R, E],
		MoqGenerics_Usual_params[S, B, R, E],
		MoqGenerics_Usual_paramsKey[S, B, R, E],
		MoqGenerics_Usual_results[S, B, R, E]]

	Runtime MoqGenerics_runtime
}

// MoqGenerics_mock isolates the mock interface of the Generics type
type MoqGenerics_mock[S, B any, R ~string, E error] struct {
	Moq *MoqGenerics[S, B, R, E]
}

// MoqGenerics_recorder isolates the recorder interface of the Generics type
type MoqGenerics_recorder[S, B any, R ~string, E error] struct {
	Moq *MoqGenerics[S, B, R, E]
}

// MoqGenerics_runtime holds runtime configuration for the Generics type
type MoqGenerics_runtime struct {
	ParameterIndexing struct {
		Usual MoqGenerics_Usual_paramIndexing
	}
}

// MoqGenerics_Usual_adaptor adapts MoqGenerics a needed by the runtime
type MoqGenerics_Usual_adaptor[S, B any, R ~string, E error] struct {
	Moq *MoqGenerics[S, B, R, E]
}

// MoqGenerics_Usual_params holds the params of the Generics type
type MoqGenerics_Usual_params[S, B any, R ~string, E error] struct {
	SParam S
	BParam B
}

// MoqGenerics_Usual_paramsKey holds the map key params of the Generics type
type MoqGenerics_Usual_paramsKey[S, B any, R ~string, E error] struct {
	Params struct{}
	Hashes struct {
		SParam hash.Hash
		BParam hash.Hash
	}
}

// MoqGenerics_Usual_results holds the results of the Generics type
type MoqGenerics_Usual_results[S, B any, R ~string, E error] struct {
	SResult R
	Err     E
}

// MoqGenerics_Usual_paramIndexing hold the parameter indexing runtime
// configuration for the Generics type
type MoqGenerics_Usual_paramIndexing struct {
	SParam moq.ParamIndexing
	BParam moq.ParamIndexing
}

// MoqGenerics_Usual_doFn defines the type of function needed when calling
// AndDo for the Generics type
type MoqGenerics_Usual_doFn[S, B any, R ~string, E error] func(sParam S, bParam B)

// MoqGenerics_Usual_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Generics type
type MoqGenerics_Usual_doReturnFn[S, B any, R ~string, E error] func(sParam S, bParam B) (sResult R, err E)

// MoqGenerics_Usual_recorder routes recorded function calls to the MoqGenerics
// moq
type MoqGenerics_Usual_recorder[S, B any, R ~string, E error] struct {
	Recorder *impl.Recorder[
		*MoqGenerics_Usual_adaptor[S, B, R, E],
		MoqGenerics_Usual_params[S, B, R, E],
		MoqGenerics_Usual_paramsKey[S, B, R, E],
		MoqGenerics_Usual_results[S, B, R, E]]
}

// MoqGenerics_Usual_anyParams isolates the any params functions of the
// Generics type
type MoqGenerics_Usual_anyParams[S, B any, R ~string, E error] struct {
	Recorder *MoqGenerics_Usual_recorder[S, B, R, E]
}

// NewMoqGenerics creates a new moq of the Generics type
func NewMoqGenerics[S, B any, R ~string, E error](scene *moq.Scene, config *moq.Config) *MoqGenerics[S, B, R, E] {
	adaptor0 := &MoqGenerics_Usual_adaptor[S, B, R, E]{}
	m := &MoqGenerics[S, B, R, E]{
		Moq: &MoqGenerics_mock[S, B, R, E]{},

		Moq_Usual: impl.NewMoq[
			*MoqGenerics_Usual_adaptor[S, B, R, E],
			MoqGenerics_Usual_params[S, B, R, E],
			MoqGenerics_Usual_paramsKey[S, B, R, E],
			MoqGenerics_Usual_results[S, B, R, E]](scene, adaptor0, config),

		Runtime: MoqGenerics_runtime{ParameterIndexing: struct {
			Usual MoqGenerics_Usual_paramIndexing
		}{
			Usual: MoqGenerics_Usual_paramIndexing{
				SParam: moq.ParamIndexByHash,
				BParam: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	adaptor0.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Generics type
func (m *MoqGenerics[S, B, R, E]) Mock() *MoqGenerics_mock[S, B, R, E] { return m.Moq }

func (m *MoqGenerics_mock[S, B, R, E]) Usual(sParam S, bParam B) (R, E) {
	m.Moq.Moq_Usual.Scene.T.Helper()
	params := MoqGenerics_Usual_params[S, B, R, E]{
		SParam: sParam,
		BParam: bParam,
	}

	var result0 R
	var result1 E
	if result := m.Moq.Moq_Usual.Function(params); result != nil {
		result0 = result.SResult
		result1 = result.Err
	}
	return result0, result1
}

// OnCall returns the recorder implementation of the Generics type
func (m *MoqGenerics[S, B, R, E]) OnCall() *MoqGenerics_recorder[S, B, R, E] {
	return &MoqGenerics_recorder[S, B, R, E]{
		Moq: m,
	}
}

func (m *MoqGenerics_recorder[S, B, R, E]) Usual(sParam S, bParam B) *MoqGenerics_Usual_recorder[S, B, R, E] {
	return &MoqGenerics_Usual_recorder[S, B, R, E]{
		Recorder: m.Moq.Moq_Usual.OnCall(MoqGenerics_Usual_params[S, B, R, E]{
			SParam: sParam,
			BParam: bParam,
		}),
	}
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) Any() *MoqGenerics_Usual_anyParams[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.IsAnyPermitted(true) {
		return nil
	}
	return &MoqGenerics_Usual_anyParams[S, B, R, E]{Recorder: r}
}

func (a *MoqGenerics_Usual_anyParams[S, B, R, E]) SParam() *MoqGenerics_Usual_recorder[S, B, R, E] {
	a.Recorder.Recorder.AnyParam(0)
	return a.Recorder
}

func (a *MoqGenerics_Usual_anyParams[S, B, R, E]) BParam() *MoqGenerics_Usual_recorder[S, B, R, E] {
	a.Recorder.Recorder.AnyParam(1)
	return a.Recorder
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) Seq() *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(true, "Seq", true) {
		return nil
	}
	return r
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) NoSeq() *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(false, "NoSeq", true) {
		return nil
	}
	return r
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) ReturnResults(
	sResult R, err E,
) *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.ReturnResults(MoqGenerics_Usual_results[S, B, R, E]{
		SResult: sResult,
		Err:     err,
	})
	return r
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) AndDo(
	fn MoqGenerics_Usual_doFn[S, B, R, E],
) *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.AndDo(func(params MoqGenerics_Usual_params[S, B, R, E]) {
		fn(params.SParam, params.BParam)
	}, true) {
		return nil
	}
	return r
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) DoReturnResults(
	fn MoqGenerics_Usual_doReturnFn[S, B, R, E],
) *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.DoReturnResults(func(params MoqGenerics_Usual_params[S, B, R, E]) *MoqGenerics_Usual_results[S, B, R, E] {
		sResult, err := fn(params.SParam, params.BParam)
		return &MoqGenerics_Usual_results[S, B, R, E]{
			SResult: sResult,
			Err:     err,
		}
	})
	return r
}

func (r *MoqGenerics_Usual_recorder[S, B, R, E]) Repeat(
	repeaters ...moq.Repeater,
) *MoqGenerics_Usual_recorder[S, B, R, E] {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Repeat(repeaters, true) {
		return nil
	}
	return r
}

func (*MoqGenerics_Usual_adaptor[S, B, R, E]) PrettyParams(params MoqGenerics_Usual_params[S, B, R, E]) string {
	return fmt.Sprintf("Usual(%#v, %#v)", params.SParam, params.BParam)
}

func (a *MoqGenerics_Usual_adaptor[S, B, R, E]) ParamsKey(
	params MoqGenerics_Usual_params[S, B, R, E], anyParams uint64,
) MoqGenerics_Usual_paramsKey[S, B, R, E] {
	a.Moq.Moq_Usual.Scene.T.Helper()
	sParamUsedHash := impl.HashOnlyParamKey(a.Moq.Moq_Usual.Scene.T,
		params.SParam, "sParam", 0, a.Moq.Runtime.ParameterIndexing.Usual.SParam, anyParams)
	bParamUsedHash := impl.HashOnlyParamKey(a.Moq.Moq_Usual.Scene.T,
		params.BParam, "bParam", 1, a.Moq.Runtime.ParameterIndexing.Usual.BParam, anyParams)
	return MoqGenerics_Usual_paramsKey[S, B, R, E]{
		Params: struct{}{},
		Hashes: struct {
			SParam hash.Hash
			BParam hash.Hash
		}{
			SParam: sParamUsedHash,
			BParam: bParamUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGenerics[S, B, R, E]) Reset() {
	m.Moq_Usual.Reset()
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGenerics[S, B, R, E]) AssertExpectationsMet() {
	m.Moq_Usual.Scene.T.Helper()
	m.Moq_Usual.AssertExpectationsMet()
}
