// AUTO-GENERATED CODE, DO NOT MODIFY

package corespotlight

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SearchableItem] class.
var SearchableItemClass = _SearchableItemClass{objc.GetClass("CSSearchableItem")}

type _SearchableItemClass struct {
	objc.Class
}

// An interface definition for the [SearchableItem] class.
type ISearchableItem interface {
	objc.IObject
	AttributeSet() SearchableItemAttributeSet
	SetAttributeSet(value ISearchableItemAttributeSet)
	UniqueIdentifier() string
	SetUniqueIdentifier(value string)
	ExpirationDate() foundation.Date
	SetExpirationDate(value foundation.IDate)
	DomainIdentifier() string
	SetDomainIdentifier(value string)
}

// An item that can be indexed and made available to users when they search on their devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem?language=objc
type SearchableItem struct {
	objc.Object
}

func SearchableItemFrom(ptr unsafe.Pointer) SearchableItem {
	return SearchableItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SearchableItem) InitWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ISearchableItemAttributeSet) SearchableItem {
	rv := objc.Call[SearchableItem](s_, objc.Sel("initWithUniqueIdentifier:domainIdentifier:attributeSet:"), uniqueIdentifier, domainIdentifier, objc.Ptr(attributeSet))
	return rv
}

// Returns a searchable item associated with the specified identifier, domain identifier, and attribute set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621647-initwithuniqueidentifier?language=objc
func SearchableItem_InitWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ISearchableItemAttributeSet) SearchableItem {
	return SearchableItemClass.Alloc().InitWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier, domainIdentifier, attributeSet)
}

func (sc _SearchableItemClass) Alloc() SearchableItem {
	rv := objc.Call[SearchableItem](sc, objc.Sel("alloc"))
	return rv
}

func SearchableItem_Alloc() SearchableItem {
	return SearchableItemClass.Alloc()
}

func (sc _SearchableItemClass) New() SearchableItem {
	rv := objc.Call[SearchableItem](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSearchableItem() SearchableItem {
	return SearchableItemClass.New()
}

func (s_ SearchableItem) Init() SearchableItem {
	rv := objc.Call[SearchableItem](s_, objc.Sel("init"))
	return rv
}

// The set of attributes that contain metadata associated with the item in a CSSearchableItemAttributeSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621649-attributeset?language=objc
func (s_ SearchableItem) AttributeSet() SearchableItemAttributeSet {
	rv := objc.Call[SearchableItemAttributeSet](s_, objc.Sel("attributeSet"))
	return rv
}

// The set of attributes that contain metadata associated with the item in a CSSearchableItemAttributeSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621649-attributeset?language=objc
func (s_ SearchableItem) SetAttributeSet(value ISearchableItemAttributeSet) {
	objc.Call[objc.Void](s_, objc.Sel("setAttributeSet:"), objc.Ptr(value))
}

// The value that uniquely identifies the searchable item within your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621672-uniqueidentifier?language=objc
func (s_ SearchableItem) UniqueIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("uniqueIdentifier"))
	return rv
}

// The value that uniquely identifies the searchable item within your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621672-uniqueidentifier?language=objc
func (s_ SearchableItem) SetUniqueIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setUniqueIdentifier:"), value)
}

// The date after which the searchable item should no longer exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621680-expirationdate?language=objc
func (s_ SearchableItem) ExpirationDate() foundation.Date {
	rv := objc.Call[foundation.Date](s_, objc.Sel("expirationDate"))
	return rv
}

// The date after which the searchable item should no longer exist. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621680-expirationdate?language=objc
func (s_ SearchableItem) SetExpirationDate(value foundation.IDate) {
	objc.Call[objc.Void](s_, objc.Sel("setExpirationDate:"), objc.Ptr(value))
}

// An optional identifier that represents the domain or owner of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621665-domainidentifier?language=objc
func (s_ SearchableItem) DomainIdentifier() string {
	rv := objc.Call[string](s_, objc.Sel("domainIdentifier"))
	return rv
}

// An optional identifier that represents the domain or owner of the item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitem/1621665-domainidentifier?language=objc
func (s_ SearchableItem) SetDomainIdentifier(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDomainIdentifier:"), value)
}
