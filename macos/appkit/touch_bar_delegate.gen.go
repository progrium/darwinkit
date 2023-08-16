// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that allows you to provide the items for a bar dynamically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate?language=objc
type PTouchBarDelegate interface {
	// optional
	TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
	HasTouchBarMakeItemForIdentifier() bool
}

// A delegate implementation builder for the [PTouchBarDelegate] protocol.
type TouchBarDelegate struct {
	_TouchBarMakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

func (di *TouchBarDelegate) HasTouchBarMakeItemForIdentifier() bool {
	return di._TouchBarMakeItemForIdentifier != nil
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (di *TouchBarDelegate) SetTouchBarMakeItemForIdentifier(f func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem) {
	di._TouchBarMakeItemForIdentifier = f
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (di *TouchBarDelegate) TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem {
	return di._TouchBarMakeItemForIdentifier(touchBar, identifier)
}

// A concrete type wrapper for the [PTouchBarDelegate] protocol.
type TouchBarDelegateWrapper struct {
	objc.Object
}

func (t_ TouchBarDelegateWrapper) HasTouchBarMakeItemForIdentifier() bool {
	return t_.RespondsToSelector(objc.Sel("touchBar:makeItemForIdentifier:"))
}

// Asks the delegate object for the bar item for the specified bar and item identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbardelegate/2544851-touchbar?language=objc
func (t_ TouchBarDelegateWrapper) TouchBarMakeItemForIdentifier(touchBar ITouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.Call[TouchBarItem](t_, objc.Sel("touchBar:makeItemForIdentifier:"), objc.Ptr(touchBar), identifier)
	return rv
}
