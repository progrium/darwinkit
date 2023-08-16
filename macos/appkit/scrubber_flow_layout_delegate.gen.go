// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A protocol that a scrubber delegate can adopt to provide the size of an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayoutdelegate?language=objc
type PScrubberFlowLayoutDelegate interface {
	// optional
	ScrubberLayoutSizeForItemAtIndex(scrubber Scrubber, layout ScrubberFlowLayout, itemIndex int) foundation.Size
	HasScrubberLayoutSizeForItemAtIndex() bool
}

// A delegate implementation builder for the [PScrubberFlowLayoutDelegate] protocol.
type ScrubberFlowLayoutDelegate struct {
	_ScrubberLayoutSizeForItemAtIndex func(scrubber Scrubber, layout ScrubberFlowLayout, itemIndex int) foundation.Size
}

func (di *ScrubberFlowLayoutDelegate) HasScrubberLayoutSizeForItemAtIndex() bool {
	return di._ScrubberLayoutSizeForItemAtIndex != nil
}

// Asks the delegate for the size of each item in a scrubber whose items are arranged in a flow layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayoutdelegate/2544630-scrubber?language=objc
func (di *ScrubberFlowLayoutDelegate) SetScrubberLayoutSizeForItemAtIndex(f func(scrubber Scrubber, layout ScrubberFlowLayout, itemIndex int) foundation.Size) {
	di._ScrubberLayoutSizeForItemAtIndex = f
}

// Asks the delegate for the size of each item in a scrubber whose items are arranged in a flow layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayoutdelegate/2544630-scrubber?language=objc
func (di *ScrubberFlowLayoutDelegate) ScrubberLayoutSizeForItemAtIndex(scrubber Scrubber, layout ScrubberFlowLayout, itemIndex int) foundation.Size {
	return di._ScrubberLayoutSizeForItemAtIndex(scrubber, layout, itemIndex)
}

// A concrete type wrapper for the [PScrubberFlowLayoutDelegate] protocol.
type ScrubberFlowLayoutDelegateWrapper struct {
	objc.Object
}

func (s_ ScrubberFlowLayoutDelegateWrapper) HasScrubberLayoutSizeForItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:layout:sizeForItemAtIndex:"))
}

// Asks the delegate for the size of each item in a scrubber whose items are arranged in a flow layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberflowlayoutdelegate/2544630-scrubber?language=objc
func (s_ ScrubberFlowLayoutDelegateWrapper) ScrubberLayoutSizeForItemAtIndex(scrubber IScrubber, layout IScrubberFlowLayout, itemIndex int) foundation.Size {
	rv := objc.Call[foundation.Size](s_, objc.Sel("scrubber:layout:sizeForItemAtIndex:"), objc.Ptr(scrubber), objc.Ptr(layout), itemIndex)
	return rv
}
