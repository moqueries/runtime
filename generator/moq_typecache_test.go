// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package generator_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/dave/dst"
	"github.com/myshkin5/moqueries/hash"
	"github.com/myshkin5/moqueries/moq"
)

// moqTypeCache holds the state of a moq of the TypeCache type
type moqTypeCache struct {
	scene                        *moq.Scene
	config                       moq.Config
	resultsByParams_Type         []moqTypeCache_Type_resultsByParams
	resultsByParams_IsComparable []moqTypeCache_IsComparable_resultsByParams
}

// moqTypeCache_mock isolates the mock interface of the TypeCache type
type moqTypeCache_mock struct {
	moq *moqTypeCache
}

// moqTypeCache_recorder isolates the recorder interface of the TypeCache type
type moqTypeCache_recorder struct {
	moq *moqTypeCache
}

// moqTypeCache_Type_params holds the params of the TypeCache type
type moqTypeCache_Type_params struct {
	id            dst.Ident
	loadTestTypes bool
}

// moqTypeCache_Type_paramsKey holds the map key params of the TypeCache type
type moqTypeCache_Type_paramsKey struct {
	id            hash.Hash
	loadTestTypes bool
}

// moqTypeCache_Type_resultsByParams contains the results for a given set of parameters for the TypeCache type
type moqTypeCache_Type_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqTypeCache_Type_paramsKey]*moqTypeCache_Type_results
}

// moqTypeCache_Type_doFn defines the type of function needed when calling andDo for the TypeCache type
type moqTypeCache_Type_doFn func(id dst.Ident, loadTestTypes bool)

// moqTypeCache_Type_doReturnFn defines the type of function needed when calling doReturnResults for the TypeCache type
type moqTypeCache_Type_doReturnFn func(id dst.Ident, loadTestTypes bool) (*dst.TypeSpec, string, error)

// moqTypeCache_Type_results holds the results of the TypeCache type
type moqTypeCache_Type_results struct {
	params  moqTypeCache_Type_params
	results []struct {
		values *struct {
			result1 *dst.TypeSpec
			result2 string
			result3 error
		}
		sequence   uint32
		doFn       moqTypeCache_Type_doFn
		doReturnFn moqTypeCache_Type_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// moqTypeCache_Type_fnRecorder routes recorded function calls to the moqTypeCache moq
type moqTypeCache_Type_fnRecorder struct {
	params    moqTypeCache_Type_params
	paramsKey moqTypeCache_Type_paramsKey
	anyParams uint64
	sequence  bool
	results   *moqTypeCache_Type_results
	moq       *moqTypeCache
}

// moqTypeCache_IsComparable_params holds the params of the TypeCache type
type moqTypeCache_IsComparable_params struct{ expr dst.Expr }

// moqTypeCache_IsComparable_paramsKey holds the map key params of the TypeCache type
type moqTypeCache_IsComparable_paramsKey struct{ expr hash.Hash }

// moqTypeCache_IsComparable_resultsByParams contains the results for a given set of parameters for the TypeCache type
type moqTypeCache_IsComparable_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqTypeCache_IsComparable_paramsKey]*moqTypeCache_IsComparable_results
}

// moqTypeCache_IsComparable_doFn defines the type of function needed when calling andDo for the TypeCache type
type moqTypeCache_IsComparable_doFn func(expr dst.Expr)

// moqTypeCache_IsComparable_doReturnFn defines the type of function needed when calling doReturnResults for the TypeCache type
type moqTypeCache_IsComparable_doReturnFn func(expr dst.Expr) (bool, error)

