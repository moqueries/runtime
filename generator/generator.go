// Package generator generates Moqueries mocks
package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/gopackages"
	"golang.org/x/tools/go/packages"

	"github.com/myshkin5/moqueries/ast"
	"github.com/myshkin5/moqueries/logs"
	"github.com/myshkin5/moqueries/metrics"
)

// GenerateRequest contains all the parameters needed to call Generate
type GenerateRequest struct {
	Types          []string `json:"types"`
	Export         bool     `json:"export"`
	Destination    string   `json:"destination"`
	DestinationDir string   `json:"destination-dir"`
	Package        string   `json:"package"`
	Import         string   `json:"import"`
	TestImport     bool     `json:"test-import"`
	// WorkingDir is the current working directory. Optional, in which case
	// os.Getwd is used. Useful in cases where a request is serialized then
	// rerun in bulk processing from a different working directory. WorkingDir
	// is used for relative-path imports and relative path destination
	// files/directories.
	WorkingDir string `json:"working-dir"`
}

// Generate generates a moq
func Generate(reqs ...GenerateRequest) error {
	m := metrics.NewMetrics(logs.IsDebug, logs.Debugf)
	cache := ast.NewCache(packages.Load, m)
	start := time.Now()
	for _, req := range reqs {
		err := GenerateWithTypeCache(cache, req)
		if err != nil {
			return err
		}
	}
	m.TotalProcessingTimeInc(time.Since(start))
	m.Finalize()
	return nil
}

//go:generate moqueries TypeCache

// TypeCache defines the interface to the Cache type
type TypeCache interface {
	Type(id dst.Ident, testImport bool) (*dst.TypeSpec, string, error)
	IsComparable(expr dst.Expr) (bool, error)
	IsDefaultComparable(expr dst.Expr) (bool, error)
	FindPackage(dir string) (string, error)
}

// GenerateWithTypeCache generates a single moq using the provided type cache.
// This function is exposed for use in bulk operations that have already loaded
// a type.
func GenerateWithTypeCache(cache TypeCache, req GenerateRequest) error {
	newConverterFn := func(typ Type, export bool) Converterer {
		return NewConverter(typ, export, cache)
	}
	gen := New(cache, os.Getwd, newConverterFn)

	file, destPath, err := gen.Generate(req)
	if err != nil {
		return fmt.Errorf("error generating moqs: %w", err)
	}

	tempFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}

	defer func() {
		err = tempFile.Close()
		if err != nil {
			logs.Error("Error closing temp file", err)
		}
	}()

	destDir := filepath.Dir(destPath)
	if _, err = os.Stat(destDir); os.IsNotExist(err) {
		err = os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			logs.Errorf(
				"Error creating destination directory %s from working director %s: %v",
				destDir, req.WorkingDir, err)
		}
	}

	restorer := decorator.NewRestorerWithImports(destDir, gopackages.New(destDir))
	err = restorer.Fprint(tempFile, file)
	if err != nil {
		return fmt.Errorf("invalid moq: %w", err)
	}

	err = os.Rename(tempFile.Name(), destPath)
	if err != nil {
		logs.Debugf("Error removing destination file: %v", err)
	}
	logs.Debugf("Wrote file: %s", destPath)

	return nil
}
