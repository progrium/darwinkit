// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a scrubber delegate implements to respond to user interactions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate?language=objc
type PScrubberDelegate interface {
	// optional
	DidFinishInteractingWithScrubber(scrubber Scrubber)
	HasDidFinishInteractingWithScrubber() bool

	// optional
	ScrubberDidSelectItemAtIndex(scrubber Scrubber, selectedIndex int)
	HasScrubberDidSelectItemAtIndex() bool

	// optional
	DidBeginInteractingWithScrubber(scrubber Scrubber)
	HasDidBeginInteractingWithScrubber() bool

	// optional
	DidCancelInteractingWithScrubber(scrubber Scrubber)
	HasDidCancelInteractingWithScrubber() bool
}

// A delegate implementation builder for the [PScrubberDelegate] protocol.
type ScrubberDelegate struct {
	_DidFinishInteractingWithScrubber func(scrubber Scrubber)
	_ScrubberDidSelectItemAtIndex     func(scrubber Scrubber, selectedIndex int)
	_DidBeginInteractingWithScrubber  func(scrubber Scrubber)
	_DidCancelInteractingWithScrubber func(scrubber Scrubber)
}

func (di *ScrubberDelegate) HasDidFinishInteractingWithScrubber() bool {
	return di._DidFinishInteractingWithScrubber != nil
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidFinishInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidFinishInteractingWithScrubber = f
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidFinishInteractingWithScrubber(scrubber Scrubber) {
	di._DidFinishInteractingWithScrubber(scrubber)
}
func (di *ScrubberDelegate) HasScrubberDidSelectItemAtIndex() bool {
	return di._ScrubberDidSelectItemAtIndex != nil
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (di *ScrubberDelegate) SetScrubberDidSelectItemAtIndex(f func(scrubber Scrubber, selectedIndex int)) {
	di._ScrubberDidSelectItemAtIndex = f
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (di *ScrubberDelegate) ScrubberDidSelectItemAtIndex(scrubber Scrubber, selectedIndex int) {
	di._ScrubberDidSelectItemAtIndex(scrubber, selectedIndex)
}
func (di *ScrubberDelegate) HasDidBeginInteractingWithScrubber() bool {
	return di._DidBeginInteractingWithScrubber != nil
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidBeginInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidBeginInteractingWithScrubber = f
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidBeginInteractingWithScrubber(scrubber Scrubber) {
	di._DidBeginInteractingWithScrubber(scrubber)
}
func (di *ScrubberDelegate) HasDidCancelInteractingWithScrubber() bool {
	return di._DidCancelInteractingWithScrubber != nil
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) SetDidCancelInteractingWithScrubber(f func(scrubber Scrubber)) {
	di._DidCancelInteractingWithScrubber = f
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (di *ScrubberDelegate) DidCancelInteractingWithScrubber(scrubber Scrubber) {
	di._DidCancelInteractingWithScrubber(scrubber)
}

// A concrete type wrapper for the [PScrubberDelegate] protocol.
type ScrubberDelegateWrapper struct {
	objc.Object
}

func (s_ ScrubberDelegateWrapper) HasDidFinishInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didFinishInteractingWithScrubber:"))
}

// Tells the delegate that a pan or scroll interaction with the scrubber has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544653-didfinishinteractingwithscrubber?language=objc
func (s_ ScrubberDelegateWrapper) DidFinishInteractingWithScrubber(scrubber IScrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didFinishInteractingWithScrubber:"), objc.Ptr(scrubber))
}

func (s_ ScrubberDelegateWrapper) HasScrubberDidSelectItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:didSelectItemAtIndex:"))
}

// Tells the delegate that the item at the specified index was selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544714-scrubber?language=objc
func (s_ ScrubberDelegateWrapper) ScrubberDidSelectItemAtIndex(scrubber IScrubber, selectedIndex int) {
	objc.Call[objc.Void](s_, objc.Sel("scrubber:didSelectItemAtIndex:"), objc.Ptr(scrubber), selectedIndex)
}

func (s_ ScrubberDelegateWrapper) HasDidBeginInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didBeginInteractingWithScrubber:"))
}

// Tells the delegate that the user is panning or scrolling the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2544657-didbegininteractingwithscrubber?language=objc
func (s_ ScrubberDelegateWrapper) DidBeginInteractingWithScrubber(scrubber IScrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didBeginInteractingWithScrubber:"), objc.Ptr(scrubber))
}

func (s_ ScrubberDelegateWrapper) HasDidCancelInteractingWithScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("didCancelInteractingWithScrubber:"))
}

// Tells the delegate that a user interaction with the scrubber has been canceled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdelegate/2646905-didcancelinteractingwithscrubber?language=objc
func (s_ ScrubberDelegateWrapper) DidCancelInteractingWithScrubber(scrubber IScrubber) {
	objc.Call[objc.Void](s_, objc.Sel("didCancelInteractingWithScrubber:"), objc.Ptr(scrubber))
}
