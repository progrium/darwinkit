// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// Optional methods that delegates implement to respond to layout changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagerdelegate?language=objc
type PTextLayoutManagerDelegate interface {
	// optional
	TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager TextLayoutManager, location TextLocationWrapper, hyphenating bool) bool
	HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool
}

// A delegate implementation builder for the [PTextLayoutManagerDelegate] protocol.
type TextLayoutManagerDelegate struct {
	_TextLayoutManagerShouldBreakLineBeforeLocationHyphenating func(textLayoutManager TextLayoutManager, location TextLocationWrapper, hyphenating bool) bool
}

func (di *TextLayoutManagerDelegate) HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool {
	return di._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating != nil
}

// The method the framework calls to determine the soft line break point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagerdelegate/3810021-textlayoutmanager?language=objc
func (di *TextLayoutManagerDelegate) SetTextLayoutManagerShouldBreakLineBeforeLocationHyphenating(f func(textLayoutManager TextLayoutManager, location TextLocationWrapper, hyphenating bool) bool) {
	di._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating = f
}

// The method the framework calls to determine the soft line break point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagerdelegate/3810021-textlayoutmanager?language=objc
func (di *TextLayoutManagerDelegate) TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager TextLayoutManager, location TextLocationWrapper, hyphenating bool) bool {
	return di._TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager, location, hyphenating)
}

// A concrete type wrapper for the [PTextLayoutManagerDelegate] protocol.
type TextLayoutManagerDelegateWrapper struct {
	objc.Object
}

func (t_ TextLayoutManagerDelegateWrapper) HasTextLayoutManagerShouldBreakLineBeforeLocationHyphenating() bool {
	return t_.RespondsToSelector(objc.Sel("textLayoutManager:shouldBreakLineBeforeLocation:hyphenating:"))
}

// The method the framework calls to determine the soft line break point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanagerdelegate/3810021-textlayoutmanager?language=objc
func (t_ TextLayoutManagerDelegateWrapper) TextLayoutManagerShouldBreakLineBeforeLocationHyphenating(textLayoutManager ITextLayoutManager, location PTextLocation, hyphenating bool) bool {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[bool](t_, objc.Sel("textLayoutManager:shouldBreakLineBeforeLocation:hyphenating:"), objc.Ptr(textLayoutManager), po1, hyphenating)
	return rv
}
