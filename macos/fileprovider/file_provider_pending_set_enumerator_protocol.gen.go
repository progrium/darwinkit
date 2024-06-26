// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol for enumerating pending items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator?language=objc
type PFileProviderPendingSetEnumerator interface {
	// optional
	DomainVersion() FileProviderDomainVersion
	HasDomainVersion() bool

	// optional
	RefreshInterval() foundation.TimeInterval
	HasRefreshInterval() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderPendingSetEnumerator = (*FileProviderPendingSetEnumeratorObject)(nil)

// A concrete type for the [PFileProviderPendingSetEnumerator] protocol.
type FileProviderPendingSetEnumeratorObject struct {
	objc.Object
}

func (f_ FileProviderPendingSetEnumeratorObject) HasDomainVersion() bool {
	return f_.RespondsToSelector(objc.Sel("domainVersion"))
}

// The domain version when the system last refreshed the pending set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator/3727815-domainversion?language=objc
func (f_ FileProviderPendingSetEnumeratorObject) DomainVersion() FileProviderDomainVersion {
	rv := objc.Call[FileProviderDomainVersion](f_, objc.Sel("domainVersion"))
	return rv
}

func (f_ FileProviderPendingSetEnumeratorObject) HasRefreshInterval() bool {
	return f_.RespondsToSelector(objc.Sel("refreshInterval"))
}

// The amount of time, in seconds, between updates to the pending set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderpendingsetenumerator/3727816-refreshinterval?language=objc
func (f_ FileProviderPendingSetEnumeratorObject) RefreshInterval() foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](f_, objc.Sel("refreshInterval"))
	return rv
}
