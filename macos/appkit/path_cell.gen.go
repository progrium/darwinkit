// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var PathCellClass = _PathCellClass{objc.GetClass("NSPathCell")}

type _PathCellClass struct {
	objc.Class
}

type IPathCell interface {
	IActionCell
	MouseEnteredWithFrameInView(event IEvent, frame foundation.Rect, view IView)
	MouseExitedWithFrameInView(event IEvent, frame foundation.Rect, view IView)
	RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect
	PathComponentCellAtPointWithFrameInView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.IAttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ClickedPathComponentCell() PathComponentCell
	PathComponentCells() []PathComponentCell
	SetPathComponentCells(value []IPathComponentCell)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	Delegate() PathCellDelegateWrapper
	SetDelegate(value IPathCellDelegate)
	SetDelegate0(value objc.IObject)
}

type PathCell struct {
	ActionCell
}

func MakePathCell(ptr unsafe.Pointer) PathCell {
	return PathCell{
		ActionCell: MakeActionCell(ptr),
	}
}

func (p_ PathCell) InitImageCell(image IImage) PathCell {
	rv := objc.CallMethod[PathCell](p_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func PathCell_InitImageCell(image IImage) PathCell {
	return PathCellClass.Alloc().InitImageCell(image)
}

func (p_ PathCell) InitTextCell(string_ string) PathCell {
	rv := objc.CallMethod[PathCell](p_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func PathCell_InitTextCell(string_ string) PathCell {
	return PathCellClass.Alloc().InitTextCell(string_)
}

func (p_ PathCell) Init() PathCell {
	rv := objc.CallMethod[PathCell](p_, objc.GetSelector("init"))
	return rv
}

func PathCell_Init() PathCell {
	return PathCellClass.Alloc().Init()
}

func (pc _PathCellClass) Alloc() PathCell {
	rv := objc.CallMethod[PathCell](pc, objc.GetSelector("alloc"))
	return rv
}

func PathCell_Alloc() PathCell {
	return PathCellClass.Alloc()
}

func (pc _PathCellClass) New() PathCell {
	rv := objc.CallMethod[PathCell](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPathCell() PathCell {
	return PathCellClass.New()
}

func PathCell_New() PathCell {
	return PathCellClass.New()
}

func (p_ PathCell) MouseEnteredWithFrameInView(event IEvent, frame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("mouseEntered:withFrame:inView:"), objc.ExtractPtr(event), frame, objc.ExtractPtr(view))
}

func (p_ PathCell) MouseExitedWithFrameInView(event IEvent, frame foundation.Rect, view IView) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("mouseExited:withFrame:inView:"), objc.ExtractPtr(event), frame, objc.ExtractPtr(view))
}

func (p_ PathCell) RectOfPathComponentCellWithFrameInView(cell IPathComponentCell, frame foundation.Rect, view IView) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](p_, objc.GetSelector("rectOfPathComponentCell:withFrame:inView:"), objc.ExtractPtr(cell), frame, objc.ExtractPtr(view))
	return rv
}

func (p_ PathCell) PathComponentCellAtPointWithFrameInView(point foundation.Point, frame foundation.Rect, view IView) PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("pathComponentCellAtPoint:withFrame:inView:"), point, frame, objc.ExtractPtr(view))
	return rv
}

func (p_ PathCell) AllowedTypes() []string {
	rv := objc.CallMethod[[]string](p_, objc.GetSelector("allowedTypes"))
	return rv
}

func (p_ PathCell) SetAllowedTypes(value []string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setAllowedTypes:"), value)
}

func (p_ PathCell) PathStyle() PathStyle {
	rv := objc.CallMethod[PathStyle](p_, objc.GetSelector("pathStyle"))
	return rv
}

func (p_ PathCell) SetPathStyle(value PathStyle) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPathStyle:"), value)
}

func (p_ PathCell) PlaceholderAttributedString() foundation.AttributedString {
	rv := objc.CallMethod[foundation.AttributedString](p_, objc.GetSelector("placeholderAttributedString"))
	return rv
}

func (p_ PathCell) SetPlaceholderAttributedString(value foundation.IAttributedString) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPlaceholderAttributedString:"), objc.ExtractPtr(value))
}

func (p_ PathCell) PlaceholderString() string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("placeholderString"))
	return rv
}

func (p_ PathCell) SetPlaceholderString(value string) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPlaceholderString:"), value)
}

func (p_ PathCell) BackgroundColor() Color {
	rv := objc.CallMethod[Color](p_, objc.GetSelector("backgroundColor"))
	return rv
}

func (p_ PathCell) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (pc _PathCellClass) PathComponentCellClass() objc.Class {
	rv := objc.CallMethod[objc.Class](pc, objc.GetSelector("pathComponentCellClass"))
	return rv
}

func PathCell_PathComponentCellClass() objc.Class {
	return PathCellClass.PathComponentCellClass()
}

func (p_ PathCell) ClickedPathComponentCell() PathComponentCell {
	rv := objc.CallMethod[PathComponentCell](p_, objc.GetSelector("clickedPathComponentCell"))
	return rv
}

func (p_ PathCell) PathComponentCells() []PathComponentCell {
	rv := objc.CallMethod[[]PathComponentCell](p_, objc.GetSelector("pathComponentCells"))
	return rv
}

func (p_ PathCell) SetPathComponentCells(value []IPathComponentCell) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setPathComponentCells:"), value)
}

func (p_ PathCell) DoubleAction() objc.Selector {
	rv := objc.CallMethod[objc.Selector](p_, objc.GetSelector("doubleAction"))
	return rv
}

func (p_ PathCell) SetDoubleAction(value objc.Selector) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDoubleAction:"), value)
}

func (p_ PathCell) URL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](p_, objc.GetSelector("URL"))
	return rv
}

func (p_ PathCell) SetURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setURL:"), objc.ExtractPtr(value))
}

func (p_ PathCell) Delegate() PathCellDelegateWrapper {
	rv := objc.CallMethod[PathCellDelegateWrapper](p_, objc.GetSelector("delegate"))
	return rv
}

func (p_ PathCell) SetDelegate(value IPathCellDelegate) {
	po := objc.WrapAsProtocol("NSPathCellDelegate", value)
	objc.SetAssociatedObject(p_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), po)
}

func (p_ PathCell) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](p_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}
