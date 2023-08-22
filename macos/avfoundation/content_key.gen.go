// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentKey] class.
var ContentKeyClass = _ContentKeyClass{objc.GetClass("AVContentKey")}

type _ContentKeyClass struct {
	objc.Class
}

// An interface definition for the [ContentKey] class.
type IContentKey interface {
	objc.IObject
	ContentKeySpecifier() ContentKeySpecifier
}

// An object that represents the content key decryptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkey?language=objc
type ContentKey struct {
	objc.Object
}

func ContentKeyFrom(ptr unsafe.Pointer) ContentKey {
	return ContentKey{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContentKeyClass) Alloc() ContentKey {
	rv := objc.Call[ContentKey](cc, objc.Sel("alloc"))
	return rv
}

func ContentKey_Alloc() ContentKey {
	return ContentKeyClass.Alloc()
}

func (cc _ContentKeyClass) New() ContentKey {
	rv := objc.Call[ContentKey](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentKey() ContentKey {
	return ContentKeyClass.New()
}

func (c_ ContentKey) Init() ContentKey {
	rv := objc.Call[ContentKey](c_, objc.Sel("init"))
	return rv
}

// The content key’s unique identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkey/3726100-contentkeyspecifier?language=objc
func (c_ ContentKey) ContentKeySpecifier() ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](c_, objc.Sel("contentKeySpecifier"))
	return rv
}
