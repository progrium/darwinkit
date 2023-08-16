// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
)

// A completion handler signature a method calls when it finishes creating a compute pipeline and reflection information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlnewcomputepipelinestatewithreflectioncompletionhandler?language=objc
type NewComputePipelineStateWithReflectionCompletionHandler = func(computePipelineState ComputePipelineStateWrapper, reflection ComputePipelineReflection, error foundation.Error)

// A completion handler signature a method calls when it finishes creating a render pipeline and reflection information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlnewrenderpipelinestatewithreflectioncompletionhandler?language=objc
type NewRenderPipelineStateWithReflectionCompletionHandler = func(renderPipelineState RenderPipelineStateWrapper, reflection RenderPipelineReflection, error foundation.Error)

// A block of code invoked after a shareable event’s signal value equals or exceeds a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedeventnotificationblock?language=objc
type SharedEventNotificationBlock = func(arg0 SharedEventWrapper, value uint64)

// A completion handler signature a method calls when it finishes creating a render pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlnewrenderpipelinestatecompletionhandler?language=objc
type NewRenderPipelineStateCompletionHandler = func(renderPipelineState RenderPipelineStateWrapper, error foundation.Error)

// A completion handler signature a method calls when it finishes creating a compute pipeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlnewcomputepipelinestatecompletionhandler?language=objc
type NewComputePipelineStateCompletionHandler = func(computePipelineState ComputePipelineStateWrapper, error foundation.Error)

// A completion handler signature a GPU device calls when it finishes scheduling a command buffer, or when the GPU finishes running it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbufferhandler?language=objc
type CommandBufferHandler = func(arg0 CommandBufferWrapper)

// A completion handler signature a method calls when it finishes creating a Metal library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlnewlibrarycompletionhandler?language=objc
type NewLibraryCompletionHandler = func(library LibraryWrapper, error foundation.Error)

// A block of code invoked after a drawable is presented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldrawablepresentedhandler?language=objc
type DrawablePresentedHandler = func(arg0 DrawableWrapper)

// A block that Metal calls when the system adds or removes a GPU device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldevicenotificationhandler?language=objc
type DeviceNotificationHandler = func(device DeviceWrapper, notifyName DeviceNotificationName)
