package phone

func Parse(expression string) PhoneNumber {
	yyErrorVerbose = true
	lexer := &Lexer{input: []byte(expression)}
	yyParse(lexer)

	return lexer.phoneNumber
}
