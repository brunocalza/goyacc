package main

type CmpOperatorNode interface {
	CmpOperator() CmpOperator
	ToString() string
}

type EqualOperator struct {
	token Token
}

func (op *EqualOperator) CmpOperator() CmpOperator {
	return EQ_OP
}

func (op *EqualOperator) ToString() string {
	return string(op.token.literal)
}

type LessThanOperator struct {
	token Token
}

func (op *LessThanOperator) CmpOperator() CmpOperator {
	return LT_OP
}

func (op *LessThanOperator) ToString() string {
	return string(op.token.literal)
}

type GreaterThanOperator struct {
	token Token
}

func (op *GreaterThanOperator) CmpOperator() CmpOperator {
	return GT_OP
}

func (op *GreaterThanOperator) ToString() string {
	return string(op.token.literal)
}

type LessEqualOperator struct {
	token Token
}

func (op *LessEqualOperator) CmpOperator() CmpOperator {
	return LE_OP
}

func (op *LessEqualOperator) ToString() string {
	return string(op.token.literal)
}

type GreaterEqualOperator struct {
	token Token
}

func (op *GreaterEqualOperator) CmpOperator() CmpOperator {
	return GE_OP
}

func (op *GreaterEqualOperator) ToString() string {
	return string(op.token.literal)
}

type NotEqualOperator struct {
	token Token
}

func (op *NotEqualOperator) CmpOperator() CmpOperator {
	return NE_OP
}

func (op *NotEqualOperator) ToString() string {
	return string(op.token.literal)
}
