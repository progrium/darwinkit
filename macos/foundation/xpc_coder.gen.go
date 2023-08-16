// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XPCCoder] class.
var XPCCoderClass = _XPCCoderClass{objc.GetClass("NSXPCCoder")}

type _XPCCoderClass struct {
	objc.Class
}

// An interface definition for the [XPCCoder] class.
type IXPCCoder interface {
	ICoder
	Connection() XPCConnection
	UserInfo() objc.Object
	SetUserInfo(value objc.IObject)
}

// A coder that encodes and decodes objects that your app sends over an XPC connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpccoder?language=objc
type XPCCoder struct {
	Coder
}

func XPCCoderFrom(ptr unsafe.Pointer) XPCCoder {
	return XPCCoder{
		Coder: CoderFrom(ptr),
	}
}

func (xc _XPCCoderClass) Alloc() XPCCoder {
	rv := objc.Call[XPCCoder](xc, objc.Sel("alloc"))
	return rv
}

func XPCCoder_Alloc() XPCCoder {
	return XPCCoderClass.Alloc()
}

func (xc _XPCCoderClass) New() XPCCoder {
	rv := objc.Call[XPCCoder](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXPCCoder() XPCCoder {
	return XPCCoderClass.New()
}

func (x_ XPCCoder) Init() XPCCoder {
	rv := objc.Call[XPCCoder](x_, objc.Sel("init"))
	return rv
}

// The connection currently performing encoding or decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpccoder/3172580-connection?language=objc
func (x_ XPCCoder) Connection() XPCConnection {
	rv := objc.Call[XPCConnection](x_, objc.Sel("connection"))
	return rv
}

// An optional user information object associated with the coder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpccoder/3172581-userinfo?language=objc
func (x_ XPCCoder) UserInfo() objc.Object {
	rv := objc.Call[objc.Object](x_, objc.Sel("userInfo"))
	return rv
}

// An optional user information object associated with the coder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpccoder/3172581-userinfo?language=objc
func (x_ XPCCoder) SetUserInfo(value objc.IObject) {
	objc.Call[objc.Void](x_, objc.Sel("setUserInfo:"), value)
}
