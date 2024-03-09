package ast

import (
	"github.com/zkpranav/imonkey/token"
)

/*
* Statements do not produce a value. eg: return x;
* Expressions do produce a value. eg: 5 (produces the value 5)
*
* These form the basis of all nodes in the AST.
*/

/*
* Base interface for all nodes in the AST.
*/
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*
* The implicit root node of the AST.
*/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET
	Name *Identifier // name the value is being bound to
	Value Expression // the value being bound
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

/*
* Although some identifiers do not produce a value (eg: let x = 10;)
* the Identifier is defined as an Expression still because some do produce a value.
* eg: let x = 5 * f;
*/
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}