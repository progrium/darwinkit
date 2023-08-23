// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNInstanceNormalizationNode] class.
var CNNInstanceNormalizationNodeClass = _CNNInstanceNormalizationNodeClass{objc.GetClass("MPSCNNInstanceNormalizationNode")}

type _CNNInstanceNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNInstanceNormalizationNode] class.
type ICNNInstanceNormalizationNode interface {
	INNFilterNode
	TrainingStyle() NNTrainingStyle
	SetTrainingStyle(value NNTrainingStyle)
}

// A representation of an instance normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode?language=objc
type CNNInstanceNormalizationNode struct {
	NNFilterNode
}

func CNNInstanceNormalizationNodeFrom(ptr unsafe.Pointer) CNNInstanceNormalizationNode {
	return CNNInstanceNormalizationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNInstanceNormalizationNodeClass) NodeWithSourceDataSource(source INNImageNode, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNInstanceNormalizationDataSource", dataSource)
	rv := objc.Call[CNNInstanceNormalizationNode](cc, objc.Sel("nodeWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951941-nodewithsource?language=objc
func CNNInstanceNormalizationNode_NodeWithSourceDataSource(source INNImageNode, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalizationNode {
	return CNNInstanceNormalizationNodeClass.NodeWithSourceDataSource(source, dataSource)
}

func (c_ CNNInstanceNormalizationNode) InitWithSourceDataSource(source INNImageNode, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNInstanceNormalizationDataSource", dataSource)
	rv := objc.Call[CNNInstanceNormalizationNode](c_, objc.Sel("initWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/2951940-initwithsource?language=objc
func NewCNNInstanceNormalizationNodeWithSourceDataSource(source INNImageNode, dataSource PCNNInstanceNormalizationDataSource) CNNInstanceNormalizationNode {
	instance := CNNInstanceNormalizationNodeClass.Alloc().InitWithSourceDataSource(source, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNInstanceNormalizationNodeClass) Alloc() CNNInstanceNormalizationNode {
	rv := objc.Call[CNNInstanceNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNInstanceNormalizationNode_Alloc() CNNInstanceNormalizationNode {
	return CNNInstanceNormalizationNodeClass.Alloc()
}

func (cc _CNNInstanceNormalizationNodeClass) New() CNNInstanceNormalizationNode {
	rv := objc.Call[CNNInstanceNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNInstanceNormalizationNode() CNNInstanceNormalizationNode {
	return CNNInstanceNormalizationNodeClass.New()
}

func (c_ CNNInstanceNormalizationNode) Init() CNNInstanceNormalizationNode {
	rv := objc.Call[CNNInstanceNormalizationNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle?language=objc
func (c_ CNNInstanceNormalizationNode) TrainingStyle() NNTrainingStyle {
	rv := objc.Call[NNTrainingStyle](c_, objc.Sel("trainingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationnode/3197824-trainingstyle?language=objc
func (c_ CNNInstanceNormalizationNode) SetTrainingStyle(value NNTrainingStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setTrainingStyle:"), value)
}
