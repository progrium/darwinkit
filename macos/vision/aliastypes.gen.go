// AUTO-GENERATED CODE, DO NOT MODIFY

package vision

import (
	"github.com/progrium/macdriver/macos/foundation"
)

// A block executed at intervals during the processing of a Vision request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestprogresshandler?language=objc
type RequestProgressHandler = func(request Request, fractionCompleted float64, error foundation.Error)

// A type alias to encapsulate the syntax for the completion handler the system calls after the request finishes processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequestcompletionhandler?language=objc
type RequestCompletionHandler = func(request Request, error foundation.Error)
