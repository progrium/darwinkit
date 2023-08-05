// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var UnitClass = _UnitClass{objc.GetClass("NSUnit")}

type _UnitClass struct {
	objc.Class
}

type IUnit interface {
	objc.IObject
	Symbol() string
}

type Unit struct {
	objc.Object
}

func MakeUnit(ptr unsafe.Pointer) Unit {
	return Unit{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ Unit) InitWithSymbol(symbol string) Unit {
	rv := objc.CallMethod[Unit](u_, objc.GetSelector("initWithSymbol:"), symbol)
	return rv
}

func Unit_InitWithSymbol(symbol string) Unit {
	return UnitClass.Alloc().InitWithSymbol(symbol)
}

func (uc _UnitClass) Alloc() Unit {
	rv := objc.CallMethod[Unit](uc, objc.GetSelector("alloc"))
	return rv
}

func Unit_Alloc() Unit {
	return UnitClass.Alloc()
}

func (uc _UnitClass) New() Unit {
	rv := objc.CallMethod[Unit](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUnit() Unit {
	return UnitClass.New()
}

func Unit_New() Unit {
	return UnitClass.New()
}

func (u_ Unit) Init() Unit {
	rv := objc.CallMethod[Unit](u_, objc.GetSelector("init"))
	return rv
}

func Unit_Init() Unit {
	return UnitClass.Alloc().Init()
}

func (u_ Unit) Symbol() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("symbol"))
	return rv
}
