// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharedEventHandle] class.
var SharedEventHandleClass = _SharedEventHandleClass{objc.GetClass("MTLSharedEventHandle")}

type _SharedEventHandleClass struct {
	objc.Class
}

// An interface definition for the [SharedEventHandle] class.
type ISharedEventHandle interface {
	objc.IObject
	Label() string
}

// An object you use to recreate a shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventhandle?language=objc
type SharedEventHandle struct {
	objc.Object
}

func SharedEventHandleFrom(ptr unsafe.Pointer) SharedEventHandle {
	return SharedEventHandle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SharedEventHandleClass) Alloc() SharedEventHandle {
	rv := objc.Call[SharedEventHandle](sc, objc.Sel("alloc"))
	return rv
}

func SharedEventHandle_Alloc() SharedEventHandle {
	return SharedEventHandleClass.Alloc()
}

func (sc _SharedEventHandleClass) New() SharedEventHandle {
	rv := objc.Call[SharedEventHandle](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharedEventHandle() SharedEventHandle {
	return SharedEventHandleClass.New()
}

func (s_ SharedEventHandle) Init() SharedEventHandle {
	rv := objc.Call[SharedEventHandle](s_, objc.Sel("init"))
	return rv
}

// A string that identifies the shareable event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventhandle/2981027-label?language=objc
func (s_ SharedEventHandle) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}
