// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An object that contains global state data about the domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate?language=objc
type PFileProviderDomainState interface {
	// optional
	UserInfo() foundation.Dictionary
	HasUserInfo() bool

	// optional
	DomainVersion() IFileProviderDomainVersion
	HasDomainVersion() bool
}

// A concrete type wrapper for the [PFileProviderDomainState] protocol.
type FileProviderDomainStateWrapper struct {
	objc.Object
}

func (f_ FileProviderDomainStateWrapper) HasUserInfo() bool {
	return f_.RespondsToSelector(objc.Sel("userInfo"))
}

// Global state information about the current domain version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/3727820-userinfo?language=objc
func (f_ FileProviderDomainStateWrapper) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](f_, objc.Sel("userInfo"))
	return rv
}

func (f_ FileProviderDomainStateWrapper) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// An opaque object that uniquely identifies the domain’s version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderdomainstate/3727819-domainversion?language=objc
func (f_ FileProviderDomainStateWrapper) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}
