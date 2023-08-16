// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Floor] class.
var FloorClass = _FloorClass{objc.GetClass("CLFloor")}

type _FloorClass struct {
	objc.Class
}

// An interface definition for the [Floor] class.
type IFloor interface {
	objc.IObject
	Level() int
}

// The floor of a building on which the user's device is located. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clfloor?language=objc
type Floor struct {
	objc.Object
}

func FloorFrom(ptr unsafe.Pointer) Floor {
	return Floor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (fc _FloorClass) Alloc() Floor {
	rv := objc.Call[Floor](fc, objc.Sel("alloc"))
	return rv
}

func Floor_Alloc() Floor {
	return FloorClass.Alloc()
}

func (fc _FloorClass) New() Floor {
	rv := objc.Call[Floor](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFloor() Floor {
	return FloorClass.New()
}

func (f_ Floor) Init() Floor {
	rv := objc.Call[Floor](f_, objc.Sel("init"))
	return rv
}

// The logical floor of the building. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clfloor/1616759-level?language=objc
func (f_ Floor) Level() int {
	rv := objc.Call[int](f_, objc.Sel("level"))
	return rv
}
