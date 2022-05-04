%{
package main

const MaxColumnNameLength = 64

%}

%union{
  bool bool
  string string
  bytes []byte
  expr expr 
  cmpOp CmpOperatorNode
  column *Column
  convertType ConvertType
}

%token <bytes> IDENTIFIER STRING INTEGRAL HEXNUM FLOAT BLOB
%token ERROR 
%token <bytes> TRUE FALSE NULL
%token <empty> '(' ',' ')'
%token <empty> NONE INTEGER NUMERIC REAL TEXT CAST AS

%left <bytes> OR
%left <bytes> AND
%right <bytes> NOT '!'
%left <bytes> '=' NE LIKE
%left <bytes> '<' '>' LE GE IS ISNULL NOTNULL

%type <expr> expr value_expr
%type <expr> function_call_keyword

%type <cmpOp> cmp_op
%type <column> column_name
%type <expr> value_literal
%type <convertType> convert_type

%%
start: 
  expr { yylex.(*Lexer).ast = &astRoot{$1} } ;


value_literal:
  STRING
  {
    $$ = &SQLValue{Token : Token{token: STRING, literal : $1}, Type: StrValue, Val: $1[1:len($1)-1]}
  }
|  INTEGRAL
  {
    $$ = &SQLValue{Token : Token{token: INTEGRAL, literal : $1}, Type: IntValue, Val: $1}
  }
|  FLOAT
  {
    $$ = &SQLValue{Token : Token{token: FLOAT, literal : $1}, Type: FloatValue, Val: $1}
  }
| BLOB
  {
    $$ = &SQLValue{Token : Token{token: BLOB, literal : $1}, Type: BlobValue, Val: $1}
  }
|  HEXNUM
  {
    $$ = &SQLValue{Token : Token{token: HEXNUM, literal : $1}, Type: HexNumValue, Val: $1}
  }
|  TRUE
  {
    $$ = BoolValue{token : Token{token : TRUE, literal : $1}, val : true}
  }
| FALSE
  {
    $$ = BoolValue{token : Token{token : FALSE, literal : $1}, val : false}
  }
| NULL
  {
    $$ = &NullValue{token : Token{token: NULL, literal : $1}}
  }


column_name:
  IDENTIFIER 
  { 
    if len($1) > MaxColumnNameLength {
      yylex.Error(__yyfmt__.Sprintf("column length greater than %d", MaxColumnNameLength))
      return 1
    }
    $$ = &Column{string($1)} 
  };

value_expr:
  value_literal { $$ = $1 }
| column_name { $$ = $1 }
| '!' value_expr 
  {
    $$ = &NotExpr{token: Token{token : NOT, literal : $1}, expr : $2}
  }

cmp_op:
  '='
  {
    $$ = &EqualOperator{token : Token{token: int('='), literal : $1}} 
  }
| '<'
  {
    $$ = &LessThanOperator{token : Token{token: int('<'), literal : $1}} 
  }
| '>'
  {
    $$ = &GreaterThanOperator{token : Token{token: int('>'), literal : $1}} 
  }
| LE
  {
    $$ = &LessEqualOperator{token : Token{token: LE, literal : $1}} 
  }
| GE
  {
    $$ = &GreaterEqualOperator{token : Token{token: GE, literal : $1}} 
  }
| NE
  {
    $$ = &NotEqualOperator{token : Token{token: NE, literal : $1}} 
  }

convert_type:
  NONE { $$ = NoneStr}
| TEXT { $$ = TextStr}
| REAL { $$ = RealStr}
| INTEGER { $$ = IntegerStr}
| NUMERIC { $$ = NumericStr}

function_call_keyword:
  CAST '(' expr AS convert_type ')'
  {
    $$ = &ConvertExpr{expr: $3, typ: $5}
  }

expr:
  value_expr { $$ = $1 }
| value_expr cmp_op value_expr 
  {  
    $$ = &CmpExpr{left: $1, op: $2, right: $3} 
  }
| expr AND expr 
  {  
    $$ = &AndExpr{token: Token{token : AND, literal : $2}, left: $1, right: $3}
  }
| expr OR expr 
  {  
    $$ = &OrExpr{token: Token{token : OR, literal : $2}, left: $1, right: $3}
  }
| NOT expr 
  {  
    $$ = &NotExpr{token: Token{token : NOT, literal : $1}, expr : $2}
  }
| expr IS expr
  {  
    $$ = &IsExpr{typ: IsStr, lhs : $1, rhs : $3}
  }
| expr ISNULL
  {  
    $$ = &IsExpr{typ: IsNullStr, lhs : $1, rhs : &NullValue{token : Token{token: ISNULL, literal : $2}}}
  }
| expr NOTNULL
  {  
    $$ = &IsExpr{typ: NotNullStr, lhs : $1, rhs : &NullValue{token : Token{token: NOTNULL, literal : $2}}}
  }
| function_call_keyword
;
%%