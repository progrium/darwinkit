// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebpagePreferences] class.
var WebpagePreferencesClass = _WebpagePreferencesClass{objc.GetClass("WKWebpagePreferences")}

type _WebpagePreferencesClass struct {
	objc.Class
}

// An interface definition for the [WebpagePreferences] class.
type IWebpagePreferences interface {
	objc.IObject
	PreferredContentMode() ContentMode
	SetPreferredContentMode(value ContentMode)
	AllowsContentJavaScript() bool
	SetAllowsContentJavaScript(value bool)
}

// An object that specifies the behaviors to use when loading and rendering page content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences?language=objc
type WebpagePreferences struct {
	objc.Object
}

func WebpagePreferencesFrom(ptr unsafe.Pointer) WebpagePreferences {
	return WebpagePreferences{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebpagePreferencesClass) Alloc() WebpagePreferences {
	rv := objc.Call[WebpagePreferences](wc, objc.Sel("alloc"))
	return rv
}

func WebpagePreferences_Alloc() WebpagePreferences {
	return WebpagePreferencesClass.Alloc()
}

func (wc _WebpagePreferencesClass) New() WebpagePreferences {
	rv := objc.Call[WebpagePreferences](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebpagePreferences() WebpagePreferences {
	return WebpagePreferencesClass.New()
}

func (w_ WebpagePreferences) Init() WebpagePreferences {
	rv := objc.Call[WebpagePreferences](w_, objc.Sel("init"))
	return rv
}

// The content mode for the web view to use when it loads and renders a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/3194426-preferredcontentmode?language=objc
func (w_ WebpagePreferences) PreferredContentMode() ContentMode {
	rv := objc.Call[ContentMode](w_, objc.Sel("preferredContentMode"))
	return rv
}

// The content mode for the web view to use when it loads and renders a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/3194426-preferredcontentmode?language=objc
func (w_ WebpagePreferences) SetPreferredContentMode(value ContentMode) {
	objc.Call[objc.Void](w_, objc.Sel("setPreferredContentMode:"), value)
}

// A Boolean value that indicates whether JavaScript from web content is allowed to run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/3552422-allowscontentjavascript?language=objc
func (w_ WebpagePreferences) AllowsContentJavaScript() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsContentJavaScript"))
	return rv
}

// A Boolean value that indicates whether JavaScript from web content is allowed to run. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/3552422-allowscontentjavascript?language=objc
func (w_ WebpagePreferences) SetAllowsContentJavaScript(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsContentJavaScript:"), value)
}
