// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StencilOpDescriptor] class.
var StencilOpDescriptorClass = _StencilOpDescriptorClass{objc.GetClass("MPSGraphStencilOpDescriptor")}

type _StencilOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [StencilOpDescriptor] class.
type IStencilOpDescriptor interface {
	objc.IObject
	PaddingConstant() float64
	SetPaddingConstant(value float64)
	Offsets() *foundation.Array
	SetOffsets(value *foundation.Array)
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	DilationRates() *foundation.Array
	SetDilationRates(value *foundation.Array)
	ExplicitPadding() *foundation.Array
	SetExplicitPadding(value *foundation.Array)
	BoundaryMode() PaddingMode
	SetBoundaryMode(value PaddingMode)
	ReductionMode() ReductionMode
	SetReductionMode(value ReductionMode)
	Strides() *foundation.Array
	SetStrides(value *foundation.Array)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor?language=objc
type StencilOpDescriptor struct {
	objc.Object
}

func StencilOpDescriptorFrom(ptr unsafe.Pointer) StencilOpDescriptor {
	return StencilOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StencilOpDescriptorClass) DescriptorWithOffsetsExplicitPadding(offsets *foundation.Array, explicitPadding *foundation.Array) StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("descriptorWithOffsets:explicitPadding:"), offsets, explicitPadding)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787609-descriptorwithoffsets?language=objc
func StencilOpDescriptor_DescriptorWithOffsetsExplicitPadding(offsets *foundation.Array, explicitPadding *foundation.Array) StencilOpDescriptor {
	return StencilOpDescriptorClass.DescriptorWithOffsetsExplicitPadding(offsets, explicitPadding)
}

func (sc _StencilOpDescriptorClass) DescriptorWithExplicitPadding(explicitPadding *foundation.Array) StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("descriptorWithExplicitPadding:"), explicitPadding)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787608-descriptorwithexplicitpadding?language=objc
func StencilOpDescriptor_DescriptorWithExplicitPadding(explicitPadding *foundation.Array) StencilOpDescriptor {
	return StencilOpDescriptorClass.DescriptorWithExplicitPadding(explicitPadding)
}

func (sc _StencilOpDescriptorClass) DescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode ReductionMode, offsets *foundation.Array, strides *foundation.Array, dilationRates *foundation.Array, explicitPadding *foundation.Array, boundaryMode PaddingMode, paddingStyle PaddingStyle, paddingConstant float64) StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("descriptorWithReductionMode:offsets:strides:dilationRates:explicitPadding:boundaryMode:paddingStyle:paddingConstant:"), reductionMode, offsets, strides, dilationRates, explicitPadding, boundaryMode, paddingStyle, paddingConstant)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787611-descriptorwithreductionmode?language=objc
func StencilOpDescriptor_DescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode ReductionMode, offsets *foundation.Array, strides *foundation.Array, dilationRates *foundation.Array, explicitPadding *foundation.Array, boundaryMode PaddingMode, paddingStyle PaddingStyle, paddingConstant float64) StencilOpDescriptor {
	return StencilOpDescriptorClass.DescriptorWithReductionModeOffsetsStridesDilationRatesExplicitPaddingBoundaryModePaddingStylePaddingConstant(reductionMode, offsets, strides, dilationRates, explicitPadding, boundaryMode, paddingStyle, paddingConstant)
}

func (sc _StencilOpDescriptorClass) DescriptorWithPaddingStyle(paddingStyle PaddingStyle) StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("descriptorWithPaddingStyle:"), paddingStyle)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787610-descriptorwithpaddingstyle?language=objc
func StencilOpDescriptor_DescriptorWithPaddingStyle(paddingStyle PaddingStyle) StencilOpDescriptor {
	return StencilOpDescriptorClass.DescriptorWithPaddingStyle(paddingStyle)
}

func (sc _StencilOpDescriptorClass) Alloc() StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("alloc"))
	return rv
}

func StencilOpDescriptor_Alloc() StencilOpDescriptor {
	return StencilOpDescriptorClass.Alloc()
}

func (sc _StencilOpDescriptorClass) New() StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStencilOpDescriptor() StencilOpDescriptor {
	return StencilOpDescriptorClass.New()
}

func (s_ StencilOpDescriptor) Init() StencilOpDescriptor {
	rv := objc.Call[StencilOpDescriptor](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787615-paddingconstant?language=objc
func (s_ StencilOpDescriptor) PaddingConstant() float64 {
	rv := objc.Call[float64](s_, objc.Sel("paddingConstant"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787615-paddingconstant?language=objc
func (s_ StencilOpDescriptor) SetPaddingConstant(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setPaddingConstant:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787614-offsets?language=objc
func (s_ StencilOpDescriptor) Offsets() *foundation.Array {
	rv := objc.Call[*foundation.Array](s_, objc.Sel("offsets"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787614-offsets?language=objc
func (s_ StencilOpDescriptor) SetOffsets(value *foundation.Array) {
	objc.Call[objc.Void](s_, objc.Sel("setOffsets:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787616-paddingstyle?language=objc
func (s_ StencilOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](s_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787616-paddingstyle?language=objc
func (s_ StencilOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](s_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787612-dilationrates?language=objc
func (s_ StencilOpDescriptor) DilationRates() *foundation.Array {
	rv := objc.Call[*foundation.Array](s_, objc.Sel("dilationRates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787612-dilationrates?language=objc
func (s_ StencilOpDescriptor) SetDilationRates(value *foundation.Array) {
	objc.Call[objc.Void](s_, objc.Sel("setDilationRates:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787613-explicitpadding?language=objc
func (s_ StencilOpDescriptor) ExplicitPadding() *foundation.Array {
	rv := objc.Call[*foundation.Array](s_, objc.Sel("explicitPadding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787613-explicitpadding?language=objc
func (s_ StencilOpDescriptor) SetExplicitPadding(value *foundation.Array) {
	objc.Call[objc.Void](s_, objc.Sel("setExplicitPadding:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787607-boundarymode?language=objc
func (s_ StencilOpDescriptor) BoundaryMode() PaddingMode {
	rv := objc.Call[PaddingMode](s_, objc.Sel("boundaryMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787607-boundarymode?language=objc
func (s_ StencilOpDescriptor) SetBoundaryMode(value PaddingMode) {
	objc.Call[objc.Void](s_, objc.Sel("setBoundaryMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787617-reductionmode?language=objc
func (s_ StencilOpDescriptor) ReductionMode() ReductionMode {
	rv := objc.Call[ReductionMode](s_, objc.Sel("reductionMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787617-reductionmode?language=objc
func (s_ StencilOpDescriptor) SetReductionMode(value ReductionMode) {
	objc.Call[objc.Void](s_, objc.Sel("setReductionMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787618-strides?language=objc
func (s_ StencilOpDescriptor) Strides() *foundation.Array {
	rv := objc.Call[*foundation.Array](s_, objc.Sel("strides"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphstencilopdescriptor/3787618-strides?language=objc
func (s_ StencilOpDescriptor) SetStrides(value *foundation.Array) {
	objc.Call[objc.Void](s_, objc.Sel("setStrides:"), value)
}
