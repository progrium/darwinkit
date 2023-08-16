// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PasteboardItem] class.
var PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	objc.Class
}

// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	objc.IObject
	StringForType(type_ PasteboardType) string
	SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType) bool
	SetStringForType(string_ string, type_ PasteboardType) bool
	DataForType(type_ PasteboardType) []byte
	SetDataForType(data []byte, type_ PasteboardType) bool
	SetDataProviderForTypes(dataProvider PPasteboardItemDataProvider, types []PasteboardType) bool
	SetDataProviderObjectForTypes(dataProviderObject objc.IObject, types []PasteboardType) bool
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	PropertyListForType(type_ PasteboardType) objc.Object
	Types() []PasteboardType
}

// An item on a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem?language=objc
type PasteboardItem struct {
	objc.Object
}

func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.Call[PasteboardItem](pc, objc.Sel("alloc"))
	return rv
}

func PasteboardItem_Alloc() PasteboardItem {
	return PasteboardItemClass.Alloc()
}

func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.Call[PasteboardItem](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPasteboardItem() PasteboardItem {
	return PasteboardItemClass.New()
}

func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.Call[PasteboardItem](p_, objc.Sel("init"))
	return rv
}

// Returns the value for the specified type as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508490-stringfortype?language=objc
func (p_ PasteboardItem) StringForType(type_ PasteboardType) string {
	rv := objc.Call[string](p_, objc.Sel("stringForType:"), type_)
	return rv
}

// Sets the value for a specified type as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508494-setpropertylist?language=objc
func (p_ PasteboardItem) SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setPropertyList:forType:"), propertyList, type_)
	return rv
}

// Sets the value for a specified type as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508498-setstring?language=objc
func (p_ PasteboardItem) SetStringForType(string_ string, type_ PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setString:forType:"), string_, type_)
	return rv
}

// Returns the value for the specified type as a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508496-datafortype?language=objc
func (p_ PasteboardItem) DataForType(type_ PasteboardType) []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("dataForType:"), type_)
	return rv
}

// Sets the value for a specified type as a data object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508501-setdata?language=objc
func (p_ PasteboardItem) SetDataForType(data []byte, type_ PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setData:forType:"), data, type_)
	return rv
}

// Sets the data provider for the specified types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508502-setdataprovider?language=objc
func (p_ PasteboardItem) SetDataProviderForTypes(dataProvider PPasteboardItemDataProvider, types []PasteboardType) bool {
	po0 := objc.WrapAsProtocol("NSPasteboardItemDataProvider", dataProvider)
	rv := objc.Call[bool](p_, objc.Sel("setDataProvider:forTypes:"), po0, types)
	return rv
}

// Sets the data provider for the specified types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508502-setdataprovider?language=objc
func (p_ PasteboardItem) SetDataProviderObjectForTypes(dataProviderObject objc.IObject, types []PasteboardType) bool {
	rv := objc.Call[bool](p_, objc.Sel("setDataProvider:forTypes:"), objc.Ptr(dataProviderObject), types)
	return rv
}

// Returns from a given array of types the first type within the pasteboard item, according to the ordering of types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508488-availabletypefromarray?language=objc
func (p_ PasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.Call[PasteboardType](p_, objc.Sel("availableTypeFromArray:"), types)
	return rv
}

// Returns the value for the specified type as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508489-propertylistfortype?language=objc
func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType) objc.Object {
	rv := objc.Call[objc.Object](p_, objc.Sel("propertyListForType:"), type_)
	return rv
}

// An array of uniform type identifier strings of the data types that the receiver supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboarditem/1508499-types?language=objc
func (p_ PasteboardItem) Types() []PasteboardType {
	rv := objc.Call[[]PasteboardType](p_, objc.Sel("types"))
	return rv
}
