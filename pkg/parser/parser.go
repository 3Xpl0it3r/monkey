package parser

import (
	"fmt"

	"github.com/3Xpl0it3r/monkey/pkg/ast"
	"github.com/3Xpl0it3r/monkey/pkg/lexer"
	"github.com/3Xpl0it3r/monkey/pkg/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: make([]string, 0)}
	// Read two tokens, so curToken  and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p

}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = make([]ast.Statement, 0)
	for !p.curTokenIs(token.EOF) {
		vtmt := p.parseStatement()
		if vtmt != nil {
			program.Statements = append(program.Statements, vtmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	var ret ast.Statement
	switch p.curToken.Type {
	case token.LET:
		ret = p.parseLetStatement()
	case token.RETURN:
		ret = p.parseReturnStatement()
	default:
		ret = nil
	}
	return ret
}

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

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs
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
