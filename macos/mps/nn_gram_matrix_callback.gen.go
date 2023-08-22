// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcallback?language=objc
type PNNGramMatrixCallback interface {
	// optional
	AlphaForSourceImageDestinationImage(sourceImage Image, destinationImage Image) float64
	HasAlphaForSourceImageDestinationImage() bool
}

// A concrete type wrapper for the [PNNGramMatrixCallback] protocol.
type NNGramMatrixCallbackWrapper struct {
	objc.Object
}

func (n_ NNGramMatrixCallbackWrapper) HasAlphaForSourceImageDestinationImage() bool {
	return n_.RespondsToSelector(objc.Sel("alphaForSourceImage:destinationImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcallback/3131846-alphaforsourceimage?language=objc
func (n_ NNGramMatrixCallbackWrapper) AlphaForSourceImageDestinationImage(sourceImage IImage, destinationImage IImage) float64 {
	rv := objc.Call[float64](n_, objc.Sel("alphaForSourceImage:destinationImage:"), objc.Ptr(sourceImage), objc.Ptr(destinationImage))
	return rv
}
