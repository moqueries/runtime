package internal_test

import (
	"fmt"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/impl"
	"moqueries.org/runtime/moq"
)

type Usual interface {
	Usual(sParam string, bParam bool) (sResult string, err error)
	Nothing()
}

// The following type assertion assures that testmoqs.Usual is mocked
// completely
var _ Usual = (*MoqUsual_mock)(nil)

// MoqUsual holds the state of a moq of the Usual type
type MoqUsual struct {
	Moq *MoqUsual_mock

	Moq_Usual *impl.Moq[
		*MoqUsual_Usual_adaptor,
		MoqUsual_Usual_params,
		MoqUsual_Usual_paramsKey,
		MoqUsual_Usual_results,
	]
	Moq_Nothing *impl.Moq[
		*MoqUsual_Nothing_adaptor,
		MoqUsual_Nothing_params,
		MoqUsual_Nothing_paramsKey,
		MoqUsual_Nothing_results,
	]

	Runtime MoqUsual_runtime
}

// MoqUsual_mock isolates the mock interface of the Usual type
type MoqUsual_mock struct {
	Moq *MoqUsual
}

// MoqUsual_recorder isolates the recorder interface of the Usual type
type MoqUsual_recorder struct {
	Moq *MoqUsual
}

// MoqUsual_runtime holds runtime configuration for the Usual type
type MoqUsual_runtime struct {
	ParameterIndexing struct {
		Usual   MoqUsual_Usual_paramIndexing
		Nothing MoqUsual_Nothing_paramIndexing
	}
}

// MoqUsual_Usual_adaptor adapts MoqUsual a needed by the runtime
type MoqUsual_Usual_adaptor struct {
	Moq *MoqUsual
}

// MoqUsual_Usual_params holds the params of the Usual type
type MoqUsual_Usual_params struct {
	SParam string
	BParam bool
}

// MoqUsual_Usual_paramsKey holds the map key params of the Usual type
type MoqUsual_Usual_paramsKey struct {
	Params struct {
		SParam string
		BParam bool
	}
	Hashes struct {
		SParam hash.Hash
		BParam hash.Hash
	}
}

// MoqUsual_Usual_results holds the results of the Usual type
type MoqUsual_Usual_results struct {
	SResult string
	Err     error
}

// MoqUsual_Usual_paramIndexing hold the parameter indexing runtime
// configuration for the Usual type
type MoqUsual_Usual_paramIndexing struct {
	SParam moq.ParamIndexing
	BParam moq.ParamIndexing
}

// MoqUsual_Usual_doFn defines the type of function needed when calling AndDo
// for the Usual type
type MoqUsual_Usual_doFn func(sParam string, bParam bool)

// MoqUsual_Usual_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Usual type
type MoqUsual_Usual_doReturnFn func(sParam string, bParam bool) (sResult string, err error)

// MoqUsual_Usual_recorder routes recorded function calls to the MoqUsual moq
type MoqUsual_Usual_recorder struct {
	Recorder *impl.Recorder[
		*MoqUsual_Usual_adaptor,
		MoqUsual_Usual_params,
		MoqUsual_Usual_paramsKey,
		MoqUsual_Usual_results,
	]
}

// MoqUsual_Usual_anyParams isolates the any params functions of the Usual type
type MoqUsual_Usual_anyParams struct {
	Recorder *MoqUsual_Usual_recorder
}

// MoqUsual_Nothing_adaptor adapts MoqUsual a needed by the runtime
type MoqUsual_Nothing_adaptor struct {
	Moq *MoqUsual
}

// MoqUsual_Nothing_params holds the params of the Usual type
type MoqUsual_Nothing_params struct{}

// MoqUsual_Nothing_paramsKey holds the map key params of the Usual type
type MoqUsual_Nothing_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUsual_Nothing_results holds the results of the Usual type
type MoqUsual_Nothing_results struct{}

// MoqUsual_Nothing_paramIndexing hold the parameter indexing runtime
// configuration for the Usual type
type MoqUsual_Nothing_paramIndexing struct{}

// MoqUsual_Nothing_doFn defines the type of function needed when calling AndDo
// for the Usual type
type MoqUsual_Nothing_doFn func()

// MoqUsual_Nothing_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Usual type
type MoqUsual_Nothing_doReturnFn func()

