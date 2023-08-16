// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled?language=objc
type PNinePartTiled interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetFlipYTiles(value bool)
	HasSetFlipYTiles() bool

	// optional
	FlipYTiles() bool
	HasFlipYTiles() bool

	// optional
	SetGrowAmount(value coregraphics.Point)
	HasSetGrowAmount() bool

	// optional
	GrowAmount() coregraphics.Point
	HasGrowAmount() bool

	// optional
	SetBreakpoint0(value coregraphics.Point)
	HasSetBreakpoint0() bool

	// optional
	Breakpoint0() coregraphics.Point
	HasBreakpoint0() bool

	// optional
	SetBreakpoint1(value coregraphics.Point)
	HasSetBreakpoint1() bool

	// optional
	Breakpoint1() coregraphics.Point
	HasBreakpoint1() bool
}

// A concrete type wrapper for the [PNinePartTiled] protocol.
type NinePartTiledWrapper struct {
	objc.Object
}

func (n_ NinePartTiledWrapper) HasSetInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600185-inputimage?language=objc
func (n_ NinePartTiledWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](n_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (n_ NinePartTiledWrapper) HasInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600185-inputimage?language=objc
func (n_ NinePartTiledWrapper) InputImage() Image {
	rv := objc.Call[Image](n_, objc.Sel("inputImage"))
	return rv
}

func (n_ NinePartTiledWrapper) HasSetFlipYTiles() bool {
	return n_.RespondsToSelector(objc.Sel("setFlipYTiles:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600183-flipytiles?language=objc
func (n_ NinePartTiledWrapper) SetFlipYTiles(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setFlipYTiles:"), value)
}

func (n_ NinePartTiledWrapper) HasFlipYTiles() bool {
	return n_.RespondsToSelector(objc.Sel("flipYTiles"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600183-flipytiles?language=objc
func (n_ NinePartTiledWrapper) FlipYTiles() bool {
	rv := objc.Call[bool](n_, objc.Sel("flipYTiles"))
	return rv
}

func (n_ NinePartTiledWrapper) HasSetGrowAmount() bool {
	return n_.RespondsToSelector(objc.Sel("setGrowAmount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600184-growamount?language=objc
func (n_ NinePartTiledWrapper) SetGrowAmount(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setGrowAmount:"), value)
}

func (n_ NinePartTiledWrapper) HasGrowAmount() bool {
	return n_.RespondsToSelector(objc.Sel("growAmount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600184-growamount?language=objc
func (n_ NinePartTiledWrapper) GrowAmount() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("growAmount"))
	return rv
}

func (n_ NinePartTiledWrapper) HasSetBreakpoint0() bool {
	return n_.RespondsToSelector(objc.Sel("setBreakpoint0:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600181-breakpoint0?language=objc
func (n_ NinePartTiledWrapper) SetBreakpoint0(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setBreakpoint0:"), value)
}

func (n_ NinePartTiledWrapper) HasBreakpoint0() bool {
	return n_.RespondsToSelector(objc.Sel("breakpoint0"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600181-breakpoint0?language=objc
func (n_ NinePartTiledWrapper) Breakpoint0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("breakpoint0"))
	return rv
}

func (n_ NinePartTiledWrapper) HasSetBreakpoint1() bool {
	return n_.RespondsToSelector(objc.Sel("setBreakpoint1:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600182-breakpoint1?language=objc
func (n_ NinePartTiledWrapper) SetBreakpoint1(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setBreakpoint1:"), value)
}

func (n_ NinePartTiledWrapper) HasBreakpoint1() bool {
	return n_.RespondsToSelector(objc.Sel("breakpoint1"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinineparttiled/3600182-breakpoint1?language=objc
func (n_ NinePartTiledWrapper) Breakpoint1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("breakpoint1"))
	return rv
}
