package ast

import "gompiler/token"

/* define the structures and interface for AST */

type Node interface {
	TokenLiteral() string
}

/* Difference between `Statement` and `Expression` is whether producing value */
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/* root node of every AST */
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
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatement struct {
	Token token.Token // the token.LET token
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode() {}
func (ls *ReturnStatement) TokenLiteral() string {
	return ls.Token.Literal
}


type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
