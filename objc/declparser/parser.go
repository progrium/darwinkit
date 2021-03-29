package declparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Statement struct {
	Method    *MethodDecl
	Property  *PropertyDecl
	Interface *InterfaceDecl
}

type InterfaceDecl struct {
	Name      string
	SuperName string
}

type PropertyDecl struct {
	TypeProperty bool // class property
	Readonly     bool
	Weak         bool
	Nonatomic    bool
	Copy         bool
	Getter       string
	Setter       string
	Name         string
	Type         TypeInfo
}

type MethodDecl struct {
	TypeMethod bool // instance method otherwise
	ReturnType TypeInfo
	NameParts  []string
	Args       []ArgInfo
}

func (m *MethodDecl) Name() string {
	if len(m.NameParts) == 0 {
		return ""
	}
	if len(m.NameParts) == 1 {
		return m.NameParts[0]
	}
	return strings.Join(append(m.NameParts, ""), ":")
}

type TypeInfo struct {
	Name  string
	IsPtr bool
}

type ArgInfo struct {
	Name string
	Type TypeInfo
}

type Parser struct {
	s   *scanner
	buf struct {
		tok token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: &scanner{r: bufio.NewReader(r)}}
}

func NewStringParser(s string) *Parser {
	return &Parser{s: &scanner{r: bufio.NewReader(bytes.NewBufferString(s))}}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok token, lit string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit = p.s.Scan()

	// ignore whitespace
	if tok == WS {
		tok, lit = p.s.Scan()
	}

	p.buf.tok, p.buf.lit = tok, lit

	return
}

func (p *Parser) unscan() { p.buf.n = 1 }

// scanType will scan a typeInfo
func (p *Parser) scanType() (*TypeInfo, error) {
	ti := &TypeInfo{}

	if tok, lit := p.scan(); tok != LEFTPAREN {
		return nil, fmt.Errorf("found %q, expected (", lit)
	}

	tok, lit := p.scan()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected identifier", lit)
	}
	ti.Name = lit

	tok, lit = p.scan()
	if tok == ASTERISK {
		ti.IsPtr = true
	} else {
		p.unscan()
	}

	if tok, lit := p.scan(); tok != RIGHTPAREN {
		return nil, fmt.Errorf("found %q, expected )", lit)
	}

	return ti, nil
}

func (p *Parser) Parse() (*Statement, error) {
	tok, lit := p.scan()
	switch tok {
	case PLUS, MINUS:
		p.unscan()
		decl, err := p.parseMethod()
		return &Statement{Method: decl}, err
	case PROPERTY:
		decl, err := p.parseProperty()
		return &Statement{Property: decl}, err
	case INTERFACE:
		decl, err := p.parseInterface()
		return &Statement{Interface: decl}, err
	default:
		return nil, fmt.Errorf("found %q, expected method (+,-) or keyword (@...)", lit)
	}
}

func (p *Parser) parseProperty() (*PropertyDecl, error) {
	decl := &PropertyDecl{}

	tok, lit := p.scan()
	if tok == LEFTPAREN {
		for {
			tok, lit = p.scan()
			if tok != IDENT {
				return nil, fmt.Errorf("found %q, expected property attribute", lit)
			}

			switch lit {
			case "readwrite":
				// readwrite is opposite of readonly,
				// this is the default, so no-op
			case "readonly":
				decl.Readonly = true
			case "class":
				decl.TypeProperty = true
			case "strong":
				// strong is opposite of weak,
				// this is the default, so no-op
			case "weak":
				decl.Weak = true
			case "copy":
				decl.Copy = true
			case "assign":
				// another default, no-ops
			case "nonatomic":
				decl.Nonatomic = true
			case "setter":
				tok, lit = p.scan()
				if tok != EQUAL {
					return nil, fmt.Errorf("found %q, expected =", lit)
				}

				tok, lit = p.scan()
				if tok != IDENT {
					return nil, fmt.Errorf("found %q, expected setter identifier", lit)
				}
				decl.Setter = lit
			case "getter":
				tok, lit = p.scan()
				if tok != EQUAL {
					return nil, fmt.Errorf("found %q, expected =", lit)
				}

				tok, lit = p.scan()
				if tok != IDENT {
					return nil, fmt.Errorf("found %q, expected getter identifier", lit)
				}
				decl.Getter = lit
			default:
				return nil, fmt.Errorf("found %q, unrecognized property attribute", lit)
			}

			tok, lit = p.scan()
			if tok == RIGHTPAREN {
				break
			}
			if tok != COMMA {
				return nil, fmt.Errorf("found %q, expected , or )", lit)
			}
		}
	} else {
		p.unscan()
	}

	tok, lit = p.scan()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected type identifier", lit)
	}
	decl.Type.Name = lit

	tok, lit = p.scan()
	if tok == ASTERISK {
		decl.Type.IsPtr = true
	} else {
		p.unscan()
	}

	tok, lit = p.scan()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected name identifier", lit)
	}
	decl.Name = lit

	tok, lit = p.scan()
	if tok != SEMICOLON {
		return nil, fmt.Errorf("found %q, expected ;", lit)
	}

	return decl, nil
}

func (p *Parser) parseInterface() (*InterfaceDecl, error) {
	decl := &InterfaceDecl{}

	tok, lit := p.scan()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected identifier", lit)
	}
	decl.Name = lit

	tok, lit = p.scan()
	if tok == COLON {
		tok, lit = p.scan()
		if tok != IDENT {
			return nil, fmt.Errorf("found %q, expected identifier", lit)
		}
		decl.SuperName = lit
	} else {
		p.unscan()
	}

	return decl, nil
}

func (p *Parser) parseMethod() (*MethodDecl, error) {
	decl := &MethodDecl{}

	tok, lit := p.scan()
	switch tok {
	case PLUS:
		decl.TypeMethod = true
	case MINUS:
		decl.TypeMethod = false
	default:
		return nil, fmt.Errorf("found %q, expected + or -", lit)
	}

	typ, err := p.scanType()
	if err != nil {
		return nil, err
	}
	decl.ReturnType = *typ

	tok, lit = p.scan()
	if tok != IDENT {
		return nil, fmt.Errorf("found %q, expected identifier", lit)
	}
	decl.NameParts = append(decl.NameParts, lit)

	tok, lit = p.scan()
	if tok == SEMICOLON {
		return decl, nil
	} else if tok == COLON {
		for {
			arg := ArgInfo{}

			typ, err := p.scanType()
			if err != nil {
				return nil, err
			}
			arg.Type = *typ

			tok, lit = p.scan()
			if tok != IDENT {
				return nil, fmt.Errorf("found %q, expected identifier", lit)
			}
			arg.Name = lit

			decl.Args = append(decl.Args, arg)

			tok, lit = p.scan()
			if tok == SEMICOLON {
				return decl, nil
			} else if tok == IDENT {
				decl.NameParts = append(decl.NameParts, lit)

				tok, lit = p.scan()
				if tok != COLON {
					return nil, fmt.Errorf("found %q, expected :", lit)
				}
			} else {
				return nil, fmt.Errorf("found %q, expected ; or more arguments", lit)
			}
		}
	} else {
		return nil, fmt.Errorf("found %q, expected : or ;", lit)
	}
}
