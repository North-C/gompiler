package parser

import (
	"gompiler/ast"
	"gompiler/lexer"
	"testing"
)

func TestReturnStatements(t *testing.T) {
	input := `
			return 10;
			return x;
			return (x+1);
			`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 { 
		t.Fatalf("program.Statements does not contain 3 statemenst. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok{
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return"{
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
						returnStmt.TokenLiteral())			
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser){
	errors := p.Errors()

	if len(errors) == 0{
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors{
		t.Errorf("parser error:%q", msg)
	}
	t.FailNow()
}