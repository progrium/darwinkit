// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SecureUnarchiveFromDataTransformer] class.
var SecureUnarchiveFromDataTransformerClass = _SecureUnarchiveFromDataTransformerClass{objc.GetClass("NSSecureUnarchiveFromDataTransformer")}

type _SecureUnarchiveFromDataTransformerClass struct {
	objc.Class
}

// An interface definition for the [SecureUnarchiveFromDataTransformer] class.
type ISecureUnarchiveFromDataTransformer interface {
	IValueTransformer
}

// A value transformer that converts data to and from classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssecureunarchivefromdatatransformer?language=objc
type SecureUnarchiveFromDataTransformer struct {
	ValueTransformer
}

func SecureUnarchiveFromDataTransformerFrom(ptr unsafe.Pointer) SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformer{
		ValueTransformer: ValueTransformerFrom(ptr),
	}
}

func (sc _SecureUnarchiveFromDataTransformerClass) Alloc() SecureUnarchiveFromDataTransformer {
	rv := objc.Call[SecureUnarchiveFromDataTransformer](sc, objc.Sel("alloc"))
	return rv
}

func SecureUnarchiveFromDataTransformer_Alloc() SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformerClass.Alloc()
}

func (sc _SecureUnarchiveFromDataTransformerClass) New() SecureUnarchiveFromDataTransformer {
	rv := objc.Call[SecureUnarchiveFromDataTransformer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSecureUnarchiveFromDataTransformer() SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformerClass.New()
}

func (s_ SecureUnarchiveFromDataTransformer) Init() SecureUnarchiveFromDataTransformer {
	rv := objc.Call[SecureUnarchiveFromDataTransformer](s_, objc.Sel("init"))
	return rv
}

// A list of allowed classes the top-level object in an archive must conform to, for encoding and decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssecureunarchivefromdatatransformer/2962887-allowedtoplevelclasses?language=objc
func (sc _SecureUnarchiveFromDataTransformerClass) AllowedTopLevelClasses() []objc.Class {
	rv := objc.Call[[]objc.Class](sc, objc.Sel("allowedTopLevelClasses"))
	return rv
}

// A list of allowed classes the top-level object in an archive must conform to, for encoding and decoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssecureunarchivefromdatatransformer/2962887-allowedtoplevelclasses?language=objc
func SecureUnarchiveFromDataTransformer_AllowedTopLevelClasses() []objc.Class {
	return SecureUnarchiveFromDataTransformerClass.AllowedTopLevelClasses()
}
