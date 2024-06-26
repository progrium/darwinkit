// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol that provides a container in which the find bar is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer?language=objc
type PTextFinderBarContainer interface {
	// optional
	FindBarViewDidChangeHeight()
	HasFindBarViewDidChangeHeight() bool

	// optional
	ContentView() View
	HasContentView() bool

	// optional
	SetFindBarView(value View)
	HasSetFindBarView() bool

	// optional
	FindBarView() View
	HasFindBarView() bool

	// optional
	SetFindBarVisible(value bool)
	HasSetFindBarVisible() bool

	// optional
	IsFindBarVisible() bool
	HasIsFindBarVisible() bool
}

// ensure impl type implements protocol interface
var _ PTextFinderBarContainer = (*TextFinderBarContainerObject)(nil)

// A concrete type for the [PTextFinderBarContainer] protocol.
type TextFinderBarContainerObject struct {
	objc.Object
}

func (t_ TextFinderBarContainerObject) HasFindBarViewDidChangeHeight() bool {
	return t_.RespondsToSelector(objc.Sel("findBarViewDidChangeHeight"))
}

// Notifies the find bar container that the find bar has changed its height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1529109-findbarviewdidchangeheight?language=objc
func (t_ TextFinderBarContainerObject) FindBarViewDidChangeHeight() {
	objc.Call[objc.Void](t_, objc.Sel("findBarViewDidChangeHeight"))
}

func (t_ TextFinderBarContainerObject) HasContentView() bool {
	return t_.RespondsToSelector(objc.Sel("contentView"))
}

// A view hierarchy that contains all the views which display the contents being searched. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1532766-contentview?language=objc
func (t_ TextFinderBarContainerObject) ContentView() View {
	rv := objc.Call[View](t_, objc.Sel("contentView"))
	return rv
}

func (t_ TextFinderBarContainerObject) HasSetFindBarView() bool {
	return t_.RespondsToSelector(objc.Sel("setFindBarView:"))
}

// The view assigned by the text bar as the find bar view for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1531692-findbarview?language=objc
func (t_ TextFinderBarContainerObject) SetFindBarView(value View) {
	objc.Call[objc.Void](t_, objc.Sel("setFindBarView:"), value)
}

func (t_ TextFinderBarContainerObject) HasFindBarView() bool {
	return t_.RespondsToSelector(objc.Sel("findBarView"))
}

// The view assigned by the text bar as the find bar view for the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1531692-findbarview?language=objc
func (t_ TextFinderBarContainerObject) FindBarView() View {
	rv := objc.Call[View](t_, objc.Sel("findBarView"))
	return rv
}

func (t_ TextFinderBarContainerObject) HasSetFindBarVisible() bool {
	return t_.RespondsToSelector(objc.Sel("setFindBarVisible:"))
}

// Returns whether the container should display its find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1528587-findbarvisible?language=objc
func (t_ TextFinderBarContainerObject) SetFindBarVisible(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFindBarVisible:"), value)
}

func (t_ TextFinderBarContainerObject) HasIsFindBarVisible() bool {
	return t_.RespondsToSelector(objc.Sel("isFindBarVisible"))
}

// Returns whether the container should display its find bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextfinderbarcontainer/1528587-findbarvisible?language=objc
func (t_ TextFinderBarContainerObject) IsFindBarVisible() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFindBarVisible"))
	return rv
}
