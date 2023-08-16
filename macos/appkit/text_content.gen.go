// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that describes specific kinds of input content types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontent?language=objc
type PTextContent interface {
	// optional
	SetContentType(value TextContentType)
	HasSetContentType() bool

	// optional
	ContentType() TextContentType
	HasContentType() bool
}

// A concrete type wrapper for the [PTextContent] protocol.
type TextContentWrapper struct {
	objc.Object
}

func (t_ TextContentWrapper) HasSetContentType() bool {
	return t_.RespondsToSelector(objc.Sel("setContentType:"))
}

// The semantic meaning for a text input area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontent/3566859-contenttype?language=objc
func (t_ TextContentWrapper) SetContentType(value TextContentType) {
	objc.Call[objc.Void](t_, objc.Sel("setContentType:"), value)
}

func (t_ TextContentWrapper) HasContentType() bool {
	return t_.RespondsToSelector(objc.Sel("contentType"))
}

// The semantic meaning for a text input area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcontent/3566859-contenttype?language=objc
func (t_ TextContentWrapper) ContentType() TextContentType {
	rv := objc.Call[TextContentType](t_, objc.Sel("contentType"))
	return rv
}
