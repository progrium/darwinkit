package lexer

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
		pos Pos
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: EOF},
		{s: `#`, tok: HASH, lit: ``},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},
		{s: "\r", tok: WS, lit: "\n"},
		{s: "\r\n", tok: WS, lit: "\n"},
		{s: "\rX", tok: WS, lit: "\n"},
		{s: "\n\r", tok: WS, lit: "\n\n"},
		{s: " \n\t \r\n\t", tok: WS, lit: " \n\t \n\t"},
		{s: " foo", tok: WS, lit: " "},

		// Numeric operators
		{s: `+`, tok: PLUS},
		{s: `-`, tok: MINUS},
		{s: `*`, tok: MUL},
		{s: `/`, tok: DIV},

		{s: "&", tok: AMPERSAND},
		{s: "^", tok: XOR},
		{s: "|", tok: PIPE},
		{s: ">>", tok: RSHIFT},
		{s: "<<", tok: LSHIFT},
		{s: "**", tok: POW},
		{s: "->", tok: ARROW},
		{s: "=>", tok: EQARROW},

		// Logical operators
		{s: `AND`, tok: AND},
		{s: `and`, tok: AND},
		{s: `OR`, tok: OR},
		{s: `or`, tok: OR},

		{s: `=`, tok: EQ},
		{s: `<>`, tok: NEQ},
		{s: `! `, tok: ILLEGAL, lit: "!"},
		{s: `!=`, tok: NEQ},
		{s: `<`, tok: LT},
		{s: `<=`, tok: LTE},
		{s: `>`, tok: GT},
		{s: `>=`, tok: GTE},

		// Misc tokens
		{s: `(`, tok: LPAREN},
		{s: `)`, tok: RPAREN},
		{s: "[", tok: LBRACKET},
		{s: "]", tok: RBRACKET},
		{s: "{", tok: LCURLY},
		{s: "}", tok: RCURLY},
		{s: `,`, tok: COMMA},
		{s: `;`, tok: SEMICOLON},
		{s: `:`, tok: COLON},
		{s: `.`, tok: DOT},
		{s: `%`, tok: PERCENT},
		{s: "$", tok: DOLLAR},
		{s: "#", tok: HASH},
		{s: "@", tok: ATSIGN},

		{s: `=~`, tok: EQREGEX},
		{s: `!~`, tok: NEQREGEX},

		// Identifiers
		{s: `foo`, tok: IDENT, lit: `foo`},
		{s: `_foo`, tok: IDENT, lit: `_foo`},
		{s: `Zx12_3U_-`, tok: IDENT, lit: `Zx12_3U_`},
		{s: `"foo"`, tok: STRING, lit: `foo`},
		{s: `"foo\\bar"`, tok: STRING, lit: `foo\bar`},
		{s: `"foo\bar"`, tok: BADESCAPE, lit: `\b`, pos: Pos{Line: 0, Char: 5}},
		{s: `"foo\"bar\""`, tok: STRING, lit: `foo"bar"`},
		{s: `test"`, tok: BADSTRING, lit: "", pos: Pos{Line: 0, Char: 3}},
		{s: `"test`, tok: BADSTRING, lit: `test`},

		{s: `true`, tok: TRUE},
		{s: `false`, tok: FALSE},

		// Strings
		{s: `'testing 123!'`, tok: STRING, lit: `testing 123!`},
		{s: `'foo\nbar'`, tok: STRING, lit: "foo\nbar"},
		{s: `'foo\\bar'`, tok: STRING, lit: "foo\\bar"},
		{s: `'test`, tok: BADSTRING, lit: `test`},
		{s: "'test\nfoo", tok: BADSTRING, lit: `test`},
		{s: `'test\g'`, tok: BADESCAPE, lit: `\g`, pos: Pos{Line: 0, Char: 6}},

		// Numbers
		{s: `100`, tok: INTEGER, lit: `100`},
		{s: `100.23`, tok: DECIMAL, lit: `100.23`},
		{s: `+100.23`, tok: DECIMAL, lit: `+100.23`},
		{s: `-100.23`, tok: DECIMAL, lit: `-100.23`},
		{s: `-100.`, tok: DECIMAL, lit: `-100`},
		{s: `.23`, tok: DECIMAL, lit: `.23`},
		{s: `+.23`, tok: DECIMAL, lit: `+.23`},
		{s: `-.23`, tok: DECIMAL, lit: `-.23`},
		//{s: `.`, tok: ILLEGAL, lit: `.`},
		{s: `-.`, tok: MINUS, lit: ``},
		{s: `+.`, tok: PLUS, lit: ``},
		{s: `10.3s`, tok: DECIMAL, lit: `10.3`},
		{s: `10.3E5`, tok: DECIMAL, lit: `10.3E5`},
		{s: `10.3E+5`, tok: DECIMAL, lit: `10.3E+5`},
		{s: `10.3E-5`, tok: DECIMAL, lit: `10.3E-5`},
		{s: `10.3E-0.5`, tok: DECIMAL, lit: `10.3E-0.5`},
		{s: `10.3E-0.5.3`, tok: DECIMAL, lit: `10.3E-0.5`},
		{s: `10E3`, tok: DECIMAL, lit: `10E3`},

		// Durations
		{s: `10u`, tok: DURATION, lit: `10u`},
		{s: `10µ`, tok: DURATION, lit: `10µ`},
		{s: `10ms`, tok: DURATION, lit: `10ms`},
		{s: `-1s`, tok: DURATION, lit: `-1s`},
		{s: `10m`, tok: DURATION, lit: `10m`},
		{s: `10h`, tok: DURATION, lit: `10h`},
		{s: `10d`, tok: DURATION, lit: `10d`},
		{s: `10w`, tok: DURATION, lit: `10w`},
		{s: `10x`, tok: INTEGER, lit: `10`}, // non-duration unit
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, pos, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.pos.Line != pos.Line || tt.pos.Char != pos.Char {
			t.Errorf("%d. %q pos mismatch: exp=%#v got=%#v", i, tt.s, tt.pos, pos)
		} else if tt.lit != "" && tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

