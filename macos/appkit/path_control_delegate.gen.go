// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that can be implemented by the delegate of a path control object to support dragging to and from the control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate?language=objc
type PPathControlDelegate interface {
	// optional
	PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
	HasPathControlShouldDragItemWithPasteboard() bool
}

// A delegate implementation builder for the [PPathControlDelegate] protocol.
type PathControlDelegate struct {
	_PathControlShouldDragItemWithPasteboard func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool
}

func (di *PathControlDelegate) HasPathControlShouldDragItemWithPasteboard() bool {
	return di._PathControlShouldDragItemWithPasteboard != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (di *PathControlDelegate) SetPathControlShouldDragItemWithPasteboard(f func(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool) {
	di._PathControlShouldDragItemWithPasteboard = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (di *PathControlDelegate) PathControlShouldDragItemWithPasteboard(pathControl PathControl, pathItem PathControlItem, pasteboard Pasteboard) bool {
	return di._PathControlShouldDragItemWithPasteboard(pathControl, pathItem, pasteboard)
}

// A concrete type wrapper for the [PPathControlDelegate] protocol.
type PathControlDelegateWrapper struct {
	objc.Object
}

func (p_ PathControlDelegateWrapper) HasPathControlShouldDragItemWithPasteboard() bool {
	return p_.RespondsToSelector(objc.Sel("pathControl:shouldDragItem:withPasteboard:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspathcontroldelegate/1526752-pathcontrol?language=objc
func (p_ PathControlDelegateWrapper) PathControlShouldDragItemWithPasteboard(pathControl IPathControl, pathItem IPathControlItem, pasteboard IPasteboard) bool {
	rv := objc.Call[bool](p_, objc.Sel("pathControl:shouldDragItem:withPasteboard:"), objc.Ptr(pathControl), objc.Ptr(pathItem), objc.Ptr(pasteboard))
	return rv
}
