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

	// Keywords
	LET = "LET"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func GetIdentType(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}