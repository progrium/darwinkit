// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var MutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}

type _MutableParagraphStyleClass struct {
	objc.Class
}

// An interface definition for the [MutableParagraphStyle] class.
type IMutableParagraphStyle interface {
	IParagraphStyle
	SetParagraphStyle(obj IParagraphStyle)
	AddTabStop(anObject ITextTab)
	RemoveTabStop(anObject ITextTab)
	SetLineSpacing(value float64)
	SetTextLists(value []ITextList)
	SetAllowsDefaultTighteningForTruncation(value bool)
	SetUsesDefaultHyphenation(value bool)
	SetMinimumLineHeight(value float64)
	SetTighteningFactorForTruncation(value float64)
	SetAlignment(value TextAlignment)
	SetFirstLineHeadIndent(value float64)
	SetLineBreakStrategy(value LineBreakStrategy)
	SetMaximumLineHeight(value float64)
	SetHeaderLevel(value int)
	SetDefaultTabInterval(value float64)
	SetTabStops(value []ITextTab)
	SetTextBlocks(value []ITextBlock)
	SetLineHeightMultiple(value float64)
	SetBaseWritingDirection(value WritingDirection)
	SetHyphenationFactor(value float64)
	SetParagraphSpacingBefore(value float64)
	SetParagraphSpacing(value float64)
	SetTailIndent(value float64)
	SetHeadIndent(value float64)
}

// An object for changing the values of the subattributes in a paragraph style attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle?language=objc
type MutableParagraphStyle struct {
	ParagraphStyle
}

func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: ParagraphStyleFrom(ptr),
	}
}

func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.Call[MutableParagraphStyle](mc, objc.Sel("alloc"))
	return rv
}

func MutableParagraphStyle_Alloc() MutableParagraphStyle {
	return MutableParagraphStyleClass.Alloc()
}

func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := objc.Call[MutableParagraphStyle](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableParagraphStyle() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := objc.Call[MutableParagraphStyle](m_, objc.Sel("init"))
	return rv
}

// Replaces the subattributes of the paragraph with those in the specified paragraph style object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1533980-setparagraphstyle?language=objc
func (m_ MutableParagraphStyle) SetParagraphStyle(obj IParagraphStyle) {
	objc.Call[objc.Void](m_, objc.Sel("setParagraphStyle:"), objc.Ptr(obj))
}

// Adds the specified tab stop to the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1525051-addtabstop?language=objc
func (m_ MutableParagraphStyle) AddTabStop(anObject ITextTab) {
	objc.Call[objc.Void](m_, objc.Sel("addTabStop:"), objc.Ptr(anObject))
}

// Removes the first text tab with a location and type equal to the specified tab stop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1535084-removetabstop?language=objc
func (m_ MutableParagraphStyle) RemoveTabStop(anObject ITextTab) {
	objc.Call[objc.Void](m_, objc.Sel("removeTabStop:"), objc.Ptr(anObject))
}

// The distance in points between the bottom of one line fragment and the top of the next. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1528742-linespacing?language=objc
func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setLineSpacing:"), value)
}

// The text lists that contain the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/1524648-textlists?language=objc
func (m_ MutableParagraphStyle) SetTextLists(value []ITextList) {
	objc.Call[objc.Void](m_, objc.Sel("setTextLists:"), value)
}

// A Boolean value that indicates whether the system tightens intercharacter spacing before truncating text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1534136-allowsdefaulttighteningfortrunca?language=objc
func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/3809912-usesdefaulthyphenation?language=objc
func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setUsesDefaultHyphenation:"), value)
}

// The paragraph’s minimum line height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1531118-minimumlineheight?language=objc
func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMinimumLineHeight:"), value)
}

// The threshold for using tightening as an alternative to truncation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/1531383-tighteningfactorfortruncation?language=objc
func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setTighteningFactorForTruncation:"), value)
}

// The text alignment of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1534368-alignment?language=objc
func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	objc.Call[objc.Void](m_, objc.Sel("setAlignment:"), value)
}

// The indentation of the first line of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1528392-firstlineheadindent?language=objc
func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setFirstLineHeadIndent:"), value)
}

// The strategies that the text system may use to break lines while laying out the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/3667462-linebreakstrategy?language=objc
func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.Call[objc.Void](m_, objc.Sel("setLineBreakStrategy:"), value)
}

// The paragraph’s maximum line height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1524351-maximumlineheight?language=objc
func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMaximumLineHeight:"), value)
}

// The paragraph’s header level for HTML generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/1533962-headerlevel?language=objc
func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setHeaderLevel:"), value)
}

// A number used as the document’s default tab spacing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1529861-defaulttabinterval?language=objc
func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setDefaultTabInterval:"), value)
}

// The text tab objects that represent the paragraph’s tab stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1531988-tabstops?language=objc
func (m_ MutableParagraphStyle) SetTabStops(value []ITextTab) {
	objc.Call[objc.Void](m_, objc.Sel("setTabStops:"), value)
}

// The text blocks that contain the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmutableparagraphstyle/1535855-textblocks?language=objc
func (m_ MutableParagraphStyle) SetTextBlocks(value []ITextBlock) {
	objc.Call[objc.Void](m_, objc.Sel("setTextBlocks:"), value)
}

// The line height multiple. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1524596-lineheightmultiple?language=objc
func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setLineHeightMultiple:"), value)
}

// The base writing direction for the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1534601-basewritingdirection?language=objc
func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.Call[objc.Void](m_, objc.Sel("setBaseWritingDirection:"), value)
}

// The paragraph’s threshold for hyphenation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1535553-hyphenationfactor?language=objc
func (m_ MutableParagraphStyle) SetHyphenationFactor(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setHyphenationFactor:"), value)
}

// The distance between the paragraph’s top and the beginning of its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1527729-paragraphspacingbefore?language=objc
func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setParagraphSpacingBefore:"), value)
}

// The space after the end of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1532528-paragraphspacing?language=objc
func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setParagraphSpacing:"), value)
}

// The trailing indentation of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1531666-tailindent?language=objc
func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setTailIndent:"), value)
}

// The indentation of the paragraph’s lines other than the first. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsmutableparagraphstyle/1525135-headindent?language=objc
func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setHeadIndent:"), value)
}
