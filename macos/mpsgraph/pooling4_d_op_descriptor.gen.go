// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Pooling4DOpDescriptor] class.
var Pooling4DOpDescriptorClass = _Pooling4DOpDescriptorClass{objc.GetClass("MPSGraphPooling4DOpDescriptor")}

type _Pooling4DOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [Pooling4DOpDescriptor] class.
type IPooling4DOpDescriptor interface {
	objc.IObject
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	CeilMode() bool
	SetCeilMode(value bool)
	KernelSizes() []foundation.Number
	SetKernelSizes(value []foundation.INumber)
	PaddingValues() []foundation.Number
	SetPaddingValues(value []foundation.INumber)
	DilationRates() []foundation.Number
	SetDilationRates(value []foundation.INumber)
	ReturnIndicesMode() PoolingReturnIndicesMode
	SetReturnIndicesMode(value PoolingReturnIndicesMode)
	ReturnIndicesDataType() mps.DataType
	SetReturnIndicesDataType(value mps.DataType)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	Strides() []foundation.Number
	SetStrides(value []foundation.INumber)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor?language=objc
type Pooling4DOpDescriptor struct {
	objc.Object
}

func Pooling4DOpDescriptorFrom(ptr unsafe.Pointer) Pooling4DOpDescriptor {
	return Pooling4DOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _Pooling4DOpDescriptorClass) DescriptorWithKernelSizesPaddingStyle(kernelSizes []foundation.INumber, paddingStyle PaddingStyle) Pooling4DOpDescriptor {
	rv := objc.Call[Pooling4DOpDescriptor](pc, objc.Sel("descriptorWithKernelSizes:paddingStyle:"), kernelSizes, paddingStyle)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750698-descriptorwithkernelsizes?language=objc
func Pooling4DOpDescriptor_DescriptorWithKernelSizesPaddingStyle(kernelSizes []foundation.INumber, paddingStyle PaddingStyle) Pooling4DOpDescriptor {
	return Pooling4DOpDescriptorClass.DescriptorWithKernelSizesPaddingStyle(kernelSizes, paddingStyle)
}

func (pc _Pooling4DOpDescriptorClass) Alloc() Pooling4DOpDescriptor {
	rv := objc.Call[Pooling4DOpDescriptor](pc, objc.Sel("alloc"))
	return rv
}

func Pooling4DOpDescriptor_Alloc() Pooling4DOpDescriptor {
	return Pooling4DOpDescriptorClass.Alloc()
}

func (pc _Pooling4DOpDescriptorClass) New() Pooling4DOpDescriptor {
	rv := objc.Call[Pooling4DOpDescriptor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPooling4DOpDescriptor() Pooling4DOpDescriptor {
	return Pooling4DOpDescriptorClass.New()
}

func (p_ Pooling4DOpDescriptor) Init() Pooling4DOpDescriptor {
	rv := objc.Call[Pooling4DOpDescriptor](p_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750703-paddingstyle?language=objc
func (p_ Pooling4DOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](p_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750703-paddingstyle?language=objc
func (p_ Pooling4DOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750697-ceilmode?language=objc
func (p_ Pooling4DOpDescriptor) CeilMode() bool {
	rv := objc.Call[bool](p_, objc.Sel("ceilMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750697-ceilmode?language=objc
func (p_ Pooling4DOpDescriptor) SetCeilMode(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCeilMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750702-kernelsizes?language=objc
func (p_ Pooling4DOpDescriptor) KernelSizes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("kernelSizes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750702-kernelsizes?language=objc
func (p_ Pooling4DOpDescriptor) SetKernelSizes(value []foundation.INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setKernelSizes:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750704-paddingvalues?language=objc
func (p_ Pooling4DOpDescriptor) PaddingValues() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("paddingValues"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750704-paddingvalues?language=objc
func (p_ Pooling4DOpDescriptor) SetPaddingValues(value []foundation.INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingValues:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750700-dilationrates?language=objc
func (p_ Pooling4DOpDescriptor) DilationRates() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("dilationRates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750700-dilationrates?language=objc
func (p_ Pooling4DOpDescriptor) SetDilationRates(value []foundation.INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setDilationRates:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3919747-returnindicesmode?language=objc
func (p_ Pooling4DOpDescriptor) ReturnIndicesMode() PoolingReturnIndicesMode {
	rv := objc.Call[PoolingReturnIndicesMode](p_, objc.Sel("returnIndicesMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3919747-returnindicesmode?language=objc
func (p_ Pooling4DOpDescriptor) SetReturnIndicesMode(value PoolingReturnIndicesMode) {
	objc.Call[objc.Void](p_, objc.Sel("setReturnIndicesMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3919746-returnindicesdatatype?language=objc
func (p_ Pooling4DOpDescriptor) ReturnIndicesDataType() mps.DataType {
	rv := objc.Call[mps.DataType](p_, objc.Sel("returnIndicesDataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3919746-returnindicesdatatype?language=objc
func (p_ Pooling4DOpDescriptor) SetReturnIndicesDataType(value mps.DataType) {
	objc.Call[objc.Void](p_, objc.Sel("setReturnIndicesDataType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750701-includezeropadtoaverage?language=objc
func (p_ Pooling4DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Call[bool](p_, objc.Sel("includeZeroPadToAverage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750701-includezeropadtoaverage?language=objc
func (p_ Pooling4DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750705-strides?language=objc
func (p_ Pooling4DOpDescriptor) Strides() []foundation.Number {
	rv := objc.Call[[]foundation.Number](p_, objc.Sel("strides"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling4dopdescriptor/3750705-strides?language=objc
func (p_ Pooling4DOpDescriptor) SetStrides(value []foundation.INumber) {
	objc.Call[objc.Void](p_, objc.Sel("setStrides:"), value)
}
