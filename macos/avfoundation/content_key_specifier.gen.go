// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContentKeySpecifier] class.
var ContentKeySpecifierClass = _ContentKeySpecifierClass{objc.GetClass("AVContentKeySpecifier")}

type _ContentKeySpecifierClass struct {
	objc.Class
}

// An interface definition for the [ContentKeySpecifier] class.
type IContentKeySpecifier interface {
	objc.IObject
	KeySystem() ContentKeySystem
	Options() map[string]objc.Object
	Identifier() objc.Object
}

// An object that uniquely identifies a content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier?language=objc
type ContentKeySpecifier struct {
	objc.Object
}

func ContentKeySpecifierFrom(ptr unsafe.Pointer) ContentKeySpecifier {
	return ContentKeySpecifier{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ContentKeySpecifier) InitForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objc.IObject, options map[string]objc.IObject) ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](c_, objc.Sel("initForKeySystem:identifier:options:"), keySystem, contentKeyIdentifier, options)
	return rv
}

// Creates a content key specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/3726107-initforkeysystem?language=objc
func NewContentKeySpecifierForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objc.IObject, options map[string]objc.IObject) ContentKeySpecifier {
	instance := ContentKeySpecifierClass.Alloc().InitForKeySystemIdentifierOptions(keySystem, contentKeyIdentifier, options)
	instance.Autorelease()
	return instance
}

func (cc _ContentKeySpecifierClass) ContentKeySpecifierForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objc.IObject, options map[string]objc.IObject) ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](cc, objc.Sel("contentKeySpecifierForKeySystem:identifier:options:"), keySystem, contentKeyIdentifier, options)
	return rv
}

// A convenience initializer to create a content key specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/3726105-contentkeyspecifierforkeysystem?language=objc
func ContentKeySpecifier_ContentKeySpecifierForKeySystemIdentifierOptions(keySystem ContentKeySystem, contentKeyIdentifier objc.IObject, options map[string]objc.IObject) ContentKeySpecifier {
	return ContentKeySpecifierClass.ContentKeySpecifierForKeySystemIdentifierOptions(keySystem, contentKeyIdentifier, options)
}

func (cc _ContentKeySpecifierClass) Alloc() ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](cc, objc.Sel("alloc"))
	return rv
}

func ContentKeySpecifier_Alloc() ContentKeySpecifier {
	return ContentKeySpecifierClass.Alloc()
}

func (cc _ContentKeySpecifierClass) New() ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContentKeySpecifier() ContentKeySpecifier {
	return ContentKeySpecifierClass.New()
}

func (c_ ContentKeySpecifier) Init() ContentKeySpecifier {
	rv := objc.Call[ContentKeySpecifier](c_, objc.Sel("init"))
	return rv
}

// The key system that generates content keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/3726108-keysystem?language=objc
func (c_ ContentKeySpecifier) KeySystem() ContentKeySystem {
	rv := objc.Call[ContentKeySystem](c_, objc.Sel("keySystem"))
	return rv
}

// A dictionary of options with which you initialized the specifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/3726109-options?language=objc
func (c_ ContentKeySpecifier) Options() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("options"))
	return rv
}

// The container and protocol-specific key identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontentkeyspecifier/3726106-identifier?language=objc
func (c_ ContentKeySpecifier) Identifier() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("identifier"))
	return rv
}
