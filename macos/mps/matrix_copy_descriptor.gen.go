// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixCopyDescriptor] class.
var MatrixCopyDescriptorClass = _MatrixCopyDescriptorClass{objc.GetClass("MPSMatrixCopyDescriptor")}

type _MatrixCopyDescriptorClass struct {
	objc.Class
}

// An interface definition for the [MatrixCopyDescriptor] class.
type IMatrixCopyDescriptor interface {
	objc.IObject
	SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets)
}

// A description of multiple matrix copy operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor?language=objc
type MatrixCopyDescriptor struct {
	objc.Object
}

func MatrixCopyDescriptorFrom(ptr unsafe.Pointer) MatrixCopyDescriptor {
	return MatrixCopyDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MatrixCopyDescriptorClass) DescriptorWithSourceMatrixDestinationMatrixOffsets(sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets) MatrixCopyDescriptor {
	rv := objc.Call[MatrixCopyDescriptor](mc, objc.Sel("descriptorWithSourceMatrix:destinationMatrix:offsets:"), objc.Ptr(sourceMatrix), objc.Ptr(destinationMatrix), offsets)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915333-descriptorwithsourcematrix?language=objc
func MatrixCopyDescriptor_DescriptorWithSourceMatrixDestinationMatrixOffsets(sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets) MatrixCopyDescriptor {
	return MatrixCopyDescriptorClass.DescriptorWithSourceMatrixDestinationMatrixOffsets(sourceMatrix, destinationMatrix, offsets)
}

func (m_ MatrixCopyDescriptor) InitWithDeviceCount(device metal.PDevice, count uint) MatrixCopyDescriptor {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixCopyDescriptor](m_, objc.Sel("initWithDevice:count:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915324-initwithdevice?language=objc
func NewMatrixCopyDescriptorWithDeviceCount(device metal.PDevice, count uint) MatrixCopyDescriptor {
	instance := MatrixCopyDescriptorClass.Alloc().InitWithDeviceCount(device, count)
	instance.Autorelease()
	return instance
}

func (m_ MatrixCopyDescriptor) InitWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices []IMatrix, destinationMatrices []IMatrix, offsets IVector, byteOffset uint) MatrixCopyDescriptor {
	rv := objc.Call[MatrixCopyDescriptor](m_, objc.Sel("initWithSourceMatrices:destinationMatrices:offsetVector:offset:"), sourceMatrices, destinationMatrices, objc.Ptr(offsets), byteOffset)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915344-initwithsourcematrices?language=objc
func NewMatrixCopyDescriptorWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices []IMatrix, destinationMatrices []IMatrix, offsets IVector, byteOffset uint) MatrixCopyDescriptor {
	instance := MatrixCopyDescriptorClass.Alloc().InitWithSourceMatricesDestinationMatricesOffsetVectorOffset(sourceMatrices, destinationMatrices, offsets, byteOffset)
	instance.Autorelease()
	return instance
}

func (mc _MatrixCopyDescriptorClass) Alloc() MatrixCopyDescriptor {
	rv := objc.Call[MatrixCopyDescriptor](mc, objc.Sel("alloc"))
	return rv
}

func MatrixCopyDescriptor_Alloc() MatrixCopyDescriptor {
	return MatrixCopyDescriptorClass.Alloc()
}

func (mc _MatrixCopyDescriptorClass) New() MatrixCopyDescriptor {
	rv := objc.Call[MatrixCopyDescriptor](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixCopyDescriptor() MatrixCopyDescriptor {
	return MatrixCopyDescriptorClass.New()
}

func (m_ MatrixCopyDescriptor) Init() MatrixCopyDescriptor {
	rv := objc.Call[MatrixCopyDescriptor](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixcopydescriptor/2915331-setcopyoperationatindex?language=objc
func (m_ MatrixCopyDescriptor) SetCopyOperationAtIndexSourceMatrixDestinationMatrixOffsets(index uint, sourceMatrix IMatrix, destinationMatrix IMatrix, offsets MatrixCopyOffsets) {
	objc.Call[objc.Void](m_, objc.Sel("setCopyOperationAtIndex:sourceMatrix:destinationMatrix:offsets:"), index, objc.Ptr(sourceMatrix), objc.Ptr(destinationMatrix), offsets)
}
