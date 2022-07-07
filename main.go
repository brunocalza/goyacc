package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/goyacc@master -l -o yy_parser.go grammar.y

func main() {

	lexer := &Lexer{input: []byte("800.555.1234")}
	yyDebug = 1
	yyParse(lexer)

	fmt.Println(lexer.phoneNumber)

}
