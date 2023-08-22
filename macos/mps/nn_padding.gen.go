// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

// The protocol that provides a description of how kernels should pad images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding?language=objc
type PNNPadding interface {
	// optional
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []Image, sourceStates []State, kernel Kernel, inDescriptor ImageDescriptor) IImageDescriptor
	HasDestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor() bool

	// optional
	PaddingMethod() NNPaddingMethod
	HasPaddingMethod() bool

	// optional
	Inverse() objc.IObject
	HasInverse() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PNNPadding] protocol.
type NNPaddingWrapper struct {
	objc.Object
}

func (n_ NNPaddingWrapper) HasDestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor() bool {
	return n_.RespondsToSelector(objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2867167-destinationimagedescriptorforsou?language=objc
func (n_ NNPaddingWrapper) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages []IImage, sourceStates []IState, kernel IKernel, inDescriptor IImageDescriptor) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](n_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), sourceImages, sourceStates, objc.Ptr(kernel), objc.Ptr(inDescriptor))
	return rv
}

func (n_ NNPaddingWrapper) HasPaddingMethod() bool {
	return n_.RespondsToSelector(objc.Sel("paddingMethod"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2866950-paddingmethod?language=objc
func (n_ NNPaddingWrapper) PaddingMethod() NNPaddingMethod {
	rv := objc.Call[NNPaddingMethod](n_, objc.Sel("paddingMethod"))
	return rv
}

func (n_ NNPaddingWrapper) HasInverse() bool {
	return n_.RespondsToSelector(objc.Sel("inverse"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2942456-inverse?language=objc
func (n_ NNPaddingWrapper) Inverse() objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("inverse"))
	return rv
}

func (n_ NNPaddingWrapper) HasLabel() bool {
	return n_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadding/2889870-label?language=objc
func (n_ NNPaddingWrapper) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}
