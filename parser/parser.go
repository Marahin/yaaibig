package parser

import (
	"github.com/marahin/yaaibig/ast"
	"github.com/marahin/yaaibig/lexer"
	"github.com/marahin/yaaibig/token"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken() // we read twice so both curToken and peekToken is set
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	 program := &ast.Program{}
	 program.Statements = []ast.Statement{}

	 for p.curToken.Type != token.EOF {
	 	stmt := p.parseStatement()
	 	if stmt != nil {
	 		program.Statements = append(program.Statements, stmt)
	 	}
	 	p.nextToken()
	 }

	 return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.ADD:
		return p.parseAddStatement()
	default: 
		return nil
	}
}

func (p *Parser) parseAddStatement() *ast.AddStatement {
	stmt := &ast.AddStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal }

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}