// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InputManager] class.
var InputManagerClass = _InputManagerClass{objc.GetClass("NSInputManager")}

type _InputManagerClass struct {
	objc.Class
}

// An interface definition for the [InputManager] class.
type IInputManager interface {
	objc.IObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsinputmanager?language=objc
type InputManager struct {
	objc.Object
}

func InputManagerFrom(ptr unsafe.Pointer) InputManager {
	return InputManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _InputManagerClass) Alloc() InputManager {
	rv := objc.Call[InputManager](ic, objc.Sel("alloc"))
	return rv
}

func InputManager_Alloc() InputManager {
	return InputManagerClass.Alloc()
}

func (ic _InputManagerClass) New() InputManager {
	rv := objc.Call[InputManager](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInputManager() InputManager {
	return InputManagerClass.New()
}

func (i_ InputManager) Init() InputManager {
	rv := objc.Call[InputManager](i_, objc.Sel("init"))
	return rv
}
