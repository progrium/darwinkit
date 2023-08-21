// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextAttachment] class.
var TextAttachmentClass = _TextAttachmentClass{objc.GetClass("NSTextAttachment")}

type _TextAttachmentClass struct {
	objc.Class
}

// An interface definition for the [TextAttachment] class.
type ITextAttachment interface {
	objc.IObject
	FileWrapper() foundation.FileWrapper
	SetFileWrapper(value foundation.IFileWrapper)
	LineLayoutPadding() float64
	SetLineLayoutPadding(value float64)
	Bounds() coregraphics.Rect
	SetBounds(value coregraphics.Rect)
	AllowsTextAttachmentView() bool
	SetAllowsTextAttachmentView(value bool)
	UsesTextAttachmentView() bool
	FileType() string
	SetFileType(value string)
	Contents() []byte
	SetContents(value []byte)
	AttachmentCell() objc.Object
	SetAttachmentCell(value objc.IObject)
	Image() Image
	SetImage(value IImage)
}

// The values for the attachment characteristics of attributed strings and related objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment?language=objc
type TextAttachment struct {
	objc.Object
}

func TextAttachmentFrom(ptr unsafe.Pointer) TextAttachment {
	return TextAttachment{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextAttachment) InitWithDataOfType(contentData []byte, uti string) TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("initWithData:ofType:"), contentData, uti)
	return rv
}

// Creates a text attachment object with the specified data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508374-initwithdata?language=objc
func NewTextAttachmentWithDataOfType(contentData []byte, uti string) TextAttachment {
	instance := TextAttachmentClass.Alloc().InitWithDataOfType(contentData, uti)
	instance.Autorelease()
	return instance
}

func (t_ TextAttachment) InitWithFileWrapper(fileWrapper foundation.IFileWrapper) TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("initWithFileWrapper:"), objc.Ptr(fileWrapper))
	return rv
}

// Creates a text attachment object to contain the specified file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/1508373-initwithfilewrapper?language=objc
func NewTextAttachmentWithFileWrapper(fileWrapper foundation.IFileWrapper) TextAttachment {
	instance := TextAttachmentClass.Alloc().InitWithFileWrapper(fileWrapper)
	instance.Autorelease()
	return instance
}

func (tc _TextAttachmentClass) Alloc() TextAttachment {
	rv := objc.Call[TextAttachment](tc, objc.Sel("alloc"))
	return rv
}

func TextAttachment_Alloc() TextAttachment {
	return TextAttachmentClass.Alloc()
}

func (tc _TextAttachmentClass) New() TextAttachment {
	rv := objc.Call[TextAttachment](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachment() TextAttachment {
	return TextAttachmentClass.New()
}

func (t_ TextAttachment) Init() TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("init"))
	return rv
}

// Returns the text attachment view provider class, if any, for the file type you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857587-textattachmentviewproviderclassf?language=objc
func (tc _TextAttachmentClass) TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	rv := objc.Call[objc.Class](tc, objc.Sel("textAttachmentViewProviderClassForFileType:"), fileType)
	return rv
}

// Returns the text attachment view provider class, if any, for the file type you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857587-textattachmentviewproviderclassf?language=objc
func TextAttachment_TextAttachmentViewProviderClassForFileType(fileType string) objc.Class {
	return TextAttachmentClass.TextAttachmentViewProviderClassForFileType(fileType)
}

// Registers a specific file type with the attachment view provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857586-registertextattachmentviewprovid?language=objc
func (tc _TextAttachmentClass) RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	objc.Call[objc.Void](tc, objc.Sel("registerTextAttachmentViewProviderClass:forFileType:"), objc.Ptr(textAttachmentViewProviderClass), fileType)
}

// Registers a specific file type with the attachment view provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857586-registertextattachmentviewprovid?language=objc
func TextAttachment_RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass objc.IClass, fileType string) {
	TextAttachmentClass.RegisterTextAttachmentViewProviderClassForFileType(textAttachmentViewProviderClass, fileType)
}

// The text attachment’s file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508398-filewrapper?language=objc
func (t_ TextAttachment) FileWrapper() foundation.FileWrapper {
	rv := objc.Call[foundation.FileWrapper](t_, objc.Sel("fileWrapper"))
	return rv
}

// The text attachment’s file wrapper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508398-filewrapper?language=objc
func (t_ TextAttachment) SetFileWrapper(value foundation.IFileWrapper) {
	objc.Call[objc.Void](t_, objc.Sel("setFileWrapper:"), objc.Ptr(value))
}

// The layout padding before and after the text attachment bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857585-linelayoutpadding?language=objc
func (t_ TextAttachment) LineLayoutPadding() float64 {
	rv := objc.Call[float64](t_, objc.Sel("lineLayoutPadding"))
	return rv
}

// The layout padding before and after the text attachment bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857585-linelayoutpadding?language=objc
func (t_ TextAttachment) SetLineLayoutPadding(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setLineLayoutPadding:"), value)
}

// The layout bounds of the text attachment’s graphical representation in the text coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508394-bounds?language=objc
func (t_ TextAttachment) Bounds() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("bounds"))
	return rv
}

// The layout bounds of the text attachment’s graphical representation in the text coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508394-bounds?language=objc
func (t_ TextAttachment) SetBounds(value coregraphics.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("setBounds:"), value)
}

// A Boolean value that determines whether the text attachment uses text attachment views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857584-allowstextattachmentview?language=objc
func (t_ TextAttachment) AllowsTextAttachmentView() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsTextAttachmentView"))
	return rv
}

// A Boolean value that determines whether the text attachment uses text attachment views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857584-allowstextattachmentview?language=objc
func (t_ TextAttachment) SetAllowsTextAttachmentView(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsTextAttachmentView:"), value)
}

// A Boolean value that indicates whether the text attachment uses text attachment views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/3857588-usestextattachmentview?language=objc
func (t_ TextAttachment) UsesTextAttachmentView() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesTextAttachmentView"))
	return rv
}

// The file type of the contents for the text attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508416-filetype?language=objc
func (t_ TextAttachment) FileType() string {
	rv := objc.Call[string](t_, objc.Sel("fileType"))
	return rv
}

// The file type of the contents for the text attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508416-filetype?language=objc
func (t_ TextAttachment) SetFileType(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setFileType:"), value)
}

// The contents for the text attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508401-contents?language=objc
func (t_ TextAttachment) Contents() []byte {
	rv := objc.Call[[]byte](t_, objc.Sel("contents"))
	return rv
}

// The contents for the text attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508401-contents?language=objc
func (t_ TextAttachment) SetContents(value []byte) {
	objc.Call[objc.Void](t_, objc.Sel("setContents:"), value)
}

// The object that draws the icon for the text attachment and handles mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/1508413-attachmentcell?language=objc
func (t_ TextAttachment) AttachmentCell() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("attachmentCell"))
	return rv
}

// The object that draws the icon for the text attachment and handles mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextattachment/1508413-attachmentcell?language=objc
func (t_ TextAttachment) SetAttachmentCell(value objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setAttachmentCell:"), value)
}

// An instance of the relevant image class that represents the contents of the text attachment object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508378-image?language=objc
func (t_ TextAttachment) Image() Image {
	rv := objc.Call[Image](t_, objc.Sel("image"))
	return rv
}

// An instance of the relevant image class that represents the contents of the text attachment object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachment/1508378-image?language=objc
func (t_ TextAttachment) SetImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setImage:"), objc.Ptr(value))
}
