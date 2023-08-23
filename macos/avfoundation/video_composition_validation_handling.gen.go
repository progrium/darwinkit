// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// Methods you can implement to indicate whether validation of a video composition should continue after specific errors are found. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionvalidationhandling?language=objc
type PVideoCompositionValidationHandling interface {
	// optional
	VideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange(videoComposition VideoComposition, timeRange coremedia.TimeRange) bool
	HasVideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange() bool
}

// A concrete type wrapper for the [PVideoCompositionValidationHandling] protocol.
type VideoCompositionValidationHandlingWrapper struct {
	objc.Object
}

func (v_ VideoCompositionValidationHandlingWrapper) HasVideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange() bool {
	return v_.RespondsToSelector(objc.Sel("videoComposition:shouldContinueValidatingAfterFindingEmptyTimeRange:"))
}

// Reports a time range that has no corresponding video composition instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositionvalidationhandling/1388620-videocomposition?language=objc
func (v_ VideoCompositionValidationHandlingWrapper) VideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange(videoComposition IVideoComposition, timeRange coremedia.TimeRange) bool {
	rv := objc.Call[bool](v_, objc.Sel("videoComposition:shouldContinueValidatingAfterFindingEmptyTimeRange:"), objc.Ptr(videoComposition), timeRange)
	return rv
}
