package internal

import (
	"fmt"
	utils2 "github.com/aman-bansal/go_perf_tracker/perf_tracker/internal/utils"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
)

//need to add code
//defer go_perf_tracker.CalculateFunctionPerf(time.Now(), "")
func FormatCode(fileset *token.FileSet, file *ast.File, funcs []string) {
	ast.Inspect(file, func(node ast.Node) bool {
		if fn, ok := node.(*ast.FuncDecl); ok {
			if utils2.Contains(funcs, fn.Name.Name) {
				//it contains add the perf code
				astutil.AddImport(fileset, file, "github.com/aman-bansal/go_perf_tracker/perf_tracker")
				astutil.AddImport(fileset, file, "time")

				fn.Body.List = append([]ast.Stmt{
					&ast.DeferStmt{
						Call: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "perf_tracker"},
								Sel: &ast.Ident{Name: "CalculateFunctionPerf"},
							},
							Args: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "time"},
										Sel: &ast.Ident{Name: "Now"},
									},
								},
								&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("\"%s\"", fn.Name.Name)},
							},
						},
					},
				}, fn.Body.List...)
			}
		}
		return true
	})
}
