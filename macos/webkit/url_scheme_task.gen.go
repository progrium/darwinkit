// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IURLSchemeTask interface {
	// required
	DidReceiveResponse(response foundation.URLResponse)
	// required
	DidReceiveData(data []byte)
	// required
	DidFinish()
	// required
	DidFailWithError(error foundation.Error)
	ImplementsRequest() bool
	// optional
	Request() foundation.IURLRequest
}

type URLSchemeTaskWrapper struct {
	objc.Object
}

func (u_ URLSchemeTaskWrapper) DidReceiveResponse(response foundation.IURLResponse) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveResponse:"), objc.ExtractPtr(response))
}

func (u_ URLSchemeTaskWrapper) DidReceiveData(data []byte) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didReceiveData:"), data)
}

func (u_ URLSchemeTaskWrapper) DidFinish() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFinish"))
}

func (u_ URLSchemeTaskWrapper) DidFailWithError(error foundation.IError) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("didFailWithError:"), objc.ExtractPtr(error))
}

func (u_ URLSchemeTaskWrapper) ImplementsRequest() bool {
	return u_.RespondsToSelector(objc.GetSelector("request"))
}

func (u_ URLSchemeTaskWrapper) Request() foundation.URLRequest {
	rv := objc.CallMethod[foundation.URLRequest](u_, objc.GetSelector("request"))
	return rv
}
