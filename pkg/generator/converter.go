package generator

import (
	"fmt"
	"go/token"
	"strings"

	"github.com/dave/dst"

	. "github.com/myshkin5/moqueries/pkg/ast"
	"github.com/myshkin5/moqueries/pkg/logs"
)

const (
	moqueriesPkg  = "github.com/myshkin5/moqueries"
	hashPkg       = moqueriesPkg + "/pkg/hash"
	moqPkg        = moqueriesPkg + "/pkg/moq"
	syncAtomicPkg = "sync/atomic"

	intType        = "int"
	mockConfigType = "MockConfig"
	moqTType       = "MoqT"
	sceneType      = "Scene"

	anyTimesIdent         = "anyTimes"
	configIdent           = "config"
	countIdent            = "count"
	expectationIdent      = "Expectation"
	iIdent                = "i"
	indexIdent            = "index"
	lastIdent             = "last"
	missingIdent          = "missing"
	mockIdent             = "mock"
	mockReceiverIdent     = "m"
	nilIdent              = "nil"
	okIdent               = "ok"
	paramsIdent           = "params"
	paramsKeyIdent        = "paramsKey"
	recorderIdent         = "recorder"
	recorderReceiverIdent = "r"
	resultsByParamsIdent  = "resultsByParams"
	resultIdent           = "result"
	resultsIdent          = "results"
	sceneIdent            = "scene"
	strictIdent           = "Strict"

	anyTimesFnName = "anyTimes"
	assertFnName   = "AssertExpectationsMet"
	errorfFnName   = "Errorf"
	fatalfFnName   = "Fatalf"
	fnFnName       = "fn"
	lenFnName      = "len"
	mockFnName     = "mock"
	onCallFnName   = "onCall"
	resetFnName    = "Reset"
	returnFnName   = "returnResults"
	timesFnName    = "times"

	fnRecorderSuffix = "fnRecorder"
	paramPrefix      = "param"
	resultMgrSuffix  = "resultMgr"
	resultPrefix     = "result"
)

// Converter converts various interface and function types to AST structs to
// build a mock
type Converter struct {
	isExported bool
}

// NewConverter creates a new Converter
func NewConverter(isExported bool) *Converter {
	return &Converter{
		isExported: isExported,
	}
}

// Func holds on to function related data
type Func struct {
	Name    string
	Params  *dst.FieldList
	Results *dst.FieldList
}

// BaseStruct generates the base structure used to store the mock's state
func (c *Converter) BaseStruct(typeSpec *dst.TypeSpec, funcs []Func) *dst.GenDecl {
	fields := []*dst.Field{
		Field(Star(IdPath(sceneType, moqPkg))).Names(Id(c.export(sceneIdent))).Obj,
		Field(IdPath(mockConfigType, moqPkg)).Names(Id(c.export(configIdent))).Obj,
	}

	mName := c.mockName(typeSpec.Name.Name)
	for _, fn := range funcs {
		typePrefix := mName
		fieldSuffix := ""
		if _, ok := typeSpec.Type.(*dst.InterfaceType); ok {
			typePrefix = fmt.Sprintf("%s_%s", mName, fn.Name)
			fieldSuffix = "_" + fn.Name
		}
		fields = append(fields,
			Field(MapType(Id(c.export(fmt.Sprintf("%s_%s", typePrefix, paramsKeyIdent)))).
				Value(Star(Id(fmt.Sprintf("%s_%s", typePrefix, resultMgrSuffix)))).Obj).
				Names(Id(c.export(resultsByParamsIdent+fieldSuffix))).Obj)
	}

	return Struct(mName).Fields(fields...).Decs(genDeclDec(
		"// %s holds the state of a mock of the %s type",
		mName, typeSpec.Name.Name)).Obj
}

// IsolationStruct generates a struct used to isolate an interface for the mock
func (c *Converter) IsolationStruct(typeName, suffix string) (structDecl *dst.GenDecl) {
	mName := c.mockName(typeName)
	iName := fmt.Sprintf("%s_%s", mName, suffix)

	return Struct(iName).
		Fields(Field(Star(Id(mName))).Names(Id(c.export(mockIdent))).Obj).
		Decs(genDeclDec("// %s isolates the %s interface of the %s type",
			iName, suffix, typeName)).Obj
}

// MethodStructs generates a structure for storing a set of parameters or
// a set of results for a method invocation of a mock
func (c *Converter) MethodStructs(typeSpec *dst.TypeSpec, fn Func) []dst.Decl {
	prefix := c.mockName(typeSpec.Name.Name)
	if _, ok := typeSpec.Type.(*dst.InterfaceType); ok {
		prefix = fmt.Sprintf("%s_%s", prefix, fn.Name)
	}

	return []dst.Decl{
		c.methodStruct(typeSpec.Name.Name, prefix, paramsIdent, fn.Params),
		c.methodStruct(typeSpec.Name.Name, prefix, paramsKeyIdent, fn.Params),
		c.resultMgrStruct(typeSpec.Name.Name, prefix),
		c.methodStruct(typeSpec.Name.Name, prefix, resultsIdent, fn.Results),
		c.fnRecorderStruct(typeSpec.Name.Name, prefix),
	}
}

