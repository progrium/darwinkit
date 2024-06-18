package declparse

import (
	"github.com/progrium/darwinkit/generate/declparse/keywords"
	"github.com/progrium/darwinkit/generate/declparse/lexer"
)

func parseProtocol(p *Parser) (next stateFn, node Node, err error) {
	decl := &ProtocolDecl{}

	if err := p.expectToken(keywords.PROTOCOL); err != nil {
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
