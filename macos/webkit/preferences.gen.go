// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PreferencesClass = _PreferencesClass{objc.GetClass("WKPreferences")}

type _PreferencesClass struct {
	objc.Class
}

type IPreferences interface {
	objc.IObject
	MinimumFontSize() float64
	SetMinimumFontSize(value float64)
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)
	IsElementFullscreenEnabled() bool
	SetElementFullscreenEnabled(value bool)
	IsSiteSpecificQuirksModeEnabled() bool
	SetSiteSpecificQuirksModeEnabled(value bool)
	IsTextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
}

type Preferences struct {
	objc.Object
}

func MakePreferences(ptr unsafe.Pointer) Preferences {
	return Preferences{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PreferencesClass) Alloc() Preferences {
	rv := objc.CallMethod[Preferences](pc, objc.GetSelector("alloc"))
	return rv
}

func Preferences_Alloc() Preferences {
	return PreferencesClass.Alloc()
}

func (pc _PreferencesClass) New() Preferences {
	rv := objc.CallMethod[Preferences](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPreferences() Preferences {
	return PreferencesClass.New()
}

func Preferences_New() Preferences {
	return PreferencesClass.New()
}

func (p_ Preferences) Init() Preferences {
	rv := objc.CallMethod[Preferences](p_, objc.GetSelector("init"))
	return rv
}

func Preferences_Init() Preferences {
	return PreferencesClass.Alloc().Init()
}

func (p_ Preferences) MinimumFontSize() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("minimumFontSize"))
	return rv
}

func (p_ Preferences) SetMinimumFontSize(value float64) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setMinimumFontSize:"), value)
}

func (p_ Preferences) TabFocusesLinks() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("tabFocusesLinks"))
	return rv
}

func (p_ Preferences) SetTabFocusesLinks(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTabFocusesLinks:"), value)
}

func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("javaScriptCanOpenWindowsAutomatically"))
	return rv
}

func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setJavaScriptCanOpenWindowsAutomatically:"), value)
}

func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isFraudulentWebsiteWarningEnabled"))
	return rv
}

func (p_ Preferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setFraudulentWebsiteWarningEnabled:"), value)
}

func (p_ Preferences) IsElementFullscreenEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isElementFullscreenEnabled"))
	return rv
}

func (p_ Preferences) SetElementFullscreenEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setElementFullscreenEnabled:"), value)
}

func (p_ Preferences) IsSiteSpecificQuirksModeEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isSiteSpecificQuirksModeEnabled"))
	return rv
}

func (p_ Preferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setSiteSpecificQuirksModeEnabled:"), value)
}

func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("isTextInteractionEnabled"))
	return rv
}

func (p_ Preferences) SetTextInteractionEnabled(value bool) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setTextInteractionEnabled:"), value)
}
