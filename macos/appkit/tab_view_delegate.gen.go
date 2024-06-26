// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// The NSTabViewDelegate protocol defines the optional methods implemented by delegates of NSTabView objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate?language=objc
type PTabViewDelegate interface {
	// optional
	TabViewDidChangeNumberOfTabViewItems(tabView TabView)
	HasTabViewDidChangeNumberOfTabViewItems() bool

	// optional
	TabViewDidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
	HasTabViewDidSelectTabViewItem() bool

	// optional
	TabViewShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool
	HasTabViewShouldSelectTabViewItem() bool

	// optional
	TabViewWillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
	HasTabViewWillSelectTabViewItem() bool
}

// A delegate implementation builder for the [PTabViewDelegate] protocol.
type TabViewDelegate struct {
	_TabViewDidChangeNumberOfTabViewItems func(tabView TabView)
	_TabViewDidSelectTabViewItem          func(tabView TabView, tabViewItem TabViewItem)
	_TabViewShouldSelectTabViewItem       func(tabView TabView, tabViewItem TabViewItem) bool
	_TabViewWillSelectTabViewItem         func(tabView TabView, tabViewItem TabViewItem)
}

func (di *TabViewDelegate) HasTabViewDidChangeNumberOfTabViewItems() bool {
	return di._TabViewDidChangeNumberOfTabViewItems != nil
}

// Informs the delegate that the number of tab view items in tabView has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391657-tabviewdidchangenumberoftabviewi?language=objc
func (di *TabViewDelegate) SetTabViewDidChangeNumberOfTabViewItems(f func(tabView TabView)) {
	di._TabViewDidChangeNumberOfTabViewItems = f
}

// Informs the delegate that the number of tab view items in tabView has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391657-tabviewdidchangenumberoftabviewi?language=objc
func (di *TabViewDelegate) TabViewDidChangeNumberOfTabViewItems(tabView TabView) {
	di._TabViewDidChangeNumberOfTabViewItems(tabView)
}
func (di *TabViewDelegate) HasTabViewDidSelectTabViewItem() bool {
	return di._TabViewDidSelectTabViewItem != nil
}

// Informs the delegate that tabView has selected tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391582-tabview?language=objc
func (di *TabViewDelegate) SetTabViewDidSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabViewDidSelectTabViewItem = f
}

// Informs the delegate that tabView has selected tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391582-tabview?language=objc
func (di *TabViewDelegate) TabViewDidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabViewDidSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegate) HasTabViewShouldSelectTabViewItem() bool {
	return di._TabViewShouldSelectTabViewItem != nil
}

// Invoked just before tabViewItem in tabView is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391651-tabview?language=objc
func (di *TabViewDelegate) SetTabViewShouldSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem) bool) {
	di._TabViewShouldSelectTabViewItem = f
}

// Invoked just before tabViewItem in tabView is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391651-tabview?language=objc
func (di *TabViewDelegate) TabViewShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool {
	return di._TabViewShouldSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegate) HasTabViewWillSelectTabViewItem() bool {
	return di._TabViewWillSelectTabViewItem != nil
}

// Informs the delegate that tabView is about to select tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391611-tabview?language=objc
func (di *TabViewDelegate) SetTabViewWillSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabViewWillSelectTabViewItem = f
}

// Informs the delegate that tabView is about to select tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391611-tabview?language=objc
func (di *TabViewDelegate) TabViewWillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabViewWillSelectTabViewItem(tabView, tabViewItem)
}

// ensure impl type implements protocol interface
var _ PTabViewDelegate = (*TabViewDelegateObject)(nil)

// A concrete type for the [PTabViewDelegate] protocol.
type TabViewDelegateObject struct {
	objc.Object
}

func (t_ TabViewDelegateObject) HasTabViewDidChangeNumberOfTabViewItems() bool {
	return t_.RespondsToSelector(objc.Sel("tabViewDidChangeNumberOfTabViewItems:"))
}

// Informs the delegate that the number of tab view items in tabView has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391657-tabviewdidchangenumberoftabviewi?language=objc
func (t_ TabViewDelegateObject) TabViewDidChangeNumberOfTabViewItems(tabView TabView) {
	objc.Call[objc.Void](t_, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), tabView)
}

func (t_ TabViewDelegateObject) HasTabViewDidSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.Sel("tabView:didSelectTabViewItem:"))
}

// Informs the delegate that tabView has selected tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391582-tabview?language=objc
func (t_ TabViewDelegateObject) TabViewDidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("tabView:didSelectTabViewItem:"), tabView, tabViewItem)
}

func (t_ TabViewDelegateObject) HasTabViewShouldSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.Sel("tabView:shouldSelectTabViewItem:"))
}

// Invoked just before tabViewItem in tabView is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391651-tabview?language=objc
func (t_ TabViewDelegateObject) TabViewShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool {
	rv := objc.Call[bool](t_, objc.Sel("tabView:shouldSelectTabViewItem:"), tabView, tabViewItem)
	return rv
}

func (t_ TabViewDelegateObject) HasTabViewWillSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.Sel("tabView:willSelectTabViewItem:"))
}

// Informs the delegate that tabView is about to select tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391611-tabview?language=objc
func (t_ TabViewDelegateObject) TabViewWillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("tabView:willSelectTabViewItem:"), tabView, tabViewItem)
}
