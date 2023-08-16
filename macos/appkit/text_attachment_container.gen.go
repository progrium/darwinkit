// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that defines the interface to text attachment objects from a layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentcontainer?language=objc
type PTextAttachmentContainer interface {
	// optional
	ImageForBoundsTextContainerCharacterIndex(imageBounds coregraphics.Rect, textContainer TextContainer, charIndex uint) IImage
	HasImageForBoundsTextContainerCharacterIndex() bool

	// optional
	AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer TextContainer, lineFrag coregraphics.Rect, position coregraphics.Point, charIndex uint) coregraphics.Rect
	HasAttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool
}

// A concrete type wrapper for the [PTextAttachmentContainer] protocol.
type TextAttachmentContainerWrapper struct {
	objc.Object
}

func (t_ TextAttachmentContainerWrapper) HasImageForBoundsTextContainerCharacterIndex() bool {
	return t_.RespondsToSelector(objc.Sel("imageForBounds:textContainer:characterIndex:"))
}

// Returns the image object that the layout manager renders in the specified image bounds rectangle inside the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentcontainer/1508386-imageforbounds?language=objc
func (t_ TextAttachmentContainerWrapper) ImageForBoundsTextContainerCharacterIndex(imageBounds coregraphics.Rect, textContainer ITextContainer, charIndex uint) Image {
	rv := objc.Call[Image](t_, objc.Sel("imageForBounds:textContainer:characterIndex:"), imageBounds, objc.Ptr(textContainer), charIndex)
	return rv
}

func (t_ TextAttachmentContainerWrapper) HasAttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return t_.RespondsToSelector(objc.Sel("attachmentBoundsForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"))
}

// Returns the layout bounds of the text attachment to the layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentcontainer/1508382-attachmentboundsfortextcontainer?language=objc
func (t_ TextAttachmentContainerWrapper) AttachmentBoundsForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer ITextContainer, lineFrag coregraphics.Rect, position coregraphics.Point, charIndex uint) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("attachmentBoundsForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), objc.Ptr(textContainer), lineFrag, position, charIndex)
	return rv
}
