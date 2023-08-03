// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var MutableParagraphStyleClass = _MutableParagraphStyleClass{objc.GetClass("NSMutableParagraphStyle")}

type _MutableParagraphStyleClass struct {
	objc.Class
}

type IMutableParagraphStyle interface {
	IParagraphStyle
	SetParagraphStyle(obj IParagraphStyle)
	AddTabStop(anObject ITextTab)
	RemoveTabStop(anObject ITextTab)
	SetAlignment(value TextAlignment)
	SetFirstLineHeadIndent(value float64)
	SetHeadIndent(value float64)
	SetTailIndent(value float64)
	SetLineHeightMultiple(value float64)
	SetMaximumLineHeight(value float64)
	SetMinimumLineHeight(value float64)
	SetLineSpacing(value float64)
	SetParagraphSpacing(value float64)
	SetParagraphSpacingBefore(value float64)
	SetBaseWritingDirection(value WritingDirection)
	SetTabStops(value []ITextTab)
	SetDefaultTabInterval(value float64)
	SetTextBlocks(value []ITextBlock)
	SetTextLists(value []ITextList)
	SetLineBreakMode(value LineBreakMode)
	SetLineBreakStrategy(value LineBreakStrategy)
	SetHyphenationFactor(value float32)
	SetUsesDefaultHyphenation(value bool)
	SetTighteningFactorForTruncation(value float32)
	SetAllowsDefaultTighteningForTruncation(value bool)
	SetHeaderLevel(value int)
}

type MutableParagraphStyle struct {
	ParagraphStyle
}

func MakeMutableParagraphStyle(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ParagraphStyle: MakeParagraphStyle(ptr),
	}
}

func (mc _MutableParagraphStyleClass) Alloc() MutableParagraphStyle {
	rv := objc.CallMethod[MutableParagraphStyle](mc, objc.GetSelector("alloc"))
	return rv
}

func MutableParagraphStyle_Alloc() MutableParagraphStyle {
	return MutableParagraphStyleClass.Alloc()
}

func (mc _MutableParagraphStyleClass) New() MutableParagraphStyle {
	rv := objc.CallMethod[MutableParagraphStyle](mc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewMutableParagraphStyle() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

func MutableParagraphStyle_New() MutableParagraphStyle {
	return MutableParagraphStyleClass.New()
}

func (m_ MutableParagraphStyle) Init() MutableParagraphStyle {
	rv := objc.CallMethod[MutableParagraphStyle](m_, objc.GetSelector("init"))
	return rv
}

func MutableParagraphStyle_Init() MutableParagraphStyle {
	return MutableParagraphStyleClass.Alloc().Init()
}

func (m_ MutableParagraphStyle) SetParagraphStyle(obj IParagraphStyle) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setParagraphStyle:"), objc.ExtractPtr(obj))
}

func (m_ MutableParagraphStyle) AddTabStop(anObject ITextTab) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("addTabStop:"), objc.ExtractPtr(anObject))
}

func (m_ MutableParagraphStyle) RemoveTabStop(anObject ITextTab) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("removeTabStop:"), objc.ExtractPtr(anObject))
}

func (m_ MutableParagraphStyle) SetAlignment(value TextAlignment) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAlignment:"), value)
}

func (m_ MutableParagraphStyle) SetFirstLineHeadIndent(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setFirstLineHeadIndent:"), value)
}

func (m_ MutableParagraphStyle) SetHeadIndent(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHeadIndent:"), value)
}

func (m_ MutableParagraphStyle) SetTailIndent(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTailIndent:"), value)
}

func (m_ MutableParagraphStyle) SetLineHeightMultiple(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setLineHeightMultiple:"), value)
}

func (m_ MutableParagraphStyle) SetMaximumLineHeight(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMaximumLineHeight:"), value)
}

func (m_ MutableParagraphStyle) SetMinimumLineHeight(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setMinimumLineHeight:"), value)
}

func (m_ MutableParagraphStyle) SetLineSpacing(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setLineSpacing:"), value)
}

func (m_ MutableParagraphStyle) SetParagraphSpacing(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setParagraphSpacing:"), value)
}

func (m_ MutableParagraphStyle) SetParagraphSpacingBefore(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setParagraphSpacingBefore:"), value)
}

func (m_ MutableParagraphStyle) SetBaseWritingDirection(value WritingDirection) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setBaseWritingDirection:"), value)
}

func (m_ MutableParagraphStyle) SetTabStops(value []ITextTab) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTabStops:"), value)
}

func (m_ MutableParagraphStyle) SetDefaultTabInterval(value float64) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setDefaultTabInterval:"), value)
}

func (m_ MutableParagraphStyle) SetTextBlocks(value []ITextBlock) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTextBlocks:"), value)
}

func (m_ MutableParagraphStyle) SetTextLists(value []ITextList) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTextLists:"), value)
}

func (m_ MutableParagraphStyle) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setLineBreakMode:"), value)
}

func (m_ MutableParagraphStyle) SetLineBreakStrategy(value LineBreakStrategy) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setLineBreakStrategy:"), value)
}

func (m_ MutableParagraphStyle) SetHyphenationFactor(value float32) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHyphenationFactor:"), value)
}

func (m_ MutableParagraphStyle) SetUsesDefaultHyphenation(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setUsesDefaultHyphenation:"), value)
}

func (m_ MutableParagraphStyle) SetTighteningFactorForTruncation(value float32) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setTighteningFactorForTruncation:"), value)
}

func (m_ MutableParagraphStyle) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setAllowsDefaultTighteningForTruncation:"), value)
}

func (m_ MutableParagraphStyle) SetHeaderLevel(value int) {
	objc.CallMethod[objc.Void](m_, objc.GetSelector("setHeaderLevel:"), value)
}
