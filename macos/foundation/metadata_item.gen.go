// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataItem] class.
var MetadataItemClass = _MetadataItemClass{objc.GetClass("NSMetadataItem")}

type _MetadataItemClass struct {
	objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objc.IObject
	ValuesForAttributes(keys []string) map[string]objc.Object
	ValueForAttribute(key string) objc.Object
	Attributes() []string
}

// The metadata associated with a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem?language=objc
type MetadataItem struct {
	objc.Object
}

func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MetadataItem) InitWithURL(url IURL) MetadataItem {
	rv := objc.Call[MetadataItem](m_, objc.Sel("initWithURL:"), objc.Ptr(url))
	return rv
}

// Initializes a metadata item with a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/1414919-initwithurl?language=objc
func NewMetadataItemWithURL(url IURL) MetadataItem {
	instance := MetadataItemClass.Alloc().InitWithURL(url)
	instance.Autorelease()
	return instance
}

func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Call[MetadataItem](mc, objc.Sel("alloc"))
	return rv
}

func MetadataItem_Alloc() MetadataItem {
	return MetadataItemClass.Alloc()
}

func (mc _MetadataItemClass) New() MetadataItem {
	rv := objc.Call[MetadataItem](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataItem() MetadataItem {
	return MetadataItemClass.New()
}

func (m_ MetadataItem) Init() MetadataItem {
	rv := objc.Call[MetadataItem](m_, objc.Sel("init"))
	return rv
}

// Returns a dictionary containing the key-value pairs for the attribute names specified by a given array of keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/1409934-valuesforattributes?language=objc
func (m_ MetadataItem) ValuesForAttributes(keys []string) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](m_, objc.Sel("valuesForAttributes:"), keys)
	return rv
}

// Returns the receiver’s metadata attribute name specified by a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/1411721-valueforattribute?language=objc
func (m_ MetadataItem) ValueForAttribute(key string) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("valueForAttribute:"), key)
	return rv
}

// An array containing the attribute keys for the metadata item’s values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/1418347-attributes?language=objc
func (m_ MetadataItem) Attributes() []string {
	rv := objc.Call[[]string](m_, objc.Sel("attributes"))
	return rv
}
