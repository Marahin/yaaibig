package lexer

import (
	"github.com/marahin/yaaibig/lexer"
	"github.com/marahin/yaaibig/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `MOV A 1
mov B 1
ADD 2 3
add 23 32`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.MOVE, "MOV"},
		{token.IDENTIFIER, "A"},
		{token.INTEGER, "1"},

		{token.MOVE, "mov"},
		{token.IDENTIFIER, "B"},
		{token.INTEGER, "1"},

		{token.ADD, "ADD"},
		{token.INTEGER, "2"},
		{token.INTEGER, "3"},

		{token.ADD, "add"},
		{token.INTEGER, "23"},
		{token.INTEGER, "32"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - tokenType wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
