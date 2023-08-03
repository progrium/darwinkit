// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}

type _PasteboardItemClass struct {
	objc.Class
}

type IPasteboardItem interface {
	objc.IObject
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	SetDataProviderForTypes(dataProvider IPasteboardItemDataProvider, types []PasteboardType) bool
	SetDataProvider0ForTypes(dataProvider objc.IObject, types []PasteboardType) bool
	SetDataForType(data []byte, type_ PasteboardType) bool
	SetStringForType(string_ string, type_ PasteboardType) bool
	SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType) bool
	DataForType(type_ PasteboardType) []byte
	StringForType(type_ PasteboardType) string
	PropertyListForType(type_ PasteboardType) objc.Object
	Types() []PasteboardType
}

type PasteboardItem struct {
	objc.Object
}

func MakePasteboardItem(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{
		Object: objc.MakeObject(ptr),
	}
}

func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.CallMethod[PasteboardItem](pc, objc.GetSelector("alloc"))
	return rv
}

func PasteboardItem_Alloc() PasteboardItem {
	return PasteboardItemClass.Alloc()
}

func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.CallMethod[PasteboardItem](pc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewPasteboardItem() PasteboardItem {
	return PasteboardItemClass.New()
}

func PasteboardItem_New() PasteboardItem {
	return PasteboardItemClass.New()
}

func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.CallMethod[PasteboardItem](p_, objc.GetSelector("init"))
	return rv
}

func PasteboardItem_Init() PasteboardItem {
	return PasteboardItemClass.Alloc().Init()
}

func (p_ PasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](p_, objc.GetSelector("availableTypeFromArray:"), types)
	return rv
}

func (p_ PasteboardItem) SetDataProviderForTypes(dataProvider IPasteboardItemDataProvider, types []PasteboardType) bool {
	po := objc.WrapAsProtocol("NSPasteboardItemDataProvider", dataProvider)
	rv := objc.CallMethod[bool](p_, objc.GetSelector("setDataProvider:forTypes:"), po, types)
	return rv
}

func (p_ PasteboardItem) SetDataProvider0ForTypes(dataProvider objc.IObject, types []PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("setDataProvider:forTypes:"), objc.ExtractPtr(dataProvider), types)
	return rv
}

func (p_ PasteboardItem) SetDataForType(data []byte, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("setData:forType:"), data, type_)
	return rv
}

func (p_ PasteboardItem) SetStringForType(string_ string, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("setString:forType:"), string_, type_)
	return rv
}

func (p_ PasteboardItem) SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](p_, objc.GetSelector("setPropertyList:forType:"), objc.ExtractPtr(propertyList), type_)
	return rv
}

func (p_ PasteboardItem) DataForType(type_ PasteboardType) []byte {
	rv := objc.CallMethod[[]byte](p_, objc.GetSelector("dataForType:"), type_)
	return rv
}

func (p_ PasteboardItem) StringForType(type_ PasteboardType) string {
	rv := objc.CallMethod[string](p_, objc.GetSelector("stringForType:"), type_)
	return rv
}

func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType) objc.Object {
	rv := objc.CallMethod[objc.Object](p_, objc.GetSelector("propertyListForType:"), type_)
	return rv
}

func (p_ PasteboardItem) Types() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](p_, objc.GetSelector("types"))
	return rv
}
