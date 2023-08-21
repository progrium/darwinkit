// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Matrix] class.
var MatrixClass = _MatrixClass{objc.GetClass("NSMatrix")}

type _MatrixClass struct {
	objc.Class
}

// An interface definition for the [Matrix] class.
type IMatrix interface {
	IControl
	SendAction() bool
	DeselectSelectedCell()
	AddColumnWithCells(newCells []ICell)
	CellFrameAtRowColumn(row int, col int) foundation.Rect
	SetToolTipForCell(toolTipString string, cell ICell)
	AddRow()
	RemoveRow(row int)
	SortUsingFunctionContext(compare func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer)
	SizeToCells()
	ScrollCellToVisibleAtRowColumn(row int, col int)
	DrawCellAtRowColumn(row int, col int)
	SelectCellWithTag(tag int) bool
	PutCellAtRowColumn(newCell ICell, row int, col int)
	HighlightCellAtRowColumn(flag bool, row int, col int)
	SelectAll(sender objc.IObject)
	SelectText(sender objc.IObject)
	SelectTextAtRowColumn(row int, col int) Cell
	TextShouldBeginEditing(textObject IText) bool
	AddRowWithCells(newCells []ICell)
	SortUsingSelector(comparator objc.Selector)
	InsertRow(row int)
	InsertColumnWithCells(column int, newCells []ICell)
	TextDidChange(notification foundation.INotification)
	CellWithTag(tag int) Cell
	CellAtRowColumn(row int, col int) Cell
	TextDidBeginEditing(notification foundation.INotification)
	SetScrollable(flag bool)
	SendDoubleAction()
	SelectCellAtRowColumn(row int, col int)
	RemoveColumn(col int)
	DeselectAllCells()
	MakeCellAtRowColumn(row int, col int) Cell
	SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool)
	ToolTipForCell(cell ICell) string
	SetValidateSize(flag bool)
	TextShouldEndEditing(textObject IText) bool
	AddColumn()
	RenewRowsColumns(newRows int, newCols int)
	GetRowColumnOfCell(row *int, col *int, cell ICell) bool
	TextDidEndEditing(notification foundation.INotification)
	GetNumberOfRowsColumns(rowCount *int, colCount *int)
	SetStateAtRowColumn(value int, row int, col int)
	MouseDownFlags() int
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	SelectedCells() []Cell
	TabKeyTraversesCells() bool
	SetTabKeyTraversesCells(value bool)
	SelectedCell() Cell
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	Cells() []Cell
	CellBackgroundColor() Color
	SetCellBackgroundColor(value IColor)
	IsSelectionByRect() bool
	SetSelectionByRect(value bool)
	IsAutoscroll() bool
	SetAutoscroll(value bool)
	DrawsCellBackground() bool
	SetDrawsCellBackground(value bool)
	AutorecalculatesCellSize() bool
	SetAutorecalculatesCellSize(value bool)
	Delegate() MatrixDelegateWrapper
	SetDelegate(value PMatrixDelegate)
	SetDelegateObject(valueObject objc.IObject)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	CellSize() foundation.Size
	SetCellSize(value foundation.Size)
	KeyCell() Cell
	SetKeyCell(value ICell)
	SelectedColumn() int
	SelectedRow() int
	Mode() MatrixMode
	SetMode(value MatrixMode)
	AutosizesCells() bool
	SetAutosizesCells(value bool)
	CellClass() objc.Class
	SetCellClass(value objc.IClass)
	Prototype() Cell
	SetPrototype(value ICell)
	NumberOfRows() int
	IntercellSpacing() foundation.Size
	SetIntercellSpacing(value foundation.Size)
	NumberOfColumns() int
}

// A legacy interface for grouping radio buttons or other types of cells together. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix?language=objc
type Matrix struct {
	Control
}

func MatrixFrom(ptr unsafe.Pointer) Matrix {
	return Matrix{
		Control: ControlFrom(ptr),
	}
}

