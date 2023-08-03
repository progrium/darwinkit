// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}

type _TableRowViewClass struct {
	objc.Class
}

type ITableRowView interface {
	IView
	DrawBackgroundInRect(dirtyRect foundation.Rect)
	DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect)
	DrawSelectionInRect(dirtyRect foundation.Rect)
	DrawSeparatorInRect(dirtyRect foundation.Rect)
	ViewAtColumn(column int) objc.Object
	IsEmphasized() bool
	SetEmphasized(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	IsFloating() bool
	SetFloating(value bool)
	IsSelected() bool
	SetSelected(value bool)
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
	IndentationForDropOperation() float64
	SetIndentationForDropOperation(value float64)
	IsTargetForDropOperation() bool
	SetTargetForDropOperation(value bool)
	IsGroupRowStyle() bool
	SetGroupRowStyle(value bool)
	NumberOfColumns() int
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	IsNextRowSelected() bool
	SetNextRowSelected(value bool)
	IsPreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
}

type TableRowView struct {
	View
}

func MakeTableRowView(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		View: MakeView(ptr),
	}
}

func (t_ TableRowView) InitWithFrame(frameRect foundation.Rect) TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TableRowView_InitWithFrame(frameRect foundation.Rect) TableRowView {
	return TableRowViewClass.Alloc().InitWithFrame(frameRect)
}

func (t_ TableRowView) Init() TableRowView {
	rv := objc.CallMethod[TableRowView](t_, objc.GetSelector("init"))
	return rv
}

func TableRowView_Init() TableRowView {
	return TableRowViewClass.Alloc().Init()
}

func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.CallMethod[TableRowView](tc, objc.GetSelector("alloc"))
	return rv
}

func TableRowView_Alloc() TableRowView {
	return TableRowViewClass.Alloc()
}

func (tc _TableRowViewClass) New() TableRowView {
	rv := objc.CallMethod[TableRowView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableRowView() TableRowView {
	return TableRowViewClass.New()
}

func TableRowView_New() TableRowView {
	return TableRowViewClass.New()
}

func (t_ TableRowView) DrawBackgroundInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawBackgroundInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawSelectionInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawSelectionInRect:"), dirtyRect)
}

func (t_ TableRowView) DrawSeparatorInRect(dirtyRect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawSeparatorInRect:"), dirtyRect)
}

func (t_ TableRowView) ViewAtColumn(column int) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("viewAtColumn:"), column)
	return rv
}

func (t_ TableRowView) IsEmphasized() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEmphasized"))
	return rv
}

func (t_ TableRowView) SetEmphasized(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEmphasized:"), value)
}

func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.CallMethod[BackgroundStyle](t_, objc.GetSelector("interiorBackgroundStyle"))
	return rv
}

func (t_ TableRowView) IsFloating() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isFloating"))
	return rv
}

func (t_ TableRowView) SetFloating(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFloating:"), value)
}

func (t_ TableRowView) IsSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isSelected"))
	return rv
}

func (t_ TableRowView) SetSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelected:"), value)
}

func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.CallMethod[TableViewSelectionHighlightStyle](t_, objc.GetSelector("selectionHighlightStyle"))
	return rv
}

func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectionHighlightStyle:"), value)
}

func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.CallMethod[TableViewDraggingDestinationFeedbackStyle](t_, objc.GetSelector("draggingDestinationFeedbackStyle"))
	return rv
}

func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDraggingDestinationFeedbackStyle:"), value)
}

func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("indentationForDropOperation"))
	return rv
}

func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIndentationForDropOperation:"), value)
}

func (t_ TableRowView) IsTargetForDropOperation() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isTargetForDropOperation"))
	return rv
}

func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTargetForDropOperation:"), value)
}

func (t_ TableRowView) IsGroupRowStyle() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isGroupRowStyle"))
	return rv
}

func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGroupRowStyle:"), value)
}

func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("numberOfColumns"))
	return rv
}

func (t_ TableRowView) BackgroundColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("backgroundColor"))
	return rv
}

func (t_ TableRowView) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (t_ TableRowView) IsNextRowSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isNextRowSelected"))
	return rv
}

func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setNextRowSelected:"), value)
}

func (t_ TableRowView) IsPreviousRowSelected() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isPreviousRowSelected"))
	return rv
}

func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setPreviousRowSelected:"), value)
}
