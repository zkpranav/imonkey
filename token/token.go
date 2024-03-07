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
	EQ = "=="
	NOT_EQ = "!="

	// Keywords
	LET = "LET"
	FUNCTION = "FUNCTION"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

/*
* The look-up table for keywords.
*/
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}