package parser

import (
	"fmt"
	"strconv"

	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// register prefixParseFn associate with token type
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// register infixParseFn associate with token type
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// parseIntegerLiteral 解析整型
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer", p.curToken.Literal))
		return nil
	}
	lit.Value = value
	return lit
}

// 前缀解析
func (p *Parser) parsePrefixExpression() ast.Expression {
	// -1 + 2 ; curToken -  peekToken 1
	express := &ast.PrefixExpression{
		Token:    p.curToken,	// -
		Operator: string(p.curToken.Literal), // -
	}
	p.nextToken()	// curToken 1 peekToken +
	express.Right = p.parseExpression(PREFIX)

	return express
}

// 中缀解析
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	express := &ast.InfixExpression{
		Token:    p.curToken,	//
		Left:     left,		//
		Operator: string(p.curToken.Literal),	//
	}
	precedence := p.curPrecedence()	// get precedence
	p.nextToken()		// curToken is 2, peekToken is
	express.Right = p.parseExpression(precedence) 		// right is 2
	return express
}
