// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The interface for loading Core Image image units. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipluginregistration?language=objc
type PPlugInRegistration interface {
	// optional
	Load(host unsafe.Pointer) bool
	HasLoad() bool
}

// A concrete type wrapper for the [PPlugInRegistration] protocol.
type PlugInRegistrationWrapper struct {
	objc.Object
}

func (p_ PlugInRegistrationWrapper) HasLoad() bool {
	return p_.RespondsToSelector(objc.Sel("load:"))
}

// Loads and initializes an image unit, performing custom tasks as needed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipluginregistration/1437823-load?language=objc
func (p_ PlugInRegistrationWrapper) Load(host unsafe.Pointer) bool {
	rv := objc.Call[bool](p_, objc.Sel("load:"), host)
	return rv
}
