package generate

import (
	"errors"
	"strings"
)

// TokenParser processes a stream of tokens to extract structured information
// from Objective-C declarations
type TokenParser struct {
	stream *TokenStream
}

// NewTokenParser creates a new token parser for the given token stream
func NewTokenParser(stream *TokenStream) *TokenParser {
	return &TokenParser{stream: stream}
}

// Parse parses the token stream and returns structured information about the declaration
func (p *TokenParser) Parse() (*ParsedDeclaration, error) {
	p.stream.Reset()

	// Determine the declaration type based on the first tokens
	token := p.stream.Current()
	if token == nil {
		return nil, errors.New("empty token stream")
	}

	// Skip attributes and annotations at the beginning
	for token != nil && (token.Kind == TokenPunctuation && (token.Text == "__" || token.Text == "@")) {
		p.skipUntilAfter(TokenPunctuation, ")")
		token = p.stream.Current()
	}

	if token == nil {
		return nil, errors.New("unexpected end of token stream")
	}

	// Detect declaration type
	switch {
	case token.Kind == TokenKeyword && (token.Text == "typedef" || token.Text == "struct" || token.Text == "enum"):
		return p.parseTypeDefinition()
	case token.Kind == TokenKeyword && token.Text == "@interface":
		return p.parseClassOrProtocol()
	case token.Kind == TokenKeyword && token.Text == "@protocol":
		return p.parseClassOrProtocol()
	case token.Kind == TokenType || token.Kind == TokenIdentifier:
		return p.parseFunctionOrVariable()
	default:
		// Try to make a best effort parse
		return p.parseGenericDeclaration()
	}
}

// ParsedDeclaration represents a structured Objective-C declaration
type ParsedDeclaration struct {
	Kind         string            // Class, Protocol, Method, Property, Function, Constant, Enum, Struct, etc.
	Name         string            // Name of the declared entity
	ReturnType   string            // Return type for methods and functions
	Parameters   []ParsedParameter // Parameters for methods and functions
	SuperClass   string            // Superclass for class declarations
	Protocols    []string          // Adopted protocols
	Properties   map[string]string // Properties with their types
	IsDeprecated bool              // Whether the declaration is marked as deprecated
	IsStatic     bool              // Whether the declaration is static (class method, etc.)
	IsNullable   bool              // Whether the return type is nullable
}

// ParsedParameter represents a function or method parameter
type ParsedParameter struct {
	Name     string
	Type     string
	IsConst  bool
	Nullable bool
}

// parseTypeDefinition parses typedef, struct, or enum declarations
func (p *TokenParser) parseTypeDefinition() (*ParsedDeclaration, error) {
	result := &ParsedDeclaration{}
	
	// Get the kind of type definition
	keyword := p.stream.Next()
	result.Kind = strings.Title(keyword.Text)
	
	// Handle typedef struct, typedef enum, etc.
	if keyword.Text == "typedef" {
		next := p.stream.Next()
		if next == nil {
			return nil, errors.New("unexpected end of token stream after typedef")
		}
		
		if next.Kind == TokenKeyword && (next.Text == "struct" || next.Text == "enum") {
			result.Kind = strings.Title(next.Text)
		} else {
			// It's a simple typedef, the last identifier is the new type name
			p.stream.Reset()
			p.stream.Next() // Skip typedef
			
			var typeName string
			for p.stream.HasMore() {
				token := p.stream.Next()
				if token.Kind == TokenIdentifier && !p.stream.HasMore() {
					typeName = token.Text
					break
				}
			}
			
			result.Kind = "Type"
			result.Name = typeName
			return result, nil
		}
	}
	
	// Get the name of the struct/enum
	for p.stream.HasMore() {
		token := p.stream.Next()
		if token.Kind == TokenIdentifier && (p.stream.Peek() == nil || 
		   (p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "{")) {
			result.Name = token.Text
			break
		}
	}
	
	return result, nil
}

// parseClassOrProtocol parses @interface or @protocol declarations
func (p *TokenParser) parseClassOrProtocol() (*ParsedDeclaration, error) {
	result := &ParsedDeclaration{}
	
	// Determine if it's a class or protocol
	directive := p.stream.Next()
	if directive.Text == "@interface" {
		result.Kind = "Class"
	} else {
		result.Kind = "Protocol"
	}
	
	// Get the name
	if !p.stream.HasMore() {
		return nil, errors.New("unexpected end of token stream")
	}
	name := p.stream.Next()
	result.Name = name.Text
	
	// Check for superclass or adopted protocols
	if p.stream.HasMore() && p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == ":" {
		p.stream.Next() // Skip the colon
		if !p.stream.HasMore() {
			return nil, errors.New("unexpected end of token stream after superclass indicator")
		}
		
		superclass := p.stream.Next()
		result.SuperClass = superclass.Text
	}
	
	// Check for protocols
	if p.stream.HasMore() && p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "<" {
		p.stream.Next() // Skip <
		
		for p.stream.HasMore() {
			token := p.stream.Next()
			if token.Kind == TokenPunctuation && token.Text == ">" {
				break
			}
			
			if token.Kind == TokenIdentifier {
				result.Protocols = append(result.Protocols, token.Text)
			}
			
			// Skip commas
			if p.stream.HasMore() && p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "," {
				p.stream.Next()
			}
		}
	}
	
	return result, nil
}

