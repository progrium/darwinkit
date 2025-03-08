package homekit

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestHomeKitValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, HomeManagerClass)
	assert.NotNil(t, HomeClass)
	assert.NotNil(t, RoomClass)
	assert.NotNil(t, AccessoryClass)
	assert.NotNil(t, ServiceClass)
	assert.NotNil(t, CharacteristicClass)

	// Test the constants
	assert.Equal(t, ServiceTypeLightbulb, "00000043-0000-1000-8000-0026BB765291")
	assert.Equal(t, ServiceTypeSwitch, "00000049-0000-1000-8000-0026BB765291")
	assert.Equal(t, CharacteristicTypePowerState, "00000025-0000-1000-8000-0026BB765291")
	assert.Equal(t, CharacteristicTypeBrightness, "00000008-0000-1000-8000-0026BB765291")
}