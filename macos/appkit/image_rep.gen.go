// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageRep] class.
var ImageRepClass = _ImageRepClass{objc.GetClass("NSImageRep")}

type _ImageRepClass struct {
	objc.Class
}

// An interface definition for the [ImageRep] class.
type IImageRep interface {
	objc.IObject
	DrawAtPoint(point foundation.Point) bool
	DrawInRect(rect foundation.Rect) bool
	Draw() bool
	CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef
	LayoutDirection() ImageLayoutDirection
	SetLayoutDirection(value ImageLayoutDirection)
	ColorSpaceName() ColorSpaceName
	SetColorSpaceName(value ColorSpaceName)
	IsOpaque() bool
	SetOpaque(value bool)
	PixelsWide() int
	SetPixelsWide(value int)
	HasAlpha() bool
	SetAlpha(value bool)
	PixelsHigh() int
	SetPixelsHigh(value int)
	BitsPerSample() int
	SetBitsPerSample(value int)
	Size() foundation.Size
	SetSize(value foundation.Size)
}

// A semiabstract superclass that provides subclasses that you use to draw an image from a particular type of source data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep?language=objc
type ImageRep struct {
	objc.Object
}

func ImageRepFrom(ptr unsafe.Pointer) ImageRep {
	return ImageRep{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ ImageRep) Init() ImageRep {
	rv := objc.Call[ImageRep](i_, objc.Sel("init"))
	return rv
}

func (ic _ImageRepClass) Alloc() ImageRep {
	rv := objc.Call[ImageRep](ic, objc.Sel("alloc"))
	return rv
}

func ImageRep_Alloc() ImageRep {
	return ImageRepClass.Alloc()
}

func (ic _ImageRepClass) New() ImageRep {
	rv := objc.Call[ImageRep](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageRep() ImageRep {
	return ImageRepClass.New()
}

// Creates and returns an image representation object using the contents of the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534589-imagerepwithcontentsoffile?language=objc
func (ic _ImageRepClass) ImageRepWithContentsOfFile(filename string) ImageRep {
	rv := objc.Call[ImageRep](ic, objc.Sel("imageRepWithContentsOfFile:"), filename)
	return rv
}

// Creates and returns an image representation object using the contents of the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534589-imagerepwithcontentsoffile?language=objc
func ImageRep_ImageRepWithContentsOfFile(filename string) ImageRep {
	return ImageRepClass.ImageRepWithContentsOfFile(filename)
}

// Draws the image representation’s image data at the specified point in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1525632-drawatpoint?language=objc
func (i_ ImageRep) DrawAtPoint(point foundation.Point) bool {
	rv := objc.Call[bool](i_, objc.Sel("drawAtPoint:"), point)
	return rv
}

// Creates and returns an array of image representation objects initialized using the contents of the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526978-imagerepswithcontentsofurl?language=objc
func (ic _ImageRepClass) ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	rv := objc.Call[[]ImageRep](ic, objc.Sel("imageRepsWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns an array of image representation objects initialized using the contents of the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526978-imagerepswithcontentsofurl?language=objc
func ImageRep_ImageRepsWithContentsOfURL(url foundation.IURL) []ImageRep {
	return ImageRepClass.ImageRepsWithContentsOfURL(url)
}

// Returns the image representation subclass that handles the specified type of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534272-imagerepclassfordata?language=objc
func (ic _ImageRepClass) ImageRepClassForData(data []byte) objc.Class {
	rv := objc.Call[objc.Class](ic, objc.Sel("imageRepClassForData:"), data)
	return rv
}

// Returns the image representation subclass that handles the specified type of data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534272-imagerepclassfordata?language=objc
func ImageRep_ImageRepClassForData(data []byte) objc.Class {
	return ImageRepClass.ImageRepClassForData(data)
}

// Draws the image, scaling it (as needed) to fit the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535066-drawinrect?language=objc
func (i_ ImageRep) DrawInRect(rect foundation.Rect) bool {
	rv := objc.Call[bool](i_, objc.Sel("drawInRect:"), rect)
	return rv
}

// Implemented by subclasses to draw the image in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526674-draw?language=objc
func (i_ ImageRep) Draw() bool {
	rv := objc.Call[bool](i_, objc.Sel("draw"))
	return rv
}

// Creates and returns an image representation object using the contents of the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533918-imagerepwithpasteboard?language=objc
func (ic _ImageRepClass) ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	rv := objc.Call[ImageRep](ic, objc.Sel("imageRepWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Creates and returns an image representation object using the contents of the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533918-imagerepwithpasteboard?language=objc
func ImageRep_ImageRepWithPasteboard(pasteboard IPasteboard) ImageRep {
	return ImageRepClass.ImageRepWithPasteboard(pasteboard)
}

// Returns the image representation subclass that handles image data for the specified UTI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1529342-imagerepclassfortype?language=objc
func (ic _ImageRepClass) ImageRepClassForType(type_ string) objc.Class {
	rv := objc.Call[objc.Class](ic, objc.Sel("imageRepClassForType:"), type_)
	return rv
}

// Returns the image representation subclass that handles image data for the specified UTI. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1529342-imagerepclassfortype?language=objc
func ImageRep_ImageRepClassForType(type_ string) objc.Class {
	return ImageRepClass.ImageRepClassForType(type_)
}

// Creates and returns an image representation object using the data at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535840-imagerepwithcontentsofurl?language=objc
func (ic _ImageRepClass) ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	rv := objc.Call[ImageRep](ic, objc.Sel("imageRepWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns an image representation object using the data at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535840-imagerepwithcontentsofurl?language=objc
func ImageRep_ImageRepWithContentsOfURL(url foundation.IURL) ImageRep {
	return ImageRepClass.ImageRepWithContentsOfURL(url)
}

// Returns a Boolean value that indicates whether the receiver can initialize itself from the data on the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1527887-caninitwithpasteboard?language=objc
func (ic _ImageRepClass) CanInitWithPasteboard(pasteboard IPasteboard) bool {
	rv := objc.Call[bool](ic, objc.Sel("canInitWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Returns a Boolean value that indicates whether the receiver can initialize itself from the data on the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1527887-caninitwithpasteboard?language=objc
func ImageRep_CanInitWithPasteboard(pasteboard IPasteboard) bool {
	return ImageRepClass.CanInitWithPasteboard(pasteboard)
}

// Removes the specified image representation subclass from the registry of available image representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535027-unregisterimagerepclass?language=objc
func (ic _ImageRepClass) UnregisterImageRepClass(imageRepClass objc.IClass) {
	objc.Call[objc.Void](ic, objc.Sel("unregisterImageRepClass:"), objc.Ptr(imageRepClass))
}

// Removes the specified image representation subclass from the registry of available image representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535027-unregisterimagerepclass?language=objc
func ImageRep_UnregisterImageRepClass(imageRepClass objc.IClass) {
	ImageRepClass.UnregisterImageRepClass(imageRepClass)
}

// Returns a Boolean value that indicates whether the image representation can initialize itself from the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535427-caninitwithdata?language=objc
func (ic _ImageRepClass) CanInitWithData(data []byte) bool {
	rv := objc.Call[bool](ic, objc.Sel("canInitWithData:"), data)
	return rv
}

// Returns a Boolean value that indicates whether the image representation can initialize itself from the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535427-caninitwithdata?language=objc
func ImageRep_CanInitWithData(data []byte) bool {
	return ImageRepClass.CanInitWithData(data)
}

// Creates and returns an array of image representation objects initialized using the contents of the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526968-imagerepswithcontentsoffile?language=objc
func (ic _ImageRepClass) ImageRepsWithContentsOfFile(filename string) []ImageRep {
	rv := objc.Call[[]ImageRep](ic, objc.Sel("imageRepsWithContentsOfFile:"), filename)
	return rv
}

// Creates and returns an array of image representation objects initialized using the contents of the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526968-imagerepswithcontentsoffile?language=objc
func ImageRep_ImageRepsWithContentsOfFile(filename string) []ImageRep {
	return ImageRepClass.ImageRepsWithContentsOfFile(filename)
}

// Adds the specified class to the registry of available image representation subclasses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535393-registerimagerepclass?language=objc
func (ic _ImageRepClass) RegisterImageRepClass(imageRepClass objc.IClass) {
	objc.Call[objc.Void](ic, objc.Sel("registerImageRepClass:"), objc.Ptr(imageRepClass))
}

// Adds the specified class to the registry of available image representation subclasses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535393-registerimagerepclass?language=objc
func ImageRep_RegisterImageRepClass(imageRepClass objc.IClass) {
	ImageRepClass.RegisterImageRepClass(imageRepClass)
}

// Returns a Core Graphics image object that captures the drawing of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533478-cgimageforproposedrect?language=objc
func (i_ ImageRep) CGImageForProposedRectContextHints(proposedDestRect *foundation.Rect, context IGraphicsContext, hints map[ImageHintKey]objc.IObject) coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](i_, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, objc.Ptr(context), hints)
	return rv
}

// Creates and returns an array of image representation objects initialized using the contents of the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535586-imagerepswithpasteboard?language=objc
func (ic _ImageRepClass) ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	rv := objc.Call[[]ImageRep](ic, objc.Sel("imageRepsWithPasteboard:"), objc.Ptr(pasteboard))
	return rv
}

// Creates and returns an array of image representation objects initialized using the contents of the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535586-imagerepswithpasteboard?language=objc
func ImageRep_ImageRepsWithPasteboard(pasteboard IPasteboard) []ImageRep {
	return ImageRepClass.ImageRepsWithPasteboard(pasteboard)
}

// The layout direction for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1690237-layoutdirection?language=objc
func (i_ ImageRep) LayoutDirection() ImageLayoutDirection {
	rv := objc.Call[ImageLayoutDirection](i_, objc.Sel("layoutDirection"))
	return rv
}

// The layout direction for the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1690237-layoutdirection?language=objc
func (i_ ImageRep) SetLayoutDirection(value ImageLayoutDirection) {
	objc.Call[objc.Void](i_, objc.Sel("setLayoutDirection:"), value)
}

// The name of the color space used by the image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535395-colorspacename?language=objc
func (i_ ImageRep) ColorSpaceName() ColorSpaceName {
	rv := objc.Call[ColorSpaceName](i_, objc.Sel("colorSpaceName"))
	return rv
}

// The name of the color space used by the image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535395-colorspacename?language=objc
func (i_ ImageRep) SetColorSpaceName(value ColorSpaceName) {
	objc.Call[objc.Void](i_, objc.Sel("setColorSpaceName:"), value)
}

// Returns an array of UTI strings identifying the image types supported by the image representation, either directly or through a user-installed filter service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535246-imagetypes?language=objc
func (ic _ImageRepClass) ImageTypes() []string {
	rv := objc.Call[[]string](ic, objc.Sel("imageTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported by the image representation, either directly or through a user-installed filter service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1535246-imagetypes?language=objc
func ImageRep_ImageTypes() []string {
	return ImageRepClass.ImageTypes()
}

// A Boolean value that indicates whether the image is opaque. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1528462-opaque?language=objc
func (i_ ImageRep) IsOpaque() bool {
	rv := objc.Call[bool](i_, objc.Sel("isOpaque"))
	return rv
}

// A Boolean value that indicates whether the image is opaque. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1528462-opaque?language=objc
func (i_ ImageRep) SetOpaque(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setOpaque:"), value)
}

// The width of the image, measured in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526995-pixelswide?language=objc
func (i_ ImageRep) PixelsWide() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsWide"))
	return rv
}

// The width of the image, measured in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1526995-pixelswide?language=objc
func (i_ ImageRep) SetPixelsWide(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setPixelsWide:"), value)
}

// A Boolean value that indicates whether the image data has an alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534506-alpha?language=objc
func (i_ ImageRep) HasAlpha() bool {
	rv := objc.Call[bool](i_, objc.Sel("hasAlpha"))
	return rv
}

// A Boolean value that indicates whether the image data has an alpha channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534506-alpha?language=objc
func (i_ ImageRep) SetAlpha(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setAlpha:"), value)
}

// The height of the image, measured in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533989-pixelshigh?language=objc
func (i_ ImageRep) PixelsHigh() int {
	rv := objc.Call[int](i_, objc.Sel("pixelsHigh"))
	return rv
}

// The height of the image, measured in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533989-pixelshigh?language=objc
func (i_ ImageRep) SetPixelsHigh(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setPixelsHigh:"), value)
}

// Returns an array of UTI strings identifying the image types supported directly by the ime representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534816-imageunfilteredtypes?language=objc
func (ic _ImageRepClass) ImageUnfilteredTypes() []string {
	rv := objc.Call[[]string](ic, objc.Sel("imageUnfilteredTypes"))
	return rv
}

// Returns an array of UTI strings identifying the image types supported directly by the ime representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1534816-imageunfilteredtypes?language=objc
func ImageRep_ImageUnfilteredTypes() []string {
	return ImageRepClass.ImageUnfilteredTypes()
}

// Returns an array containing the registered image representation classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1525139-registeredimagerepclasses?language=objc
func (ic _ImageRepClass) RegisteredImageRepClasses() []objc.Class {
	rv := objc.Call[[]objc.Class](ic, objc.Sel("registeredImageRepClasses"))
	return rv
}

// Returns an array containing the registered image representation classes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1525139-registeredimagerepclasses?language=objc
func ImageRep_RegisteredImageRepClasses() []objc.Class {
	return ImageRepClass.RegisteredImageRepClasses()
}

// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533157-bitspersample?language=objc
func (i_ ImageRep) BitsPerSample() int {
	rv := objc.Call[int](i_, objc.Sel("bitsPerSample"))
	return rv
}

// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1533157-bitspersample?language=objc
func (i_ ImageRep) SetBitsPerSample(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setBitsPerSample:"), value)
}

// The size of the image representation, measured in points in the user coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1524374-size?language=objc
func (i_ ImageRep) Size() foundation.Size {
	rv := objc.Call[foundation.Size](i_, objc.Sel("size"))
	return rv
}

// The size of the image representation, measured in points in the user coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsimagerep/1524374-size?language=objc
func (i_ ImageRep) SetSize(value foundation.Size) {
	objc.Call[objc.Void](i_, objc.Sel("setSize:"), value)
}