// NewFunc generates a function for constructing a mock
func (c *Converter) NewFunc(typeSpec *dst.TypeSpec) (funcDecl *dst.FuncDecl) {
	fnName := c.export("newMock" + typeSpec.Name.Name)
	mName := c.mockName(typeSpec.Name.Name)
	return Fn(fnName).
		Params(
			Field(Star(IdPath(sceneType, moqPkg))).Names(Id(sceneIdent)).Obj,
			Field(Star(IdPath(mockConfigType, moqPkg))).Names(Id(configIdent)).Obj,
		).
		Results(Field(Star(Id(mName))).Obj).
		Body(
			If(Bin(Id(configIdent)).Op(token.EQL).Y(Id(nilIdent)).Obj).Body(
				Assign(Id(configIdent)).Tok(token.ASSIGN).Rhs(Un(token.AND,
					Comp(IdPath(mockConfigType, moqPkg)).Obj)).Obj).Obj,
			Assign(Id(mockReceiverIdent)).Tok(token.DEFINE).
				Rhs(Un(token.AND, Comp(Id(mName)).Elts(
					Key(Id(c.export(sceneIdent))).
						Value(Id(sceneIdent)).Decs(kvExprDec(dst.None)).Obj,
					Key(Id(c.export(configIdent))).
						Value(Star(Id(configIdent))).Decs(kvExprDec(dst.None)).Obj,
				).Decs(litDec()).Obj)).Obj,
			Expr(Call(Sel(Id(mockReceiverIdent)).Dot(Id(resetFnName)).Obj).Obj),
			Expr(Call(Sel(Id(sceneIdent)).Dot(Id("AddMock")).Obj).
				Args(Id(mockReceiverIdent)).Obj),
			Return(Id(mockReceiverIdent)),
		).
		Decs(fnDeclDec("// %s creates a new mock of the %s type",
			fnName, typeSpec.Name.Name)).Obj
}

// IsolationAccessor generates a function to access an isolation interface
func (c *Converter) IsolationAccessor(typeName, suffix, fnName string) (funcDecl *dst.FuncDecl) {
	fnName = c.export(fnName)
	mName := c.mockName(typeName)
	iName := fmt.Sprintf("%s_%s", mName, suffix)
	return Fn(fnName).
		Recv(Field(Star(Id(mName))).Names(Id(mockReceiverIdent)).Obj).
		Results(Field(Star(Id(iName))).Obj).
		Body(Return(Un(token.AND, Comp(Id(iName)).Elts(
			Key(Id(c.export(mockIdent))).Value(Id(mockReceiverIdent)).
				Decs(kvExprDec(dst.None)).Obj).Decs(litDec()).Obj,
		))).
		Decs(fnDeclDec("// %s returns the %s implementation of the %s type",
			fnName, suffix, typeName)).Obj
}

// FuncClosure generates a mock implementation of function type wrapped in a
// closure
func (c *Converter) FuncClosure(typeName, pkgPath string, fn Func) (
	funcDecl *dst.FuncDecl,
) {
	mName := c.mockName(typeName)
	ellipsis := false
	if len(fn.Params.List) > 0 {
		if _, ok := fn.Params.List[len(fn.Params.List)-1].Type.(*dst.Ellipsis); ok {
			ellipsis = true
		}
	}
	fnLitCall := Call(Sel(Id(mockIdent)).Dot(Id(c.export(fnFnName))).Obj).
		Args(passthroughFields(fn.Params)...).
		Ellipsis(ellipsis).Obj
	var fnLitRetStmt dst.Stmt
	fnLitRetStmt = Return(fnLitCall)
	if fn.Results == nil {
		fnLitRetStmt = Expr(fnLitCall)
	}

	return Fn(c.export(mockFnName)).
		Recv(Field(Star(Id(mName))).Names(Id(mockReceiverIdent)).Obj).
		Results(Field(IdPath(typeName, pkgPath)).Obj).
		Body(Return(FnLit(FnType(dst.Clone(fn.Params).(*dst.FieldList)).
			Results(cloneNilableFieldList(fn.Results)).Obj).
			Body(Assign(Id(mockIdent)).
				Tok(token.DEFINE).
				Rhs(Un(
					token.AND,
					Comp(Idf("%s_%s", mName, mockIdent)).
						Elts(Key(Id(c.export(mockIdent))).
							Value(Id(mockReceiverIdent)).Obj).Obj,
				)).Obj,
				fnLitRetStmt,
			).Obj)).
		Decs(fnDeclDec("// %s returns the %s implementation of the %s type",
			c.export(mockFnName), mockIdent, typeName)).Obj
}

