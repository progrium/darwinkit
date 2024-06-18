package declparse

import (
	"fmt"

	"github.com/progrium/darwinkit/generate/declparse/keywords"
	"github.com/progrium/darwinkit/generate/declparse/lexer"
)

func parseProperty(p *Parser) (next stateFn, node Node, err error) {
	decl := &PropertyDecl{Attrs: make(map[PropAttr]string)}

	if err := p.expectToken(keywords.PROPERTY); err != nil {
		return nil, nil, err
	}

	if tok, _, _ := p.tb.Scan(); tok == lexer.LPAREN {
		for {
			lit, err := p.expectIdent()
			if err != nil {
				return nil, nil, err
			}

			for _, attr := range PropAttrs() {
				if lit == attr.String() {
					switch attr {
					case PropAttrGetter, PropAttrSetter:
						if err := p.expectToken(lexer.EQ); err != nil {
							return nil, nil, err
						}
						val, err := p.expectIdent()
						if err != nil {
							return nil, nil, err
						}
						t, _, _ := p.tb.Scan()
						if t.String() == ":" {
							val = val + ":"
						} else {
							p.tb.Unscan()
						}
						decl.Attrs[attr] = val
					default:
						decl.Attrs[attr] = ""
					}
				}
			}

			tok, _, lit := p.tb.Scan()
			if tok == lexer.RPAREN {
				break
			}
			if tok != lexer.COMMA {
				return nil, nil, fmt.Errorf("found %q, expected , or )", lit)
			}
		}
	} else {
		p.tb.Unscan()
	}

	if tok, _, lit := p.tb.Scan(); tok == lexer.IDENT && lit == "IBOutlet" {
		decl.IsOutlet = true
	} else {
		p.tb.Unscan()
	}

	typ, err := p.expectType(false)
	if err != nil {
		return nil, nil, err
	}
	decl.Type = *typ

	if typ.Func != nil {
		decl.Name = typ.Func.Name
		typ.Func.Name = ""
		return nil, decl, nil
	}

	decl.Name, err = p.expectIdent()
	if err != nil {
		return nil, nil, err
	}

	return nil, decl, nil
}
