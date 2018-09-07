package parser

import (
	"testing"

	"github.com/marahin/yaaibig/ast"
	"github.com/marahin/yaaibig/lexer"
)

func TestAddStatements(t *testing.T) {
	input := `
ADD a 10
ADD b 20
ADD c 30
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"a"},
		{"b"},
		{"c"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "ADD" {
		t.Errorf("s.TokenLiteral not 'ADD'. got=%q", s.TokenLiteral())
		return false
	}

	addStmt, ok := s.(*ast.AddStatement)
	if !ok {
		t.Errorf("s not *ast.AddStatement. got=%T", s)
		return false
	}

	if addStmt.Name.Value != name {
		t.Errorf("addStmt.Name.Value not '%s'. got=%s", name, addStmt.Name.Value)
		return false
	}

	return true
}