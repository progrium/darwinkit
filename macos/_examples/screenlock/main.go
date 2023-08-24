package main

import (
	"log"

	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		dnc := foundation.DistributedNotificationCenter_NotificationCenterForType(foundation.LocalNotificationCenterType)
		dnc.AddObserverForNameObjectQueueUsingBlock("com.apple.screenIsLocked", nil, foundation.OperationQueue_MainQueue(), func(notification foundation.Notification) {
			log.Println("screen is locked")
		})
		dnc.AddObserverForNameObjectQueueUsingBlock("com.apple.screenIsUnlocked", nil, foundation.OperationQueue_MainQueue(), func(notification foundation.Notification) {
			log.Println("screen is unlocked")
		})
	})
}
