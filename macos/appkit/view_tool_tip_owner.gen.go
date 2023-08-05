// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IViewToolTipOwner interface {
	// required
	ViewStringForToolTipPointUserData(view View, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string
}

type ViewToolTipOwnerWrapper struct {
	objc.Object
}

func (v_ ViewToolTipOwnerWrapper) ViewStringForToolTipPointUserData(view IView, tag ToolTipTag, point foundation.Point, data unsafe.Pointer) string {
	rv := objc.CallMethod[string](v_, objc.GetSelector("view:stringForToolTip:point:userData:"), objc.ExtractPtr(view), tag, point, data)
	return rv
}
