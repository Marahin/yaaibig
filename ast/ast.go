package ast

import "github.com/marahin/yaaibig/token"

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

type AddStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (as *AddStatement) statementNode() {}
func (as *AddStatement) TokenLiteral() string { return as.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}