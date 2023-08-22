// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGroupNormalizationNode] class.
var CNNGroupNormalizationNodeClass = _CNNGroupNormalizationNodeClass{objc.GetClass("MPSCNNGroupNormalizationNode")}

type _CNNGroupNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNGroupNormalizationNode] class.
type ICNNGroupNormalizationNode interface {
	INNFilterNode
	TrainingStyle() NNTrainingStyle
	SetTrainingStyle(value NNTrainingStyle)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode?language=objc
type CNNGroupNormalizationNode struct {
	NNFilterNode
}

func CNNGroupNormalizationNodeFrom(ptr unsafe.Pointer) CNNGroupNormalizationNode {
	return CNNGroupNormalizationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNGroupNormalizationNodeClass) NodeWithSourceDataSource(source INNImageNode, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNGroupNormalizationDataSource", dataSource)
	rv := objc.Call[CNNGroupNormalizationNode](cc, objc.Sel("nodeWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152573-nodewithsource?language=objc
func CNNGroupNormalizationNode_NodeWithSourceDataSource(source INNImageNode, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalizationNode {
	return CNNGroupNormalizationNodeClass.NodeWithSourceDataSource(source, dataSource)
}

func (c_ CNNGroupNormalizationNode) InitWithSourceDataSource(source INNImageNode, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNGroupNormalizationDataSource", dataSource)
	rv := objc.Call[CNNGroupNormalizationNode](c_, objc.Sel("initWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3152572-initwithsource?language=objc
func NewCNNGroupNormalizationNodeWithSourceDataSource(source INNImageNode, dataSource PCNNGroupNormalizationDataSource) CNNGroupNormalizationNode {
	instance := CNNGroupNormalizationNodeClass.Alloc().InitWithSourceDataSource(source, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNGroupNormalizationNodeClass) Alloc() CNNGroupNormalizationNode {
	rv := objc.Call[CNNGroupNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNGroupNormalizationNode_Alloc() CNNGroupNormalizationNode {
	return CNNGroupNormalizationNodeClass.Alloc()
}

func (cc _CNNGroupNormalizationNodeClass) New() CNNGroupNormalizationNode {
	rv := objc.Call[CNNGroupNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGroupNormalizationNode() CNNGroupNormalizationNode {
	return CNNGroupNormalizationNodeClass.New()
}

func (c_ CNNGroupNormalizationNode) Init() CNNGroupNormalizationNode {
	rv := objc.Call[CNNGroupNormalizationNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle?language=objc
func (c_ CNNGroupNormalizationNode) TrainingStyle() NNTrainingStyle {
	rv := objc.Call[NNTrainingStyle](c_, objc.Sel("trainingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationnode/3197823-trainingstyle?language=objc
func (c_ CNNGroupNormalizationNode) SetTrainingStyle(value NNTrainingStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setTrainingStyle:"), value)
}
