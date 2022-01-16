package parser

import (
	"testing"

	"github.com/3Xpl0it3r/monkey/pkg/lexer"
)

// Let 语句测试
func TestLetStatements(t *testing.T) {
	input := `
	let x  = 5;
	let y = 10;
	let foobar = 83354;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram () return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements doesn;t contain 3 statement, got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}
