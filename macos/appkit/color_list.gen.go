// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}

type _ColorListClass struct {
	objc.Class
}

type IColorList interface {
	objc.IObject
	ColorWithKey(key ColorName) Color
	InsertColorKeyAtIndex(color IColor, key ColorName, loc uint)
	RemoveColorWithKey(key ColorName)
	SetColorForKey(color IColor, key ColorName)
	WriteToURLError(url foundation.IURL, errPtr *foundation.Error) bool
	RemoveFile()
	Name() ColorListName
	IsEditable() bool
	AllKeys() []ColorName
}

type ColorList struct {
	objc.Object
}

func MakeColorList(ptr unsafe.Pointer) ColorList {
	return ColorList{
		Object: objc.MakeObject(ptr),
	}
}

func (c_ ColorList) InitWithName(name ColorListName) ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.GetSelector("initWithName:"), name)
	return rv
}

func ColorList_InitWithName(name ColorListName) ColorList {
	return ColorListClass.Alloc().InitWithName(name)
}

func (c_ ColorList) InitWithNameFromFile(name ColorListName, path string) ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.GetSelector("initWithName:fromFile:"), name, path)
	return rv
}

func ColorList_InitWithNameFromFile(name ColorListName, path string) ColorList {
	return ColorListClass.Alloc().InitWithNameFromFile(name, path)
}

func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.GetSelector("alloc"))
	return rv
}

func ColorList_Alloc() ColorList {
	return ColorListClass.Alloc()
}

func (cc _ColorListClass) New() ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewColorList() ColorList {
	return ColorListClass.New()
}

func ColorList_New() ColorList {
	return ColorListClass.New()
}

func (c_ ColorList) Init() ColorList {
	rv := objc.CallMethod[ColorList](c_, objc.GetSelector("init"))
	return rv
}

func ColorList_Init() ColorList {
	return ColorListClass.Alloc().Init()
}

func (cc _ColorListClass) ColorListNamed(name ColorListName) ColorList {
	rv := objc.CallMethod[ColorList](cc, objc.GetSelector("colorListNamed:"), name)
	return rv
}

func ColorList_ColorListNamed(name ColorListName) ColorList {
	return ColorListClass.ColorListNamed(name)
}

func (c_ ColorList) ColorWithKey(key ColorName) Color {
	rv := objc.CallMethod[Color](c_, objc.GetSelector("colorWithKey:"), key)
	return rv
}

func (c_ ColorList) InsertColorKeyAtIndex(color IColor, key ColorName, loc uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("insertColor:key:atIndex:"), objc.ExtractPtr(color), key, loc)
}

func (c_ ColorList) RemoveColorWithKey(key ColorName) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("removeColorWithKey:"), key)
}

func (c_ ColorList) SetColorForKey(color IColor, key ColorName) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setColor:forKey:"), objc.ExtractPtr(color), key)
}

func (c_ ColorList) WriteToURLError(url foundation.IURL, errPtr *foundation.Error) bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("writeToURL:error:"), objc.ExtractPtr(url), unsafe.Pointer(errPtr))
	return rv
}

func (c_ ColorList) RemoveFile() {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("removeFile"))
}

func (cc _ColorListClass) AvailableColorLists() []ColorList {
	rv := objc.CallMethod[[]ColorList](cc, objc.GetSelector("availableColorLists"))
	return rv
}

func ColorList_AvailableColorLists() []ColorList {
	return ColorListClass.AvailableColorLists()
}

func (c_ ColorList) Name() ColorListName {
	rv := objc.CallMethod[ColorListName](c_, objc.GetSelector("name"))
	return rv
}

func (c_ ColorList) IsEditable() bool {
	rv := objc.CallMethod[bool](c_, objc.GetSelector("isEditable"))
	return rv
}

func (c_ ColorList) AllKeys() []ColorName {
	rv := objc.CallMethod[[]ColorName](c_, objc.GetSelector("allKeys"))
	return rv
}
