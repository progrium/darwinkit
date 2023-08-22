// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Pooling2DOpDescriptor] class.
var Pooling2DOpDescriptorClass = _Pooling2DOpDescriptorClass{objc.GetClass("MPSGraphPooling2DOpDescriptor")}

type _Pooling2DOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [Pooling2DOpDescriptor] class.
type IPooling2DOpDescriptor interface {
	objc.IObject
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
	StrideInX() uint
	SetStrideInX(value uint)
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	PaddingTop() uint
	SetPaddingTop(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	CeilMode() bool
	SetCeilMode(value bool)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
	ReturnIndicesMode() PoolingReturnIndicesMode
	SetReturnIndicesMode(value PoolingReturnIndicesMode)
	DataLayout() TensorNamedDataLayout
	SetDataLayout(value TensorNamedDataLayout)
	ReturnIndicesDataType() mps.DataType
	SetReturnIndicesDataType(value mps.DataType)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	KernelHeight() uint
	SetKernelHeight(value uint)
	IncludeZeroPadToAverage() bool
	SetIncludeZeroPadToAverage(value bool)
	KernelWidth() uint
	SetKernelWidth(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor?language=objc
type Pooling2DOpDescriptor struct {
	objc.Object
}

func Pooling2DOpDescriptorFrom(ptr unsafe.Pointer) Pooling2DOpDescriptor {
	return Pooling2DOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _Pooling2DOpDescriptorClass) DescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout) Pooling2DOpDescriptor {
	rv := objc.Call[Pooling2DOpDescriptor](pc, objc.Sel("descriptorWithKernelWidth:kernelHeight:strideInX:strideInY:paddingStyle:dataLayout:"), kernelWidth, kernelHeight, strideInX, strideInY, paddingStyle, dataLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564712-descriptorwithkernelwidth?language=objc
func Pooling2DOpDescriptor_DescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth uint, kernelHeight uint, strideInX uint, strideInY uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout) Pooling2DOpDescriptor {
	return Pooling2DOpDescriptorClass.DescriptorWithKernelWidthKernelHeightStrideInXStrideInYPaddingStyleDataLayout(kernelWidth, kernelHeight, strideInX, strideInY, paddingStyle, dataLayout)
}

func (pc _Pooling2DOpDescriptorClass) Alloc() Pooling2DOpDescriptor {
	rv := objc.Call[Pooling2DOpDescriptor](pc, objc.Sel("alloc"))
	return rv
}

func Pooling2DOpDescriptor_Alloc() Pooling2DOpDescriptor {
	return Pooling2DOpDescriptorClass.Alloc()
}

func (pc _Pooling2DOpDescriptorClass) New() Pooling2DOpDescriptor {
	rv := objc.Call[Pooling2DOpDescriptor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPooling2DOpDescriptor() Pooling2DOpDescriptor {
	return Pooling2DOpDescriptorClass.New()
}

func (p_ Pooling2DOpDescriptor) Init() Pooling2DOpDescriptor {
	rv := objc.Call[Pooling2DOpDescriptor](p_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564722-setexplicitpaddingwithpaddinglef?language=objc
func (p_ Pooling2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Call[objc.Void](p_, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564723-strideinx?language=objc
func (p_ Pooling2DOpDescriptor) StrideInX() uint {
	rv := objc.Call[uint](p_, objc.Sel("strideInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564723-strideinx?language=objc
func (p_ Pooling2DOpDescriptor) SetStrideInX(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setStrideInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564720-paddingstyle?language=objc
func (p_ Pooling2DOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](p_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564720-paddingstyle?language=objc
func (p_ Pooling2DOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564717-paddingbottom?language=objc
func (p_ Pooling2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Call[uint](p_, objc.Sel("paddingBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564717-paddingbottom?language=objc
func (p_ Pooling2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingBottom:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564721-paddingtop?language=objc
func (p_ Pooling2DOpDescriptor) PaddingTop() uint {
	rv := objc.Call[uint](p_, objc.Sel("paddingTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564721-paddingtop?language=objc
func (p_ Pooling2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingTop:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564724-strideiny?language=objc
func (p_ Pooling2DOpDescriptor) StrideInY() uint {
	rv := objc.Call[uint](p_, objc.Sel("strideInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564724-strideiny?language=objc
func (p_ Pooling2DOpDescriptor) SetStrideInY(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setStrideInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3861842-ceilmode?language=objc
func (p_ Pooling2DOpDescriptor) CeilMode() bool {
	rv := objc.Call[bool](p_, objc.Sel("ceilMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3861842-ceilmode?language=objc
func (p_ Pooling2DOpDescriptor) SetCeilMode(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setCeilMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564718-paddingleft?language=objc
func (p_ Pooling2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Call[uint](p_, objc.Sel("paddingLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564718-paddingleft?language=objc
func (p_ Pooling2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingLeft:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564713-dilationrateinx?language=objc
func (p_ Pooling2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Call[uint](p_, objc.Sel("dilationRateInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564713-dilationrateinx?language=objc
func (p_ Pooling2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setDilationRateInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3919745-returnindicesmode?language=objc
func (p_ Pooling2DOpDescriptor) ReturnIndicesMode() PoolingReturnIndicesMode {
	rv := objc.Call[PoolingReturnIndicesMode](p_, objc.Sel("returnIndicesMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3919745-returnindicesmode?language=objc
func (p_ Pooling2DOpDescriptor) SetReturnIndicesMode(value PoolingReturnIndicesMode) {
	objc.Call[objc.Void](p_, objc.Sel("setReturnIndicesMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564710-datalayout?language=objc
func (p_ Pooling2DOpDescriptor) DataLayout() TensorNamedDataLayout {
	rv := objc.Call[TensorNamedDataLayout](p_, objc.Sel("dataLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564710-datalayout?language=objc
func (p_ Pooling2DOpDescriptor) SetDataLayout(value TensorNamedDataLayout) {
	objc.Call[objc.Void](p_, objc.Sel("setDataLayout:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3919744-returnindicesdatatype?language=objc
func (p_ Pooling2DOpDescriptor) ReturnIndicesDataType() mps.DataType {
	rv := objc.Call[mps.DataType](p_, objc.Sel("returnIndicesDataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3919744-returnindicesdatatype?language=objc
func (p_ Pooling2DOpDescriptor) SetReturnIndicesDataType(value mps.DataType) {
	objc.Call[objc.Void](p_, objc.Sel("setReturnIndicesDataType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564714-dilationrateiny?language=objc
func (p_ Pooling2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Call[uint](p_, objc.Sel("dilationRateInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564714-dilationrateiny?language=objc
func (p_ Pooling2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setDilationRateInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564715-kernelheight?language=objc
func (p_ Pooling2DOpDescriptor) KernelHeight() uint {
	rv := objc.Call[uint](p_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564715-kernelheight?language=objc
func (p_ Pooling2DOpDescriptor) SetKernelHeight(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setKernelHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3861843-includezeropadtoaverage?language=objc
func (p_ Pooling2DOpDescriptor) IncludeZeroPadToAverage() bool {
	rv := objc.Call[bool](p_, objc.Sel("includeZeroPadToAverage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3861843-includezeropadtoaverage?language=objc
func (p_ Pooling2DOpDescriptor) SetIncludeZeroPadToAverage(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setIncludeZeroPadToAverage:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564716-kernelwidth?language=objc
func (p_ Pooling2DOpDescriptor) KernelWidth() uint {
	rv := objc.Call[uint](p_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564716-kernelwidth?language=objc
func (p_ Pooling2DOpDescriptor) SetKernelWidth(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setKernelWidth:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564719-paddingright?language=objc
func (p_ Pooling2DOpDescriptor) PaddingRight() uint {
	rv := objc.Call[uint](p_, objc.Sel("paddingRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphpooling2dopdescriptor/3564719-paddingright?language=objc
func (p_ Pooling2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setPaddingRight:"), value)
}
