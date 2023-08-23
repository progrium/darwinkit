// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to handle resource-loading requests coming from a URL asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate?language=objc
type PAssetResourceLoaderDelegate interface {
	// optional
	ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
	HasResourceLoaderDidCancelAuthenticationChallenge() bool
}

// A delegate implementation builder for the [PAssetResourceLoaderDelegate] protocol.
type AssetResourceLoaderDelegate struct {
	_ResourceLoaderDidCancelAuthenticationChallenge func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
}

func (di *AssetResourceLoaderDelegate) HasResourceLoaderDidCancelAuthenticationChallenge() bool {
	return di._ResourceLoaderDidCancelAuthenticationChallenge != nil
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderDidCancelAuthenticationChallenge(f func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)) {
	di._ResourceLoaderDidCancelAuthenticationChallenge = f
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) {
	di._ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader, authenticationChallenge)
}

// A concrete type wrapper for the [PAssetResourceLoaderDelegate] protocol.
type AssetResourceLoaderDelegateWrapper struct {
	objc.Object
}

func (a_ AssetResourceLoaderDelegateWrapper) HasResourceLoaderDidCancelAuthenticationChallenge() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:didCancelAuthenticationChallenge:"))
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateWrapper) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader IAssetResourceLoader, authenticationChallenge foundation.IURLAuthenticationChallenge) {
	objc.Call[objc.Void](a_, objc.Sel("resourceLoader:didCancelAuthenticationChallenge:"), objc.Ptr(resourceLoader), objc.Ptr(authenticationChallenge))
}
