// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextStorage] class.
var TextStorageClass = _TextStorageClass{objc.GetClass("NSTextStorage")}

type _TextStorageClass struct {
	objc.Class
}

// An interface definition for the [TextStorage] class.
type ITextStorage interface {
	foundation.IMutableAttributedString
	EnsureAttributesAreFixedInRange(range_ foundation.Range)
	InvalidateAttributesInRange(range_ foundation.Range)
	EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int)
	ProcessEditing()
	RemoveLayoutManager(aLayoutManager ILayoutManager)
	AddLayoutManager(aLayoutManager ILayoutManager)
	Paragraphs() []TextStorage
	SetParagraphs(value []ITextStorage)
	LayoutManagers() []LayoutManager
	ChangeInLength() int
	Characters() []TextStorage
	SetCharacters(value []ITextStorage)
	EditedRange() foundation.Range
	Delegate() TextStorageDelegateWrapper
	SetDelegate(value PTextStorageDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ForegroundColor() Color
	SetForegroundColor(value IColor)
	Font() Font
	SetFont(value IFont)
	EditedMask() TextStorageEditActions
	AttributeRuns() []TextStorage
	SetAttributeRuns(value []ITextStorage)
	Words() []TextStorage
	SetWords(value []ITextStorage)
	TextStorageObserver() TextStorageObservingWrapper
	SetTextStorageObserver(value PTextStorageObserving)
	SetTextStorageObserverObject(valueObject objc.IObject)
	FixesAttributesLazily() bool
}

// The fundamental storage mechanism of TextKit that contains the text managed by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage?language=objc
type TextStorage struct {
	foundation.MutableAttributedString
}

func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		MutableAttributedString: foundation.MutableAttributedStringFrom(ptr),
	}
}

func (tc _TextStorageClass) Alloc() TextStorage {
	rv := objc.Call[TextStorage](tc, objc.Sel("alloc"))
	return rv
}

func TextStorage_Alloc() TextStorage {
	return TextStorageClass.Alloc()
}

func (tc _TextStorageClass) New() TextStorage {
	rv := objc.Call[TextStorage](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextStorage() TextStorage {
	return TextStorageClass.New()
}

func (t_ TextStorage) Init() TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("init"))
	return rv
}

func (t_ TextStorage) InitWithHTMLOptionsDocumentAttributes(data []byte, options map[AttributedStringDocumentReadingOptionKey]objc.IObject, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithHTML:options:documentAttributes:"), data, options, dict)
	return rv
}

// Creates an attributed string from the HTML in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1535412-initwithhtml?language=objc
func TextStorage_InitWithHTMLOptionsDocumentAttributes(data []byte, options map[AttributedStringDocumentReadingOptionKey]objc.IObject, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	return TextStorageClass.Alloc().InitWithHTMLOptionsDocumentAttributes(data, options, dict)
}

func (tc _TextStorageClass) LocalizedAttributedStringWithFormat(format foundation.IAttributedString) TextStorage {
	rv := objc.Call[TextStorage](tc, objc.Sel("localizedAttributedStringWithFormat:"), objc.Ptr(format))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746877-localizedattributedstringwithfor?language=objc
func TextStorage_LocalizedAttributedStringWithFormat(format foundation.IAttributedString) TextStorage {
	return TextStorageClass.LocalizedAttributedStringWithFormat(format)
}

func (t_ TextStorage) InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithAttributedString:"), objc.Ptr(attrStr))
	return rv
}

// Creates an attributed string with the characters and attributes of the specified attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1415342-initwithattributedstring?language=objc
func TextStorage_InitWithAttributedString(attrStr foundation.IAttributedString) TextStorage {
	return TextStorageClass.Alloc().InitWithAttributedString(attrStr)
}

func (t_ TextStorage) InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile foundation.IURL, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithContentsOfMarkdownFileAtURL:options:baseURL:error:"), objc.Ptr(markdownFile), objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from the contents of a specified URL that contains Markdown-formatted data using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746872-initwithcontentsofmarkdownfileat?language=objc
func TextStorage_InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile foundation.IURL, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	return TextStorageClass.Alloc().InitWithContentsOfMarkdownFileAtURLOptionsBaseURLError(markdownFile, options, baseURL, error)
}

func (t_ TextStorage) InitWithFormatOptionsLocale(format foundation.IAttributedString, options foundation.AttributedStringFormattingOptions, locale foundation.ILocale) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithFormat:options:locale:"), objc.Ptr(format), options, objc.Ptr(locale))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746873-initwithformat?language=objc
func TextStorage_InitWithFormatOptionsLocale(format foundation.IAttributedString, options foundation.AttributedStringFormattingOptions, locale foundation.ILocale) TextStorage {
	return TextStorageClass.Alloc().InitWithFormatOptionsLocale(format, options, locale)
}

func (t_ TextStorage) InitWithDataOptionsDocumentAttributesError(data []byte, options map[AttributedStringDocumentReadingOptionKey]objc.IObject, dict map[AttributedStringDocumentAttributeKey]objc.IObject, error foundation.IError) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithData:options:documentAttributes:error:"), data, options, dict, objc.Ptr(error))
	return rv
}

