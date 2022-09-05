package expression

func Parse(expression string) (float32, error) {
	lexer := &Lexer{input: []byte(expression)}
	lexer.readByte()
	yyParse(lexer)

	return lexer.result, lexer.err
}