func (m_ Matrix) InitWithFrame(frameRect foundation.Rect) Matrix {
	rv := objc.Call[Matrix](m_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a newly allocated matrix with the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436428-initwithframe?language=objc
func NewMatrixWithFrame(frameRect foundation.Rect) Matrix {
	instance := MatrixClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (mc _MatrixClass) Alloc() Matrix {
	rv := objc.Call[Matrix](mc, objc.Sel("alloc"))
	return rv
}

func Matrix_Alloc() Matrix {
	return MatrixClass.Alloc()
}

func (mc _MatrixClass) New() Matrix {
	rv := objc.Call[Matrix](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrix() Matrix {
	return MatrixClass.New()
}

func (m_ Matrix) Init() Matrix {
	rv := objc.Call[Matrix](m_, objc.Sel("init"))
	return rv
}

// If the selected cell has both an action and a target, sends its action to its target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436413-sendaction?language=objc
func (m_ Matrix) SendAction() bool {
	rv := objc.Call[bool](m_, objc.Sel("sendAction"))
	return rv
}

// Deselects the selected cell or cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436440-deselectselectedcell?language=objc
func (m_ Matrix) DeselectSelectedCell() {
	objc.Call[objc.Void](m_, objc.Sel("deselectSelectedCell"))
}

// Adds a new column of cells to the right of the last column, using the given cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436453-addcolumnwithcells?language=objc
func (m_ Matrix) AddColumnWithCells(newCells []ICell) {
	objc.Call[objc.Void](m_, objc.Sel("addColumnWithCells:"), newCells)
}

// Returns the frame rectangle of the cell that would be drawn at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436417-cellframeatrow?language=objc
func (m_ Matrix) CellFrameAtRowColumn(row int, col int) foundation.Rect {
	rv := objc.Call[foundation.Rect](m_, objc.Sel("cellFrameAtRow:column:"), row, col)
	return rv
}

// Sets the tooltip for the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436372-settooltip?language=objc
func (m_ Matrix) SetToolTipForCell(toolTipString string, cell ICell) {
	objc.Call[objc.Void](m_, objc.Sel("setToolTip:forCell:"), toolTipString, objc.Ptr(cell))
}

// Adds a new row of cells below the last row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436414-addrow?language=objc
func (m_ Matrix) AddRow() {
	objc.Call[objc.Void](m_, objc.Sel("addRow"))
}

// Removes the specified row from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436403-removerow?language=objc
func (m_ Matrix) RemoveRow(row int) {
	objc.Call[objc.Void](m_, objc.Sel("removeRow:"), row)
}

// Sorts the receiver’s cells in ascending order as defined by the specified comparison function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436468-sortusingfunction?language=objc
func (m_ Matrix) SortUsingFunctionContext(compare func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingFunction:context:"), compare, context)
}

// Changes the width and the height of the receiver’s frame so it exactly contains the cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436408-sizetocells?language=objc
func (m_ Matrix) SizeToCells() {
	objc.Call[objc.Void](m_, objc.Sel("sizeToCells"))
}

// Scrolls the receiver so the specified cell is visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436455-scrollcelltovisibleatrow?language=objc
func (m_ Matrix) ScrollCellToVisibleAtRowColumn(row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("scrollCellToVisibleAtRow:column:"), row, col)
}

// Displays the cell at the specified row and column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436495-drawcellatrow?language=objc
func (m_ Matrix) DrawCellAtRowColumn(row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("drawCellAtRow:column:"), row, col)
}

// Selects the last cell with the given tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436446-selectcellwithtag?language=objc
func (m_ Matrix) SelectCellWithTag(tag int) bool {
	rv := objc.Call[bool](m_, objc.Sel("selectCellWithTag:"), tag)
	return rv
}

// Replaces the cell at the specified row and column with the new cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436384-putcell?language=objc
func (m_ Matrix) PutCellAtRowColumn(newCell ICell, row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("putCell:atRow:column:"), objc.Ptr(newCell), row, col)
}

// Highlights or unhighlights the cell at the specified row and column location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436499-highlightcell?language=objc
func (m_ Matrix) HighlightCellAtRowColumn(flag bool, row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("highlightCell:atRow:column:"), flag, row, col)
}

// Selects and highlights all cells in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436378-selectall?language=objc
func (m_ Matrix) SelectAll(sender objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("selectAll:"), sender)
}

// Selects text in the currently selected cell or in the key cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436470-selecttext?language=objc
func (m_ Matrix) SelectText(sender objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("selectText:"), sender)
}

