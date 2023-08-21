// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PredicateEditor] class.
var PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}

type _PredicateEditorClass struct {
	objc.Class
}

// An interface definition for the [PredicateEditor] class.
type IPredicateEditor interface {
	IRuleEditor
	RowTemplates() []PredicateEditorRowTemplate
	SetRowTemplates(value []IPredicateEditorRowTemplate)
}

// A defined set of rules that allows the editing of predicate objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor?language=objc
type PredicateEditor struct {
	RuleEditor
}

func PredicateEditorFrom(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: RuleEditorFrom(ptr),
	}
}

func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.Call[PredicateEditor](pc, objc.Sel("alloc"))
	return rv
}

func PredicateEditor_Alloc() PredicateEditor {
	return PredicateEditorClass.Alloc()
}

func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.Call[PredicateEditor](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPredicateEditor() PredicateEditor {
	return PredicateEditorClass.New()
}

func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.Call[PredicateEditor](p_, objc.Sel("init"))
	return rv
}

func (p_ PredicateEditor) InitWithFrame(frameRect foundation.Rect) PredicateEditor {
	rv := objc.Call[PredicateEditor](p_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewPredicateEditorWithFrame(frameRect foundation.Rect) PredicateEditor {
	instance := PredicateEditorClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// The row templates for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/1474094-rowtemplates?language=objc
func (p_ PredicateEditor) RowTemplates() []PredicateEditorRowTemplate {
	rv := objc.Call[[]PredicateEditorRowTemplate](p_, objc.Sel("rowTemplates"))
	return rv
}

// The row templates for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspredicateeditor/1474094-rowtemplates?language=objc
func (p_ PredicateEditor) SetRowTemplates(value []IPredicateEditorRowTemplate) {
	objc.Call[objc.Void](p_, objc.Sel("setRowTemplates:"), value)
}
