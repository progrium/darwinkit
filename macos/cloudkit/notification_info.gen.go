// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NotificationInfo] class.
var NotificationInfoClass = _NotificationInfoClass{objc.GetClass("CKNotificationInfo")}

type _NotificationInfoClass struct {
	objc.Class
}

// An interface definition for the [NotificationInfo] class.
type INotificationInfo interface {
	objc.IObject
	TitleLocalizationKey() string
	SetTitleLocalizationKey(value string)
	AlertBody() string
	SetAlertBody(value string)
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	TitleLocalizationArgs() []RecordFieldKey
	SetTitleLocalizationArgs(value []RecordFieldKey)
	SoundName() string
	SetSoundName(value string)
	AlertLaunchImage() string
	SetAlertLaunchImage(value string)
	SubtitleLocalizationKey() string
	SetSubtitleLocalizationKey(value string)
	AlertActionLocalizationKey() string
	SetAlertActionLocalizationKey(value string)
	AlertLocalizationKey() string
	SetAlertLocalizationKey(value string)
	Subtitle() string
	SetSubtitle(value string)
	SubtitleLocalizationArgs() []RecordFieldKey
	SetSubtitleLocalizationArgs(value []RecordFieldKey)
	Category() string
	SetCategory(value string)
	ShouldSendMutableContent() bool
	SetShouldSendMutableContent(value bool)
	CollapseIDKey() string
	SetCollapseIDKey(value string)
	Title() string
	SetTitle(value string)
	ShouldBadge() bool
	SetShouldBadge(value bool)
	AlertLocalizationArgs() []RecordFieldKey
	SetAlertLocalizationArgs(value []RecordFieldKey)
	DesiredKeys() []RecordFieldKey
	SetDesiredKeys(value []RecordFieldKey)
}

// An object that describes the configuration of a subscription’s push notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo?language=objc
type NotificationInfo struct {
	objc.Object
}

func NotificationInfoFrom(ptr unsafe.Pointer) NotificationInfo {
	return NotificationInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NotificationInfoClass) Alloc() NotificationInfo {
	rv := objc.Call[NotificationInfo](nc, objc.Sel("alloc"))
	return rv
}

func NotificationInfo_Alloc() NotificationInfo {
	return NotificationInfoClass.Alloc()
}

func (nc _NotificationInfoClass) New() NotificationInfo {
	rv := objc.Call[NotificationInfo](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNotificationInfo() NotificationInfo {
	return NotificationInfoClass.New()
}

func (n_ NotificationInfo) Init() NotificationInfo {
	rv := objc.Call[NotificationInfo](n_, objc.Sel("init"))
	return rv
}

// The key that identifies the localized string for the notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869865-titlelocalizationkey?language=objc
func (n_ NotificationInfo) TitleLocalizationKey() string {
	rv := objc.Call[string](n_, objc.Sel("titleLocalizationKey"))
	return rv
}

// The key that identifies the localized string for the notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869865-titlelocalizationkey?language=objc
func (n_ NotificationInfo) SetTitleLocalizationKey(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setTitleLocalizationKey:"), value)
}

// The text for the notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515270-alertbody?language=objc
func (n_ NotificationInfo) AlertBody() string {
	rv := objc.Call[string](n_, objc.Sel("alertBody"))
	return rv
}

// The text for the notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515270-alertbody?language=objc
func (n_ NotificationInfo) SetAlertBody(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setAlertBody:"), value)
}

// A Boolean value that indicates whether the push notification includes the content available flag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515110-shouldsendcontentavailable?language=objc
func (n_ NotificationInfo) ShouldSendContentAvailable() bool {
	rv := objc.Call[bool](n_, objc.Sel("shouldSendContentAvailable"))
	return rv
}

// A Boolean value that indicates whether the push notification includes the content available flag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515110-shouldsendcontentavailable?language=objc
func (n_ NotificationInfo) SetShouldSendContentAvailable(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setShouldSendContentAvailable:"), value)
}

// The fields for building a notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869866-titlelocalizationargs?language=objc
func (n_ NotificationInfo) TitleLocalizationArgs() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](n_, objc.Sel("titleLocalizationArgs"))
	return rv
}

// The fields for building a notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869866-titlelocalizationargs?language=objc
func (n_ NotificationInfo) SetTitleLocalizationArgs(value []RecordFieldKey) {
	objc.Call[objc.Void](n_, objc.Sel("setTitleLocalizationArgs:"), value)
}

// The filename of the sound file to play when a notification arrives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514987-soundname?language=objc
func (n_ NotificationInfo) SoundName() string {
	rv := objc.Call[string](n_, objc.Sel("soundName"))
	return rv
}

// The filename of the sound file to play when a notification arrives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514987-soundname?language=objc
func (n_ NotificationInfo) SetSoundName(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setSoundName:"), value)
}

// The filename of an image to use as a launch image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515075-alertlaunchimage?language=objc
func (n_ NotificationInfo) AlertLaunchImage() string {
	rv := objc.Call[string](n_, objc.Sel("alertLaunchImage"))
	return rv
}

// The filename of an image to use as a launch image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515075-alertlaunchimage?language=objc
func (n_ NotificationInfo) SetAlertLaunchImage(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setAlertLaunchImage:"), value)
}

