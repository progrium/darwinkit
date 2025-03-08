package mapkit

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestMapKitValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, MapViewClass)
	assert.NotNil(t, AnnotationClass)
	assert.NotNil(t, PointAnnotationClass)

	// Test the constants
	assert.Equal(t, MapTypeStandard, 0)
	assert.Equal(t, MapTypeSatellite, 1)
	assert.Equal(t, MapTypeHybrid, 2)
	assert.Equal(t, MapTypeSatelliteFlyover, 3)
	assert.Equal(t, MapTypeHybridFlyover, 4)
	assert.Equal(t, MapTypeMutedStandard, 5)
}
