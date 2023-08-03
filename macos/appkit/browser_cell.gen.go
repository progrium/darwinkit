// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var BrowserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}

type _BrowserCellClass struct {
	objc.Class
}

type IBrowserCell interface {
	ICell
	Reset()
	Set()
	HighlightColorInView(controlView IView) Color
	AlternateImage() Image
	SetAlternateImage(value IImage)
	IsLeaf() bool
	SetLeaf(value bool)
	IsLoaded() bool
	SetLoaded(value bool)
}

type BrowserCell struct {
	Cell
}

func MakeBrowserCell(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		Cell: MakeCell(ptr),
	}
}

func (b_ BrowserCell) InitImageCell(image IImage) BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func BrowserCell_InitImageCell(image IImage) BrowserCell {
	return BrowserCellClass.Alloc().InitImageCell(image)
}

func (b_ BrowserCell) InitTextCell(string_ string) BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func BrowserCell_InitTextCell(string_ string) BrowserCell {
	return BrowserCellClass.Alloc().InitTextCell(string_)
}

func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.CallMethod[BrowserCell](b_, objc.GetSelector("init"))
	return rv
}

func BrowserCell_Init() BrowserCell {
	return BrowserCellClass.Alloc().Init()
}

func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, objc.GetSelector("alloc"))
	return rv
}

func BrowserCell_Alloc() BrowserCell {
	return BrowserCellClass.Alloc()
}

func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.CallMethod[BrowserCell](bc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

func BrowserCell_New() BrowserCell {
	return BrowserCellClass.New()
}

func (b_ BrowserCell) Reset() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("reset"))
}

func (b_ BrowserCell) Set() {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("set"))
}

func (b_ BrowserCell) HighlightColorInView(controlView IView) Color {
	rv := objc.CallMethod[Color](b_, objc.GetSelector("highlightColorInView:"), objc.ExtractPtr(controlView))
	return rv
}

func (bc _BrowserCellClass) BranchImage() Image {
	rv := objc.CallMethod[Image](bc, objc.GetSelector("branchImage"))
	return rv
}

func BrowserCell_BranchImage() Image {
	return BrowserCellClass.BranchImage()
}

func (bc _BrowserCellClass) HighlightedBranchImage() Image {
	rv := objc.CallMethod[Image](bc, objc.GetSelector("highlightedBranchImage"))
	return rv
}

func BrowserCell_HighlightedBranchImage() Image {
	return BrowserCellClass.HighlightedBranchImage()
}

func (b_ BrowserCell) AlternateImage() Image {
	rv := objc.CallMethod[Image](b_, objc.GetSelector("alternateImage"))
	return rv
}

func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setAlternateImage:"), objc.ExtractPtr(value))
}

func (b_ BrowserCell) IsLeaf() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isLeaf"))
	return rv
}

func (b_ BrowserCell) SetLeaf(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLeaf:"), value)
}

func (b_ BrowserCell) IsLoaded() bool {
	rv := objc.CallMethod[bool](b_, objc.GetSelector("isLoaded"))
	return rv
}

func (b_ BrowserCell) SetLoaded(value bool) {
	objc.CallMethod[objc.Void](b_, objc.GetSelector("setLoaded:"), value)
}
