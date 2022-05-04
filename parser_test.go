package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleExpr(t *testing.T) {
	type testCase struct {
		name        string
		expr        string
		expectedAst *astRoot
	}

	tests := []testCase{
		{
			name: "true",
			expr: "true",
			expectedAst: &astRoot{
				expr: BoolValue{
					token: Token{TRUE, []byte("true")},
					val:   true,
				},
			},
		},
		{
			name: "false",
			expr: "false",
			expectedAst: &astRoot{
				expr: BoolValue{
					token: Token{FALSE, []byte("false")},
					val:   false,
				},
			},
		},
		{
			name: "null",
			expr: "null",
			expectedAst: &astRoot{
				expr: &NullValue{
					token: Token{NULL, []byte("null")},
				},
			},
		},
		{
			name: "string",
			expr: "'blablabla'",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{STRING, []byte("'blablabla'")},
					Type:  StrValue,
					Val:   []byte("blablabla"),
				},
			},
		},
		{
			name: "integral",
			expr: "123",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{INTEGRAL, []byte("123")},
					Type:  IntValue,
					Val:   []byte("123"),
				},
			},
		},
		{
			name: "float",
			expr: ".2",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{FLOAT, []byte(".2")},
					Type:  FloatValue,
					Val:   []byte(".2"),
				},
			},
		},
		{
			name: "float-mantissa",
			expr: "1.2",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{FLOAT, []byte("1.2")},
					Type:  FloatValue,
					Val:   []byte("1.2"),
				},
			},
		},
		{
			name: "float-expoent",
			expr: "1e10",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{FLOAT, []byte("1e10")},
					Type:  FloatValue,
					Val:   []byte("1e10"),
				},
			},
		},
		{
			name: "hexnum",
			expr: "0xA6",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{HEXNUM, []byte("0xA6")},
					Type:  HexNumValue,
					Val:   []byte("0xA6"),
				},
			},
		},
		{
			name: "string-single-quote",
			expr: "'joe''s car'",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{STRING, []byte("'joe''s car'")},
					Type:  StrValue,
					Val:   []byte("joe''s car"),
				},
			},
		},
		{
			name: "blob",
			expr: "X'53514C697465'",
			expectedAst: &astRoot{
				expr: &SQLValue{
					Token: Token{BLOB, []byte("X'53514C697465'")},
					Type:  BlobValue,
					Val:   []byte("X'53514C697465'"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tc testCase) func(t *testing.T) {
			return func(t *testing.T) {
				t.Parallel()
				parser := NewParser()

				ast, err := parser.Parse(tc.expr)

				require.NoError(t, err)
				require.Equal(t, tc.expectedAst, ast)
				require.Equal(t, tc.expr, Deparse(ast))
			}
		}(tc))
	}
}

