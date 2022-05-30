build 参数：
go build -debug-actiongraph=build.json -debug-trace=out.trace main.go 


enter: src/cmd/go/main.go

source code:

src/cmd/go/internal/run/run.go : run.go

modload.InitWorkfile()

p = load.GoFilesPackage(ctx, pkgOpts, files)// files=main.go
pkg.load(ctx, opts, "command-line-arguments", &stk, nil, bp, err)// 找到所有的依赖，构建package的信息
useBindir = true
用栈来保存import path的信息
	stk.Push(path)
	defer stk.Pop()



// Build list of imported packages and full dependency list.
	imports := make([]*Package, 0, len(p.Imports))
	for i, path := range importPaths { // 导入所有的import
		if path == "C" {
			continue
		}
        // 递归的遍历
		p1 := LoadImport(ctx, opts, path, p.Dir, p, stk, p.Internal.Build.ImportPos[path], ResolveImport)

		path = p1.ImportPath
		importPaths[i] = path
		if i < len(p.Imports) {
			p.Imports[i] = path
		}

		imports = append(imports, p1)
		if p1.Incomplete {
			p.Incomplete = true
		}
	}



=>
func loadImport(ctx context.Context, opts PackageOpts, pre *preload, path, srcDir string, parent *Package, stk *ImportStack, importPos []token.Position, mode int) *Package {
}

=> 	bp, loaded, err := loadPackageData(ctx, path, parentPath, srcDir, parentRoot, parentIsStd, mode)


++++> link 程序
	a1 := b.LinkAction(work.ModeBuild, work.ModeBuild, p)
	a := &work.Action{Mode: "go run", Func: buildRunProgram, Args: cmdArgs, Deps: []*work.Action{a1}}
	b.Do(ctx, a)



// A Package describes the Go package found in a directory.
type Package struct {
	Dir           string   // directory containing package sources
	Name          string   // package name
	ImportComment string   // path in import comment on package statement
	Doc           string   // documentation synopsis
	ImportPath    string   // import path of package ("" if unknown)
	Root          string   // root of Go tree where this package lives
	SrcRoot       string   // package source root directory ("" if unknown)
	PkgRoot       string   // package install root directory ("" if unknown)
	PkgTargetRoot string   // architecture dependent install root directory ("" if unknown)
	BinDir        string   // command install directory ("" if unknown)
	Goroot        bool     // package found in Go root
	PkgObj        string   // installed .a file
	AllTags       []string // tags that can influence file selection in this directory
	ConflictDir   string   // this directory shadows Dir in $GOPATH
	BinaryOnly    bool     // cannot be rebuilt from source (has //go:binary-only-package comment)

	// Source files
	GoFiles           []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles          []string // .go source files that import "C"
	IgnoredGoFiles    []string // .go source files ignored for this build (including ignored _test.go files)
	InvalidGoFiles    []string // .go source files with detected problems (parse error, wrong package name, and so on)
	IgnoredOtherFiles []string // non-.go source files ignored for this build
	CFiles            []string // .c source files
	CXXFiles          []string // .cc, .cpp and .cxx source files
	MFiles            []string // .m (Objective-C) source files
	HFiles            []string // .h, .hh, .hpp and .hxx source files
	FFiles            []string // .f, .F, .for and .f90 Fortran source files
	SFiles            []string // .s source files
	SwigFiles         []string // .swig files
	SwigCXXFiles      []string // .swigcxx files
	SysoFiles         []string // .syso system object files to add to archive

	// Cgo directives
	CgoCFLAGS    []string // Cgo CFLAGS directives
	CgoCPPFLAGS  []string // Cgo CPPFLAGS directives
	CgoCXXFLAGS  []string // Cgo CXXFLAGS directives
	CgoFFLAGS    []string // Cgo FFLAGS directives
	CgoLDFLAGS   []string // Cgo LDFLAGS directives
	CgoPkgConfig []string // Cgo pkg-config directives

	// Test information
	TestGoFiles  []string // _test.go files in package
	XTestGoFiles []string // _test.go files outside package

	// Dependency information
	Imports        []string                    // import paths from GoFiles, CgoFiles
	ImportPos      map[string][]token.Position // line information for Imports
	TestImports    []string                    // import paths from TestGoFiles
	TestImportPos  map[string][]token.Position // line information for TestImports
	XTestImports   []string                    // import paths from XTestGoFiles
	XTestImportPos map[string][]token.Position // line information for XTestImports

	// //go:embed patterns found in Go source files
	// For example, if a source file says
	//	//go:embed a* b.c
	// then the list will contain those two strings as separate entries.
	// (See package embed for more details about //go:embed.)
	EmbedPatterns        []string                    // patterns from GoFiles, CgoFiles
	EmbedPatternPos      map[string][]token.Position // line information for EmbedPatterns
	TestEmbedPatterns    []string                    // patterns from TestGoFiles
	TestEmbedPatternPos  map[string][]token.Position // line information for TestEmbedPatterns
	XTestEmbedPatterns   []string                    // patterns from XTestGoFiles
	XTestEmbedPatternPos map[string][]token.Position // line information for XTestEmbedPatternPos
}