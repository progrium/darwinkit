// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationNode] class.
var CNNBatchNormalizationNodeClass = _CNNBatchNormalizationNodeClass{objc.GetClass("MPSCNNBatchNormalizationNode")}

type _CNNBatchNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationNode] class.
type ICNNBatchNormalizationNode interface {
	INNFilterNode
	TrainingStyle() NNTrainingStyle
	SetTrainingStyle(value NNTrainingStyle)
	Flags() CNNBatchNormalizationFlags
	SetFlags(value CNNBatchNormalizationFlags)
}

// A representation of a batch normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode?language=objc
type CNNBatchNormalizationNode struct {
	NNFilterNode
}

func CNNBatchNormalizationNodeFrom(ptr unsafe.Pointer) CNNBatchNormalizationNode {
	return CNNBatchNormalizationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNBatchNormalizationNodeClass) NodeWithSourceDataSource(source INNImageNode, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNBatchNormalizationDataSource", dataSource)
	rv := objc.Call[CNNBatchNormalizationNode](cc, objc.Sel("nodeWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948033-nodewithsource?language=objc
func CNNBatchNormalizationNode_NodeWithSourceDataSource(source INNImageNode, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalizationNode {
	return CNNBatchNormalizationNodeClass.NodeWithSourceDataSource(source, dataSource)
}

func (c_ CNNBatchNormalizationNode) InitWithSourceDataSource(source INNImageNode, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalizationNode {
	po1 := objc.WrapAsProtocol("MPSCNNBatchNormalizationDataSource", dataSource)
	rv := objc.Call[CNNBatchNormalizationNode](c_, objc.Sel("initWithSource:dataSource:"), objc.Ptr(source), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2948004-initwithsource?language=objc
func NewCNNBatchNormalizationNodeWithSourceDataSource(source INNImageNode, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalizationNode {
	instance := CNNBatchNormalizationNodeClass.Alloc().InitWithSourceDataSource(source, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationNodeClass) Alloc() CNNBatchNormalizationNode {
	rv := objc.Call[CNNBatchNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationNode_Alloc() CNNBatchNormalizationNode {
	return CNNBatchNormalizationNodeClass.Alloc()
}

func (cc _CNNBatchNormalizationNodeClass) New() CNNBatchNormalizationNode {
	rv := objc.Call[CNNBatchNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationNode() CNNBatchNormalizationNode {
	return CNNBatchNormalizationNodeClass.New()
}

func (c_ CNNBatchNormalizationNode) Init() CNNBatchNormalizationNode {
	rv := objc.Call[CNNBatchNormalizationNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle?language=objc
func (c_ CNNBatchNormalizationNode) TrainingStyle() NNTrainingStyle {
	rv := objc.Call[NNTrainingStyle](c_, objc.Sel("trainingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/3197821-trainingstyle?language=objc
func (c_ CNNBatchNormalizationNode) SetTrainingStyle(value NNTrainingStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setTrainingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags?language=objc
func (c_ CNNBatchNormalizationNode) Flags() CNNBatchNormalizationFlags {
	rv := objc.Call[CNNBatchNormalizationFlags](c_, objc.Sel("flags"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationnode/2953940-flags?language=objc
func (c_ CNNBatchNormalizationNode) SetFlags(value CNNBatchNormalizationFlags) {
	objc.Call[objc.Void](c_, objc.Sel("setFlags:"), value)
}
