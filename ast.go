package main

import "fmt"

type expr interface {
	ToString() string
}

type astRoot struct {
	expr expr
}

type Token struct {
	token   int
	literal []byte
}

type CmpExpr struct {
	op    CmpOperatorNode
	left  expr
	right expr
}

func (e *CmpExpr) ToString() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", e.left.ToString(), e.op.ToString(), e.right.ToString())
}

type AndExpr struct {
	token Token
	left  expr
	right expr
}

func (e *AndExpr) ToString() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", e.left.ToString(), e.token.literal, e.right.ToString())
}

type OrExpr struct {
	token Token
	left  expr
	right expr
}

func (e *OrExpr) ToString() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", e.left.ToString(), e.token.literal, e.right.ToString())
}

type NotExpr struct {
	token Token
	expr  expr
}

func (e *NotExpr) ToString() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s %s", e.token.literal, e.expr.ToString())
}

type BoolValue struct {
	token Token
	val   bool
}

func (v BoolValue) ToString() string {
	return string(v.token.literal)
}

type NullValue struct {
	token Token
}

func (v NullValue) ToString() string {
	return string(v.token.literal)
}

// ValueType specifies the type for SQLVal.
type ValueType int

const (
	StrValue = ValueType(iota)
	IntValue
	FloatValue
	HexNumValue
	BlobValue
)

// SQLValue represents a single value.
type SQLValue struct {
	Token Token
	Type  ValueType
	Val   []byte
}

func (v *SQLValue) ToString() string {
	return string(v.Token.literal)
}

type Column struct {
	name string
}

func (c *Column) ToString() string {
	if c == nil {
		return ""
	}
	return c.name
}

type TableRef struct {
	name string
}

func (c *TableRef) ToString() string {
	if c == nil {
		return ""
	}
	return c.name
}

type Number struct {
	val string
}

func (n *Number) ToString() string {
	if n == nil {
		return ""
	}
	return n.val
}

// ConvertType specifies the type for ConvertExpr.
type ConvertType string

const (
	NoneStr    = ConvertType("NONE")
	RealStr    = ConvertType("REAL")
	NumericStr = ConvertType("NUMERIC")
	TextStr    = ConvertType("TEXT")
	IntegerStr = ConvertType("INTEGER")
)

type ConvertExpr struct {
	expr expr
	typ  ConvertType
}

func (e *ConvertExpr) ToString() string {
	return fmt.Sprintf("CAST (%s AS %s)", e.expr.ToString(), string(e.typ))
}

// IsExprType specifies the type for IsExprType.
type IsExprType string

const (
	IsStr      = IsExprType("IS")
	IsNullStr  = IsExprType("ISNULL")
	NotNullStr = IsExprType("NOTNULL")
)

type IsExpr struct {
	lhs expr
	rhs expr
	typ IsExprType
}

func (e *IsExpr) ToString() string {
	if e.typ == IsNullStr || e.typ == NotNullStr {
		return fmt.Sprintf("%s %s", e.lhs.ToString(), e.typ)
	}
	return fmt.Sprintf("%s %s %s", e.lhs.ToString(), e.typ, e.rhs.ToString())
}
