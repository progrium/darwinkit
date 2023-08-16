// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that define the interface to attachment objects from a text layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentlayout?language=objc
type PTextAttachmentLayout interface {
	// optional
	ImageForBoundsAttributesLocationTextContainer(bounds coregraphics.Rect, attributes map[foundation.AttributedStringKey]objc.Object, location TextLocationWrapper, textContainer TextContainer) IImage
	HasImageForBoundsAttributesLocationTextContainer() bool

	// optional
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.Object, location TextLocationWrapper, textContainer TextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect
	HasAttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition() bool

	// optional
	ViewProviderForParentViewLocationTextContainer(parentView View, location TextLocationWrapper, textContainer TextContainer) ITextAttachmentViewProvider
	HasViewProviderForParentViewLocationTextContainer() bool
}

// A concrete type wrapper for the [PTextAttachmentLayout] protocol.
type TextAttachmentLayoutWrapper struct {
	objc.Object
}

func (t_ TextAttachmentLayoutWrapper) HasImageForBoundsAttributesLocationTextContainer() bool {
	return t_.RespondsToSelector(objc.Sel("imageForBounds:attributes:location:textContainer:"))
}

// Returns the image object rendered at the bounds and inside the text container you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentlayout/3857591-imageforbounds?language=objc
func (t_ TextAttachmentLayoutWrapper) ImageForBoundsAttributesLocationTextContainer(bounds coregraphics.Rect, attributes map[foundation.AttributedStringKey]objc.IObject, location PTextLocation, textContainer ITextContainer) Image {
	po2 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[Image](t_, objc.Sel("imageForBounds:attributes:location:textContainer:"), bounds, attributes, po2, objc.Ptr(textContainer))
	return rv
}

func (t_ TextAttachmentLayoutWrapper) HasAttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition() bool {
	return t_.RespondsToSelector(objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"))
}

// Returns the layout bounds of the attachment you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentlayout/3857590-attachmentboundsforattributes?language=objc
func (t_ TextAttachmentLayoutWrapper) AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.IObject, location PTextLocation, textContainer ITextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, po1, objc.Ptr(textContainer), proposedLineFragment, position)
	return rv
}

func (t_ TextAttachmentLayoutWrapper) HasViewProviderForParentViewLocationTextContainer() bool {
	return t_.RespondsToSelector(objc.Sel("viewProviderForParentView:location:textContainer:"))
}

// Returns the text attachment view provider corresponding to the file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentlayout/3857592-viewproviderforparentview?language=objc
func (t_ TextAttachmentLayoutWrapper) ViewProviderForParentViewLocationTextContainer(parentView IView, location PTextLocation, textContainer ITextContainer) TextAttachmentViewProvider {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextAttachmentViewProvider](t_, objc.Sel("viewProviderForParentView:location:textContainer:"), objc.Ptr(parentView), po1, objc.Ptr(textContainer))
	return rv
}
