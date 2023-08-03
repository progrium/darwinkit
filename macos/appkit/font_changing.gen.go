// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IFontChanging interface {
	ImplementsChangeFont() bool
	// optional
	ChangeFont(sender FontManager)
	ImplementsValidModesForFontPanel() bool
	// optional
	ValidModesForFontPanel(fontPanel FontPanel) FontPanelModeMask
}

type FontChangingWrapper struct {
	objc.Object
}

func (f_ FontChangingWrapper) ImplementsChangeFont() bool {
	return f_.RespondsToSelector(objc.GetSelector("changeFont:"))
}

func (f_ FontChangingWrapper) ChangeFont(sender IFontManager) {
	objc.CallMethod[objc.Void](f_, objc.GetSelector("changeFont:"), objc.ExtractPtr(sender))
}

func (f_ FontChangingWrapper) ImplementsValidModesForFontPanel() bool {
	return f_.RespondsToSelector(objc.GetSelector("validModesForFontPanel:"))
}

func (f_ FontChangingWrapper) ValidModesForFontPanel(fontPanel IFontPanel) FontPanelModeMask {
	rv := objc.CallMethod[FontPanelModeMask](f_, objc.GetSelector("validModesForFontPanel:"), objc.ExtractPtr(fontPanel))
	return rv
}
