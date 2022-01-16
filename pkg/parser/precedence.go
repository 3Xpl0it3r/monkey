package parser

import "github.com/3Xpl0it3r/monkey/pkg/token"

const (
	// precedence define the order or operation
	//  定义运算符号的优先级
	_ int = iota
	LOWEST
	EQUALS     // ==
	LESSGRATER // > or <
	SUM        // +
	PRODUCT    // *
	PREFIX     // -X OR !X
	CALL       // myFunction(X)
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGRATER,
	token.GT:       LESSGRATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
