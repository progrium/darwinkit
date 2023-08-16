// AUTO-GENERATED CODE, DO NOT MODIFY

package coreml

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ParameterKey] class.
var ParameterKeyClass = _ParameterKeyClass{objc.GetClass("MLParameterKey")}

type _ParameterKeyClass struct {
	objc.Class
}

// An interface definition for the [ParameterKey] class.
type IParameterKey interface {
	IKey
	ScopedTo(scope string) ParameterKey
}

// The keys for the parameter dictionary in a model configuration or a model update context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey?language=objc
type ParameterKey struct {
	Key
}

func ParameterKeyFrom(ptr unsafe.Pointer) ParameterKey {
	return ParameterKey{
		Key: KeyFrom(ptr),
	}
}

func (pc _ParameterKeyClass) Alloc() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("alloc"))
	return rv
}

func ParameterKey_Alloc() ParameterKey {
	return ParameterKeyClass.Alloc()
}

func (pc _ParameterKeyClass) New() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewParameterKey() ParameterKey {
	return ParameterKeyClass.New()
}

func (p_ ParameterKey) Init() ParameterKey {
	rv := objc.Call[ParameterKey](p_, objc.Sel("init"))
	return rv
}

// Creates a copy of a parameter key and adds the scope to it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3333251-scopedto?language=objc
func (p_ ParameterKey) ScopedTo(scope string) ParameterKey {
	rv := objc.Call[ParameterKey](p_, objc.Sel("scopedTo:"), scope)
	return rv
}

// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3197646-momentum?language=objc
func (pc _ParameterKeyClass) Momentum() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("momentum"))
	return rv
}

// The key you use to access the stochastic gradient descent (SGD) optimizer’s momentum parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3197646-momentum?language=objc
func ParameterKey_Momentum() ParameterKey {
	return ParameterKeyClass.Momentum()
}

// The key you use to access the linked model’s filename. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3227893-linkedmodelfilename?language=objc
func (pc _ParameterKeyClass) LinkedModelFileName() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("linkedModelFileName"))
	return rv
}

// The key you use to access the linked model’s filename. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3227893-linkedmodelfilename?language=objc
func ParameterKey_LinkedModelFileName() ParameterKey {
	return ParameterKeyClass.LinkedModelFileName()
}

// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362529-numberofneighbors?language=objc
func (pc _ParameterKeyClass) NumberOfNeighbors() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("numberOfNeighbors"))
	return rv
}

// The key you use to access the number of neighbors that adjusts the affinity of a k-nearest-neighbor model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362529-numberofneighbors?language=objc
func ParameterKey_NumberOfNeighbors() ParameterKey {
	return ParameterKeyClass.NumberOfNeighbors()
}

// The key you use to access the weights of a layer in a neural network model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362530-weights?language=objc
func (pc _ParameterKeyClass) Weights() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("weights"))
	return rv
}

// The key you use to access the weights of a layer in a neural network model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362530-weights?language=objc
func ParameterKey_Weights() ParameterKey {
	return ParameterKeyClass.Weights()
}

// The key you use to access the Adam optimizer’s second beta parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194197-beta2?language=objc
func (pc _ParameterKeyClass) Beta2() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("beta2"))
	return rv
}

// The key you use to access the Adam optimizer’s second beta parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194197-beta2?language=objc
func ParameterKey_Beta2() ParameterKey {
	return ParameterKeyClass.Beta2()
}

// The key you use to access the Adam optimizer’s first beta parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194196-beta1?language=objc
func (pc _ParameterKeyClass) Beta1() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("beta1"))
	return rv
}

// The key you use to access the Adam optimizer’s first beta parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194196-beta1?language=objc
func ParameterKey_Beta1() ParameterKey {
	return ParameterKeyClass.Beta1()
}

// The key you use to access the biases of a layer in a neural network model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362528-biases?language=objc
func (pc _ParameterKeyClass) Biases() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("biases"))
	return rv
}

// The key you use to access the biases of a layer in a neural network model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3362528-biases?language=objc
func ParameterKey_Biases() ParameterKey {
	return ParameterKeyClass.Biases()
}

// The key you use to access the optimizer’s learning rate parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180073-learningrate?language=objc
func (pc _ParameterKeyClass) LearningRate() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("learningRate"))
	return rv
}

// The key you use to access the optimizer’s learning rate parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180073-learningrate?language=objc
func ParameterKey_LearningRate() ParameterKey {
	return ParameterKeyClass.LearningRate()
}

// The key you use to access the Adam optimizer’s epsilon parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194198-eps?language=objc
func (pc _ParameterKeyClass) Eps() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("eps"))
	return rv
}

// The key you use to access the Adam optimizer’s epsilon parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3194198-eps?language=objc
func ParameterKey_Eps() ParameterKey {
	return ParameterKeyClass.Eps()
}

// The key you use to access the linked model’s search path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3227894-linkedmodelsearchpath?language=objc
func (pc _ParameterKeyClass) LinkedModelSearchPath() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("linkedModelSearchPath"))
	return rv
}

// The key you use to access the linked model’s search path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3227894-linkedmodelsearchpath?language=objc
func ParameterKey_LinkedModelSearchPath() ParameterKey {
	return ParameterKeyClass.LinkedModelSearchPath()
}

// The key you use to access the optimizer’s mini batch-size parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180074-minibatchsize?language=objc
func (pc _ParameterKeyClass) MiniBatchSize() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("miniBatchSize"))
	return rv
}

// The key you use to access the optimizer’s mini batch-size parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180074-minibatchsize?language=objc
func ParameterKey_MiniBatchSize() ParameterKey {
	return ParameterKeyClass.MiniBatchSize()
}

// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3333253-shuffle?language=objc
func (pc _ParameterKeyClass) Shuffle() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("shuffle"))
	return rv
}

// The key you use to access the shuffle parameter, a Boolean value that determines whether the model randomizes the data between epochs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3333253-shuffle?language=objc
func ParameterKey_Shuffle() ParameterKey {
	return ParameterKeyClass.Shuffle()
}

// The key you use to access the seed parameter that initializes the random number generator for the shuffle option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3333252-seed?language=objc
func (pc _ParameterKeyClass) Seed() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("seed"))
	return rv
}

// The key you use to access the seed parameter that initializes the random number generator for the shuffle option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3333252-seed?language=objc
func ParameterKey_Seed() ParameterKey {
	return ParameterKeyClass.Seed()
}

// The key you use to access the optimizer’s epochs parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180072-epochs?language=objc
func (pc _ParameterKeyClass) Epochs() ParameterKey {
	rv := objc.Call[ParameterKey](pc, objc.Sel("epochs"))
	return rv
}

// The key you use to access the optimizer’s epochs parameter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlparameterkey/3180072-epochs?language=objc
func ParameterKey_Epochs() ParameterKey {
	return ParameterKeyClass.Epochs()
}
