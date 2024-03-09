package parser

import (
	"log"
	"testing"
	"github.com/zkpranav/imonkey/ast"
	"github.com/zkpranav/imonkey/lexer"
)

func TestVariableBinding(t *testing.T) {
	ip := `
		let x = 5;
		let y = 10;
		let foo = 420;
	`

	l := lexer.New(ip)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		log.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		log.Fatalf("incorrect number of statements. expected=3, got=%d", len(program.Statements))
	}

	testTable := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, test := range testTable {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
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