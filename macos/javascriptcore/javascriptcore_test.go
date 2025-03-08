package javascriptcore

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestJavaScriptCoreValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, ContextClass)
	assert.NotNil(t, ValueClass)
	assert.NotNil(t, VirtualMachineClass)
}