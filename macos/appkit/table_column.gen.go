// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}

type _TableColumnClass struct {
	objc.Class
}

type ITableColumn interface {
	objc.IObject
	SizeToFit()
	TableView() TableView
	SetTableView(value ITableView)
	Width() float64
	SetWidth(value float64)
	MinWidth() float64
	SetMinWidth(value float64)
	MaxWidth() float64
	SetMaxWidth(value float64)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	Title() string
	SetTitle(value string)
	HeaderCell() TableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	IsEditable() bool
	SetEditable(value bool)
	SortDescriptorPrototype() foundation.SortDescriptor
	SetSortDescriptorPrototype(value foundation.ISortDescriptor)
	IsHidden() bool
	SetHidden(value bool)
	HeaderToolTip() string
	SetHeaderToolTip(value string)
}

type TableColumn struct {
	objc.Object
}

func MakeTableColumn(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TableColumn) InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("initWithIdentifier:"), identifier)
	return rv
}

func TableColumn_InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	return TableColumnClass.Alloc().InitWithIdentifier(identifier)
}

func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.CallMethod[TableColumn](tc, objc.GetSelector("alloc"))
	return rv
}

func TableColumn_Alloc() TableColumn {
	return TableColumnClass.Alloc()
}

func (tc _TableColumnClass) New() TableColumn {
	rv := objc.CallMethod[TableColumn](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTableColumn() TableColumn {
	return TableColumnClass.New()
}

func TableColumn_New() TableColumn {
	return TableColumnClass.New()
}

func (t_ TableColumn) Init() TableColumn {
	rv := objc.CallMethod[TableColumn](t_, objc.GetSelector("init"))
	return rv
}

func TableColumn_Init() TableColumn {
	return TableColumnClass.Alloc().Init()
}

func (t_ TableColumn) SizeToFit() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("sizeToFit"))
}

func (t_ TableColumn) TableView() TableView {
	rv := objc.CallMethod[TableView](t_, objc.GetSelector("tableView"))
	return rv
}

func (t_ TableColumn) SetTableView(value ITableView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTableView:"), objc.ExtractPtr(value))
}

func (t_ TableColumn) Width() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("width"))
	return rv
}

func (t_ TableColumn) SetWidth(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setWidth:"), value)
}

func (t_ TableColumn) MinWidth() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("minWidth"))
	return rv
}

func (t_ TableColumn) SetMinWidth(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMinWidth:"), value)
}

func (t_ TableColumn) MaxWidth() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("maxWidth"))
	return rv
}

func (t_ TableColumn) SetMaxWidth(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMaxWidth:"), value)
}

func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := objc.CallMethod[TableColumnResizingOptions](t_, objc.GetSelector("resizingMask"))
	return rv
}

func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setResizingMask:"), value)
}

func (t_ TableColumn) Title() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("title"))
	return rv
}

func (t_ TableColumn) SetTitle(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTitle:"), value)
}

func (t_ TableColumn) HeaderCell() TableHeaderCell {
	rv := objc.CallMethod[TableHeaderCell](t_, objc.GetSelector("headerCell"))
	return rv
}

func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHeaderCell:"), objc.ExtractPtr(value))
}

func (t_ TableColumn) Identifier() UserInterfaceItemIdentifier {
	rv := objc.CallMethod[UserInterfaceItemIdentifier](t_, objc.GetSelector("identifier"))
	return rv
}

func (t_ TableColumn) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIdentifier:"), value)
}

func (t_ TableColumn) IsEditable() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isEditable"))
	return rv
}

func (t_ TableColumn) SetEditable(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEditable:"), value)
}

func (t_ TableColumn) SortDescriptorPrototype() foundation.SortDescriptor {
	rv := objc.CallMethod[foundation.SortDescriptor](t_, objc.GetSelector("sortDescriptorPrototype"))
	return rv
}

func (t_ TableColumn) SetSortDescriptorPrototype(value foundation.ISortDescriptor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSortDescriptorPrototype:"), objc.ExtractPtr(value))
}

func (t_ TableColumn) IsHidden() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isHidden"))
	return rv
}

func (t_ TableColumn) SetHidden(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHidden:"), value)
}

func (t_ TableColumn) HeaderToolTip() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("headerToolTip"))
	return rv
}

func (t_ TableColumn) SetHeaderToolTip(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setHeaderToolTip:"), value)
}
