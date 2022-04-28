package main

import (
	"fmt"
)

//go:generate go run golang.org/x/tools/cmd/goyacc@master -l -o parser.go grammar.y

func main() {
	yyErrorVerbose = true
	l := Lexer{input: "c OR b AND 12121 = 3"}
	l.readChar()

	yyParse(&l)

	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).left.(*LogicalExpr).left)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).left.(*LogicalExpr).op)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).left.(*LogicalExpr).right)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).op)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).right.(*CmpExpr).left)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).right.(*CmpExpr).op)
	fmt.Printf("%v\n", l.ast.(*astRoot).expr.(*LogicalExpr).right.(*CmpExpr).right)
}
