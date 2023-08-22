// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
)

// A type alias for a closure that provides the result of an image generation request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetimagegeneratorcompletionhandler?language=objc
type AssetImageGeneratorCompletionHandler = func(requestedTime coremedia.Time, image coregraphics.ImageRef, actualTime coremedia.Time, result AssetImageGeneratorResult, error foundation.Error)
