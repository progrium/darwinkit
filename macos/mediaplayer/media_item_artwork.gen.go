// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MediaItemArtwork] class.
var MediaItemArtworkClass = _MediaItemArtworkClass{objc.GetClass("MPMediaItemArtwork")}

type _MediaItemArtworkClass struct {
	objc.Class
}

// An interface definition for the [MediaItemArtwork] class.
type IMediaItemArtwork interface {
	objc.IObject
	ImageWithSize(size coregraphics.Size) appkit.Image
	Bounds() coregraphics.Rect
	ImageCropRect() coregraphics.Rect
}

// A graphical image, such as music album cover art, associated with a media item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork?language=objc
type MediaItemArtwork struct {
	objc.Object
}

func MediaItemArtworkFrom(ptr unsafe.Pointer) MediaItemArtwork {
	return MediaItemArtwork{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MediaItemArtwork) InitWithBoundsSizeRequestHandler(boundsSize coregraphics.Size, requestHandler func(size coregraphics.Size) appkit.Image) MediaItemArtwork {
	rv := objc.Call[MediaItemArtwork](m_, objc.Sel("initWithBoundsSize:requestHandler:"), boundsSize, requestHandler)
	return rv
}

// Creates a new image from existing artwork with the specified bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/1649704-initwithboundssize?language=objc
func NewMediaItemArtworkWithBoundsSizeRequestHandler(boundsSize coregraphics.Size, requestHandler func(size coregraphics.Size) appkit.Image) MediaItemArtwork {
	instance := MediaItemArtworkClass.Alloc().InitWithBoundsSizeRequestHandler(boundsSize, requestHandler)
	instance.Autorelease()
	return instance
}

func (mc _MediaItemArtworkClass) Alloc() MediaItemArtwork {
	rv := objc.Call[MediaItemArtwork](mc, objc.Sel("alloc"))
	return rv
}

func MediaItemArtwork_Alloc() MediaItemArtwork {
	return MediaItemArtworkClass.Alloc()
}

func (mc _MediaItemArtworkClass) New() MediaItemArtwork {
	rv := objc.Call[MediaItemArtwork](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMediaItemArtwork() MediaItemArtwork {
	return MediaItemArtworkClass.New()
}

func (m_ MediaItemArtwork) Init() MediaItemArtwork {
	rv := objc.Call[MediaItemArtwork](m_, objc.Sel("init"))
	return rv
}

// Returns the artwork image for an item at the given size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/1621736-imagewithsize?language=objc
func (m_ MediaItemArtwork) ImageWithSize(size coregraphics.Size) appkit.Image {
	rv := objc.Call[appkit.Image](m_, objc.Sel("imageWithSize:"), size)
	return rv
}

// The maximum size, in points, of the image associated with the media item artwork. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/1621723-bounds?language=objc
func (m_ MediaItemArtwork) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](m_, objc.Sel("bounds"))
	return rv
}

// The bounds, in points, of the content area for the full size image associated with the media item artwork. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitemartwork/1621760-imagecroprect?language=objc
func (m_ MediaItemArtwork) ImageCropRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](m_, objc.Sel("imageCropRect"))
	return rv
}
