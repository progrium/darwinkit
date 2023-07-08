package cocoa

import (
	"github.com/progrium/macdriver/core"
)

type NSTextAlignment uint

const (
	NSTextAlignmentLeft      NSTextAlignment = 0
	NSTextAlignmentRight     NSTextAlignment = 1
	NSTextAlignmentCenter    NSTextAlignment = 2
	NSTextAlignmentJustified NSTextAlignment = 3
	NSTextAlignmentNatural   NSTextAlignment = 4
)

type NSTextView struct {
	gen_NSTextView
}

func NSTextView_Init(frame core.NSRect) NSTextView {
	return NSTextView_Alloc().InitWithFrame(frame)
}

func (v NSTextView) String() string {
	return v.gen_NSTextView.String().String()
}

func (v NSTextView) SetString(s string) {
	v.gen_NSTextView.SetString(core.String(s))
}

func (v NSTextView) Selectable() bool {
	return v.IsSelectable()
}

func (v NSTextView) RichText() bool {
	return v.IsRichText()
}

func (v NSTextView) Editable() bool {
	return v.IsEditable()
}

func (v NSTextView) FieldEditor() bool {
	return v.IsFieldEditor()
}

// TODO(mgood): once `NSTextAlignment` is defined in the schema we can use the
// generated methods on the `NSText` parent class here
func (v NSTextView) Alignment() NSTextAlignment {
	return NSTextAlignment(v.Get("alignment").Int())
}

func (v NSTextView) SetAlignment(d NSTextAlignment) {
	v.Set("alignment:", d)
}
