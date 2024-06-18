package declparse

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/declparse/lexer"
)

func parseMethod(p *Parser) (next stateFn, node Node, err error) {
	decl := &MethodDecl{}

	tok, pos, lit := p.tb.Scan()
	switch tok {
	case lexer.PLUS:
		decl.TypeMethod = true
	case lexer.MINUS:
		decl.TypeMethod = false
	default:
		return nil, nil, fmt.Errorf("found %q, expected + or - at %v", lit, pos)
	}

	if tok, _, _ := p.tb.Scan(); tok == lexer.LPAREN {
		p.tb.Unscan()
		typ, err := p.expectType(true)
		if err != nil {
			return nil, nil, err
		}
		decl.ReturnType = *typ
	} else {
		p.tb.Unscan()
	}

	name, err := p.expectIdent()
	if err != nil {
		return nil, nil, err
	}
	decl.NameParts = append(decl.NameParts, name)

	if tok, _, _ = p.tb.Scan(); tok == lexer.COLON {
		for {
			arg := ArgInfo{}

			typ, err := p.expectType(true)
			if err != nil {
				return nil, nil, err
			}
			arg.Type = *typ

			if arg.Name, err = p.expectIdent(); err != nil {
				return nil, nil, err
			}

			decl.Args = append(decl.Args, arg)

			if tok, _, _ = p.tb.Scan(); tok == lexer.COMMA {
				if tok, _, _ = p.tb.Scan(); tok == lexer.VARARG {
					decl.Variadic = true
				} else {
					p.tb.Unscan()
				}
			} else {
				p.tb.Unscan()
			}

			if tok, pos, lit = p.tb.Scan(); tok == lexer.IDENT {
				decl.NameParts = append(decl.NameParts, lit)

				if err := p.expectToken(lexer.COLON); err != nil {
					return nil, nil, err
				}
			} else {
				p.tb.Unscan()
				break
			}
		}
	} else {
		p.tb.Unscan()
	}

	tok, _, lit = p.tb.Scan()
	if lit == "NS_UNAVAILABLE" {
		// Handle NS_UNAVAILABLE.
		decl.Unavailable = true
	} else {
		p.tb.Unscan()
	}

	return nil, decl, nil
}
