// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ColorList] class.
var ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}

type _ColorListClass struct {
	objc.Class
}

// An interface definition for the [ColorList] class.
type IColorList interface {
	objc.IObject
	RemoveFile()
	InsertColorKeyAtIndex(color IColor, key ColorName, loc uint)
	RemoveColorWithKey(key ColorName)
	ColorWithKey(key ColorName) Color
	WriteToURLError(url foundation.IURL, errPtr foundation.IError) bool
	SetColorForKey(color IColor, key ColorName)
	IsEditable() bool
	Name() ColorListName
	AllKeys() []ColorName
}

// An ordered list of color objects, identified by keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist?language=objc
type ColorList struct {
	objc.Object
}

func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ColorList) InitWithNameFromFile(name ColorListName, path string) ColorList {
	rv := objc.Call[ColorList](c_, objc.Sel("initWithName:fromFile:"), name, path)
	return rv
}

// Initializes and returns a color list from the specified file, registering it under the specified name if it isn’t in use already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522133-initwithname?language=objc
func NewColorListWithNameFromFile(name ColorListName, path string) ColorList {
	instance := ColorListClass.Alloc().InitWithNameFromFile(name, path)
	instance.Autorelease()
	return instance
}

func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.Call[ColorList](cc, objc.Sel("alloc"))
	return rv
}

func ColorList_Alloc() ColorList {
	return ColorListClass.Alloc()
}

func (cc _ColorListClass) New() ColorList {
	rv := objc.Call[ColorList](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewColorList() ColorList {
	return ColorListClass.New()
}

func (c_ ColorList) Init() ColorList {
	rv := objc.Call[ColorList](c_, objc.Sel("init"))
	return rv
}

// Removes the file from which the list was created, if the file is in a standard search path and owned by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522132-removefile?language=objc
func (c_ ColorList) RemoveFile() {
	objc.Call[objc.Void](c_, objc.Sel("removeFile"))
}

// Inserts the specified color at the specified location in the color list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522137-insertcolor?language=objc
func (c_ ColorList) InsertColorKeyAtIndex(color IColor, key ColorName, loc uint) {
	objc.Call[objc.Void](c_, objc.Sel("insertColor:key:atIndex:"), objc.Ptr(color), key, loc)
}

// Searches the available color lists array and returns the color list with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522126-colorlistnamed?language=objc
func (cc _ColorListClass) ColorListNamed(name ColorListName) ColorList {
	rv := objc.Call[ColorList](cc, objc.Sel("colorListNamed:"), name)
	return rv
}

// Searches the available color lists array and returns the color list with the specified name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522126-colorlistnamed?language=objc
func ColorList_ColorListNamed(name ColorListName) ColorList {
	return ColorListClass.ColorListNamed(name)
}

// Removes the color associated with the specified key from the color list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522123-removecolorwithkey?language=objc
func (c_ ColorList) RemoveColorWithKey(key ColorName) {
	objc.Call[objc.Void](c_, objc.Sel("removeColorWithKey:"), key)
}

// Returns the color object associated with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522143-colorwithkey?language=objc
func (c_ ColorList) ColorWithKey(key ColorName) Color {
	rv := objc.Call[Color](c_, objc.Sel("colorWithKey:"), key)
	return rv
}

// Saves the color list to the file at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/2269695-writetourl?language=objc
func (c_ ColorList) WriteToURLError(url foundation.IURL, errPtr foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("writeToURL:error:"), objc.Ptr(url), objc.Ptr(errPtr))
	return rv
}

// Associates the specified color object with the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522130-setcolor?language=objc
func (c_ ColorList) SetColorForKey(color IColor, key ColorName) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:forKey:"), objc.Ptr(color), key)
}

// Returns an array of all color lists found in the standard color list directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522127-availablecolorlists?language=objc
func (cc _ColorListClass) AvailableColorLists() []ColorList {
	rv := objc.Call[[]ColorList](cc, objc.Sel("availableColorLists"))
	return rv
}

// Returns an array of all color lists found in the standard color list directories. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522127-availablecolorlists?language=objc
func ColorList_AvailableColorLists() []ColorList {
	return ColorListClass.AvailableColorLists()
}

// A Boolean value that indicates whether the color list can be modified. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522125-editable?language=objc
func (c_ ColorList) IsEditable() bool {
	rv := objc.Call[bool](c_, objc.Sel("isEditable"))
	return rv
}

// The name of the color list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522138-name?language=objc
func (c_ ColorList) Name() ColorListName {
	rv := objc.Call[ColorListName](c_, objc.Sel("name"))
	return rv
}

// An array of the keys by which the color objects are stored in the color list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorlist/1522141-allkeys?language=objc
func (c_ ColorList) AllKeys() []ColorName {
	rv := objc.Call[[]ColorName](c_, objc.Sel("allKeys"))
	return rv
}
