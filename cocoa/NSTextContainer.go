package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSTextContainer struct {
	gen_NSTextContainer
}

func (tc NSTextContainer) Size() (size core.NSSize) {
	tc.GetSt("size", &size)
	return size
}

func (tc NSTextContainer) SetHeightTracksTextView(b bool) {
	tc.SetHeightTracksTextView_(b)
}