func errstring(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

// Ensure the library can correctly scan strings.
func TestScanString(t *testing.T) {
	var tests = []struct {
		in  string
		out string
		err string
	}{
		{in: `""`, out: ``},
		{in: `"foo bar"`, out: `foo bar`},
		{in: `'foo bar'`, out: `foo bar`},
		{in: `"foo\nbar"`, out: "foo\nbar"},
		{in: `"foo\\bar"`, out: `foo\bar`},
		{in: `"foo\"bar"`, out: `foo"bar`},
		{in: `'foo\'bar'`, out: `foo'bar`},

		{in: `"foo` + "\n", out: `foo`, err: "bad string"}, // newline in string
		{in: `"foo`, out: `foo`, err: "bad string"},        // unclosed quotes
		{in: `"foo\xbar"`, out: `\x`, err: "bad escape"},   // invalid escape
	}

	for i, tt := range tests {
		out, err := ScanString(strings.NewReader(tt.in))
		if tt.err != errstring(err) {
			t.Errorf("%d. %s: error: exp=%s, got=%s", i, tt.in, tt.err, err)
		} else if tt.out != out {
			t.Errorf("%d. %s: out: exp=%s, got=%s", i, tt.in, tt.out, out)
		}
	}
}

// Test scanning regex
func TestScanRegex(t *testing.T) {
	var tests = []struct {
		in  string
		tok Token
		lit string
		err string
	}{
		{in: `/^payments\./`, tok: REGEX, lit: `^payments\.`},
		{in: `/foo\/bar/`, tok: REGEX, lit: `foo/bar`},
		{in: `/foo\\/bar/`, tok: REGEX, lit: `foo\/bar`},
		{in: `/foo\\bar/`, tok: REGEX, lit: `foo\\bar`},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.in))
		tok, _, lit := s.ScanRegex()
		if tok != tt.tok {
			t.Errorf("%d. %s: error:\n\texp=%s\n\tgot=%s\n", i, tt.in, tt.tok.String(), tok.String())
		}
		if lit != tt.lit {
			t.Errorf("%d. %s: error:\n\texp=%s\n\tgot=%s\n", i, tt.in, tt.lit, lit)
		}
	}
}
