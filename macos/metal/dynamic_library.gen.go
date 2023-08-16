// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A dynamically linkable representation of compiled shader code for a specific Metal device object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary?language=objc
type PDynamicLibrary interface {
	// optional
	SerializeToURLError(url foundation.URL, error foundation.Error) bool
	HasSerializeToURLError() bool

	// optional
	InstallName() string
	HasInstallName() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PDynamicLibrary] protocol.
type DynamicLibraryWrapper struct {
	objc.Object
}

func (d_ DynamicLibraryWrapper) HasSerializeToURLError() bool {
	return d_.RespondsToSelector(objc.Sel("serializeToURL:error:"))
}

// Writes the contents of the dynamic library to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553982-serializetourl?language=objc
func (d_ DynamicLibraryWrapper) SerializeToURLError(url foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](d_, objc.Sel("serializeToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

func (d_ DynamicLibraryWrapper) HasInstallName() bool {
	return d_.RespondsToSelector(objc.Sel("installName"))
}

// A file path for this dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553980-installname?language=objc
func (d_ DynamicLibraryWrapper) InstallName() string {
	rv := objc.Call[string](d_, objc.Sel("installName"))
	return rv
}

func (d_ DynamicLibraryWrapper) HasDevice() bool {
	return d_.RespondsToSelector(objc.Sel("device"))
}

// The Metal device object that created the dynamic library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553979-device?language=objc
func (d_ DynamicLibraryWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](d_, objc.Sel("device"))
	return rv
}

func (d_ DynamicLibraryWrapper) HasSetLabel() bool {
	return d_.RespondsToSelector(objc.Sel("setLabel:"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553981-label?language=objc
func (d_ DynamicLibraryWrapper) SetLabel(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLabel:"), value)
}

func (d_ DynamicLibraryWrapper) HasLabel() bool {
	return d_.RespondsToSelector(objc.Sel("label"))
}

// A string that identifies the library. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldynamiclibrary/3553981-label?language=objc
func (d_ DynamicLibraryWrapper) Label() string {
	rv := objc.Call[string](d_, objc.Sel("label"))
	return rv
}
