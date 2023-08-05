// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IOutlineViewDelegate interface {
	IControlTextEditingDelegate
	ImplementsOutlineViewShouldExpandItem() bool
	// optional
	OutlineViewShouldExpandItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewShouldCollapseItem() bool
	// optional
	OutlineViewShouldCollapseItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewTypeSelectStringForTableColumnItem() bool
	// optional
	OutlineViewTypeSelectStringForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	ImplementsOutlineViewNextTypeSelectMatchFromItemToItemForString() bool
	// optional
	OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject
	ImplementsOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool
	// optional
	OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool
	ImplementsOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool
	// optional
	OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string
	ImplementsOutlineViewShouldSelectTableColumn() bool
	// optional
	OutlineViewShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool
	ImplementsOutlineViewShouldSelectItem() bool
	// optional
	OutlineViewShouldSelectItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewSelectionIndexesForProposedSelection() bool
	// optional
	OutlineViewSelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	ImplementsSelectionShouldChangeInOutlineView() bool
	// optional
	SelectionShouldChangeInOutlineView(outlineView OutlineView) bool
	ImplementsOutlineViewSelectionIsChanging() bool
	// optional
	OutlineViewSelectionIsChanging(notification foundation.Notification)
	ImplementsOutlineViewSelectionDidChange() bool
	// optional
	OutlineViewSelectionDidChange(notification foundation.Notification)
	ImplementsOutlineViewWillDisplayCellForTableColumnItem() bool
	// optional
	OutlineViewWillDisplayCellForTableColumnItem(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineViewWillDisplayOutlineCellForTableColumnItem() bool
	// optional
	OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	ImplementsOutlineViewDataCellForTableColumnItem() bool
	// optional
	OutlineViewDataCellForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell
	ImplementsOutlineViewShouldShowOutlineCellForItem() bool
	// optional
	OutlineViewShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewShouldShowCellExpansionForTableColumnItem() bool
	// optional
	OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineViewShouldReorderColumnToColumn() bool
	// optional
	OutlineViewShouldReorderColumnToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	ImplementsOutlineViewColumnDidMove() bool
	// optional
	OutlineViewColumnDidMove(notification foundation.Notification)
	ImplementsOutlineViewColumnDidResize() bool
	// optional
	OutlineViewColumnDidResize(notification foundation.Notification)
	ImplementsOutlineViewItemWillExpand() bool
	// optional
	OutlineViewItemWillExpand(notification foundation.Notification)
	ImplementsOutlineViewItemDidExpand() bool
	// optional
	OutlineViewItemDidExpand(notification foundation.Notification)
	ImplementsOutlineViewItemWillCollapse() bool
	// optional
	OutlineViewItemWillCollapse(notification foundation.Notification)
	ImplementsOutlineViewItemDidCollapse() bool
	// optional
	OutlineViewItemDidCollapse(notification foundation.Notification)
	ImplementsOutlineViewShouldEditTableColumnItem() bool
	// optional
	OutlineViewShouldEditTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineViewMouseDownInHeaderOfTableColumn() bool
	// optional
	OutlineViewMouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineViewDidClickTableColumn() bool
	// optional
	OutlineViewDidClickTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineViewDidDragTableColumn() bool
	// optional
	OutlineViewDidDragTableColumn(outlineView OutlineView, tableColumn TableColumn)
	ImplementsOutlineViewHeightOfRowByItem() bool
	// optional
	OutlineViewHeightOfRowByItem(outlineView OutlineView, item objc.Object) float64
	ImplementsOutlineViewSizeToFitWidthOfColumn() bool
	// optional
	OutlineViewSizeToFitWidthOfColumn(outlineView OutlineView, column int) float64
	ImplementsOutlineViewTintConfigurationForItem() bool
	// optional
	OutlineViewTintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration
	ImplementsOutlineViewShouldTrackCellForTableColumnItem() bool
	// optional
	OutlineViewShouldTrackCellForTableColumnItem(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	ImplementsOutlineViewIsGroupItem() bool
	// optional
	OutlineViewIsGroupItem(outlineView OutlineView, item objc.Object) bool
	ImplementsOutlineViewDidAddRowViewForRow() bool
	// optional
	OutlineViewDidAddRowViewForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineViewDidRemoveRowViewForRow() bool
	// optional
	OutlineViewDidRemoveRowViewForRow(outlineView OutlineView, rowView TableRowView, row int)
	ImplementsOutlineViewRowViewForItem() bool
	// optional
	OutlineViewRowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView
	ImplementsOutlineViewViewForTableColumnItem() bool
	// optional
	OutlineViewViewForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView
}

type OutlineViewDelegate struct {
	ControlTextEditingDelegate
	_OutlineViewShouldExpandItem                                func(outlineView OutlineView, item objc.Object) bool
	_OutlineViewShouldCollapseItem                              func(outlineView OutlineView, item objc.Object) bool
	_OutlineViewTypeSelectStringForTableColumnItem              func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string
	_OutlineViewNextTypeSelectMatchFromItemToItemForString      func(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject
	_OutlineViewShouldTypeSelectForEventWithCurrentSearchString func(outlineView OutlineView, event Event, searchString string) bool
	_OutlineViewToolTipForCellRectTableColumnItemMouseLocation  func(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string
	_OutlineViewShouldSelectTableColumn                         func(outlineView OutlineView, tableColumn TableColumn) bool
	_OutlineViewShouldSelectItem                                func(outlineView OutlineView, item objc.Object) bool
	_OutlineViewSelectionIndexesForProposedSelection            func(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet
	_SelectionShouldChangeInOutlineView                         func(outlineView OutlineView) bool
	_OutlineViewSelectionIsChanging                             func(notification foundation.Notification)
	_OutlineViewSelectionDidChange                              func(notification foundation.Notification)
	_OutlineViewWillDisplayCellForTableColumnItem               func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	_OutlineViewWillDisplayOutlineCellForTableColumnItem        func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)
	_OutlineViewDataCellForTableColumnItem                      func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell
	_OutlineViewShouldShowOutlineCellForItem                    func(outlineView OutlineView, item objc.Object) bool
	_OutlineViewShouldShowCellExpansionForTableColumnItem       func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	_OutlineViewShouldReorderColumnToColumn                     func(outlineView OutlineView, columnIndex int, newColumnIndex int) bool
	_OutlineViewColumnDidMove                                   func(notification foundation.Notification)
	_OutlineViewColumnDidResize                                 func(notification foundation.Notification)
	_OutlineViewItemWillExpand                                  func(notification foundation.Notification)
	_OutlineViewItemDidExpand                                   func(notification foundation.Notification)
	_OutlineViewItemWillCollapse                                func(notification foundation.Notification)
	_OutlineViewItemDidCollapse                                 func(notification foundation.Notification)
	_OutlineViewShouldEditTableColumnItem                       func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool
	_OutlineViewMouseDownInHeaderOfTableColumn                  func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineViewDidClickTableColumn                             func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineViewDidDragTableColumn                              func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineViewHeightOfRowByItem                               func(outlineView OutlineView, item objc.Object) float64
	_OutlineViewSizeToFitWidthOfColumn                          func(outlineView OutlineView, column int) float64
	_OutlineViewTintConfigurationForItem                        func(outlineView OutlineView, item objc.Object) ITintConfiguration
	_OutlineViewShouldTrackCellForTableColumnItem               func(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool
	_OutlineViewIsGroupItem                                     func(outlineView OutlineView, item objc.Object) bool
	_OutlineViewDidAddRowViewForRow                             func(outlineView OutlineView, rowView TableRowView, row int)
	_OutlineViewDidRemoveRowViewForRow                          func(outlineView OutlineView, rowView TableRowView, row int)
	_OutlineViewRowViewForItem                                  func(outlineView OutlineView, item objc.Object) ITableRowView
	_OutlineViewViewForTableColumnItem                          func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView
}

func (di *OutlineViewDelegate) ImplementsOutlineViewShouldExpandItem() bool {
	return di._OutlineViewShouldExpandItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldExpandItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineViewShouldExpandItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldExpandItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineViewShouldExpandItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldCollapseItem() bool {
	return di._OutlineViewShouldCollapseItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldCollapseItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineViewShouldCollapseItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldCollapseItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineViewShouldCollapseItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewTypeSelectStringForTableColumnItem() bool {
	return di._OutlineViewTypeSelectStringForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewTypeSelectStringForTableColumnItem(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string) {
	di._OutlineViewTypeSelectStringForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewTypeSelectStringForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) string {
	return di._OutlineViewTypeSelectStringForTableColumnItem(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewNextTypeSelectMatchFromItemToItemForString() bool {
	return di._OutlineViewNextTypeSelectMatchFromItemToItemForString != nil
}

func (di *OutlineViewDelegate) SetOutlineViewNextTypeSelectMatchFromItemToItemForString(f func(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject) {
	di._OutlineViewNextTypeSelectMatchFromItemToItemForString = f
}

func (di *OutlineViewDelegate) OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView OutlineView, startItem objc.Object, endItem objc.Object, searchString string) objc.IObject {
	return di._OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView, startItem, endItem, searchString)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return di._OutlineViewShouldTypeSelectForEventWithCurrentSearchString != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldTypeSelectForEventWithCurrentSearchString(f func(outlineView OutlineView, event Event, searchString string) bool) {
	di._OutlineViewShouldTypeSelectForEventWithCurrentSearchString = f
}

func (di *OutlineViewDelegate) OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView OutlineView, event Event, searchString string) bool {
	return di._OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView, event, searchString)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool {
	return di._OutlineViewToolTipForCellRectTableColumnItemMouseLocation != nil
}

func (di *OutlineViewDelegate) SetOutlineViewToolTipForCellRectTableColumnItemMouseLocation(f func(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string) {
	di._OutlineViewToolTipForCellRectTableColumnItemMouseLocation = f
}

func (di *OutlineViewDelegate) OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView OutlineView, cell Cell, rect *foundation.Rect, tableColumn TableColumn, item objc.Object, mouseLocation foundation.Point) string {
	return di._OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView, cell, rect, tableColumn, item, mouseLocation)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldSelectTableColumn() bool {
	return di._OutlineViewShouldSelectTableColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldSelectTableColumn(f func(outlineView OutlineView, tableColumn TableColumn) bool) {
	di._OutlineViewShouldSelectTableColumn = f
}

func (di *OutlineViewDelegate) OutlineViewShouldSelectTableColumn(outlineView OutlineView, tableColumn TableColumn) bool {
	return di._OutlineViewShouldSelectTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldSelectItem() bool {
	return di._OutlineViewShouldSelectItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldSelectItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineViewShouldSelectItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldSelectItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineViewShouldSelectItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewSelectionIndexesForProposedSelection() bool {
	return di._OutlineViewSelectionIndexesForProposedSelection != nil
}

func (di *OutlineViewDelegate) SetOutlineViewSelectionIndexesForProposedSelection(f func(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet) {
	di._OutlineViewSelectionIndexesForProposedSelection = f
}

func (di *OutlineViewDelegate) OutlineViewSelectionIndexesForProposedSelection(outlineView OutlineView, proposedSelectionIndexes foundation.IndexSet) foundation.IIndexSet {
	return di._OutlineViewSelectionIndexesForProposedSelection(outlineView, proposedSelectionIndexes)
}
func (di *OutlineViewDelegate) ImplementsSelectionShouldChangeInOutlineView() bool {
	return di._SelectionShouldChangeInOutlineView != nil
}

func (di *OutlineViewDelegate) SetSelectionShouldChangeInOutlineView(f func(outlineView OutlineView) bool) {
	di._SelectionShouldChangeInOutlineView = f
}

func (di *OutlineViewDelegate) SelectionShouldChangeInOutlineView(outlineView OutlineView) bool {
	return di._SelectionShouldChangeInOutlineView(outlineView)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewSelectionIsChanging() bool {
	return di._OutlineViewSelectionIsChanging != nil
}

func (di *OutlineViewDelegate) SetOutlineViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionIsChanging = f
}

func (di *OutlineViewDelegate) OutlineViewSelectionIsChanging(notification foundation.Notification) {
	di._OutlineViewSelectionIsChanging(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewSelectionDidChange() bool {
	return di._OutlineViewSelectionDidChange != nil
}

func (di *OutlineViewDelegate) SetOutlineViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionDidChange = f
}

func (di *OutlineViewDelegate) OutlineViewSelectionDidChange(notification foundation.Notification) {
	di._OutlineViewSelectionDidChange(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewWillDisplayCellForTableColumnItem() bool {
	return di._OutlineViewWillDisplayCellForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewWillDisplayCellForTableColumnItem(f func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)) {
	di._OutlineViewWillDisplayCellForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewWillDisplayCellForTableColumnItem(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	di._OutlineViewWillDisplayCellForTableColumnItem(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewWillDisplayOutlineCellForTableColumnItem() bool {
	return di._OutlineViewWillDisplayOutlineCellForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewWillDisplayOutlineCellForTableColumnItem(f func(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object)) {
	di._OutlineViewWillDisplayOutlineCellForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView OutlineView, cell objc.Object, tableColumn TableColumn, item objc.Object) {
	di._OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewDataCellForTableColumnItem() bool {
	return di._OutlineViewDataCellForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewDataCellForTableColumnItem(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell) {
	di._OutlineViewDataCellForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewDataCellForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) ICell {
	return di._OutlineViewDataCellForTableColumnItem(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldShowOutlineCellForItem() bool {
	return di._OutlineViewShouldShowOutlineCellForItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldShowOutlineCellForItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineViewShouldShowOutlineCellForItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldShowOutlineCellForItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineViewShouldShowOutlineCellForItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldShowCellExpansionForTableColumnItem() bool {
	return di._OutlineViewShouldShowCellExpansionForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldShowCellExpansionForTableColumnItem(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineViewShouldShowCellExpansionForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldReorderColumnToColumn() bool {
	return di._OutlineViewShouldReorderColumnToColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldReorderColumnToColumn(f func(outlineView OutlineView, columnIndex int, newColumnIndex int) bool) {
	di._OutlineViewShouldReorderColumnToColumn = f
}

func (di *OutlineViewDelegate) OutlineViewShouldReorderColumnToColumn(outlineView OutlineView, columnIndex int, newColumnIndex int) bool {
	return di._OutlineViewShouldReorderColumnToColumn(outlineView, columnIndex, newColumnIndex)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewColumnDidMove() bool {
	return di._OutlineViewColumnDidMove != nil
}

func (di *OutlineViewDelegate) SetOutlineViewColumnDidMove(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidMove = f
}

func (di *OutlineViewDelegate) OutlineViewColumnDidMove(notification foundation.Notification) {
	di._OutlineViewColumnDidMove(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewColumnDidResize() bool {
	return di._OutlineViewColumnDidResize != nil
}

func (di *OutlineViewDelegate) SetOutlineViewColumnDidResize(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidResize = f
}

func (di *OutlineViewDelegate) OutlineViewColumnDidResize(notification foundation.Notification) {
	di._OutlineViewColumnDidResize(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewItemWillExpand() bool {
	return di._OutlineViewItemWillExpand != nil
}

func (di *OutlineViewDelegate) SetOutlineViewItemWillExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillExpand = f
}

func (di *OutlineViewDelegate) OutlineViewItemWillExpand(notification foundation.Notification) {
	di._OutlineViewItemWillExpand(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewItemDidExpand() bool {
	return di._OutlineViewItemDidExpand != nil
}

func (di *OutlineViewDelegate) SetOutlineViewItemDidExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidExpand = f
}

func (di *OutlineViewDelegate) OutlineViewItemDidExpand(notification foundation.Notification) {
	di._OutlineViewItemDidExpand(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewItemWillCollapse() bool {
	return di._OutlineViewItemWillCollapse != nil
}

func (di *OutlineViewDelegate) SetOutlineViewItemWillCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillCollapse = f
}

func (di *OutlineViewDelegate) OutlineViewItemWillCollapse(notification foundation.Notification) {
	di._OutlineViewItemWillCollapse(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewItemDidCollapse() bool {
	return di._OutlineViewItemDidCollapse != nil
}

func (di *OutlineViewDelegate) SetOutlineViewItemDidCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidCollapse = f
}

func (di *OutlineViewDelegate) OutlineViewItemDidCollapse(notification foundation.Notification) {
	di._OutlineViewItemDidCollapse(notification)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldEditTableColumnItem() bool {
	return di._OutlineViewShouldEditTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldEditTableColumnItem(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineViewShouldEditTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldEditTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineViewShouldEditTableColumnItem(outlineView, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewMouseDownInHeaderOfTableColumn() bool {
	return di._OutlineViewMouseDownInHeaderOfTableColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewMouseDownInHeaderOfTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineViewMouseDownInHeaderOfTableColumn = f
}

func (di *OutlineViewDelegate) OutlineViewMouseDownInHeaderOfTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineViewMouseDownInHeaderOfTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewDidClickTableColumn() bool {
	return di._OutlineViewDidClickTableColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewDidClickTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineViewDidClickTableColumn = f
}

func (di *OutlineViewDelegate) OutlineViewDidClickTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineViewDidClickTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewDidDragTableColumn() bool {
	return di._OutlineViewDidDragTableColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewDidDragTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineViewDidDragTableColumn = f
}

func (di *OutlineViewDelegate) OutlineViewDidDragTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineViewDidDragTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewHeightOfRowByItem() bool {
	return di._OutlineViewHeightOfRowByItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewHeightOfRowByItem(f func(outlineView OutlineView, item objc.Object) float64) {
	di._OutlineViewHeightOfRowByItem = f
}

func (di *OutlineViewDelegate) OutlineViewHeightOfRowByItem(outlineView OutlineView, item objc.Object) float64 {
	return di._OutlineViewHeightOfRowByItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewSizeToFitWidthOfColumn() bool {
	return di._OutlineViewSizeToFitWidthOfColumn != nil
}

func (di *OutlineViewDelegate) SetOutlineViewSizeToFitWidthOfColumn(f func(outlineView OutlineView, column int) float64) {
	di._OutlineViewSizeToFitWidthOfColumn = f
}

func (di *OutlineViewDelegate) OutlineViewSizeToFitWidthOfColumn(outlineView OutlineView, column int) float64 {
	return di._OutlineViewSizeToFitWidthOfColumn(outlineView, column)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewTintConfigurationForItem() bool {
	return di._OutlineViewTintConfigurationForItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewTintConfigurationForItem(f func(outlineView OutlineView, item objc.Object) ITintConfiguration) {
	di._OutlineViewTintConfigurationForItem = f
}

func (di *OutlineViewDelegate) OutlineViewTintConfigurationForItem(outlineView OutlineView, item objc.Object) ITintConfiguration {
	return di._OutlineViewTintConfigurationForItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewShouldTrackCellForTableColumnItem() bool {
	return di._OutlineViewShouldTrackCellForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewShouldTrackCellForTableColumnItem(f func(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool) {
	di._OutlineViewShouldTrackCellForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewShouldTrackCellForTableColumnItem(outlineView OutlineView, cell Cell, tableColumn TableColumn, item objc.Object) bool {
	return di._OutlineViewShouldTrackCellForTableColumnItem(outlineView, cell, tableColumn, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewIsGroupItem() bool {
	return di._OutlineViewIsGroupItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewIsGroupItem(f func(outlineView OutlineView, item objc.Object) bool) {
	di._OutlineViewIsGroupItem = f
}

func (di *OutlineViewDelegate) OutlineViewIsGroupItem(outlineView OutlineView, item objc.Object) bool {
	return di._OutlineViewIsGroupItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewDidAddRowViewForRow() bool {
	return di._OutlineViewDidAddRowViewForRow != nil
}

func (di *OutlineViewDelegate) SetOutlineViewDidAddRowViewForRow(f func(outlineView OutlineView, rowView TableRowView, row int)) {
	di._OutlineViewDidAddRowViewForRow = f
}

func (di *OutlineViewDelegate) OutlineViewDidAddRowViewForRow(outlineView OutlineView, rowView TableRowView, row int) {
	di._OutlineViewDidAddRowViewForRow(outlineView, rowView, row)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewDidRemoveRowViewForRow() bool {
	return di._OutlineViewDidRemoveRowViewForRow != nil
}

func (di *OutlineViewDelegate) SetOutlineViewDidRemoveRowViewForRow(f func(outlineView OutlineView, rowView TableRowView, row int)) {
	di._OutlineViewDidRemoveRowViewForRow = f
}

func (di *OutlineViewDelegate) OutlineViewDidRemoveRowViewForRow(outlineView OutlineView, rowView TableRowView, row int) {
	di._OutlineViewDidRemoveRowViewForRow(outlineView, rowView, row)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewRowViewForItem() bool {
	return di._OutlineViewRowViewForItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewRowViewForItem(f func(outlineView OutlineView, item objc.Object) ITableRowView) {
	di._OutlineViewRowViewForItem = f
}

func (di *OutlineViewDelegate) OutlineViewRowViewForItem(outlineView OutlineView, item objc.Object) ITableRowView {
	return di._OutlineViewRowViewForItem(outlineView, item)
}
func (di *OutlineViewDelegate) ImplementsOutlineViewViewForTableColumnItem() bool {
	return di._OutlineViewViewForTableColumnItem != nil
}

func (di *OutlineViewDelegate) SetOutlineViewViewForTableColumnItem(f func(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView) {
	di._OutlineViewViewForTableColumnItem = f
}

func (di *OutlineViewDelegate) OutlineViewViewForTableColumnItem(outlineView OutlineView, tableColumn TableColumn, item objc.Object) IView {
	return di._OutlineViewViewForTableColumnItem(outlineView, tableColumn, item)
}

type OutlineViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldExpandItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldExpandItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldExpandItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldExpandItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldCollapseItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldCollapseItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldCollapseItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldCollapseItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewTypeSelectStringForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:typeSelectStringForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewTypeSelectStringForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("outlineView:typeSelectStringForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewNextTypeSelectMatchFromItemToItemForString() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewNextTypeSelectMatchFromItemToItemForString(outlineView IOutlineView, startItem objc.IObject, endItem objc.IObject, searchString string) objc.Object {
	rv := objc.CallMethod[objc.Object](o_, objc.GetSelector("outlineView:nextTypeSelectMatchFromItem:toItem:forString:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(startItem), objc.ExtractPtr(endItem), searchString)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldTypeSelectForEventWithCurrentSearchString() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldTypeSelectForEventWithCurrentSearchString(outlineView IOutlineView, event IEvent, searchString string) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldTypeSelectForEvent:withCurrentSearchString:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(event), searchString)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewToolTipForCellRectTableColumnItemMouseLocation() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewToolTipForCellRectTableColumnItemMouseLocation(outlineView IOutlineView, cell ICell, rect *foundation.Rect, tableColumn ITableColumn, item objc.IObject, mouseLocation foundation.Point) string {
	rv := objc.CallMethod[string](o_, objc.GetSelector("outlineView:toolTipForCell:rect:tableColumn:item:mouseLocation:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), rect, objc.ExtractPtr(tableColumn), objc.ExtractPtr(item), mouseLocation)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldSelectTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldSelectTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldSelectTableColumn(outlineView IOutlineView, tableColumn ITableColumn) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldSelectTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldSelectItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldSelectItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldSelectItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldSelectItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewSelectionIndexesForProposedSelection() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:selectionIndexesForProposedSelection:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionIndexesForProposedSelection(outlineView IOutlineView, proposedSelectionIndexes foundation.IIndexSet) foundation.IndexSet {
	rv := objc.CallMethod[foundation.IndexSet](o_, objc.GetSelector("outlineView:selectionIndexesForProposedSelection:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(proposedSelectionIndexes))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsSelectionShouldChangeInOutlineView() bool {
	return o_.RespondsToSelector(objc.GetSelector("selectionShouldChangeInOutlineView:"))
}

func (o_ OutlineViewDelegateWrapper) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("selectionShouldChangeInOutlineView:"), objc.ExtractPtr(outlineView))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewSelectionIsChanging() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewSelectionIsChanging:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionIsChanging(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewSelectionIsChanging:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewSelectionDidChange:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewSelectionDidChange:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewWillDisplayCellForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:willDisplayCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewWillDisplayCellForTableColumnItem(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:willDisplayCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewWillDisplayOutlineCellForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:willDisplayOutlineCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewWillDisplayOutlineCellForTableColumnItem(outlineView IOutlineView, cell objc.IObject, tableColumn ITableColumn, item objc.IObject) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:willDisplayOutlineCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewDataCellForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:dataCellForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewDataCellForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) Cell {
	rv := objc.CallMethod[Cell](o_, objc.GetSelector("outlineView:dataCellForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldShowOutlineCellForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldShowOutlineCellForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldShowOutlineCellForItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldShowOutlineCellForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldShowCellExpansionForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldShowCellExpansionForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldShowCellExpansionForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldShowCellExpansionForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldReorderColumnToColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldReorderColumn:toColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldReorderColumnToColumn(outlineView IOutlineView, columnIndex int, newColumnIndex int) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldReorderColumn:toColumn:"), objc.ExtractPtr(outlineView), columnIndex, newColumnIndex)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewColumnDidMove() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewColumnDidMove:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewColumnDidMove:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewColumnDidResize() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewColumnDidResize:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewColumnDidResize:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewItemWillExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemWillExpand:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillExpand(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemWillExpand:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewItemDidExpand() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemDidExpand:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidExpand(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemDidExpand:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewItemWillCollapse() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemWillCollapse:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillCollapse(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemWillCollapse:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewItemDidCollapse() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineViewItemDidCollapse:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidCollapse(notification foundation.INotification) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineViewItemDidCollapse:"), objc.ExtractPtr(notification))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldEditTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldEditTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldEditTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldEditTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewMouseDownInHeaderOfTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:mouseDownInHeaderOfTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewMouseDownInHeaderOfTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:mouseDownInHeaderOfTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewDidClickTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didClickTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewDidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didClickTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewDidDragTableColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didDragTableColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewDidDragTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didDragTableColumn:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewHeightOfRowByItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:heightOfRowByItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewHeightOfRowByItem(outlineView IOutlineView, item objc.IObject) float64 {
	rv := objc.CallMethod[float64](o_, objc.GetSelector("outlineView:heightOfRowByItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewSizeToFitWidthOfColumn() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:sizeToFitWidthOfColumn:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewSizeToFitWidthOfColumn(outlineView IOutlineView, column int) float64 {
	rv := objc.CallMethod[float64](o_, objc.GetSelector("outlineView:sizeToFitWidthOfColumn:"), objc.ExtractPtr(outlineView), column)
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewTintConfigurationForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:tintConfigurationForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewTintConfigurationForItem(outlineView IOutlineView, item objc.IObject) TintConfiguration {
	rv := objc.CallMethod[TintConfiguration](o_, objc.GetSelector("outlineView:tintConfigurationForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewShouldTrackCellForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:shouldTrackCell:forTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewShouldTrackCellForTableColumnItem(outlineView IOutlineView, cell ICell, tableColumn ITableColumn, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:shouldTrackCell:forTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(cell), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewIsGroupItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:isGroupItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewIsGroupItem(outlineView IOutlineView, item objc.IObject) bool {
	rv := objc.CallMethod[bool](o_, objc.GetSelector("outlineView:isGroupItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewDidAddRowViewForRow() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didAddRowView:forRow:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewDidAddRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didAddRowView:forRow:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(rowView), row)
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewDidRemoveRowViewForRow() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:didRemoveRowView:forRow:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewDidRemoveRowViewForRow(outlineView IOutlineView, rowView ITableRowView, row int) {
	objc.CallMethod[objc.Void](o_, objc.GetSelector("outlineView:didRemoveRowView:forRow:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(rowView), row)
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewRowViewForItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:rowViewForItem:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewRowViewForItem(outlineView IOutlineView, item objc.IObject) TableRowView {
	rv := objc.CallMethod[TableRowView](o_, objc.GetSelector("outlineView:rowViewForItem:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(item))
	return rv
}

func (o_ OutlineViewDelegateWrapper) ImplementsOutlineViewViewForTableColumnItem() bool {
	return o_.RespondsToSelector(objc.GetSelector("outlineView:viewForTableColumn:item:"))
}

func (o_ OutlineViewDelegateWrapper) OutlineViewViewForTableColumnItem(outlineView IOutlineView, tableColumn ITableColumn, item objc.IObject) View {
	rv := objc.CallMethod[View](o_, objc.GetSelector("outlineView:viewForTableColumn:item:"), objc.ExtractPtr(outlineView), objc.ExtractPtr(tableColumn), objc.ExtractPtr(item))
	return rv
}
