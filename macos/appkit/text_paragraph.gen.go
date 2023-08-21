// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextParagraph] class.
var TextParagraphClass = _TextParagraphClass{objc.GetClass("NSTextParagraph")}

type _TextParagraphClass struct {
	objc.Class
}

// An interface definition for the [TextParagraph] class.
type ITextParagraph interface {
	ITextElement
	AttributedString() foundation.AttributedString
	ParagraphContentRange() TextRange
	ParagraphSeparatorRange() TextRange
}

// A class that represents a single paragraph backed by an attributed string as the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextparagraph?language=objc
type TextParagraph struct {
	TextElement
}

func TextParagraphFrom(ptr unsafe.Pointer) TextParagraph {
	return TextParagraph{
		TextElement: TextElementFrom(ptr),
	}
}

func (t_ TextParagraph) InitWithAttributedString(attributedString foundation.IAttributedString) TextParagraph {
	rv := objc.Call[TextParagraph](t_, objc.Sel("initWithAttributedString:"), objc.Ptr(attributedString))
	return rv
}

// Creates a new paragraph with the attributed string you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextparagraph/3809959-initwithattributedstring?language=objc
func NewTextParagraphWithAttributedString(attributedString foundation.IAttributedString) TextParagraph {
	instance := TextParagraphClass.Alloc().InitWithAttributedString(attributedString)
	instance.Autorelease()
	return instance
}

func (tc _TextParagraphClass) Alloc() TextParagraph {
	rv := objc.Call[TextParagraph](tc, objc.Sel("alloc"))
	return rv
}

func TextParagraph_Alloc() TextParagraph {
	return TextParagraphClass.Alloc()
}

func (tc _TextParagraphClass) New() TextParagraph {
	rv := objc.Call[TextParagraph](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextParagraph() TextParagraph {
	return TextParagraphClass.New()
}

func (t_ TextParagraph) Init() TextParagraph {
	rv := objc.Call[TextParagraph](t_, objc.Sel("init"))
	return rv
}

func (t_ TextParagraph) InitWithTextContentManager(textContentManager ITextContentManager) TextParagraph {
	rv := objc.Call[TextParagraph](t_, objc.Sel("initWithTextContentManager:"), objc.Ptr(textContentManager))
	return rv
}

// Creates a new text element with the content manager you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextelement/3809955-initwithtextcontentmanager?language=objc
func NewTextParagraphWithTextContentManager(textContentManager ITextContentManager) TextParagraph {
	instance := TextParagraphClass.Alloc().InitWithTextContentManager(textContentManager)
	instance.Autorelease()
	return instance
}

// Returns the source attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextparagraph/3809958-attributedstring?language=objc
func (t_ TextParagraph) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedString"))
	return rv
}

// Returns the range of the paragraph in the attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextparagraph/3809960-paragraphcontentrange?language=objc
func (t_ TextParagraph) ParagraphContentRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("paragraphContentRange"))
	return rv
}

// Returns the range of the paragraph separator in the attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextparagraph/3809961-paragraphseparatorrange?language=objc
func (t_ TextParagraph) ParagraphSeparatorRange() TextRange {
	rv := objc.Call[TextRange](t_, objc.Sel("paragraphSeparatorRange"))
	return rv
}
