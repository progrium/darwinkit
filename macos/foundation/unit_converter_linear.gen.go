// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitConverterLinear] class.
var UnitConverterLinearClass = _UnitConverterLinearClass{objc.GetClass("NSUnitConverterLinear")}

type _UnitConverterLinearClass struct {
	objc.Class
}

// An interface definition for the [UnitConverterLinear] class.
type IUnitConverterLinear interface {
	IUnitConverter
	Constant() float64
	Coefficient() float64
}

// A description of how to convert between units using a linear equation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverterlinear?language=objc
type UnitConverterLinear struct {
	UnitConverter
}

func UnitConverterLinearFrom(ptr unsafe.Pointer) UnitConverterLinear {
	return UnitConverterLinear{
		UnitConverter: UnitConverterFrom(ptr),
	}
}

func (uc _UnitConverterLinearClass) Alloc() UnitConverterLinear {
	rv := objc.Call[UnitConverterLinear](uc, objc.Sel("alloc"))
	return rv
}

func UnitConverterLinear_Alloc() UnitConverterLinear {
	return UnitConverterLinearClass.Alloc()
}

func (uc _UnitConverterLinearClass) New() UnitConverterLinear {
	rv := objc.Call[UnitConverterLinear](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitConverterLinear() UnitConverterLinear {
	return UnitConverterLinearClass.New()
}

func (u_ UnitConverterLinear) Init() UnitConverterLinear {
	rv := objc.Call[UnitConverterLinear](u_, objc.Sel("init"))
	return rv
}

// The constant to use in the linear unit conversion calculation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverterlinear/1823598-constant?language=objc
func (u_ UnitConverterLinear) Constant() float64 {
	rv := objc.Call[float64](u_, objc.Sel("constant"))
	return rv
}

// The coefficient to use in the linear unit conversion calculation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitconverterlinear/1823683-coefficient?language=objc
func (u_ UnitConverterLinear) Coefficient() float64 {
	rv := objc.Call[float64](u_, objc.Sel("coefficient"))
	return rv
}
