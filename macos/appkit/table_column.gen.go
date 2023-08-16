// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableColumn] class.
var TableColumnClass = _TableColumnClass{objc.GetClass("NSTableColumn")}

type _TableColumnClass struct {
	objc.Class
}

// An interface definition for the [TableColumn] class.
type ITableColumn interface {
	objc.IObject
	SizeToFit()
	Width() float64
	SetWidth(value float64)
	IsHidden() bool
	SetHidden(value bool)
	IsEditable() bool
	SetEditable(value bool)
	MinWidth() float64
	SetMinWidth(value float64)
	HeaderToolTip() string
	SetHeaderToolTip(value string)
	HeaderCell() TableHeaderCell
	SetHeaderCell(value ITableHeaderCell)
	MaxWidth() float64
	SetMaxWidth(value float64)
	TableView() TableView
	SetTableView(value ITableView)
	Title() string
	SetTitle(value string)
	ResizingMask() TableColumnResizingOptions
	SetResizingMask(value TableColumnResizingOptions)
	Identifier() UserInterfaceItemIdentifier
	SetIdentifier(value UserInterfaceItemIdentifier)
	SortDescriptorPrototype() foundation.SortDescriptor
	SetSortDescriptorPrototype(value foundation.ISortDescriptor)
}

// The display characteristics and identifier for a column in a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn?language=objc
type TableColumn struct {
	objc.Object
}

func TableColumnFrom(ptr unsafe.Pointer) TableColumn {
	return TableColumn{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TableColumn) InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	rv := objc.Call[TableColumn](t_, objc.Sel("initWithIdentifier:"), identifier)
	return rv
}

// Initializes a newly created table column with a string identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1526749-initwithidentifier?language=objc
func TableColumn_InitWithIdentifier(identifier UserInterfaceItemIdentifier) TableColumn {
	return TableColumnClass.Alloc().InitWithIdentifier(identifier)
}

func (tc _TableColumnClass) Alloc() TableColumn {
	rv := objc.Call[TableColumn](tc, objc.Sel("alloc"))
	return rv
}

func TableColumn_Alloc() TableColumn {
	return TableColumnClass.Alloc()
}

func (tc _TableColumnClass) New() TableColumn {
	rv := objc.Call[TableColumn](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableColumn() TableColumn {
	return TableColumnClass.New()
}

func (t_ TableColumn) Init() TableColumn {
	rv := objc.Call[TableColumn](t_, objc.Sel("init"))
	return rv
}

// Resizes the table column to fit the width of its header cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1532997-sizetofit?language=objc
func (t_ TableColumn) SizeToFit() {
	objc.Call[objc.Void](t_, objc.Sel("sizeToFit"))
}

// The table column’s width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1528774-width?language=objc
func (t_ TableColumn) Width() float64 {
	rv := objc.Call[float64](t_, objc.Sel("width"))
	return rv
}

// The table column’s width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1528774-width?language=objc
func (t_ TableColumn) SetWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setWidth:"), value)
}

// A Boolean that indicates whether the table column is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1524681-hidden?language=objc
func (t_ TableColumn) IsHidden() bool {
	rv := objc.Call[bool](t_, objc.Sel("isHidden"))
	return rv
}

// A Boolean that indicates whether the table column is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1524681-hidden?language=objc
func (t_ TableColumn) SetHidden(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setHidden:"), value)
}

// A Boolean that indicates whether a cell-based table’s column cells are user editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1528412-editable?language=objc
func (t_ TableColumn) IsEditable() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEditable"))
	return rv
}

// A Boolean that indicates whether a cell-based table’s column cells are user editable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1528412-editable?language=objc
func (t_ TableColumn) SetEditable(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setEditable:"), value)
}

// The table column’s minimum width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1525126-minwidth?language=objc
func (t_ TableColumn) MinWidth() float64 {
	rv := objc.Call[float64](t_, objc.Sel("minWidth"))
	return rv
}

