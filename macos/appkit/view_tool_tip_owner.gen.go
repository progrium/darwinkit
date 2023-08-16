// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods for dynamically associating a tool tip with a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewtooltipowner?language=objc
type PViewToolTipOwner interface {
	// optional
	ViewStringForToolTipPointUserData(view View, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string
	HasViewStringForToolTipPointUserData() bool
}

// A concrete type wrapper for the [PViewToolTipOwner] protocol.
type ViewToolTipOwnerWrapper struct {
	objc.Object
}

func (v_ ViewToolTipOwnerWrapper) HasViewStringForToolTipPointUserData() bool {
	return v_.RespondsToSelector(objc.Sel("view:stringForToolTip:point:userData:"))
}

// Returns the tool tip string to be displayed due to the cursor pausing at location point within the tool tip rectangle identified by tag in the view view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewtooltipowner/3005296-view?language=objc
func (v_ ViewToolTipOwnerWrapper) ViewStringForToolTipPointUserData(view IView, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string {
	rv := objc.Call[string](v_, objc.Sel("view:stringForToolTip:point:userData:"), objc.Ptr(view), tag, point, data)
	return rv
}
