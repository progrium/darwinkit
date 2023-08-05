// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var WebpagePreferencesClass = _WebpagePreferencesClass{objc.GetClass("WKWebpagePreferences")}

type _WebpagePreferencesClass struct {
	objc.Class
}

type IWebpagePreferences interface {
	objc.IObject
	AllowsContentJavaScript() bool
	SetAllowsContentJavaScript(value bool)
	PreferredContentMode() ContentMode
	SetPreferredContentMode(value ContentMode)
	IsLockdownModeEnabled() bool
	SetLockdownModeEnabled(value bool)
}

type WebpagePreferences struct {
	objc.Object
}

func MakeWebpagePreferences(ptr unsafe.Pointer) WebpagePreferences {
	return WebpagePreferences{
		Object: objc.MakeObject(ptr),
	}
}

func (wc _WebpagePreferencesClass) Alloc() WebpagePreferences {
	rv := objc.CallMethod[WebpagePreferences](wc, objc.GetSelector("alloc"))
	return rv
}

func WebpagePreferences_Alloc() WebpagePreferences {
	return WebpagePreferencesClass.Alloc()
}

func (wc _WebpagePreferencesClass) New() WebpagePreferences {
	rv := objc.CallMethod[WebpagePreferences](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWebpagePreferences() WebpagePreferences {
	return WebpagePreferencesClass.New()
}

func WebpagePreferences_New() WebpagePreferences {
	return WebpagePreferencesClass.New()
}

func (w_ WebpagePreferences) Init() WebpagePreferences {
	rv := objc.CallMethod[WebpagePreferences](w_, objc.GetSelector("init"))
	return rv
}

func WebpagePreferences_Init() WebpagePreferences {
	return WebpagePreferencesClass.Alloc().Init()
}

func (w_ WebpagePreferences) AllowsContentJavaScript() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsContentJavaScript"))
	return rv
}

func (w_ WebpagePreferences) SetAllowsContentJavaScript(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsContentJavaScript:"), value)
}

func (w_ WebpagePreferences) PreferredContentMode() ContentMode {
	rv := objc.CallMethod[ContentMode](w_, objc.GetSelector("preferredContentMode"))
	return rv
}

func (w_ WebpagePreferences) SetPreferredContentMode(value ContentMode) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setPreferredContentMode:"), value)
}

func (w_ WebpagePreferences) IsLockdownModeEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isLockdownModeEnabled"))
	return rv
}

func (w_ WebpagePreferences) SetLockdownModeEnabled(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setLockdownModeEnabled:"), value)
}
