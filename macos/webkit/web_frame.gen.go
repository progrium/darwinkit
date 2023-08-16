// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [WebFrame] class.
var WebFrameClass = _WebFrameClass{objc.GetClass("WebFrame")}

type _WebFrameClass struct {
	objc.Class
}

// An interface definition for the [WebFrame] class.
type IWebFrame interface {
	objc.IObject
}

// A WebFrame object encapsulates the data displayed in a WebFrameView object. There is one WebFrame object per frame displayed in a WebView. An entire webpage is represented by a hierarchy of WebFrame objects in which the root object is called the main frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webframe?language=objc
type WebFrame struct {
	objc.Object
}

func WebFrameFrom(ptr unsafe.Pointer) WebFrame {
	return WebFrame{
		Object: objc.ObjectFrom(ptr),
	}
}

func (wc _WebFrameClass) Alloc() WebFrame {
	rv := objc.Call[WebFrame](wc, objc.Sel("alloc"))
	return rv
}

func WebFrame_Alloc() WebFrame {
	return WebFrameClass.Alloc()
}

func (wc _WebFrameClass) New() WebFrame {
	rv := objc.Call[WebFrame](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWebFrame() WebFrame {
	return WebFrameClass.New()
}

func (w_ WebFrame) Init() WebFrame {
	rv := objc.Call[WebFrame](w_, objc.Sel("init"))
	return rv
}