func TestCastExpr(t *testing.T) {
	type testCase struct {
		name        string
		expr        string
		expectedAst *astRoot
	}

	tests := []testCase{
		{
			name: "cast-to-text",
			expr: "CAST (1 AS TEXT)",
			expectedAst: &astRoot{
				expr: &ConvertExpr{
					expr: &SQLValue{
						Token: Token{token: INTEGRAL, literal: []byte{'1'}},
						Type:  IntValue,
						Val:   []byte{'1'},
					},
					typ: TextStr,
				},
			},
		},
		{
			name: "cast-to-text",
			expr: "CAST (column AS REAL)",
			expectedAst: &astRoot{
				expr: &ConvertExpr{
					expr: &Column{
						name: "column",
					},
					typ: RealStr,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tc testCase) func(t *testing.T) {
			return func(t *testing.T) {
				t.Parallel()
				parser := NewParser()

				ast, err := parser.Parse(tc.expr)

				require.NoError(t, err)
				require.Equal(t, tc.expectedAst, ast)
				require.Equal(t, tc.expr, Deparse(ast))
			}
		}(tc))
	}
}

func TestIsExpr(t *testing.T) {
	type testCase struct {
		name        string
		expr        string
		expectedAst *astRoot
	}

	tests := []testCase{
		{
			name: "expr-is-expr",
			expr: "column IS true",
			expectedAst: &astRoot{
				expr: &IsExpr{
					rhs: BoolValue{
						token: Token{token: TRUE, literal: []byte("true")},
						val:   true,
					},
					lhs: &Column{name: "column"},
					typ: IsStr,
				},
			},
		},
		{
			name: "expr-isnull",
			expr: "column ISNULL",
			expectedAst: &astRoot{
				expr: &IsExpr{
					rhs: &NullValue{
						token: Token{token: ISNULL, literal: []byte("ISNULL")},
					},
					lhs: &Column{name: "column"},
					typ: IsNullStr,
				},
			},
		},
		{
			name: "expr-notnull",
			expr: "column NOTNULL",
			expectedAst: &astRoot{
				expr: &IsExpr{
					rhs: &NullValue{
						token: Token{token: NOTNULL, literal: []byte("NOTNULL")},
					},
					lhs: &Column{name: "column"},
					typ: NotNullStr,
				},
			},
		},
		{
			name: "expr-is-not-null",
			expr: "column IS not null",
			expectedAst: &astRoot{
				expr: &IsExpr{
					rhs: &NotExpr{
						token: Token{token: NOT, literal: []byte("not")},
						expr: &NullValue{
							token: Token{token: NULL, literal: []byte("null")},
						},
					},
					lhs: &Column{name: "column"},
					typ: IsStr,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tc testCase) func(t *testing.T) {
			return func(t *testing.T) {
				t.Parallel()
				parser := NewParser()

				ast, err := parser.Parse(tc.expr)

				require.NoError(t, err)
				require.Equal(t, tc.expectedAst, ast)
				require.Equal(t, tc.expr, Deparse(ast))
			}
		}(tc))
	}
}

func TestExpr(t *testing.T) {
	type testCase struct {
		name        string
		expr        string
		expectedAst *astRoot
	}

	tests := []testCase{
		{
			name: "eq-operator-true",
			expr: "c = true",
			expectedAst: &astRoot{
				expr: &CmpExpr{
					op:    &EqualOperator{Token{int('='), []byte("=")}},
					left:  &Column{"c"},
					right: BoolValue{Token{TRUE, []byte("true")}, true},
				},
			},
		},
		{
			name: "double-eq-operator-true",
			expr: "c == true",
			expectedAst: &astRoot{
				expr: &CmpExpr{
					op:    &EqualOperator{Token{int('='), []byte("==")}},
					left:  &Column{"c"},
					right: BoolValue{Token{TRUE, []byte("true")}, true},
				},
			},
		},
		{
			name: "eq-operator-false",
			expr: "c = false",
			expectedAst: &astRoot{
				expr: &CmpExpr{
					op:    &EqualOperator{Token{int('='), []byte("=")}},
					left:  &Column{"c"},
					right: BoolValue{Token{FALSE, []byte("false")}, false},
				},
			},
		},
		{
			name: "not-operator",
			expr: "not c",
			expectedAst: &astRoot{
				expr: &NotExpr{
					token: Token{NOT, []byte("not")},
					expr:  &Column{"c"},
				},
			},
		},
		{
			name: "bang-operator",
			expr: "! c",
			expectedAst: &astRoot{
				expr: &NotExpr{
					token: Token{NOT, []byte("!")},
					expr:  &Column{"c"},
				},
			},
		},
		{
			name: "and-operator",
			expr: "a AND b",
			expectedAst: &astRoot{
				expr: &AndExpr{
					token: Token{AND, []byte("AND")},
					left:  &Column{"a"},
					right: &Column{"b"},
				},
			},
		},
		{
			name: "or-operator",
			expr: "a OR b",
			expectedAst: &astRoot{
				expr: &OrExpr{
					token: Token{OR, []byte("OR")},
					left:  &Column{"a"},
					right: &Column{"b"},
				},
			},
		},
		{
			name: "not-and-not",
			expr: "not c AND NOT b",
			expectedAst: &astRoot{
				expr: &AndExpr{
					token: Token{AND, []byte("AND")},
					left: &NotExpr{
						token: Token{NOT, []byte("not")},
						expr:  &Column{"c"},
					},
					right: &NotExpr{
						token: Token{NOT, []byte("NOT")},
						expr:  &Column{"b"},
					},
				},
			},
		},
		{
			name: "not-or-not",
			expr: "not c OR NOT b",
			expectedAst: &astRoot{
				expr: &OrExpr{
					token: Token{OR, []byte("OR")},
					left: &NotExpr{
						token: Token{NOT, []byte("not")},
						expr:  &Column{"c"},
					},
					right: &NotExpr{
						token: Token{NOT, []byte("NOT")},
						expr:  &Column{"b"},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tc testCase) func(t *testing.T) {
			return func(t *testing.T) {
				t.Parallel()
				parser := NewParser()

				ast, err := parser.Parse(tc.expr)

				require.NoError(t, err)
				require.Equal(t, tc.expectedAst, ast)
				require.Equal(t, tc.expr, Deparse(ast))
			}
		}(tc))
	}
}

func TestPrecedence(t *testing.T) {
	type testCase struct {
		name        string
		expr        string
		expectedAst *astRoot
	}

	tests := []testCase{
		{
			name: "and-not",
			expr: "a and not b",
			expectedAst: &astRoot{
				expr: &AndExpr{
					token: Token{AND, []byte("and")},
					left:  &Column{"a"},
					right: &NotExpr{
						token: Token{NOT, []byte("not")},
						expr:  &Column{"b"},
					},
				},
			},
		},
		{
			name: "or-and",
			expr: "a OR b AND not c",
			expectedAst: &astRoot{
				expr: &OrExpr{
					token: Token{OR, []byte("OR")},
					left:  &Column{"a"},
					right: &AndExpr{
						token: Token{AND, []byte("AND")},
						left:  &Column{"b"},
						right: &NotExpr{
							token: Token{NOT, []byte("not")},
							expr:  &Column{"c"},
						},
					},
				},
			},
		},
		{
			name: "and-or",
			expr: "a AND b OR not c",
			expectedAst: &astRoot{
				expr: &OrExpr{
					token: Token{OR, []byte("OR")},
					left: &AndExpr{
						token: Token{AND, []byte("AND")},
						left:  &Column{"a"},
						right: &Column{"b"},
					},
					right: &NotExpr{
						token: Token{NOT, []byte("not")},
						expr:  &Column{"c"},
					},
				},
			},
		},
		{
			name: "and-eq",
			expr: "a AND b = c",
			expectedAst: &astRoot{
				expr: &AndExpr{
					token: Token{AND, []byte("AND")},
					right: &CmpExpr{
						op:    &EqualOperator{Token{int('='), []byte("=")}},
						left:  &Column{"b"},
						right: &Column{"c"},
					},
					left: &Column{"a"},
				},
			},
		},
		{
			name: "eq-and",
			expr: "a = b AND c",
			expectedAst: &astRoot{
				expr: &AndExpr{
					token: Token{AND, []byte("AND")},
					left: &CmpExpr{
						op:    &EqualOperator{Token{int('='), []byte("=")}},
						left:  &Column{"a"},
						right: &Column{"b"},
					},
					right: &Column{"c"},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tc testCase) func(t *testing.T) {
			return func(t *testing.T) {
				t.Parallel()
				parser := NewParser()

				ast, err := parser.Parse(tc.expr)

				require.NoError(t, err)
				require.Equal(t, tc.expectedAst, ast)
				require.Equal(t, tc.expr, Deparse(ast))
			}
		}(tc))
	}

}