// Selects the text in the cell at the specified location and returns the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436459-selecttextatrow?language=objc
func (m_ Matrix) SelectTextAtRowColumn(row int, col int) Cell {
	rv := objc.Call[Cell](m_, objc.Sel("selectTextAtRow:column:"), row, col)
	return rv
}

// Requests permission to begin editing text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436393-textshouldbeginediting?language=objc
func (m_ Matrix) TextShouldBeginEditing(textObject IText) bool {
	rv := objc.Call[bool](m_, objc.Sel("textShouldBeginEditing:"), objc.Ptr(textObject))
	return rv
}

// Adds a new row of cells below the last row, using the specified cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436450-addrowwithcells?language=objc
func (m_ Matrix) AddRowWithCells(newCells []ICell) {
	objc.Call[objc.Void](m_, objc.Sel("addRowWithCells:"), newCells)
}

// Sorts the receiver’s cells in ascending order as defined by the comparison method. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436475-sortusingselector?language=objc
func (m_ Matrix) SortUsingSelector(comparator objc.Selector) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingSelector:"), comparator)
}

// Inserts a new row of cells before the specified row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436427-insertrow?language=objc
func (m_ Matrix) InsertRow(row int) {
	objc.Call[objc.Void](m_, objc.Sel("insertRow:"), row)
}

// Inserts a new column of cells before the specified column, using the given cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436483-insertcolumn?language=objc
func (m_ Matrix) InsertColumnWithCells(column int, newCells []ICell) {
	objc.Call[objc.Void](m_, objc.Sel("insertColumn:withCells:"), column, newCells)
}

// Invoked when a key-down event or paste operation occurs that changes the receiver’s contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436489-textdidchange?language=objc
func (m_ Matrix) TextDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](m_, objc.Sel("textDidChange:"), objc.Ptr(notification))
}

// Searches the receiver and returns the last cell matching the specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436481-cellwithtag?language=objc
func (m_ Matrix) CellWithTag(tag int) Cell {
	rv := objc.Call[Cell](m_, objc.Sel("cellWithTag:"), tag)
	return rv
}

// Returns the cell at the specified row and column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436424-cellatrow?language=objc
func (m_ Matrix) CellAtRowColumn(row int, col int) Cell {
	rv := objc.Call[Cell](m_, objc.Sel("cellAtRow:column:"), row, col)
	return rv
}

// Invoked when there’s a change in the text after the receiver gains first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436395-textdidbeginediting?language=objc
func (m_ Matrix) TextDidBeginEditing(notification foundation.INotification) {
	objc.Call[objc.Void](m_, objc.Sel("textDidBeginEditing:"), objc.Ptr(notification))
}

// Specifies whether the cells in the matrix are scrollable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436420-setscrollable?language=objc
func (m_ Matrix) SetScrollable(flag bool) {
	objc.Call[objc.Void](m_, objc.Sel("setScrollable:"), flag)
}

// Sends the double-click action message to the target of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436436-senddoubleaction?language=objc
func (m_ Matrix) SendDoubleAction() {
	objc.Call[objc.Void](m_, objc.Sel("sendDoubleAction"))
}

// Selects the cell at the specified row and column within the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436399-selectcellatrow?language=objc
func (m_ Matrix) SelectCellAtRowColumn(row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("selectCellAtRow:column:"), row, col)
}

// Removes the specified column at from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436397-removecolumn?language=objc
func (m_ Matrix) RemoveColumn(col int) {
	objc.Call[objc.Void](m_, objc.Sel("removeColumn:"), col)
}

// Deselects all cells in the receiver and, if necessary, redisplays the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436396-deselectallcells?language=objc
func (m_ Matrix) DeselectAllCells() {
	objc.Call[objc.Void](m_, objc.Sel("deselectAllCells"))
}

// Creates a new cell at the location specified by the given row and column in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436439-makecellatrow?language=objc
func (m_ Matrix) MakeCellAtRowColumn(row int, col int) Cell {
	rv := objc.Call[Cell](m_, objc.Sel("makeCellAtRow:column:"), row, col)
	return rv
}

