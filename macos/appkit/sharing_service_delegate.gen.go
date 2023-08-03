// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ISharingServiceDelegate interface {
	ImplementsSharingServiceWillShareItems() bool
	// optional
	SharingServiceWillShareItems(sharingService SharingService, items []objc.Object)
	ImplementsSharingServiceDidShareItems() bool
	// optional
	SharingServiceDidShareItems(sharingService SharingService, items []objc.Object)
	ImplementsSharingServiceDidFailToShareItemsError() bool
	// optional
	SharingServiceDidFailToShareItemsError(sharingService SharingService, items []objc.Object, error foundation.Error)
	ImplementsSharingServiceSourceFrameOnScreenForShareItem() bool
	// optional
	SharingServiceSourceFrameOnScreenForShareItem(sharingService SharingService, item objc.Object) foundation.Rect
	ImplementsSharingServiceTransitionImageForShareItemContentRect() bool
	// optional
	SharingServiceTransitionImageForShareItemContentRect(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage
	ImplementsSharingServiceSourceWindowForShareItemsSharingContentScope() bool
	// optional
	SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow
	ImplementsAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool
	// optional
	AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
}

type SharingServiceDelegate struct {
	_SharingServiceWillShareItems                                  func(sharingService SharingService, items []objc.Object)
	_SharingServiceDidShareItems                                   func(sharingService SharingService, items []objc.Object)
	_SharingServiceDidFailToShareItemsError                        func(sharingService SharingService, items []objc.Object, error foundation.Error)
	_SharingServiceSourceFrameOnScreenForShareItem                 func(sharingService SharingService, item objc.Object) foundation.Rect
	_SharingServiceTransitionImageForShareItemContentRect          func(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage
	_SharingServiceSourceWindowForShareItemsSharingContentScope    func(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow
	_AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
}

func (di *SharingServiceDelegate) ImplementsSharingServiceWillShareItems() bool {
	return di._SharingServiceWillShareItems != nil
}

func (di *SharingServiceDelegate) SetSharingServiceWillShareItems(f func(sharingService SharingService, items []objc.Object)) {
	di._SharingServiceWillShareItems = f
}

func (di *SharingServiceDelegate) SharingServiceWillShareItems(sharingService SharingService, items []objc.Object) {
	di._SharingServiceWillShareItems(sharingService, items)
}
func (di *SharingServiceDelegate) ImplementsSharingServiceDidShareItems() bool {
	return di._SharingServiceDidShareItems != nil
}

func (di *SharingServiceDelegate) SetSharingServiceDidShareItems(f func(sharingService SharingService, items []objc.Object)) {
	di._SharingServiceDidShareItems = f
}

func (di *SharingServiceDelegate) SharingServiceDidShareItems(sharingService SharingService, items []objc.Object) {
	di._SharingServiceDidShareItems(sharingService, items)
}
func (di *SharingServiceDelegate) ImplementsSharingServiceDidFailToShareItemsError() bool {
	return di._SharingServiceDidFailToShareItemsError != nil
}

func (di *SharingServiceDelegate) SetSharingServiceDidFailToShareItemsError(f func(sharingService SharingService, items []objc.Object, error foundation.Error)) {
	di._SharingServiceDidFailToShareItemsError = f
}

func (di *SharingServiceDelegate) SharingServiceDidFailToShareItemsError(sharingService SharingService, items []objc.Object, error foundation.Error) {
	di._SharingServiceDidFailToShareItemsError(sharingService, items, error)
}
func (di *SharingServiceDelegate) ImplementsSharingServiceSourceFrameOnScreenForShareItem() bool {
	return di._SharingServiceSourceFrameOnScreenForShareItem != nil
}

func (di *SharingServiceDelegate) SetSharingServiceSourceFrameOnScreenForShareItem(f func(sharingService SharingService, item objc.Object) foundation.Rect) {
	di._SharingServiceSourceFrameOnScreenForShareItem = f
}

func (di *SharingServiceDelegate) SharingServiceSourceFrameOnScreenForShareItem(sharingService SharingService, item objc.Object) foundation.Rect {
	return di._SharingServiceSourceFrameOnScreenForShareItem(sharingService, item)
}
func (di *SharingServiceDelegate) ImplementsSharingServiceTransitionImageForShareItemContentRect() bool {
	return di._SharingServiceTransitionImageForShareItemContentRect != nil
}

func (di *SharingServiceDelegate) SetSharingServiceTransitionImageForShareItemContentRect(f func(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage) {
	di._SharingServiceTransitionImageForShareItemContentRect = f
}

func (di *SharingServiceDelegate) SharingServiceTransitionImageForShareItemContentRect(sharingService SharingService, item objc.Object, contentRect *foundation.Rect) IImage {
	return di._SharingServiceTransitionImageForShareItemContentRect(sharingService, item, contentRect)
}
func (di *SharingServiceDelegate) ImplementsSharingServiceSourceWindowForShareItemsSharingContentScope() bool {
	return di._SharingServiceSourceWindowForShareItemsSharingContentScope != nil
}

func (di *SharingServiceDelegate) SetSharingServiceSourceWindowForShareItemsSharingContentScope(f func(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow) {
	di._SharingServiceSourceWindowForShareItemsSharingContentScope = f
}

func (di *SharingServiceDelegate) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService SharingService, items []objc.Object, sharingContentScope *SharingContentScope) IWindow {
	return di._SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService, items, sharingContentScope)
}
func (di *SharingServiceDelegate) ImplementsAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil
}

func (di *SharingServiceDelegate) SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(f func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView) {
	di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge = f
}

func (di *SharingServiceDelegate) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView {
	return di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService, positioningRect, preferredEdge)
}

type SharingServiceDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceWillShareItems() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:willShareItems:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceWillShareItems(sharingService ISharingService, items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:willShareItems:"), objc.ExtractPtr(sharingService), items)
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceDidShareItems() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:didShareItems:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceDidShareItems(sharingService ISharingService, items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:didShareItems:"), objc.ExtractPtr(sharingService), items)
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceDidFailToShareItemsError() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:didFailToShareItems:error:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceDidFailToShareItemsError(sharingService ISharingService, items []objc.IObject, error foundation.IError) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("sharingService:didFailToShareItems:error:"), objc.ExtractPtr(sharingService), items, objc.ExtractPtr(error))
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceSourceFrameOnScreenForShareItem() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:sourceFrameOnScreenForShareItem:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceSourceFrameOnScreenForShareItem(sharingService ISharingService, item objc.IObject) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](s_, objc.GetSelector("sharingService:sourceFrameOnScreenForShareItem:"), objc.ExtractPtr(sharingService), objc.ExtractPtr(item))
	return rv
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceTransitionImageForShareItemContentRect() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:transitionImageForShareItem:contentRect:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceTransitionImageForShareItemContentRect(sharingService ISharingService, item objc.IObject, contentRect *foundation.Rect) Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("sharingService:transitionImageForShareItem:contentRect:"), objc.ExtractPtr(sharingService), objc.ExtractPtr(item), contentRect)
	return rv
}

func (s_ SharingServiceDelegateWrapper) ImplementsSharingServiceSourceWindowForShareItemsSharingContentScope() bool {
	return s_.RespondsToSelector(objc.GetSelector("sharingService:sourceWindowForShareItems:sharingContentScope:"))
}

func (s_ SharingServiceDelegateWrapper) SharingServiceSourceWindowForShareItemsSharingContentScope(sharingService ISharingService, items []objc.IObject, sharingContentScope *SharingContentScope) Window {
	rv := objc.CallMethod[Window](s_, objc.GetSelector("sharingService:sourceWindowForShareItems:sharingContentScope:"), objc.ExtractPtr(sharingService), items, sharingContentScope)
	return rv
}

func (s_ SharingServiceDelegateWrapper) ImplementsAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return s_.RespondsToSelector(objc.GetSelector("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"))
}

func (s_ SharingServiceDelegateWrapper) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) View {
	rv := objc.CallMethod[View](s_, objc.GetSelector("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), objc.ExtractPtr(sharingService), positioningRect, preferredEdge)
	return rv
}
