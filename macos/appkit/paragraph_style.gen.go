// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ParagraphStyle] class.
var ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}

type _ParagraphStyleClass struct {
	objc.Class
}

// An interface definition for the [ParagraphStyle] class.
type IParagraphStyle interface {
	objc.IObject
	LineSpacing() float64
	TextLists() []TextList
	AllowsDefaultTighteningForTruncation() bool
	UsesDefaultHyphenation() bool
	MinimumLineHeight() float64
	TighteningFactorForTruncation() float64
	Alignment() TextAlignment
	FirstLineHeadIndent() float64
	LineBreakStrategy() LineBreakStrategy
	MaximumLineHeight() float64
	HeaderLevel() int
	DefaultTabInterval() float64
	TabStops() []TextTab
	TextBlocks() []TextBlock
	LineHeightMultiple() float64
	BaseWritingDirection() WritingDirection
	HyphenationFactor() float64
	ParagraphSpacingBefore() float64
	ParagraphSpacing() float64
	TailIndent() float64
	HeadIndent() float64
}

// The paragraph or ruler attributes for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle?language=objc
type ParagraphStyle struct {
	objc.Object
}

func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](pc, objc.Sel("alloc"))
	return rv
}

func ParagraphStyle_Alloc() ParagraphStyle {
	return ParagraphStyleClass.Alloc()
}

func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](p_, objc.Sel("init"))
	return rv
}

// Returns the default writing direction for the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1535327-defaultwritingdirectionforlangua?language=objc
func (pc _ParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	rv := objc.Call[WritingDirection](pc, objc.Sel("defaultWritingDirectionForLanguage:"), languageName)
	return rv
}

// Returns the default writing direction for the specified language. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1535327-defaultwritingdirectionforlangua?language=objc
func ParagraphStyle_DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	return ParagraphStyleClass.DefaultWritingDirectionForLanguage(languageName)
}

// The default paragraph style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1532681-defaultparagraphstyle?language=objc
func (pc _ParagraphStyleClass) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](pc, objc.Sel("defaultParagraphStyle"))
	return rv
}

// The default paragraph style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1532681-defaultparagraphstyle?language=objc
func ParagraphStyle_DefaultParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.DefaultParagraphStyle()
}

// The distance in points between the bottom of one line fragment and the top of the next. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1524635-linespacing?language=objc
func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := objc.Call[float64](p_, objc.Sel("lineSpacing"))
	return rv
}

// The text lists that contain the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/1534193-textlists?language=objc
func (p_ ParagraphStyle) TextLists() []TextList {
	rv := objc.Call[[]TextList](p_, objc.Sel("textLists"))
	return rv
}

// A Boolean value that indicates whether the system tightens character spacing before truncating text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1528994-allowsdefaulttighteningfortrunca?language=objc
func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Call[bool](p_, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}

// A Boolean value that indicates whether the paragraph style uses the system hyphenation settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/3809913-usesdefaulthyphenation?language=objc
func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.Call[bool](p_, objc.Sel("usesDefaultHyphenation"))
	return rv
}

// The paragraph’s minimum line height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1535639-minimumlineheight?language=objc
func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.Call[float64](p_, objc.Sel("minimumLineHeight"))
	return rv
}

// The threshold for using tightening as an alternative to truncation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/1529278-tighteningfactorfortruncation?language=objc
func (p_ ParagraphStyle) TighteningFactorForTruncation() float64 {
	rv := objc.Call[float64](p_, objc.Sel("tighteningFactorForTruncation"))
	return rv
}

// The text alignment of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1532321-alignment?language=objc
func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := objc.Call[TextAlignment](p_, objc.Sel("alignment"))
	return rv
}

// The indentation of the first line of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1527764-firstlineheadindent?language=objc
func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.Call[float64](p_, objc.Sel("firstLineHeadIndent"))
	return rv
}

// The strategy for breaking lines while laying out paragraphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/3667463-linebreakstrategy?language=objc
func (p_ ParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.Call[LineBreakStrategy](p_, objc.Sel("lineBreakStrategy"))
	return rv
}

// The paragraph’s maximum line height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1533343-maximumlineheight?language=objc
func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.Call[float64](p_, objc.Sel("maximumLineHeight"))
	return rv
}

// The paragraph’s header level for HTML generation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/1535869-headerlevel?language=objc
func (p_ ParagraphStyle) HeaderLevel() int {
	rv := objc.Call[int](p_, objc.Sel("headerLevel"))
	return rv
}

// The documentwide default tab interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1535614-defaulttabinterval?language=objc
func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.Call[float64](p_, objc.Sel("defaultTabInterval"))
	return rv
}

// The text tab objects that represent the paragraph’s tab stops. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1532841-tabstops?language=objc
func (p_ ParagraphStyle) TabStops() []TextTab {
	rv := objc.Call[[]TextTab](p_, objc.Sel("tabStops"))
	return rv
}

// The text blocks that contain the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsparagraphstyle/1528053-textblocks?language=objc
func (p_ ParagraphStyle) TextBlocks() []TextBlock {
	rv := objc.Call[[]TextBlock](p_, objc.Sel("textBlocks"))
	return rv
}

// The line height multiple. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1528614-lineheightmultiple?language=objc
func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.Call[float64](p_, objc.Sel("lineHeightMultiple"))
	return rv
}

// The base writing direction for the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1527354-basewritingdirection?language=objc
func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.Call[WritingDirection](p_, objc.Sel("baseWritingDirection"))
	return rv
}

// The paragraph’s threshold for hyphenation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1529275-hyphenationfactor?language=objc
func (p_ ParagraphStyle) HyphenationFactor() float64 {
	rv := objc.Call[float64](p_, objc.Sel("hyphenationFactor"))
	return rv
}

// The distance between the paragraph’s top and the beginning of its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1533011-paragraphspacingbefore?language=objc
func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.Call[float64](p_, objc.Sel("paragraphSpacingBefore"))
	return rv
}

// Distance between the bottom of this paragraph and top of next. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1530912-paragraphspacing?language=objc
func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.Call[float64](p_, objc.Sel("paragraphSpacing"))
	return rv
}

// The trailing indentation of the paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1525556-tailindent?language=objc
func (p_ ParagraphStyle) TailIndent() float64 {
	rv := objc.Call[float64](p_, objc.Sel("tailIndent"))
	return rv
}

// The indentation of the paragraph’s lines other than the first. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsparagraphstyle/1530760-headindent?language=objc
func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := objc.Call[float64](p_, objc.Sel("headIndent"))
	return rv
}
