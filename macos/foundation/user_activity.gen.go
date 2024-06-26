// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [UserActivity] class.
var UserActivityClass = _UserActivityClass{objc.GetClass("NSUserActivity")}

type _UserActivityClass struct {
	objc.Class
}

// An interface definition for the [UserActivity] class.
type IUserActivity interface {
	objc.IObject
	Invalidate()
	GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error))
	BecomeCurrent()
	ResignCurrent()
	AddUserInfoEntriesFromDictionary(otherDictionary Dictionary)
	ContextIdentifierPath() []string
	UserInfo() Dictionary
	SetUserInfo(value Dictionary)
	NeedsSave() bool
	SetNeedsSave(value bool)
	ContentAttributeSet() objc.Object
	SetContentAttributeSet(value objc.IObject)
	ExpirationDate() Date
	SetExpirationDate(value IDate)
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	DetectedBarcodeDescriptor() objc.Object
	RequiredUserInfoKeys() Set
	SetRequiredUserInfoKeys(value ISet)
	IsEligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	ActivityType() string
	IsEligibleForSearch() bool
	SetEligibleForSearch(value bool)
	Title() string
	SetTitle(value string)
	SuggestedInvocationPhrase() string
	SetSuggestedInvocationPhrase(value string)
	Keywords() Set
	SetKeywords(value ISet)
	IsEligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	ReferrerURL() URL
	SetReferrerURL(value IURL)
	Delegate() UserActivityDelegateObject
	SetDelegate(value PUserActivityDelegate)
	SetDelegateObject(valueObject objc.IObject)
	IsClassKitDeepLink() bool
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	PersistentIdentifier() UserActivityPersistentIdentifier
	SetPersistentIdentifier(value UserActivityPersistentIdentifier)
	WebpageURL() URL
	SetWebpageURL(value IURL)
}

// A representation of the state of your app at a moment in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity?language=objc
type UserActivity struct {
	objc.Object
}

