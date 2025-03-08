package generate

// Token represents a single token in an Objective-C declaration
type Token struct {
	Kind string // identifier, keyword, operator, punctuation, etc.
	Text string // actual text of the token
}

// TokenKind constants for different types of tokens
const (
	TokenIdentifier  = "identifier"
	TokenKeyword     = "keyword"
	TokenOperator    = "operator"
	TokenPunctuation = "punctuation"
	TokenType        = "type"
	TokenString      = "string"
	TokenNumber      = "number"
)

// TokenStream represents a sequence of tokens
type TokenStream struct {
	tokens []Token
	pos    int
}

// NewTokenStream creates a new token stream from a slice of tokens
func NewTokenStream(tokens []Token) *TokenStream {
	return &TokenStream{tokens: tokens}
}

// Current returns the current token without advancing
func (s *TokenStream) Current() *Token {
	if s.pos >= len(s.tokens) {
		return nil
	}
	return &s.tokens[s.pos]
}

// Next advances to and returns the next token, or nil if at end
func (s *TokenStream) Next() *Token {
	if s.pos >= len(s.tokens) {
		return nil
	}
	tok := &s.tokens[s.pos]
	s.pos++
	return tok
}

// Peek returns the next token without advancing
func (s *TokenStream) Peek() *Token {
	if s.pos >= len(s.tokens) {
		return nil
	}
	return &s.tokens[s.pos]
}

// HasMore returns true if there are more tokens to process
func (s *TokenStream) HasMore() bool {
	return s.pos < len(s.tokens)
}

// Reset resets the stream to the beginning
func (s *TokenStream) Reset() {
	s.pos = 0
}
