// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}

type _TextAttachmentClass struct {
	objc.Class
}

type ITextAttachment interface {
	objc.IObject
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	Contents() []byte
	SetContents(value []byte)
	FileType() string
	SetFileType(value string)
	Image() Image
	SetImage(value IImage)
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.IFileWrapper)
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	UsesTextAttachmentView() bool
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)
}

type TextAttachment struct {
	objc.Object
}

func MakeTextAttachment(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextAttachment) InitWithFileWrapper(fileWrapper foundation.IFileWrapper) TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.GetSelector("initWithFileWrapper:"), objc.ExtractPtr(fileWrapper))
	return rv
}

func TextAttachment_InitWithFileWrapper(fileWrapper foundation.IFileWrapper) TextAttachment {
	return TextAttachmentClass.Alloc().InitWithFileWrapper(fileWrapper)
}

func (t_ TextAttachment) InitWithDataOfType(contentData []byte, uti string) TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.GetSelector("initWithData:ofType:"), contentData, uti)
	return rv
}

func TextAttachment_InitWithDataOfType(contentData []byte, uti string) TextAttachment {
	return TextAttachmentClass.Alloc().InitWithDataOfType(contentData, uti)
}

func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.CallMethod[TextAttachment](tc, objc.GetSelector("alloc"))
	return rv
}

func TextAttachment_Alloc() TextAttachment {
	return TextAttachmentClass.Alloc()
}

func (tc _TextAttachmentClass) New() TextAttachment {
	rv := objc.CallMethod[TextAttachment](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachment() TextAttachment {
	return TextAttachmentClass.New()
}

func TextAttachment_New() TextAttachment {
	return TextAttachmentClass.New()
}

func (t_ TextAttachment) Init() TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.GetSelector("init"))
	return rv
}

func TextAttachment_Init() TextAttachment {
	return TextAttachmentClass.Alloc().Init()
}

func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("registerTextAttachmentViewProviderClass:forFileType:"), objc.ExtractPtr(textAttachmentViewProviderClass), fileType)
}

func TextAttachment_RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	TextAttachmentClass.RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass, fileType)
}

func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	rv := objc.CallMethod[objc.Class](tc, objc.GetSelector("textAttachmentViewProviderClassForFileType:"), fileType)
	return rv
}

func TextAttachment_TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	return TextAttachmentClass.TextAttachmentViewProviderClassForFileType(fileType)
}

func (t_ TextAttachment) Bounds() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.GetSelector("bounds"))
	return rv
}

func (t_ TextAttachment) SetBounds(value coregraphics.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBounds:"), value)
}

func (t_ TextAttachment) Contents() []byte {
	rv := objc.CallMethod[[]byte](t_, objc.GetSelector("contents"))
	return rv
}

func (t_ TextAttachment) SetContents(value []byte) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setContents:"), value)
}

func (t_ TextAttachment) FileType() string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("fileType"))
	return rv
}

func (t_ TextAttachment) SetFileType(value string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFileType:"), value)
}

func (t_ TextAttachment) Image() Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("image"))
	return rv
}

func (t_ TextAttachment) SetImage(value IImage) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setImage:"), objc.ExtractPtr(value))
}

func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := objc.CallMethod[foundation.FileWrapper](t_, objc.GetSelector("fileWrapper"))
	return rv
}

func (t_ TextAttachment) SetFileWrapper(value foundation.IFileWrapper) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setFileWrapper:"), objc.ExtractPtr(value))
}

func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsTextAttachmentView"))
	return rv
}

func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsTextAttachmentView:"), value)
}

func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesTextAttachmentView"))
	return rv
}

func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.CallMethod[float64](t_, objc.GetSelector("lineLayoutPadding"))
	return rv
}

func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLineLayoutPadding:"), value)
}
