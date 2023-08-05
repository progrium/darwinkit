// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ISplitViewDelegate interface {
	ImplementsSplitViewWillResizeSubviews() bool
	// optional
	SplitViewWillResizeSubviews(notification foundation.Notification)
	ImplementsSplitViewDidResizeSubviews() bool
	// optional
	SplitViewDidResizeSubviews(notification foundation.Notification)
	ImplementsSplitViewCanCollapseSubview() bool
	// optional
	SplitViewCanCollapseSubview(splitView SplitView, subview View) bool
	ImplementsSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool
	// optional
	// deprecated
	SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView SplitView, subview View, dividerIndex int) bool
	ImplementsSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool
	// optional
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	ImplementsSplitViewShouldHideDividerAtIndex() bool
	// optional
	SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool
	ImplementsSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool
	// optional
	SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect
	ImplementsSplitViewConstrainSplitPositionOfSubviewAt() bool
	// optional
	SplitViewConstrainSplitPositionOfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	ImplementsSplitViewConstrainMinCoordinateOfSubviewAt() bool
	// optional
	SplitViewConstrainMinCoordinateOfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
	ImplementsSplitViewConstrainMaxCoordinateOfSubviewAt() bool
	// optional
	SplitViewConstrainMaxCoordinateOfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	ImplementsSplitViewResizeSubviewsWithOldSize() bool
	// optional
	SplitViewResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size)
	ImplementsSplitViewShouldAdjustSizeOfSubview() bool
	// optional
	SplitViewShouldAdjustSizeOfSubview(splitView SplitView, view View) bool
}

type SplitViewDelegate struct {
	_SplitViewWillResizeSubviews                                  func(notification foundation.Notification)
	_SplitViewDidResizeSubviews                                   func(notification foundation.Notification)
	_SplitViewCanCollapseSubview                                  func(splitView SplitView, subview View) bool
	_SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex func(splitView SplitView, subview View, dividerIndex int) bool
	_SplitViewEffectiveRectForDrawnRectOfDividerAtIndex           func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect
	_SplitViewShouldHideDividerAtIndex                            func(splitView SplitView, dividerIndex int) bool
	_SplitViewAdditionalEffectiveRectOfDividerAtIndex             func(splitView SplitView, dividerIndex int) foundation.Rect
	_SplitViewConstrainSplitPositionOfSubviewAt                   func(splitView SplitView, proposedPosition float64, dividerIndex int) float64
	_SplitViewConstrainMinCoordinateOfSubviewAt                   func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64
	_SplitViewConstrainMaxCoordinateOfSubviewAt                   func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64
	_SplitViewResizeSubviewsWithOldSize                           func(splitView SplitView, oldSize foundation.Size)
	_SplitViewShouldAdjustSizeOfSubview                           func(splitView SplitView, view View) bool
}

func (di *SplitViewDelegate) ImplementsSplitViewWillResizeSubviews() bool {
	return di._SplitViewWillResizeSubviews != nil
}

func (di *SplitViewDelegate) SetSplitViewWillResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewWillResizeSubviews = f
}

func (di *SplitViewDelegate) SplitViewWillResizeSubviews(notification foundation.Notification) {
	di._SplitViewWillResizeSubviews(notification)
}
func (di *SplitViewDelegate) ImplementsSplitViewDidResizeSubviews() bool {
	return di._SplitViewDidResizeSubviews != nil
}

func (di *SplitViewDelegate) SetSplitViewDidResizeSubviews(f func(notification foundation.Notification)) {
	di._SplitViewDidResizeSubviews = f
}

func (di *SplitViewDelegate) SplitViewDidResizeSubviews(notification foundation.Notification) {
	di._SplitViewDidResizeSubviews(notification)
}
func (di *SplitViewDelegate) ImplementsSplitViewCanCollapseSubview() bool {
	return di._SplitViewCanCollapseSubview != nil
}

func (di *SplitViewDelegate) SetSplitViewCanCollapseSubview(f func(splitView SplitView, subview View) bool) {
	di._SplitViewCanCollapseSubview = f
}

func (di *SplitViewDelegate) SplitViewCanCollapseSubview(splitView SplitView, subview View) bool {
	return di._SplitViewCanCollapseSubview(splitView, subview)
}
func (di *SplitViewDelegate) ImplementsSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool {
	return di._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex != nil
}

// deprecated
func (di *SplitViewDelegate) SetSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(f func(splitView SplitView, subview View, dividerIndex int) bool) {
	di._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex = f
}

