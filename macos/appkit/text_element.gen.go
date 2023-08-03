// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	objc.Class
}

type ITextElement interface {
	objc.IObject
	TextContentManager() TextContentManager
	SetTextContentManager(value ITextContentManager)
	ElementRange() TextRange
	SetElementRange(value ITextRange)
	IsRepresentedElement() bool
	ParentElement() TextElement
	ChildElements() []TextElement
}

type TextElement struct {
	objc.Object
}

func MakeTextElement(ptr unsafe.Pointer) TextElement {
	return TextElement{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextElement) InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("initWithTextContentManager:"), objc.ExtractPtr(textContentManager))
	return rv
}

func TextElement_InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	return TextElementClass.Alloc().InitWithTextContentManager(textContentManager)
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.GetSelector("alloc"))
	return rv
}

func TextElement_Alloc() TextElement {
	return TextElementClass.Alloc()
}

func (tc _TextElementClass) New() TextElement {
	rv := objc.CallMethod[TextElement](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func TextElement_New() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) Init() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("init"))
	return rv
}

func TextElement_Init() TextElement {
	return TextElementClass.Alloc().Init()
}

func (t_ TextElement) TextContentManager() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.GetSelector("textContentManager"))
	return rv
}

func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextContentManager:"), objc.ExtractPtr(value))
}

func (t_ TextElement) ElementRange() TextRange {
	rv := objc.CallMethod[TextRange](t_, objc.GetSelector("elementRange"))
	return rv
}

func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setElementRange:"), objc.ExtractPtr(value))
}

func (t_ TextElement) IsRepresentedElement() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isRepresentedElement"))
	return rv
}

func (t_ TextElement) ParentElement() TextElement {
	rv := objc.CallMethod[TextElement](t_, objc.GetSelector("parentElement"))
	return rv
}

func (t_ TextElement) ChildElements() []TextElement {
	rv := objc.CallMethod[[]TextElement](t_, objc.GetSelector("childElements"))
	// TODO: convert slice items...
	return rv
}
