// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var CollectionViewGridLayoutClass = _CollectionViewGridLayoutClass{objc.GetClass("NSCollectionViewGridLayout")}

type _CollectionViewGridLayoutClass struct {
	objc.Class
}

type ICollectionViewGridLayout interface {
	ICollectionViewLayout
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	MinimumItemSize() foundation.Size
	SetMinimumItemSize(value foundation.Size)
	MaximumItemSize() foundation.Size
	SetMaximumItemSize(value foundation.Size)
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	Margins() foundation.EdgeInsets
	SetMargins(value foundation.EdgeInsets)
	BackgroundColors() []Color
	SetBackgroundColors(value []IColor)
}

type CollectionViewGridLayout struct {
	CollectionViewLayout
}

func MakeCollectionViewGridLayout(ptr unsafe.Pointer) CollectionViewGridLayout {
	return CollectionViewGridLayout{
		CollectionViewLayout: MakeCollectionViewLayout(ptr),
	}
}

func (cc _CollectionViewGridLayoutClass) Alloc() CollectionViewGridLayout {
	rv := objc.CallMethod[CollectionViewGridLayout](cc, objc.GetSelector("alloc"))
	return rv
}

func CollectionViewGridLayout_Alloc() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.Alloc()
}

func (cc _CollectionViewGridLayoutClass) New() CollectionViewGridLayout {
	rv := objc.CallMethod[CollectionViewGridLayout](cc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewGridLayout() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.New()
}

func CollectionViewGridLayout_New() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.New()
}

func (c_ CollectionViewGridLayout) Init() CollectionViewGridLayout {
	rv := objc.CallMethod[CollectionViewGridLayout](c_, objc.GetSelector("init"))
	return rv
}

func CollectionViewGridLayout_Init() CollectionViewGridLayout {
	return CollectionViewGridLayoutClass.Alloc().Init()
}

func (c_ CollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("maximumNumberOfRows"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaximumNumberOfRows:"), value)
}

func (c_ CollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := objc.CallMethod[uint](c_, objc.GetSelector("maximumNumberOfColumns"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaximumNumberOfColumns:"), value)
}

func (c_ CollectionViewGridLayout) MinimumItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("minimumItemSize"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumItemSize:"), value)
}

func (c_ CollectionViewGridLayout) MaximumItemSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](c_, objc.GetSelector("maximumItemSize"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMaximumItemSize(value foundation.Size) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMaximumItemSize:"), value)
}

func (c_ CollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("minimumInteritemSpacing"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumInteritemSpacing:"), value)
}

func (c_ CollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := objc.CallMethod[float64](c_, objc.GetSelector("minimumLineSpacing"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMinimumLineSpacing:"), value)
}

func (c_ CollectionViewGridLayout) Margins() foundation.EdgeInsets {
	rv := objc.CallMethod[foundation.EdgeInsets](c_, objc.GetSelector("margins"))
	return rv
}

func (c_ CollectionViewGridLayout) SetMargins(value foundation.EdgeInsets) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setMargins:"), value)
}

func (c_ CollectionViewGridLayout) BackgroundColors() []Color {
	rv := objc.CallMethod[[]Color](c_, objc.GetSelector("backgroundColors"))
	// TODO: convert slice items...
	return rv
}

func (c_ CollectionViewGridLayout) SetBackgroundColors(value []IColor) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("setBackgroundColors:"), value)
}