// MockMethod generates a mock implementation of a method
func (c *Converter) MockMethod(typeName string, fn Func) *dst.FuncDecl {
	mName := c.mockName(typeName)
	recv := fmt.Sprintf("%s_%s", mName, mockIdent)

	fnName := fn.Name
	fieldSuffix := "_" + fn.Name
	typePrefix := fmt.Sprintf("%s_%s", mName, fn.Name)
	if fnName == "" {
		fnName = c.export(fnFnName)
		fieldSuffix = ""
		typePrefix = mName
	}

	return Fn(fnName).
		Recv(Field(Star(Id(recv))).Names(Id(mockReceiverIdent)).Obj).
		ParamFieldList(cloneAndNameUnnamed(paramPrefix, fn.Params)).
		ResultFieldList(cloneAndNameUnnamed(resultPrefix, fn.Results)).
		Body(c.mockFunc(typePrefix, fieldSuffix, fn)...).
		Decs(stdFuncDec()).Obj
}

// RecorderMethods generates a recorder implementation of a method and
// associated return method
func (c *Converter) RecorderMethods(typeName string, fn Func) (funcDecls []dst.Decl) {
	decls := []dst.Decl{
		c.recorderFn(typeName, fn),
	}

	decls = append(decls,
		c.recorderReturnFn(typeName, fn),
		c.recorderTimesFn(typeName, fn),
		c.recorderAnyTimesFn(typeName, fn),
	)

	return decls
}

// ResetMethod generates a method to reset the mock's state
func (c *Converter) ResetMethod(typeSpec *dst.TypeSpec, funcs []Func) (funcDecl *dst.FuncDecl) {
	mName := c.mockName(typeSpec.Name.Name)

	var stmts []dst.Stmt
	for _, fn := range funcs {
		typePrefix := mName
		fieldSuffix := ""
		if _, ok := typeSpec.Type.(*dst.InterfaceType); ok {
			typePrefix = fmt.Sprintf("%s_%s", mName, fn.Name)
			fieldSuffix = "_" + fn.Name
		}

		pName := fmt.Sprintf("%s_%s", typePrefix, paramsKeyIdent)

		stmts = append(stmts, Assign(Sel(Id(mockReceiverIdent)).
			Dot(Id(c.export(resultsByParamsIdent+fieldSuffix))).Obj).
			Tok(token.ASSIGN).
			Rhs(Comp(MapType(Id(pName)).
				Value(Star(Idf("%s_%s", typePrefix, resultMgrSuffix))).Obj).Obj).Obj)
	}

	return Fn(resetFnName).
		Recv(Field(Star(Id(mName))).Names(Id(mockReceiverIdent)).Obj).
		Body(stmts...).
		Decs(fnDeclDec("// %s resets the state of the mock", resetFnName)).Obj
}

// AssertMethod generates a method to assert all expectations are met
func (c *Converter) AssertMethod(typeSpec *dst.TypeSpec, funcs []Func) (funcDecl *dst.FuncDecl) {
	mName := c.mockName(typeSpec.Name.Name)

	var stmts []dst.Stmt
	for _, fn := range funcs {
		fieldSuffix := ""
		if _, ok := typeSpec.Type.(*dst.InterfaceType); ok {
			fieldSuffix = "_" + fn.Name
		}

		stmts = append(stmts, Range(Sel(Id(mockReceiverIdent)).
			Dot(Id(c.export(resultsByParamsIdent+fieldSuffix))).Obj).
			Key(Id("_")).
			Value(Id(resultsIdent)).
			Tok(token.DEFINE).
			Body(
				Assign(Id(missingIdent)).
					Tok(token.DEFINE).
					Rhs(Bin(Call(Id(lenFnName)).Args(Sel(Id(resultsIdent)).
						Dot(Id(c.export(resultsIdent))).Obj).Obj).
						Op(token.SUB).
						Y(Call(Id(intType)).Args(
							Call(IdPath("LoadUint32", syncAtomicPkg)).Args(Un(
								token.AND,
								Sel(Id(resultsIdent)).
									Dot(Id(c.export(indexIdent))).Obj)).Obj).Obj).Obj).Obj,
				If(Bin(Bin(Id(missingIdent)).Op(token.EQL).
					Y(LitInt(1)).Obj).
					Op(token.LAND).
					Y(Bin(Sel(Id(resultsIdent)).
						Dot(Id(c.export(anyTimesIdent))).Obj).
						Op(token.EQL).
						Y(Id("true")).Obj).Obj).
					Body(Continue()).Obj,
				If(Bin(Id(missingIdent)).Op(token.GTR).Y(LitInt(0)).Obj).
					Body(
						Expr(Call(Sel(Sel(Sel(Id(mockReceiverIdent)).
							Dot(Id(c.export(sceneIdent))).Obj).
							Dot(Id(moqTType)).Obj).Dot(Id(errorfFnName)).Obj).
							Args(
								LitString("Expected %d additional call(s) with parameters %#v"),
								Id(missingIdent),
								Sel(Id(resultsIdent)).Dot(Id(c.export(paramsIdent))).Obj).Obj),
					).Obj,
			).Obj)
	}

	return Fn(assertFnName).
		Recv(Field(Star(Id(mName))).Names(Id(mockReceiverIdent)).Obj).
		Body(stmts...).
		Decs(fnDeclDec("// %s asserts that all expectations have been met",
			assertFnName)).Obj
}

