package ast

import (
	"testing"

	"github.com/3Xpl0it3r/monkey/pkg/token"
)

func TestString(t *testing.T) {
	// let myVar = anotherVar;
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name:  &Identifier{Token: token.Token{Type: token.IDENT, Literal: "myVar"}, Value: "myVar"},
				// for Identifier
				Value: &Identifier{Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, Value: "anotherVar"},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Fatalf("program String Failed: expected %s ,but got %#q", "let myVar = anotheVar;", program.String())
	}
}
