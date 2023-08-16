// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that provides a container in which the find bar is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer?language=objc
type PTextFinderBarContainer interface {
	// optional
	ContentView() IView
	HasContentView() bool

	// optional
	FindBarViewDidChangeHeight()
	HasFindBarViewDidChangeHeight() bool

	// optional
	SetFindBarVisible(value bool)
	HasSetFindBarVisible() bool

	// optional
	IsFindBarVisible() bool
	HasIsFindBarVisible() bool

	// optional
	SetFindBarView(value View)
	HasSetFindBarView() bool

	// optional
	FindBarView() IView
	HasFindBarView() bool
}

// A concrete type wrapper for the [PTextFinderBarContainer] protocol.
type TextFinderBarContainerWrapper struct {
	objc.Object
}

func (t_ TextFinderBarContainerWrapper) HasContentView() bool {
	return t_.RespondsToSelector(objc.Sel("contentView"))
}

// A view hierarchy that contains all the views which display the contents being searched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1532766-contentview?language=objc
func (t_ TextFinderBarContainerWrapper) ContentView() View {
	rv := objc.Call[View](t_, objc.Sel("contentView"))
	return rv
}

func (t_ TextFinderBarContainerWrapper) HasFindBarViewDidChangeHeight() bool {
	return t_.RespondsToSelector(objc.Sel("findBarViewDidChangeHeight"))
}

// Notifies the find bar container that the find bar has changed its height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1529109-findbarviewdidchangeheight?language=objc
func (t_ TextFinderBarContainerWrapper) FindBarViewDidChangeHeight() {
	objc.Call[objc.Void](t_, objc.Sel("findBarViewDidChangeHeight"))
}

func (t_ TextFinderBarContainerWrapper) HasSetFindBarVisible() bool {
	return t_.RespondsToSelector(objc.Sel("setFindBarVisible:"))
}

// Returns whether the container should display its find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1528587-findbarvisible?language=objc
func (t_ TextFinderBarContainerWrapper) SetFindBarVisible(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFindBarVisible:"), value)
}

func (t_ TextFinderBarContainerWrapper) HasIsFindBarVisible() bool {
	return t_.RespondsToSelector(objc.Sel("isFindBarVisible"))
}

// Returns whether the container should display its find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1528587-findbarvisible?language=objc
func (t_ TextFinderBarContainerWrapper) IsFindBarVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFindBarVisible"))
	return rv
}

func (t_ TextFinderBarContainerWrapper) HasSetFindBarView() bool {
	return t_.RespondsToSelector(objc.Sel("setFindBarView:"))
}

// The view assigned by the text bar as the find bar view for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1531692-findbarview?language=objc
func (t_ TextFinderBarContainerWrapper) SetFindBarView(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setFindBarView:"), objc.Ptr(value))
}

func (t_ TextFinderBarContainerWrapper) HasFindBarView() bool {
	return t_.RespondsToSelector(objc.Sel("findBarView"))
}

// The view assigned by the text bar as the find bar view for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1531692-findbarview?language=objc
func (t_ TextFinderBarContainerWrapper) FindBarView() View {
	rv := objc.Call[View](t_, objc.Sel("findBarView"))
	return rv
}
