package generator_test

import (
	"errors"

	"github.com/dave/dst"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/myshkin5/moqueries/pkg/generator"
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
		_, file, err := gen.Generate(nil, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file).NotTo(BeNil())
		Expect(len(file.Decs.Start)).To(BeNumerically(">", 0))
		Expect(file.Decs.Start[0]).To(Equal(
			"// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!",
		))
	})

	It("defaults the package when it isn't specified", func() {
		// ASSEMBLE
		gen := generator.New(false, "", "dir/file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate(nil, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Name.Name).To(Equal("dir_test"))
	})

	It("defaults the package to a name based on the current directory when it isn't specified", func() {
		// ASSEMBLE
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate(nil, ".", false)

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
		converterMock.onCall().BaseStruct(ifaceSpec1, ifaceFuncs).ret(pubDecl)
		pubMockDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-mock-decl")}},
		}
		converterMock.onCall().IsolationStruct("PublicInterface", "mock").ret(pubMockDecl)
		pubRecorderDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-rec-decl")}},
		}
		converterMock.onCall().IsolationStruct("PublicInterface", "recorder").ret(pubRecorderDecl)
		pubStructDecls := []dst.Decl{
			&dst.GenDecl{
				Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-struct-decl")}},
			},
		}
		converterMock.onCall().MethodStructs(
			"PublicInterface",
			"mockPublicInterface_Func1",
			ifaceFuncs[0],
		).ret(pubStructDecls)

		privateDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("private-decl")}},
		}
		converterMock.onCall().BaseStruct(ifaceSpec2, []generator.Func{}).ret(privateDecl)

		fnDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-decl")}},
		}
		fnFuncs := []generator.Func{
			{
				Params: func1Params,
			},
		}
		converterMock.onCall().BaseStruct(fnSpec, fnFuncs).ret(fnDecl)
		fnMockDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-mock-decl")}},
		}
		converterMock.onCall().IsolationStruct("PublicFn", "mock").ret(fnMockDecl)
		fnRecorderDecl := &dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-rec-decl")}},
		}
		converterMock.onCall().IsolationStruct("PublicFn", "recorder").ret(fnRecorderDecl)
		fnStructDecls := []dst.Decl{
			&dst.GenDecl{
				Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("fn-struct-decl")}},
			},
		}
		converterMock.onCall().MethodStructs("PublicFn", "mockPublicFn", fnFuncs[0]).ret(fnStructDecls)

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface", "privateInterface", "PublicFn"}, ".", false)

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
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{
				ifaceSpec1,
				ifaceSpec2,
			},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		loadTypesFnMock.onCall("io", false).ret([]*dst.TypeSpec{readerSpec}, "io", nil)
		converterMock.onCall().BaseStruct(ifaceSpec1, []generator.Func{}).ret(&dst.GenDecl{
			Specs: []dst.Spec{&dst.TypeSpec{Name: dst.NewIdent("pub-decl")}},
		})

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, _, err := gen.Generate([]string{"PublicInterface", "privateInterface"}, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
	})

	It("loads tests types when requested", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall(".", true).ret(
			[]*dst.TypeSpec{
				ifaceSpec1,
				ifaceSpec2,
			},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, _, err := gen.Generate([]string{"PublicInterface", "privateInterface"}, ".", true)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(loadTypesFnMock.params).To(Receive(&mockLoadTypesFn_params{
			pkg:           ".",
			loadTestTypes: true,
		}))
	})

	It("loads tests types when importing a test package", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall("github.com/myshkin5/moqueries/pkg/generator", true).ret(
			[]*dst.TypeSpec{
				ifaceSpec1,
				ifaceSpec2,
			},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)

		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, _, err := gen.Generate(
			[]string{"PublicInterface", "privateInterface"},
			"github.com/myshkin5/moqueries/pkg/generator_test",
			false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(loadTypesFnMock.params).To(Receive(&mockLoadTypesFn_params{
			pkg:           ".",
			loadTestTypes: true,
		}))
	})

	It("creates a new mock function", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{ifaceSpec1},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		newFunc := &dst.FuncDecl{Name: dst.NewIdent("newMockFn")}
		converterMock.onCall().NewFunc(
			ifaceSpec1,
			[]generator.Func{
				{
					Name:   "Func1",
					Params: func1Params,
				},
			},
		).ret(newFunc)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".", false)

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
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{ifaceSpec1},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		mockFn := &dst.FuncDecl{Name: dst.NewIdent("mock")}
		converterMock.onCall().IsolationAccessor("PublicInterface", "mock", "mock").ret(mockFn)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".", false)

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
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{ifaceSpec1},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		methodFn := &dst.FuncDecl{Name: dst.NewIdent("Func1")}
		converterMock.onCall().MockMethod(
			"PublicInterface",
			generator.Func{
				Name:   "Func1",
				Params: func1Params,
			},
		).ret(methodFn)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFn))
	})

	It("creates a closure for a mock func", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{fnSpec},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		methodFn := &dst.FuncDecl{Name: dst.NewIdent("Func1")}
		converterMock.onCall().FuncClosure(
			"PublicFn",
			"github.com/myshkin5/moqueries/pkg/generator",
			generator.Func{Params: func1Params},
		).ret(methodFn)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicFn"}, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFn))
	})

	It("creates a recorder accessor", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{ifaceSpec1},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		recFn := &dst.FuncDecl{Name: dst.NewIdent("onCall")}
		converterMock.onCall().IsolationAccessor("PublicInterface", "recorder", "onCall").ret(recFn)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".", false)

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
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{ifaceSpec1},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		methodFns := []dst.Decl{&dst.FuncDecl{Name: dst.NewIdent("Func1")}}
		converterMock.onCall().RecorderMethods(
			"PublicInterface",
			generator.Func{
				Name:   "Func1",
				Params: func1Params,
			},
		).ret(methodFns)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		_, file, err := gen.Generate([]string{"PublicInterface"}, ".", false)

		// ASSERT
		Expect(err).NotTo(HaveOccurred())
		Expect(file.Decls).To(ContainElement(methodFns[0]))
	})

	It("returns an error when the interface can't be found", func() {
		// ASSEMBLE
		loadTypesFnMock.onCall(".", false).ret(
			[]*dst.TypeSpec{{Name: dst.NewIdent("SomethingElseInterface")}},
			"github.com/myshkin5/moqueries/pkg/generator",
			nil,
		)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		fSet, file, err := gen.Generate([]string{"NotThereInterface"}, ".", false)

		// ASSERT
		Expect(err).To(Equal(errors.New("type was not found: NotThereInterface")))
		Expect(fSet).To(BeNil())
		Expect(file).To(BeNil())
	})

	It("returns an ast error when the interfaces can't be loaded", func() {
		// ASSEMBLE
		loadErr := errors.New("ast is not happy")
		loadTypesFnMock.onCall(".", false).ret(nil, "github.com/myshkin5/moqueries/pkg/generator", loadErr)
		gen := generator.New(false, "", "file_test.go", loadTypesFnMock.mock(), converterMock.mock())

		// ACT
		fSet, file, err := gen.Generate([]string{"NotThereInterface"}, ".", false)

		// ASSERT
		Expect(err).To(Equal(loadErr))
		Expect(fSet).To(BeNil())
		Expect(file).To(BeNil())
	})
})
