// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Navigation] class.
var NavigationClass = _NavigationClass{objc.GetClass("WKNavigation")}

type _NavigationClass struct {
	objc.Class
}

// An interface definition for the [Navigation] class.
type INavigation interface {
	objc.IObject
	EffectiveContentMode() ContentMode
}

// An object that tracks the loading progress of a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigation?language=objc
type Navigation struct {
	objc.Object
}

func NavigationFrom(ptr unsafe.Pointer) Navigation {
	return Navigation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NavigationClass) Alloc() Navigation {
	rv := objc.Call[Navigation](nc, objc.Sel("alloc"))
	return rv
}

func Navigation_Alloc() Navigation {
	return NavigationClass.Alloc()
}

func (nc _NavigationClass) New() Navigation {
	rv := objc.Call[Navigation](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNavigation() Navigation {
	return NavigationClass.New()
}

func (n_ Navigation) Init() Navigation {
	rv := objc.Call[Navigation](n_, objc.Sel("init"))
	return rv
}

// The content mode WebKit uses to load the webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigation/3229061-effectivecontentmode?language=objc
func (n_ Navigation) EffectiveContentMode() ContentMode {
	rv := objc.Call[ContentMode](n_, objc.Sel("effectiveContentMode"))
	return rv
}
