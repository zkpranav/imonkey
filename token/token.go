package token

type TokenType string // TODO: change to an enum-like

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifies & literals
	IDENT = "IDENT"
	INT = "INT"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	// Keywords
	LET = "LET"
	FUNCTION = "FUNCTION"
)

/*
* The look-up table for keywords.
*/
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}