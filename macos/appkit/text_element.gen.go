// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextElement] class.
var TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	objc.IObject
	ElementRange() TextRange
	SetElementRange(value ITextRange)
	TextContentManager() TextContentManager
	SetTextContentManager(value ITextContentManager)
}

// An abstract base class that represents the smallest units of text layout such as paragraphs or attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement?language=objc
type TextElement struct {
	objc.Object
}

func TextElementFrom(ptr unsafe.Pointer) TextElement {
	return TextElement{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextElement) InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	rv := objc.Call[TextElement](t_, objc.Sel("initWithTextContentManager:"), objc.Ptr(textContentManager))
	return rv
}

// Creates a new text element with the content manager you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809955-initwithtextcontentmanager?language=objc
func TextElement_InitWithTextContentManager(textContentManager ITextContentManager) TextElement {
	return TextElementClass.Alloc().InitWithTextContentManager(textContentManager)
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Call[TextElement](tc, objc.Sel("alloc"))
	return rv
}

func TextElement_Alloc() TextElement {
	return TextElementClass.Alloc()
}

func (tc _TextElementClass) New() TextElement {
	rv := objc.Call[TextElement](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextElement() TextElement {
	return TextElementClass.New()
}

func (t_ TextElement) Init() TextElement {
	rv := objc.Call[TextElement](t_, objc.Sel("init"))
	return rv
}

// A range value that represents the range of the element inside the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809954-elementrange?language=objc
func (t_ TextElement) ElementRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("elementRange"))
	return rv
}

// A range value that represents the range of the element inside the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809954-elementrange?language=objc
func (t_ TextElement) SetElementRange(value ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("setElementRange:"), objc.Ptr(value))
}

// The value that represents the current content manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809956-textcontentmanager?language=objc
func (t_ TextElement) TextContentManager() TextContentManager {
	rv := objc.Call[TextContentManager](t_, objc.Sel("textContentManager"))
	return rv
}

// The value that represents the current content manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809956-textcontentmanager?language=objc
func (t_ TextElement) SetTextContentManager(value ITextContentManager) {
	objc.Call[objc.Void](t_, objc.Sel("setTextContentManager:"), objc.Ptr(value))
}
