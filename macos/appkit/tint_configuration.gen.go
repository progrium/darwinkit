// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TintConfigurationClass = _TintConfigurationClass{objc.GetClass("NSTintConfiguration")}

type _TintConfigurationClass struct {
	objc.Class
}

type ITintConfiguration interface {
	objc.IObject
	AdaptsToUserAccentColor() bool
	BaseTintColor() Color
	EquivalentContentTintColor() Color
}

type TintConfiguration struct {
	objc.Object
}

func MakeTintConfiguration(ptr unsafe.Pointer) TintConfiguration {
	return TintConfiguration{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TintConfigurationClass) TintConfigurationWithFixedColor(color IColor) TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("tintConfigurationWithFixedColor:"), objc.ExtractPtr(color))
	return rv
}

func TintConfiguration_TintConfigurationWithFixedColor(color IColor) TintConfiguration {
	return TintConfigurationClass.TintConfigurationWithFixedColor(color)
}

func (tc _TintConfigurationClass) TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("tintConfigurationWithPreferredColor:"), objc.ExtractPtr(color))
	return rv
}

func TintConfiguration_TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	return TintConfigurationClass.TintConfigurationWithPreferredColor(color)
}

func (tc _TintConfigurationClass) Alloc() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("alloc"))
	return rv
}

func TintConfiguration_Alloc() TintConfiguration {
	return TintConfigurationClass.Alloc()
}

func (tc _TintConfigurationClass) New() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTintConfiguration() TintConfiguration {
	return TintConfigurationClass.New()
}

func TintConfiguration_New() TintConfiguration {
	return TintConfigurationClass.New()
}

func (t_ TintConfiguration) Init() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](t_, objc.GetSelector("init"))
	return rv
}

func TintConfiguration_Init() TintConfiguration {
	return TintConfigurationClass.Alloc().Init()
}

func (t_ TintConfiguration) AdaptsToUserAccentColor() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("adaptsToUserAccentColor"))
	return rv
}

func (tc _TintConfigurationClass) DefaultTintConfiguration() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("defaultTintConfiguration"))
	return rv
}

func TintConfiguration_DefaultTintConfiguration() TintConfiguration {
	return TintConfigurationClass.DefaultTintConfiguration()
}

func (tc _TintConfigurationClass) MonochromeTintConfiguration() TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](tc, objc.GetSelector("monochromeTintConfiguration"))
	return rv
}

func TintConfiguration_MonochromeTintConfiguration() TintConfiguration {
	return TintConfigurationClass.MonochromeTintConfiguration()
}

func (t_ TintConfiguration) BaseTintColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("baseTintColor"))
	return rv
}

func (t_ TintConfiguration) EquivalentContentTintColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("equivalentContentTintColor"))
	return rv
}