// Programmatically selects a range of cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436405-setselectionfrom?language=objc
func (m_ Matrix) SetSelectionFromToAnchorHighlight(startPos int, endPos int, anchorPos int, lit bool) {
	objc.Call[objc.Void](m_, objc.Sel("setSelectionFrom:to:anchor:highlight:"), startPos, endPos, anchorPos, lit)
}

// Returns the tooltip for the specified cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436401-tooltipforcell?language=objc
func (m_ Matrix) ToolTipForCell(cell ICell) string {
	rv := objc.Call[string](m_, objc.Sel("toolTipForCell:"), objc.Ptr(cell))
	return rv
}

// Specifies whether the receiver's size information is validated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436503-setvalidatesize?language=objc
func (m_ Matrix) SetValidateSize(flag bool) {
	objc.Call[objc.Void](m_, objc.Sel("setValidateSize:"), flag)
}

// Requests permission to end editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436509-textshouldendediting?language=objc
func (m_ Matrix) TextShouldEndEditing(textObject IText) bool {
	rv := objc.Call[bool](m_, objc.Sel("textShouldEndEditing:"), objc.Ptr(textObject))
	return rv
}

// Adds a new column of cells to the right of the last column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436419-addcolumn?language=objc
func (m_ Matrix) AddColumn() {
	objc.Call[objc.Void](m_, objc.Sel("addColumn"))
}

// Changes the number of rows and columns in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436452-renewrows?language=objc
func (m_ Matrix) RenewRowsColumns(newRows int, newCols int) {
	objc.Call[objc.Void](m_, objc.Sel("renewRows:columns:"), newRows, newCols)
}

// Searches the receiver for the specified cell and returns the row and column of the cell [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436421-getrow?language=objc
func (m_ Matrix) GetRowColumnOfCell(row *int, col *int, cell ICell) bool {
	rv := objc.Call[bool](m_, objc.Sel("getRow:column:ofCell:"), row, col, objc.Ptr(cell))
	return rv
}

// Invoked when text editing ends. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436426-textdidendediting?language=objc
func (m_ Matrix) TextDidEndEditing(notification foundation.INotification) {
	objc.Call[objc.Void](m_, objc.Sel("textDidEndEditing:"), objc.Ptr(notification))
}

// Obtains the number of rows and columns in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436423-getnumberofrows?language=objc
func (m_ Matrix) GetNumberOfRowsColumns(rowCount *int, colCount *int) {
	objc.Call[objc.Void](m_, objc.Sel("getNumberOfRows:columns:"), rowCount, colCount)
}

// Sets the state of the cell at specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436479-setstate?language=objc
func (m_ Matrix) SetStateAtRowColumn(value int, row int, col int) {
	objc.Call[objc.Void](m_, objc.Sel("setState:atRow:column:"), value, row, col)
}

// The flags in effect at the mouse-down event that started the current tracking session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436487-mousedownflags?language=objc
func (m_ Matrix) MouseDownFlags() int {
	rv := objc.Call[int](m_, objc.Sel("mouseDownFlags"))
	return rv
}

// The action sent to the target of the receiver when the user double-clicks a cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436469-doubleaction?language=objc
func (m_ Matrix) DoubleAction() objc.Selector {
	rv := objc.Call[objc.Selector](m_, objc.Sel("doubleAction"))
	return rv
}

// The action sent to the target of the receiver when the user double-clicks a cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436469-doubleaction?language=objc
func (m_ Matrix) SetDoubleAction(value objc.Selector) {
	objc.Call[objc.Void](m_, objc.Sel("setDoubleAction:"), value)
}

// An array containing all of the matrix’s highlighted cells plus its selected cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436434-selectedcells?language=objc
func (m_ Matrix) SelectedCells() []Cell {
	rv := objc.Call[[]Cell](m_, objc.Sel("selectedCells"))
	return rv
}

// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436491-tabkeytraversescells?language=objc
func (m_ Matrix) TabKeyTraversesCells() bool {
	rv := objc.Call[bool](m_, objc.Sel("tabKeyTraversesCells"))
	return rv
}

// A Boolean that indicates whether pressing the Tab key advances the key cell to the next selectable cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436491-tabkeytraversescells?language=objc
func (m_ Matrix) SetTabKeyTraversesCells(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setTabKeyTraversesCells:"), value)
}

