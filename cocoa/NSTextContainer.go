package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSTextContainer struct {
	objc.Object
}

func (tc NSTextContainer) Size() (size core.NSSize) {
	tc.GetSt("size", &size)
	return size
}

func (tc NSTextContainer) HeightTracksTextView() bool {
	return tc.Get("heightTracksTextView").Bool()
}

func (tc NSTextContainer) SetHeightTracksTextView(b bool) {
	tc.Set("heightTracksTextView:", b)
}
