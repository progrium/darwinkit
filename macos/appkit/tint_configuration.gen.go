// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TintConfiguration] class.
var TintConfigurationClass = _TintConfigurationClass{objc.GetClass("NSTintConfiguration")}

type _TintConfigurationClass struct {
	objc.Class
}

// An interface definition for the [TintConfiguration] class.
type ITintConfiguration interface {
	objc.IObject
	BaseTintColor() Color
	EquivalentContentTintColor() Color
	AdaptsToUserAccentColor() bool
}

// An object that gives you the ability to choose from system-provided tinting behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration?language=objc
type TintConfiguration struct {
	objc.Object
}

func TintConfigurationFrom(ptr unsafe.Pointer) TintConfiguration {
	return TintConfiguration{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TintConfigurationClass) TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("tintConfigurationWithPreferredColor:"), objc.Ptr(color))
	return rv
}

// Creates a new tint configuration for the system to use when the app’s preferred accent color is in use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626824-tintconfigurationwithpreferredco?language=objc
func TintConfiguration_TintConfigurationWithPreferredColor(color IColor) TintConfiguration {
	return TintConfigurationClass.TintConfigurationWithPreferredColor(color)
}

func (tc _TintConfigurationClass) TintConfigurationWithFixedColor(color IColor) TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("tintConfigurationWithFixedColor:"), objc.Ptr(color))
	return rv
}

// Creates a new tint configuration using a specific color value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626823-tintconfigurationwithfixedcolor?language=objc
func TintConfiguration_TintConfigurationWithFixedColor(color IColor) TintConfiguration {
	return TintConfigurationClass.TintConfigurationWithFixedColor(color)
}

func (tc _TintConfigurationClass) Alloc() TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("alloc"))
	return rv
}

func TintConfiguration_Alloc() TintConfiguration {
	return TintConfigurationClass.Alloc()
}

func (tc _TintConfigurationClass) New() TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTintConfiguration() TintConfiguration {
	return TintConfigurationClass.New()
}

func (t_ TintConfiguration) Init() TintConfiguration {
	rv := objc.Call[TintConfiguration](t_, objc.Sel("init"))
	return rv
}

// The color the system supplies when you create a tint configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626819-basetintcolor?language=objc
func (t_ TintConfiguration) BaseTintColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("baseTintColor"))
	return rv
}

// A color object that matches the effective content tint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626821-equivalentcontenttintcolor?language=objc
func (t_ TintConfiguration) EquivalentContentTintColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("equivalentContentTintColor"))
	return rv
}

// The system tints the content using the system default value for its context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626820-defaulttintconfiguration?language=objc
func (tc _TintConfigurationClass) DefaultTintConfiguration() TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("defaultTintConfiguration"))
	return rv
}

// The system tints the content using the system default value for its context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626820-defaulttintconfiguration?language=objc
func TintConfiguration_DefaultTintConfiguration() TintConfiguration {
	return TintConfigurationClass.DefaultTintConfiguration()
}

// The content always displays in monochrome. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626822-monochrometintconfiguration?language=objc
func (tc _TintConfigurationClass) MonochromeTintConfiguration() TintConfiguration {
	rv := objc.Call[TintConfiguration](tc, objc.Sel("monochromeTintConfiguration"))
	return rv
}

// The content always displays in monochrome. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626822-monochrometintconfiguration?language=objc
func TintConfiguration_MonochromeTintConfiguration() TintConfiguration {
	return TintConfigurationClass.MonochromeTintConfiguration()
}

// A Boolean value that indicates whether the tint configuration alters its effect based on the user’s preferred accent color choice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstintconfiguration/3626818-adaptstouseraccentcolor?language=objc
func (t_ TintConfiguration) AdaptsToUserAccentColor() bool {
	rv := objc.Call[bool](t_, objc.Sel("adaptsToUserAccentColor"))
	return rv
}
