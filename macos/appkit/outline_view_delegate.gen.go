// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSOutlineView objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate?language=objc
type POutlineViewDelegate interface {
	PControlTextEditingDelegate
	// optional
	OutlineViewItemDidCollapse(notification foundation.Notification)
	HasOutlineViewItemDidCollapse() bool

	// optional
	OutlineViewItemWillExpand(notification foundation.Notification)
	HasOutlineViewItemWillExpand() bool

	// optional
	OutlineViewColumnDidMove(notification foundation.Notification)
	HasOutlineViewColumnDidMove() bool

	// optional
	OutlineViewItemWillCollapse(notification foundation.Notification)
	HasOutlineViewItemWillCollapse() bool

	// optional
	OutlineViewColumnDidResize(notification foundation.Notification)
	HasOutlineViewColumnDidResize() bool

	// optional
	OutlineViewDidClickTableColumn(outlineView OutlineView, tableColumn TableColumn)
	HasOutlineViewDidClickTableColumn() bool

	// optional
	OutlineViewSelectionDidChange(notification foundation.Notification)
	HasOutlineViewSelectionDidChange() bool

	// optional
	OutlineViewItemDidExpand(notification foundation.Notification)
	HasOutlineViewItemDidExpand() bool

	// optional
	OutlineViewSelectionIsChanging(notification foundation.Notification)
	HasOutlineViewSelectionIsChanging() bool

	// optional
	SelectionShouldChangeInOutlineView(outlineView OutlineView) bool
	HasSelectionShouldChangeInOutlineView() bool
}

// A delegate implementation builder for the [POutlineViewDelegate] protocol.
type OutlineViewDelegate struct {
	ControlTextEditingDelegate
	_OutlineViewItemDidCollapse         func(notification foundation.Notification)
	_OutlineViewItemWillExpand          func(notification foundation.Notification)
	_OutlineViewColumnDidMove           func(notification foundation.Notification)
	_OutlineViewItemWillCollapse        func(notification foundation.Notification)
	_OutlineViewColumnDidResize         func(notification foundation.Notification)
	_OutlineViewDidClickTableColumn     func(outlineView OutlineView, tableColumn TableColumn)
	_OutlineViewSelectionDidChange      func(notification foundation.Notification)
	_OutlineViewItemDidExpand           func(notification foundation.Notification)
	_OutlineViewSelectionIsChanging     func(notification foundation.Notification)
	_SelectionShouldChangeInOutlineView func(outlineView OutlineView) bool
}

func (di *OutlineViewDelegate) HasOutlineViewItemDidCollapse() bool {
	return di._OutlineViewItemDidCollapse != nil
}

// Invoked when the did collapse notification is posted—that is, whenever the user collapses an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535557-outlineviewitemdidcollapse?language=objc
func (di *OutlineViewDelegate) SetOutlineViewItemDidCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidCollapse = f
}

