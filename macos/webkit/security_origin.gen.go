// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SecurityOrigin] class.
var SecurityOriginClass = _SecurityOriginClass{objc.GetClass("WKSecurityOrigin")}

type _SecurityOriginClass struct {
	objc.Class
}

// An interface definition for the [SecurityOrigin] class.
type ISecurityOrigin interface {
	objc.IObject
	Host() string
	Port() int
	Protocol() string
}

// An object that identifies the origin of a particular resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin?language=objc
type SecurityOrigin struct {
	objc.Object
}

func SecurityOriginFrom(ptr unsafe.Pointer) SecurityOrigin {
	return SecurityOrigin{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SecurityOriginClass) Alloc() SecurityOrigin {
	rv := objc.Call[SecurityOrigin](sc, objc.Sel("alloc"))
	return rv
}

func SecurityOrigin_Alloc() SecurityOrigin {
	return SecurityOriginClass.Alloc()
}

func (sc _SecurityOriginClass) New() SecurityOrigin {
	rv := objc.Call[SecurityOrigin](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSecurityOrigin() SecurityOrigin {
	return SecurityOriginClass.New()
}

func (s_ SecurityOrigin) Init() SecurityOrigin {
	rv := objc.Call[SecurityOrigin](s_, objc.Sel("init"))
	return rv
}

// The security origin’s host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/1536794-host?language=objc
func (s_ SecurityOrigin) Host() string {
	rv := objc.Call[string](s_, objc.Sel("host"))
	return rv
}

// The security origin's port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/1536403-port?language=objc
func (s_ SecurityOrigin) Port() int {
	rv := objc.Call[int](s_, objc.Sel("port"))
	return rv
}

// The security origin's protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wksecurityorigin/1537470-protocol?language=objc
func (s_ SecurityOrigin) Protocol() string {
	rv := objc.Call[string](s_, objc.Sel("protocol"))
	return rv
}
