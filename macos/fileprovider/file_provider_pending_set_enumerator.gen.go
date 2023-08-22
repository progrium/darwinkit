// AUTO-GENERATED CODE, DO NOT MODIFY

package fileprovider

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol for enumerating pending items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator?language=objc
type PFileProviderPendingSetEnumerator interface {
	// optional
	DomainVersion() IFileProviderDomainVersion
	HasDomainVersion() bool

	// optional
	RefreshInterval() foundation.TimeInterval
	HasRefreshInterval() bool
}

// A concrete type wrapper for the [PFileProviderPendingSetEnumerator] protocol.
type FileProviderPendingSetEnumeratorWrapper struct {
	objc.Object
}

func (f_ FileProviderPendingSetEnumeratorWrapper) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain version when the system last refreshed the pending set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator/3727815-domainversion?language=objc
func (f_ FileProviderPendingSetEnumeratorWrapper) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

func (f_ FileProviderPendingSetEnumeratorWrapper) HasRefreshInterval() bool {
	return f_.RespondsToSelector(objc.Sel("refreshInterval"))
}

// The amount of time, in seconds, between updates to the pending set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator/3727816-refreshinterval?language=objc
func (f_ FileProviderPendingSetEnumeratorWrapper) RefreshInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](f_, objc.Sel("refreshInterval"))
	return rv
}
