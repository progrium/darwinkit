// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CustomAttributeKey] class.
var CustomAttributeKeyClass = _CustomAttributeKeyClass{objc.GetClass("CSCustomAttributeKey")}

type _CustomAttributeKeyClass struct {
	objc.Class
}

// An interface definition for the [CustomAttributeKey] class.
type ICustomAttributeKey interface {
	objc.IObject
	IsUnique() bool
	IsSearchable() bool
	IsMultiValued() bool
	KeyName() string
	IsSearchableByDefault() bool
}

// A key associated with a custom attribute for a searchable item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey?language=objc
type CustomAttributeKey struct {
	objc.Object
}

func CustomAttributeKeyFrom(ptr unsafe.Pointer) CustomAttributeKey {
	return CustomAttributeKey{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CustomAttributeKey) InitWithKeyName(keyName string) CustomAttributeKey {
	rv := objc.Call[CustomAttributeKey](c_, objc.Sel("initWithKeyName:"), keyName)
	return rv
}

// Returns a new custom attribute key with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616402-initwithkeyname?language=objc
func NewCustomAttributeKeyWithKeyName(keyName string) CustomAttributeKey {
	instance := CustomAttributeKeyClass.Alloc().InitWithKeyName(keyName)
	instance.Autorelease()
	return instance
}

func (cc _CustomAttributeKeyClass) Alloc() CustomAttributeKey {
	rv := objc.Call[CustomAttributeKey](cc, objc.Sel("alloc"))
	return rv
}

func CustomAttributeKey_Alloc() CustomAttributeKey {
	return CustomAttributeKeyClass.Alloc()
}

func (cc _CustomAttributeKeyClass) New() CustomAttributeKey {
	rv := objc.Call[CustomAttributeKey](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCustomAttributeKey() CustomAttributeKey {
	return CustomAttributeKeyClass.New()
}

func (c_ CustomAttributeKey) Init() CustomAttributeKey {
	rv := objc.Call[CustomAttributeKey](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates if duplicate custom attribute values should be treated as the same value to save storage space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616409-unique?language=objc
func (c_ CustomAttributeKey) IsUnique() bool {
	rv := objc.Call[bool](c_, objc.Sel("isUnique"))
	return rv
}

// A Boolean value that indicates if the custom attribute can be specified as a search term. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616397-searchable?language=objc
func (c_ CustomAttributeKey) IsSearchable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSearchable"))
	return rv
}

// A Boolean value that indicates if the custom attribute is likely to have multiple values, such as arrays, associated with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616400-multivalued?language=objc
func (c_ CustomAttributeKey) IsMultiValued() bool {
	rv := objc.Call[bool](c_, objc.Sel("isMultiValued"))
	return rv
}

// The name of the custom attribute key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616405-keyname?language=objc
func (c_ CustomAttributeKey) KeyName() string {
	rv := objc.Call[string](c_, objc.Sel("keyName"))
	return rv
}

// A Boolean value that indicates if the custom attribute should be searchable by default. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cscustomattributekey/1616396-searchablebydefault?language=objc
func (c_ CustomAttributeKey) IsSearchableByDefault() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSearchableByDefault"))
	return rv
}