// The most recently selected cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436472-selectedcell?language=objc
func (m_ Matrix) SelectedCell() Cell {
	rv := objc.Call[Cell](m_, objc.Sel("selectedCell"))
	return rv
}

// A Boolean that indicates whether a radio-mode matrix supports an empty selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436394-allowsemptyselection?language=objc
func (m_ Matrix) AllowsEmptySelection() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsEmptySelection"))
	return rv
}

// A Boolean that indicates whether a radio-mode matrix supports an empty selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436394-allowsemptyselection?language=objc
func (m_ Matrix) SetAllowsEmptySelection(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsEmptySelection:"), value)
}

// An array containing the cells of the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436464-cells?language=objc
func (m_ Matrix) Cells() []Cell {
	rv := objc.Call[[]Cell](m_, objc.Sel("cells"))
	return rv
}

// The background color of the matrix’s cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436449-cellbackgroundcolor?language=objc
func (m_ Matrix) CellBackgroundColor() Color {
	rv := objc.Call[Color](m_, objc.Sel("cellBackgroundColor"))
	return rv
}

// The background color of the matrix’s cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436449-cellbackgroundcolor?language=objc
func (m_ Matrix) SetCellBackgroundColor(value IColor) {
	objc.Call[objc.Void](m_, objc.Sel("setCellBackgroundColor:"), objc.Ptr(value))
}

// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436431-selectionbyrect?language=objc
func (m_ Matrix) IsSelectionByRect() bool {
	rv := objc.Call[bool](m_, objc.Sel("isSelectionByRect"))
	return rv
}

// A Boolean that indicates whether the user can select a rectangle of cells in the receiver by dragging the cursor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436431-selectionbyrect?language=objc
func (m_ Matrix) SetSelectionByRect(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setSelectionByRect:"), value)
}

// A Boolean that indicates whether the receiver is automatically scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436410-autoscroll?language=objc
func (m_ Matrix) IsAutoscroll() bool {
	rv := objc.Call[bool](m_, objc.Sel("isAutoscroll"))
	return rv
}

// A Boolean that indicates whether the receiver is automatically scrolled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436410-autoscroll?language=objc
func (m_ Matrix) SetAutoscroll(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutoscroll:"), value)
}

// A Boolean that indicates whether the matrix draws the background within each of its cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436379-drawscellbackground?language=objc
func (m_ Matrix) DrawsCellBackground() bool {
	rv := objc.Call[bool](m_, objc.Sel("drawsCellBackground"))
	return rv
}

// A Boolean that indicates whether the matrix draws the background within each of its cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436379-drawscellbackground?language=objc
func (m_ Matrix) SetDrawsCellBackground(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setDrawsCellBackground:"), value)
}

// A Boolean that indicates whether the matrix auto-recalculates its cell size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436501-autorecalculatescellsize?language=objc
func (m_ Matrix) AutorecalculatesCellSize() bool {
	rv := objc.Call[bool](m_, objc.Sel("autorecalculatesCellSize"))
	return rv
}

// A Boolean that indicates whether the matrix auto-recalculates its cell size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436501-autorecalculatescellsize?language=objc
func (m_ Matrix) SetAutorecalculatesCellSize(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutorecalculatesCellSize:"), value)
}

// The delegate for messages from the field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436404-delegate?language=objc
func (m_ Matrix) Delegate() MatrixDelegateWrapper {
	rv := objc.Call[MatrixDelegateWrapper](m_, objc.Sel("delegate"))
	return rv
}

// The delegate for messages from the field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436404-delegate?language=objc
func (m_ Matrix) SetDelegate(value PMatrixDelegate) {
	po0 := objc.WrapAsProtocol("NSMatrixDelegate", value)
	objc.SetAssociatedObject(m_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), po0)
}

