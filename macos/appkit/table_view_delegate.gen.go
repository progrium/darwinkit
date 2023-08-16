// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods you implement in a table view delegate to customize the behavior of the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate?language=objc
type PTableViewDelegate interface {
	PControlTextEditingDelegate
	// optional
	TableViewSelectionDidChange(notification foundation.Notification)
	HasTableViewSelectionDidChange() bool

	// optional
	TableViewColumnDidResize(notification foundation.Notification)
	HasTableViewColumnDidResize() bool

	// optional
	TableViewSelectionIsChanging(notification foundation.Notification)
	HasTableViewSelectionIsChanging() bool

	// optional
	TableViewDidClickTableColumn(tableView TableView, tableColumn TableColumn)
	HasTableViewDidClickTableColumn() bool

	// optional
	SelectionShouldChangeInTableView(tableView TableView) bool
	HasSelectionShouldChangeInTableView() bool

	// optional
	TableViewColumnDidMove(notification foundation.Notification)
	HasTableViewColumnDidMove() bool
}

// A delegate implementation builder for the [PTableViewDelegate] protocol.
type TableViewDelegate struct {
	ControlTextEditingDelegate
	_TableViewSelectionDidChange      func(notification foundation.Notification)
	_TableViewColumnDidResize         func(notification foundation.Notification)
	_TableViewSelectionIsChanging     func(notification foundation.Notification)
	_TableViewDidClickTableColumn     func(tableView TableView, tableColumn TableColumn)
	_SelectionShouldChangeInTableView func(tableView TableView) bool
	_TableViewColumnDidMove           func(notification foundation.Notification)
}

func (di *TableViewDelegate) HasTableViewSelectionDidChange() bool {
	return di._TableViewSelectionDidChange != nil
}

// Tells the delegate that the table view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1528567-tableviewselectiondidchange?language=objc
func (di *TableViewDelegate) SetTableViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._TableViewSelectionDidChange = f
}

// Tells the delegate that the table view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1528567-tableviewselectiondidchange?language=objc
func (di *TableViewDelegate) TableViewSelectionDidChange(notification foundation.Notification) {
	di._TableViewSelectionDidChange(notification)
}
func (di *TableViewDelegate) HasTableViewColumnDidResize() bool {
	return di._TableViewColumnDidResize != nil
}

// Tells the delegate that a table column was resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1535901-tableviewcolumndidresize?language=objc
func (di *TableViewDelegate) SetTableViewColumnDidResize(f func(notification foundation.Notification)) {
	di._TableViewColumnDidResize = f
}

// Tells the delegate that a table column was resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1535901-tableviewcolumndidresize?language=objc
func (di *TableViewDelegate) TableViewColumnDidResize(notification foundation.Notification) {
	di._TableViewColumnDidResize(notification)
}
func (di *TableViewDelegate) HasTableViewSelectionIsChanging() bool {
	return di._TableViewSelectionIsChanging != nil
}

// Tells the delegate that the table view’s selection is in the process of changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1530812-tableviewselectionischanging?language=objc
func (di *TableViewDelegate) SetTableViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._TableViewSelectionIsChanging = f
}

// Tells the delegate that the table view’s selection is in the process of changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1530812-tableviewselectionischanging?language=objc
func (di *TableViewDelegate) TableViewSelectionIsChanging(notification foundation.Notification) {
	di._TableViewSelectionIsChanging(notification)
}
func (di *TableViewDelegate) HasTableViewDidClickTableColumn() bool {
	return di._TableViewDidClickTableColumn != nil
}

// Tells the delegate that the mouse button was clicked in the specified table column, but the column was not dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533923-tableview?language=objc
func (di *TableViewDelegate) SetTableViewDidClickTableColumn(f func(tableView TableView, tableColumn TableColumn)) {
	di._TableViewDidClickTableColumn = f
}

// Tells the delegate that the mouse button was clicked in the specified table column, but the column was not dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533923-tableview?language=objc
func (di *TableViewDelegate) TableViewDidClickTableColumn(tableView TableView, tableColumn TableColumn) {
	di._TableViewDidClickTableColumn(tableView, tableColumn)
}
func (di *TableViewDelegate) HasSelectionShouldChangeInTableView() bool {
	return di._SelectionShouldChangeInTableView != nil
}

