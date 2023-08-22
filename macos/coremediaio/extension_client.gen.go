// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionClient] class.
var ExtensionClientClass = _ExtensionClientClass{objc.GetClass("CMIOExtensionClient")}

type _ExtensionClientClass struct {
	objc.Class
}

// An interface definition for the [ExtensionClient] class.
type IExtensionClient interface {
	objc.IObject
	ClientID() foundation.UUID
}

// An object that represents a client of the extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionclient?language=objc
type ExtensionClient struct {
	objc.Object
}

func ExtensionClientFrom(ptr unsafe.Pointer) ExtensionClient {
	return ExtensionClient{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionClientClass) Alloc() ExtensionClient {
	rv := objc.Call[ExtensionClient](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionClient_Alloc() ExtensionClient {
	return ExtensionClientClass.Alloc()
}

func (ec _ExtensionClientClass) New() ExtensionClient {
	rv := objc.Call[ExtensionClient](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionClient() ExtensionClient {
	return ExtensionClientClass.New()
}

func (e_ ExtensionClient) Init() ExtensionClient {
	rv := objc.Call[ExtensionClient](e_, objc.Sel("init"))
	return rv
}

// A unique client identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionclient/3915852-clientid?language=objc
func (e_ ExtensionClient) ClientID() foundation.UUID {
	rv := objc.Call[foundation.UUID](e_, objc.Sel("clientID"))
	return rv
}
