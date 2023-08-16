// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging?language=objc
type PFontChanging interface {
	// optional
	ChangeFont(sender FontManager)
	HasChangeFont() bool

	// optional
	ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask
	HasValidModesForFontPanel() bool
}

// A concrete type wrapper for the [PFontChanging] protocol.
type FontChangingWrapper struct {
	objc.Object
}

func (f_ FontChangingWrapper) HasChangeFont() bool {
	return f_.RespondsToSelector(objc.Sel("changeFont:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging/3005180-changefont?language=objc
func (f_ FontChangingWrapper) ChangeFont(sender IFontManager) {
	objc.Call[objc.Void](f_, objc.Sel("changeFont:"), objc.Ptr(sender))
}

func (f_ FontChangingWrapper) HasValidModesForFontPanel() bool {
	return f_.RespondsToSelector(objc.Sel("validModesForFontPanel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontchanging/3005181-validmodesforfontpanel?language=objc
func (f_ FontChangingWrapper) ValidModesForFontPanel(fontPanel IFontPanel) FontPanelModeMask {
	rv := objc.Call[FontPanelModeMask](f_, objc.Sel("validModesForFontPanel:"), objc.Ptr(fontPanel))
	return rv
}
