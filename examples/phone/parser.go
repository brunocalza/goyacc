package phone

func Parse(expression string) PhoneNumber {
	lexer := &Lexer{input: []byte(expression)}
	yyParse(lexer)

	return lexer.phoneNumber
}
