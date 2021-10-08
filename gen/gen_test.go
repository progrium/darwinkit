package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportedName(t *testing.T) {
	t.Run("lower-case", func(t *testing.T) {
		assert.Equal(t, "FooBar", toExportedName("fooBar"))
	})
	t.Run("already exported", func(t *testing.T) {
		assert.Equal(t, "FooBar", toExportedName("FooBar"))
	})
}

func TestParseClassDeclaration(t *testing.T) {
	d := parseClassDeclaration("@interface NSString : NSObject")
	assert.Equal(t, declaration{
		Name: "NSString",
		Base: "NSObject",
	}, d)
}
