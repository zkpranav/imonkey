package lexer

import (
	"github.com/zkpranav/imonkey/token"
)

/*
* Supports only ASCII text as input.
*
* TODO: Proccess the input as a stream of bytes from a file.
*/
type Lexer struct {
	input string // The source
	pointer int // Pointer to the current character ch
	lookAhead int // The look ahead pointer
	ch byte // The character being examined
}

func New(input string) *Lexer {
	l := &Lexer {
		input: input,
		pointer: 0,
		lookAhead: 0,
		ch: 0,
	}
	l.readChar()

	return l
}

/*
* Reads the character at lookAhead and advances lookAhead.
*/
func (l *Lexer) readChar() {
	if l.lookAhead >= len(l.input) {
		// We're at or beyond EOF
		l.ch = 0 // \0
	} else {
		l.ch = l.input[l.lookAhead]
	}

	l.pointer = l.lookAhead
	l.lookAhead += 1
}

/*
* Reads a characters until it encounters a character that is not a valid identifier character
*/
func (l *Lexer) readIdentifier() string {
	start := l.pointer
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[start: l.pointer]
}

func (l *Lexer) readNumber() string {
	start := l.pointer
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[start: l.pointer]
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
* Tokenizes the character under examination
*/
func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.eatWhitespace()

	switch l.ch {
		case '=':
			tk = newToken(token.ASSIGN, l.ch)
		case '+':
			tk = newToken(token.PLUS, l.ch)
		case '-':
			tk = newToken(token.MINUS, l.ch)
		case '!':
			tk = newToken(token.BANG, l.ch)
		case '*':
			tk = newToken(token.ASTERISK, l.ch)
		case '/':
			tk = newToken(token.SLASH, l.ch)
		case '<':
			tk = newToken(token.LT, l.ch)
		case '>':
			tk = newToken(token.GT, l.ch)
		case ',':
			tk = newToken(token.COMMA, l.ch)
		case ';':
			tk = newToken(token.SEMICOLON, l.ch)
		case '(':
			tk = newToken(token.LPAREN, l.ch)
		case ')':
			tk = newToken(token.RPAREN, l.ch)
		case '{':
			tk = newToken(token.LBRACE, l.ch)
		case '}':
			tk = newToken(token.RBRACE, l.ch)
		case 0:
			tk.Type = token.EOF
			tk.Literal = ""
		default:
			if isLetter(l.ch) {
				tk.Literal = l.readIdentifier()
				tk.Type = token.LookupIdent(tk.Literal)
				
				return tk // returns here to avoid the upcoming readChar
			} else if isDigit(l.ch) {
				tk.Type = token.INT
				tk.Literal = l.readNumber()

				return tk
			} else {
				tk = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tk
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token {
		Type: tokenType,
		Literal: string(ch),
	}
}

func isLetter(ch byte) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_' { // allows underscore
		return true
	}

	return false
}

func isDigit(ch byte) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}

	return false
}