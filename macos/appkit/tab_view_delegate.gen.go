// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
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
}

// A delegate implementation builder for the [PTabViewDelegate] protocol.
type TabViewDelegate struct {
	_TabViewDidChangeNumberOfTabViewItems func(tabView TabView)
	_TabViewDidSelectTabViewItem          func(tabView TabView, tabViewItem TabViewItem)
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

// A concrete type wrapper for the [PTabViewDelegate] protocol.
type TabViewDelegateWrapper struct {
	objc.Object
}

func (t_ TabViewDelegateWrapper) HasTabViewDidChangeNumberOfTabViewItems() bool {
	return t_.RespondsToSelector(objc.Sel("tabViewDidChangeNumberOfTabViewItems:"))
}

// Informs the delegate that the number of tab view items in tabView has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391657-tabviewdidchangenumberoftabviewi?language=objc
func (t_ TabViewDelegateWrapper) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	objc.Call[objc.Void](t_, objc.Sel("tabViewDidChangeNumberOfTabViewItems:"), objc.Ptr(tabView))
}

func (t_ TabViewDelegateWrapper) HasTabViewDidSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.Sel("tabView:didSelectTabViewItem:"))
}

// Informs the delegate that tabView has selected tabViewItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabviewdelegate/1391582-tabview?language=objc
func (t_ TabViewDelegateWrapper) TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("tabView:didSelectTabViewItem:"), objc.Ptr(tabView), objc.Ptr(tabViewItem))
}