// Invoked when the did collapse notification is posted—that is, whenever the user collapses an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535557-outlineviewitemdidcollapse?language=objc
func (di *OutlineViewDelegate) OutlineViewItemDidCollapse(notification foundation.Notification) {
	di._OutlineViewItemDidCollapse(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewItemWillExpand() bool {
	return di._OutlineViewItemWillExpand != nil
}

// Invoked when notification is posted—that is, whenever the user is about to expand an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535847-outlineviewitemwillexpand?language=objc
func (di *OutlineViewDelegate) SetOutlineViewItemWillExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillExpand = f
}

// Invoked when notification is posted—that is, whenever the user is about to expand an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535847-outlineviewitemwillexpand?language=objc
func (di *OutlineViewDelegate) OutlineViewItemWillExpand(notification foundation.Notification) {
	di._OutlineViewItemWillExpand(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewColumnDidMove() bool {
	return di._OutlineViewColumnDidMove != nil
}

// Invoked whenever the user moves a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1525297-outlineviewcolumndidmove?language=objc
func (di *OutlineViewDelegate) SetOutlineViewColumnDidMove(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidMove = f
}

// Invoked whenever the user moves a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1525297-outlineviewcolumndidmove?language=objc
func (di *OutlineViewDelegate) OutlineViewColumnDidMove(notification foundation.Notification) {
	di._OutlineViewColumnDidMove(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewItemWillCollapse() bool {
	return di._OutlineViewItemWillCollapse != nil
}

// Invoked when notification is posted—that is, whenever the user is about to collapse an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526896-outlineviewitemwillcollapse?language=objc
func (di *OutlineViewDelegate) SetOutlineViewItemWillCollapse(f func(notification foundation.Notification)) {
	di._OutlineViewItemWillCollapse = f
}

// Invoked when notification is posted—that is, whenever the user is about to collapse an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526896-outlineviewitemwillcollapse?language=objc
func (di *OutlineViewDelegate) OutlineViewItemWillCollapse(notification foundation.Notification) {
	di._OutlineViewItemWillCollapse(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewColumnDidResize() bool {
	return di._OutlineViewColumnDidResize != nil
}

// Invoked whenever the user resizes a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1533372-outlineviewcolumndidresize?language=objc
func (di *OutlineViewDelegate) SetOutlineViewColumnDidResize(f func(notification foundation.Notification)) {
	di._OutlineViewColumnDidResize = f
}

// Invoked whenever the user resizes a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1533372-outlineviewcolumndidresize?language=objc
func (di *OutlineViewDelegate) OutlineViewColumnDidResize(notification foundation.Notification) {
	di._OutlineViewColumnDidResize(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewDidClickTableColumn() bool {
	return di._OutlineViewDidClickTableColumn != nil
}

// Sent at the time the mouse button subsequently goes up in outlineView and tableColumn has been “clicked” without having been dragged anywhere. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1534040-outlineview?language=objc
func (di *OutlineViewDelegate) SetOutlineViewDidClickTableColumn(f func(outlineView OutlineView, tableColumn TableColumn)) {
	di._OutlineViewDidClickTableColumn = f
}

// Sent at the time the mouse button subsequently goes up in outlineView and tableColumn has been “clicked” without having been dragged anywhere. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1534040-outlineview?language=objc
func (di *OutlineViewDelegate) OutlineViewDidClickTableColumn(outlineView OutlineView, tableColumn TableColumn) {
	di._OutlineViewDidClickTableColumn(outlineView, tableColumn)
}
func (di *OutlineViewDelegate) HasOutlineViewSelectionDidChange() bool {
	return di._OutlineViewSelectionDidChange != nil
}

// Invoked when the selection did change notification is posted—that is, immediately after the outline view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526913-outlineviewselectiondidchange?language=objc
func (di *OutlineViewDelegate) SetOutlineViewSelectionDidChange(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionDidChange = f
}

// Invoked when the selection did change notification is posted—that is, immediately after the outline view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526913-outlineviewselectiondidchange?language=objc
func (di *OutlineViewDelegate) OutlineViewSelectionDidChange(notification foundation.Notification) {
	di._OutlineViewSelectionDidChange(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewItemDidExpand() bool {
	return di._OutlineViewItemDidExpand != nil
}

// Invoked when notification is posted—that is, whenever the user expands an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530869-outlineviewitemdidexpand?language=objc
func (di *OutlineViewDelegate) SetOutlineViewItemDidExpand(f func(notification foundation.Notification)) {
	di._OutlineViewItemDidExpand = f
}

// Invoked when notification is posted—that is, whenever the user expands an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530869-outlineviewitemdidexpand?language=objc
func (di *OutlineViewDelegate) OutlineViewItemDidExpand(notification foundation.Notification) {
	di._OutlineViewItemDidExpand(notification)
}
func (di *OutlineViewDelegate) HasOutlineViewSelectionIsChanging() bool {
	return di._OutlineViewSelectionIsChanging != nil
}

// Invoked when notification is posted—that is, whenever the outline view’s selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1532481-outlineviewselectionischanging?language=objc
func (di *OutlineViewDelegate) SetOutlineViewSelectionIsChanging(f func(notification foundation.Notification)) {
	di._OutlineViewSelectionIsChanging = f
}

// Invoked when notification is posted—that is, whenever the outline view’s selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1532481-outlineviewselectionischanging?language=objc
func (di *OutlineViewDelegate) OutlineViewSelectionIsChanging(notification foundation.Notification) {
	di._OutlineViewSelectionIsChanging(notification)
}
func (di *OutlineViewDelegate) HasSelectionShouldChangeInOutlineView() bool {
	return di._SelectionShouldChangeInOutlineView != nil
}

// Returns a Boolean value that indicates whether the outline view should change its selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530734-selectionshouldchangeinoutlinevi?language=objc
func (di *OutlineViewDelegate) SetSelectionShouldChangeInOutlineView(f func(outlineView OutlineView) bool) {
	di._SelectionShouldChangeInOutlineView = f
}

// Returns a Boolean value that indicates whether the outline view should change its selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530734-selectionshouldchangeinoutlinevi?language=objc
func (di *OutlineViewDelegate) SelectionShouldChangeInOutlineView(outlineView OutlineView) bool {
	return di._SelectionShouldChangeInOutlineView(outlineView)
}

// A concrete type wrapper for the [POutlineViewDelegate] protocol.
type OutlineViewDelegateWrapper struct {
	ControlTextEditingDelegateWrapper
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewItemDidCollapse() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewItemDidCollapse:"))
}

// Invoked when the did collapse notification is posted—that is, whenever the user collapses an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535557-outlineviewitemdidcollapse?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidCollapse(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewItemDidCollapse:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewItemWillExpand() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewItemWillExpand:"))
}

// Invoked when notification is posted—that is, whenever the user is about to expand an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1535847-outlineviewitemwillexpand?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillExpand(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewItemWillExpand:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewColumnDidMove() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewColumnDidMove:"))
}

// Invoked whenever the user moves a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1525297-outlineviewcolumndidmove?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidMove(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewColumnDidMove:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewItemWillCollapse() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewItemWillCollapse:"))
}

// Invoked when notification is posted—that is, whenever the user is about to collapse an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526896-outlineviewitemwillcollapse?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewItemWillCollapse(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewItemWillCollapse:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewColumnDidResize() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewColumnDidResize:"))
}

// Invoked whenever the user resizes a column in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1533372-outlineviewcolumndidresize?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewColumnDidResize(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewColumnDidResize:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewDidClickTableColumn() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:didClickTableColumn:"))
}

// Sent at the time the mouse button subsequently goes up in outlineView and tableColumn has been “clicked” without having been dragged anywhere. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1534040-outlineview?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewDidClickTableColumn(outlineView IOutlineView, tableColumn ITableColumn) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:didClickTableColumn:"), objc.Ptr(outlineView), objc.Ptr(tableColumn))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewSelectionDidChange() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewSelectionDidChange:"))
}

// Invoked when the selection did change notification is posted—that is, immediately after the outline view’s selection has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1526913-outlineviewselectiondidchange?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewSelectionDidChange:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewItemDidExpand() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewItemDidExpand:"))
}

// Invoked when notification is posted—that is, whenever the user expands an item in the outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530869-outlineviewitemdidexpand?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewItemDidExpand(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewItemDidExpand:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasOutlineViewSelectionIsChanging() bool {
	return o_.RespondsToSelector(objc.Sel("outlineViewSelectionIsChanging:"))
}

// Invoked when notification is posted—that is, whenever the outline view’s selection changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1532481-outlineviewselectionischanging?language=objc
func (o_ OutlineViewDelegateWrapper) OutlineViewSelectionIsChanging(notification foundation.INotification) {
	objc.Call[objc.Void](o_, objc.Sel("outlineViewSelectionIsChanging:"), objc.Ptr(notification))
}

func (o_ OutlineViewDelegateWrapper) HasSelectionShouldChangeInOutlineView() bool {
	return o_.RespondsToSelector(objc.Sel("selectionShouldChangeInOutlineView:"))
}

// Returns a Boolean value that indicates whether the outline view should change its selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdelegate/1530734-selectionshouldchangeinoutlinevi?language=objc
func (o_ OutlineViewDelegateWrapper) SelectionShouldChangeInOutlineView(outlineView IOutlineView) bool {
	rv := objc.Call[bool](o_, objc.Sel("selectionShouldChangeInOutlineView:"), objc.Ptr(outlineView))
	return rv
}
