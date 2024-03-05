package lexer

import "github.com/zkpranav/imonkey/token"

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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token {
		Type: tokenType,
		Literal: string(ch),
	}
}

/*
* Tokenizes the character under examination
*/
func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	switch l.ch {
		case '=':
			tk = newToken(token.ASSIGN, l.ch)
		case '+':
			tk = newToken(token.PLUS, l.ch)
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
	}

	l.readChar()
	return tk
}