// moqTypeCache_IsComparable_results holds the results of the TypeCache type
type moqTypeCache_IsComparable_results struct {
	params  moqTypeCache_IsComparable_params
	results []struct {
		values *struct {
			result1 bool
			result2 error
		}
		sequence   uint32
		doFn       moqTypeCache_IsComparable_doFn
		doReturnFn moqTypeCache_IsComparable_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// moqTypeCache_IsComparable_fnRecorder routes recorded function calls to the moqTypeCache moq
type moqTypeCache_IsComparable_fnRecorder struct {
	params    moqTypeCache_IsComparable_params
	paramsKey moqTypeCache_IsComparable_paramsKey
	anyParams uint64
	sequence  bool
	results   *moqTypeCache_IsComparable_results
	moq       *moqTypeCache
}

// newMoqTypeCache creates a new moq of the TypeCache type
func newMoqTypeCache(scene *moq.Scene, config *moq.Config) *moqTypeCache {
	if config == nil {
		config = &moq.Config{}
	}
	m := &moqTypeCache{
		scene:  scene,
		config: *config,
	}
	scene.AddMoq(m)
	return m
}

// mock returns the mock implementation of the TypeCache type
func (m *moqTypeCache) mock() *moqTypeCache_mock {
	return &moqTypeCache_mock{
		moq: m,
	}
}

func (m *moqTypeCache_mock) Type(id dst.Ident, loadTestTypes bool) (result1 *dst.TypeSpec, result2 string, result3 error) {
	params := moqTypeCache_Type_params{
		id:            id,
		loadTestTypes: loadTestTypes,
	}
	var results *moqTypeCache_Type_results
	for _, resultsByParams := range m.moq.resultsByParams_Type {
		var idUsed hash.Hash
		if resultsByParams.anyParams&(1<<0) == 0 {
			idUsed = hash.DeepHash(id)
		}
		var loadTestTypesUsed bool
		if resultsByParams.anyParams&(1<<1) == 0 {
			loadTestTypesUsed = loadTestTypes
		}
		paramsKey := moqTypeCache_Type_paramsKey{
			id:            idUsed,
			loadTestTypes: loadTestTypesUsed,
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
		result.doFn(id, loadTestTypes)
	}

	if result.values != nil {
		result1 = result.values.result1
		result2 = result.values.result2
		result3 = result.values.result3
	}
	if result.doReturnFn != nil {
		result1, result2, result3 = result.doReturnFn(id, loadTestTypes)
	}
	return
}

func (m *moqTypeCache_mock) IsComparable(expr dst.Expr) (result1 bool, result2 error) {
	params := moqTypeCache_IsComparable_params{
		expr: expr,
	}
	var results *moqTypeCache_IsComparable_results
	for _, resultsByParams := range m.moq.resultsByParams_IsComparable {
		var exprUsed hash.Hash
		if resultsByParams.anyParams&(1<<0) == 0 {
			exprUsed = hash.DeepHash(expr)
		}
		paramsKey := moqTypeCache_IsComparable_paramsKey{
			expr: exprUsed,
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
		result.doFn(expr)
	}

	if result.values != nil {
		result1 = result.values.result1
		result2 = result.values.result2
	}
	if result.doReturnFn != nil {
		result1, result2 = result.doReturnFn(expr)
	}
	return
}

// onCall returns the recorder implementation of the TypeCache type
func (m *moqTypeCache) onCall() *moqTypeCache_recorder {
	return &moqTypeCache_recorder{
		moq: m,
	}
}

func (m *moqTypeCache_recorder) Type(id dst.Ident, loadTestTypes bool) *moqTypeCache_Type_fnRecorder {
	return &moqTypeCache_Type_fnRecorder{
		params: moqTypeCache_Type_params{
			id:            id,
			loadTestTypes: loadTestTypes,
		},
		paramsKey: moqTypeCache_Type_paramsKey{
			id:            hash.DeepHash(id),
			loadTestTypes: loadTestTypes,
		},
		sequence: m.moq.config.Sequence == moq.SeqDefaultOn,
		moq:      m.moq,
	}
}

func (r *moqTypeCache_Type_fnRecorder) anyId() *moqTypeCache_Type_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *moqTypeCache_Type_fnRecorder) anyLoadTestTypes() *moqTypeCache_Type_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *moqTypeCache_Type_fnRecorder) seq() *moqTypeCache_Type_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqTypeCache_Type_fnRecorder) noSeq() *moqTypeCache_Type_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqTypeCache_Type_fnRecorder) returnResults(result1 *dst.TypeSpec, result2 string, result3 error) *moqTypeCache_Type_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 *dst.TypeSpec
			result2 string
			result3 error
		}
		sequence   uint32
		doFn       moqTypeCache_Type_doFn
		doReturnFn moqTypeCache_Type_doReturnFn
	}{
		values: &struct {
			result1 *dst.TypeSpec
			result2 string
			result3 error
		}{
			result1: result1,
			result2: result2,
			result3: result3,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqTypeCache_Type_fnRecorder) andDo(fn moqTypeCache_Type_doFn) *moqTypeCache_Type_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqTypeCache_Type_fnRecorder) doReturnResults(fn moqTypeCache_Type_doReturnFn) *moqTypeCache_Type_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 *dst.TypeSpec
			result2 string
			result3 error
		}
		sequence   uint32
		doFn       moqTypeCache_Type_doFn
		doReturnFn moqTypeCache_Type_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqTypeCache_Type_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *moqTypeCache_Type_resultsByParams
		for n, res := range r.moq.resultsByParams_Type {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &moqTypeCache_Type_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[moqTypeCache_Type_paramsKey]*moqTypeCache_Type_results{},
			}
			r.moq.resultsByParams_Type = append(r.moq.resultsByParams_Type, *results)
			if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams_Type) {
				copy(r.moq.resultsByParams_Type[insertAt+1:], r.moq.resultsByParams_Type[insertAt:0])
				r.moq.resultsByParams_Type[insertAt] = *results
			}
		}

		var idUsed hash.Hash
		if r.anyParams&(1<<0) == 0 {
			idUsed = r.paramsKey.id
		}
		var loadTestTypesUsed bool
		if r.anyParams&(1<<1) == 0 {
			loadTestTypesUsed = r.paramsKey.loadTestTypes
		}
		paramsKey := moqTypeCache_Type_paramsKey{
			id:            idUsed,
			loadTestTypes: loadTestTypesUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &moqTypeCache_Type_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *moqTypeCache_Type_fnRecorder) times(count int) *moqTypeCache_Type_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					result1 *dst.TypeSpec
					result2 string
					result3 error
				}
				sequence   uint32
				doFn       moqTypeCache_Type_doFn
				doReturnFn moqTypeCache_Type_doReturnFn
			}{
				values: &struct {
					result1 *dst.TypeSpec
					result2 string
					result3 error
				}{
					result1: last.values.result1,
					result2: last.values.result2,
					result3: last.values.result3,
				},
				sequence: r.moq.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *moqTypeCache_Type_fnRecorder) anyTimes() {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *moqTypeCache_recorder) IsComparable(expr dst.Expr) *moqTypeCache_IsComparable_fnRecorder {
	return &moqTypeCache_IsComparable_fnRecorder{
		params: moqTypeCache_IsComparable_params{
			expr: expr,
		},
		paramsKey: moqTypeCache_IsComparable_paramsKey{
			expr: hash.DeepHash(expr),
		},
		sequence: m.moq.config.Sequence == moq.SeqDefaultOn,
		moq:      m.moq,
	}
}

