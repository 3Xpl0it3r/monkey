package parser

import (
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// parse

// curTokenIs 判断当前token是不是符合期望的tokenType
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs 判断下一个tokenType 是不是符合期望的tokenType
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek is used to check the type of peek token is expected or not
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekTokenError(t)
		return false
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
