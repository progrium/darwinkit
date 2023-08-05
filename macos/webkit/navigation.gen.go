// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var NavigationClass = _NavigationClass{objc.GetClass("WKNavigation")}

type _NavigationClass struct {
	objc.Class
}

type INavigation interface {
	objc.IObject
	EffectiveContentMode() ContentMode
}

type Navigation struct {
	objc.Object
}

func MakeNavigation(ptr unsafe.Pointer) Navigation {
	return Navigation{
		Object: objc.MakeObject(ptr),
	}
}

func (nc _NavigationClass) Alloc() Navigation {
	rv := objc.CallMethod[Navigation](nc, objc.GetSelector("alloc"))
	return rv
}

func Navigation_Alloc() Navigation {
	return NavigationClass.Alloc()
}

func (nc _NavigationClass) New() Navigation {
	rv := objc.CallMethod[Navigation](nc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewNavigation() Navigation {
	return NavigationClass.New()
}

func Navigation_New() Navigation {
	return NavigationClass.New()
}

func (n_ Navigation) Init() Navigation {
	rv := objc.CallMethod[Navigation](n_, objc.GetSelector("init"))
	return rv
}

func Navigation_Init() Navigation {
	return NavigationClass.Alloc().Init()
}

func (n_ Navigation) EffectiveContentMode() ContentMode {
	rv := objc.CallMethod[ContentMode](n_, objc.GetSelector("effectiveContentMode"))
	return rv
}
