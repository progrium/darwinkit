// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// This protocol is adopted by document representation classes that handle specific MIME types. You can implement your own document view classes and document representation classes to render data for specific MIME types, and register those classes using the WebFrame registerViewClass:representationClass:forMIMEType: method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdocumentrepresentation?language=objc
type PWebDocumentRepresentation interface {
}

// A concrete type wrapper for the [PWebDocumentRepresentation] protocol.
type WebDocumentRepresentationWrapper struct {
	objc.Object
}
