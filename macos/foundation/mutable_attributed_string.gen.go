// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableAttributedString] class.
var MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}

type _MutableAttributedStringClass struct {
	objc.Class
}

// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	IAttributedString
	AppendLocalizedFormat(format IAttributedString)
	AppendAttributedString(attrString IAttributedString)
	DeleteCharactersInRange(range_ Range)
	SetAlignmentRange(alignment objc.IObject, range_ Range)
	InsertAttributedStringAtIndex(attrString IAttributedString, loc uint)
	FixFontAttributeInRange(range_ Range)
	ApplyFontTraitsRange(traitMask objc.IObject, range_ Range)
	UnscriptRange(range_ Range)
	BeginEditing()
	ReadFromDataOptionsDocumentAttributesError(data []byte, opts Dictionary, dict Dictionary, error IError) bool
	FixAttributesInRange(range_ Range)
	AddAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	UpdateAttachmentsFromPath(path string)
	ReplaceCharactersInRangeWithAttributedString(range_ Range, attrString IAttributedString)
	EndEditing()
	SetAttributedString(attrString IAttributedString)
	RemoveAttributeRange(name AttributedStringKey, range_ Range)
	AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ Range)
	SetAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range)
	SuperscriptRange(range_ Range)
	FixParagraphStyleAttributeInRange(range_ Range)
	SubscriptRange(range_ Range)
	SetBaseWritingDirectionRange(writingDirection objc.IObject, range_ Range)
	FixAttachmentAttributeInRange(range_ Range)
	MutableString() MutableString
}

// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring?language=objc
type MutableAttributedString struct {
	AttributedString
}

func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: AttributedStringFrom(ptr),
	}
}

func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Call[MutableAttributedString](mc, objc.Sel("alloc"))
	return rv
}

func MutableAttributedString_Alloc() MutableAttributedString {
	return MutableAttributedStringClass.Alloc()
}

func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.Call[MutableAttributedString](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAttributedString() MutableAttributedString {
	return MutableAttributedStringClass.New()
}

func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("init"))
	return rv
}

func (m_ MutableAttributedString) InitWithHTMLOptionsDocumentAttributes(data []byte, options Dictionary, dict Dictionary) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	return rv
}

// Creates an attributed string from the HTML in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1535412-initwithhtml?language=objc
func MutableAttributedString_InitWithHTMLOptionsDocumentAttributes(data []byte, options Dictionary, dict Dictionary) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithHTMLOptionsDocumentAttributes(data, options, dict)
}

func (mc _MutableAttributedStringClass) LocalizedAttributedStringWithFormat(format IAttributedString) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](mc, objc.Sel("localizedAttributedStringWithFormat:"), objc.Ptr(format))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746877-localizedattributedstringwithfor?language=objc
func MutableAttributedString_LocalizedAttributedStringWithFormat(format IAttributedString) MutableAttributedString {
	return MutableAttributedStringClass.LocalizedAttributedStringWithFormat(format)
}

func (m_ MutableAttributedString) InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithAttributedString:"), objc.Ptr(attrStr))
	return rv
}

// Creates an attributed string with the characters and attributes of the specified attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1415342-initwithattributedstring?language=objc
func MutableAttributedString_InitWithAttributedString(attrStr IAttributedString) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithAttributedString(attrStr)
}

func (m_ MutableAttributedString) InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile IURL, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), objc.Ptr(markdownFile), objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from the contents of a specified URL that contains Markdown-formatted data using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746872-initwithcontentsofmarkdownfileat?language=objc
func MutableAttributedString_InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile IURL, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile, options, baseURL, error)
}

func (m_ MutableAttributedString) InitWithFormatOptionsLocale(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithFormat:options:locale:"), objc.Ptr(format), options, objc.Ptr(locale))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746873-initwithformat?language=objc
func MutableAttributedString_InitWithFormatOptionsLocale(format IAttributedString, options AttributedStringFormattingOptions, locale ILocale) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithFormatOptionsLocale(format, options, locale)
}

func (m_ MutableAttributedString) InitWithDataOptionsDocumentAttributesError(data []byte, options Dictionary, dict Dictionary, error IError) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, objc.Ptr(error))
	return rv
}

// Creates an attributed string from the data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1524613-initwithdata?language=objc
func MutableAttributedString_InitWithDataOptionsDocumentAttributesError(data []byte, options Dictionary, dict Dictionary, error IError) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithDataOptionsDocumentAttributesError(data, options, dict, error)
}

func (m_ MutableAttributedString) InitWithString(str string) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithString:"), str)
	return rv
}

