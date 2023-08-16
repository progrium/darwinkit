// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NotificationQueue] class.
var NotificationQueueClass = _NotificationQueueClass{objc.GetClass("NSNotificationQueue")}

type _NotificationQueueClass struct {
	objc.Class
}

// An interface definition for the [NotificationQueue] class.
type INotificationQueue interface {
	objc.IObject
	DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint)
	EnqueueNotificationPostingStyle(notification INotification, postingStyle PostingStyle)
}

// A notification center buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue?language=objc
type NotificationQueue struct {
	objc.Object
}

func NotificationQueueFrom(ptr unsafe.Pointer) NotificationQueue {
	return NotificationQueue{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ NotificationQueue) InitWithNotificationCenter(notificationCenter INotificationCenter) NotificationQueue {
	rv := objc.Call[NotificationQueue](n_, objc.Sel("initWithNotificationCenter:"), objc.Ptr(notificationCenter))
	return rv
}

// Initializes and returns a notification queue for the specified notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1410062-initwithnotificationcenter?language=objc
func NotificationQueue_InitWithNotificationCenter(notificationCenter INotificationCenter) NotificationQueue {
	return NotificationQueueClass.Alloc().InitWithNotificationCenter(notificationCenter)
}

func (nc _NotificationQueueClass) Alloc() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("alloc"))
	return rv
}

func NotificationQueue_Alloc() NotificationQueue {
	return NotificationQueueClass.Alloc()
}

func (nc _NotificationQueueClass) New() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationQueue() NotificationQueue {
	return NotificationQueueClass.New()
}

func (n_ NotificationQueue) Init() NotificationQueue {
	rv := objc.Call[NotificationQueue](n_, objc.Sel("init"))
	return rv
}

// Removes all notifications from the queue that match a provided notification using provided matching criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1416688-dequeuenotificationsmatching?language=objc
func (n_ NotificationQueue) DequeueNotificationsMatchingCoalesceMask(notification INotification, coalesceMask uint) {
	objc.Call[objc.Void](n_, objc.Sel("dequeueNotificationsMatching:coalesceMask:"), objc.Ptr(notification), coalesceMask)
}

// Adds a notification to the notification queue with a specified posting style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1416340-enqueuenotification?language=objc
func (n_ NotificationQueue) EnqueueNotificationPostingStyle(notification INotification, postingStyle PostingStyle) {
	objc.Call[objc.Void](n_, objc.Sel("enqueueNotification:postingStyle:"), objc.Ptr(notification), postingStyle)
}

// Returns the default notification queue for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1412392-defaultqueue?language=objc
func (nc _NotificationQueueClass) DefaultQueue() NotificationQueue {
	rv := objc.Call[NotificationQueue](nc, objc.Sel("defaultQueue"))
	return rv
}

// Returns the default notification queue for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotificationqueue/1412392-defaultqueue?language=objc
func NotificationQueue_DefaultQueue() NotificationQueue {
	return NotificationQueueClass.DefaultQueue()
}
