package expression

import (
	"bytes"
	"fmt"
)

type Lexer struct {
	input        []byte
	readPosition int  // current position being read
	ch           byte // current char being read

	result float32 // result of evaluating expression
	err    error   // any errors we found along the way
}

func (l *Lexer) Error(e string) {
	l.err = fmt.Errorf("error: %s", e)
}

func (l *Lexer) Lex(lval *yySymType) int {
	if l.ch == 0 {
		return 0
	}

	l.skipWhitespace()

	if isDigit(l.ch) {
		token, literal := l.readNumber()
		lval.str = literal

		return token
	}

	if l.ch == '.' {
		if isDigit(l.peekByte()) {
			var buf bytes.Buffer
			buf.WriteByte('.')
			l.readByte()
			l.readDigits(&buf)
			if l.ch == 'e' || l.ch == 'E' {
				l.readExpoent(&buf)
			}

			lval.str = buf.String()
			return FLOAT
		}

		l.readByte()
		return int('.')
	}

	switch ch := l.ch; ch {
	case '+', '-', '*', '/', '(', ')':
		l.readByte()
		return int(ch)
	}

	return 0
}

func isDigit(byt byte) bool {
	return '0' <= byt && byt <= '9'
}

func (l *Lexer) readNumber() (int, string) {
	var buf bytes.Buffer
	isFloat := false

	l.readDigits(&buf)
	if l.ch == '.' {
		isFloat = true
		buf.WriteByte(l.ch)
		l.readByte()
		l.readDigits(&buf)
	}

	if l.ch == 'e' || l.ch == 'E' {
		isFloat = true
		l.readExpoent(&buf)
	}

	if isFloat {
		return FLOAT, buf.String()
	}

	return INTEGER, buf.String()
}

func (l *Lexer) readDigits(buf *bytes.Buffer) {
	for isDigit(l.ch) {
		buf.WriteByte(l.ch)
		l.readByte()
	}
}

func (l *Lexer) readExpoent(buf *bytes.Buffer) {
	buf.WriteByte(l.ch)
	l.readByte()
	if l.ch == '+' || l.ch == '-' {
		buf.WriteByte(l.ch)
		l.readByte()
	}
	l.readDigits(buf)
}

func (l *Lexer) readByte() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.readPosition += 1
}

func (l *Lexer) peekByte() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readByte()
	}
}