// Creates an attributed string with the characters of the specified string and no attribute information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1407481-initwithstring?language=objc
func MutableAttributedString_InitWithString(str string) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithString(str)
}

func (m_ MutableAttributedString) InitWithMarkdownStringOptionsBaseURLError(markdownString string, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithMarkdownString:options:baseURL:error:"), markdownString, objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from a Markdown-formatted string using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746876-initwithmarkdownstring?language=objc
func MutableAttributedString_InitWithMarkdownStringOptionsBaseURLError(markdownString string, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithMarkdownStringOptionsBaseURLError(markdownString, options, baseURL, error)
}

func (m_ MutableAttributedString) InitWithMarkdownOptionsBaseURLError(markdown []byte, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from Markdown-formatted data using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746875-initwithmarkdown?language=objc
func MutableAttributedString_InitWithMarkdownOptionsBaseURLError(markdown []byte, options IAttributedStringMarkdownParsingOptions, baseURL IURL, error IError) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithMarkdownOptionsBaseURLError(markdown, options, baseURL, error)
}

func (m_ MutableAttributedString) InitWithDocFormatDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string from Microsoft Word format data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1534329-initwithdocformat?language=objc
func MutableAttributedString_InitWithDocFormatDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithDocFormatDocumentAttributes(data, dict)
}

func (m_ MutableAttributedString) InitWithRTFDDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1530987-initwithrtfd?language=objc
func MutableAttributedString_InitWithRTFDDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithRTFDDocumentAttributes(data, dict)
}

func (m_ MutableAttributedString) InitWithRTFDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string by decoding the stream of RTF commands and data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1532912-initwithrtf?language=objc
func MutableAttributedString_InitWithRTFDocumentAttributes(data []byte, dict Dictionary) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithRTFDocumentAttributes(data, dict)
}

func (m_ MutableAttributedString) InitWithRTFDFileWrapperDocumentAttributes(wrapper IFileWrapper, dict Dictionary) MutableAttributedString {
	rv := objc.Call[MutableAttributedString](m_, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), objc.Ptr(wrapper), dict)
	return rv
}

// Creates an attributed string from the specified file wrapper that contains an RTFD document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1533594-initwithrtfdfilewrapper?language=objc
func MutableAttributedString_InitWithRTFDFileWrapperDocumentAttributes(wrapper IFileWrapper, dict Dictionary) MutableAttributedString {
	return MutableAttributedStringClass.Alloc().InitWithRTFDFileWrapperDocumentAttributes(wrapper, dict)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/3746902-appendlocalizedformat?language=objc
func (m_ MutableAttributedString) AppendLocalizedFormat(format IAttributedString) {
	objc.Call[objc.Void](m_, objc.Sel("appendLocalizedFormat:"), objc.Ptr(format))
}

// Adds the characters and attributes of a given attributed string to the end of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1417879-appendattributedstring?language=objc
func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.Call[objc.Void](m_, objc.Sel("appendAttributedString:"), objc.Ptr(attrString))
}

// Deletes the characters in the given range along with their associated attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1410610-deletecharactersinrange?language=objc
func (m_ MutableAttributedString) DeleteCharactersInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("deleteCharactersInRange:"), range_)
}

// Sets the alignment characteristic of the paragraph style attribute for the characters in aRange to alignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1530762-setalignment?language=objc
func (m_ MutableAttributedString) SetAlignmentRange(alignment objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("setAlignment:range:"), objc.Ptr(alignment), range_)
}

// Inserts the characters and attributes of the given attributed string into the receiver at the given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1414947-insertattributedstring?language=objc
func (m_ MutableAttributedString) InsertAttributedStringAtIndex(attrString IAttributedString, loc uint) {
	objc.Call[objc.Void](m_, objc.Sel("insertAttributedString:atIndex:"), objc.Ptr(attrString), loc)
}

// Fixes the font attribute in aRange, assigning default fonts to characters with illegal fonts for their scripts and otherwise correcting font attribute assignments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1534303-fixfontattributeinrange?language=objc
func (m_ MutableAttributedString) FixFontAttributeInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("fixFontAttributeInRange:"), range_)
}

// Applies the font attributes specified by mask to the characters in aRange. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1532734-applyfonttraits?language=objc
func (m_ MutableAttributedString) ApplyFontTraitsRange(traitMask objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("applyFontTraits:range:"), objc.Ptr(traitMask), range_)
}

// Removes the superscript attribute from the characters in aRange. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1530210-unscriptrange?language=objc
func (m_ MutableAttributedString) UnscriptRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("unscriptRange:"), range_)
}

