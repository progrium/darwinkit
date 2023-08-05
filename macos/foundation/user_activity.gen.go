// AUTO-GENERATED CODE, DO NOT MODIFY
package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var UserActivityClass = _UserActivityClass{objc.GetClass("NSUserActivity")}

type _UserActivityClass struct {
	objc.Class
}

type IUserActivity interface {
	objc.IObject
	BecomeCurrent()
	ResignCurrent()
	Invalidate()
	GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error))
	ActivityType() string
	Title() string
	SetTitle(value string)
	RequiredUserInfoKeys() Set
	SetRequiredUserInfoKeys(value ISet)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	NeedsSave() bool
	SetNeedsSave(value bool)
	Keywords() Set
	SetKeywords(value ISet)
	PersistentIdentifier() UserActivityPersistentIdentifier
	SetPersistentIdentifier(value UserActivityPersistentIdentifier)
	IsEligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	IsEligibleForSearch() bool
	SetEligibleForSearch(value bool)
	IsEligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	ExpirationDate() Date
	SetExpirationDate(value IDate)
	Delegate() UserActivityDelegateWrapper
	SetDelegate(value IUserActivityDelegate)
	SetDelegate0(value objc.IObject)
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	WebpageURL() URL
	SetWebpageURL(value IURL)
	ReferrerURL() URL
	SetReferrerURL(value IURL)
}

type UserActivity struct {
	objc.Object
}

func MakeUserActivity(ptr unsafe.Pointer) UserActivity {
	return UserActivity{
		Object: objc.MakeObject(ptr),
	}
}

func (u_ UserActivity) InitWithActivityType(activityType string) UserActivity {
	rv := objc.CallMethod[UserActivity](u_, objc.GetSelector("initWithActivityType:"), activityType)
	return rv
}

func UserActivity_InitWithActivityType(activityType string) UserActivity {
	return UserActivityClass.Alloc().InitWithActivityType(activityType)
}

func (u_ UserActivity) Init() UserActivity {
	rv := objc.CallMethod[UserActivity](u_, objc.GetSelector("init"))
	return rv
}

func UserActivity_Init() UserActivity {
	return UserActivityClass.Alloc().Init()
}

func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, objc.GetSelector("alloc"))
	return rv
}

func UserActivity_Alloc() UserActivity {
	return UserActivityClass.Alloc()
}

func (uc _UserActivityClass) New() UserActivity {
	rv := objc.CallMethod[UserActivity](uc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewUserActivity() UserActivity {
	return UserActivityClass.New()
}

func UserActivity_New() UserActivity {
	return UserActivityClass.New()
}

func (u_ UserActivity) BecomeCurrent() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("becomeCurrent"))
}

func (u_ UserActivity) ResignCurrent() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("resignCurrent"))
}

func (u_ UserActivity) Invalidate() {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("invalidate"))
}

func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	objc.CallMethod[objc.Void](uc, objc.GetSelector("deleteAllSavedUserActivitiesWithCompletionHandler:"), handler)
}

func UserActivity_DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	UserActivityClass.DeleteAllSavedUserActivitiesWithCompletionHandler(handler)
}

func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	objc.CallMethod[objc.Void](uc, objc.GetSelector("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, handler)
}

func UserActivity_DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	UserActivityClass.DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers, handler)
}

func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("getContinuationStreamsWithCompletionHandler:"), completionHandler)
}

func (u_ UserActivity) ActivityType() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("activityType"))
	return rv
}

func (u_ UserActivity) Title() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("title"))
	return rv
}

func (u_ UserActivity) SetTitle(value string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setTitle:"), value)
}

func (u_ UserActivity) RequiredUserInfoKeys() Set {
	rv := objc.CallMethod[Set](u_, objc.GetSelector("requiredUserInfoKeys"))
	return rv
}

func (u_ UserActivity) SetRequiredUserInfoKeys(value ISet) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setRequiredUserInfoKeys:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) TargetContentIdentifier() string {
	rv := objc.CallMethod[string](u_, objc.GetSelector("targetContentIdentifier"))
	return rv
}

func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setTargetContentIdentifier:"), value)
}

func (u_ UserActivity) NeedsSave() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("needsSave"))
	return rv
}

func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setNeedsSave:"), value)
}

func (u_ UserActivity) Keywords() Set {
	rv := objc.CallMethod[Set](u_, objc.GetSelector("keywords"))
	return rv
}

func (u_ UserActivity) SetKeywords(value ISet) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setKeywords:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := objc.CallMethod[UserActivityPersistentIdentifier](u_, objc.GetSelector("persistentIdentifier"))
	return rv
}

func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setPersistentIdentifier:"), value)
}

func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isEligibleForHandoff"))
	return rv
}

func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setEligibleForHandoff:"), value)
}

func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isEligibleForSearch"))
	return rv
}

func (u_ UserActivity) SetEligibleForSearch(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setEligibleForSearch:"), value)
}

func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("isEligibleForPublicIndexing"))
	return rv
}

func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setEligibleForPublicIndexing:"), value)
}

func (u_ UserActivity) ExpirationDate() Date {
	rv := objc.CallMethod[Date](u_, objc.GetSelector("expirationDate"))
	return rv
}

func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setExpirationDate:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) Delegate() UserActivityDelegateWrapper {
	rv := objc.CallMethod[UserActivityDelegateWrapper](u_, objc.GetSelector("delegate"))
	return rv
}

func (u_ UserActivity) SetDelegate(value IUserActivityDelegate) {
	po := objc.WrapAsProtocol("NSUserActivityDelegate", value)
	objc.SetAssociatedObject(u_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setDelegate:"), po)
}

func (u_ UserActivity) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.CallMethod[bool](u_, objc.GetSelector("supportsContinuationStreams"))
	return rv
}

func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setSupportsContinuationStreams:"), value)
}

func (u_ UserActivity) WebpageURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("webpageURL"))
	return rv
}

func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setWebpageURL:"), objc.ExtractPtr(value))
}

func (u_ UserActivity) ReferrerURL() URL {
	rv := objc.CallMethod[URL](u_, objc.GetSelector("referrerURL"))
	return rv
}

func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("setReferrerURL:"), objc.ExtractPtr(value))
}
