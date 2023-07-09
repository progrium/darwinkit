package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSView struct {
	gen_NSView
}

func NSView_Init(frame core.NSRect) NSView {
	return NSView_Alloc().InitWithFrame(frame)
}

func (v NSView) AddSubviewPositionedRelativeTo(subview NSViewRef, positioned int, relativeTo NSViewRef) {
	v.gen_NSView.AddSubviewPositionedRelativeTo(subview, core.NSUInteger(positioned), relativeTo)
}

func (v NSView) SetOpaque(b bool) {
	// TODO NSView declares this property as `readonly`, do we really need this
	// setter?
	v.Set("opaque:", b)
}

func (v NSView) Opaque() bool {
	return v.IsOpaque()
}

func (v NSView) SetValueForKey(value, key objc.Object) {
	v.Send("setValue:forKey:", value, key)
}
