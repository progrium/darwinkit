// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A completion handler for getting an asynchronous attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/nsattributedstringcompletionhandler?language=objc
type AttributedStringCompletionHandler = func(arg0 foundation.AttributedString, arg1 map[appkit.AttributedStringDocumentAttributeKey]objc.Object, arg2 foundation.Error)
