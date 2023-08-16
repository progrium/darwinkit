// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebUndefined] class.
var WebUndefinedClass = _WebUndefinedClass{objc.GetClass("WebUndefined")}

type _WebUndefinedClass struct {
	objc.Class
}

// An interface definition for the [WebUndefined] class.
type IWebUndefined interface {
	objc.IObject
}

// WebUndefined objects are simply used to represent the JavaScript “undefined” value in methods when bridging between JavaScript and Objective-C. For example, if you invoke a JavaScript function that returns the JavaScript “undefined” value, then a WebUndefined object is returned to the Objective-C calling context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webundefined?language=objc
type WebUndefined struct {
	objc.Object
}

func WebUndefinedFrom(ptr unsafe.Pointer) WebUndefined {
	return WebUndefined{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebUndefinedClass) Alloc() WebUndefined {
	rv := objc.Call[WebUndefined](wc, objc.Sel("alloc"))
	return rv
}

func WebUndefined_Alloc() WebUndefined {
	return WebUndefinedClass.Alloc()
}

func (wc _WebUndefinedClass) New() WebUndefined {
	rv := objc.Call[WebUndefined](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebUndefined() WebUndefined {
	return WebUndefinedClass.New()
}

func (w_ WebUndefined) Init() WebUndefined {
	rv := objc.Call[WebUndefined](w_, objc.Sel("init"))
	return rv
}
