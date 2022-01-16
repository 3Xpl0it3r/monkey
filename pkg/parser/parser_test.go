package parser

import (
	"testing"

	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/lexer"
)

// Identifier 表达式测试
func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statement, expected 1, but got=%d\n", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] expected *ast.ExpressionStatement, but got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not *ast.Identifier. got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value is not foobar, but got=%s", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() is not foobar, but got=%s", ident.TokenLiteral())
	}
}

// 整型测试
func TestIntegerLiteralExpression(t *testing.T) {
	input := `5;`
	p := New(lexer.New(input))
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not engouh statement , expected 1 , but got=%d\n", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] expecpted *ast.ExpressionStament, but got=%T", program.Statements[0])
	}
	integer, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp is not *ast.IntegerLiteral, got=%T\n", stmt.Expression)
	}

	if integer.Value != 5 {
		t.Fatalf("integer.value is not 5, but got=%d", integer.Value)
	}

	if integer.TokenLiteral() != "5" {
		t.Fatalf("Integer.TokenLiteral is not '5', but got %s", integer.String())
	}
}

// 前缀表达式测试
func TestParsePrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParseErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statement, expected 1, but go=%d\n", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Program.Statements[0] expected *ast.ExpressionStatement, but got=%T", program.Statements[0])
		}
		pe, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Express expected got *ast.PrefixExpression, but got %T", stmt.Expression)
		}

		if pe.Operator != tt.operator {
			t.Fatalf("exp operator expected %s, but got %s", tt.operator, pe.Operator)
		}
		if !testIntegerLiteral(pe.Right, tt.integerValue) {
			t.Fatalf("testIntegerLiteral failed at %s, expect got %d", pe.Right.String(), tt.integerValue)
		}
	}
}

func TestParseInfixExpression(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}
	for _, tt := range infixTests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParseErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program has not engouh statement, expected 1 ,but got=%d\n", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Program.Statements[0] expected got *ast.ExpressionStatement, but got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp expected got *ast.InfixExpression, but got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp operator expected %s, but got %s", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral( exp.Left, tt.leftValue) {
			t.Fatalf("left value testIntegerLiteral failed,")
		}
		if !testIntegerLiteral(exp.Right, tt.rightValue) {
			t.Fatalf("right value testIntegerLiteral failed")
		}

	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
	}

	for index, tt := range tests {
		p := New(lexer.New(tt.input))
		program := p.ParseProgram()
		checkParseErrors(t, p)


		if tt.expected != program.String() {
			t.Fatalf("item[%d]->expect %s, but got %s", index, tt.expected, program.String())
		}
	}

}

func testIntegerLiteral( il ast.Expression, value int64) bool {
	exp, ok := il.(*ast.IntegerLiteral)
	if !ok {
		return false
	}
	if exp.Value != value {
		return false
	}
	return true
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', but got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T\n", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("LetStatement.Name.Value not '%s',got='%s'\n", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s', got=%s\n", name, letStmt.Name.Value)
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parse has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error: %q\n", msg)
	}
	t.FailNow()
}
