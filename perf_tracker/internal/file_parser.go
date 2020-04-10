package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(sourceFile string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, sourceFile, nil, 0)
	if err != nil {
		return nil, nil, err
	}
	return fset, file, nil
}
