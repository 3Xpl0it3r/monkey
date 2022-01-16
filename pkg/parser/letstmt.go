/*
let statement
*/

package parser

import (
	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

// parseLetStatement 解析Let语句
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	// identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Skipping the expression

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}
