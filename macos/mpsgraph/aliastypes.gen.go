// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"github.com/progrium/macdriver/macos/foundation"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphscheduledhandler?language=objc
type ScheduledHandler = func(resultsDictionary *foundation.Dictionary, error foundation.Error)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutablescheduledhandler?language=objc
type ExecutableScheduledHandler = func(results []TensorData, error foundation.Error)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcontrolflowdependencyblock?language=objc
type ControlFlowDependencyBlock = func() []Tensor

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphwhilebeforeblock?language=objc
type WhileBeforeBlock = func(inputTensors []Tensor, resultTensors foundation.MutableArray) Tensor

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphifthenelseblock?language=objc
type IfThenElseBlock = func() []Tensor

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphwhileafterblock?language=objc
type WhileAfterBlock = func(bodyBlockArguments []Tensor) []Tensor

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcompletionhandler?language=objc
type CompletionHandler = func(resultsDictionary *foundation.Dictionary, error foundation.Error)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphexecutablecompletionhandler?language=objc
type ExecutableCompletionHandler = func(results []TensorData, error foundation.Error)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphforloopbodyblock?language=objc
type ForLoopBodyBlock = func(index Tensor, iterationArguments []Tensor) []Tensor
