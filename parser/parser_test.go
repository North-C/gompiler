package parser

import (
	"fmt"
	"gompiler/ast"
	"gompiler/lexer"
	"testing"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input        string
		leftValue 	 int64
		operator     string
		rightValue   int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
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

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not  '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}

}

func testIntegerLiteral(t *testing.T, e ast.Expression, v int64) bool {
	integ, ok := e.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("e not ast.IntegerLiteral. got=%T", e)
		return false
	}

	if integ.Value != v {
		t.Errorf("integ.Value not %d, got=%d", v, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", v) {
		t.Errorf("integ.Tokenliteral not %d. got=%s", v, integ.TokenLiteral())
		return false
	}

	return true
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
