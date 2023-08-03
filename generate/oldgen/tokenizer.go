package oldgen

// Tokenizer is simple tokenizer to parse objc signature
type Tokenizer struct {
	Text  string
	index int
}

func (t *Tokenizer) nextChar() (char byte, ok bool) {
	if t.index >= len(t.Text) {
		return 0, false
	}
	c := t.Text[t.index]
	t.index++
	return c, true
}

func (t *Tokenizer) putBackChar() {
	if t.index < 1 {
		panic("index <= 0")
	}
	t.index--
}

func (t *Tokenizer) MustNextToken() string {
	token, ok := t.NextToken()
	if !ok {
		panic("unexpected end")
	}
	return token
}

func (t *Tokenizer) MustNextTokenIs(expect string) {
	token, ok := t.NextToken()
	if !ok {
		panic("unexpected end")
	}
	if token != expect {
		panic("expect token: " + expect)
	}
}

func (t *Tokenizer) MustNextTill(c byte) string {
	begin := t.index
	for {
		if t.Text[t.index] == c {
			break
		}
		t.index++
		if t.index >= len(t.Text) {
			panic("unexpected end")
		}
	}
	return t.Text[begin:t.index]
}

func (t *Tokenizer) MustNextTillClosingAngle() string {
	var openingCount = 1
	begin := t.index
	for {
		c := t.Text[t.index]
		if c == '>' {
			openingCount--
			if openingCount == 0 {
				break
			}
		} else if c == '<' {
			openingCount++
		}
		t.index++
		if t.index >= len(t.Text) {
			panic("unexpected end")
		}
	}
	return t.Text[begin:t.index]
}

func (t *Tokenizer) NextToken() (token string, ok bool) {
	begin := t.index
	c, ok := t.nextChar()
	if !ok {
		return "", false
	}
	// skip white space
Outer:
	for ok {
		switch c {
		case ' ', '\t', '\r', '\n':
			c, ok = t.nextChar()
		default:
			begin = t.index - 1
			break Outer
		}
	}

	if !ok {
		return "", false
	}

	// next token
	switch c {
	case '*', '<', '>', '(', '^', ')', ',':
		return t.Text[begin:t.index], true
	default:
		for {
			c, ok = t.nextChar()
			if !ok {
				return t.Text[begin:t.index], true
			}
			switch c {
			case ' ', '\t', '\r', '\n':
				fallthrough
			case '*', '<', '>', '(', '^', ')', ',':
				t.putBackChar()
				return t.Text[begin:t.index], true
			default:
			}
		}
	}
}

func (t *Tokenizer) PutBackToken(token string) {
	if t.index < len(token) {
		panic("index <= 0")
	}
	t.index = t.index - len(token)
Outer:
	for t.index > 0 {
		c := t.Text[t.index-1]
		switch c {
		case ' ', '\t', '\r', '\n':
			t.index--
		default:
			break Outer
		}
	}
}
