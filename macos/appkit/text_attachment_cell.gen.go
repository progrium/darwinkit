// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextAttachmentCellClass = _TextAttachmentCellClass{objc.GetClass("NSTextAttachmentCell")}

type _TextAttachmentCellClass struct {
	objc.Class
}

type ITextAttachmentCell interface {
	ICell
}

type TextAttachmentCell struct {
	Cell
}

func MakeTextAttachmentCell(ptr unsafe.Pointer) TextAttachmentCell {
	return TextAttachmentCell{
		Cell: MakeCell(ptr),
	}
}

func (t_ TextAttachmentCell) InitImageCell(image IImage) TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.GetSelector("initImageCell:"), objc.ExtractPtr(image))
	return rv
}

func TextAttachmentCell_InitImageCell(image IImage) TextAttachmentCell {
	return TextAttachmentCellClass.Alloc().InitImageCell(image)
}

func (t_ TextAttachmentCell) InitTextCell(string_ string) TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.GetSelector("initTextCell:"), string_)
	return rv
}

func TextAttachmentCell_InitTextCell(string_ string) TextAttachmentCell {
	return TextAttachmentCellClass.Alloc().InitTextCell(string_)
}

func (t_ TextAttachmentCell) Init() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](t_, objc.GetSelector("init"))
	return rv
}

func TextAttachmentCell_Init() TextAttachmentCell {
	return TextAttachmentCellClass.Alloc().Init()
}

func (tc _TextAttachmentCellClass) Alloc() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](tc, objc.GetSelector("alloc"))
	return rv
}

func TextAttachmentCell_Alloc() TextAttachmentCell {
	return TextAttachmentCellClass.Alloc()
}

func (tc _TextAttachmentCellClass) New() TextAttachmentCell {
	rv := objc.CallMethod[TextAttachmentCell](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachmentCell() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}

func TextAttachmentCell_New() TextAttachmentCell {
	return TextAttachmentCellClass.New()
}
