// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// Web view resource load delegates implement this protocol to be notified on the progress of loading individual resources. Note that there can be hundreds of resources, such as images and other media, per page. So, if you just want to get page loading status see the WebFrameLoadDelegate protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webresourceloaddelegate?language=objc
type PWebResourceLoadDelegate interface {
}

// A delegate implementation builder for the [PWebResourceLoadDelegate] protocol.
type WebResourceLoadDelegate struct {
}

// A concrete type wrapper for the [PWebResourceLoadDelegate] protocol.
type WebResourceLoadDelegateWrapper struct {
	objc.Object
}
