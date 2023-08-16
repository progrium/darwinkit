// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// This protocol is adopted by the document view of a WebFrameView. You can extend WebKit to support additional MIME types by implementing your own document view and document representation classes to render data for specific MIME types. You register those classes using the WebFrame registerViewClass:representationClass:forMIMEType: method. Classes that adopt this protocol are expected to be subclasses of NSView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdocumentview?language=objc
type PWebDocumentView interface {
}

// A concrete type wrapper for the [PWebDocumentView] protocol.
type WebDocumentViewWrapper struct {
	objc.Object
}
