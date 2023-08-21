// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CustomImageRep] class.
var CustomImageRepClass = _CustomImageRepClass{objc.GetClass("NSCustomImageRep")}

type _CustomImageRepClass struct {
	objc.Class
}

// An interface definition for the [CustomImageRep] class.
type ICustomImageRep interface {
	IImageRep
	DrawSelector() objc.Selector
	DrawingHandler() func(arg0 foundation.Rect) bool
	Delegate() objc.Object
}

// An object that uses a delegate object to render an image from a custom format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep?language=objc
type CustomImageRep struct {
	ImageRep
}

func CustomImageRepFrom(ptr unsafe.Pointer) CustomImageRep {
	return CustomImageRep{
		ImageRep: ImageRepFrom(ptr),
	}
}

func (c_ CustomImageRep) InitWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) CustomImageRep {
	rv := objc.Call[CustomImageRep](c_, objc.Sel("initWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return rv
}

// Initializes a representation of an image of the specified size and flipped status, using a block to draw its content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep/1526521-initwithsize?language=objc
func NewCustomImageRepWithSizeFlippedDrawingHandler(size foundation.Size, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler func(dstRect foundation.Rect) bool) CustomImageRep {
	instance := CustomImageRepClass.Alloc().InitWithSizeFlippedDrawingHandler(size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	instance.Autorelease()
	return instance
}

func (c_ CustomImageRep) InitWithDrawSelectorDelegate(selector objc.Selector, delegate objc.IObject) CustomImageRep {
	rv := objc.Call[CustomImageRep](c_, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	return rv
}

// Returns a representation of an image initialized with the specified delegate information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep/1533328-initwithdrawselector?language=objc
func NewCustomImageRepWithDrawSelectorDelegate(selector objc.Selector, delegate objc.IObject) CustomImageRep {
	instance := CustomImageRepClass.Alloc().InitWithDrawSelectorDelegate(selector, delegate)
	instance.Autorelease()
	return instance
}

func (cc _CustomImageRepClass) Alloc() CustomImageRep {
	rv := objc.Call[CustomImageRep](cc, objc.Sel("alloc"))
	return rv
}

func CustomImageRep_Alloc() CustomImageRep {
	return CustomImageRepClass.Alloc()
}

func (cc _CustomImageRepClass) New() CustomImageRep {
	rv := objc.Call[CustomImageRep](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCustomImageRep() CustomImageRep {
	return CustomImageRepClass.New()
}

func (c_ CustomImageRep) Init() CustomImageRep {
	rv := objc.Call[CustomImageRep](c_, objc.Sel("init"))
	return rv
}

// The selector for the delegate's drawing method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep/1529935-drawselector?language=objc
func (c_ CustomImageRep) DrawSelector() objc.Selector {
	rv := objc.Call[objc.Selector](c_, objc.Sel("drawSelector"))
	return rv
}

// The destination rectangle of the drawing handler block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep/1527316-drawinghandler?language=objc
func (c_ CustomImageRep) DrawingHandler() func(arg0 foundation.Rect) bool {
	rv := objc.Call[func(arg0 foundation.Rect) bool](c_, objc.Sel("drawingHandler"))
	return rv
}

// The delegate object that renders the image for the image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscustomimagerep/1534716-delegate?language=objc
func (c_ CustomImageRep) Delegate() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("delegate"))
	return rv
}
