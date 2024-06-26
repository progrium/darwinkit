// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A completion handler for getting an asynchronous attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/nsattributedstringcompletionhandler?language=objc
type AttributedStringCompletionHandler = func(arg0 foundation.AttributedString, arg1 map[appkit.AttributedStringDocumentAttributeKey]objc.Object, arg2 foundation.Error)
