// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeGradientStateNode] class.
var CNNConvolutionTransposeGradientStateNodeClass = _CNNConvolutionTransposeGradientStateNodeClass{objc.GetClass("MPSCNNConvolutionTransposeGradientStateNode")}

type _CNNConvolutionTransposeGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradientStateNode] class.
type ICNNConvolutionTransposeGradientStateNode interface {
	ICNNConvolutionGradientStateNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstatenode?language=objc
type CNNConvolutionTransposeGradientStateNode struct {
	CNNConvolutionGradientStateNode
}

func CNNConvolutionTransposeGradientStateNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientStateNode {
	return CNNConvolutionTransposeGradientStateNode{
		CNNConvolutionGradientStateNode: CNNConvolutionGradientStateNodeFrom(ptr),
	}
}

func (cc _CNNConvolutionTransposeGradientStateNodeClass) Alloc() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Call[CNNConvolutionTransposeGradientStateNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTransposeGradientStateNode_Alloc() CNNConvolutionTransposeGradientStateNode {
	return CNNConvolutionTransposeGradientStateNodeClass.Alloc()
}

func (cc _CNNConvolutionTransposeGradientStateNodeClass) New() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Call[CNNConvolutionTransposeGradientStateNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeGradientStateNode() CNNConvolutionTransposeGradientStateNode {
	return CNNConvolutionTransposeGradientStateNodeClass.New()
}

func (c_ CNNConvolutionTransposeGradientStateNode) Init() CNNConvolutionTransposeGradientStateNode {
	rv := objc.Call[CNNConvolutionTransposeGradientStateNode](c_, objc.Sel("init"))
	return rv
}
