package main

type parser struct {
	l *Lexer
}

func NewParser() *parser {
	return &parser{
		l: &Lexer{},
	}
}

func (p *parser) Parse(sql string) (*astRoot, error) {
	yyErrorVerbose = true

	p.l.input = []byte(sql)
	p.l.readByte()

	yyParse(p.l)

	if p.l.err != nil {
		return nil, p.l.err
	}

	return p.l.ast, nil
}

func Deparse(ast *astRoot) string {
	if ast == nil {
		return ""
	}
	return ast.expr.ToString()
}
