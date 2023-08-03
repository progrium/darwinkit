package oldgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer_NextToken(t *testing.T) {
	tokenizer := Tokenizer{Text: "NSUInteger"}
	token, ok := tokenizer.NextToken()
	assert.True(t, ok)
	assert.Equal(t, "NSUInteger", token)

	tokenizer = Tokenizer{Text: "NSArray<Foundation.NSRunLoopMode> * "}
	token, ok = tokenizer.NextToken()
	assert.Equal(t, "NSArray", token)
	token, ok = tokenizer.NextToken()
	assert.Equal(t, "<", token)
	token, ok = tokenizer.NextToken()
	assert.Equal(t, "Foundation.NSRunLoopMode", token)
	token, ok = tokenizer.NextToken()
	assert.Equal(t, ">", token)
	token, ok = tokenizer.NextToken()
	assert.Equal(t, "*", token)

	tokenizer.PutBackToken(token)
	assert.Equal(t, byte(' '), tokenizer.Text[tokenizer.index])
	token, ok = tokenizer.NextToken()
	assert.Equal(t, "*", token)
	token, ok = tokenizer.NextToken()
	assert.False(t, ok)
}

func TestTokenizer_PutBackToken(t *testing.T) {

}
