package generator_test

import (
	"errors"

	"github.com/dave/dst"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/myshkin5/moqueries/pkg/generator"
	"github.com/myshkin5/moqueries/pkg/hash"
)

var _ = Describe("MockGen", func() {
	var (
		loadTypesFnMock *mockLoadTypesFn
		converterMock   *mockConverterer

		ifaceSpec1    *dst.TypeSpec
		ifaceSpec2    *dst.TypeSpec
		ifaceMethods1 *dst.FieldList
		ifaceMethods2 *dst.FieldList
		func1         *dst.Field
		func1Params   *dst.FieldList

		fnSpec *dst.TypeSpec
		fnType *dst.FuncType

		readerSpec *dst.TypeSpec
	)

	BeforeEach(func() {
		loadTypesFnMock = newMockLoadTypesFn()
		converterMock = newMockConverterer()

		func1Params = &dst.FieldList{
			List: []*dst.Field{
				{
					Names: []*dst.Ident{dst.NewIdent("firstParm")},
					Type: &dst.StarExpr{
						X: &dst.SelectorExpr{
							X:   dst.NewIdent("cobra"),
							Sel: dst.NewIdent("Command"),
						},
					},
				},
				{
					Type: dst.NewIdent("string"),
				},
				{
					Type: &dst.StarExpr{
						X: &dst.SelectorExpr{
							X:   dst.NewIdent("dst"),
							Sel: dst.NewIdent("InterfaceType"),
						},
					},
				},
			},
		}
		func1 = &dst.Field{
			Names: []*dst.Ident{dst.NewIdent("Func1")},
			Type: &dst.FuncType{
				Params:  func1Params,
				Results: nil,
			},
		}
		ifaceMethods1 = &dst.FieldList{
			List: []*dst.Field{func1},
		}
		ifaceSpec1 = &dst.TypeSpec{
			Name: dst.NewIdent("PublicInterface"),
			Type: &dst.InterfaceType{Methods: ifaceMethods1},
		}
		ifaceMethods2 = &dst.FieldList{}
		ifaceSpec2 = &dst.TypeSpec{
			Name: dst.NewIdent("privateInterface"),
			Type: &dst.InterfaceType{Methods: ifaceMethods2},
		}

		fnType = &dst.FuncType{
			Params:  func1Params,
			Results: nil,
		}
		fnSpec = &dst.TypeSpec{
			Name: dst.NewIdent("PublicFn"),
			Type: fnType,
		}

		readerSpec = &dst.TypeSpec{
			Name: &dst.Ident{
				Name: "Reader",
				Path: "io",
			},
			Type: &dst.InterfaceType{
				Methods: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("Read")},
							Type: &dst.FuncType{
								Params: &dst.FieldList{
									List: []*dst.Field{
										{
											Names: []*dst.Ident{dst.NewIdent("p")},
											Type:  &dst.ArrayType{Elt: dst.NewIdent("byte")},
										},
									},
								},
								Results: &dst.FieldList{
									List: []*dst.Field{
										{
											Names: []*dst.Ident{dst.NewIdent("n")},
											Type:  dst.NewIdent("int"),
										},
										{
											Names: []*dst.Ident{dst.NewIdent("err")},
											Type:  dst.NewIdent("error"),
										},
									},
								},
							},
						},
					},
				},
			},
		}
	})

	It("always returns a header comment", func() {
		// ASSEMBLE
		gen := generator.New(false, "", "dir/file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate(nil, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file).NotTo(BeNil())
		Expect(len(file.Decs.Start)).To(BeNumerically(">", 0))
		Expect(file.Decs.Start[0]).To(Equal(
			"// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!"))
	})

	It("defaults the package when it isn't specified", func() {
		// ASSEMBLE
		gen := generator.New(false, "", "dir/file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate(nil, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Name.Name).To(Equal("dir_test"))
	})

	It("defaults the package to a name based on the current directory when it isn't specified", func() {
		// ASSEMBLE
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate(nil, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Name.Name).To(Equal("generator_test"))
	})

	It("creates structs for each mock", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{
				ifaceSpec1,
				ifaceSpec2,
				fnSpec,
			},
			pkgPath: "github.com/myshkin5/moqueries/pkg/generator",
			err:     nil,
		}
		pubDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-decl")}},
		}
		ifaceFuncs := []generator.Func{
			{
				Name:   "Func1",
				Params: func1Params,
			},
		}
		converterMock.resultsByParams_BaseStruct[mockConverterer_BaseStruct_params{
			typeSpec: ifaceSpec1,
			funcs:    hash.DeepHash(ifaceFuncs),
		}] = mockConverterer_BaseStruct_results{structDecl: pubDecl}
		pubMockDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-mock-decl")}},
		}
		converterMock.resultsByParams_IsolationStruct[mockConverterer_IsolationStruct_params{
			typeName: "PublicInterface",
			suffix:   "mock",
		}] = mockConverterer_IsolationStruct_results{structDecl: pubMockDecl}
		pubRecorderDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-rec-decl")}},
		}
		converterMock.resultsByParams_IsolationStruct[mockConverterer_IsolationStruct_params{
			typeName: "PublicInterface",
			suffix:   "recorder",
		}] = mockConverterer_IsolationStruct_results{structDecl: pubRecorderDecl}
		pubStructDecls := []dst.Decl{
			&dst.GenDecl{
				Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-struct-decl")}},
			},
		}
		converterMock.resultsByParams_MethodStructs[mockConverterer_MethodStructs_params{
			typeName: "PublicInterface",
			prefix:   "mockPublicInterface_Func1",
			fn:       ifaceFuncs[0],
		}] = mockConverterer_MethodStructs_results{structDecls: pubStructDecls}

		privateDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("private-decl")}},
		}
		converterMock.resultsByParams_BaseStruct[mockConverterer_BaseStruct_params{
			typeSpec: ifaceSpec2,
			funcs:    hash.DeepHash([]generator.Func{}),
		}] = mockConverterer_BaseStruct_results{structDecl: privateDecl}

		fnDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-decl")}},
		}
		fnFuncs := []generator.Func{
			{
				Params: func1Params,
			},
		}
		converterMock.resultsByParams_BaseStruct[mockConverterer_BaseStruct_params{
			typeSpec: fnSpec,
			funcs:    hash.DeepHash(fnFuncs),
		}] = mockConverterer_BaseStruct_results{structDecl: fnDecl}
		fnMockDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-mock-decl")}},
		}
		converterMock.resultsByParams_IsolationStruct[mockConverterer_IsolationStruct_params{
			typeName: "PublicFn",
			suffix:   "mock",
		}] = mockConverterer_IsolationStruct_results{structDecl: fnMockDecl}
		fnRecorderDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-rec-decl")}},
		}
		converterMock.resultsByParams_IsolationStruct[mockConverterer_IsolationStruct_params{
			typeName: "PublicFn",
			suffix:   "recorder",
		}] = mockConverterer_IsolationStruct_results{structDecl: fnRecorderDecl}
		fnStructDecls := []dst.Decl{
			&dst.GenDecl{
				Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-struct-decl")}},
			},
		}
		converterMock.resultsByParams_MethodStructs[mockConverterer_MethodStructs_params{
			typeName: "PublicFn",
			prefix:   "mockPublicFn",
			fn:       fnFuncs[0],
		}] = mockConverterer_MethodStructs_results{structDecls: fnStructDecls}

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface", "privateInterface", "PublicFn"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(len(file.Decls)).To(BeNumerically(">", 8))
		// Only looking at the first three structs for PublicInterface
		Expect(file.Decls[:4]).To(Equal([]dst.Decl{
			pubDecl,
			pubMockDecl,
			pubRecorderDecl,
			pubStructDecls[0],
		}))

		fnStart := -1
		for n, decl := range file.Decls {
			if decl == fnDecl {
				fnStart = n
			}
		}
		Expect(fnStart).NotTo(Equal(-1))
		// Only looking at the structs for PublicFn
		Expect(file.Decls[fnStart : fnStart+4]).To(Equal([]dst.Decl{
			fnDecl,
			fnMockDecl,
			fnRecorderDecl,
			fnStructDecls[0],
		}))
		Expect(file.Decls).To(ContainElement(privateDecl))
	})

	It("recursively looks up nested interfaces", func() {
		// ASSEMBLE
		// PublicInterface embeds privateInterface which embeds io.Reader
		ifaceMethods1.List = append(ifaceMethods1.List, &dst.Field{
			Type: &dst.Ident{
				Name: "privateInterface",
				Path: "github.com/myshkin5/moqueries/pkg/generator",
			},
		})
		ifaceMethods2.List = append(ifaceMethods2.List, &dst.Field{
			Type: &dst.Ident{
				Name: "Reader",
				Path: "io",
			},
		})
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{
				ifaceSpec1,
				ifaceSpec2,
			},
			pkgPath: "github.com/myshkin5/moqueries/pkg/generator",
			err:     nil,
		}
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "io"}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{readerSpec},
			pkgPath:   "io",
			err:       nil,
		}
		pubDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-decl")}},
		}
		converterMock.resultsByParams_BaseStruct[mockConverterer_BaseStruct_params{
			typeSpec: ifaceSpec1,
			funcs:    hash.DeepHash([]generator.Func{}),
		}] = mockConverterer_BaseStruct_results{structDecl: pubDecl}

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, _, err := gen.Generate([]string{"PublicInterface", "privateInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
	})

	It("creates a new mock function", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{ifaceSpec1},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		newFunc := &dst.FuncDecl{Name: dst.NewIdent("newMockFn")}
		funcs := []generator.Func{
			{
				Name:   "Func1",
				Params: func1Params,
			},
		}
		converterMock.resultsByParams_NewFunc[mockConverterer_NewFunc_params{
			typeSpec: ifaceSpec1,
			funcs:    hash.DeepHash(funcs),
		}] = mockConverterer_NewFunc_results{funcDecl: newFunc}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		fnStart := -1
		for n, decl := range file.Decls {
			if decl == newFunc {
				fnStart = n
			}
		}
		Expect(fnStart).NotTo(Equal(-1))
	})

	It("creates a mock accessor", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{ifaceSpec1},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		mockFn := &dst.FuncDecl{Name: dst.NewIdent("mock")}
		converterMock.resultsByParams_IsolationAccessor[mockConverterer_IsolationAccessor_params{
			typeName: "PublicInterface",
			suffix:   "mock",
			fnName:   "mock",
		}] = mockConverterer_IsolationAccessor_results{funcDecl: mockFn}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		fnStart := -1
		for n, decl := range file.Decls {
			if decl == mockFn {
				fnStart = n
			}
		}
		Expect(fnStart).NotTo(Equal(-1))
	})

	It("creates mock functions for each method in the interface", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{ifaceSpec1},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		methodFn := &dst.FuncDecl{Name: dst.NewIdent("Func1")}
		converterMock.resultsByParams_MockMethod[mockConverterer_MockMethod_params{
			typeName: "PublicInterface",
			fn: generator.Func{
				Name:   "Func1",
				Params: func1Params,
			},
		}] = mockConverterer_MockMethod_results{funcDecl: methodFn}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFn))
	})

	It("creates a closure for a mock func", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{fnSpec},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		methodFn := &dst.FuncDecl{Name: dst.NewIdent("Func1")}
		converterMock.resultsByParams_FuncClosure[mockConverterer_FuncClosure_params{
			typeName: "PublicFn",
			pkgPath:  "github.com/myshkin5/moqueries/pkg/generator",
			fn: generator.Func{
				Params: func1Params,
			},
		}] = mockConverterer_FuncClosure_results{funcDecl: methodFn}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicFn"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFn))
	})

	It("creates a recorder accessor", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{ifaceSpec1},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		recFn := &dst.FuncDecl{Name: dst.NewIdent("onCall")}
		converterMock.resultsByParams_IsolationAccessor[mockConverterer_IsolationAccessor_params{
			typeName: "PublicInterface",
			suffix:   "recorder",
			fnName:   "onCall",
		}] = mockConverterer_IsolationAccessor_results{funcDecl: recFn}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		fnStart := -1
		for n, decl := range file.Decls {
			if decl == recFn {
				fnStart = n
			}
		}
		Expect(fnStart).NotTo(Equal(-1))
	})

	It("creates recorder functions for each method in the interface", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{ifaceSpec1},
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       nil,
		}
		methodFns := []dst.Decl{&dst.FuncDecl{Name: dst.NewIdent("Func1")}}
		converterMock.resultsByParams_RecorderMethods[mockConverterer_RecorderMethods_params{
			typeName: "PublicInterface",
			fn: generator.Func{
				Name:   "Func1",
				Params: func1Params,
			},
		}] = mockConverterer_RecorderMethods_results{funcDecls: methodFns}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".")

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFns[0]))
	})

	It("returns an error when the interface can't be found", func() {
		// ASSEMBLE
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: []*dst.TypeSpec{
				{
					Name: dst.NewIdent("SomethingElseInterface"),
				},
			},
			pkgPath: "github.com/myshkin5/moqueries/pkg/generator",
			err:     nil,
		}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		fSet, file, err := gen.Generate([]string{"NotThereInterface"}, ".")

		// ASSERT
		Expect(err).To(Equal(errors.New("type was not found: NotThereInterface")))
		Expect(fSet).To(BeNil())
		Expect(file).To(BeNil())
	})

	It("returns an ast error when the interfaces can't be loaded", func() {
		// ASSEMBLE
		loadErr := errors.New("ast is not happy")
		loadTypesFnMock.resultsByParams[mockLoadTypesFn_params{pkg: "."}] = mockLoadTypesFn_results{
			typeSpecs: nil,
			pkgPath:   "github.com/myshkin5/moqueries/pkg/generator",
			err:       loadErr,
		}
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		fSet, file, err := gen.Generate([]string{"NotThereInterface"}, ".")

		// ASSERT
		Expect(err).To(Equal(loadErr))
		Expect(fSet).To(BeNil())
		Expect(file).To(BeNil())
	})
})
