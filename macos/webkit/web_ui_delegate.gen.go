// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// Web view user interface delegates implement this protocol to control the opening of new windows, augment the behavior of default menu items displayed when the user clicks elements, and perform other user interface–related tasks. These methods can be invoked as a result of handling JavaScript or other plug-in content. Delegates that display more than one web view per window, for example, need to implement some of these methods to handle that case. The default implementation assumes one window per web view, so non-conventional user interfaces might implement a user interface delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webuidelegate?language=objc
type PWebUIDelegate interface {
}

// A delegate implementation builder for the [PWebUIDelegate] protocol.
type WebUIDelegate struct {
}

// A concrete type wrapper for the [PWebUIDelegate] protocol.
type WebUIDelegateWrapper struct {
	objc.Object
}
