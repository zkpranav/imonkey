package parser

import (
	"github.com/zkpranav/imonkey/ast"
	"github.com/zkpranav/imonkey/lexer"
	"github.com/zkpranav/imonkey/token"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	lookAheadToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := Parser {
		l: l,
	}

	// Sets curToken and peekToken
	p.nextToken()
	p.nextToken()

	return &p
}

func (p *Parser) nextToken() {
	p.curToken = p.lookAheadToken
	p.lookAheadToken = p.l.NextToken()
}

/*
* Implements a Recursive Decent parser.
*
* Parse one token --> realize --> within context parse one token --> realize --> ...
* This defines the recursive nature of this parser and provides intuition for how it generates the AST top-down.
*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
		case token.LET:
			return p.parseLetStatement()
		default:
			return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: p.curToken,
	}

	if !p.expectPeekAndAdvance(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}

	if !p.expectPeekAndAdvance(token.ASSIGN) {
		return nil
	}
	
	// TODO: Expect expression & store expression

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func (p *Parser) curTokenIs(expectedType token.TokenType) bool {
	return p.curToken.Type == expectedType
}

func (p *Parser) peekTokenIs(expectedType token.TokenType) bool {
	return p.lookAheadToken.Type == expectedType
}

func (p *Parser) expectPeekAndAdvance(expectedType token.TokenType) bool {
	if p.peekTokenIs(expectedType) {
		p.nextToken()
		return true
	} else {
		return false
	}
}