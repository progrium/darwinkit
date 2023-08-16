// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods for getting and setting the appearance attributes of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization?language=objc
type PAppearanceCustomization interface {
	// optional
	EffectiveAppearance() IAppearance
	HasEffectiveAppearance() bool

	// optional
	SetAppearance(value Appearance)
	HasSetAppearance() bool

	// optional
	Appearance() IAppearance
	HasAppearance() bool
}

// A concrete type wrapper for the [PAppearanceCustomization] protocol.
type AppearanceCustomizationWrapper struct {
	objc.Object
}

func (a_ AppearanceCustomizationWrapper) HasEffectiveAppearance() bool {
	return a_.RespondsToSelector(objc.Sel("effectiveAppearance"))
}

// The appearance that will be used when the receiver is drawn onscreen, in an NSAppearance object. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization/1535147-effectiveappearance?language=objc
func (a_ AppearanceCustomizationWrapper) EffectiveAppearance() Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("effectiveAppearance"))
	return rv
}

func (a_ AppearanceCustomizationWrapper) HasSetAppearance() bool {
	return a_.RespondsToSelector(objc.Sel("setAppearance:"))
}

// The appearance of the receiver, in an NSAppearance object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization/1533925-appearance?language=objc
func (a_ AppearanceCustomizationWrapper) SetAppearance(value IAppearance) {
	objc.Call[objc.Void](a_, objc.Sel("setAppearance:"), objc.Ptr(value))
}

func (a_ AppearanceCustomizationWrapper) HasAppearance() bool {
	return a_.RespondsToSelector(objc.Sel("appearance"))
}

// The appearance of the receiver, in an NSAppearance object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsappearancecustomization/1533925-appearance?language=objc
func (a_ AppearanceCustomizationWrapper) Appearance() Appearance {
	rv := objc.Call[Appearance](a_, objc.Sel("appearance"))
	return rv
}
