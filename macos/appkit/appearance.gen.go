// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var AppearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}

type _AppearanceClass struct {
	objc.Class
}

type IAppearance interface {
	objc.IObject
	BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName
	PerformAsCurrentDrawingAppearance(block func())
	Name() AppearanceName
	AllowsVibrancy() bool
}

type Appearance struct {
	objc.Object
}

func MakeAppearance(ptr unsafe.Pointer) Appearance {
	return Appearance{
		Object: objc.MakeObject(ptr),
	}
}

func (a_ Appearance) InitWithAppearanceNamedBundle(name AppearanceName, bundle foundation.IBundle) Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("initWithAppearanceNamed:bundle:"), name, objc.ExtractPtr(bundle))
	return rv
}

func Appearance_InitWithAppearanceNamedBundle(name AppearanceName, bundle foundation.IBundle) Appearance {
	return AppearanceClass.Alloc().InitWithAppearanceNamedBundle(name, bundle)
}

func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.GetSelector("alloc"))
	return rv
}

func Appearance_Alloc() Appearance {
	return AppearanceClass.Alloc()
}

func (ac _AppearanceClass) New() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewAppearance() Appearance {
	return AppearanceClass.New()
}

func Appearance_New() Appearance {
	return AppearanceClass.New()
}

func (a_ Appearance) Init() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("init"))
	return rv
}

func Appearance_Init() Appearance {
	return AppearanceClass.Alloc().Init()
}

func (ac _AppearanceClass) AppearanceNamed(name AppearanceName) Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.GetSelector("appearanceNamed:"), name)
	return rv
}

func Appearance_AppearanceNamed(name AppearanceName) Appearance {
	return AppearanceClass.AppearanceNamed(name)
}

func (a_ Appearance) BestMatchFromAppearancesWithNames(appearances []AppearanceName) AppearanceName {
	rv := objc.CallMethod[AppearanceName](a_, objc.GetSelector("bestMatchFromAppearancesWithNames:"), appearances)
	return rv
}

func (a_ Appearance) PerformAsCurrentDrawingAppearance(block func()) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("performAsCurrentDrawingAppearance:"), block)
}

func (a_ Appearance) Name() AppearanceName {
	rv := objc.CallMethod[AppearanceName](a_, objc.GetSelector("name"))
	return rv
}

func (ac _AppearanceClass) CurrentDrawingAppearance() Appearance {
	rv := objc.CallMethod[Appearance](ac, objc.GetSelector("currentDrawingAppearance"))
	return rv
}

func Appearance_CurrentDrawingAppearance() Appearance {
	return AppearanceClass.CurrentDrawingAppearance()
}

func (a_ Appearance) AllowsVibrancy() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("allowsVibrancy"))
	return rv
}
