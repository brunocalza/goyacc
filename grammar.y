%{
package main
%}

%union{
  string string
  expr expr 
  logicOp LogicalOperator
  cmpOp CmpOperator
}

%token <string> NUMBER IDENTIFIER
%token <logicOp> LOP
%token <cmpOp> CMPOP

%type <expr> expr

%left  LOP
%left  CMPOP


%%
start: expr { yylex.(*Lexer).ast = &astRoot{$1} };

expr:
      NUMBER  { $$ = &number{$1} }
    | IDENTIFIER { $$ = &column{$1}}
    | expr CMPOP expr {  $$ = &CmpExpr{left: $1, op: $2, right: $3} }
    | expr LOP expr {  $$ = &LogicalExpr{left: $1, op: $2, right: $3} }
    ;
%%