// MoqUsual_Nothing_recorder routes recorded function calls to the MoqUsual moq
type MoqUsual_Nothing_recorder struct {
	Recorder *impl.Recorder[
		*MoqUsual_Nothing_adaptor,
		MoqUsual_Nothing_params,
		MoqUsual_Nothing_paramsKey,
		MoqUsual_Nothing_results,
	]
}

// MoqUsual_Nothing_anyParams isolates the any params functions of the Usual
// type
type MoqUsual_Nothing_anyParams struct {
	Recorder *MoqUsual_Nothing_recorder
}

// NewMoqUsual creates a new moq of the Usual type
func NewMoqUsual(scene *moq.Scene, config *moq.Config) *MoqUsual {
	adaptor0 := &MoqUsual_Usual_adaptor{}
	adaptor1 := &MoqUsual_Nothing_adaptor{}
	m := &MoqUsual{
		Moq: &MoqUsual_mock{},

		Moq_Usual: impl.NewMoq[
			*MoqUsual_Usual_adaptor,
			MoqUsual_Usual_params,
			MoqUsual_Usual_paramsKey,
			MoqUsual_Usual_results,
		](scene, adaptor0, config),
		Moq_Nothing: impl.NewMoq[
			*MoqUsual_Nothing_adaptor,
			MoqUsual_Nothing_params,
			MoqUsual_Nothing_paramsKey,
			MoqUsual_Nothing_results,
		](scene, adaptor1, config),

		Runtime: MoqUsual_runtime{ParameterIndexing: struct {
			Usual   MoqUsual_Usual_paramIndexing
			Nothing MoqUsual_Nothing_paramIndexing
		}{
			Usual: MoqUsual_Usual_paramIndexing{
				SParam: moq.ParamIndexByValue,
				BParam: moq.ParamIndexByValue,
			},
			Nothing: MoqUsual_Nothing_paramIndexing{},
		}},
	}
	m.Moq.Moq = m

	adaptor0.Moq = m
	adaptor1.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Usual type
func (m *MoqUsual) Mock() *MoqUsual_mock { return m.Moq }

func (m *MoqUsual_mock) Usual(sParam string, bParam bool) (string, error) {
	m.Moq.Moq_Usual.Scene.T.Helper()
	params := MoqUsual_Usual_params{
		SParam: sParam,
		BParam: bParam,
	}

	var result0 string
	var result1 error
	if result := m.Moq.Moq_Usual.Function(params); result != nil {
		result0 = result.SResult
		result1 = result.Err
	}
	return result0, result1
}

func (m *MoqUsual_mock) Nothing() {
	m.Moq.Moq_Nothing.Scene.T.Helper()
	params := MoqUsual_Nothing_params{}

	m.Moq.Moq_Nothing.Function(params)
}

// OnCall returns the recorder implementation of the Usual type
func (m *MoqUsual) OnCall() *MoqUsual_recorder {
	return &MoqUsual_recorder{
		Moq: m,
	}
}

func (m *MoqUsual_recorder) Usual(sParam string, bParam bool) *MoqUsual_Usual_recorder {
	return &MoqUsual_Usual_recorder{
		Recorder: m.Moq.Moq_Usual.OnCall(MoqUsual_Usual_params{
			SParam: sParam,
			BParam: bParam,
		}),
	}
}

func (r *MoqUsual_Usual_recorder) Any() *MoqUsual_Usual_anyParams {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.IsAnyPermitted(true) {
		return nil
	}
	return &MoqUsual_Usual_anyParams{Recorder: r}
}

func (a *MoqUsual_Usual_anyParams) SParam() *MoqUsual_Usual_recorder {
	a.Recorder.Recorder.AnyParam(0)
	return a.Recorder
}

func (a *MoqUsual_Usual_anyParams) BParam() *MoqUsual_Usual_recorder {
	a.Recorder.Recorder.AnyParam(1)
	return a.Recorder
}

func (r *MoqUsual_Usual_recorder) Seq() *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(true, "Seq", true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Usual_recorder) NoSeq() *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(false, "NoSeq", true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Usual_recorder) ReturnResults(sResult string, err error) *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.ReturnResults(MoqUsual_Usual_results{
		SResult: sResult,
		Err:     err,
	})
	return r
}