func isComparable(expr dst.Expr) bool {
	// TODO this logic needs to be expanded -- also should check structs recursively
	switch expr.(type) {
	case *dst.ArrayType, *dst.MapType, *dst.Ellipsis:
		return false
	}

	return true
}

func (c *Converter) methodStruct(
	typeName, prefix, label string, fieldList *dst.FieldList,
) *dst.GenDecl {
	unnamedPrefix, comparable := labelDirection(label)
	fieldList = cloneNilableFieldList(fieldList)

	if fieldList == nil {
		// Result field lists can be nil (rather than containing an empty
		// list). Struct field lists cannot be nil.
		fieldList = FieldList()
	} else {
		for n, f := range fieldList.List {
			if len(f.Names) == 0 {
				f.Names = []*dst.Ident{Idf("%s%d", unnamedPrefix, n+1)}
			}

			for nn := range f.Names {
				f.Names[nn] = Id(c.export(f.Names[nn].Name))
			}

			f.Type = comparableType(comparable, f.Type)
		}
	}

	goDocDesc := label
	if label == paramsKeyIdent {
		goDocDesc = "map key params"
	}

	structName := fmt.Sprintf("%s_%s", prefix, label)
	return Struct(structName).FieldList(fieldList).Decs(genDeclDec(
		"// %s holds the %s of the %s type",
		structName, goDocDesc, typeName)).Obj
}

func comparableType(needComparable bool, typ dst.Expr) dst.Expr {
	if needComparable && !isComparable(typ) {
		// Non-comparable params are represented as a deep hash
		return IdPath("Hash", hashPkg)
	} else if ellipsis, ok := typ.(*dst.Ellipsis); ok {
		// Ellipsis params are represented as a slice (when not comparable)
		return Slice(ellipsis.Elt)
	}

	return typ
}

func (c *Converter) resultMgrStruct(typeName, prefix string) *dst.GenDecl {
	structName := fmt.Sprintf("%s_%s", prefix, resultMgrSuffix)

	return Struct(structName).Fields(
		Field(Idf("%s_%s", prefix, paramsIdent)).
			Names(Id(c.export(paramsIdent))).Obj,
		Field(Slice(Star(Id(c.export(fmt.Sprintf(
			"%s_%s", prefix, resultsIdent)))))).
			Names(Id(c.export(resultsIdent))).Obj,
		Field(Id("uint32")).Names(Id(c.export(indexIdent))).Obj,
		Field(Id("bool")).Names(Id(c.export(anyTimesIdent))).Obj,
	).Decs(genDeclDec(
		"// %s manages multiple results and the state of the %s type",
		structName,
		typeName)).Obj
}

func (c *Converter) fnRecorderStruct(typeName string, prefix string) *dst.GenDecl {
	mName := c.mockName(typeName)
	structName := fmt.Sprintf("%s_%s", prefix, fnRecorderSuffix)
	return Struct(structName).Fields(
		Field(Idf("%s_%s", prefix, paramsIdent)).
			Names(Id(c.export(paramsIdent))).Obj,
		Field(Idf("%s_%s", prefix, paramsKeyIdent)).
			Names(Id(c.export(paramsKeyIdent))).Obj,
		Field(Star(Idf("%s_%s", prefix, resultMgrSuffix))).
			Names(Id(c.export(resultsIdent))).Obj,
		Field(Star(Id(mName))).
			Names(Id(c.export(mockIdent))).Obj,
	).Decs(genDeclDec("// %s routes recorded function calls to the %s mock",
		structName, mName)).Obj
}

