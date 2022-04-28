package main

import "fmt"

type LogicalOperator = int

const (
	AND LogicalOperator = iota
	OR
)

var logicOps = map[string]LogicalOperator{
	"AND": AND,
	"OR":  OR,
}

type CmpOperator = int

const (
	EQ CmpOperator = iota
	NE
	LT
	LE
	GT
	GE
)

var cmpOps = map[string]CmpOperator{
	"=":  EQ,
	"!=": NE,
	"<>": NE,
	"<":  LT,
	"<=": LE,
	">":  GT,
	">=": GE,
}

const EOF = 0

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte

	ast expr
}

func (i *Lexer) Error(e string) {
	fmt.Println(e)
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.skipWhitespace()

	if l.ch == 0 {
		return EOF
	}

	if isComparison(l.ch) {
		literal := l.readComparison()

		if cmpOperator, ok := cmpOps[literal]; ok {
			lval.cmpOp = cmpOperator
			l.readChar()
			return CMPOP
		}

		lval.string = literal
		return 2
	}

	if isLetter(l.ch) {
		literal := l.readIdentifier()

		if logicalOperator, ok := logicOps[literal]; ok {
			lval.logicOp = logicalOperator
			return LOP
		}

		lval.string = literal
		l.readChar()
		return IDENTIFIER
	}

	if isDigit(l.ch) {
		lval.string = l.readNumber()
		l.readChar()
		return NUMBER
	}

	return 2
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readComparison() string {
	position := l.position
	for isComparison(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isComparison(ch byte) bool {
	return ch == '=' || ch == '!' || ch == '<' || ch == '>'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
