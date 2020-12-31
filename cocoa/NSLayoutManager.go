package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSLayoutManager struct {
	objc.Object
}

func (lm NSLayoutManager) EnsureLayoutForTextContainer(tc NSTextContainer) {
	lm.Send("ensureLayoutForTextContainer:", tc)
}

func (lm NSLayoutManager) UsedRectForTextContainer(tc NSTextContainer) (rect core.NSRect) {
	lm.Send("usedRectForTextContainer:", tc, &rect)
	return rect
}
