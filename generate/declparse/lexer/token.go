package lexer

import "strings"

// Token represents each lexer symbol
type Token int

// Token enums
const (
	ILLEGAL Token = iota
	EOF
	WS

	// Punctuation

	LPAREN      // (
	RPAREN      // )
	LBRACKET    // [
	RBRACKET    // ]
	LCURLY      // {
	RCURLY      // }
	COMMA       // ,
	SEMICOLON   // ;
	COLON       // :
	DOT         // .
	SINGLEQUOTE // '
	DOUBLEQUOTE // "
	PERCENT     // %
	DOLLAR      // $
	HASH        // #
	ATSIGN      // @
	ARROW       // ->
	EQARROW     // =>

	// Literals

	startLiterals
	IDENT
	INTEGER
	DECIMAL
	STRING
	BADSTRING
	BADESCAPE
	TRUE
	FALSE
	REGEX
	BADREGEX
	DURATION
	endLiterals

	// Operators

	startOperators
	PLUS       // +
	PLUSPLUS   // ++
	MINUS      // -
	MINUSMINUS // --
	MUL        // *
	DIV        // /
	AMPERSAND  // &
	XOR        // ^
	PIPE       // |
	LSHIFT     // <<
	RSHIFT     // >>
	POW        // **

	AND // AND
	OR  // OR

	EQ       // =
	NEQ      // !=
	EQEQ     // ==
	EQREGEX  // =~
	NEQREGEX // !~
	LT       // <
	LTE      // <=
	GT       // >
	GTE      // >=

	VARARG // ...
	endOperators
)

var tokens = map[Token]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	LPAREN:      "(",
	RPAREN:      ")",
	LBRACKET:    "[",
	RBRACKET:    "]",
	LCURLY:      "{",
	RCURLY:      "}",
	COMMA:       ",",
	SEMICOLON:   ";",
	COLON:       ":",
	DOT:         ".",
	SINGLEQUOTE: "'",
	DOUBLEQUOTE: "\"",
	PERCENT:     "%%",
	DOLLAR:      "$",
	HASH:        "#",
	ATSIGN:      "@",
	VARARG:      "...",

	IDENT:     "IDENT",
	INTEGER:   "INTEGER",
	DECIMAL:   "DECIMAL",
	STRING:    "TEXTUAL",
	DURATION:  "DURATION",
	BADSTRING: "BADSTRING",
	BADESCAPE: "BADESCAPE",
	REGEX:     "REGEX",
	BADREGEX:  "BADREGEX",

	PLUS:       "+",
	PLUSPLUS:   "++",
	MINUS:      "-",
	MINUSMINUS: "--",
	MUL:        "*",
	DIV:        "/",
	AMPERSAND:  "&",
	XOR:        "^",
	PIPE:       "|",
	RSHIFT:     ">>",
	LSHIFT:     "<<",
	POW:        "**",
	ARROW:      "->",
	EQARROW:    "=>",

	AND: "AND",
	OR:  "OR",

	TRUE:  "true",
	FALSE: "false",

	EQ:       "=",
	NEQ:      "!=",
	EQEQ:     "==",
	EQREGEX:  "=~",
	NEQREGEX: "!~",

	LT:  "<",
	LTE: "<=",
	GT:  ">",
	GTE: ">=",
}

var (
	keywords = map[string]Token{}
	prefixed = map[rune][]string{}
)

func init() {
	keywords["and"] = AND
	keywords["or"] = OR
	keywords["true"] = TRUE
	keywords["false"] = FALSE
}

// LoadTokenMap allows for extra keywords to be added to the lexer
func LoadTokenMap(keywordTokens map[Token]string) {

	// Combine built-in tokens and keywords
	for k, v := range keywordTokens {
		tokens[k] = v
	}

	// Load Keywords
	for k, v := range keywordTokens {
		keywords[strings.ToLower(v)] = k
		for _, tv := range tokens {
			if len(tv) == 1 && v[0] == tv[0] {
				r := rune(v[0])
				prefixed[r] = append(prefixed[r], v)
			}
		}
	}
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if _, ok := tokens[tok]; ok {
		return tokens[tok]
	}
	return ""
}

// Precedence returns the operator precedence of the binary operator token.
func (tok Token) Precedence() int {
	switch tok {
	case OR:
		return 1
	case AND:
		return 2
	case EQ, NEQ, EQREGEX, NEQREGEX, LT, LTE, GT, GTE:
		return 3
	case PLUS, MINUS:
		return 4
	case MUL, DIV:
		return 5
	case PIPE, XOR, RSHIFT, LSHIFT, POW, PLUSPLUS, MINUSMINUS:
		return 6
	}
	return 0
}

// isOperator returns true for operator tokens.
func (tok Token) IsOperator() bool { return tok > startOperators && tok < endOperators }

// tokstr returns a literal if provided, otherwise returns the token string.
func tokstr(tok Token, lit string) string {
	if lit != "" {
		return lit
	}
	return tok.String()
}

// Lookup returns the token associated with a given string.
func Lookup(ident string) Token {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return IDENT
}
