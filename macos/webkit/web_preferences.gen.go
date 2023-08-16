// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebPreferences] class.
var WebPreferencesClass = _WebPreferencesClass{objc.GetClass("WebPreferences")}

type _WebPreferencesClass struct {
	objc.Class
}

// An interface definition for the [WebPreferences] class.
type IWebPreferences interface {
	objc.IObject
}

// WebPreferences encapsulates the preferences you can change per WebView object. These preferences include font, text encoding, and image settings. Normally a WebView object uses the standard preferences returned by the standardPreferences class method. However, you can modify the preferences for individual WebView instances too. Use the preferencesIdentifier WebView method to change a WebView object’s preferences, or to share preferences between WebView objects. Use the autosaves method to specify if the preferences object should be automatically saved to the user defaults database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webpreferences?language=objc
type WebPreferences struct {
	objc.Object
}

func WebPreferencesFrom(ptr unsafe.Pointer) WebPreferences {
	return WebPreferences{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebPreferencesClass) Alloc() WebPreferences {
	rv := objc.Call[WebPreferences](wc, objc.Sel("alloc"))
	return rv
}

func WebPreferences_Alloc() WebPreferences {
	return WebPreferencesClass.Alloc()
}

func (wc _WebPreferencesClass) New() WebPreferences {
	rv := objc.Call[WebPreferences](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebPreferences() WebPreferences {
	return WebPreferencesClass.New()
}

func (w_ WebPreferences) Init() WebPreferences {
	rv := objc.Call[WebPreferences](w_, objc.Sel("init"))
	return rv
}
