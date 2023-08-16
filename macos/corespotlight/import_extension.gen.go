// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImportExtension] class.
var ImportExtensionClass = _ImportExtensionClass{objc.GetClass("CSImportExtension")}

type _ImportExtensionClass struct {
	objc.Class
}

// An interface definition for the [ImportExtension] class.
type IImportExtension interface {
	objc.IObject
	UpdateAttributesForFileAtURLError(attributes ISearchableItemAttributeSet, contentURL foundation.IURL, error foundation.IError) bool
}

// An object that provides searchable attributes for file types that the app supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csimportextension?language=objc
type ImportExtension struct {
	objc.Object
}

func ImportExtensionFrom(ptr unsafe.Pointer) ImportExtension {
	return ImportExtension{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ic _ImportExtensionClass) Alloc() ImportExtension {
	rv := objc.Call[ImportExtension](ic, objc.Sel("alloc"))
	return rv
}

func ImportExtension_Alloc() ImportExtension {
	return ImportExtensionClass.Alloc()
}

func (ic _ImportExtensionClass) New() ImportExtension {
	rv := objc.Call[ImportExtension](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImportExtension() ImportExtension {
	return ImportExtensionClass.New()
}

func (i_ ImportExtension) Init() ImportExtension {
	rv := objc.Call[ImportExtension](i_, objc.Sel("init"))
	return rv
}

// Provides searchable attributes for a file at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csimportextension/3752010-updateattributes?language=objc
func (i_ ImportExtension) UpdateAttributesForFileAtURLError(attributes ISearchableItemAttributeSet, contentURL foundation.IURL, error foundation.IError) bool {
	rv := objc.Call[bool](i_, objc.Sel("updateAttributes:forFileAtURL:error:"), objc.Ptr(attributes), objc.Ptr(contentURL), objc.Ptr(error))
	return rv
}
