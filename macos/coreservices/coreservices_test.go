package coreservices

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestCoreServicesValid(t *testing.T) {
	// Test constants
	assert.Equal(t, LSLaunchDefaults, 0x00000001)
	assert.Equal(t, LSLaunchAndPrint, 0x00000002)
	assert.Equal(t, LSLaunchAsync, 0x00010000)
	
	// Test FSEvents constants
	assert.Equal(t, FSEventStreamCreateFlagNone, FSEventStreamCreateFlags(0x00000000))
	assert.Equal(t, FSEventStreamCreateFlagUseCFTypes, FSEventStreamCreateFlags(0x00000001))
	assert.Equal(t, FSEventStreamCreateFlagFileEvents, FSEventStreamCreateFlags(0x00000010))
	
	// Test event flags
	assert.Equal(t, FSEventStreamEventFlagNone, FSEventStreamEventFlags(0x00000000))
	assert.Equal(t, FSEventStreamEventFlagItemCreated, FSEventStreamEventFlags(0x00000100))
	assert.Equal(t, FSEventStreamEventFlagItemModified, FSEventStreamEventFlags(0x00001000))
}