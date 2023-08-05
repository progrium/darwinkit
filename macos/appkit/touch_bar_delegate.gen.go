// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ITouchBarDelegate interface {
	ImplementsTouchBarMakeItemForIdentifier() bool
	// optional
	TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

type TouchBarDelegate struct {
	_TouchBarMakeItemForIdentifier func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem
}

func (di *TouchBarDelegate) ImplementsTouchBarMakeItemForIdentifier() bool {
	return di._TouchBarMakeItemForIdentifier != nil
}

func (di *TouchBarDelegate) SetTouchBarMakeItemForIdentifier(f func(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem) {
	di._TouchBarMakeItemForIdentifier = f
}

func (di *TouchBarDelegate) TouchBarMakeItemForIdentifier(touchBar TouchBar, identifier TouchBarItemIdentifier) ITouchBarItem {
	return di._TouchBarMakeItemForIdentifier(touchBar, identifier)
}

type TouchBarDelegateWrapper struct {
	objc.Object
}

func (t_ TouchBarDelegateWrapper) ImplementsTouchBarMakeItemForIdentifier() bool {
	return t_.RespondsToSelector(objc.GetSelector("touchBar:makeItemForIdentifier:"))
}

func (t_ TouchBarDelegateWrapper) TouchBarMakeItemForIdentifier(touchBar ITouchBar, identifier TouchBarItemIdentifier) TouchBarItem {
	rv := objc.CallMethod[TouchBarItem](t_, objc.GetSelector("touchBar:makeItemForIdentifier:"), objc.ExtractPtr(touchBar), identifier)
	return rv
}
