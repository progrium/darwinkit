package eventkit

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestEventKitValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, EventStoreClass)
	assert.NotNil(t, EventClass)
	assert.NotNil(t, CalendarClass)

	// Test the constants
	assert.Equal(t, EntityTypeEvent, 0)
	assert.Equal(t, EntityTypeReminder, 1)
	assert.Equal(t, SpanThisEvent, 0)
	assert.Equal(t, SpanFutureEvents, 1)
	assert.Equal(t, SpanAllEvents, 2)
}
