package security

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestSecurityValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, CertificateClass)
	assert.NotNil(t, IdentityClass)
	assert.NotNil(t, PolicyClass)
	assert.NotNil(t, TrustClass)

	// Test the constants
	assert.Equal(t, ErrSecSuccess, OSStatus(0))
	assert.Equal(t, ErrSecItemNotFound, OSStatus(-25300))
	assert.Equal(t, TrustResultProceed, TrustResultType(1))
	assert.Equal(t, TrustResultDeny, TrustResultType(3))
}