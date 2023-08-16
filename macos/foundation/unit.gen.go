// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Unit] class.
var UnitClass = _UnitClass{objc.GetClass("NSUnit")}

type _UnitClass struct {
	objc.Class
}

// An interface definition for the [Unit] class.
type IUnit interface {
	objc.IObject
	Symbol() string
}

// An abstract class representing a unit of measure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit?language=objc
type Unit struct {
	objc.Object
}

func UnitFrom(ptr unsafe.Pointer) Unit {
	return Unit{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ Unit) InitWithSymbol(symbol string) Unit {
	rv := objc.Call[Unit](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func Unit_InitWithSymbol(symbol string) Unit {
	return UnitClass.Alloc().InitWithSymbol(symbol)
}

func (uc _UnitClass) Alloc() Unit {
	rv := objc.Call[Unit](uc, objc.Sel("alloc"))
	return rv
}

func Unit_Alloc() Unit {
	return UnitClass.Alloc()
}

func (uc _UnitClass) New() Unit {
	rv := objc.Call[Unit](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnit() Unit {
	return UnitClass.New()
}

func (u_ Unit) Init() Unit {
	rv := objc.Call[Unit](u_, objc.Sel("init"))
	return rv
}

// The symbolic representation of the unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1642700-symbol?language=objc
func (u_ Unit) Symbol() string {
	rv := objc.Call[string](u_, objc.Sel("symbol"))
	return rv
}
