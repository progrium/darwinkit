// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PredicateEditorClass = _PredicateEditorClass{objc.GetClass("NSPredicateEditor")}

type _PredicateEditorClass struct {
	objc.Class
}

type IPredicateEditor interface {
	IRuleEditor
	RowTemplates() []PredicateEditorRowTemplate
	SetRowTemplates(value []IPredicateEditorRowTemplate)
}

type PredicateEditor struct {
	RuleEditor
}

func MakePredicateEditor(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		RuleEditor: MakeRuleEditor(ptr),
	}
}

func (p_ PredicateEditor) InitWithFrame(frameRect foundation.Rect) PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](p_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func PredicateEditor_InitWithFrame(frameRect foundation.Rect) PredicateEditor {
	return PredicateEditorClass.Alloc().InitWithFrame(frameRect)
}

func (p_ PredicateEditor) Init() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](p_, objc.GetSelector("init"))
	return rv
}

func PredicateEditor_Init() PredicateEditor {
	return PredicateEditorClass.Alloc().Init()
}

func (pc _PredicateEditorClass) Alloc() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](pc, objc.GetSelector("alloc"))
	return rv
}

func PredicateEditor_Alloc() PredicateEditor {
	return PredicateEditorClass.Alloc()
}

func (pc _PredicateEditorClass) New() PredicateEditor {
	rv := objc.CallMethod[PredicateEditor](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPredicateEditor() PredicateEditor {
	return PredicateEditorClass.New()
}

func PredicateEditor_New() PredicateEditor {
	return PredicateEditorClass.New()
}

func (p_ PredicateEditor) RowTemplates() []PredicateEditorRowTemplate {
	rv := objc.CallMethod[[]PredicateEditorRowTemplate](p_, objc.GetSelector("rowTemplates"))
	// TODO: convert slice items...
	return rv
}

func (p_ PredicateEditor) SetRowTemplates(value []IPredicateEditorRowTemplate) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setRowTemplates:"), value)
}
