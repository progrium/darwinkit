package declparse

import (
	"strings"

	"github.com/progrium/macdriver/generate/declparse/lexer"
)

func parseVariable(p *Parser) (next stateFn, node Node, err error) {
	decl := &VariableDecl{}

	typ, err := p.expectType(false)
	if err != nil {
		return nil, nil, err
	}
	decl.Type = *typ

	decl.Name, err = p.expectIdent()
	if err != nil {
		return nil, nil, err
	}

	if err := p.expectToken(lexer.EQ); err != nil {
		p.tb.Unscan()
		return nil, decl, nil
	}

	var rest []string
	for {
		tok, _, lit := p.tb.Scan()
		if tok == lexer.EOF || tok == lexer.SEMICOLON {
			break
		}
		if lit == "" {
			lit = tok.String()
		}
		rest = append(rest, lit)
	}
	decl.Value = strings.Join(rest, "")

	return nil, decl, nil
}

func parseEnumCase(p *Parser) (next stateFn, node Node, err error) {
	decl := &VariableDecl{}

	decl.Name, err = p.expectIdent()
	if err != nil {
		return nil, nil, err
	}

	if err := p.expectToken(lexer.EQ); err != nil {
		p.tb.Unscan()
		return nil, decl, nil
	}

	var rest []string
	for {
		tok, _, lit := p.tb.Scan()
		if tok == lexer.EOF || tok == lexer.SEMICOLON {
			break
		}
		if lit == "" {
			lit = tok.String()
		}
		rest = append(rest, lit)
	}
	decl.Value = strings.Join(rest, "")

	return nil, decl, nil
}