// Creates an attributed string from the data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1524613-initwithdata?language=objc
func TextStorage_InitWithDataOptionsDocumentAttributesError(data []byte, options map[AttributedStringDocumentReadingOptionKey]objc.IObject, dict map[AttributedStringDocumentAttributeKey]objc.IObject, error foundation.IError) TextStorage {
	return TextStorageClass.Alloc().InitWithDataOptionsDocumentAttributesError(data, options, dict, error)
}

func (t_ TextStorage) InitWithString(str string) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithString:"), str)
	return rv
}

// Creates an attributed string with the characters of the specified string and no attribute information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1407481-initwithstring?language=objc
func TextStorage_InitWithString(str string) TextStorage {
	return TextStorageClass.Alloc().InitWithString(str)
}

func (t_ TextStorage) InitWithMarkdownStringOptionsBaseURLError(markdownString string, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithMarkdownString:options:baseURL:error:"), markdownString, objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from a Markdown-formatted string using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746876-initwithmarkdownstring?language=objc
func TextStorage_InitWithMarkdownStringOptionsBaseURLError(markdownString string, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	return TextStorageClass.Alloc().InitWithMarkdownStringOptionsBaseURLError(markdownString, options, baseURL, error)
}

func (t_ TextStorage) InitWithMarkdownOptionsBaseURLError(markdown []byte, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithMarkdown:options:baseURL:error:"), markdown, objc.Ptr(options), objc.Ptr(baseURL), objc.Ptr(error))
	return rv
}

// Creates an attributed string from Markdown-formatted data using the provided options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/3746875-initwithmarkdown?language=objc
func TextStorage_InitWithMarkdownOptionsBaseURLError(markdown []byte, options foundation.IAttributedStringMarkdownParsingOptions, baseURL foundation.IURL, error foundation.IError) TextStorage {
	return TextStorageClass.Alloc().InitWithMarkdownOptionsBaseURLError(markdown, options, baseURL, error)
}

func (t_ TextStorage) InitWithDocFormatDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithDocFormat:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string from Microsoft Word format data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1534329-initwithdocformat?language=objc
func TextStorage_InitWithDocFormatDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	return TextStorageClass.Alloc().InitWithDocFormatDocumentAttributes(data, dict)
}

func (t_ TextStorage) InitWithRTFDDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithRTFD:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string by decoding the stream of RTFD commands and data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1530987-initwithrtfd?language=objc
func TextStorage_InitWithRTFDDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	return TextStorageClass.Alloc().InitWithRTFDDocumentAttributes(data, dict)
}

func (t_ TextStorage) InitWithRTFDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithRTF:documentAttributes:"), data, dict)
	return rv
}

// Creates an attributed string by decoding the stream of RTF commands and data in the specified data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1532912-initwithrtf?language=objc
func TextStorage_InitWithRTFDocumentAttributes(data []byte, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	return TextStorageClass.Alloc().InitWithRTFDocumentAttributes(data, dict)
}

func (t_ TextStorage) InitWithRTFDFileWrapperDocumentAttributes(wrapper foundation.IFileWrapper, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("initWithRTFDFileWrapper:documentAttributes:"), objc.Ptr(wrapper), dict)
	return rv
}

// Creates an attributed string from the specified file wrapper that contains an RTFD document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/1533594-initwithrtfdfilewrapper?language=objc
func TextStorage_InitWithRTFDFileWrapperDocumentAttributes(wrapper foundation.IFileWrapper, dict map[AttributedStringDocumentAttributeKey]objc.IObject) TextStorage {
	return TextStorageClass.Alloc().InitWithRTFDFileWrapperDocumentAttributes(wrapper, dict)
}

// Ensures that attribute fixing occurs in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1533947-ensureattributesarefixedinrange?language=objc
func (t_ TextStorage) EnsureAttributesAreFixedInRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("ensureAttributesAreFixedInRange:"), range_)
}

// Invalidates attributes in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1534025-invalidateattributesinrange?language=objc
func (t_ TextStorage) InvalidateAttributesInRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("invalidateAttributesInRange:"), range_)
}

// Tracks changes made to the text storage object, allowing the text storage to record the full extent of changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1529793-edited?language=objc
func (t_ TextStorage) EditedRangeChangeInLength(editedMask TextStorageEditActions, editedRange foundation.Range, delta int) {
	objc.Call[objc.Void](t_, objc.Sel("edited:range:changeInLength:"), editedMask, editedRange, delta)
}

// Cleans up changes to the text storage object and notifies its delegate and layout managers of changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1525980-processediting?language=objc
func (t_ TextStorage) ProcessEditing() {
	objc.Call[objc.Void](t_, objc.Sel("processEditing"))
}

// Removes a layout manager from the text storage object’s set of layout managers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1528755-removelayoutmanager?language=objc
func (t_ TextStorage) RemoveLayoutManager(aLayoutManager ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("removeLayoutManager:"), objc.Ptr(aLayoutManager))
}

