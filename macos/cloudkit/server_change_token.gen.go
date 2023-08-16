// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ServerChangeToken] class.
var ServerChangeTokenClass = _ServerChangeTokenClass{objc.GetClass("CKServerChangeToken")}

type _ServerChangeTokenClass struct {
	objc.Class
}

// An interface definition for the [ServerChangeToken] class.
type IServerChangeToken interface {
	objc.IObject
}

// An opaque token that represents a specific point in a database’s history. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckserverchangetoken?language=objc
type ServerChangeToken struct {
	objc.Object
}

func ServerChangeTokenFrom(ptr unsafe.Pointer) ServerChangeToken {
	return ServerChangeToken{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ServerChangeTokenClass) Alloc() ServerChangeToken {
	rv := objc.Call[ServerChangeToken](sc, objc.Sel("alloc"))
	return rv
}

func ServerChangeToken_Alloc() ServerChangeToken {
	return ServerChangeTokenClass.Alloc()
}

func (sc _ServerChangeTokenClass) New() ServerChangeToken {
	rv := objc.Call[ServerChangeToken](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewServerChangeToken() ServerChangeToken {
	return ServerChangeTokenClass.New()
}

func (s_ ServerChangeToken) Init() ServerChangeToken {
	rv := objc.Call[ServerChangeToken](s_, objc.Sel("init"))
	return rv
}
