// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ParagraphStyleClass = _ParagraphStyleClass{objc.GetClass("NSParagraphStyle")}

type _ParagraphStyleClass struct {
	objc.Class
}

type IParagraphStyle interface {
	objc.IObject
	Alignment() TextAlignment
	FirstLineHeadIndent() float64
	HeadIndent() float64
	TailIndent() float64
	LineHeightMultiple() float64
	MaximumLineHeight() float64
	MinimumLineHeight() float64
	LineSpacing() float64
	ParagraphSpacing() float64
	ParagraphSpacingBefore() float64
	TabStops() []TextTab
	DefaultTabInterval() float64
	TextBlocks() []TextBlock
	TextLists() []TextList
	LineBreakMode() LineBreakMode
	LineBreakStrategy() LineBreakStrategy
	HyphenationFactor() float32
	UsesDefaultHyphenation() bool
	TighteningFactorForTruncation() float32
	AllowsDefaultTighteningForTruncation() bool
	HeaderLevel() int
	BaseWritingDirection() WritingDirection
}

type ParagraphStyle struct {
	objc.Object
}

func MakeParagraphStyle(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _ParagraphStyleClass) Alloc() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](pc, objc.GetSelector("alloc"))
	return rv
}

func ParagraphStyle_Alloc() ParagraphStyle {
	return ParagraphStyleClass.Alloc()
}

func (pc _ParagraphStyleClass) New() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func ParagraphStyle_New() ParagraphStyle {
	return ParagraphStyleClass.New()
}

func (p_ ParagraphStyle) Init() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](p_, objc.GetSelector("init"))
	return rv
}

func ParagraphStyle_Init() ParagraphStyle {
	return ParagraphStyleClass.Alloc().Init()
}

func (pc _ParagraphStyleClass) DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	rv := objc.CallMethod[WritingDirection](pc, objc.GetSelector("defaultWritingDirectionForLanguage:"), languageName)
	return rv
}

func ParagraphStyle_DefaultWritingDirectionForLanguage(languageName string) WritingDirection {
	return ParagraphStyleClass.DefaultWritingDirectionForLanguage(languageName)
}

func (pc _ParagraphStyleClass) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](pc, objc.GetSelector("defaultParagraphStyle"))
	return rv
}

func ParagraphStyle_DefaultParagraphStyle() ParagraphStyle {
	return ParagraphStyleClass.DefaultParagraphStyle()
}

func (p_ ParagraphStyle) Alignment() TextAlignment {
	rv := objc.CallMethod[TextAlignment](p_, objc.GetSelector("alignment"))
	return rv
}

func (p_ ParagraphStyle) FirstLineHeadIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("firstLineHeadIndent"))
	return rv
}

func (p_ ParagraphStyle) HeadIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("headIndent"))
	return rv
}

func (p_ ParagraphStyle) TailIndent() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("tailIndent"))
	return rv
}

func (p_ ParagraphStyle) LineHeightMultiple() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("lineHeightMultiple"))
	return rv
}

func (p_ ParagraphStyle) MaximumLineHeight() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("maximumLineHeight"))
	return rv
}

func (p_ ParagraphStyle) MinimumLineHeight() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("minimumLineHeight"))
	return rv
}

func (p_ ParagraphStyle) LineSpacing() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("lineSpacing"))
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacing() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("paragraphSpacing"))
	return rv
}

func (p_ ParagraphStyle) ParagraphSpacingBefore() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("paragraphSpacingBefore"))
	return rv
}

func (p_ ParagraphStyle) TabStops() []TextTab {
	rv := objc.CallMethod[[]TextTab](p_, objc.GetSelector("tabStops"))
	return rv
}

func (p_ ParagraphStyle) DefaultTabInterval() float64 {
	rv := objc.CallMethod[float64](p_, objc.GetSelector("defaultTabInterval"))
	return rv
}

func (p_ ParagraphStyle) TextBlocks() []TextBlock {
	rv := objc.CallMethod[[]TextBlock](p_, objc.GetSelector("textBlocks"))
	return rv
}

func (p_ ParagraphStyle) TextLists() []TextList {
	rv := objc.CallMethod[[]TextList](p_, objc.GetSelector("textLists"))
	return rv
}

func (p_ ParagraphStyle) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](p_, objc.GetSelector("lineBreakMode"))
	return rv
}

func (p_ ParagraphStyle) LineBreakStrategy() LineBreakStrategy {
	rv := objc.CallMethod[LineBreakStrategy](p_, objc.GetSelector("lineBreakStrategy"))
	return rv
}

func (p_ ParagraphStyle) HyphenationFactor() float32 {
	rv := objc.CallMethod[float32](p_, objc.GetSelector("hyphenationFactor"))
	return rv
}

func (p_ ParagraphStyle) UsesDefaultHyphenation() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("usesDefaultHyphenation"))
	return rv
}

func (p_ ParagraphStyle) TighteningFactorForTruncation() float32 {
	rv := objc.CallMethod[float32](p_, objc.GetSelector("tighteningFactorForTruncation"))
	return rv
}

func (p_ ParagraphStyle) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("allowsDefaultTighteningForTruncation"))
	return rv
}

func (p_ ParagraphStyle) HeaderLevel() int {
	rv := objc.CallMethod[int](p_, objc.GetSelector("headerLevel"))
	return rv
}

func (p_ ParagraphStyle) BaseWritingDirection() WritingDirection {
	rv := objc.CallMethod[WritingDirection](p_, objc.GetSelector("baseWritingDirection"))
	return rv
}