// The table column’s minimum width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1525126-minwidth?language=objc
func (t_ TableColumn) SetMinWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setMinWidth:"), value)
}

// The string that’s displayed in a help tag over the table column header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1524685-headertooltip?language=objc
func (t_ TableColumn) HeaderToolTip() string {
	rv := objc.Call[string](t_, objc.Sel("headerToolTip"))
	return rv
}

// The string that’s displayed in a help tag over the table column header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1524685-headertooltip?language=objc
func (t_ TableColumn) SetHeaderToolTip(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setHeaderToolTip:"), value)
}

// The cell used to draw the table column’s header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1525137-headercell?language=objc
func (t_ TableColumn) HeaderCell() TableHeaderCell {
	rv := objc.Call[TableHeaderCell](t_, objc.Sel("headerCell"))
	return rv
}

// The cell used to draw the table column’s header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1525137-headercell?language=objc
func (t_ TableColumn) SetHeaderCell(value ITableHeaderCell) {
	objc.Call[objc.Void](t_, objc.Sel("setHeaderCell:"), objc.Ptr(value))
}

// The table column’s maximum width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1526342-maxwidth?language=objc
func (t_ TableColumn) MaxWidth() float64 {
	rv := objc.Call[float64](t_, objc.Sel("maxWidth"))
	return rv
}

// The table column’s maximum width, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1526342-maxwidth?language=objc
func (t_ TableColumn) SetMaxWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setMaxWidth:"), value)
}

// The table view that contains the table column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1535908-tableview?language=objc
func (t_ TableColumn) TableView() TableView {
	rv := objc.Call[TableView](t_, objc.Sel("tableView"))
	return rv
}

// The table view that contains the table column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1535908-tableview?language=objc
func (t_ TableColumn) SetTableView(value ITableView) {
	objc.Call[objc.Void](t_, objc.Sel("setTableView:"), objc.Ptr(value))
}

// The title of the table column’s header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1526875-title?language=objc
func (t_ TableColumn) Title() string {
	rv := objc.Call[string](t_, objc.Sel("title"))
	return rv
}

// The title of the table column’s header. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1526875-title?language=objc
func (t_ TableColumn) SetTitle(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setTitle:"), value)
}

// The table column’s resizing mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1529591-resizingmask?language=objc
func (t_ TableColumn) ResizingMask() TableColumnResizingOptions {
	rv := objc.Call[TableColumnResizingOptions](t_, objc.Sel("resizingMask"))
	return rv
}

// The table column’s resizing mask. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1529591-resizingmask?language=objc
func (t_ TableColumn) SetResizingMask(value TableColumnResizingOptions) {
	objc.Call[objc.Void](t_, objc.Sel("setResizingMask:"), value)
}

// The identifier string for the table column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1531113-identifier?language=objc
func (t_ TableColumn) Identifier() UserInterfaceItemIdentifier {
	rv := objc.Call[UserInterfaceItemIdentifier](t_, objc.Sel("identifier"))
	return rv
}

// The identifier string for the table column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1531113-identifier?language=objc
func (t_ TableColumn) SetIdentifier(value UserInterfaceItemIdentifier) {
	objc.Call[objc.Void](t_, objc.Sel("setIdentifier:"), value)
}

// The table column’s sort descriptor prototype. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1534663-sortdescriptorprototype?language=objc
func (t_ TableColumn) SortDescriptorPrototype() foundation.SortDescriptor {
	rv := objc.Call[foundation.SortDescriptor](t_, objc.Sel("sortDescriptorPrototype"))
	return rv
}

// The table column’s sort descriptor prototype. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablecolumn/1534663-sortdescriptorprototype?language=objc
func (t_ TableColumn) SetSortDescriptorPrototype(value foundation.ISortDescriptor) {
	objc.Call[objc.Void](t_, objc.Sel("setSortDescriptorPrototype:"), objc.Ptr(value))
}
