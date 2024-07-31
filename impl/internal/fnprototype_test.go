package internal_test

import (
	"fmt"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/impl"
	"moqueries.org/runtime/moq"
)

type UsualFn func(sParam string, bParam bool) (sResult string, err error)

// MoqUsualFn holds the state of a moq of the UsualFn type
type MoqUsualFn struct {
	Moq *impl.Moq[
		*MoqUsualFn_adaptor,
		MoqUsualFn_params,
		MoqUsualFn_paramsKey,
		MoqUsualFn_results,
	]

	Runtime MoqUsualFn_runtime
}

// MoqUsualFn_runtime holds runtime configuration for the UsualFn type
type MoqUsualFn_runtime struct {
	ParameterIndexing MoqUsualFn_paramIndexing
}

// MoqUsualFn_adaptor adapts MoqUsualFn a needed by the runtime
type MoqUsualFn_adaptor struct {
	Moq *MoqUsualFn
}

// MoqUsualFn_params holds the params of the UsualFn type
type MoqUsualFn_params struct {
	SParam string
	BParam bool
}

// MoqUsualFn_paramsKey holds the map key params of the UsualFn type
type MoqUsualFn_paramsKey struct {
	Params struct {
		SParam string
		BParam bool
	}
	Hashes struct {
		SParam hash.Hash
		BParam hash.Hash
	}
}

// MoqUsualFn_results holds the results of the UsualFn type
type MoqUsualFn_results struct {
	SResult string
	Err     error
}

// MoqUsualFn_paramIndexing hold the parameter indexing runtime configuration
// for the UsualFn type
type MoqUsualFn_paramIndexing struct {
	SParam moq.ParamIndexing
	BParam moq.ParamIndexing
}

// MoqUsualFn_doFn defines the type of function needed when calling AndDo for
// the UsualFn type
type MoqUsualFn_doFn func(sParam string, bParam bool)

// MoqUsualFn_doReturnFn defines the type of function needed when calling
// DoReturnResults for the UsualFn type
type MoqUsualFn_doReturnFn func(sParam string, bParam bool) (sResult string, err error)

// MoqUsualFn_recorder routes recorded function calls to the MoqUsualFn moq
type MoqUsualFn_recorder struct {
	Recorder *impl.Recorder[
		*MoqUsualFn_adaptor,
		MoqUsualFn_params,
		MoqUsualFn_paramsKey,
		MoqUsualFn_results,
	]
}

// MoqUsualFn_anyParams isolates the any params functions of the UsualFn type
type MoqUsualFn_anyParams struct {
	Recorder *MoqUsualFn_recorder
}

// NewMoqUsualFn creates a new moq of the UsualFn type
func NewMoqUsualFn(scene *moq.Scene, config *moq.Config) *MoqUsualFn {
	adaptor0 := &MoqUsualFn_adaptor{}
	m := &MoqUsualFn{
		Moq: impl.NewMoq[
			*MoqUsualFn_adaptor,
			MoqUsualFn_params,
			MoqUsualFn_paramsKey,
			MoqUsualFn_results,
		](scene, adaptor0, config),

		Runtime: MoqUsualFn_runtime{ParameterIndexing: MoqUsualFn_paramIndexing{
			SParam: moq.ParamIndexByValue,
			BParam: moq.ParamIndexByValue,
		}},
	}
	adaptor0.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the UsualFn type
func (m *MoqUsualFn) Mock() UsualFn {
	return func(sParam string, bParam bool) (string, error) {
		m.Moq.Scene.T.Helper()
		params := MoqUsualFn_params{
			SParam: sParam,
			BParam: bParam,
		}

		var result0 string
		var result1 error
		if result := m.Moq.Function(params); result != nil {
			result0 = result.SResult
			result1 = result.Err
		}
		return result0, result1
	}
}

func (m *MoqUsualFn) OnCall(sParam string, bParam bool) *MoqUsualFn_recorder {
	return &MoqUsualFn_recorder{
		Recorder: m.Moq.OnCall(MoqUsualFn_params{
			SParam: sParam,
			BParam: bParam,
		}),
	}
}

func (r *MoqUsualFn_recorder) Any() *MoqUsualFn_anyParams {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.IsAnyPermitted(true) {
		return nil
	}
	return &MoqUsualFn_anyParams{Recorder: r}
}

func (a *MoqUsualFn_anyParams) SParam() *MoqUsualFn_recorder {
	a.Recorder.Recorder.AnyParam(0)
	return a.Recorder
}

func (a *MoqUsualFn_anyParams) BParam() *MoqUsualFn_recorder {
	a.Recorder.Recorder.AnyParam(1)
	return a.Recorder
}

func (r *MoqUsualFn_recorder) Seq() *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(true, "Seq", true) {
		return nil
	}
	return r
}

func (r *MoqUsualFn_recorder) NoSeq() *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Seq(false, "NoSeq", true) {
		return nil
	}
	return r
}

func (r *MoqUsualFn_recorder) ReturnResults(sResult string, err error) *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.ReturnResults(MoqUsualFn_results{
		SResult: sResult,
		Err:     err,
	})
	return r
}

func (r *MoqUsualFn_recorder) AndDo(fn MoqUsualFn_doFn) *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.AndDo(func(params MoqUsualFn_params) {
		fn(params.SParam, params.BParam)
	}, true) {
		return nil
	}
	return r
}

func (r *MoqUsualFn_recorder) DoReturnResults(fn MoqUsualFn_doReturnFn) *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	r.Recorder.DoReturnResults(func(params MoqUsualFn_params) *MoqUsualFn_results {
		sResult, err := fn(params.SParam, params.BParam)
		return &MoqUsualFn_results{
			SResult: sResult,
			Err:     err,
		}
	})
	return r
}

func (r *MoqUsualFn_recorder) Repeat(repeaters ...moq.Repeater) *MoqUsualFn_recorder {
	r.Recorder.Moq.Scene.T.Helper()
	if !r.Recorder.Repeat(repeaters, true) {
		return nil
	}
	return r
}

func (*MoqUsualFn_adaptor) PrettyParams(params MoqUsualFn_params) string {
	return fmt.Sprintf("UsualFn(%#v, %#v)", params.SParam, params.BParam)
}

func (a *MoqUsualFn_adaptor) ParamsKey(params MoqUsualFn_params, anyParams uint64) MoqUsualFn_paramsKey {
	a.Moq.Moq.Scene.T.Helper()
	sParamUsed, sParamUsedHash := impl.ParamKey(
		params.SParam, 0, a.Moq.Runtime.ParameterIndexing.SParam, anyParams)
	bParamUsed, bParamUsedHash := impl.ParamKey(
		params.BParam, 1, a.Moq.Runtime.ParameterIndexing.BParam, anyParams)
	return MoqUsualFn_paramsKey{
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

// Reset resets the state of the moq
func (m *MoqUsualFn) Reset() {
	m.Moq.Reset()
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUsualFn) AssertExpectationsMet() {
	m.Moq.Scene.T.Helper()
	m.Moq.AssertExpectationsMet()
}
