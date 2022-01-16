package parser

import (
	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// parseReturnStatement  解析return语句
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
