// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ITabViewDelegate interface {
	ImplementsTabViewDidChangeNumberOfTabViewItems() bool
	// optional
	TabViewDidChangeNumberOfTabViewItems(tabView TabView)
	ImplementsTabViewShouldSelectTabViewItem() bool
	// optional
	TabViewShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool
	ImplementsTabViewWillSelectTabViewItem() bool
	// optional
	TabViewWillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
	ImplementsTabViewDidSelectTabViewItem() bool
	// optional
	TabViewDidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem)
}

type TabViewDelegate struct {
	_TabViewDidChangeNumberOfTabViewItems func(tabView TabView)
	_TabViewShouldSelectTabViewItem       func(tabView TabView, tabViewItem TabViewItem) bool
	_TabViewWillSelectTabViewItem         func(tabView TabView, tabViewItem TabViewItem)
	_TabViewDidSelectTabViewItem          func(tabView TabView, tabViewItem TabViewItem)
}

func (di *TabViewDelegate) ImplementsTabViewDidChangeNumberOfTabViewItems() bool {
	return di._TabViewDidChangeNumberOfTabViewItems != nil
}

func (di *TabViewDelegate) SetTabViewDidChangeNumberOfTabViewItems(f func(tabView TabView)) {
	di._TabViewDidChangeNumberOfTabViewItems = f
}

func (di *TabViewDelegate) TabViewDidChangeNumberOfTabViewItems(tabView TabView) {
	di._TabViewDidChangeNumberOfTabViewItems(tabView)
}
func (di *TabViewDelegate) ImplementsTabViewShouldSelectTabViewItem() bool {
	return di._TabViewShouldSelectTabViewItem != nil
}

func (di *TabViewDelegate) SetTabViewShouldSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem) bool) {
	di._TabViewShouldSelectTabViewItem = f
}

func (di *TabViewDelegate) TabViewShouldSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) bool {
	return di._TabViewShouldSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegate) ImplementsTabViewWillSelectTabViewItem() bool {
	return di._TabViewWillSelectTabViewItem != nil
}

func (di *TabViewDelegate) SetTabViewWillSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabViewWillSelectTabViewItem = f
}

func (di *TabViewDelegate) TabViewWillSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabViewWillSelectTabViewItem(tabView, tabViewItem)
}
func (di *TabViewDelegate) ImplementsTabViewDidSelectTabViewItem() bool {
	return di._TabViewDidSelectTabViewItem != nil
}

func (di *TabViewDelegate) SetTabViewDidSelectTabViewItem(f func(tabView TabView, tabViewItem TabViewItem)) {
	di._TabViewDidSelectTabViewItem = f
}

func (di *TabViewDelegate) TabViewDidSelectTabViewItem(tabView TabView, tabViewItem TabViewItem) {
	di._TabViewDidSelectTabViewItem(tabView, tabViewItem)
}

type TabViewDelegateWrapper struct {
	objc.Object
}

func (t_ TabViewDelegateWrapper) ImplementsTabViewDidChangeNumberOfTabViewItems() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabViewDidChangeNumberOfTabViewItems:"))
}

func (t_ TabViewDelegateWrapper) TabViewDidChangeNumberOfTabViewItems(tabView ITabView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabViewDidChangeNumberOfTabViewItems:"), objc.ExtractPtr(tabView))
}

func (t_ TabViewDelegateWrapper) ImplementsTabViewShouldSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:shouldSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabViewShouldSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tabView:shouldSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
	return rv
}

func (t_ TabViewDelegateWrapper) ImplementsTabViewWillSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:willSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabViewWillSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabView:willSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
}

func (t_ TabViewDelegateWrapper) ImplementsTabViewDidSelectTabViewItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tabView:didSelectTabViewItem:"))
}

func (t_ TabViewDelegateWrapper) TabViewDidSelectTabViewItem(tabView ITabView, tabViewItem ITabViewItem) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tabView:didSelectTabViewItem:"), objc.ExtractPtr(tabView), objc.ExtractPtr(tabViewItem))
}
