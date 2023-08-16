// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that defines the interface for initializing an object from a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreading?language=objc
type PPasteboardReading interface {
	// optional
	InitWithPasteboardPropertyListOfType(propertyList objc.Object, type_ PasteboardType) objc.IObject
	HasInitWithPasteboardPropertyListOfType() bool
}

// A concrete type wrapper for the [PPasteboardReading] protocol.
type PasteboardReadingWrapper struct {
	objc.Object
}

func (p_ PasteboardReadingWrapper) HasInitWithPasteboardPropertyListOfType() bool {
	return p_.RespondsToSelector(objc.Sel("initWithPasteboardPropertyList:ofType:"))
}

// Initializes an instance with a property list object and a type string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboardreading/1528252-initwithpasteboardpropertylist?language=objc
func (p_ PasteboardReadingWrapper) InitWithPasteboardPropertyListOfType(propertyList objc.IObject, type_ PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, type_)
	return rv
}
