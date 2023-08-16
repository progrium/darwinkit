// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutGuide] class.
var LayoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}

type _LayoutGuideClass struct {
	objc.Class
}

// An interface definition for the [LayoutGuide] class.
type ILayoutGuide interface {
	objc.IObject
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	TopAnchor() LayoutYAxisAnchor
	LeftAnchor() LayoutXAxisAnchor
	CenterYAnchor() LayoutYAxisAnchor
	RightAnchor() LayoutXAxisAnchor
	BottomAnchor() LayoutYAxisAnchor
	HeightAnchor() LayoutDimension
	TrailingAnchor() LayoutXAxisAnchor
	LeadingAnchor() LayoutXAxisAnchor
	Frame() foundation.Rect
	OwningView() View
	SetOwningView(value IView)
	HasAmbiguousLayout() bool
	CenterXAnchor() LayoutXAxisAnchor
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	WidthAnchor() LayoutDimension
}

// A rectangular area that can interact with Auto Layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide?language=objc
type LayoutGuide struct {
	objc.Object
}

func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{
		Object: objc.ObjectFrom(ptr),
	}
}

func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.Call[LayoutGuide](lc, objc.Sel("alloc"))
	return rv
}

func LayoutGuide_Alloc() LayoutGuide {
	return LayoutGuideClass.Alloc()
}

func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.Call[LayoutGuide](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutGuide() LayoutGuide {
	return LayoutGuideClass.New()
}

func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.Call[LayoutGuide](l_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1641956-constraintsaffectinglayoutforori?language=objc
func (l_ LayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](l_, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}

// A layout anchor representing the top edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1534159-topanchor?language=objc
func (l_ LayoutGuide) TopAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](l_, objc.Sel("topAnchor"))
	return rv
}

// A layout anchor representing the left edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1530723-leftanchor?language=objc
func (l_ LayoutGuide) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("leftAnchor"))
	return rv
}

// A layout anchor representing the vertical center of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1530346-centeryanchor?language=objc
func (l_ LayoutGuide) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](l_, objc.Sel("centerYAnchor"))
	return rv
}

// A layout anchor representing the right edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1534892-rightanchor?language=objc
func (l_ LayoutGuide) RightAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("rightAnchor"))
	return rv
}

// A layout anchor representing the bottom edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1529133-bottomanchor?language=objc
func (l_ LayoutGuide) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](l_, objc.Sel("bottomAnchor"))
	return rv
}

// A layout anchor representing the height of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1532150-heightanchor?language=objc
func (l_ LayoutGuide) HeightAnchor() LayoutDimension {
	rv := objc.Call[LayoutDimension](l_, objc.Sel("heightAnchor"))
	return rv
}

// A layout anchor representing the trailing edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1528204-trailinganchor?language=objc
func (l_ LayoutGuide) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("trailingAnchor"))
	return rv
}

// A layout anchor representing the leading edge of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1534720-leadinganchor?language=objc
func (l_ LayoutGuide) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("leadingAnchor"))
	return rv
}

// The layout guide’s frame in its owning view’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1533744-frame?language=objc
func (l_ LayoutGuide) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](l_, objc.Sel("frame"))
	return rv
}

// The view that owns this layout guide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1533971-owningview?language=objc
func (l_ LayoutGuide) OwningView() View {
	rv := objc.Call[View](l_, objc.Sel("owningView"))
	return rv
}

// The view that owns this layout guide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1533971-owningview?language=objc
func (l_ LayoutGuide) SetOwningView(value IView) {
	objc.Call[objc.Void](l_, objc.Sel("setOwningView:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1641955-hasambiguouslayout?language=objc
func (l_ LayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.Call[bool](l_, objc.Sel("hasAmbiguousLayout"))
	return rv
}

// A layout anchor representing the horizontal center of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1528997-centerxanchor?language=objc
func (l_ LayoutGuide) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](l_, objc.Sel("centerXAnchor"))
	return rv
}

// A string used to identify the layout guide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1524460-identifier?language=objc
func (l_ LayoutGuide) Identifier() UserInterfaceItemIdentifier {
	rv := objc.Call[UserInterfaceItemIdentifier](l_, objc.Sel("identifier"))
	return rv
}

// A string used to identify the layout guide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1524460-identifier?language=objc
func (l_ LayoutGuide) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](l_, objc.Sel("setIdentifier:"), value)
}

// A layout anchor representing the width of the layout guide’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutguide/1527215-widthanchor?language=objc
func (l_ LayoutGuide) WidthAnchor() LayoutDimension {
	rv := objc.Call[LayoutDimension](l_, objc.Sel("widthAnchor"))
	return rv
}