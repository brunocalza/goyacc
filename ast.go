package main

type expr interface{}

type astRoot struct {
	expr expr
}

type CmpExpr struct {
	op    CmpOperator
	left  expr
	right expr
}

type LogicalExpr struct {
	op    LogicalOperator
	left  expr
	right expr
}

type column struct {
	name string
}

type number struct {
	val string
}
