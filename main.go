package main

import (
	"fmt"
)

//go:generate go run golang.org/x/tools/cmd/goyacc@master -l -o yy_parser.go grammar.y

func main() {
	yyErrorVerbose = true
	l := Lexer{input: []byte("_1 OR b AND 12121 = 3")}
	l.readByte()

	yyParse(&l)

	fmt.Println(Deparse(l.ast))
}
