// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var TextAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}

type _TextAttachmentViewProviderClass struct {
	objc.Class
}

type ITextAttachmentViewProvider interface {
	objc.IObject
	LoadView()
	TextAttachment() TextAttachment
	TextLayoutManager() TextLayoutManager
	TracksTextAttachmentViewBounds() bool
	SetTracksTextAttachmentViewBounds(value bool)
	View() View
	SetView(value IView)
}

type TextAttachmentViewProvider struct {
	objc.Object
}

func MakeTextAttachmentViewProvider(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{
		Object: objc.MakeObject(ptr),
	}
}

func (tc _TextAttachmentViewProviderClass) Alloc() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](tc, objc.GetSelector("alloc"))
	return rv
}

func TextAttachmentViewProvider_Alloc() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.Alloc()
}

func (tc _TextAttachmentViewProviderClass) New() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachmentViewProvider() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.New()
}

func TextAttachmentViewProvider_New() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.New()
}

func (t_ TextAttachmentViewProvider) Init() TextAttachmentViewProvider {
	rv := objc.CallMethod[TextAttachmentViewProvider](t_, objc.GetSelector("init"))
	return rv
}

func TextAttachmentViewProvider_Init() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.Alloc().Init()
}

func (t_ TextAttachmentViewProvider) LoadView() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("loadView"))
}

func (t_ TextAttachmentViewProvider) TextAttachment() TextAttachment {
	rv := objc.CallMethod[TextAttachment](t_, objc.GetSelector("textAttachment"))
	return rv
}

func (t_ TextAttachmentViewProvider) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("textLayoutManager"))
	return rv
}

func (t_ TextAttachmentViewProvider) TracksTextAttachmentViewBounds() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tracksTextAttachmentViewBounds"))
	return rv
}

func (t_ TextAttachmentViewProvider) SetTracksTextAttachmentViewBounds(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTracksTextAttachmentViewBounds:"), value)
}

func (t_ TextAttachmentViewProvider) View() View {
	rv := objc.CallMethod[View](t_, objc.GetSelector("view"))
	return rv
}

func (t_ TextAttachmentViewProvider) SetView(value IView) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setView:"), objc.ExtractPtr(value))
}
