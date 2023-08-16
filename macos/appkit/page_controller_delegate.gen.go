// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// The NSPageControllerDelegate protocol allows you to customize the behavior of instances of the NSPageController class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate?language=objc
type PPageControllerDelegate interface {
	// optional
	PageControllerDidEndLiveTransition(pageController PageController)
	HasPageControllerDidEndLiveTransition() bool

	// optional
	PageControllerWillStartLiveTransition(pageController PageController)
	HasPageControllerWillStartLiveTransition() bool

	// optional
	PageControllerDidTransitionToObject(pageController PageController, object objc.Object)
	HasPageControllerDidTransitionToObject() bool
}

// A delegate implementation builder for the [PPageControllerDelegate] protocol.
type PageControllerDelegate struct {
	_PageControllerDidEndLiveTransition    func(pageController PageController)
	_PageControllerWillStartLiveTransition func(pageController PageController)
	_PageControllerDidTransitionToObject   func(pageController PageController, object objc.Object)
}

func (di *PageControllerDelegate) HasPageControllerDidEndLiveTransition() bool {
	return di._PageControllerDidEndLiveTransition != nil
}

// This message is sent when a transition animation completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1434985-pagecontrollerdidendlivetransiti?language=objc
func (di *PageControllerDelegate) SetPageControllerDidEndLiveTransition(f func(pageController PageController)) {
	di._PageControllerDidEndLiveTransition = f
}

// This message is sent when a transition animation completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1434985-pagecontrollerdidendlivetransiti?language=objc
func (di *PageControllerDelegate) PageControllerDidEndLiveTransition(pageController PageController) {
	di._PageControllerDidEndLiveTransition(pageController)
}
func (di *PageControllerDelegate) HasPageControllerWillStartLiveTransition() bool {
	return di._PageControllerWillStartLiveTransition != nil
}

// This message is sent when the user begins a transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435009-pagecontrollerwillstartlivetrans?language=objc
func (di *PageControllerDelegate) SetPageControllerWillStartLiveTransition(f func(pageController PageController)) {
	di._PageControllerWillStartLiveTransition = f
}

// This message is sent when the user begins a transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435009-pagecontrollerwillstartlivetrans?language=objc
func (di *PageControllerDelegate) PageControllerWillStartLiveTransition(pageController PageController) {
	di._PageControllerWillStartLiveTransition(pageController)
}
func (di *PageControllerDelegate) HasPageControllerDidTransitionToObject() bool {
	return di._PageControllerDidTransitionToObject != nil
}

// This message is sent when any page transition is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435021-pagecontroller?language=objc
func (di *PageControllerDelegate) SetPageControllerDidTransitionToObject(f func(pageController PageController, object objc.Object)) {
	di._PageControllerDidTransitionToObject = f
}

// This message is sent when any page transition is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435021-pagecontroller?language=objc
func (di *PageControllerDelegate) PageControllerDidTransitionToObject(pageController PageController, object objc.Object) {
	di._PageControllerDidTransitionToObject(pageController, object)
}

// A concrete type wrapper for the [PPageControllerDelegate] protocol.
type PageControllerDelegateWrapper struct {
	objc.Object
}

func (p_ PageControllerDelegateWrapper) HasPageControllerDidEndLiveTransition() bool {
	return p_.RespondsToSelector(objc.Sel("pageControllerDidEndLiveTransition:"))
}

// This message is sent when a transition animation completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1434985-pagecontrollerdidendlivetransiti?language=objc
func (p_ PageControllerDelegateWrapper) PageControllerDidEndLiveTransition(pageController IPageController) {
	objc.Call[objc.Void](p_, objc.Sel("pageControllerDidEndLiveTransition:"), objc.Ptr(pageController))
}

func (p_ PageControllerDelegateWrapper) HasPageControllerWillStartLiveTransition() bool {
	return p_.RespondsToSelector(objc.Sel("pageControllerWillStartLiveTransition:"))
}

// This message is sent when the user begins a transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435009-pagecontrollerwillstartlivetrans?language=objc
func (p_ PageControllerDelegateWrapper) PageControllerWillStartLiveTransition(pageController IPageController) {
	objc.Call[objc.Void](p_, objc.Sel("pageControllerWillStartLiveTransition:"), objc.Ptr(pageController))
}

func (p_ PageControllerDelegateWrapper) HasPageControllerDidTransitionToObject() bool {
	return p_.RespondsToSelector(objc.Sel("pageController:didTransitionToObject:"))
}

// This message is sent when any page transition is completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate/1435021-pagecontroller?language=objc
func (p_ PageControllerDelegateWrapper) PageControllerDidTransitionToObject(pageController IPageController, object objc.IObject) {
	objc.Call[objc.Void](p_, objc.Sel("pageController:didTransitionToObject:"), objc.Ptr(pageController), object)
}
