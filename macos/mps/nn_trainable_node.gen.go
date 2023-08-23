// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines methods that determine whether and when neural network training parameters are updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainablenode?language=objc
type PNNTrainableNode interface {
	// optional
	SetTrainingStyle(value NNTrainingStyle)
	HasSetTrainingStyle() bool

	// optional
	TrainingStyle() NNTrainingStyle
	HasTrainingStyle() bool
}

// A concrete type wrapper for the [PNNTrainableNode] protocol.
type NNTrainableNodeWrapper struct {
	objc.Object
}

func (n_ NNTrainableNodeWrapper) HasSetTrainingStyle() bool {
	return n_.RespondsToSelector(objc.Sel("setTrainingStyle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainablenode/2952971-trainingstyle?language=objc
func (n_ NNTrainableNodeWrapper) SetTrainingStyle(value NNTrainingStyle) {
	objc.Call[objc.Void](n_, objc.Sel("setTrainingStyle:"), value)
}

func (n_ NNTrainableNodeWrapper) HasTrainingStyle() bool {
	return n_.RespondsToSelector(objc.Sel("trainingStyle"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnntrainablenode/2952971-trainingstyle?language=objc
func (n_ NNTrainableNodeWrapper) TrainingStyle() NNTrainingStyle {
	rv := objc.Call[NNTrainingStyle](n_, objc.Sel("trainingStyle"))
	return rv
}