// Adds a layout manager to the text storage object’s set of layout managers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1533459-addlayoutmanager?language=objc
func (t_ TextStorage) AddLayoutManager(aLayoutManager ILayoutManager) {
	objc.Call[objc.Void](t_, objc.Sel("addLayoutManager:"), objc.Ptr(aLayoutManager))
}

// The text storage contents as an array of paragraphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1525943-paragraphs?language=objc
func (t_ TextStorage) Paragraphs() []TextStorage {
	rv := objc.Call[[]TextStorage](t_, objc.Sel("paragraphs"))
	return rv
}

// The text storage contents as an array of paragraphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1525943-paragraphs?language=objc
func (t_ TextStorage) SetParagraphs(value []ITextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setParagraphs:"), value)
}

// The layout managers for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1527938-layoutmanagers?language=objc
func (t_ TextStorage) LayoutManagers() []LayoutManager {
	rv := objc.Call[[]LayoutManager](t_, objc.Sel("layoutManagers"))
	return rv
}

// The difference between the current length of the edited range and its length before editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1528400-changeinlength?language=objc
func (t_ TextStorage) ChangeInLength() int {
	rv := objc.Call[int](t_, objc.Sel("changeInLength"))
	return rv
}

// The text storage contents as an array of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1535788-characters?language=objc
func (t_ TextStorage) Characters() []TextStorage {
	rv := objc.Call[[]TextStorage](t_, objc.Sel("characters"))
	return rv
}

// The text storage contents as an array of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1535788-characters?language=objc
func (t_ TextStorage) SetCharacters(value []ITextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setCharacters:"), value)
}

// The range of text that contains changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1524379-editedrange?language=objc
func (t_ TextStorage) EditedRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("editedRange"))
	return rv
}

// The delegate for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1532704-delegate?language=objc
func (t_ TextStorage) Delegate() TextStorageDelegateWrapper {
	rv := objc.Call[TextStorageDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The delegate for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1532704-delegate?language=objc
func (t_ TextStorage) SetDelegate(value PTextStorageDelegate) {
	po0 := objc.WrapAsProtocol("NSTextStorageDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1532704-delegate?language=objc
func (t_ TextStorage) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The color for the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1527175-foregroundcolor?language=objc
func (t_ TextStorage) ForegroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("foregroundColor"))
	return rv
}

// The color for the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1527175-foregroundcolor?language=objc
func (t_ TextStorage) SetForegroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setForegroundColor:"), objc.Ptr(value))
}

// The font for the text storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1535365-font?language=objc
func (t_ TextStorage) Font() Font {
	rv := objc.Call[Font](t_, objc.Sel("font"))
	return rv
}

// The font for the text storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1535365-font?language=objc
func (t_ TextStorage) SetFont(value IFont) {
	objc.Call[objc.Void](t_, objc.Sel("setFont:"), objc.Ptr(value))
}

// A mask that describes the kinds of edits pending for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1525323-editedmask?language=objc
func (t_ TextStorage) EditedMask() TextStorageEditActions {
	rv := objc.Call[TextStorageEditActions](t_, objc.Sel("editedMask"))
	return rv
}

// The text storage contents as an array of attribute runs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1532095-attributeruns?language=objc
func (t_ TextStorage) AttributeRuns() []TextStorage {
	rv := objc.Call[[]TextStorage](t_, objc.Sel("attributeRuns"))
	return rv
}

// The text storage contents as an array of attribute runs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1532095-attributeruns?language=objc
func (t_ TextStorage) SetAttributeRuns(value []ITextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setAttributeRuns:"), value)
}

// The text storage contents as an array of words. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1524565-words?language=objc
func (t_ TextStorage) Words() []TextStorage {
	rv := objc.Call[[]TextStorage](t_, objc.Sel("words"))
	return rv
}

// The text storage contents as an array of words. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextstorage/1524565-words?language=objc
func (t_ TextStorage) SetWords(value []ITextStorage) {
	objc.Call[objc.Void](t_, objc.Sel("setWords:"), value)
}

// The observer for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/3824756-textstorageobserver?language=objc
func (t_ TextStorage) TextStorageObserver() TextStorageObservingWrapper {
	rv := objc.Call[TextStorageObservingWrapper](t_, objc.Sel("textStorageObserver"))
	return rv
}

// The observer for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/3824756-textstorageobserver?language=objc
func (t_ TextStorage) SetTextStorageObserver(value PTextStorageObserving) {
	po0 := objc.WrapAsProtocol("NSTextStorageObserving", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setTextStorageObserver"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setTextStorageObserver:"), po0)
}

// The observer for the text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/3824756-textstorageobserver?language=objc
func (t_ TextStorage) SetTextStorageObserverObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setTextStorageObserver:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether the text storage object fixes attributes lazily. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextstorage/1532043-fixesattributeslazily?language=objc
func (t_ TextStorage) FixesAttributesLazily() bool {
	rv := objc.Call[bool](t_, objc.Sel("fixesAttributesLazily"))
	return rv
}
