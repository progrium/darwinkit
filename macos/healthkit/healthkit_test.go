package healthkit

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestHealthKitValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, HealthStoreClass)
	assert.NotNil(t, QuantityTypeClass)
	assert.NotNil(t, QuantitySampleClass)
	assert.NotNil(t, QuantityClass)
	assert.NotNil(t, UnitClass)
	assert.NotNil(t, WorkoutTypeClass)
	assert.NotNil(t, WorkoutClass)

	// Test the constants
	assert.Equal(t, WorkoutActivityTypeRunning, 8)
	assert.Equal(t, WorkoutActivityTypeWalking, 7)
	assert.Equal(t, WorkoutActivityTypeCycling, 13)
	assert.Equal(t, MetricPrefixKilo, 9)
	assert.Equal(t, MetricPrefixMilli, 4)
}