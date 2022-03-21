// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package ast_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/ast"
	"github.com/myshkin5/moqueries/hash"
	"github.com/myshkin5/moqueries/moq"
	"golang.org/x/tools/go/packages"
)

// moqLoadFn holds the state of a moq of the LoadFn type
type moqLoadFn struct {
	scene  *moq.Scene
	config moq.Config
	moq    *moqLoadFn_mock

	resultsByParams []moqLoadFn_resultsByParams

	runtime struct {
		parameterIndexing struct {
			cfg      moq.ParamIndexing
			patterns moq.ParamIndexing
		}
	}
}

// moqLoadFn_mock isolates the mock interface of the LoadFn type
type moqLoadFn_mock struct {
	moq *moqLoadFn
}

// moqLoadFn_params holds the params of the LoadFn type
type moqLoadFn_params struct {
	cfg      *packages.Config
	patterns []string
}

// moqLoadFn_paramsKey holds the map key params of the LoadFn type
type moqLoadFn_paramsKey struct {
	params struct{ cfg *packages.Config }
	hashes struct {
		cfg      hash.Hash
		patterns hash.Hash
	}
}

// moqLoadFn_resultsByParams contains the results for a given set of parameters for the LoadFn type
type moqLoadFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqLoadFn_paramsKey]*moqLoadFn_results
}

// moqLoadFn_doFn defines the type of function needed when calling andDo for the LoadFn type
type moqLoadFn_doFn func(cfg *packages.Config, patterns ...string)

// moqLoadFn_doReturnFn defines the type of function needed when calling doReturnResults for the LoadFn type
type moqLoadFn_doReturnFn func(cfg *packages.Config, patterns ...string) ([]*packages.Package, error)

// moqLoadFn_results holds the results of the LoadFn type
type moqLoadFn_results struct {
	params  moqLoadFn_params
	results []struct {
		values *struct {
			result1 []*packages.Package
			result2 error
		}
		sequence   uint32
		doFn       moqLoadFn_doFn
		doReturnFn moqLoadFn_doReturnFn
	}
	index  uint32
	repeat *moq.RepeatVal
}

// moqLoadFn_fnRecorder routes recorded function calls to the moqLoadFn moq
type moqLoadFn_fnRecorder struct {
	params    moqLoadFn_params
	anyParams uint64
	sequence  bool
	results   *moqLoadFn_results
	moq       *moqLoadFn
}

// moqLoadFn_anyParams isolates the any params functions of the LoadFn type
type moqLoadFn_anyParams struct {
	recorder *moqLoadFn_fnRecorder
}

// newMoqLoadFn creates a new moq of the LoadFn type
func newMoqLoadFn(scene *moq.Scene, config *moq.Config) *moqLoadFn {
	if config == nil {
		config = &moq.Config{}
	}
	m := &moqLoadFn{
		scene:  scene,
		config: *config,
		moq:    &moqLoadFn_mock{},

		runtime: struct {
			parameterIndexing struct {
				cfg      moq.ParamIndexing
				patterns moq.ParamIndexing
			}
		}{parameterIndexing: struct {
			cfg      moq.ParamIndexing
			patterns moq.ParamIndexing
		}{
			cfg:      moq.ParamIndexByHash,
			patterns: moq.ParamIndexByHash,
		}},
	}
	m.moq.moq = m

	scene.AddMoq(m)
	return m
}

// mock returns the moq implementation of the LoadFn type
func (m *moqLoadFn) mock() ast.LoadFn {
	return func(cfg *packages.Config, patterns ...string) ([]*packages.Package, error) {
		moq := &moqLoadFn_mock{moq: m}
		return moq.fn(cfg, patterns...)
	}
}

func (m *moqLoadFn_mock) fn(cfg *packages.Config, patterns ...string) (result1 []*packages.Package, result2 error) {
	m.moq.scene.T.Helper()
	params := moqLoadFn_params{
		cfg:      cfg,
		patterns: patterns,
	}
	var results *moqLoadFn_results
	for _, resultsByParams := range m.moq.resultsByParams {
		paramsKey := m.moq.paramsKey(params, resultsByParams.anyParams)
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
		result.doFn(cfg, patterns...)
	}

	if result.values != nil {
		result1 = result.values.result1
		result2 = result.values.result2
	}
	if result.doReturnFn != nil {
		result1, result2 = result.doReturnFn(cfg, patterns...)
	}
	return
}

