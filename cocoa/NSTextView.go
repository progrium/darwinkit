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
	return NSTextView_alloc().InitWithFrame__asNSTextView(frame)
}

func (v NSTextView) String() string {
	return v.gen_NSTextView.String().String()
}

func (v NSTextView) SetString(s string) {
	v.SetString_(core.String(s))
}

func (v NSTextView) Selectable() bool {
	return v.IsSelectable()
}

func (v NSTextView) SetSelectable(b bool) {
	v.SetSelectable_(b)
}

func (v NSTextView) RichText() bool {
	return v.IsRichText()
}

func (v NSTextView) SetRichText(b bool) {
	v.SetRichText_(b)
}

func (v NSTextView) Editable() bool {
	return v.IsEditable()
}

func (v NSTextView) SetEditable(b bool) {
	v.SetEditable_(b)
}

func (v NSTextView) FieldEditor() bool {
	return v.IsFieldEditor()
}

func (v NSTextView) SetFieldEditor(b bool) {
	v.SetFieldEditor_(b)
}

func (v NSTextView) SetImportsGraphics(b bool) {
	v.SetImportsGraphics_(b)
}

func (v NSTextView) SetDrawsBackground(b bool) {
	v.SetDrawsBackground_(b)
}

func (v NSTextView) SetFont(f NSFont) {
	v.SetFont_(f)
}

func (v NSTextView) SetTextColor(c NSColorRef) {
	v.SetTextColor_(c)
}

// TODO(mgood): once `NSTextAlignment` is defined in the schema we can use the
// generated methods on the `NSText` parent class here
func (v NSTextView) Alignment() NSTextAlignment {
	return NSTextAlignment(v.Get("alignment").Int())
}

func (v NSTextView) SetAlignment(d NSTextAlignment) {
	v.Set("alignment:", d)
}

func (v NSTextView) SetTextContainer(tc NSTextContainer) {
	v.SetTextContainer_(tc)
}
