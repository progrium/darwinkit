// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BrowserCell] class.
var BrowserCellClass = _BrowserCellClass{objc.GetClass("NSBrowserCell")}

type _BrowserCellClass struct {
	objc.Class
}

// An interface definition for the [BrowserCell] class.
type IBrowserCell interface {
	ICell
	HighlightColorInView(controlView IView) Color
	Set()
	Reset()
	IsLeaf() bool
	SetLeaf(value bool)
	AlternateImage() Image
	SetAlternateImage(value IImage)
	IsLoaded() bool
	SetLoaded(value bool)
}

// The user interface of a browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell?language=objc
type BrowserCell struct {
	Cell
}

func BrowserCellFrom(ptr unsafe.Pointer) BrowserCell {
	return BrowserCell{
		Cell: CellFrom(ptr),
	}
}

func (b_ BrowserCell) InitImageCell(image IImage) BrowserCell {
	rv := objc.Call[BrowserCell](b_, objc.Sel("initImageCell:"), objc.Ptr(image))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1644593-initimagecell?language=objc
func NewBrowserCellImageCell(image IImage) BrowserCell {
	instance := BrowserCellClass.Alloc().InitImageCell(image)
	instance.Autorelease()
	return instance
}

func (b_ BrowserCell) InitTextCell(string_ string) BrowserCell {
	rv := objc.Call[BrowserCell](b_, objc.Sel("initTextCell:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1644701-inittextcell?language=objc
func NewBrowserCellTextCell(string_ string) BrowserCell {
	instance := BrowserCellClass.Alloc().InitTextCell(string_)
	instance.Autorelease()
	return instance
}

func (bc _BrowserCellClass) Alloc() BrowserCell {
	rv := objc.Call[BrowserCell](bc, objc.Sel("alloc"))
	return rv
}

func BrowserCell_Alloc() BrowserCell {
	return BrowserCellClass.Alloc()
}

func (bc _BrowserCellClass) New() BrowserCell {
	rv := objc.Call[BrowserCell](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBrowserCell() BrowserCell {
	return BrowserCellClass.New()
}

func (b_ BrowserCell) Init() BrowserCell {
	rv := objc.Call[BrowserCell](b_, objc.Sel("init"))
	return rv
}

// Returns the highlight color that the receiver wants to display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435767-highlightcolorinview?language=objc
func (b_ BrowserCell) HighlightColorInView(controlView IView) Color {
	rv := objc.Call[Color](b_, objc.Sel("highlightColorInView:"), objc.Ptr(controlView))
	return rv
}

// Highlights the receiver and sets its state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435770-set?language=objc
func (b_ BrowserCell) Set() {
	objc.Call[objc.Void](b_, objc.Sel("set"))
}

// Unhighlights the receiver and unsets its state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435773-reset?language=objc
func (b_ BrowserCell) Reset() {
	objc.Call[objc.Void](b_, objc.Sel("reset"))
}

// A Boolean that indicates whether the browser cell is a leaf or a branch cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435771-leaf?language=objc
func (b_ BrowserCell) IsLeaf() bool {
	rv := objc.Call[bool](b_, objc.Sel("isLeaf"))
	return rv
}

// A Boolean that indicates whether the browser cell is a leaf or a branch cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435771-leaf?language=objc
func (b_ BrowserCell) SetLeaf(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setLeaf:"), value)
}

// The browser cell’s image for the highlighted state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435768-alternateimage?language=objc
func (b_ BrowserCell) AlternateImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("alternateImage"))
	return rv
}

// The browser cell’s image for the highlighted state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435768-alternateimage?language=objc
func (b_ BrowserCell) SetAlternateImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setAlternateImage:"), objc.Ptr(value))
}

// A Boolean that indicates whether the cell is ready to display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435772-loaded?language=objc
func (b_ BrowserCell) IsLoaded() bool {
	rv := objc.Call[bool](b_, objc.Sel("isLoaded"))
	return rv
}

// A Boolean that indicates whether the cell is ready to display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435772-loaded?language=objc
func (b_ BrowserCell) SetLoaded(value bool) {
	objc.Call[objc.Void](b_, objc.Sel("setLoaded:"), value)
}

// Returns the default image for branch browser cells that are highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435769-highlightedbranchimage?language=objc
func (bc _BrowserCellClass) HighlightedBranchImage() Image {
	rv := objc.Call[Image](bc, objc.Sel("highlightedBranchImage"))
	return rv
}

// Returns the default image for branch browser cells that are highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435769-highlightedbranchimage?language=objc
func BrowserCell_HighlightedBranchImage() Image {
	return BrowserCellClass.HighlightedBranchImage()
}

// Returns the default image for branch cells in a browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435775-branchimage?language=objc
func (bc _BrowserCellClass) BranchImage() Image {
	rv := objc.Call[Image](bc, objc.Sel("branchImage"))
	return rv
}

// Returns the default image for branch cells in a browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbrowsercell/1435775-branchimage?language=objc
func BrowserCell_BranchImage() Image {
	return BrowserCellClass.BranchImage()
}
