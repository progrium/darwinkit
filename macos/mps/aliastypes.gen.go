// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/foundation"
)

// A notification when an asynchronous graph execution has finished. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraphcompletionhandler?language=objc
type NNGraphCompletionHandler = func(result Image, error foundation.Error)

// A block of code that's invoked when an operation on an acceleration structure has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructurecompletionhandler?language=objc
type AccelerationStructureCompletionHandler = func(arg0 AccelerationStructure)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgradientnodeblock?language=objc
type GradientNodeBlock = func(gradientNode NNFilterNode, inferenceNode NNFilterNode, inferenceSource NNImageNode, gradientSource NNImageNode)
