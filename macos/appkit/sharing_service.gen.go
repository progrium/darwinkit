// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SharingService] class.
var SharingServiceClass = _SharingServiceClass{objc.GetClass("NSSharingService")}

type _SharingServiceClass struct {
	objc.Class
}

// An interface definition for the [SharingService] class.
type ISharingService interface {
	objc.IObject
	CanPerformWithItems(items []objc.IObject) bool
	PerformWithItems(items []objc.IObject)
	PermanentLink() foundation.URL
	MessageBody() string
	AlternateImage() Image
	Recipients() []string
	SetRecipients(value []string)
	Delegate() SharingServiceDelegateWrapper
	SetDelegate(value PSharingServiceDelegate)
	SetDelegateObject(valueObject objc.IObject)
	MenuItemTitle() string
	SetMenuItemTitle(value string)
	AttachmentFileURLs() []foundation.URL
	Title() string
	Subject() string
	SetSubject(value string)
	Image() Image
	AccountName() string
}

// An object that facilitates the sharing of content with social media services, or with apps like Mail or Safari. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice?language=objc
type SharingService struct {
	objc.Object
}

func SharingServiceFrom(ptr unsafe.Pointer) SharingService {
	return SharingService{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SharingService) InitWithTitleImageAlternateImageHandler(title string, image IImage, alternateImage IImage, block func()) SharingService {
	rv := objc.Call[SharingService](s_, objc.Sel("initWithTitle:image:alternateImage:handler:"), title, objc.Ptr(image), objc.Ptr(alternateImage), block)
	return rv
}

// Creates a custom sharing service object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402614-initwithtitle?language=objc
func NewSharingServiceWithTitleImageAlternateImageHandler(title string, image IImage, alternateImage IImage, block func()) SharingService {
	instance := SharingServiceClass.Alloc().InitWithTitleImageAlternateImageHandler(title, image, alternateImage, block)
	instance.Autorelease()
	return instance
}

func (sc _SharingServiceClass) Alloc() SharingService {
	rv := objc.Call[SharingService](sc, objc.Sel("alloc"))
	return rv
}

func SharingService_Alloc() SharingService {
	return SharingServiceClass.Alloc()
}

func (sc _SharingServiceClass) New() SharingService {
	rv := objc.Call[SharingService](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSharingService() SharingService {
	return SharingServiceClass.New()
}

func (s_ SharingService) Init() SharingService {
	rv := objc.Call[SharingService](s_, objc.Sel("init"))
	return rv
}

// Returns whether the service can share all the specified items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402662-canperformwithitems?language=objc
func (s_ SharingService) CanPerformWithItems(items []objc.IObject) bool {
	rv := objc.Call[bool](s_, objc.Sel("canPerformWithItems:"), items)
	return rv
}

// Manually performs the service on the provided items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402669-performwithitems?language=objc
func (s_ SharingService) PerformWithItems(items []objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("performWithItems:"), items)
}

// Returns a sharing service instance representing the specified service name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402673-sharingservicenamed?language=objc
func (sc _SharingServiceClass) SharingServiceNamed(serviceName SharingServiceName) SharingService {
	rv := objc.Call[SharingService](sc, objc.Sel("sharingServiceNamed:"), serviceName)
	return rv
}

// Returns a sharing service instance representing the specified service name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402673-sharingservicenamed?language=objc
func SharingService_SharingServiceNamed(serviceName SharingServiceName) SharingService {
	return SharingServiceClass.SharingServiceNamed(serviceName)
}

// A permanent URL (permalink) that your app can use to access the post. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402685-permanentlink?language=objc
func (s_ SharingService) PermanentLink() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("permanentLink"))
	return rv
}

// The message body as a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402644-messagebody?language=objc
func (s_ SharingService) MessageBody() string {
	rv := objc.Call[string](s_, objc.Sel("messageBody"))
	return rv
}

// The alternate image representing the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402650-alternateimage?language=objc
func (s_ SharingService) AlternateImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("alternateImage"))
	return rv
}

// An array containing the user handles of the desired recipients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402652-recipients?language=objc
func (s_ SharingService) Recipients() []string {
	rv := objc.Call[[]string](s_, objc.Sel("recipients"))
	return rv
}

// An array containing the user handles of the desired recipients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402652-recipients?language=objc
func (s_ SharingService) SetRecipients(value []string) {
	objc.Call[objc.Void](s_, objc.Sel("setRecipients:"), value)
}

// Specifies the delegate of the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402681-delegate?language=objc
func (s_ SharingService) Delegate() SharingServiceDelegateWrapper {
	rv := objc.Call[SharingServiceDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// Specifies the delegate of the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402681-delegate?language=objc
func (s_ SharingService) SetDelegate(value PSharingServiceDelegate) {
	po0 := objc.WrapAsProtocol("NSSharingServiceDelegate", value)
	objc.SetAssociatedObject(s_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// Specifies the delegate of the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402681-delegate?language=objc
func (s_ SharingService) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The title of the service in the Share menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402693-menuitemtitle?language=objc
func (s_ SharingService) MenuItemTitle() string {
	rv := objc.Call[string](s_, objc.Sel("menuItemTitle"))
	return rv
}

// The title of the service in the Share menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402693-menuitemtitle?language=objc
func (s_ SharingService) SetMenuItemTitle(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setMenuItemTitle:"), value)
}

// An array of NSURL objects representing the files that were shared. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402707-attachmentfileurls?language=objc
func (s_ SharingService) AttachmentFileURLs() []foundation.URL {
	rv := objc.Call[[]foundation.URL](s_, objc.Sel("attachmentFileURLs"))
	return rv
}

// The title of the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402637-title?language=objc
func (s_ SharingService) Title() string {
	rv := objc.Call[string](s_, objc.Sel("title"))
	return rv
}

// The subject of the post. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402626-subject?language=objc
func (s_ SharingService) Subject() string {
	rv := objc.Call[string](s_, objc.Sel("subject"))
	return rv
}

// The subject of the post. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402626-subject?language=objc
func (s_ SharingService) SetSubject(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setSubject:"), value)
}

// The primary image representing the sharing service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402654-image?language=objc
func (s_ SharingService) Image() Image {
	rv := objc.Call[Image](s_, objc.Sel("image"))
	return rv
}

// The account name used for posting on Twitter or Sina Weibo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservice/1402683-accountname?language=objc
func (s_ SharingService) AccountName() string {
	rv := objc.Call[string](s_, objc.Sel("accountName"))
	return rv
}