func UserActivityFrom(ptr unsafe.Pointer) UserActivity {
	return UserActivity{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ UserActivity) InitWithActivityType(activityType string) UserActivity {
	rv := objc.Call[UserActivity](u_, objc.Sel("initWithActivityType:"), activityType)
	return rv
}

// Creates a user activity object with the specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1410714-initwithactivitytype?language=objc
func NewUserActivityWithActivityType(activityType string) UserActivity {
	instance := UserActivityClass.Alloc().InitWithActivityType(activityType)
	instance.Autorelease()
	return instance
}

func (uc _UserActivityClass) Alloc() UserActivity {
	rv := objc.Call[UserActivity](uc, objc.Sel("alloc"))
	return rv
}

func (uc _UserActivityClass) New() UserActivity {
	rv := objc.Call[UserActivity](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserActivity() UserActivity {
	return UserActivityClass.New()
}

func (u_ UserActivity) Init() UserActivity {
	rv := objc.Call[UserActivity](u_, objc.Sel("init"))
	return rv
}

// Invalidates an activity and marks it as no longer eligible for continuation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1416401-invalidate?language=objc
func (u_ UserActivity) Invalidate() {
	objc.Call[objc.Void](u_, objc.Sel("invalidate"))
}

// Requests streams back to the originating app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1409931-getcontinuationstreamswithcomple?language=objc
func (u_ UserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler func(inputStream InputStream, outputStream OutputStream, error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("getContinuationStreamsWithCompletionHandler:"), completionHandler)
}

// Marks the activity as currently in use by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1413665-becomecurrent?language=objc
func (u_ UserActivity) BecomeCurrent() {
	objc.Call[objc.Void](u_, objc.Sel("becomeCurrent"))
}

// Marks this activity object as inactive without invalidating it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1409596-resigncurrent?language=objc
func (u_ UserActivity) ResignCurrent() {
	objc.Call[objc.Void](u_, objc.Sel("resignCurrent"))
}

// Deletes all user activities created by your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980672-deleteallsaveduseractivitieswith?language=objc
func (uc _UserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	objc.Call[objc.Void](uc, objc.Sel("deleteAllSavedUserActivitiesWithCompletionHandler:"), handler)
}

// Deletes all user activities created by your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980672-deleteallsaveduseractivitieswith?language=objc
func UserActivity_DeleteAllSavedUserActivitiesWithCompletionHandler(handler func()) {
	UserActivityClass.DeleteAllSavedUserActivitiesWithCompletionHandler(handler)
}

// Adds the contents of the specified dictionary to the user info dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1410066-adduserinfoentriesfromdictionary?language=objc
func (u_ UserActivity) AddUserInfoEntriesFromDictionary(otherDictionary Dictionary) {
	objc.Call[objc.Void](u_, objc.Sel("addUserInfoEntriesFromDictionary:"), otherDictionary)
}

// Deletes user activities created by your app that have the specified persistent identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980673-deletesaveduseractivitieswithper?language=objc
func (uc _UserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	objc.Call[objc.Void](uc, objc.Sel("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, handler)
}

// Deletes user activities created by your app that have the specified persistent identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980673-deletesaveduseractivitieswithper?language=objc
func UserActivity_DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []UserActivityPersistentIdentifier, handler func()) {
	UserActivityClass.DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers, handler)
}

// The identifier path associated with a user activity generated by an app that adopts ClassKit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2866826-contextidentifierpath?language=objc
func (u_ UserActivity) ContextIdentifierPath() []string {
	rv := objc.Call[[]string](u_, objc.Sel("contextIdentifierPath"))
	return rv
}

// A dictionary containing app-specific state information needed to continue an activity on another device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1411706-userinfo?language=objc
func (u_ UserActivity) UserInfo() Dictionary {
	rv := objc.Call[Dictionary](u_, objc.Sel("userInfo"))
	return rv
}

// A dictionary containing app-specific state information needed to continue an activity on another device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1411706-userinfo?language=objc
func (u_ UserActivity) SetUserInfo(value Dictionary) {
	objc.Call[objc.Void](u_, objc.Sel("setUserInfo:"), value)
}

// A Boolean value that indicates whether the state of the activity needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1408791-needssave?language=objc
func (u_ UserActivity) NeedsSave() bool {
	rv := objc.Call[bool](u_, objc.Sel("needsSave"))
	return rv
}

// A Boolean value that indicates whether the state of the activity needs to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1408791-needssave?language=objc
func (u_ UserActivity) SetNeedsSave(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setNeedsSave:"), value)
}

// A set of properties that describe the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1616398-contentattributeset?language=objc
func (u_ UserActivity) ContentAttributeSet() objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("contentAttributeSet"))
	return rv
}

// A set of properties that describe the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1616398-contentattributeset?language=objc
func (u_ UserActivity) SetContentAttributeSet(value objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setContentAttributeSet:"), value)
}

// The date after which the activity is no longer eligible for Handoff or indexing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1413745-expirationdate?language=objc
func (u_ UserActivity) ExpirationDate() Date {
	rv := objc.Call[Date](u_, objc.Sel("expirationDate"))
	return rv
}

// The date after which the activity is no longer eligible for Handoff or indexing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1413745-expirationdate?language=objc
func (u_ UserActivity) SetExpirationDate(value IDate) {
	objc.Call[objc.Void](u_, objc.Sel("setExpirationDate:"), value)
}

// A string that identifies the user activity’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/3238062-targetcontentidentifier?language=objc
func (u_ UserActivity) TargetContentIdentifier() string {
	rv := objc.Call[string](u_, objc.Sel("targetContentIdentifier"))
	return rv
}

// A string that identifies the user activity’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/3238062-targetcontentidentifier?language=objc
func (u_ UserActivity) SetTargetContentIdentifier(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setTargetContentIdentifier:"), value)
}

// The barcode that the system scanner passes in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2935570-detectedbarcodedescriptor?language=objc
func (u_ UserActivity) DetectedBarcodeDescriptor() objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("detectedBarcodeDescriptor"))
	return rv
}

// A set of keys that represent the minimal information about the activity that should be stored for later restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1417256-requireduserinfokeys?language=objc
func (u_ UserActivity) RequiredUserInfoKeys() Set {
	rv := objc.Call[Set](u_, objc.Sel("requiredUserInfoKeys"))
	return rv
}

// A set of keys that represent the minimal information about the activity that should be stored for later restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1417256-requireduserinfokeys?language=objc
func (u_ UserActivity) SetRequiredUserInfoKeys(value ISet) {
	objc.Call[objc.Void](u_, objc.Sel("setRequiredUserInfoKeys:"), value)
}

// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1414701-eligibleforpublicindexing?language=objc
func (u_ UserActivity) IsEligibleForPublicIndexing() bool {
	rv := objc.Call[bool](u_, objc.Sel("isEligibleForPublicIndexing"))
	return rv
}

// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1414701-eligibleforpublicindexing?language=objc
func (u_ UserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setEligibleForPublicIndexing:"), value)
}

// The user activity object’s activity type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1409611-activitytype?language=objc
func (u_ UserActivity) ActivityType() string {
	rv := objc.Call[string](u_, objc.Sel("activityType"))
	return rv
}

// A Boolean value that indicates whether the activity should be added to the on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1417761-eligibleforsearch?language=objc
func (u_ UserActivity) IsEligibleForSearch() bool {
	rv := objc.Call[bool](u_, objc.Sel("isEligibleForSearch"))
	return rv
}

// A Boolean value that indicates whether the activity should be added to the on-device index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1417761-eligibleforsearch?language=objc
func (u_ UserActivity) SetEligibleForSearch(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setEligibleForSearch:"), value)
}

