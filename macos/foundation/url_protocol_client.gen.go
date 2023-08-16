// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface used by NSURLProtocol subclasses to communicate with the URL Loading System. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocolclient?language=objc
type PURLProtocolClient interface {
	// optional
	URLProtocolDidFailWithError(protocol URLProtocol, error Error)
	HasURLProtocolDidFailWithError() bool

	// optional
	URLProtocolDidFinishLoading(protocol URLProtocol)
	HasURLProtocolDidFinishLoading() bool
}

// A concrete type wrapper for the [PURLProtocolClient] protocol.
type URLProtocolClientWrapper struct {
	objc.Object
}

func (u_ URLProtocolClientWrapper) HasURLProtocolDidFailWithError() bool {
	return u_.RespondsToSelector(objc.Sel("URLProtocol:didFailWithError:"))
}

// Tells the client that the load request failed due to an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocolclient/1413131-urlprotocol?language=objc
func (u_ URLProtocolClientWrapper) URLProtocolDidFailWithError(protocol IURLProtocol, error IError) {
	objc.Call[objc.Void](u_, objc.Sel("URLProtocol:didFailWithError:"), objc.Ptr(protocol), objc.Ptr(error))
}

func (u_ URLProtocolClientWrapper) HasURLProtocolDidFinishLoading() bool {
	return u_.RespondsToSelector(objc.Sel("URLProtocolDidFinishLoading:"))
}

// Tells the client that the protocol implementation has finished loading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlprotocolclient/1411543-urlprotocoldidfinishloading?language=objc
func (u_ URLProtocolClientWrapper) URLProtocolDidFinishLoading(protocol IURLProtocol) {
	objc.Call[objc.Void](u_, objc.Sel("URLProtocolDidFinishLoading:"), objc.Ptr(protocol))
}
