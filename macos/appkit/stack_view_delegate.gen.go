// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods you use to respond to a stack view detaching and reattaching views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewdelegate?language=objc
type PStackViewDelegate interface {
	// optional
	StackViewDidReattachViews(stackView StackView, views []View)
	HasStackViewDidReattachViews() bool
}

// A delegate implementation builder for the [PStackViewDelegate] protocol.
type StackViewDelegate struct {
	_StackViewDidReattachViews func(stackView StackView, views []View)
}

func (di *StackViewDelegate) HasStackViewDidReattachViews() bool {
	return di._StackViewDidReattachViews != nil
}

// Called when the stack view has automatically reattached one or more previously-detached views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewdelegate/1488921-stackview?language=objc
func (di *StackViewDelegate) SetStackViewDidReattachViews(f func(stackView StackView, views []View)) {
	di._StackViewDidReattachViews = f
}

// Called when the stack view has automatically reattached one or more previously-detached views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewdelegate/1488921-stackview?language=objc
func (di *StackViewDelegate) StackViewDidReattachViews(stackView StackView, views []View) {
	di._StackViewDidReattachViews(stackView, views)
}

// A concrete type wrapper for the [PStackViewDelegate] protocol.
type StackViewDelegateWrapper struct {
	objc.Object
}

func (s_ StackViewDelegateWrapper) HasStackViewDidReattachViews() bool {
	return s_.RespondsToSelector(objc.Sel("stackView:didReattachViews:"))
}

// Called when the stack view has automatically reattached one or more previously-detached views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstackviewdelegate/1488921-stackview?language=objc
func (s_ StackViewDelegateWrapper) StackViewDidReattachViews(stackView IStackView, views []IView) {
	objc.Call[objc.Void](s_, objc.Sel("stackView:didReattachViews:"), objc.Ptr(stackView), views)
}
