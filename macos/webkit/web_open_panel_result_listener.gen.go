// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// WebView user interface delegates that implement the webView:runOpenPanelForFileButtonWithResultListener: method use the methods defined in this protocol to communicate with the listener object. The methods allow the delegate to send a cancel message, or set the selected file name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webopenpanelresultlistener?language=objc
type PWebOpenPanelResultListener interface {
}

// A concrete type wrapper for the [PWebOpenPanelResultListener] protocol.
type WebOpenPanelResultListenerWrapper struct {
	objc.Object
}
