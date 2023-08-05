// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ContentWorldClass = _ContentWorldClass{objc.GetClass("WKContentWorld")}

type _ContentWorldClass struct {
	objc.Class
}

type IContentWorld interface {
	objc.IObject
	Name() string
}

type ContentWorld struct {
	objc.Object
}

func MakeContentWorld(ptr unsafe.Pointer) ContentWorld {
	return ContentWorld{
		Object: objc.MakeObject(ptr),
	}
}

func (cc _ContentWorldClass) Alloc() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.GetSelector("alloc"))
	return rv
}

func ContentWorld_Alloc() ContentWorld {
	return ContentWorldClass.Alloc()
}

func (cc _ContentWorldClass) New() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewContentWorld() ContentWorld {
	return ContentWorldClass.New()
}

func ContentWorld_New() ContentWorld {
	return ContentWorldClass.New()
}

func (c_ ContentWorld) Init() ContentWorld {
	rv := objc.CallMethod[ContentWorld](c_, objc.GetSelector("init"))
	return rv
}

func ContentWorld_Init() ContentWorld {
	return ContentWorldClass.Alloc().Init()
}

func (cc _ContentWorldClass) WorldWithName(name string) ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.GetSelector("worldWithName:"), name)
	return rv
}

func ContentWorld_WorldWithName(name string) ContentWorld {
	return ContentWorldClass.WorldWithName(name)
}

func (cc _ContentWorldClass) DefaultClientWorld() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.GetSelector("defaultClientWorld"))
	return rv
}

func ContentWorld_DefaultClientWorld() ContentWorld {
	return ContentWorldClass.DefaultClientWorld()
}

func (cc _ContentWorldClass) PageWorld() ContentWorld {
	rv := objc.CallMethod[ContentWorld](cc, objc.GetSelector("pageWorld"))
	return rv
}

func ContentWorld_PageWorld() ContentWorld {
	return ContentWorldClass.PageWorld()
}

func (c_ ContentWorld) Name() string {
	rv := objc.CallMethod[string](c_, objc.GetSelector("name"))
	return rv
}
