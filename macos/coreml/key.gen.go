// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Key] class.
var KeyClass = _KeyClass{objc.GetClass("MLKey")}

type _KeyClass struct {
	objc.Class
}

// An interface definition for the [Key] class.
type IKey interface {
	objc.IObject
	Name() string
	Scope() string
}

// An abstract base class for machine learning key types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlkey?language=objc
type Key struct {
	objc.Object
}

func KeyFrom(ptr unsafe.Pointer) Key {
	return Key{
		Object: objc.ObjectFrom(ptr),
	}
}

func (kc _KeyClass) Alloc() Key {
	rv := objc.Call[Key](kc, objc.Sel("alloc"))
	return rv
}

func Key_Alloc() Key {
	return KeyClass.Alloc()
}

func (kc _KeyClass) New() Key {
	rv := objc.Call[Key](kc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewKey() Key {
	return KeyClass.New()
}

func (k_ Key) Init() Key {
	rv := objc.Call[Key](k_, objc.Sel("init"))
	return rv
}

// The name of the machine learning key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlkey/3180057-name?language=objc
func (k_ Key) Name() string {
	rv := objc.Call[string](k_, objc.Sel("name"))
	return rv
}

// The applicable scope of the machine learning key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlkey/3333248-scope?language=objc
func (k_ Key) Scope() string {
	rv := objc.Call[string](k_, objc.Sel("scope"))
	return rv
}
