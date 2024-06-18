package keywords

import (
	"github.com/progrium/darwinkit/generate/declparse/lexer"
)

func init() {
	// Loads keyword tokens into lexer
	lexer.LoadTokenMap(tokenMap)
}

const (
	// Starts the keywords with an offset from the built in tokens
	startKeywords lexer.Token = iota + 1000

	PROPERTY
	INTERFACE
	PROTOCOL

	ENUM
	CONST
	TYPEDEF
	STRUCT

	endKeywords
)

var tokenMap = map[lexer.Token]string{
	PROPERTY:  "@property",
	INTERFACE: "@interface",
	PROTOCOL:  "@protocol",
	ENUM:      "enum",
	CONST:     "const",
	TYPEDEF:   "typedef",
	STRUCT:    "struct",
}

// IsKeyword returns true if the token is a keyword.
func IsKeyword(tok lexer.Token) bool {
	return tok > startKeywords && tok < endKeywords
}