// The key that identifies the localized string for the notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869864-subtitlelocalizationkey?language=objc
func (n_ NotificationInfo) SubtitleLocalizationKey() string {
	rv := objc.Call[string](n_, objc.Sel("subtitleLocalizationKey"))
	return rv
}

// The key that identifies the localized string for the notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869864-subtitlelocalizationkey?language=objc
func (n_ NotificationInfo) SetSubtitleLocalizationKey(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setSubtitleLocalizationKey:"), value)
}

// The key that identifies the localized string for the notification’s action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514945-alertactionlocalizationkey?language=objc
func (n_ NotificationInfo) AlertActionLocalizationKey() string {
	rv := objc.Call[string](n_, objc.Sel("alertActionLocalizationKey"))
	return rv
}

// The key that identifies the localized string for the notification’s action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514945-alertactionlocalizationkey?language=objc
func (n_ NotificationInfo) SetAlertActionLocalizationKey(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setAlertActionLocalizationKey:"), value)
}

// The key that identifies the localized string for the notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514968-alertlocalizationkey?language=objc
func (n_ NotificationInfo) AlertLocalizationKey() string {
	rv := objc.Call[string](n_, objc.Sel("alertLocalizationKey"))
	return rv
}

// The key that identifies the localized string for the notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514968-alertlocalizationkey?language=objc
func (n_ NotificationInfo) SetAlertLocalizationKey(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setAlertLocalizationKey:"), value)
}

// The notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869863-subtitle?language=objc
func (n_ NotificationInfo) Subtitle() string {
	rv := objc.Call[string](n_, objc.Sel("subtitle"))
	return rv
}

// The notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869863-subtitle?language=objc
func (n_ NotificationInfo) SetSubtitle(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setSubtitle:"), value)
}

// The fields for building a notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869869-subtitlelocalizationargs?language=objc
func (n_ NotificationInfo) SubtitleLocalizationArgs() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](n_, objc.Sel("subtitleLocalizationArgs"))
	return rv
}

// The fields for building a notification’s subtitle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869869-subtitlelocalizationargs?language=objc
func (n_ NotificationInfo) SetSubtitleLocalizationArgs(value []RecordFieldKey) {
	objc.Call[objc.Void](n_, objc.Sel("setSubtitleLocalizationArgs:"), value)
}

// The name of the action group that corresponds to this notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515082-category?language=objc
func (n_ NotificationInfo) Category() string {
	rv := objc.Call[string](n_, objc.Sel("category"))
	return rv
}

// The name of the action group that corresponds to this notification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515082-category?language=objc
func (n_ NotificationInfo) SetCategory(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCategory:"), value)
}

// A Boolean value that indicates whether the push notification sets the mutable content flag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2868500-shouldsendmutablecontent?language=objc
func (n_ NotificationInfo) ShouldSendMutableContent() bool {
	rv := objc.Call[bool](n_, objc.Sel("shouldSendMutableContent"))
	return rv
}

// A Boolean value that indicates whether the push notification sets the mutable content flag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2868500-shouldsendmutablecontent?language=objc
func (n_ NotificationInfo) SetShouldSendMutableContent(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setShouldSendMutableContent:"), value)
}

// A value that the system uses to coalesce unseen push notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2887430-collapseidkey?language=objc
func (n_ NotificationInfo) CollapseIDKey() string {
	rv := objc.Call[string](n_, objc.Sel("collapseIDKey"))
	return rv
}

// A value that the system uses to coalesce unseen push notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2887430-collapseidkey?language=objc
func (n_ NotificationInfo) SetCollapseIDKey(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setCollapseIDKey:"), value)
}

// The notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869870-title?language=objc
func (n_ NotificationInfo) Title() string {
	rv := objc.Call[string](n_, objc.Sel("title"))
	return rv
}

// The notification’s title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/2869870-title?language=objc
func (n_ NotificationInfo) SetTitle(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setTitle:"), value)
}

// A Boolean value that determines whether an app’s icon badge increments its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514996-shouldbadge?language=objc
func (n_ NotificationInfo) ShouldBadge() bool {
	rv := objc.Call[bool](n_, objc.Sel("shouldBadge"))
	return rv
}

// A Boolean value that determines whether an app’s icon badge increments its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514996-shouldbadge?language=objc
func (n_ NotificationInfo) SetShouldBadge(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setShouldBadge:"), value)
}

// The fields for building a notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515182-alertlocalizationargs?language=objc
func (n_ NotificationInfo) AlertLocalizationArgs() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](n_, objc.Sel("alertLocalizationArgs"))
	return rv
}

// The fields for building a notification’s alert. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1515182-alertlocalizationargs?language=objc
func (n_ NotificationInfo) SetAlertLocalizationArgs(value []RecordFieldKey) {
	objc.Call[objc.Void](n_, objc.Sel("setAlertLocalizationArgs:"), value)
}

// The names of fields to include in the push notification’s payload. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514931-desiredkeys?language=objc
func (n_ NotificationInfo) DesiredKeys() []RecordFieldKey {
	rv := objc.Call[[]RecordFieldKey](n_, objc.Sel("desiredKeys"))
	return rv
}

// The names of fields to include in the push notification’s payload. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cknotificationinfo/1514931-desiredkeys?language=objc
func (n_ NotificationInfo) SetDesiredKeys(value []RecordFieldKey) {
	objc.Call[objc.Void](n_, objc.Sel("setDesiredKeys:"), value)
}
