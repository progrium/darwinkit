// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IStackViewDelegate interface {
	ImplementsStackViewDidReattachViews() bool
	// optional
	StackViewDidReattachViews(stackView StackView, views []View)
	ImplementsStackViewWillDetachViews() bool
	// optional
	StackViewWillDetachViews(stackView StackView, views []View)
}

type StackViewDelegate struct {
	_StackViewDidReattachViews func(stackView StackView, views []View)
	_StackViewWillDetachViews  func(stackView StackView, views []View)
}

func (di *StackViewDelegate) ImplementsStackViewDidReattachViews() bool {
	return di._StackViewDidReattachViews != nil
}

func (di *StackViewDelegate) SetStackViewDidReattachViews(f func(stackView StackView, views []View)) {
	di._StackViewDidReattachViews = f
}

func (di *StackViewDelegate) StackViewDidReattachViews(stackView StackView, views []View) {
	di._StackViewDidReattachViews(stackView, views)
}
func (di *StackViewDelegate) ImplementsStackViewWillDetachViews() bool {
	return di._StackViewWillDetachViews != nil
}

func (di *StackViewDelegate) SetStackViewWillDetachViews(f func(stackView StackView, views []View)) {
	di._StackViewWillDetachViews = f
}

func (di *StackViewDelegate) StackViewWillDetachViews(stackView StackView, views []View) {
	di._StackViewWillDetachViews(stackView, views)
}

type StackViewDelegateWrapper struct {
	objc.Object
}

func (s_ StackViewDelegateWrapper) ImplementsStackViewDidReattachViews() bool {
	return s_.RespondsToSelector(objc.GetSelector("stackView:didReattachViews:"))
}

func (s_ StackViewDelegateWrapper) StackViewDidReattachViews(stackView IStackView, views []IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stackView:didReattachViews:"), objc.ExtractPtr(stackView), views)
}

func (s_ StackViewDelegateWrapper) ImplementsStackViewWillDetachViews() bool {
	return s_.RespondsToSelector(objc.GetSelector("stackView:willDetachViews:"))
}

func (s_ StackViewDelegateWrapper) StackViewWillDetachViews(stackView IStackView, views []IView) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("stackView:willDetachViews:"), objc.ExtractPtr(stackView), views)
}
