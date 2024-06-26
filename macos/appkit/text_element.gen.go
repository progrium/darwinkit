// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TextElement] class.
var TextElementClass = _TextElementClass{objc.GetClass("NSTextElement")}

type _TextElementClass struct {
	objc.Class
}

// An interface definition for the [TextElement] class.
type ITextElement interface {
	objc.IObject
	TextContentManager() TextContentManager
	SetTextContentManager(value ITextContentManager)
	ElementRange() TextRange
	SetElementRange(value ITextRange)
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
	rv := objc.Call[TextElement](t_, objc.Sel("initWithTextContentManager:"), textContentManager)
	return rv
}

// Creates a new text element with the content manager you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809955-initwithtextcontentmanager?language=objc
func NewTextElementWithTextContentManager(textContentManager ITextContentManager) TextElement {
	instance := TextElementClass.Alloc().InitWithTextContentManager(textContentManager)
	instance.Autorelease()
	return instance
}

func (tc _TextElementClass) Alloc() TextElement {
	rv := objc.Call[TextElement](tc, objc.Sel("alloc"))
	return rv
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
	objc.Call[objc.Void](t_, objc.Sel("setTextContentManager:"), value)
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
	objc.Call[objc.Void](t_, objc.Sel("setElementRange:"), value)
}
