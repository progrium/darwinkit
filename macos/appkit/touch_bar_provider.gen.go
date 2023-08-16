// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that an object adopts to create a bar object in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider?language=objc
type PTouchBarProvider interface {
	// optional
	TouchBar() ITouchBar
	HasTouchBar() bool
}

// A concrete type wrapper for the [PTouchBarProvider] protocol.
type TouchBarProviderWrapper struct {
	objc.Object
}

func (t_ TouchBarProviderWrapper) HasTouchBar() bool {
	return t_.RespondsToSelector(objc.Sel("touchBar"))
}

// The property you implement to provide a Touch Bar object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider/2544662-touchbar?language=objc
func (t_ TouchBarProviderWrapper) TouchBar() TouchBar {
	rv := objc.Call[TouchBar](t_, objc.Sel("touchBar"))
	return rv
}
