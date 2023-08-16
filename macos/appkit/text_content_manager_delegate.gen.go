// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// The optional methods that delegates of content manager objects implement for customizing or validating text elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanagerdelegate?language=objc
type PTextContentManagerDelegate interface {
	// optional
	TextContentManagerTextElementAtLocation(textContentManager TextContentManager, location TextLocationWrapper) ITextElement
	HasTextContentManagerTextElementAtLocation() bool
}

// A delegate implementation builder for the [PTextContentManagerDelegate] protocol.
type TextContentManagerDelegate struct {
	_TextContentManagerTextElementAtLocation func(textContentManager TextContentManager, location TextLocationWrapper) ITextElement
}

func (di *TextContentManagerDelegate) HasTextContentManagerTextElementAtLocation() bool {
	return di._TextContentManagerTextElementAtLocation != nil
}

// The method the framework calls to return the text element at a specific location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanagerdelegate/3809932-textcontentmanager?language=objc
func (di *TextContentManagerDelegate) SetTextContentManagerTextElementAtLocation(f func(textContentManager TextContentManager, location TextLocationWrapper) ITextElement) {
	di._TextContentManagerTextElementAtLocation = f
}

// The method the framework calls to return the text element at a specific location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanagerdelegate/3809932-textcontentmanager?language=objc
func (di *TextContentManagerDelegate) TextContentManagerTextElementAtLocation(textContentManager TextContentManager, location TextLocationWrapper) ITextElement {
	return di._TextContentManagerTextElementAtLocation(textContentManager, location)
}

// A concrete type wrapper for the [PTextContentManagerDelegate] protocol.
type TextContentManagerDelegateWrapper struct {
	objc.Object
}

func (t_ TextContentManagerDelegateWrapper) HasTextContentManagerTextElementAtLocation() bool {
	return t_.RespondsToSelector(objc.Sel("textContentManager:textElementAtLocation:"))
}

// The method the framework calls to return the text element at a specific location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextcontentmanagerdelegate/3809932-textcontentmanager?language=objc
func (t_ TextContentManagerDelegateWrapper) TextContentManagerTextElementAtLocation(textContentManager ITextContentManager, location PTextLocation) TextElement {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextElement](t_, objc.Sel("textContentManager:textElementAtLocation:"), objc.Ptr(textContentManager), po1)
	return rv
}
