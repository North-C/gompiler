package parser

import (
	"gompiler/ast"
	"gompiler/lexer"
	"testing"
)

func TestIntregerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not ast.Identifier. got=%T", program.Statements[0])
	}
	if literal.Value != 5 {
		t.Fatalf("ident.Value is not %s. got=%d", "5", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral not %s. got %s", "5", literal.TokenLiteral())
	}
}

/* Print the error messages */
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
