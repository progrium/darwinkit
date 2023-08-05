// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var NotificationClass = _NotificationClass{objc.GetClass("NSNotification")}

type _NotificationClass struct {
	objc.Class
}

type INotification interface {
	objc.IObject
	Name() NotificationName
	Object_() objc.Object
}

type Notification struct {
	objc.Object
}

func MakeNotification(ptr unsafe.Pointer) Notification {
	return Notification{
		Object: objc.MakeObject(ptr),
	}
}

func (n_ Notification) Init() Notification {
	rv := objc.CallMethod[Notification](n_, objc.GetSelector("init"))
	return rv
}

func Notification_Init() Notification {
	return NotificationClass.Alloc().Init()
}

func (nc _NotificationClass) NotificationWithNameObject(aName NotificationName, anObject objc.IObject) Notification {
	rv := objc.CallMethod[Notification](nc, objc.GetSelector("notificationWithName:object:"), aName, objc.ExtractPtr(anObject))
	return rv
}

func Notification_NotificationWithNameObject(aName NotificationName, anObject objc.IObject) Notification {
	return NotificationClass.NotificationWithNameObject(aName, anObject)
}

func (nc _NotificationClass) Alloc() Notification {
	rv := objc.CallMethod[Notification](nc, objc.GetSelector("alloc"))
	return rv
}

func Notification_Alloc() Notification {
	return NotificationClass.Alloc()
}

func (nc _NotificationClass) New() Notification {
	rv := objc.CallMethod[Notification](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNotification() Notification {
	return NotificationClass.New()
}

func Notification_New() Notification {
	return NotificationClass.New()
}

func (n_ Notification) Name() NotificationName {
	rv := objc.CallMethod[NotificationName](n_, objc.GetSelector("name"))
	return rv
}

func (n_ Notification) Object_() objc.Object {
	rv := objc.CallMethod[objc.Object](n_, objc.GetSelector("object"))
	return rv
}
