package lexer

import (
	"github.com/zkpranav/imonkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	ip := `
		let valX = 1;
		let valY = 2;

		let add = fn(x, y) {
			return x + y;
		};

		let result = add(valX, valY);

		!-42*/;
		5 < 10 * 1 > 5;

		if (0 < 1) {
			return true;
		} else {
			return false;
		}
	`

	expectedTokens := []struct {
		expectedTokenType token.TokenType
		expectedLiteral   string
	}{
		{token.LET, "let"},
		{token.IDENT, "valX"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "valY"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "valX"},
		{token.COMMA, ","},
		{token.IDENT, "valY"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.INT, "42"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.ASTERISK, "*"},
		{token.INT, "1"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "0"},
		{token.LT, "<"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
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
