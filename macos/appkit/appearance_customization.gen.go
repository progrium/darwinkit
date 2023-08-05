// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IAppearanceCustomization interface {
	ImplementsSetAppearance() bool
	// optional
	SetAppearance(value Appearance)
	ImplementsAppearance() bool
	// optional
	Appearance() IAppearance
	ImplementsEffectiveAppearance() bool
	// optional
	EffectiveAppearance() IAppearance
}

type AppearanceCustomizationWrapper struct {
	objc.Object
}

func (a_ AppearanceCustomizationWrapper) ImplementsSetAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("setAppearance:"))
}

func (a_ AppearanceCustomizationWrapper) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAppearance:"), objc.ExtractPtr(value))
}

func (a_ AppearanceCustomizationWrapper) ImplementsAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("appearance"))
}

func (a_ AppearanceCustomizationWrapper) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("appearance"))
	return rv
}

func (a_ AppearanceCustomizationWrapper) ImplementsEffectiveAppearance() bool {
	return a_.RespondsToSelector(objc.GetSelector("effectiveAppearance"))
}

func (a_ AppearanceCustomizationWrapper) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("effectiveAppearance"))
	return rv
}
