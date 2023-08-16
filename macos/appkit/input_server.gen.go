// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InputServer] class.
var InputServerClass = _InputServerClass{objc.GetClass("NSInputServer")}

type _InputServerClass struct {
	objc.Class
}

// An interface definition for the [InputServer] class.
type IInputServer interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsinputserver?language=objc
type InputServer struct {
	objc.Object
}

func InputServerFrom(ptr unsafe.Pointer) InputServer {
	return InputServer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _InputServerClass) Alloc() InputServer {
	rv := objc.Call[InputServer](ic, objc.Sel("alloc"))
	return rv
}

func InputServer_Alloc() InputServer {
	return InputServerClass.Alloc()
}

func (ic _InputServerClass) New() InputServer {
	rv := objc.Call[InputServer](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInputServer() InputServer {
	return InputServerClass.New()
}

func (i_ InputServer) Init() InputServer {
	rv := objc.Call[InputServer](i_, objc.Sel("init"))
	return rv
}
