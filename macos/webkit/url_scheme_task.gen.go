// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that WebKit uses to request custom resources from your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask?language=objc
type PURLSchemeTask interface {
	// optional
	DidFailWithError(error foundation.Error)
	HasDidFailWithError() bool

	// optional
	DidFinish()
	HasDidFinish() bool

	// optional
	DidReceiveData(data []byte)
	HasDidReceiveData() bool

	// optional
	DidReceiveResponse(response foundation.URLResponse)
	HasDidReceiveResponse() bool

	// optional
	Request() foundation.IURLRequest
	HasRequest() bool
}

// A concrete type wrapper for the [PURLSchemeTask] protocol.
type URLSchemeTaskWrapper struct {
	objc.Object
}

func (u_ URLSchemeTaskWrapper) HasDidFailWithError() bool {
	return u_.RespondsToSelector(objc.Sel("didFailWithError:"))
}

// Completes the task and reports the specified error back to WebKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask/2890841-didfailwitherror?language=objc
func (u_ URLSchemeTaskWrapper) DidFailWithError(error foundation.IError) {
	objc.Call[objc.Void](u_, objc.Sel("didFailWithError:"), objc.Ptr(error))
}

func (u_ URLSchemeTaskWrapper) HasDidFinish() bool {
	return u_.RespondsToSelector(objc.Sel("didFinish"))
}

// Signals the successful completion of the task. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask/2890837-didfinish?language=objc
func (u_ URLSchemeTaskWrapper) DidFinish() {
	objc.Call[objc.Void](u_, objc.Sel("didFinish"))
}

func (u_ URLSchemeTaskWrapper) HasDidReceiveData() bool {
	return u_.RespondsToSelector(objc.Sel("didReceiveData:"))
}

// Sends some or all of the resource data to WebKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask/2890836-didreceivedata?language=objc
func (u_ URLSchemeTaskWrapper) DidReceiveData(data []byte) {
	objc.Call[objc.Void](u_, objc.Sel("didReceiveData:"), data)
}

func (u_ URLSchemeTaskWrapper) HasDidReceiveResponse() bool {
	return u_.RespondsToSelector(objc.Sel("didReceiveResponse:"))
}

// Returns a URL response to WebKit with information about the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask/2890839-didreceiveresponse?language=objc
func (u_ URLSchemeTaskWrapper) DidReceiveResponse(response foundation.IURLResponse) {
	objc.Call[objc.Void](u_, objc.Sel("didReceiveResponse:"), objc.Ptr(response))
}

func (u_ URLSchemeTaskWrapper) HasRequest() bool {
	return u_.RespondsToSelector(objc.Sel("request"))
}

// Information about the resource to load. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemetask/2890840-request?language=objc
func (u_ URLSchemeTaskWrapper) Request() foundation.URLRequest {
	rv := objc.Call[foundation.URLRequest](u_, objc.Sel("request"))
	return rv
}
