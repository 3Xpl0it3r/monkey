package parser

import (
	"fmt"

	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// errors

// Errors return the errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekTokenError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, but got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) currTokenError(t token.TokenType) {
	msg := fmt.Sprintf("expect current token to be %s, but got %s instead", t, p.curToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", t))
}

func (p *Parser) noInfixParseFnError(t token.TokenType) {
	p.errors = append(p.errors, fmt.Sprintf("no infix parse function for %s found", t))
}
