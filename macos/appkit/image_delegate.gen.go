// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IImageDelegate interface {
	ImplementsImageDidNotDrawInRect() bool
	// optional
	ImageDidNotDrawInRect(sender Image, rect foundation.Rect) IImage
	ImplementsImageDidLoadPartOfRepresentationWithValidRows() bool
	// optional
	ImageDidLoadPartOfRepresentationWithValidRows(image Image, rep ImageRep, rows int)
	ImplementsImageDidLoadRepresentationWithStatus() bool
	// optional
	ImageDidLoadRepresentationWithStatus(image Image, rep ImageRep, status ImageLoadStatus)
	ImplementsImageDidLoadRepresentationHeader() bool
	// optional
	ImageDidLoadRepresentationHeader(image Image, rep ImageRep)
	ImplementsImageWillLoadRepresentation() bool
	// optional
	ImageWillLoadRepresentation(image Image, rep ImageRep)
}

type ImageDelegate struct {
	_ImageDidNotDrawInRect                         func(sender Image, rect foundation.Rect) IImage
	_ImageDidLoadPartOfRepresentationWithValidRows func(image Image, rep ImageRep, rows int)
	_ImageDidLoadRepresentationWithStatus          func(image Image, rep ImageRep, status ImageLoadStatus)
	_ImageDidLoadRepresentationHeader              func(image Image, rep ImageRep)
	_ImageWillLoadRepresentation                   func(image Image, rep ImageRep)
}

func (di *ImageDelegate) ImplementsImageDidNotDrawInRect() bool {
	return di._ImageDidNotDrawInRect != nil
}

func (di *ImageDelegate) SetImageDidNotDrawInRect(f func(sender Image, rect foundation.Rect) IImage) {
	di._ImageDidNotDrawInRect = f
}

func (di *ImageDelegate) ImageDidNotDrawInRect(sender Image, rect foundation.Rect) IImage {
	return di._ImageDidNotDrawInRect(sender, rect)
}
func (di *ImageDelegate) ImplementsImageDidLoadPartOfRepresentationWithValidRows() bool {
	return di._ImageDidLoadPartOfRepresentationWithValidRows != nil
}

func (di *ImageDelegate) SetImageDidLoadPartOfRepresentationWithValidRows(f func(image Image, rep ImageRep, rows int)) {
	di._ImageDidLoadPartOfRepresentationWithValidRows = f
}

func (di *ImageDelegate) ImageDidLoadPartOfRepresentationWithValidRows(image Image, rep ImageRep, rows int) {
	di._ImageDidLoadPartOfRepresentationWithValidRows(image, rep, rows)
}
func (di *ImageDelegate) ImplementsImageDidLoadRepresentationWithStatus() bool {
	return di._ImageDidLoadRepresentationWithStatus != nil
}

func (di *ImageDelegate) SetImageDidLoadRepresentationWithStatus(f func(image Image, rep ImageRep, status ImageLoadStatus)) {
	di._ImageDidLoadRepresentationWithStatus = f
}

func (di *ImageDelegate) ImageDidLoadRepresentationWithStatus(image Image, rep ImageRep, status ImageLoadStatus) {
	di._ImageDidLoadRepresentationWithStatus(image, rep, status)
}
func (di *ImageDelegate) ImplementsImageDidLoadRepresentationHeader() bool {
	return di._ImageDidLoadRepresentationHeader != nil
}

func (di *ImageDelegate) SetImageDidLoadRepresentationHeader(f func(image Image, rep ImageRep)) {
	di._ImageDidLoadRepresentationHeader = f
}

func (di *ImageDelegate) ImageDidLoadRepresentationHeader(image Image, rep ImageRep) {
	di._ImageDidLoadRepresentationHeader(image, rep)
}
func (di *ImageDelegate) ImplementsImageWillLoadRepresentation() bool {
	return di._ImageWillLoadRepresentation != nil
}

func (di *ImageDelegate) SetImageWillLoadRepresentation(f func(image Image, rep ImageRep)) {
	di._ImageWillLoadRepresentation = f
}

func (di *ImageDelegate) ImageWillLoadRepresentation(image Image, rep ImageRep) {
	di._ImageWillLoadRepresentation(image, rep)
}

type ImageDelegateWrapper struct {
	objc.Object
}

func (i_ ImageDelegateWrapper) ImplementsImageDidNotDrawInRect() bool {
	return i_.RespondsToSelector(objc.GetSelector("imageDidNotDraw:inRect:"))
}

func (i_ ImageDelegateWrapper) ImageDidNotDrawInRect(sender IImage, rect foundation.Rect) Image {
	rv := objc.CallMethod[Image](i_, objc.GetSelector("imageDidNotDraw:inRect:"), objc.ExtractPtr(sender), rect)
	return rv
}

func (i_ ImageDelegateWrapper) ImplementsImageDidLoadPartOfRepresentationWithValidRows() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadPartOfRepresentation:withValidRows:"))
}

func (i_ ImageDelegateWrapper) ImageDidLoadPartOfRepresentationWithValidRows(image IImage, rep IImageRep, rows int) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("image:didLoadPartOfRepresentation:withValidRows:"), objc.ExtractPtr(image), objc.ExtractPtr(rep), rows)
}

func (i_ ImageDelegateWrapper) ImplementsImageDidLoadRepresentationWithStatus() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentation:withStatus:"))
}

func (i_ ImageDelegateWrapper) ImageDidLoadRepresentationWithStatus(image IImage, rep IImageRep, status ImageLoadStatus) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("image:didLoadRepresentation:withStatus:"), objc.ExtractPtr(image), objc.ExtractPtr(rep), status)
}

func (i_ ImageDelegateWrapper) ImplementsImageDidLoadRepresentationHeader() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:didLoadRepresentationHeader:"))
}

func (i_ ImageDelegateWrapper) ImageDidLoadRepresentationHeader(image IImage, rep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("image:didLoadRepresentationHeader:"), objc.ExtractPtr(image), objc.ExtractPtr(rep))
}

func (i_ ImageDelegateWrapper) ImplementsImageWillLoadRepresentation() bool {
	return i_.RespondsToSelector(objc.GetSelector("image:willLoadRepresentation:"))
}

func (i_ ImageDelegateWrapper) ImageWillLoadRepresentation(image IImage, rep IImageRep) {
	objc.CallMethod[objc.Void](i_, objc.GetSelector("image:willLoadRepresentation:"), objc.ExtractPtr(image), objc.ExtractPtr(rep))
}
