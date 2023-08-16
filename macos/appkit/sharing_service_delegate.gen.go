// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that you use to customize the position and animation of a share sheet, and to be notified whether the item is successfully shared. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate?language=objc
type PSharingServiceDelegate interface {
	// optional
	SharingServiceDidShareItems(sharingService SharingService, items []objc.Object)
	HasSharingServiceDidShareItems() bool

	// optional
	AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
	HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool
}

// A delegate implementation builder for the [PSharingServiceDelegate] protocol.
type SharingServiceDelegate struct {
	_SharingServiceDidShareItems                                   func(sharingService SharingService, items []objc.Object)
	_AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView
}

func (di *SharingServiceDelegate) HasSharingServiceDidShareItems() bool {
	return di._SharingServiceDidShareItems != nil
}

// Invoked when the sharing service has finished sharing the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1402638-sharingservice?language=objc
func (di *SharingServiceDelegate) SetSharingServiceDidShareItems(f func(sharingService SharingService, items []objc.Object)) {
	di._SharingServiceDidShareItems = f
}

// Invoked when the sharing service has finished sharing the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1402638-sharingservice?language=objc
func (di *SharingServiceDelegate) SharingServiceDidShareItems(sharingService SharingService, items []objc.Object) {
	di._SharingServiceDidShareItems(sharingService, items)
}
func (di *SharingServiceDelegate) HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1644711-anchoringviewforsharingservice?language=objc
func (di *SharingServiceDelegate) SetAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(f func(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView) {
	di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1644711-anchoringviewforsharingservice?language=objc
func (di *SharingServiceDelegate) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService SharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) IView {
	return di._AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService, positioningRect, preferredEdge)
}

// A concrete type wrapper for the [PSharingServiceDelegate] protocol.
type SharingServiceDelegateWrapper struct {
	objc.Object
}

func (s_ SharingServiceDelegateWrapper) HasSharingServiceDidShareItems() bool {
	return s_.RespondsToSelector(objc.Sel("sharingService:didShareItems:"))
}

// Invoked when the sharing service has finished sharing the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1402638-sharingservice?language=objc
func (s_ SharingServiceDelegateWrapper) SharingServiceDidShareItems(sharingService ISharingService, items []objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("sharingService:didShareItems:"), objc.Ptr(sharingService), items)
}

func (s_ SharingServiceDelegateWrapper) HasAnchoringViewForSharingServiceShowRelativeToRectPreferredEdge() bool {
	return s_.RespondsToSelector(objc.Sel("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicedelegate/1644711-anchoringviewforsharingservice?language=objc
func (s_ SharingServiceDelegateWrapper) AnchoringViewForSharingServiceShowRelativeToRectPreferredEdge(sharingService ISharingService, positioningRect *foundation.Rect, preferredEdge *foundation.RectEdge) View {
	rv := objc.Call[View](s_, objc.Sel("anchoringViewForSharingService:showRelativeToRect:preferredEdge:"), objc.Ptr(sharingService), positioningRect, preferredEdge)
	return rv
}
