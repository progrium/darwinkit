// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebScriptObject] class.
var WebScriptObjectClass = _WebScriptObjectClass{objc.GetClass("WebScriptObject")}

type _WebScriptObjectClass struct {
	objc.Class
}

// An interface definition for the [WebScriptObject] class.
type IWebScriptObject interface {
	objc.IObject
}

// A WebScriptObject object is an Objective-C wrapper for a scripting object passed to your application from the scripting environment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webscriptobject?language=objc
type WebScriptObject struct {
	objc.Object
}

func WebScriptObjectFrom(ptr unsafe.Pointer) WebScriptObject {
	return WebScriptObject{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebScriptObjectClass) Alloc() WebScriptObject {
	rv := objc.Call[WebScriptObject](wc, objc.Sel("alloc"))
	return rv
}

func WebScriptObject_Alloc() WebScriptObject {
	return WebScriptObjectClass.Alloc()
}

func (wc _WebScriptObjectClass) New() WebScriptObject {
	rv := objc.Call[WebScriptObject](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebScriptObject() WebScriptObject {
	return WebScriptObjectClass.New()
}

func (w_ WebScriptObject) Init() WebScriptObject {
	rv := objc.Call[WebScriptObject](w_, objc.Sel("init"))
	return rv
}
