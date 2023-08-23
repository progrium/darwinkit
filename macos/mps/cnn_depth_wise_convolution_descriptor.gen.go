// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDepthWiseConvolutionDescriptor] class.
var CNNDepthWiseConvolutionDescriptorClass = _CNNDepthWiseConvolutionDescriptorClass{objc.GetClass("MPSCNNDepthWiseConvolutionDescriptor")}

type _CNNDepthWiseConvolutionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNDepthWiseConvolutionDescriptor] class.
type ICNNDepthWiseConvolutionDescriptor interface {
	ICNNConvolutionDescriptor
	ChannelMultiplier() uint
}

// A description of a convolution object that does depthwise convolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor?language=objc
type CNNDepthWiseConvolutionDescriptor struct {
	CNNConvolutionDescriptor
}

func CNNDepthWiseConvolutionDescriptorFrom(ptr unsafe.Pointer) CNNDepthWiseConvolutionDescriptor {
	return CNNDepthWiseConvolutionDescriptor{
		CNNConvolutionDescriptor: CNNConvolutionDescriptorFrom(ptr),
	}
}

func (cc _CNNDepthWiseConvolutionDescriptorClass) Alloc() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Call[CNNDepthWiseConvolutionDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CNNDepthWiseConvolutionDescriptor_Alloc() CNNDepthWiseConvolutionDescriptor {
	return CNNDepthWiseConvolutionDescriptorClass.Alloc()
}

func (cc _CNNDepthWiseConvolutionDescriptorClass) New() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Call[CNNDepthWiseConvolutionDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDepthWiseConvolutionDescriptor() CNNDepthWiseConvolutionDescriptor {
	return CNNDepthWiseConvolutionDescriptorClass.New()
}

func (c_ CNNDepthWiseConvolutionDescriptor) Init() CNNDepthWiseConvolutionDescriptor {
	rv := objc.Call[CNNDepthWiseConvolutionDescriptor](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNDepthWiseConvolutionDescriptorClass) CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNDepthWiseConvolutionDescriptor {
	rv := objc.Call[CNNDepthWiseConvolutionDescriptor](cc, objc.Sel("cnnConvolutionDescriptorWithKernelWidth:kernelHeight:inputFeatureChannels:outputFeatureChannels:"), kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondescriptor/1648813-cnnconvolutiondescriptorwithkern?language=objc
func CNNDepthWiseConvolutionDescriptor_CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth uint, kernelHeight uint, inputFeatureChannels uint, outputFeatureChannels uint) CNNDepthWiseConvolutionDescriptor {
	return CNNDepthWiseConvolutionDescriptorClass.CnnConvolutionDescriptorWithKernelWidthKernelHeightInputFeatureChannelsOutputFeatureChannels(kernelWidth, kernelHeight, inputFeatureChannels, outputFeatureChannels)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndepthwiseconvolutiondescriptor/2919731-channelmultiplier?language=objc
func (c_ CNNDepthWiseConvolutionDescriptor) ChannelMultiplier() uint {
	rv := objc.Call[uint](c_, objc.Sel("channelMultiplier"))
	return rv
}
