// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// A block that receives the item provider’s data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemprovidercompletionhandler?language=objc
type ItemProviderCompletionHandler = func(item objc.Object, error Error)

// Implement this block to retrieve the result of the AppleScript executed by executeWithAppleEvent:completionHandler:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserapplescripttaskcompletionhandler?language=objc
type UserAppleScriptTaskCompletionHandler = func(result AppleEventDescriptor, error Error)

// A block that the system calls when an observed progress object terminates the subscription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogressunpublishinghandler?language=objc
type ProgressUnpublishingHandler = func()

// A block that loads the item provider’s data and coerces it to the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsitemproviderloadhandler?language=objc
type ItemProviderLoadHandler = func(completionHandler ItemProviderCompletionHandler, expectedValueClass objc.Class, options Dictionary)

// A block that the system calls when an observed progress object matches the subscription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsprogresspublishinghandler?language=objc
type ProgressPublishingHandler = func(progress Progress) ProgressUnpublishingHandler

// Defines the signature for a block object used for comparison operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscomparator?language=objc
type Comparator = func(obj1 objc.Object, obj2 objc.Object) ComparisonResult

// Implement this block to retrieve the error of the script executed by executeWithCompletionHandler:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserscripttaskcompletionhandler?language=objc
type UserScriptTaskCompletionHandler = func(error Error)

// Implement this block to retrieve the output of the Automator workflow executed by executeWithInput:completionHandler:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserautomatortaskcompletionhandler?language=objc
type UserAutomatorTaskCompletionHandler = func(result objc.Object, error Error)

// Implement this block to retrieve an error from the Unix scripted executed by executeWithArguments:completionHandler:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuserunixtaskcompletionhandler?language=objc
type UserUnixTaskCompletionHandler = func(error Error)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsbackgroundactivitycompletionhandler?language=objc
type BackgroundActivityCompletionHandler = func(result BackgroundActivityResult)
