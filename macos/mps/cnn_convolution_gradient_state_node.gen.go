// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionGradientStateNode] class.
var CNNConvolutionGradientStateNodeClass = _CNNConvolutionGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionGradientStateNode")}

type _CNNConvolutionGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionGradientStateNode] class.
type ICNNConvolutionGradientStateNode interface {
	INNGradientStateNode
}

// A representation of a gradient convolution state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstatenode?language=objc
type CNNConvolutionGradientStateNode struct {
	NNGradientStateNode
}

func CNNConvolutionGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionGradientStateNode {
	return CNNConvolutionGradientStateNode{
		NNGradientStateNode: NNGradientStateNodeFrom(ptr),
	}
}

func (cc _CNNConvolutionGradientStateNodeClass) Alloc() CNNConvolutionGradientStateNode {
	rv := objc.Call[CNNConvolutionGradientStateNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionGradientStateNode_Alloc() CNNConvolutionGradientStateNode {
	return CNNConvolutionGradientStateNodeClass.Alloc()
}

func (cc _CNNConvolutionGradientStateNodeClass) New() CNNConvolutionGradientStateNode {
	rv := objc.Call[CNNConvolutionGradientStateNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionGradientStateNode() CNNConvolutionGradientStateNode {
	return CNNConvolutionGradientStateNodeClass.New()
}

func (c_ CNNConvolutionGradientStateNode) Init() CNNConvolutionGradientStateNode {
	rv := objc.Call[CNNConvolutionGradientStateNode](c_, objc.Sel("init"))
	return rv
}
