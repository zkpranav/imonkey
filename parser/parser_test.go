package parser

import (
	"github.com/zkpranav/imonkey/ast"
	"github.com/zkpranav/imonkey/lexer"
	"log"
	"testing"
)

func TestVariableBinding(t *testing.T) {
	ip := `
		return 1;
		return 420;
		return 1 * 5;
	`

	l := lexer.New(ip)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		log.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		log.Fatalf("incorrect number of statements. expected=3, got=%d", len(program.Statements))
	}

	for _, s := range program.Statements {
		testReturnStatement(t, s)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not \"let\". expected=let, got=%s", s.TokenLiteral())
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not an ast.LetStatement. got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("s.Name.Value is incorrect. expected=%s, got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("s.Name.TokenLiteral() returned incorrect value. expected=%s, got=%s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func testReturnStatement(t *testing.T, s ast.Statement) {
	rs, ok := s.(*ast.ReturnStatement)

	if !ok {
		t.Errorf("rs is not an ast.ReturnStatement. got=%T", rs)
		return
	}

	if rs.TokenLiteral() != "return" {
		t.Errorf("rs.TokenLiteral incorrect. expected=return, got=%s", rs.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser errors: %d", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