// The delegate for messages from the field editor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436404-delegate?language=objc
func (m_ Matrix) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The background color of the matrix (the space between the cells). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436442-backgroundcolor?language=objc
func (m_ Matrix) BackgroundColor() Color {
	rv := objc.Call[Color](m_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the matrix (the space between the cells). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436442-backgroundcolor?language=objc
func (m_ Matrix) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](m_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean that indicates whether the matrix draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436447-drawsbackground?language=objc
func (m_ Matrix) DrawsBackground() bool {
	rv := objc.Call[bool](m_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean that indicates whether the matrix draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436447-drawsbackground?language=objc
func (m_ Matrix) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setDrawsBackground:"), value)
}

// The size of each cell in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436497-cellsize?language=objc
func (m_ Matrix) CellSize() foundation.Size {
	rv := objc.Call[foundation.Size](m_, objc.Sel("cellSize"))
	return rv
}

// The size of each cell in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436497-cellsize?language=objc
func (m_ Matrix) SetCellSize(value foundation.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setCellSize:"), value)
}

// The cell that will be clicked when the user presses the Space bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436411-keycell?language=objc
func (m_ Matrix) KeyCell() Cell {
	rv := objc.Call[Cell](m_, objc.Sel("keyCell"))
	return rv
}

// The cell that will be clicked when the user presses the Space bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436411-keycell?language=objc
func (m_ Matrix) SetKeyCell(value ICell) {
	objc.Call[objc.Void](m_, objc.Sel("setKeyCell:"), objc.Ptr(value))
}

// The column number of the selected cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436389-selectedcolumn?language=objc
func (m_ Matrix) SelectedColumn() int {
	rv := objc.Call[int](m_, objc.Sel("selectedColumn"))
	return rv
}

// The row number of the selected cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436377-selectedrow?language=objc
func (m_ Matrix) SelectedRow() int {
	rv := objc.Call[int](m_, objc.Sel("selectedRow"))
	return rv
}

// The selection mode of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436390-mode?language=objc
func (m_ Matrix) Mode() MatrixMode {
	rv := objc.Call[MatrixMode](m_, objc.Sel("mode"))
	return rv
}

// The selection mode of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436390-mode?language=objc
func (m_ Matrix) SetMode(value MatrixMode) {
	objc.Call[objc.Void](m_, objc.Sel("setMode:"), value)
}

// A Boolean that indicates whether the cell sizes change when the receiver is resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436375-autosizescells?language=objc
func (m_ Matrix) AutosizesCells() bool {
	rv := objc.Call[bool](m_, objc.Sel("autosizesCells"))
	return rv
}

// A Boolean that indicates whether the cell sizes change when the receiver is resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436375-autosizescells?language=objc
func (m_ Matrix) SetAutosizesCells(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAutosizesCells:"), value)
}

// The subclass of NSCell that the matrix uses when creating new (empty) cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436445-cellclass?language=objc
func (m_ Matrix) CellClass() objc.Class {
	rv := objc.Call[objc.Class](m_, objc.Sel("cellClass"))
	return rv
}

// The subclass of NSCell that the matrix uses when creating new (empty) cells. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436445-cellclass?language=objc
func (m_ Matrix) SetCellClass(value objc.IClass) {
	objc.Call[objc.Void](m_, objc.Sel("setCellClass:"), objc.Ptr(value))
}

// The prototype cell that’s copied whenever the matrix creates a new cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436406-prototype?language=objc
func (m_ Matrix) Prototype() Cell {
	rv := objc.Call[Cell](m_, objc.Sel("prototype"))
	return rv
}

// The prototype cell that’s copied whenever the matrix creates a new cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436406-prototype?language=objc
func (m_ Matrix) SetPrototype(value ICell) {
	objc.Call[objc.Void](m_, objc.Sel("setPrototype:"), objc.Ptr(value))
}

// The number of rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436507-numberofrows?language=objc
func (m_ Matrix) NumberOfRows() int {
	rv := objc.Call[int](m_, objc.Sel("numberOfRows"))
	return rv
}

// The vertical and horizontal spacing between cells in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436398-intercellspacing?language=objc
func (m_ Matrix) IntercellSpacing() foundation.Size {
	rv := objc.Call[foundation.Size](m_, objc.Sel("intercellSpacing"))
	return rv
}

// The vertical and horizontal spacing between cells in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436398-intercellspacing?language=objc
func (m_ Matrix) SetIntercellSpacing(value foundation.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setIntercellSpacing:"), value)
}

// The number of columns in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmatrix/1436461-numberofcolumns?language=objc
func (m_ Matrix) NumberOfColumns() int {
	rv := objc.Call[int](m_, objc.Sel("numberOfColumns"))
	return rv
}
