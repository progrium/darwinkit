// auto-generated code, do not modify
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var NotificationCenterClass = _NotificationCenterClass{objc.GetClass("NSNotificationCenter")}

type _NotificationCenterClass struct {
	objc.Class
}

type INotificationCenter interface {
	objc.IObject
	AddObserverSelectorNameObject(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject)
	RemoveObserverNameObject(observer objc.IObject, aName NotificationName, anObject objc.IObject)
	RemoveObserver(observer objc.IObject)
	PostNotification(notification INotification)
	PostNotificationNameObject(aName NotificationName, anObject objc.IObject)
}

type NotificationCenter struct {
	objc.Object
}

func MakeNotificationCenter(ptr unsafe.Pointer) NotificationCenter {
	return NotificationCenter{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NotificationCenterClass) Alloc() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](nc, objc.GetSelector("alloc"))
	return rv
}

func NotificationCenter_Alloc() NotificationCenter {
	return NotificationCenterClass.Alloc()
}

func (nc _NotificationCenterClass) New() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationCenter() NotificationCenter {
	return NotificationCenterClass.New()
}

func NotificationCenter_New() NotificationCenter {
	return NotificationCenterClass.New()
}

func (n_ NotificationCenter) Init() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](n_, objc.GetSelector("init"))
	return rv
}

func NotificationCenter_Init() NotificationCenter {
	return NotificationCenterClass.Alloc().Init()
}

func (n_ NotificationCenter) AddObserverSelectorNameObject(observer objc.IObject, aSelector objc.Selector, aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("addObserver:selector:name:object:"), objc.ExtractPtr(observer), aSelector, aName, objc.ExtractPtr(anObject))
}

func (n_ NotificationCenter) RemoveObserverNameObject(observer objc.IObject, aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("removeObserver:name:object:"), objc.ExtractPtr(observer), aName, objc.ExtractPtr(anObject))
}

func (n_ NotificationCenter) RemoveObserver(observer objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("removeObserver:"), objc.ExtractPtr(observer))
}

func (n_ NotificationCenter) PostNotification(notification INotification) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("postNotification:"), objc.ExtractPtr(notification))
}

func (n_ NotificationCenter) PostNotificationNameObject(aName NotificationName, anObject objc.IObject) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("postNotificationName:object:"), aName, objc.ExtractPtr(anObject))
}

func (nc _NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := objc.CallMethod[NotificationCenter](nc, objc.GetSelector("defaultCenter"))
	return rv
}

func NotificationCenter_DefaultCenter() NotificationCenter {
	return NotificationCenterClass.DefaultCenter()
}
