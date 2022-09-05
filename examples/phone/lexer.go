package phone

import "fmt"

type Lexer struct {
	input       []byte
	pos         int
	phoneNumber PhoneNumber
}

func (l *Lexer) Error(e string) {
	fmt.Println(e)
}

func (l *Lexer) Lex(lval *yySymType) int {
	if l.pos >= len(l.input) {
		return 0
	}

	ch := l.input[l.pos]

	switch ch {
	case 0:
		return 0
	case '(', '.', ')', ' ', '-':
		l.pos++
		return int(ch)
	}

	if isDigit(ch) {
		lval.byt = ch
		l.pos++
		return int(ch)
	}

	return 0
}

func isDigit(byt byte) bool {
	return '0' <= byt && byt <= '9'
}