// Asks the delegate if the user is allowed to change the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533949-selectionshouldchangeintableview?language=objc
func (di *TableViewDelegate) SetSelectionShouldChangeInTableView(f func(tableView TableView) bool) {
	di._SelectionShouldChangeInTableView = f
}

// Asks the delegate if the user is allowed to change the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533949-selectionshouldchangeintableview?language=objc
func (di *TableViewDelegate) SelectionShouldChangeInTableView(tableView TableView) bool {
	return di._SelectionShouldChangeInTableView(tableView)
}
func (di *TableViewDelegate) HasTableViewColumnDidMove() bool {
	return di._TableViewColumnDidMove != nil
}

// Tells the delegate that a table column was moved by user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1534237-tableviewcolumndidmove?language=objc
func (di *TableViewDelegate) SetTableViewColumnDidMove(f func(notification foundation.Notification)) {
	di._TableViewColumnDidMove = f
}

// Tells the delegate that a table column was moved by user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1534237-tableviewcolumndidmove?language=objc
func (di *TableViewDelegate) TableViewColumnDidMove(notification foundation.Notification) {
	di._TableViewColumnDidMove(notification)
}

// A concrete type wrapper for the [PTableViewDelegate] protocol.
type TableViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (t_ TableViewDelegateWrapper) HasTableViewSelectionDidChange() bool {
	return t_.RespondsToSelector(objc.Sel("tableViewSelectionDidChange:"))
}

// Tells the delegate that the table view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1528567-tableviewselectiondidchange?language=objc
func (t_ TableViewDelegateWrapper) TableViewSelectionDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("tableViewSelectionDidChange:"), objc.Ptr(notification))
}

func (t_ TableViewDelegateWrapper) HasTableViewColumnDidResize() bool {
	return t_.RespondsToSelector(objc.Sel("tableViewColumnDidResize:"))
}

// Tells the delegate that a table column was resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1535901-tableviewcolumndidresize?language=objc
func (t_ TableViewDelegateWrapper) TableViewColumnDidResize(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("tableViewColumnDidResize:"), objc.Ptr(notification))
}

func (t_ TableViewDelegateWrapper) HasTableViewSelectionIsChanging() bool {
	return t_.RespondsToSelector(objc.Sel("tableViewSelectionIsChanging:"))
}

// Tells the delegate that the table view’s selection is in the process of changing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1530812-tableviewselectionischanging?language=objc
func (t_ TableViewDelegateWrapper) TableViewSelectionIsChanging(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("tableViewSelectionIsChanging:"), objc.Ptr(notification))
}

func (t_ TableViewDelegateWrapper) HasTableViewDidClickTableColumn() bool {
	return t_.RespondsToSelector(objc.Sel("tableView:didClickTableColumn:"))
}

// Tells the delegate that the mouse button was clicked in the specified table column, but the column was not dragged. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533923-tableview?language=objc
func (t_ TableViewDelegateWrapper) TableViewDidClickTableColumn(tableView ITableView, tableColumn ITableColumn) {
	objc.Call[objc.Void](t_, objc.Sel("tableView:didClickTableColumn:"), objc.Ptr(tableView), objc.Ptr(tableColumn))
}

func (t_ TableViewDelegateWrapper) HasSelectionShouldChangeInTableView() bool {
	return t_.RespondsToSelector(objc.Sel("selectionShouldChangeInTableView:"))
}

// Asks the delegate if the user is allowed to change the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1533949-selectionshouldchangeintableview?language=objc
func (t_ TableViewDelegateWrapper) SelectionShouldChangeInTableView(tableView ITableView) bool {
	rv := objc.Call[bool](t_, objc.Sel("selectionShouldChangeInTableView:"), objc.Ptr(tableView))
	return rv
}

func (t_ TableViewDelegateWrapper) HasTableViewColumnDidMove() bool {
	return t_.RespondsToSelector(objc.Sel("tableViewColumnDidMove:"))
}

// Tells the delegate that a table column was moved by user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdelegate/1534237-tableviewcolumndidmove?language=objc
func (t_ TableViewDelegateWrapper) TableViewColumnDidMove(notification foundation.INotification) {
	objc.Call[objc.Void](t_, objc.Sel("tableViewColumnDidMove:"), objc.Ptr(notification))
}
