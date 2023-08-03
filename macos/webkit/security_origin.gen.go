// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var SecurityOriginClass = _SecurityOriginClass{objc.GetClass("WKSecurityOrigin")}

type _SecurityOriginClass struct {
	objc.Class
}

type ISecurityOrigin interface {
	objc.IObject
	Host() string
	Port() int
	Protocol() string
}

type SecurityOrigin struct {
	objc.Object
}

func MakeSecurityOrigin(ptr unsafe.Pointer) SecurityOrigin {
	return SecurityOrigin{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _SecurityOriginClass) Alloc() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](sc, objc.GetSelector("alloc"))
	return rv
}

func SecurityOrigin_Alloc() SecurityOrigin {
	return SecurityOriginClass.Alloc()
}

func (sc _SecurityOriginClass) New() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSecurityOrigin() SecurityOrigin {
	return SecurityOriginClass.New()
}

func SecurityOrigin_New() SecurityOrigin {
	return SecurityOriginClass.New()
}

func (s_ SecurityOrigin) Init() SecurityOrigin {
	rv := objc.CallMethod[SecurityOrigin](s_, objc.GetSelector("init"))
	return rv
}

func SecurityOrigin_Init() SecurityOrigin {
	return SecurityOriginClass.Alloc().Init()
}

func (s_ SecurityOrigin) Host() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("host"))
	return rv
}

func (s_ SecurityOrigin) Port() int {
	rv := objc.CallMethod[int](s_, objc.GetSelector("port"))
	return rv
}

func (s_ SecurityOrigin) Protocol() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("protocol"))
	return rv
}
