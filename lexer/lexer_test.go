package lexer

import (
	"testing"
	"github.com/zkpranav/imonkey/token"
)

func TestNextToken(t *testing.T) {
	ip := `
		let valX = 1;
		let valY = 2;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(valX, valY);
	`

	expectedTokens := [] struct {
		expectedTokenType token.TokenType
		expectedLiteral string
	} {
		{ token.LET, "let" },
		{ token.IDENT, "valX" },
		{ token.ASSIGN, "=" },
		{ token.INT, "1" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "valY" },
		{ token.ASSIGN, "=" },
		{ token.INT, "2" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "add" },
		{ token.ASSIGN, "=" },
		{ token.FUNCTION, "fn" },
		{ token.LPAREN, "(" },
		{ token.IDENT, "x" },
		{ token.COMMA, "," },
		{ token.IDENT, "y" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.IDENT, "x" },
		{ token.PLUS, "+" },
		{ token.IDENT, "y" },
		{ token.SEMICOLON, ";" },
		{ token.RBRACE, "}" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "result" },
		{ token.ASSIGN, "=" },
		{ token.IDENT, "add" },
		{ token.LPAREN, "(" },
		{ token.IDENT, "valX" },
		{ token.COMMA, "," },
		{ token.IDENT, "valY" },
		{ token.RPAREN, ")" },
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
