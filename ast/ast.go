package ast

import "friston/lexer"

// Visitor interface (all other visitors must implement this)
type Visitor interface {
	visitBinary(b Binary) interface{}
	visitLogic(l Logic) interface{}
	visitUnary(u Unary) interface{}
	visitGroup(g Group) interface{}
	visitLiteral(l Literal) interface{}
	visitVariable(vr Variable) interface{}
	visitAssignment(a Assignment) interface{}
	visitCall(c Call) interface{}

	visitExprStmt(e ExprStmt)
	visitIfStmt(stmt IfStmt)
	visitWhileStmt(stmt WhileStmt)
	visitVarDecl(d VarDecl)
	visitBlock(b Block)
}

// Node types:
type Expression interface {
	Accept(v Visitor) interface{}
}

type Binary struct {
	X Expression
	Op lexer.Token
	Y Expression
}

func (b Binary) Accept(v Visitor) interface{} {
	return v.visitBinary(b)
}

type Logic struct {
	X Expression
	Op lexer.Token
	Y Expression
}

func (l Logic) Accept(v Visitor) interface{} {
	return v.visitLogic(l)
}

type Unary struct {
	Op lexer.Token
	X Expression
}

func (u Unary) Accept(v Visitor) interface{} {
	 return v.visitUnary(u)
}

type Group struct {
	Left lexer.Token
	X Expression
	Right lexer.Token
}

func (g Group) Accept(v Visitor) interface{} {
	 return v.visitGroup(g)
}

type Literal struct {
	X lexer.Token
}

func (l Literal) Accept(v Visitor) interface{} {
	 return v.visitLiteral(l)
}

type Variable struct {
	Name lexer.Token
}

func (vr Variable) Accept(v Visitor) interface{} {
	return v.visitVariable(vr)
}

type Assignment struct {
	Name lexer.Token
	Value Expression
}

func (a Assignment) Accept(v Visitor) interface{} {
	return v.visitAssignment(a)
}

type Call struct {
	Callee Expression
	Paren lexer.Token
	Arguments []Expression
}

func (c Call) Accept(v Visitor) interface{} {
	return v.visitCall(c)
}

//Statement types:

type Statement interface {
	Accept(v Visitor) interface{}
}

type ExprStmt struct {
	Expr Expression
}

func (e ExprStmt) Accept(v Visitor) interface{} {
	v.visitExprStmt(e)
	return nil
}

type IfStmt struct {
	Condition Expression
	ThenBranch Statement
	ElseBranch Statement
}

func (i IfStmt) Accept(v Visitor) interface{} {
	v.visitIfStmt(i)
	return nil
}

type WhileStmt struct {
	Condition Expression
	LoopBranch Statement
}

func (w WhileStmt) Accept(v Visitor) interface {} {
	v.visitWhileStmt(w)
	return nil
}

type VarDecl struct {
	Name lexer.Token
	Initializer Expression
}

func (d VarDecl) Accept(v Visitor) interface{} {
	v.visitVarDecl(d)
	return nil
}

type Block struct {
	Stmts []Statement
}

func (b Block) Accept(v Visitor) interface{} {
	v.visitBlock(b)
	return nil
}