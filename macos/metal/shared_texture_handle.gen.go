// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharedTextureHandle] class.
var SharedTextureHandleClass = _SharedTextureHandleClass{objc.GetClass("MTLSharedTextureHandle")}

type _SharedTextureHandleClass struct {
	objc.Class
}

// An interface definition for the [SharedTextureHandle] class.
type ISharedTextureHandle interface {
	objc.IObject
	Device() DeviceWrapper
	Label() string
}

// A texture handle that can be shared across process address space boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedtexturehandle?language=objc
type SharedTextureHandle struct {
	objc.Object
}

func SharedTextureHandleFrom(ptr unsafe.Pointer) SharedTextureHandle {
	return SharedTextureHandle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SharedTextureHandleClass) Alloc() SharedTextureHandle {
	rv := objc.Call[SharedTextureHandle](sc, objc.Sel("alloc"))
	return rv
}

func SharedTextureHandle_Alloc() SharedTextureHandle {
	return SharedTextureHandleClass.Alloc()
}

func (sc _SharedTextureHandleClass) New() SharedTextureHandle {
	rv := objc.Call[SharedTextureHandle](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharedTextureHandle() SharedTextureHandle {
	return SharedTextureHandleClass.New()
}

func (s_ SharedTextureHandle) Init() SharedTextureHandle {
	rv := objc.Call[SharedTextureHandle](s_, objc.Sel("init"))
	return rv
}

// The device object that created the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedtexturehandle/2967448-device?language=objc
func (s_ SharedTextureHandle) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](s_, objc.Sel("device"))
	return rv
}

// A string that identifies the texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedtexturehandle/2981031-label?language=objc
func (s_ SharedTextureHandle) Label() string {
	rv := objc.Call[string](s_, objc.Sel("label"))
	return rv
}