// deprecated
func (di *SplitViewDelegate) SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView SplitView, subview View, dividerIndex int) bool {
	return di._SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(splitView, subview, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex != nil
}

func (di *SplitViewDelegate) SetSplitViewEffectiveRectForDrawnRectOfDividerAtIndex(f func(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect) {
	di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex = f
}

func (di *SplitViewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView SplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	return di._SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView, proposedEffectiveRect, drawnRect, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewShouldHideDividerAtIndex() bool {
	return di._SplitViewShouldHideDividerAtIndex != nil
}

func (di *SplitViewDelegate) SetSplitViewShouldHideDividerAtIndex(f func(splitView SplitView, dividerIndex int) bool) {
	di._SplitViewShouldHideDividerAtIndex = f
}

func (di *SplitViewDelegate) SplitViewShouldHideDividerAtIndex(splitView SplitView, dividerIndex int) bool {
	return di._SplitViewShouldHideDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return di._SplitViewAdditionalEffectiveRectOfDividerAtIndex != nil
}

func (di *SplitViewDelegate) SetSplitViewAdditionalEffectiveRectOfDividerAtIndex(f func(splitView SplitView, dividerIndex int) foundation.Rect) {
	di._SplitViewAdditionalEffectiveRectOfDividerAtIndex = f
}

func (di *SplitViewDelegate) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView SplitView, dividerIndex int) foundation.Rect {
	return di._SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return di._SplitViewConstrainSplitPositionOfSubviewAt != nil
}

func (di *SplitViewDelegate) SetSplitViewConstrainSplitPositionOfSubviewAt(f func(splitView SplitView, proposedPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainSplitPositionOfSubviewAt = f
}

func (di *SplitViewDelegate) SplitViewConstrainSplitPositionOfSubviewAt(splitView SplitView, proposedPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainSplitPositionOfSubviewAt(splitView, proposedPosition, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return di._SplitViewConstrainMinCoordinateOfSubviewAt != nil
}

func (di *SplitViewDelegate) SetSplitViewConstrainMinCoordinateOfSubviewAt(f func(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainMinCoordinateOfSubviewAt = f
}

func (di *SplitViewDelegate) SplitViewConstrainMinCoordinateOfSubviewAt(splitView SplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainMinCoordinateOfSubviewAt(splitView, proposedMinimumPosition, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return di._SplitViewConstrainMaxCoordinateOfSubviewAt != nil
}

func (di *SplitViewDelegate) SetSplitViewConstrainMaxCoordinateOfSubviewAt(f func(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64) {
	di._SplitViewConstrainMaxCoordinateOfSubviewAt = f
}

func (di *SplitViewDelegate) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView SplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	return di._SplitViewConstrainMaxCoordinateOfSubviewAt(splitView, proposedMaximumPosition, dividerIndex)
}
func (di *SplitViewDelegate) ImplementsSplitViewResizeSubviewsWithOldSize() bool {
	return di._SplitViewResizeSubviewsWithOldSize != nil
}

func (di *SplitViewDelegate) SetSplitViewResizeSubviewsWithOldSize(f func(splitView SplitView, oldSize foundation.Size)) {
	di._SplitViewResizeSubviewsWithOldSize = f
}

func (di *SplitViewDelegate) SplitViewResizeSubviewsWithOldSize(splitView SplitView, oldSize foundation.Size) {
	di._SplitViewResizeSubviewsWithOldSize(splitView, oldSize)
}
func (di *SplitViewDelegate) ImplementsSplitViewShouldAdjustSizeOfSubview() bool {
	return di._SplitViewShouldAdjustSizeOfSubview != nil
}

func (di *SplitViewDelegate) SetSplitViewShouldAdjustSizeOfSubview(f func(splitView SplitView, view View) bool) {
	di._SplitViewShouldAdjustSizeOfSubview = f
}

func (di *SplitViewDelegate) SplitViewShouldAdjustSizeOfSubview(splitView SplitView, view View) bool {
	return di._SplitViewShouldAdjustSizeOfSubview(splitView, view)
}

type SplitViewDelegateWrapper struct {
	objc.Object
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewWillResizeSubviews() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitViewWillResizeSubviews:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewWillResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewWillResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewDidResizeSubviews() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitViewDidResizeSubviews:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewDidResizeSubviews(notification foundation.INotification) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitViewDidResizeSubviews:"), objc.ExtractPtr(notification))
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewCanCollapseSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:canCollapseSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewCanCollapseSubview(splitView ISplitView, subview IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:canCollapseSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(subview))
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewEffectiveRectForDrawnRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView ISplitView, proposedEffectiveRect foundation.Rect, drawnRect foundation.Rect, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), objc.ExtractPtr(splitView), proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewShouldHideDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldHideDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewShouldHideDividerAtIndex(splitView ISplitView, dividerIndex int) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldHideDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewAdditionalEffectiveRectOfDividerAtIndex() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView ISplitView, dividerIndex int) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("splitView:additionalEffectiveRectOfDividerAtIndex:"), objc.ExtractPtr(splitView), dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewConstrainSplitPositionOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewConstrainSplitPositionOfSubviewAt(splitView ISplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainSplitPosition:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewConstrainMinCoordinateOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewConstrainMinCoordinateOfSubviewAt(splitView ISplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMinCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMinimumPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewConstrainMaxCoordinateOfSubviewAt() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView ISplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.CallMethod[float64](s_, objc.GetSelector("splitView:constrainMaxCoordinate:ofSubviewAt:"), objc.ExtractPtr(splitView), proposedMaximumPosition, dividerIndex)
	return rv
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewResizeSubviewsWithOldSize() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:resizeSubviewsWithOldSize:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewResizeSubviewsWithOldSize(splitView ISplitView, oldSize foundation.Size) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("splitView:resizeSubviewsWithOldSize:"), objc.ExtractPtr(splitView), oldSize)
}

func (s_ SplitViewDelegateWrapper) ImplementsSplitViewShouldAdjustSizeOfSubview() bool {
	return s_.RespondsToSelector(objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"))
}

func (s_ SplitViewDelegateWrapper) SplitViewShouldAdjustSizeOfSubview(splitView ISplitView, view IView) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("splitView:shouldAdjustSizeOfSubview:"), objc.ExtractPtr(splitView), objc.ExtractPtr(view))
	return rv
}
