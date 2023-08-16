// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorchanging?language=objc
type PColorChanging interface {
	// optional
	ChangeColor(sender ColorPanel)
	HasChangeColor() bool
}

// A concrete type wrapper for the [PColorChanging] protocol.
type ColorChangingWrapper struct {
	objc.Object
}

func (c_ ColorChangingWrapper) HasChangeColor() bool {
	return c_.RespondsToSelector(objc.Sel("changeColor:"))
}

// Sent to the first responder when the user selects a color in an NSColorPanel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorchanging/3005175-changecolor?language=objc
func (c_ ColorChangingWrapper) ChangeColor(sender IColorPanel) {
	objc.Call[objc.Void](c_, objc.Sel("changeColor:"), objc.Ptr(sender))
}