func (m *moqLoadFn) onCall(cfg *packages.Config, patterns ...string) *moqLoadFn_fnRecorder {
	return &moqLoadFn_fnRecorder{
		params: moqLoadFn_params{
			cfg:      cfg,
			patterns: patterns,
		},
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		moq:      m,
	}
}

func (r *moqLoadFn_fnRecorder) any() *moqLoadFn_anyParams {
	r.moq.scene.T.Helper()
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	return &moqLoadFn_anyParams{recorder: r}
}

func (a *moqLoadFn_anyParams) cfg() *moqLoadFn_fnRecorder {
	a.recorder.anyParams |= 1 << 0
	return a.recorder
}

func (a *moqLoadFn_anyParams) patterns() *moqLoadFn_fnRecorder {
	a.recorder.anyParams |= 1 << 1
	return a.recorder
}

func (r *moqLoadFn_fnRecorder) seq() *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqLoadFn_fnRecorder) noSeq() *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqLoadFn_fnRecorder) returnResults(result1 []*packages.Package, result2 error) *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 []*packages.Package
			result2 error
		}
		sequence   uint32
		doFn       moqLoadFn_doFn
		doReturnFn moqLoadFn_doReturnFn
	}{
		values: &struct {
			result1 []*packages.Package
			result2 error
		}{
			result1: result1,
			result2: result2,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqLoadFn_fnRecorder) andDo(fn moqLoadFn_doFn) *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqLoadFn_fnRecorder) doReturnResults(fn moqLoadFn_doReturnFn) *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 []*packages.Package
			result2 error
		}
		sequence   uint32
		doFn       moqLoadFn_doFn
		doReturnFn moqLoadFn_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqLoadFn_fnRecorder) findResults() {
	if r.results != nil {
		r.results.repeat.Increment(r.moq.scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.anyParams)
	insertAt := -1
	var results *moqLoadFn_resultsByParams
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
		results = &moqLoadFn_resultsByParams{
			anyCount:  anyCount,
			anyParams: r.anyParams,
			results:   map[moqLoadFn_paramsKey]*moqLoadFn_results{},
		}
		r.moq.resultsByParams = append(r.moq.resultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams) {
			copy(r.moq.resultsByParams[insertAt+1:], r.moq.resultsByParams[insertAt:0])
			r.moq.resultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.moq.paramsKey(r.params, r.anyParams)

	var ok bool
	r.results, ok = results.results[paramsKey]
	if !ok {
		r.results = &moqLoadFn_results{
			params:  r.params,
			results: nil,
			index:   0,
			repeat:  &moq.RepeatVal{},
		}
		results.results[paramsKey] = r.results
	}

	r.results.repeat.Increment(r.moq.scene.T)
}

func (r *moqLoadFn_fnRecorder) repeat(repeaters ...moq.Repeater) *moqLoadFn_fnRecorder {
	r.moq.scene.T.Helper()
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
					result1 []*packages.Package
					result2 error
				}
				sequence   uint32
				doFn       moqLoadFn_doFn
				doReturnFn moqLoadFn_doReturnFn
			}{
				values: &struct {
					result1 []*packages.Package
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

func (m *moqLoadFn) paramsKey(params moqLoadFn_params, anyParams uint64) moqLoadFn_paramsKey {
	var cfgUsed *packages.Config
	var cfgUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.runtime.parameterIndexing.cfg == moq.ParamIndexByValue {
			cfgUsed = params.cfg
		} else {
			cfgUsedHash = hash.DeepHash(params.cfg)
		}
	}
	var patternsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.runtime.parameterIndexing.patterns == moq.ParamIndexByValue {
			m.scene.T.Fatalf("The patterns parameter can't be indexed by value")
		}
		patternsUsedHash = hash.DeepHash(params.patterns)
	}
	return moqLoadFn_paramsKey{
		params: struct{ cfg *packages.Config }{
			cfg: cfgUsed,
		},
		hashes: struct {
			cfg      hash.Hash
			patterns hash.Hash
		}{
			cfg:      cfgUsedHash,
			patterns: patternsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *moqLoadFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *moqLoadFn) AssertExpectationsMet() {
	m.scene.T.Helper()
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := results.repeat.MinTimes - int(atomic.LoadUint32(&results.index))
			if missing > 0 {
				m.scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}