// parseFunctionOrVariable parses function or variable/constant declarations
func (p *TokenParser) parseFunctionOrVariable() (*ParsedDeclaration, error) {
	result := &ParsedDeclaration{}
	
	// Collect return type tokens
	var returnType strings.Builder
	
	// Check for static keyword
	if p.stream.Current().Kind == TokenKeyword && p.stream.Current().Text == "static" {
		result.IsStatic = true
		p.stream.Next()
	}
	
	// Collect return type
	for p.stream.HasMore() {
		token := p.stream.Current()
		if token.Kind == TokenIdentifier && p.stream.Peek() != nil && 
		   p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "(" {
			break
		}
		
		// Handle nullability
		if token.Kind == TokenKeyword && (token.Text == "_Nullable" || token.Text == "nullable") {
			result.IsNullable = true
		}
		
		returnType.WriteString(token.Text + " ")
		p.stream.Next()
	}
	
	result.ReturnType = strings.TrimSpace(returnType.String())
	
	// Get function/variable name
	if !p.stream.HasMore() {
		result.Kind = "Constant"
		return result, nil
	}
	
	functionName := p.stream.Next()
	result.Name = functionName.Text
	
	// Check if it's a function by looking for opening parenthesis
	if p.stream.HasMore() && p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "(" {
		result.Kind = "Function"
		p.stream.Next() // Skip (
		
		// Parse parameters
		for p.stream.HasMore() {
			if p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == ")" {
				p.stream.Next() // Skip )
				break
			}
			
			param, err := p.parseParameter()
			if err != nil {
				return nil, err
			}
			
			result.Parameters = append(result.Parameters, *param)
			
			// Skip comma
			if p.stream.HasMore() && p.stream.Peek().Kind == TokenPunctuation && p.stream.Peek().Text == "," {
				p.stream.Next()
			}
		}
	} else {
		result.Kind = "Constant"
	}
	
	return result, nil
}

// parseParameter parses a single function parameter
func (p *TokenParser) parseParameter() (*ParsedParameter, error) {
	param := &ParsedParameter{}
	
	// Collect parameter type
	var paramType strings.Builder
	var paramName string
	
	// Check for const keyword
	if p.stream.HasMore() && p.stream.Peek().Kind == TokenKeyword && p.stream.Peek().Text == "const" {
		param.IsConst = true
		p.stream.Next()
		paramType.WriteString("const ")
	}
	
	// Collect type and name
	for p.stream.HasMore() {
		token := p.stream.Next()
		
		// Check for nullability
		if token.Kind == TokenKeyword && (token.Text == "_Nullable" || token.Text == "nullable") {
			param.Nullable = true
			paramType.WriteString(token.Text + " ")
			continue
		}
		
		// Exit conditions
		if token.Kind == TokenPunctuation && (token.Text == "," || token.Text == ")") {
			p.stream.Next() // Skip , or )
			break
		}
		
		// If we hit an identifier followed by comma or closing paren, it's the parameter name
		if token.Kind == TokenIdentifier && p.stream.HasMore() && 
		   p.stream.Peek().Kind == TokenPunctuation && (p.stream.Peek().Text == "," || p.stream.Peek().Text == ")") {
			paramName = token.Text
			break
		}
		
		paramType.WriteString(token.Text + " ")
	}
	
	param.Type = strings.TrimSpace(paramType.String())
	param.Name = paramName
	
	return param, nil
}

// parseGenericDeclaration makes a best effort to parse any declaration
func (p *TokenParser) parseGenericDeclaration() (*ParsedDeclaration, error) {
	result := &ParsedDeclaration{}
	result.Kind = "Unknown"
	
	// Extract any identifier that could be the name
	p.stream.Reset()
	for p.stream.HasMore() {
		token := p.stream.Next()
		if token.Kind == TokenIdentifier {
			result.Name = token.Text
			break
		}
	}
	
	return result, nil
}

// skipUntilAfter skips tokens until after the specified token
func (p *TokenParser) skipUntilAfter(kind, text string) {
	for p.stream.HasMore() {
		token := p.stream.Next()
		if token.Kind == kind && token.Text == text {
			return
		}
	}
}

// ParseTokens parses an array of token structs and returns a parsed declaration
func ParseTokens(tokens []Token) (*ParsedDeclaration, error) {
	stream := NewTokenStream(tokens)
	parser := NewTokenParser(stream)
	return parser.Parse()
}