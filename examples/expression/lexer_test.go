package expression

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `2 + 2`

	expTokens := []int{
		INTEGER, '+', INTEGER,
	}

	lval := &yySymType{}

	lexer := &Lexer{}
	lexer.input = []byte(input)
	lexer.readByte()

	token, i := lexer.Lex(lval), 0
	for token != 0 {
		fmt.Println(token)
		if token != expTokens[i] {
			t.Fatalf("expected %d, got %d, at index %d", expTokens[i], token, i)
		}
		token, i = lexer.Lex(lval), i+1
	}
}
