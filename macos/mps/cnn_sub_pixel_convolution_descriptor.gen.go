// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSubPixelConvolutionDescriptor] class.
var CNNSubPixelConvolutionDescriptorClass = _CNNSubPixelConvolutionDescriptorClass{objc.GetClass("MPSCNNSubPixelConvolutionDescriptor")}

type _CNNSubPixelConvolutionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNSubPixelConvolutionDescriptor] class.
type ICNNSubPixelConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	SubPixelScaleFactor() uint
	SetSubPixelScaleFactor(value uint)
}

// A description of a convolution object that does subpixel upsampling and reshaping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor?language=objc
type CNNSubPixelConvolutionDescriptor struct {
	CNNConvolutionDescriptor
}

func CNNSubPixelConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNSubPixelConvolutionDescriptor {
	return CNNSubPixelConvolutionDescriptor{
		CNNConvolutionDescriptor: CNNConvolutionDescriptorFrom(ptr),
	}
}

func (cc _CNNSubPixelConvolutionDescriptorClass) Alloc() CNNSubPixelConvolutionDescriptor {
	rv := objc.Call[CNNSubPixelConvolutionDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CNNSubPixelConvolutionDescriptor_Alloc() CNNSubPixelConvolutionDescriptor {
	return CNNSubPixelConvolutionDescriptorClass.Alloc()
}

func (cc _CNNSubPixelConvolutionDescriptorClass) New() CNNSubPixelConvolutionDescriptor {
	rv := objc.Call[CNNSubPixelConvolutionDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSubPixelConvolutionDescriptor() CNNSubPixelConvolutionDescriptor {
	return CNNSubPixelConvolutionDescriptorClass.New()
}

func (c_ CNNSubPixelConvolutionDescriptor) Init() CNNSubPixelConvolutionDescriptor {
	rv := objc.Call[CNNSubPixelConvolutionDescriptor](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNSubPixelConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNSubPixelConvolutionDescriptor {
	rv := objc.Call[CNNSubPixelConvolutionDescriptor](cc, objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648813-cnnconvolutiondescriptorwithkern?language=objc
func CNNSubPixelConvolutionDescriptor_CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNSubPixelConvolutionDescriptor {
	return CNNSubPixelConvolutionDescriptorClass.CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor?language=objc
func (c_ CNNSubPixelConvolutionDescriptor) SubPixelScaleFactor() uint {
	rv := objc.Call[uint](c_, objc.Sel("subPixelScaleFactor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubpixelconvolutiondescriptor/2875156-subpixelscalefactor?language=objc
func (c_ CNNSubPixelConvolutionDescriptor) SetSubPixelScaleFactor(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSubPixelScaleFactor:"), value)
}
