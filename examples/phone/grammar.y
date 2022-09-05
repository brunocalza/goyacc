%{
package phone

type PhoneNumber struct {
  AreaCode AreaCode
  Exchange Exchange
  Subscriber Subscriber
}

type AreaCode string
type Exchange string
type Subscriber string

%}

%union{
  areaCode AreaCode
  exchange Exchange
  subscriber Subscriber
  phoneNumber PhoneNumber
  byt byte
  string string
}

%token <empty> '(' '-' ')' '.' 
%token <byt> '0' '1' '2' '3' '4' '5' '6' '7' '8' '9'

%type <string> digit_2_9 digit
%type <areaCode> area_code
%type <exchange> exchange
%type <subscriber> subscriber
%type <phoneNumber> phone_number


// n##-n##-####
// (n##) n##-####
// n## n## ####
// n##.n##.####

// n=digits 2–9, #=digits 0–9

%%
start:
  phone_number
  {
    yylex.(*Lexer).phoneNumber = $1
  }

phone_number:
  area_code '-' exchange '-' subscriber
  {
    $$ = PhoneNumber{AreaCode: $1, Exchange: $3, Subscriber: $5}
  }
| '(' area_code ')' ' ' exchange '-' subscriber
  {
    $$ = PhoneNumber{AreaCode: $2, Exchange: $5, Subscriber: $7}
  }
| area_code ' ' exchange ' ' subscriber
  {
    $$ = PhoneNumber{AreaCode: $1, Exchange: $3, Subscriber: $5}
  }
| area_code '.' exchange '.' subscriber
  {
    $$ = PhoneNumber{AreaCode: $1, Exchange: $3, Subscriber: $5}
  }

area_code:
  digit_2_9 digit digit
  {
    $$ = AreaCode($1 + $2 + $3)
  }

exchange:
  digit_2_9 digit digit
  {
    $$ = Exchange($1 + $2 + $3)
  }

subscriber:
  digit digit digit digit
  {
    $$ = Subscriber($1 + $2 + $3 + $4)
  }

digit_2_9: 
  '2'
  {
    $$ = string($1)
  }
| '3'
  {
    $$ = string($1)
  }
| '4'
  {
    $$ = string($1)
  }
| '5'
  {
    $$ = string($1)
  }
| '6'
  {
    $$ = string($1)
  }
| '7'
  {
    $$ = string($1)
  }
| '8'
  {
    $$ = string($1)
  }
| '9'
  {
    $$ = string($1)
  }

digit:
  '0'
  {
    $$ = string($1)
  }
| '1'
  {
    $$ = string($1)
  }
| digit_2_9
  {
    $$ = string($1)
  }
%%