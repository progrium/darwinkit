package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// NOTE: The code below is HEAVILY influenced by the InfluxQL parser and lexer.
// You can find the InfluxDB code at https://github.com/influxdb/influxdb/blob/master/influxql/scanner.go

// Scanner represents a lexical scanner for InfluxQL.
type Scanner struct {
	OneRuneOperators bool // only scan one rune operators (+, not ++)

	r *reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: &reader{r: bufio.NewReader(r)}}
}

// Scan returns the next token and position from the underlying reader.
// Also returns the literal text read for strings and numbers tokens
// since these token types can have different literal representations.
func (s *Scanner) Scan() (tok Token, pos Pos, lit string) {

	// Read next code point.
	ch0, pos := s.r.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter, or certain acceptable special characters, then consume
	// as an ident or reserved word.
	if isWhitespace(ch0) {
		return s.scanWhitespace()
	} else if isLetter(ch0) || ch0 == '_' || isPrefixed(ch0) {
		s.r.unread()
		return s.scanIdent()
	} else if isDigit(ch0) {
		return s.scanNumber()
	}

	// Otherwise parse individual characters.
	switch ch0 {
	case eof:
		return EOF, pos, ""
	case '"':
		return s.scanString()
	case '\'':
		return s.scanString()
	case '.':
		ch1, _ := s.r.read()
		ch2, _ := s.r.read()
		if ch1 == '.' && ch2 == '.' {
			return VARARG, pos, "..."
		}
		s.r.unread()
		if isDigit(ch1) {
			s.r.unread()
			return s.scanNumber()
		}
		return DOT, pos, ""
	case '+':
		if ch1, _ := s.r.read(); ch1 == '+' && !s.OneRuneOperators {
			return PLUSPLUS, pos, ""
		}
		s.r.unread()
		return s.scanNumber()
	case '-':
		if ch1, _ := s.r.read(); ch1 == '>' && !s.OneRuneOperators {
			return ARROW, pos, ""
		} else if ch1 == '-' && !s.OneRuneOperators {
			return MINUSMINUS, pos, ""
		}
		s.r.unread()
		return s.scanNumber()
	case '*':
		if ch1, _ := s.r.read(); ch1 == '*' && !s.OneRuneOperators {
			return POW, pos, ""
		}
		s.r.unread()
		return MUL, pos, ""
	case '/':
		return DIV, pos, ""
	case '=':
		if ch1, _ := s.r.read(); ch1 == '~' && !s.OneRuneOperators {
			return EQREGEX, pos, ""
		} else if ch1 == '>' && !s.OneRuneOperators {
			return EQARROW, pos, ""
		} else if ch1 == '=' && !s.OneRuneOperators {
			return EQEQ, pos, ""
		}
		s.r.unread()
		return EQ, pos, ""
	case '!':
		if ch1, _ := s.r.read(); ch1 == '=' && !s.OneRuneOperators {
			return NEQ, pos, ""
		} else if ch1 == '~' && !s.OneRuneOperators {
			return NEQREGEX, pos, ""
		}
		s.r.unread()
	case '>':
		if ch1, _ := s.r.read(); ch1 == '=' && !s.OneRuneOperators {
			return GTE, pos, ""
		} else if ch1 == '>' && !s.OneRuneOperators {
			return RSHIFT, pos, ""
		}
		s.r.unread()
		return GT, pos, ""
	case '<':
		if ch1, _ := s.r.read(); ch1 == '=' && !s.OneRuneOperators {
			return LTE, pos, ""
		} else if ch1 == '>' && !s.OneRuneOperators {
			return NEQ, pos, ""
		} else if ch1 == '<' && !s.OneRuneOperators {
			return LSHIFT, pos, ""
		}
		s.r.unread()
		return LT, pos, ""
	case '(':
		return LPAREN, pos, ""
	case ')':
		return RPAREN, pos, ""
	case '[':
		return LBRACKET, pos, ""
	case ']':
		return RBRACKET, pos, ""
	case '{':
		return LCURLY, pos, ""
	case '}':
		return RCURLY, pos, ""
	case ',':
		return COMMA, pos, ""
	case ';':
		return SEMICOLON, pos, ""
	case ':':
		return COLON, pos, ""
	case '^':
		return XOR, pos, ""
	case '|':
		return PIPE, pos, ""
	case '&':
		return AMPERSAND, pos, ""
	case '%':
		return PERCENT, pos, ``
	case '$':
		return DOLLAR, pos, ""
	case '#':
		return HASH, pos, ""
	case '@':
		return ATSIGN, pos, ""
	}

	return ILLEGAL, pos, string(ch0)
}

// Peek returns the next rune without advancing the scanner
func (s *Scanner) Peek() rune {
	r, _, _ := s.r.ReadRune()
	if r != eof {
		_ = s.r.UnreadRune()
	}

	return r
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, pos Pos, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	ch, pos := s.r.curr()
	_, _ = buf.WriteRune(ch)

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		ch, _ = s.r.read()
		if ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.r.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return WS, pos, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, pos Pos, lit string) {
	// Save the starting position of the identifier.
	_, pos = s.r.read()
	s.r.unread()

	var buf bytes.Buffer
	for {
		if ch, _ := s.r.read(); ch == eof {
			break
		} else if ch == '"' {
			tok0, pos0, lit0 := s.scanString()
			if tok0 == BADSTRING || tok0 == BADESCAPE {
				return tok0, pos0, lit0
			}
			return IDENT, pos, lit0
		} else if isIdentChar(ch) {
			s.r.unread()
			buf.WriteString(ScanBareIdent(s.r))
		} else if isPrefixed(ch) {
			buf.WriteRune(ch)
			buf.WriteString(ScanBareIdent(s.r))
		} else {
			s.r.unread()
			break
		}
	}
	lit = buf.String()

	// If the literal matches a keyword then return that keyword.
	if tok = Lookup(lit); tok != IDENT {
		return tok, pos, ""
	}

	return IDENT, pos, lit
}

