package ast

import (
	"bytes"

	"github.com/3Xpl0it3r/monkey/pkg/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// statement don't produce value
type Statement interface {
	Node
	statementNode()
}

// expression can produce value
type Expression interface {
	Node
	expressionNode()
}

var _ Node = new(Program)

// the root of ast
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement represent the spec of let statement
// Let 语句

var _ Statement = new(LetStatement)

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // the identifier in the left let statement
	Value Expression  // pointer an expression node
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier标识符， 一个语句有三部分组成 identifieer operator expression/identifier
type Identifier struct {
	Token token.Token
	Value string
}

// identifier 继承了Expression的接口，例如下面例子, identidier 也有可能会成为expression
// let addFn = func(int,int)int
// let sum = addFn(1,2)
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement return 语句
type ReturnStatement struct {
	Token       token.Token // the return token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement expression表达式
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

var _ Expression = new(IntegerLiteral)

//  整型表达式 <integer>;
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) String() string       { return il.Token.Literal }
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}

// 前缀表达式 //  as the following format
// <prefix operator><expression>;
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

var _ Expression = new(PrefixExpression)

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

//  中缀表达式
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
