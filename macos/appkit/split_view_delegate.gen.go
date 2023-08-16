// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that a delegate of a split view implements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate?language=objc
type PSplitViewDelegate interface {
	// optional
	SplitViewWillResizeSubviews(notification foundation.Notification)
	HasSplitViewWillResizeSubviews() bool

	// optional
	SplitViewDidResizeSubviews(notification foundation.Notification)
	HasSplitViewDidResizeSubviews() bool

	// optional
	SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool
	HasSplitViewShouldHideDividerAtIndex() bool
}

// A delegate implementation builder for the [PSplitViewDelegate] protocol.
type SplitViewDelegate struct {
	_SplitViewWillResizeSubviews       func(notification foundation.Notification)
	_SplitViewDidResizeSubviews        func(notification foundation.Notification)
	_SplitViewShouldHideDividerAtIndex func(splitView SplitView, dividerIndex int) bool
}

func (di *SplitViewDelegate) HasSplitViewWillResizeSubviews() bool {
	return di._SplitViewWillResizeSubviews != nil
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (di *SplitViewDelegate) SetSplitViewWillResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewWillResizeSubviews = f
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (di *SplitViewDelegate) SplitViewWillResizeSubviews(notification foundation.Notification) {
	di._SplitViewWillResizeSubviews(notification)
}
func (di *SplitViewDelegate) HasSplitViewDidResizeSubviews() bool {
	return di._SplitViewDidResizeSubviews != nil
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (di *SplitViewDelegate) SetSplitViewDidResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewDidResizeSubviews = f
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (di *SplitViewDelegate) SplitViewDidResizeSubviews(notification foundation.Notification) {
	di._SplitViewDidResizeSubviews(notification)
}
func (di *SplitViewDelegate) HasSplitViewShouldHideDividerAtIndex() bool {
	return di._SplitViewShouldHideDividerAtIndex != nil
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (di *SplitViewDelegate) SetSplitViewShouldHideDividerAtIndex(f func(splitView SplitView, dividerIndex int) bool) {
	di._SplitViewShouldHideDividerAtIndex = f
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (di *SplitViewDelegate) SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	return di._SplitViewShouldHideDividerAtIndex(splitView, dividerIndex)
}

// A concrete type wrapper for the [PSplitViewDelegate] protocol.
type SplitViewDelegateWrapper struct {
	objc.Object
}

func (s_ SplitViewDelegateWrapper) HasSplitViewWillResizeSubviews() bool {
	return s_.RespondsToSelector(objc.Sel("splitViewWillResizeSubviews:"))
}

// Notifies the delegate when the split view is about to resize its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455289-splitviewwillresizesubviews?language=objc
func (s_ SplitViewDelegateWrapper) SplitViewWillResizeSubviews(notification foundation.INotification) {
	objc.Call[objc.Void](s_, objc.Sel("splitViewWillResizeSubviews:"), objc.Ptr(notification))
}

func (s_ SplitViewDelegateWrapper) HasSplitViewDidResizeSubviews() bool {
	return s_.RespondsToSelector(objc.Sel("splitViewDidResizeSubviews:"))
}

// Notifies the delegate when the split view resizes its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455314-splitviewdidresizesubviews?language=objc
func (s_ SplitViewDelegateWrapper) SplitViewDidResizeSubviews(notification foundation.INotification) {
	objc.Call[objc.Void](s_, objc.Sel("splitViewDidResizeSubviews:"), objc.Ptr(notification))
}

func (s_ SplitViewDelegateWrapper) HasSplitViewShouldHideDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("splitView:shouldHideDividerAtIndex:"))
}

// Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate/1455280-splitview?language=objc
func (s_ SplitViewDelegateWrapper) SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	rv := objc.Call[bool](s_, objc.Sel("splitView:shouldHideDividerAtIndex:"), objc.Ptr(splitView), dividerIndex)
	return rv
}
