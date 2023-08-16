// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched?language=objc
type PNinePartStretched interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

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

// A concrete type wrapper for the [PNinePartStretched] protocol.
type NinePartStretchedWrapper struct {
	objc.Object
}

func (n_ NinePartStretchedWrapper) HasSetInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600179-inputimage?language=objc
func (n_ NinePartStretchedWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](n_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (n_ NinePartStretchedWrapper) HasInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600179-inputimage?language=objc
func (n_ NinePartStretchedWrapper) InputImage() Image {
	rv := objc.Call[Image](n_, objc.Sel("inputImage"))
	return rv
}

func (n_ NinePartStretchedWrapper) HasSetGrowAmount() bool {
	return n_.RespondsToSelector(objc.Sel("setGrowAmount:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600178-growamount?language=objc
func (n_ NinePartStretchedWrapper) SetGrowAmount(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setGrowAmount:"), value)
}

func (n_ NinePartStretchedWrapper) HasGrowAmount() bool {
	return n_.RespondsToSelector(objc.Sel("growAmount"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600178-growamount?language=objc
func (n_ NinePartStretchedWrapper) GrowAmount() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("growAmount"))
	return rv
}

func (n_ NinePartStretchedWrapper) HasSetBreakpoint0() bool {
	return n_.RespondsToSelector(objc.Sel("setBreakpoint0:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600176-breakpoint0?language=objc
func (n_ NinePartStretchedWrapper) SetBreakpoint0(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setBreakpoint0:"), value)
}

func (n_ NinePartStretchedWrapper) HasBreakpoint0() bool {
	return n_.RespondsToSelector(objc.Sel("breakpoint0"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600176-breakpoint0?language=objc
func (n_ NinePartStretchedWrapper) Breakpoint0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("breakpoint0"))
	return rv
}

func (n_ NinePartStretchedWrapper) HasSetBreakpoint1() bool {
	return n_.RespondsToSelector(objc.Sel("setBreakpoint1:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600177-breakpoint1?language=objc
func (n_ NinePartStretchedWrapper) SetBreakpoint1(value coregraphics.Point) {
	objc.Call[objc.Void](n_, objc.Sel("setBreakpoint1:"), value)
}

func (n_ NinePartStretchedWrapper) HasBreakpoint1() bool {
	return n_.RespondsToSelector(objc.Sel("breakpoint1"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cininepartstretched/3600177-breakpoint1?language=objc
func (n_ NinePartStretchedWrapper) Breakpoint1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](n_, objc.Sel("breakpoint1"))
	return rv
}