func (r *MoqUsual_Usual_recorder) AndDo(fn MoqUsual_Usual_doFn) *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.AndDo(func(params MoqUsual_Usual_params) {
		fn(params.SParam, params.BParam)
	}, true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Usual_recorder) DoReturnResults(fn MoqUsual_Usual_doReturnFn) *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.DoReturnResults(func(params MoqUsual_Usual_params) *MoqUsual_Usual_results {
		sResult, err := fn(params.SParam, params.BParam)
		return &MoqUsual_Usual_results{
			SResult: sResult,
			Err:     err,
		}
	})
	return r
}

func (r *MoqUsual_Usual_recorder) Repeat(repeaters ...moq.Repeater) *MoqUsual_Usual_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Repeat(repeaters, true) {
		return nil
	}
	return r
}

func (*MoqUsual_Usual_adaptor) PrettyParams(params MoqUsual_Usual_params) string {
	return fmt.Sprintf("Usual(%#v, %#v)", params.SParam, params.BParam)
}

func (a *MoqUsual_Usual_adaptor) ParamsKey(params MoqUsual_Usual_params, anyParams uint64) MoqUsual_Usual_paramsKey {
	a.Moq.Moq_Usual.Scene.T.Helper()
	sParamUsed, sParamUsedHash := impl.ParamKey(
		params.SParam, 0, a.Moq.Runtime.ParameterIndexing.Usual.SParam, anyParams)
	bParamUsed, bParamUsedHash := impl.ParamKey(
		params.BParam, 1, a.Moq.Runtime.ParameterIndexing.Usual.BParam, anyParams)
	return MoqUsual_Usual_paramsKey{
		Params: struct {
			SParam string
			BParam bool
		}{
			SParam: sParamUsed,
			BParam: bParamUsed,
		},
		Hashes: struct {
			SParam hash.Hash
			BParam hash.Hash
		}{
			SParam: sParamUsedHash,
			BParam: bParamUsedHash,
		},
	}
}

func (m *MoqUsual_recorder) Nothing() *MoqUsual_Nothing_recorder {
	return &MoqUsual_Nothing_recorder{
		Recorder: m.Moq.Moq_Nothing.OnCall(MoqUsual_Nothing_params{}),
	}
}

func (r *MoqUsual_Nothing_recorder) Any() *MoqUsual_Nothing_anyParams {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.IsAnyPermitted(true) {
		return nil
	}
	return &MoqUsual_Nothing_anyParams{Recorder: r}
}

func (r *MoqUsual_Nothing_recorder) Seq() *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(true, "Seq", true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Nothing_recorder) NoSeq() *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(false, "NoSeq", true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Nothing_recorder) ReturnResults() *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.ReturnResults(MoqUsual_Nothing_results{})
	return r
}

func (r *MoqUsual_Nothing_recorder) AndDo(fn MoqUsual_Nothing_doFn) *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.AndDo(func(params MoqUsual_Nothing_params) {
		fn()
	}, true) {
		return nil
	}
	return r
}

func (r *MoqUsual_Nothing_recorder) DoReturnResults(fn MoqUsual_Nothing_doReturnFn) *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.DoReturnResults(func(params MoqUsual_Nothing_params) *MoqUsual_Nothing_results {
		fn()
		return &MoqUsual_Nothing_results{}
	})
	return r
}

func (r *MoqUsual_Nothing_recorder) Repeat(repeaters ...moq.Repeater) *MoqUsual_Nothing_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Repeat(repeaters, true) {
		return nil
	}
	return r
}

func (*MoqUsual_Nothing_adaptor) PrettyParams(params MoqUsual_Nothing_params) string {
	return "Nothing()"
}

func (a *MoqUsual_Nothing_adaptor) ParamsKey(
	params MoqUsual_Nothing_params, anyParams uint64,
) MoqUsual_Nothing_paramsKey {
	a.Moq.Moq_Nothing.Scene.T.Helper()
	return MoqUsual_Nothing_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqUsual) Reset() {
	m.Moq_Usual.Reset()
	m.Moq_Nothing.Reset()
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUsual) AssertExpectationsMet() {
	m.Moq_Usual.Scene.T.Helper()
	m.Moq_Usual.AssertExpectationsMet()
	m.Moq_Nothing.AssertExpectationsMet()
}