// An optional, user-visible title for this activity, such as a document name or web page title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1413375-title?language=objc
func (u_ UserActivity) Title() string {
	rv := objc.Call[string](u_, objc.Sel("title"))
	return rv
}

// An optional, user-visible title for this activity, such as a document name or web page title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1413375-title?language=objc
func (u_ UserActivity) SetTitle(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setTitle:"), value)
}

// A phrase suggested to the user when they create a shortcut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2976237-suggestedinvocationphrase?language=objc
func (u_ UserActivity) SuggestedInvocationPhrase() string {
	rv := objc.Call[string](u_, objc.Sel("suggestedInvocationPhrase"))
	return rv
}

// A phrase suggested to the user when they create a shortcut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2976237-suggestedinvocationphrase?language=objc
func (u_ UserActivity) SetSuggestedInvocationPhrase(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setSuggestedInvocationPhrase:"), value)
}

// A set of localized keywords that can help users find the activity in search results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1408023-keywords?language=objc
func (u_ UserActivity) Keywords() Set {
	rv := objc.Call[Set](u_, objc.Sel("keywords"))
	return rv
}

// A set of localized keywords that can help users find the activity in search results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1408023-keywords?language=objc
func (u_ UserActivity) SetKeywords(value ISet) {
	objc.Call[objc.Void](u_, objc.Sel("setKeywords:"), value)
}

// A Boolean value that indicates whether the activity can be continued on another device using Handoff. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1410971-eligibleforhandoff?language=objc
func (u_ UserActivity) IsEligibleForHandoff() bool {
	rv := objc.Call[bool](u_, objc.Sel("isEligibleForHandoff"))
	return rv
}

// A Boolean value that indicates whether the activity can be continued on another device using Handoff. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1410971-eligibleforhandoff?language=objc
func (u_ UserActivity) SetEligibleForHandoff(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setEligibleForHandoff:"), value)
}

// The URL of the webpage that linked to the webpage URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2875762-referrerurl?language=objc
func (u_ UserActivity) ReferrerURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("referrerURL"))
	return rv
}

// The URL of the webpage that linked to the webpage URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2875762-referrerurl?language=objc
func (u_ UserActivity) SetReferrerURL(value IURL) {
	objc.Call[objc.Void](u_, objc.Sel("setReferrerURL:"), value)
}

// The user activity object’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1412329-delegate?language=objc
func (u_ UserActivity) Delegate() UserActivityDelegateObject {
	rv := objc.Call[UserActivityDelegateObject](u_, objc.Sel("delegate"))
	return rv
}

// The user activity object’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1412329-delegate?language=objc
func (u_ UserActivity) SetDelegate(value PUserActivityDelegate) {
	po0 := objc.WrapAsProtocol("NSUserActivityDelegate", value)
	objc.SetAssociatedObject(u_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](u_, objc.Sel("setDelegate:"), po0)
}

// The user activity object’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1412329-delegate?language=objc
func (u_ UserActivity) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setDelegate:"), valueObject)
}

// A Boolean value that indicates whether a user activity represents a ClassKit context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2953056-isclasskitdeeplink?language=objc
func (u_ UserActivity) IsClassKitDeepLink() bool {
	rv := objc.Call[bool](u_, objc.Sel("isClassKitDeepLink"))
	return rv
}

// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1409195-supportscontinuationstreams?language=objc
func (u_ UserActivity) SupportsContinuationStreams() bool {
	rv := objc.Call[bool](u_, objc.Sel("supportsContinuationStreams"))
	return rv
}

// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1409195-supportscontinuationstreams?language=objc
func (u_ UserActivity) SetSupportsContinuationStreams(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setSupportsContinuationStreams:"), value)
}

// A value used to identify the user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980675-persistentidentifier?language=objc
func (u_ UserActivity) PersistentIdentifier() UserActivityPersistentIdentifier {
	rv := objc.Call[UserActivityPersistentIdentifier](u_, objc.Sel("persistentIdentifier"))
	return rv
}

// A value used to identify the user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/2980675-persistentidentifier?language=objc
func (u_ UserActivity) SetPersistentIdentifier(value UserActivityPersistentIdentifier) {
	objc.Call[objc.Void](u_, objc.Sel("setPersistentIdentifier:"), value)
}

// The URL of the webpage to load in a browser to continue the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1418086-webpageurl?language=objc
func (u_ UserActivity) WebpageURL() URL {
	rv := objc.Call[URL](u_, objc.Sel("webpageURL"))
	return rv
}

// The URL of the webpage to load in a browser to continue the activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsuseractivity/1418086-webpageurl?language=objc
func (u_ UserActivity) SetWebpageURL(value IURL) {
	objc.Call[objc.Void](u_, objc.Sel("setWebpageURL:"), value)
}
