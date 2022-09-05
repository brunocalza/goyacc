%{
package expression

import "strconv"

%}

%union{
  expression float32
  number float32
  str string
}

%token <empty> '(' ')'
%left <empty> '-' '+' 
%left <empty> '*' '/' 
%token <str> INTEGER FLOAT

%type <expression> expression
%type <number> number

%%
start:
  expression
  {
    yylex.(*Lexer).result = $1
  }

expression:
  '(' expression ')'
  {
    $$ = $2
  }
| expression '+' expression
  {
    $$ = $1 + $3
  }
| expression '-' expression
  {
    $$ = $1 - $3
  }
| expression '*' expression
  {
    $$ = $1 * $3
  }
| expression '/' expression
  {
    $$ = $1 / $3
  }
| number
  {
    $$ = $1
  }
| '-' expression
  {
    $$ = $2 * -1
  }
;

number:
  INTEGER
  {
    n, _ := strconv.ParseInt($1, 10, 32)
    $$ = float32(n)
  }
| FLOAT
  {
    n, _ := strconv.ParseFloat($1, 32)
    $$ = float32(n)
  }
;
 
%%