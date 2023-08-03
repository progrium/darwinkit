package declparse

import (
	"github.com/progrium/macdriver/generate/declparse/keywords"
	"github.com/progrium/macdriver/generate/declparse/lexer"
)

func parseInterface(p *Parser) (next stateFn, node Node, err error) {
	decl := &InterfaceDecl{}

	if err := p.expectToken(keywords.INTERFACE); err != nil {
		return nil, nil, err
	}

	decl.Name, err = p.expectIdent()
	if err != nil {
		return nil, nil, err
	}

	if tok, _, _ := p.tb.Scan(); tok == lexer.COLON {
		if decl.SuperName, err = p.expectIdent(); err != nil {
			return nil, nil, err
		}
	} else {
		p.tb.Unscan()
	}

	return nil, decl, nil
}
