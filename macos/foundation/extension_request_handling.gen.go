// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface an app extension uses to respond to a request from a host app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionrequesthandling?language=objc
type PExtensionRequestHandling interface {
	// optional
	BeginRequestWithExtensionContext(context ExtensionContext)
	HasBeginRequestWithExtensionContext() bool
}

// A concrete type wrapper for the [PExtensionRequestHandling] protocol.
type ExtensionRequestHandlingWrapper struct {
	objc.Object
}

func (e_ ExtensionRequestHandlingWrapper) HasBeginRequestWithExtensionContext() bool {
	return e_.RespondsToSelector(objc.Sel("beginRequestWithExtensionContext:"))
}

// Tells the extension to prepare for a host app’s request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionrequesthandling/1413395-beginrequestwithextensioncontext?language=objc
func (e_ ExtensionRequestHandlingWrapper) BeginRequestWithExtensionContext(context IExtensionContext) {
	objc.Call[objc.Void](e_, objc.Sel("beginRequestWithExtensionContext:"), objc.Ptr(context))
}
