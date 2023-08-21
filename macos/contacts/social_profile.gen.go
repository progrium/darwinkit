// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SocialProfile] class.
var SocialProfileClass = _SocialProfileClass{objc.GetClass("CNSocialProfile")}

type _SocialProfileClass struct {
	objc.Class
}

// An interface definition for the [SocialProfile] class.
type ISocialProfile interface {
	objc.IObject
	UrlString() string
	Service() string
	Username() string
	UserIdentifier() string
}

// An immutable object that represents one of the user's social profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile?language=objc
type SocialProfile struct {
	objc.Object
}

func SocialProfileFrom(ptr unsafe.Pointer) SocialProfile {
	return SocialProfile{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SocialProfile) InitWithUrlStringUsernameUserIdentifierService(urlString string, username string, userIdentifier string, service string) SocialProfile {
	rv := objc.Call[SocialProfile](s_, objc.Sel("initWithUrlString:username:userIdentifier:service:"), urlString, username, userIdentifier, service)
	return rv
}

// Initializes a new social profile object with the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402989-initwithurlstring?language=objc
func NewSocialProfileWithUrlStringUsernameUserIdentifierService(urlString string, username string, userIdentifier string, service string) SocialProfile {
	instance := SocialProfileClass.Alloc().InitWithUrlStringUsernameUserIdentifierService(urlString, username, userIdentifier, service)
	instance.Autorelease()
	return instance
}

func (sc _SocialProfileClass) Alloc() SocialProfile {
	rv := objc.Call[SocialProfile](sc, objc.Sel("alloc"))
	return rv
}

func SocialProfile_Alloc() SocialProfile {
	return SocialProfileClass.Alloc()
}

func (sc _SocialProfileClass) New() SocialProfile {
	rv := objc.Call[SocialProfile](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSocialProfile() SocialProfile {
	return SocialProfileClass.New()
}

func (s_ SocialProfile) Init() SocialProfile {
	rv := objc.Call[SocialProfile](s_, objc.Sel("init"))
	return rv
}

// Returns the localized name of the property for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402857-localizedstringforkey?language=objc
func (sc _SocialProfileClass) LocalizedStringForKey(key string) string {
	rv := objc.Call[string](sc, objc.Sel("localizedStringForKey:"), key)
	return rv
}

// Returns the localized name of the property for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402857-localizedstringforkey?language=objc
func SocialProfile_LocalizedStringForKey(key string) string {
	return SocialProfileClass.LocalizedStringForKey(key)
}

// Returns the localized name of the specified service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402819-localizedstringforservice?language=objc
func (sc _SocialProfileClass) LocalizedStringForService(service string) string {
	rv := objc.Call[string](sc, objc.Sel("localizedStringForService:"), service)
	return rv
}

// Returns the localized name of the specified service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402819-localizedstringforservice?language=objc
func SocialProfile_LocalizedStringForService(service string) string {
	return SocialProfileClass.LocalizedStringForService(service)
}

// The URL associated with the social profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1403044-urlstring?language=objc
func (s_ SocialProfile) UrlString() string {
	rv := objc.Call[string](s_, objc.Sel("urlString"))
	return rv
}

// The social profile’s service name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1402947-service?language=objc
func (s_ SocialProfile) Service() string {
	rv := objc.Call[string](s_, objc.Sel("service"))
	return rv
}

// The user name for the social profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1403351-username?language=objc
func (s_ SocialProfile) Username() string {
	rv := objc.Call[string](s_, objc.Sel("username"))
	return rv
}

// The service’s user identifier associated with the social profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsocialprofile/1403383-useridentifier?language=objc
func (s_ SocialProfile) UserIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("userIdentifier"))
	return rv
}
