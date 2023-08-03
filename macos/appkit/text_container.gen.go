// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextContainerClass = _TextContainerClass{objc.GetClass("NSTextContainer")}

type _TextContainerClass struct {
	objc.Class
}

type ITextContainer interface {
	objc.IObject
	ReplaceLayoutManager(newLayoutManager ILayoutManager)
	LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect foundation.Rect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect *foundation.Rect) foundation.Rect
	LayoutManager() LayoutManager
	SetLayoutManager(value ILayoutManager)
	TextLayoutManager() TextLayoutManager
	TextView() TextView
	SetTextView(value ITextView)
	Size() foundation.Size
	SetSize(value foundation.Size)
	ExclusionPaths() []BezierPath
	SetExclusionPaths(value []IBezierPath)
	LineBreakMode() LineBreakMode
	SetLineBreakMode(value LineBreakMode)
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	IsSimpleRectangularTextContainer() bool
}

type TextContainer struct {
	objc.Object
}

func MakeTextContainer(ptr unsafe.Pointer) TextContainer {
	return TextContainer{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextContainer) InitWithSize(size foundation.Size) TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("initWithSize:"), size)
	return rv
}

func TextContainer_InitWithSize(size foundation.Size) TextContainer {
	return TextContainerClass.Alloc().InitWithSize(size)
}

func (t_ TextContainer) InitWithContainerSize(aContainerSize foundation.Size) TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("initWithContainerSize:"), aContainerSize)
	return rv
}

func TextContainer_InitWithContainerSize(aContainerSize foundation.Size) TextContainer {
	return TextContainerClass.Alloc().InitWithContainerSize(aContainerSize)
}

func (tc _TextContainerClass) Alloc() TextContainer {
	rv := objc.CallMethod[TextContainer](tc, objc.GetSelector("alloc"))
	return rv
}

func TextContainer_Alloc() TextContainer {
	return TextContainerClass.Alloc()
}

func (tc _TextContainerClass) New() TextContainer {
	rv := objc.CallMethod[TextContainer](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextContainer() TextContainer {
	return TextContainerClass.New()
}

func TextContainer_New() TextContainer {
	return TextContainerClass.New()
}

func (t_ TextContainer) Init() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("init"))
	return rv
}

func TextContainer_Init() TextContainer {
	return TextContainerClass.Alloc().Init()
}

func (t_ TextContainer) ReplaceLayoutManager(newLayoutManager ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceLayoutManager:"), objc.ExtractPtr(newLayoutManager))
}

func (t_ TextContainer) LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect foundation.Rect, characterIndex uint, baseWritingDirection WritingDirection, remainingRect *foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](t_, objc.GetSelector("lineFragmentRectForProposedRect:atIndex:writingDirection:remainingRect:"), proposedRect, characterIndex, baseWritingDirection, remainingRect)
	return rv
}

func (t_ TextContainer) LayoutManager() LayoutManager {
	rv := objc.CallMethod[LayoutManager](t_, objc.GetSelector("layoutManager"))
	return rv
}

func (t_ TextContainer) SetLayoutManager(value ILayoutManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutManager:"), objc.ExtractPtr(value))
}

func (t_ TextContainer) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("textLayoutManager"))
	return rv
}

func (t_ TextContainer) TextView() TextView {
	rv := objc.CallMethod[TextView](t_, objc.GetSelector("textView"))
	return rv
}

func (t_ TextContainer) SetTextView(value ITextView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextView:"), objc.ExtractPtr(value))
}

func (t_ TextContainer) Size() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("size"))
	return rv
}

func (t_ TextContainer) SetSize(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSize:"), value)
}

func (t_ TextContainer) ExclusionPaths() []BezierPath {
	rv := objc.CallMethod[[]BezierPath](t_, objc.GetSelector("exclusionPaths"))
	// TODO: convert slice items...
	return rv
}

func (t_ TextContainer) SetExclusionPaths(value []IBezierPath) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setExclusionPaths:"), value)
}

func (t_ TextContainer) LineBreakMode() LineBreakMode {
	rv := objc.CallMethod[LineBreakMode](t_, objc.GetSelector("lineBreakMode"))
	return rv
}

func (t_ TextContainer) SetLineBreakMode(value LineBreakMode) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineBreakMode:"), value)
}

func (t_ TextContainer) WidthTracksTextView() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("widthTracksTextView"))
	return rv
}

func (t_ TextContainer) SetWidthTracksTextView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWidthTracksTextView:"), value)
}

func (t_ TextContainer) HeightTracksTextView() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("heightTracksTextView"))
	return rv
}

func (t_ TextContainer) SetHeightTracksTextView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHeightTracksTextView:"), value)
}

func (t_ TextContainer) MaximumNumberOfLines() uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("maximumNumberOfLines"))
	return rv
}

func (t_ TextContainer) SetMaximumNumberOfLines(value uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMaximumNumberOfLines:"), value)
}

func (t_ TextContainer) LineFragmentPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("lineFragmentPadding"))
	return rv
}

func (t_ TextContainer) SetLineFragmentPadding(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineFragmentPadding:"), value)
}

func (t_ TextContainer) IsSimpleRectangularTextContainer() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isSimpleRectangularTextContainer"))
	return rv
}
