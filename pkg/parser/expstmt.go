package parser

import (
	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// parseExpressionStatement 解析expression语句
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression 根据tokenType 生成对应的expression
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	//
	leftExp := prefix()	// -
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {	// peekPrecedcence +
		//  这里peekTokenIs 其实不一定需要，因为p.peekPrecedence 没有记录，会返回最低优先级, 放在这里是为了更容易理解
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()	//
		leftExp = infix(leftExp)
	}

	return leftExp
}