// scanString consumes a contiguous string of non-quote characters.
// Quote characters can be consumed if they're first escaped with a backslash.
func (s *Scanner) scanString() (tok Token, pos Pos, lit string) {
	s.r.unread()
	_, pos = s.r.curr()

	var err error
	lit, err = ScanString(s.r)
	if err == errBadString {
		return BADSTRING, pos, lit
	} else if err == errBadEscape {
		_, pos = s.r.curr()
		return BADESCAPE, pos, lit
	}
	return STRING, pos, lit
}

// scanNumber consumes anything that looks like the start of a number.
// Numbers start with a digit, full stop, plus sign or minus sign.
// This function can return non-number tokens if a scan is a false positive.
// For example, a minus sign followed by a letter will just return a minus sign.
func (s *Scanner) scanNumber() (tok Token, pos Pos, lit string) {
	var buf bytes.Buffer
	// var isFloat bool

	// Check if the initial rune is a "+" or "-".
	ch, pos := s.r.curr()
	if ch == '+' || ch == '-' {
		// Peek at the next two runes.
		ch1, _ := s.r.read()
		ch2, _ := s.r.read()
		s.r.unread()
		s.r.unread()

		// This rune must be followed by a digit or a full stop and a digit.
		if isDigit(ch1) || (ch1 == '.' && isDigit(ch2)) {
			_, _ = buf.WriteRune(ch)
		} else if ch == '+' {
			return PLUS, pos, ""
		} else if ch == '-' {
			return MINUS, pos, ""
		}
	} else if ch == '.' {
		// Peek and see if the next rune is a digit.
		ch1, _ := s.r.read()
		s.r.unread()
		if !isDigit(ch1) {
			return ILLEGAL, pos, "."
		}

		// Unread the full stop so we can read it later.
		s.r.unread()
	} else {
		s.r.unread()
	}

	// Read as many digits as possible.
	buf.WriteString(s.scanDigits())

	// Exponent?
	if exp, _ := s.r.read(); exp == 'E' || exp == 'e' {
		buf.WriteRune('E')
		buf.WriteString(s.scanExponent())
		return DECIMAL, pos, buf.String()
	} else {
		s.r.unread()
	}

	// If next code points are a full stop and digit then consume them.
	if ch0, _ := s.r.read(); ch0 == '.' {
		if ch1, _ := s.r.read(); isDigit(ch1) {
			_, _ = buf.WriteRune(ch0)
			_, _ = buf.WriteRune(ch1)
			_, _ = buf.WriteString(s.scanDigits())
		}

		// Exponent?
		if exp, _ := s.r.read(); exp == 'E' || exp == 'e' {
			buf.WriteRune('E')
			buf.WriteString(s.scanExponent())
			return DECIMAL, pos, buf.String()
		} else {
			s.r.unread()
		}
		return DECIMAL, pos, buf.String()

	} else {
		s.r.unread()
	}

	// Attempt to read as a duration if it doesn't have a fractional part.
	if !strings.Contains(buf.String(), ".") {
		// If the next rune is a duration unit (u,µ,ms,s) then return a duration token
		if ch0, _ := s.r.read(); ch0 == 'u' || ch0 == 'µ' || ch0 == 's' || ch0 == 'h' || ch0 == 'd' || ch0 == 'w' {
			_, _ = buf.WriteRune(ch0)
			return DURATION, pos, buf.String()
		} else if ch0 == 'm' {
			_, _ = buf.WriteRune(ch0)
			if ch1, _ := s.r.read(); ch1 == 's' {
				_, _ = buf.WriteRune(ch1)
			} else {
				s.r.unread()
			}
			return DURATION, pos, buf.String()
		}
		s.r.unread()
	}
	return INTEGER, pos, buf.String()
}

// scanDigits consume a contiguous series of digits.
func (s *Scanner) scanExponent() string {
	var buf bytes.Buffer

	// Sign?
	ch1, _ := s.r.read()
	if ch1 == '+' || ch1 == '-' {
		_, _ = buf.WriteRune(ch1)
	} else {
		s.r.unread()
	}

	// Digits
	var period bool
	for {
		ch3, _ := s.r.read()
		if isDigit(ch3) {
			_, _ = buf.WriteRune(ch3)
		} else if ch3 == '.' {

			// Only allow one period
			if !period {
				period = true
				buf.WriteRune(ch3)
			} else {
				s.r.unread()
				break
			}
		} else {
			s.r.unread()
			break
		}
	}
	return buf.String()
}

// scanDigits consume a contiguous series of digits.
func (s *Scanner) scanDigits() string {
	var buf bytes.Buffer
	for {
		ch, _ := s.r.read()
		if isDigit(ch) {
			buf.WriteRune(ch)
		} else {
			s.r.unread()
			break
		}
	}
	return buf.String()
}

// ScanRegex tokenizes a regex expression
func (s *Scanner) ScanRegex() (tok Token, pos Pos, lit string) {
	_, pos = s.r.curr()

	// Start & end sentinels.
	start, end := '/', '/'
	// Valid escape chars.
	escapes := map[rune]rune{'/': '/'}

	b, err := ScanDelimited(s.r, start, end, escapes, true)

	if err == errBadEscape {
		_, pos = s.r.curr()
		return BADESCAPE, pos, lit
	} else if err != nil {
		return BADREGEX, pos, lit
	}
	return REGEX, pos, string(b)
}
