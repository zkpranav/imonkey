package lexer

import (
	"testing"
	"github.com/zkpranav/imonkey/token"
)

func TestNextToken(t *testing.T) {
	ip := "=+(){},;"

	expectedTokens := [] struct {
		expectedTokenType token.TokenType
		expectedLiteral string
	} {
		{ token.ASSIGN, "=" },
		{ token.PLUS, "+" },
		{ token.LPAREN, "(" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.RBRACE, "}" },
		{ token.COMMA, "," },
		{ token.SEMICOLON, ";" },
		{ token.EOF, "" },
	}

	l := New(ip)
	for index, expectedTk := range expectedTokens {
		tk := l.NextToken()
		if tk.Type != expectedTk.expectedTokenType {
			t.Fatalf("test[%d]: wrong tokentype. expected=%q, got=%q", index, expectedTk.expectedTokenType, tk.Type)
		}
		if tk.Literal != expectedTk.expectedLiteral {
			t.Fatalf("test[%d]: wrong literal. expected=%q, got=%q", index, expectedTk.expectedLiteral, tk.Literal)
		}
	}
}