func (r *moqTypeCache_IsComparable_fnRecorder) anyExpr() *moqTypeCache_IsComparable_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) seq() *moqTypeCache_IsComparable_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) noSeq() *moqTypeCache_IsComparable_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) returnResults(result1 bool, result2 error) *moqTypeCache_IsComparable_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 bool
			result2 error
		}
		sequence   uint32
		doFn       moqTypeCache_IsComparable_doFn
		doReturnFn moqTypeCache_IsComparable_doReturnFn
	}{
		values: &struct {
			result1 bool
			result2 error
		}{
			result1: result1,
			result2: result2,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) andDo(fn moqTypeCache_IsComparable_doFn) *moqTypeCache_IsComparable_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) doReturnResults(fn moqTypeCache_IsComparable_doReturnFn) *moqTypeCache_IsComparable_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			result1 bool
			result2 error
		}
		sequence   uint32
		doFn       moqTypeCache_IsComparable_doFn
		doReturnFn moqTypeCache_IsComparable_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqTypeCache_IsComparable_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *moqTypeCache_IsComparable_resultsByParams
		for n, res := range r.moq.resultsByParams_IsComparable {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &moqTypeCache_IsComparable_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[moqTypeCache_IsComparable_paramsKey]*moqTypeCache_IsComparable_results{},
			}
			r.moq.resultsByParams_IsComparable = append(r.moq.resultsByParams_IsComparable, *results)
			if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams_IsComparable) {
				copy(r.moq.resultsByParams_IsComparable[insertAt+1:], r.moq.resultsByParams_IsComparable[insertAt:0])
				r.moq.resultsByParams_IsComparable[insertAt] = *results
			}
		}

		var exprUsed hash.Hash
		if r.anyParams&(1<<0) == 0 {
			exprUsed = r.paramsKey.expr
		}
		paramsKey := moqTypeCache_IsComparable_paramsKey{
			expr: exprUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &moqTypeCache_IsComparable_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *moqTypeCache_IsComparable_fnRecorder) times(count int) *moqTypeCache_IsComparable_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					result1 bool
					result2 error
				}
				sequence   uint32
				doFn       moqTypeCache_IsComparable_doFn
				doReturnFn moqTypeCache_IsComparable_doReturnFn
			}{
				values: &struct {
					result1 bool
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

func (r *moqTypeCache_IsComparable_fnRecorder) anyTimes() {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the moq
func (m *moqTypeCache) Reset() { m.resultsByParams_Type = nil; m.resultsByParams_IsComparable = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *moqTypeCache) AssertExpectationsMet() {
	for _, res := range m.resultsByParams_Type {
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
	for _, res := range m.resultsByParams_IsComparable {
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