func (c *Converter) mockFunc(typePrefix, fieldSuffix string, fn Func) []dst.Stmt {
	stateSelector := Sel(Id(mockReceiverIdent)).Dot(Id(c.export(mockIdent))).Obj

	stmts := []dst.Stmt{
		Assign(Id(paramsIdent)).
			Tok(token.DEFINE).
			Rhs(Comp(Idf("%s_%s", typePrefix, paramsIdent)).
				Elts(c.passthroughElements(fn.Params, paramsIdent)...).Obj).Obj,
		Assign(Id(paramsKeyIdent)).
			Tok(token.DEFINE).
			Rhs(Comp(Idf("%s_%s", typePrefix, paramsKeyIdent)).
				Elts(c.passthroughElements(fn.Params, paramsKeyIdent)...).Obj).Obj,
		Assign(Id(resultsIdent), Id(okIdent)).
			Tok(token.DEFINE).
			Rhs(Index(Sel(Sel(Id(mockReceiverIdent)).Dot(Id(c.export(mockIdent))).Obj).
				Dot(Idf("%s%s", resultsByParamsIdent, fieldSuffix)).Obj).
				Sub(Id(paramsKeyIdent)).Obj).Obj,
	}

	stmts = append(stmts,
		If(Un(token.NOT, Id(okIdent))).Body(
			If(Bin(Sel(Sel(dst.Clone(stateSelector).(dst.Expr)).
				Dot(Id(c.export(configIdent))).Obj).
				Dot(Id(expectationIdent)).Obj).
				Op(token.EQL).
				Y(IdPath(strictIdent, moqPkg)).Obj).
				Body(
					Expr(Call(Sel(Sel(Sel(dst.Clone(stateSelector).(dst.Expr)).
						Dot(Id(c.export(sceneIdent))).Obj).
						Dot(Id(moqTType)).Obj).
						Dot(Id(fatalfFnName)).Obj).
						Args(LitString("Unexpected call with parameters %#v"),
							Id(paramsIdent)).Obj)).Obj,
			Return(),
		).Obj)

	stmts = append(stmts, Assign(Id(iIdent)).
		Tok(token.DEFINE).
		Rhs(Bin(Call(Id(intType)).
			Args(Call(IdPath("AddUint32", syncAtomicPkg)).Args(Un(
				token.AND,
				Sel(Id(resultsIdent)).Dot(Id(c.export(indexIdent))).Obj),
				LitInt(1)).Obj).Obj).
			Op(token.SUB).
			Y(LitInt(1)).Obj).
		Decs(AssignDecs(dst.EmptyLine).Obj).Obj)
	stmts = append(stmts,
		If(Bin(Id(iIdent)).Op(token.GEQ).Y(Call(Id(lenFnName)).
			Args(Sel(Id(resultsIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).Obj).Obj).
			Body(
				If(Un(token.NOT, Sel(Id(resultsIdent)).
					Dot(Id(c.export(anyTimesIdent))).Obj)).
					Body(
						If(Bin(Sel(Sel(dst.Clone(stateSelector).(dst.Expr)).
							Dot(Id(c.export(configIdent))).Obj).
							Dot(Id(expectationIdent)).Obj).
							Op(token.EQL).
							Y(IdPath(strictIdent, moqPkg)).Obj).
							Body(Expr(Call(Sel(Sel(Sel(dst.Clone(stateSelector).(dst.Expr)).
								Dot(Id(c.export(sceneIdent))).Obj).
								Dot(Id(moqTType)).Obj).
								Dot(Id(fatalfFnName)).Obj).
								Args(
									LitString("Too many calls to mock with parameters %#v"),
									Id(paramsIdent),
								).Obj)).Obj,
						Return(),
					).Obj,
				Assign(Id(iIdent)).
					Tok(token.ASSIGN).
					Rhs(Bin(Call(Id(lenFnName)).
						Args(Sel(Id(resultsIdent)).
							Dot(Id(c.export(resultsIdent))).Obj).Obj).
						Op(token.SUB).
						Y(LitInt(1)).Obj).Obj,
			).Obj)

	if fn.Results != nil {
		stmts = append(stmts, Assign(Id(resultIdent)).
			Tok(token.DEFINE).
			Rhs(Index(Sel(Id(resultsIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).Sub(Id(iIdent)).Obj).Obj)
		stmts = append(stmts, c.assignResult(fn.Results)...)
	}

	stmts = append(stmts, Return())

	return stmts
}

func (c *Converter) recorderFn(typeName string, fn Func) *dst.FuncDecl {
	mName := c.mockName(typeName)

	recvType := fmt.Sprintf("%s_%s", mName, recorderIdent)
	fnName := fn.Name
	fnRecName := fmt.Sprintf("%s_%s_%s", mName, fn.Name, fnRecorderSuffix)
	typePrefix := fmt.Sprintf("%s_%s", mName, fn.Name)
	var mockVal dst.Expr = Sel(Id(mockReceiverIdent)).
		Dot(Id(c.export(mockIdent))).Obj
	if fn.Name == "" {
		recvType = mName
		fnName = c.export(onCallFnName)
		fnRecName = fmt.Sprintf("%s_%s", mName, fnRecorderSuffix)
		typePrefix = mName
		mockVal = Id(mockReceiverIdent)
	}

	return Fn(fnName).
		Recv(Field(Star(Id(recvType))).Names(Id(mockReceiverIdent)).Obj).
		ParamFieldList(cloneAndNameUnnamed(paramPrefix, fn.Params)).
		Results(Field(Star(Id(fnRecName))).Obj).
		Body(c.recorderFnInterfaceBody(fnRecName, typePrefix, mockVal, fn)...).
		Decs(stdFuncDec()).Obj
}

func (c *Converter) recorderFnInterfaceBody(
	fnRecName, typePrefix string, mockValue dst.Expr, fn Func,
) []dst.Stmt {
	return []dst.Stmt{Return(Un(
		token.AND,
		Comp(Id(fnRecName)).
			Elts(
				Key(Id(c.export(paramsIdent))).
					Value(Comp(Idf("%s_%s", typePrefix, paramsIdent)).
						Elts(c.passthroughElements(fn.Params, paramsIdent)...).Obj,
					).Decs(kvExprDec(dst.None)).Obj,
				Key(Id(c.export(paramsKeyIdent))).
					Value(Comp(Idf("%s_%s", typePrefix, paramsKeyIdent)).
						Elts(c.passthroughElements(fn.Params, paramsKeyIdent)...).Obj,
					).Decs(kvExprDec(dst.None)).Obj,
				Key(Id(c.export(mockIdent))).
					Value(mockValue).Decs(kvExprDec(dst.None)).Obj,
			).Decs(litDec()).Obj,
	))}
}

func (c *Converter) recorderReturnFn(typeName string, fn Func) *dst.FuncDecl {
	mName := c.mockName(typeName)

	fnRecName := fmt.Sprintf("%s_%s_%s", mName, fn.Name, fnRecorderSuffix)
	results := fmt.Sprintf("%s_%s_%s", mName, fn.Name, resultsIdent)
	resultMgr := fmt.Sprintf("%s_%s_%s", mName, fn.Name, resultMgrSuffix)
	resultsByParams := fmt.Sprintf("%s_%s", resultsByParamsIdent, fn.Name)
	if fn.Name == "" {
		fnRecName = fmt.Sprintf("%s_%s", mName, fnRecorderSuffix)
		results = fmt.Sprintf("%s_%s", mName, resultsIdent)
		resultMgr = fmt.Sprintf("%s_%s", mName, resultMgrSuffix)
		resultsByParams = resultsByParamsIdent
		// TODO ???
		//mockSel := Sel(Sel(NewIdent(recorderReceiverIdent)).
		//	Dot(NewIdent(c.export(mockIdent))).Obj).Obj
	}

	mockSel := Sel(Sel(Id(recorderReceiverIdent)).
		Dot(Id(c.export(mockIdent))).Obj).Obj

	return Fn(c.export(returnFnName)).
		Recv(Field(Star(Id(fnRecName))).Names(Id(recorderReceiverIdent)).Obj).
		ParamFieldList(cloneAndNameUnnamed(resultPrefix, fn.Results)).
		Results(Field(Star(Id(fnRecName))).Obj).
		Body(
			If(Bin(Sel(Id(recorderReceiverIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).
				Op(token.EQL).
				Y(Id(nilIdent)).Obj).
				Body(
					IfInit(Assign(Id("_"), Id(okIdent)).
						Tok(token.DEFINE).
						Rhs(Index(cloneSelect(mockSel, c.export(resultsByParams))).
							Sub(Sel(Id(recorderReceiverIdent)).
								Dot(Id(c.export(paramsKeyIdent))).Obj).Obj).Obj).
						Cond(Id(okIdent)).
						Body(
							Expr(Call(Sel(Sel(cloneSelect(mockSel, c.export(sceneIdent))).
								Dot(Id(moqTType)).Obj).
								Dot(Id(fatalfFnName)).Obj).
								Args(LitString(
									"Expectations already recorded for mock with parameters %#v"),
									Sel(Id(recorderReceiverIdent)).
										Dot(Id(c.export(paramsIdent))).Obj,
								).Obj),
							Return(Id(nilIdent)),
						).Decs(IfDecs(dst.EmptyLine).Obj).Obj,
					Assign(Sel(Id(recorderReceiverIdent)).
						Dot(Id(c.export(resultsIdent))).Obj).
						Tok(token.ASSIGN).
						Rhs(Un(
							token.AND,
							Comp(Id(c.export(resultMgr))).
								Elts(
									Key(Id(c.export(paramsIdent))).
										Value(Sel(Id(recorderReceiverIdent)).
											Dot(Id(c.export(paramsIdent))).Obj).
										Decs(kvExprDec(dst.NewLine)).Obj,
									Key(Id(c.export(resultsIdent))).Value(
										Comp(Slice(Star(Id(c.export(results))))).Obj).
										Decs(kvExprDec(dst.None)).Obj,
									Key(Id(c.export(indexIdent))).Value(
										LitInt(0)).Decs(kvExprDec(dst.None)).Obj,
									Key(Id(c.export(anyTimesIdent))).Value(
										Id("false")).Decs(kvExprDec(dst.None)).Obj,
								).Obj,
						)).Obj,
					Assign(Index(Sel(Sel(Id(recorderReceiverIdent)).
						Dot(Id(c.export(mockIdent))).Obj).
						Dot(Id(c.export(resultsByParams))).Obj).
						Sub(Sel(Id(recorderReceiverIdent)).
							Dot(Id(c.export(paramsKeyIdent))).Obj).Obj).
						Tok(token.ASSIGN).
						Rhs(Sel(Id(recorderReceiverIdent)).
							Dot(Id(c.export(resultsIdent))).Obj).Obj,
				).Obj,
			Assign(
				Sel(Sel(Id(recorderReceiverIdent)).
					Dot(Id(c.export(resultsIdent))).Obj).
					Dot(Id(c.export(resultsIdent))).Obj,
			).
				Tok(token.ASSIGN).
				Rhs(Call(Id("append")).Args(Sel(Sel(Id(recorderReceiverIdent)).
					Dot(Id(c.export(resultsIdent))).Obj).
					Dot(Id(c.export(resultsIdent))).Obj,
					Un(token.AND, Comp(Id(c.export(results))).
						Elts(c.passthroughElements(fn.Results, resultsIdent)...).Obj)).Obj).Obj,
			Return(Id(recorderReceiverIdent)),
		).
		Decs(stdFuncDec()).Obj
}

func (c *Converter) recorderTimesFn(typeName string, fn Func) *dst.FuncDecl {
	mName := c.mockName(typeName)

	fnRecName := fmt.Sprintf("%s_%s_%s", mName, fn.Name, fnRecorderSuffix)
	if fn.Name == "" {
		fnRecName = fmt.Sprintf("%s_%s", mName, fnRecorderSuffix)
	}

	return Fn(c.export(timesFnName)).
		Recv(Field(Star(Id(fnRecName))).Names(Id(recorderReceiverIdent)).Obj).
		Params(Field(Id(intType)).Names(Id(countIdent)).Obj).
		Results(Field(Star(Id(fnRecName))).Obj).
		Body(
			If(Bin(Sel(Id(recorderReceiverIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).
				Op(token.EQL).
				Y(Id(nilIdent)).Obj).
				Body(
					Expr(Call(Sel(Sel(Sel(Sel(Id(recorderReceiverIdent)).
						Dot(Id(c.export(mockIdent))).Obj).
						Dot(Id(c.export(sceneIdent))).Obj).
						Dot(Id(moqTType)).Obj).Dot(Id(fatalfFnName)).Obj).
						Args(LitString(
							"Return must be called before calling Times")).Obj),
					Return(Id(nilIdent)),
				).Obj,
			Assign(Id(lastIdent)).
				Tok(token.DEFINE).
				Rhs(Index(Sel(Sel(Id(recorderReceiverIdent)).
					Dot(Id(c.export(resultsIdent))).Obj).
					Dot(Id(c.export(resultsIdent))).Obj).
					Sub(Bin(Call(Id(lenFnName)).
						Args(Sel(Sel(Id(recorderReceiverIdent)).
							Dot(Id(c.export(resultsIdent))).Obj).
							Dot(Id(c.export(resultsIdent))).Obj).Obj).
						Op(token.SUB).
						Y(LitInt(1)).Obj).Obj).Obj,
			For(Assign(Id("n")).Tok(token.DEFINE).Rhs(LitInt(0)).Obj).
				Cond(Bin(Id("n")).Op(token.LSS).
					Y(Bin(Id(countIdent)).Op(token.SUB).Y(LitInt(1)).Obj).Obj).
				Post(IncStmt(Id("n"))).
				Body(Assign(Sel(Sel(Id(recorderReceiverIdent)).
					Dot(Id(c.export(resultsIdent))).Obj).
					Dot(Id(c.export(resultsIdent))).Obj).
					Tok(token.ASSIGN).
					Rhs(Call(Id("append")).Args(Sel(Sel(Id(recorderReceiverIdent)).
						Dot(Id(c.export(resultsIdent))).Obj).
						Dot(Id(c.export(resultsIdent))).Obj,
						Id(lastIdent)).Obj).Obj).Obj,
			Return(Id(recorderReceiverIdent)),
		).
		Decs(stdFuncDec()).Obj
}

func (c *Converter) recorderAnyTimesFn(typeName string, fn Func) *dst.FuncDecl {
	mName := c.mockName(typeName)

	fnRecName := fmt.Sprintf("%s_%s_%s", mName, fn.Name, fnRecorderSuffix)
	if fn.Name == "" {
		fnRecName = fmt.Sprintf("%s_%s", mName, fnRecorderSuffix)
	}

	return Fn(c.export(anyTimesFnName)).
		Recv(Field(Star(Id(fnRecName))).Names(Id(recorderReceiverIdent)).Obj).
		Body(
			If(Bin(Sel(Id(recorderReceiverIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).
				Op(token.EQL).
				Y(Id(nilIdent)).Obj).
				Body(
					Expr(Call(Sel(Sel(Sel(Sel(Id(recorderReceiverIdent)).
						Dot(Id(c.export(mockIdent))).Obj).
						Dot(Id(c.export(sceneIdent))).Obj).
						Dot(Id(moqTType)).Obj).
						Dot(Id(fatalfFnName)).Obj).
						Args(LitString(
							"Return must be called before calling AnyTimes")).Obj),
					Return(),
				).Obj,
			Assign(Sel(Sel(Id(recorderReceiverIdent)).
				Dot(Id(c.export(resultsIdent))).Obj).
				Dot(Id(c.export(anyTimesIdent))).Obj).
				Tok(token.ASSIGN).
				Rhs(Id("true")).Obj,
		).
		Decs(stdFuncDec()).Obj
}

func (c *Converter) passthroughElements(fl *dst.FieldList, label string) []dst.Expr {
	unnamedPrefix, comparable := labelDirection(label)
	var elts []dst.Expr
	if fl != nil {
		beforeDec := dst.NewLine
		fields := fl.List
		for n, field := range fields {
			if len(field.Names) == 0 {
				pName := fmt.Sprintf("%s%d", unnamedPrefix, n+1)
				elts = append(elts, Key(Id(c.export(pName))).Value(
					passthroughValue(pName, field.Type, comparable)).
					Decs(kvExprDec(beforeDec)).Obj)
				beforeDec = dst.None
			}

			for _, name := range field.Names {
				elts = append(elts, Key(Id(c.export(name.Name))).Value(
					passthroughValue(name.Name, field.Type, comparable)).
					Decs(kvExprDec(beforeDec)).Obj)
				beforeDec = dst.None
			}
		}
	}

	return elts
}

func passthroughValue(name string, typ dst.Expr, comparable bool) dst.Expr {
	var val dst.Expr
	val = Id(name)
	if comparable && !isComparable(typ) {
		val = Call(IdPath("DeepHash", hashPkg)).Args(val).Obj
	}
	return val
}

func passthroughFields(fields *dst.FieldList) []dst.Expr {
	var exprs []dst.Expr
	for _, f := range fields.List {
		for _, n := range f.Names {
			exprs = append(exprs, Id(n.Name))
		}
	}
	return exprs
}

func (c *Converter) assignResult(resFL *dst.FieldList) []dst.Stmt {
	var assigns []dst.Stmt
	if resFL != nil {
		results := resFL.List
		for n, result := range results {
			if len(result.Names) == 0 {
				rName := fmt.Sprintf("%s%d", resultPrefix, n+1)
				assigns = append(assigns, Assign(Id(rName)).
					Tok(token.ASSIGN).
					Rhs(Sel(Id(resultIdent)).
						Dot(Id(c.export(rName))).Obj).Obj)
			}

			for _, name := range result.Names {
				assigns = append(assigns, Assign(Id(name.Name)).
					Tok(token.ASSIGN).
					Rhs(Sel(Id(resultIdent)).
						Dot(Id(c.export(name.Name))).Obj).Obj)
			}
		}
	}
	return assigns
}

func cloneAndNameUnnamed(prefix string, fieldList *dst.FieldList) *dst.FieldList {
	fieldList = cloneNilableFieldList(fieldList)
	if fieldList != nil {
		for n, f := range fieldList.List {
			if len(f.Names) == 0 {
				f.Names = []*dst.Ident{Idf("%s%d", prefix, n+1)}
			}
		}
	}
	return fieldList
}

func (c *Converter) mockName(typeName string) string {
	return c.export(mockIdent + strings.Title(typeName))
}

func (c *Converter) export(name string) string {
	if c.isExported {
		name = strings.Title(name)
	}
	return name
}

func stdFuncDec() dst.FuncDeclDecorations {
	return dst.FuncDeclDecorations{
		NodeDecs: dst.NodeDecs{Before: dst.EmptyLine, After: dst.EmptyLine},
	}
}

func labelDirection(label string) (unnamedPrefix string, comparable bool) {
	switch label {
	case paramsIdent:
		unnamedPrefix = paramPrefix
		comparable = false
	case paramsKeyIdent:
		unnamedPrefix = paramPrefix
		comparable = true
	case resultsIdent:
		unnamedPrefix = resultPrefix
		comparable = false
	default:
		logs.Panicf("Unknown label: %s", label)
	}

	return unnamedPrefix, comparable
}

func cloneNilableFieldList(fl *dst.FieldList) *dst.FieldList {
	if fl != nil {
		fl = dst.Clone(fl).(*dst.FieldList)
	}
	return fl
}

func cloneSelect(x *dst.SelectorExpr, sel string) *dst.SelectorExpr {
	x = dst.Clone(x).(*dst.SelectorExpr)
	x.Sel = Id(sel)
	return x
}

func genDeclDec(format string, a ...interface{}) dst.GenDeclDecorations {
	return dst.GenDeclDecorations{
		NodeDecs: nodeDec(format, a...),
	}
}

func fnDeclDec(format string, a ...interface{}) dst.FuncDeclDecorations {
	return dst.FuncDeclDecorations{
		NodeDecs: nodeDec(format, a...),
	}
}

func nodeDec(format string, a ...interface{}) dst.NodeDecs {
	return dst.NodeDecs{
		Before: dst.NewLine,
		Start:  dst.Decorations{fmt.Sprintf(format, a...)},
	}
}

func litDec() dst.CompositeLitDecorations {
	return dst.CompositeLitDecorations{Lbrace: []string{"\n"}}
}

func kvExprDec(before dst.SpaceType) dst.KeyValueExprDecorations {
	return KeyValueDecs(before).After(dst.NewLine).Obj
}
