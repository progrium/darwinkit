// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLProtectionSpace] class.
var URLProtectionSpaceClass = _URLProtectionSpaceClass{objc.GetClass("NSURLProtectionSpace")}

type _URLProtectionSpaceClass struct {
	objc.Class
}

// An interface definition for the [URLProtectionSpace] class.
type IURLProtectionSpace interface {
	objc.IObject
	ProxyType() string
	Host() string
	Realm() string
	IsProxy() bool
	Port() int
	Protocol() string
	DistinguishedNames() [][]byte
	AuthenticationMethod() string
	ReceivesCredentialSecurely() bool
}

// A server or an area on a server, commonly referred to as a realm, that requires authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace?language=objc
type URLProtectionSpace struct {
	objc.Object
}

func URLProtectionSpaceFrom(ptr unsafe.Pointer) URLProtectionSpace {
	return URLProtectionSpace{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLProtectionSpace) InitWithProxyHostPortTypeRealmAuthenticationMethod(host string, port int, type_ string, realm string, authenticationMethod string) URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](u_, objc.Sel("initWithProxyHost:port:type:realm:authenticationMethod:"), host, port, type_, realm, authenticationMethod)
	return rv
}

// Creates a protection space object representing a proxy server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1417998-initwithproxyhost?language=objc
func NewURLProtectionSpaceWithProxyHostPortTypeRealmAuthenticationMethod(host string, port int, type_ string, realm string, authenticationMethod string) URLProtectionSpace {
	instance := URLProtectionSpaceClass.Alloc().InitWithProxyHostPortTypeRealmAuthenticationMethod(host, port, type_, realm, authenticationMethod)
	instance.Autorelease()
	return instance
}

func (u_ URLProtectionSpace) InitWithHostPortProtocolRealmAuthenticationMethod(host string, port int, protocol string, realm string, authenticationMethod string) URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](u_, objc.Sel("initWithHost:port:protocol:realm:authenticationMethod:"), host, port, protocol, realm, authenticationMethod)
	return rv
}

// Creates a protection space object from the given host, port, protocol, realm, and authentication method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1414165-initwithhost?language=objc
func NewURLProtectionSpaceWithHostPortProtocolRealmAuthenticationMethod(host string, port int, protocol string, realm string, authenticationMethod string) URLProtectionSpace {
	instance := URLProtectionSpaceClass.Alloc().InitWithHostPortProtocolRealmAuthenticationMethod(host, port, protocol, realm, authenticationMethod)
	instance.Autorelease()
	return instance
}

func (uc _URLProtectionSpaceClass) Alloc() URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](uc, objc.Sel("alloc"))
	return rv
}

func URLProtectionSpace_Alloc() URLProtectionSpace {
	return URLProtectionSpaceClass.Alloc()
}

func (uc _URLProtectionSpaceClass) New() URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLProtectionSpace() URLProtectionSpace {
	return URLProtectionSpaceClass.New()
}

func (u_ URLProtectionSpace) Init() URLProtectionSpace {
	rv := objc.Call[URLProtectionSpace](u_, objc.Sel("init"))
	return rv
}

// The receiver's proxy type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1411924-proxytype?language=objc
func (u_ URLProtectionSpace) ProxyType() string {
	rv := objc.Call[string](u_, objc.Sel("proxyType"))
	return rv
}

// The receiver’s host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1418205-host?language=objc
func (u_ URLProtectionSpace) Host() string {
	rv := objc.Call[string](u_, objc.Sel("host"))
	return rv
}

// The receiver’s authentication realm [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1416007-realm?language=objc
func (u_ URLProtectionSpace) Realm() string {
	rv := objc.Call[string](u_, objc.Sel("realm"))
	return rv
}

// A Boolean value that indicates whether the receiver represents a proxy server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1561656-isproxy?language=objc
func (u_ URLProtectionSpace) IsProxy() bool {
	rv := objc.Call[bool](u_, objc.Sel("isProxy"))
	return rv
}

// The receiver’s port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1408716-port?language=objc
func (u_ URLProtectionSpace) Port() int {
	rv := objc.Call[int](u_, objc.Sel("port"))
	return rv
}

// The receiver’s protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1411191-protocol?language=objc
func (u_ URLProtectionSpace) Protocol() string {
	rv := objc.Call[string](u_, objc.Sel("protocol"))
	return rv
}

// The acceptable certificate-issuing authorities for client certificate authentication. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1417061-distinguishednames?language=objc
func (u_ URLProtectionSpace) DistinguishedNames() [][]byte {
	rv := objc.Call[[][]byte](u_, objc.Sel("distinguishedNames"))
	return rv
}

// The authentication method used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1415028-authenticationmethod?language=objc
func (u_ URLProtectionSpace) AuthenticationMethod() string {
	rv := objc.Call[string](u_, objc.Sel("authenticationMethod"))
	return rv
}

// A Boolean value that indicates whether the credentials for the protection space can be sent securely. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotectionspace/1415176-receivescredentialsecurely?language=objc
func (u_ URLProtectionSpace) ReceivesCredentialSecurely() bool {
	rv := objc.Call[bool](u_, objc.Sel("receivesCredentialSecurely"))
	return rv
}
