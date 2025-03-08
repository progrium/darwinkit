package usernotifications

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestUserNotificationsValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, NotificationCenterClass)
	assert.NotNil(t, NotificationRequestClass)
	assert.NotNil(t, NotificationContentClass)
	assert.NotNil(t, MutableNotificationContentClass)
	assert.NotNil(t, NotificationTriggerClass)
	assert.NotNil(t, TimeIntervalNotificationTriggerClass)
	assert.NotNil(t, CalendarNotificationTriggerClass)
	assert.NotNil(t, LocationNotificationTriggerClass)
	assert.NotNil(t, PushNotificationTriggerClass)
	assert.NotNil(t, NotificationSoundClass)

	// Test the constants
	assert.Equal(t, uint(AuthorizationOptionBadge), uint(1))
	assert.Equal(t, uint(AuthorizationOptionSound), uint(2))
	assert.Equal(t, uint(AuthorizationOptionAlert), uint(4))
	assert.Equal(t, uint(AuthorizationOptionCarPlay), uint(8))
}