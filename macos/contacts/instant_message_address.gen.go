// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InstantMessageAddress] class.
var InstantMessageAddressClass = _InstantMessageAddressClass{objc.GetClass("CNInstantMessageAddress")}

type _InstantMessageAddressClass struct {
	objc.Class
}

// An interface definition for the [InstantMessageAddress] class.
type IInstantMessageAddress interface {
	objc.IObject
	Service() string
	Username() string
}

// An immutable object representing an instant message address for the contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress?language=objc
type InstantMessageAddress struct {
	objc.Object
}

func InstantMessageAddressFrom(ptr unsafe.Pointer) InstantMessageAddress {
	return InstantMessageAddress{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ InstantMessageAddress) InitWithUsernameService(username string, service string) InstantMessageAddress {
	rv := objc.Call[InstantMessageAddress](i_, objc.Sel("initWithUsername:service:"), username, service)
	return rv
}

// Returns a CNInstantMessageAddress object initialized with the specified user name and service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1402808-initwithusername?language=objc
func InstantMessageAddress_InitWithUsernameService(username string, service string) InstantMessageAddress {
	return InstantMessageAddressClass.Alloc().InitWithUsernameService(username, service)
}

func (ic _InstantMessageAddressClass) Alloc() InstantMessageAddress {
	rv := objc.Call[InstantMessageAddress](ic, objc.Sel("alloc"))
	return rv
}

func InstantMessageAddress_Alloc() InstantMessageAddress {
	return InstantMessageAddressClass.Alloc()
}

func (ic _InstantMessageAddressClass) New() InstantMessageAddress {
	rv := objc.Call[InstantMessageAddress](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInstantMessageAddress() InstantMessageAddress {
	return InstantMessageAddressClass.New()
}

func (i_ InstantMessageAddress) Init() InstantMessageAddress {
	rv := objc.Call[InstantMessageAddress](i_, objc.Sel("init"))
	return rv
}

// Returns a string containing the localized property name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403174-localizedstringforkey?language=objc
func (ic _InstantMessageAddressClass) LocalizedStringForKey(key string) string {
	rv := objc.Call[string](ic, objc.Sel("localizedStringForKey:"), key)
	return rv
}

// Returns a string containing the localized property name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403174-localizedstringforkey?language=objc
func InstantMessageAddress_LocalizedStringForKey(key string) string {
	return InstantMessageAddressClass.LocalizedStringForKey(key)
}

// Returns a string containing the localized name of the specified service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403367-localizedstringforservice?language=objc
func (ic _InstantMessageAddressClass) LocalizedStringForService(service string) string {
	rv := objc.Call[string](ic, objc.Sel("localizedStringForService:"), service)
	return rv
}

// Returns a string containing the localized name of the specified service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403367-localizedstringforservice?language=objc
func InstantMessageAddress_LocalizedStringForService(service string) string {
	return InstantMessageAddressClass.LocalizedStringForService(service)
}

// The name of the instant message address service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403423-service?language=objc
func (i_ InstantMessageAddress) Service() string {
	rv := objc.Call[string](i_, objc.Sel("service"))
	return rv
}

// The user name for instant message service address. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cninstantmessageaddress/1403129-username?language=objc
func (i_ InstantMessageAddress) Username() string {
	rv := objc.Call[string](i_, objc.Sel("username"))
	return rv
}
