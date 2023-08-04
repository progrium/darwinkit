// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}

type _SharingServiceClass struct {
	objc.Class
}

type ISharingService interface {
	objc.IObject
	CanPerformWithItems(items []objc.IObject) bool
	PerformWithItems(items []objc.IObject)
	Delegate() SharingServiceDelegateWrapper
	SetDelegate(value ISharingServiceDelegate)
	SetDelegate0(value objc.IObject)
	AccountName() string
	AlternateImage() Image
	Image() Image
	Title() string
	MenuItemTitle() string
	SetMenuItemTitle(value string)
	Recipients() []string
	SetRecipients(value []string)
	Subject() string
	SetSubject(value string)
	AttachmentFileURLs() []foundation.URL
	MessageBody() string
	PermanentLink() foundation.URL
}

type SharingService struct {
	objc.Object
}

func MakeSharingService(ptr unsafe.Pointer) SharingService {
	return SharingService{
		Object: objc.MakeObject(ptr),
	}
}

func (s_ SharingService) InitWithTitleImageAlternateImageHandler(title string, image IImage, alternateImage IImage, block func()) SharingService {
	rv := objc.CallMethod[SharingService](s_, objc.GetSelector("initWithTitle:image:alternateImage:handler:"), title, objc.ExtractPtr(image), objc.ExtractPtr(alternateImage), block)
	return rv
}

func SharingService_InitWithTitleImageAlternateImageHandler(title string, image IImage, alternateImage IImage, block func()) SharingService {
	return SharingServiceClass.Alloc().InitWithTitleImageAlternateImageHandler(title, image, alternateImage, block)
}

func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.GetSelector("alloc"))
	return rv
}

func SharingService_Alloc() SharingService {
	return SharingServiceClass.Alloc()
}

func (sc _SharingServiceClass) New() SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

func SharingService_New() SharingService {
	return SharingServiceClass.New()
}

func (s_ SharingService) Init() SharingService {
	rv := objc.CallMethod[SharingService](s_, objc.GetSelector("init"))
	return rv
}

func SharingService_Init() SharingService {
	return SharingServiceClass.Alloc().Init()
}

func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := objc.CallMethod[SharingService](sc, objc.GetSelector("sharingServiceNamed:"), serviceName)
	return rv
}

func SharingService_SharingServiceNamed(serviceName SharingServiceName) SharingService {
	return SharingServiceClass.SharingServiceNamed(serviceName)
}

func (s_ SharingService) CanPerformWithItems(items []objc.IObject) bool {
	rv := objc.CallMethod[bool](s_, objc.GetSelector("canPerformWithItems:"), items)
	return rv
}

func (s_ SharingService) PerformWithItems(items []objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("performWithItems:"), items)
}

func (s_ SharingService) Delegate() SharingServiceDelegateWrapper {
	rv := objc.CallMethod[SharingServiceDelegateWrapper](s_, objc.GetSelector("delegate"))
	return rv
}

func (s_ SharingService) SetDelegate(value ISharingServiceDelegate) {
	po := objc.WrapAsProtocol("NSSharingServiceDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), po)
}

func (s_ SharingService) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (s_ SharingService) AccountName() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("accountName"))
	return rv
}

func (s_ SharingService) AlternateImage() Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("alternateImage"))
	return rv
}

func (s_ SharingService) Image() Image {
	rv := objc.CallMethod[Image](s_, objc.GetSelector("image"))
	return rv
}

func (s_ SharingService) Title() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("title"))
	return rv
}

func (s_ SharingService) MenuItemTitle() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("menuItemTitle"))
	return rv
}

func (s_ SharingService) SetMenuItemTitle(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setMenuItemTitle:"), value)
}

func (s_ SharingService) Recipients() []string {
	rv := objc.CallMethod[[]string](s_, objc.GetSelector("recipients"))
	return rv
}

func (s_ SharingService) SetRecipients(value []string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setRecipients:"), value)
}

func (s_ SharingService) Subject() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("subject"))
	return rv
}

func (s_ SharingService) SetSubject(value string) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("setSubject:"), value)
}

func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.CallMethod[[]foundation.URL](s_, objc.GetSelector("attachmentFileURLs"))
	return rv
}

func (s_ SharingService) MessageBody() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("messageBody"))
	return rv
}

func (s_ SharingService) PermanentLink() foundation.URL {
	rv := objc.CallMethod[foundation.URL](s_, objc.GetSelector("permanentLink"))
	return rv
}
