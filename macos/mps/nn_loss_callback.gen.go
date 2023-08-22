// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlosscallback?language=objc
type PNNLossCallback interface {
	// optional
	ScalarWeightForSourceImageDestinationImage(sourceImage Image, destinationImage Image) float64
	HasScalarWeightForSourceImageDestinationImage() bool
}

// A concrete type wrapper for the [PNNLossCallback] protocol.
type NNLossCallbackWrapper struct {
	objc.Object
}

func (n_ NNLossCallbackWrapper) HasScalarWeightForSourceImageDestinationImage() bool {
	return n_.RespondsToSelector(objc.Sel("scalarWeightForSourceImage:destinationImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlosscallback/3131851-scalarweightforsourceimage?language=objc
func (n_ NNLossCallbackWrapper) ScalarWeightForSourceImageDestinationImage(sourceImage IImage, destinationImage IImage) float64 {
	rv := objc.Call[float64](n_, objc.Sel("scalarWeightForSourceImage:destinationImage:"), objc.Ptr(sourceImage), objc.Ptr(destinationImage))
	return rv
}