// Overridden by subclasses to buffer or optimize a series of changes to the receiver’s characters or attributes, until it receives a matching endEditing message, upon which it can consolidate changes and notify any observers that it has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1411853-beginediting?language=objc
func (m_ MutableAttributedString) BeginEditing() {
	objc.Call[objc.Void](m_, objc.Sel("beginEditing"))
}

// Sets the contents of the receiver from the stream at data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1535465-readfromdata?language=objc
func (m_ MutableAttributedString) ReadFromDataOptionsDocumentAttributesError(data []byte, opts Dictionary, dict Dictionary, error IError) bool {
	rv := objc.Call[bool](m_, objc.Sel("readFromData:options:documentAttributes:error:"), data, opts, dict, objc.Ptr(error))
	return rv
}

// Cleans up font, paragraph style, and attachment attributes within the given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1533823-fixattributesinrange?language=objc
func (m_ MutableAttributedString) FixAttributesInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("fixAttributesInRange:"), range_)
}

// Adds the given collection of attributes to the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1414304-addattributes?language=objc
func (m_ MutableAttributedString) AddAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("addAttributes:range:"), attrs, range_)
}

// Updates all attachments based on files contained in the RTFD file package at path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1508389-updateattachmentsfrompath?language=objc
func (m_ MutableAttributedString) UpdateAttachmentsFromPath(path string) {
	objc.Call[objc.Void](m_, objc.Sel("updateAttachmentsFromPath:"), path)
}

// Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1417045-replacecharactersinrange?language=objc
func (m_ MutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ Range, attrString IAttributedString) {
	objc.Call[objc.Void](m_, objc.Sel("replaceCharactersInRange:withAttributedString:"), range_, objc.Ptr(attrString))
}

// Overridden by subclasses to consolidate changes made since a previous beginEditing message and to notify any observers of the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1416707-endediting?language=objc
func (m_ MutableAttributedString) EndEditing() {
	objc.Call[objc.Void](m_, objc.Sel("endEditing"))
}

// Replaces the receiver’s entire contents with the characters and attributes of the given attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1411894-setattributedstring?language=objc
func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.Call[objc.Void](m_, objc.Sel("setAttributedString:"), objc.Ptr(attrString))
}

// Removes the named attribute from the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1409691-removeattribute?language=objc
func (m_ MutableAttributedString) RemoveAttributeRange(name AttributedStringKey, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeAttribute:range:"), name, range_)
}

// Adds an attribute with the given name and value to the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1417080-addattribute?language=objc
func (m_ MutableAttributedString) AddAttributeValueRange(name AttributedStringKey, value objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("addAttribute:value:range:"), name, value, range_)
}

// Sets the attributes for the characters in the specified range to the specified attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1412179-setattributes?language=objc
func (m_ MutableAttributedString) SetAttributesRange(attrs map[AttributedStringKey]objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("setAttributes:range:"), attrs, range_)
}

// Increments the value of the superscript attribute for characters in aRange by 1. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1533592-superscriptrange?language=objc
func (m_ MutableAttributedString) SuperscriptRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("superscriptRange:"), range_)
}

// Fixes the paragraph style attributes in aRange, assigning the first paragraph style attribute value in each paragraph to all characters of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1527082-fixparagraphstyleattributeinrang?language=objc
func (m_ MutableAttributedString) FixParagraphStyleAttributeInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("fixParagraphStyleAttributeInRange:"), range_)
}

// Decrements the value of the superscript attribute for characters in aRange by 1. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1529906-subscriptrange?language=objc
func (m_ MutableAttributedString) SubscriptRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("subscriptRange:"), range_)
}

// Sets the base writing direction for the characters in range to writingDirection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1532923-setbasewritingdirection?language=objc
func (m_ MutableAttributedString) SetBaseWritingDirectionRange(writingDirection objc.IObject, range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("setBaseWritingDirection:range:"), objc.Ptr(writingDirection), range_)
}

// Cleans up attachment attributes in aRange, removing all attachment attributes assigned to characters other than NSAttachmentCharacter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1526389-fixattachmentattributeinrange?language=objc
func (m_ MutableAttributedString) FixAttachmentAttributeInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("fixAttachmentAttributeInRange:"), range_)
}

// The character contents of the receiver as an NSMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/1416955-mutablestring?language=objc
func (m_ MutableAttributedString) MutableString() MutableString {
	rv := objc.Call[MutableString](m_, objc.Sel("mutableString"))
	return rv
}
