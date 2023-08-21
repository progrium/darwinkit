// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharedEventListener] class.
var SharedEventListenerClass = _SharedEventListenerClass{objc.GetClass("MTLSharedEventListener")}

type _SharedEventListenerClass struct {
	objc.Class
}

// An interface definition for the [SharedEventListener] class.
type ISharedEventListener interface {
	objc.IObject
	DispatchQueue() dispatch.Queue
}

// A listener for shareable event notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventlistener?language=objc
type SharedEventListener struct {
	objc.Object
}

func SharedEventListenerFrom(ptr unsafe.Pointer) SharedEventListener {
	return SharedEventListener{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SharedEventListener) InitWithDispatchQueue(dispatchQueue dispatch.Queue) SharedEventListener {
	rv := objc.Call[SharedEventListener](s_, objc.Sel("initWithDispatchQueue:"), dispatchQueue)
	return rv
}

// Creates a new shareable event listener with a specific dispatch queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventlistener/2966579-initwithdispatchqueue?language=objc
func NewSharedEventListenerWithDispatchQueue(dispatchQueue dispatch.Queue) SharedEventListener {
	instance := SharedEventListenerClass.Alloc().InitWithDispatchQueue(dispatchQueue)
	instance.Autorelease()
	return instance
}

func (s_ SharedEventListener) Init() SharedEventListener {
	rv := objc.Call[SharedEventListener](s_, objc.Sel("init"))
	return rv
}

func (sc _SharedEventListenerClass) Alloc() SharedEventListener {
	rv := objc.Call[SharedEventListener](sc, objc.Sel("alloc"))
	return rv
}

func SharedEventListener_Alloc() SharedEventListener {
	return SharedEventListenerClass.Alloc()
}

func (sc _SharedEventListenerClass) New() SharedEventListener {
	rv := objc.Call[SharedEventListener](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharedEventListener() SharedEventListener {
	return SharedEventListenerClass.New()
}

// The dispatch queue used to dispatch any notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventlistener/2966577-dispatchqueue?language=objc
func (s_ SharedEventListener) DispatchQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](s_, objc.Sel("dispatchQueue"))
	return rv
}
