// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLHandle] class.
var URLHandleClass = _URLHandleClass{objc.GetClass("NSURLHandle")}

type _URLHandleClass struct {
	objc.Class
}

// An interface definition for the [URLHandle] class.
type IURLHandle interface {
	objc.IObject
}

// An object that accesses and manages resource data indicated by a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlhandle?language=objc
type URLHandle struct {
	objc.Object
}

func URLHandleFrom(ptr unsafe.Pointer) URLHandle {
	return URLHandle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLHandleClass) Alloc() URLHandle {
	rv := objc.Call[URLHandle](uc, objc.Sel("alloc"))
	return rv
}

func URLHandle_Alloc() URLHandle {
	return URLHandleClass.Alloc()
}

func (uc _URLHandleClass) New() URLHandle {
	rv := objc.Call[URLHandle](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLHandle() URLHandle {
	return URLHandleClass.New()
}

func (u_ URLHandle) Init() URLHandle {
	rv := objc.Call[URLHandle](u_, objc.Sel("init"))
	return rv
}
