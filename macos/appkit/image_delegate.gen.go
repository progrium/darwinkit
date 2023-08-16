// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that you can use to respond to drawing failures and manage incremental loads. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagedelegate?language=objc
type PImageDelegate interface {
	// optional
	ImageDidNotDrawInRect(sender Image, rect foundation.Rect) IImage
	HasImageDidNotDrawInRect() bool
}

// A delegate implementation builder for the [PImageDelegate] protocol.
type ImageDelegate struct {
	_ImageDidNotDrawInRect func(sender Image, rect foundation.Rect) IImage
}

func (di *ImageDelegate) HasImageDidNotDrawInRect() bool {
	return di._ImageDidNotDrawInRect != nil
}

// Tells the delegate that the image object is unable, for whatever reason, to lock focus on its image or draw in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagedelegate/1519927-imagedidnotdraw?language=objc
func (di *ImageDelegate) SetImageDidNotDrawInRect(f func(sender Image, rect foundation.Rect) IImage) {
	di._ImageDidNotDrawInRect = f
}

// Tells the delegate that the image object is unable, for whatever reason, to lock focus on its image or draw in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagedelegate/1519927-imagedidnotdraw?language=objc
func (di *ImageDelegate) ImageDidNotDrawInRect(sender Image, rect foundation.Rect) IImage {
	return di._ImageDidNotDrawInRect(sender, rect)
}

// A concrete type wrapper for the [PImageDelegate] protocol.
type ImageDelegateWrapper struct {
	objc.Object
}

func (i_ ImageDelegateWrapper) HasImageDidNotDrawInRect() bool {
	return i_.RespondsToSelector(objc.Sel("imageDidNotDraw:inRect:"))
}

// Tells the delegate that the image object is unable, for whatever reason, to lock focus on its image or draw in the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagedelegate/1519927-imagedidnotdraw?language=objc
func (i_ ImageDelegateWrapper) ImageDidNotDrawInRect(sender IImage, rect foundation.Rect) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageDidNotDraw:inRect:"), objc.Ptr(sender), rect)
	return rv
}
