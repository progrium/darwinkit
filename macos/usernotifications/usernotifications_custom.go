package usernotifications

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// NotificationCenter represents a notification center
type NotificationCenter struct {
	objc.Object
}

// NotificationCenterClass is the class instance for NotificationCenter
var NotificationCenterClass = objc.GetClass("UNNotificationCenter")

// CurrentNotificationCenter returns the current notification center
func CurrentNotificationCenter() NotificationCenter {
	return NotificationCenter{objc.Call[objc.Object](NotificationCenterClass, objc.Sel("currentNotificationCenter"))}
}

// SetDelegate sets the delegate for the notification center
func (n NotificationCenter) SetDelegate(delegate objc.Object) {
	n.Send(objc.Sel("setDelegate:"), delegate)
}

// RequestAuthorizationWithOptionsCompletionHandler requests authorization for notifications
func (n NotificationCenter) RequestAuthorizationWithOptionsCompletionHandler(options AuthorizationOptions, completionHandler foundation.CompletionHandler) {
	n.Send(objc.Sel("requestAuthorizationWithOptions:completionHandler:"), uint(options), completionHandler)
}

// AddNotificationRequestWithCompletionHandler adds a notification request
func (n NotificationCenter) AddNotificationRequestWithCompletionHandler(request NotificationRequest, completionHandler foundation.CompletionHandler) {
	n.Send(objc.Sel("addNotificationRequest:withCompletionHandler:"), request, completionHandler)
}

// RemoveDeliveredNotificationsWithIdentifiers removes delivered notifications
func (n NotificationCenter) RemoveDeliveredNotificationsWithIdentifiers(identifiers foundation.Array) {
	n.Send(objc.Sel("removeDeliveredNotificationsWithIdentifiers:"), identifiers)
}

// RemoveAllDeliveredNotifications removes all delivered notifications
func (n NotificationCenter) RemoveAllDeliveredNotifications() {
	n.Send(objc.Sel("removeAllDeliveredNotifications"))
}

// RemovePendingNotificationRequestsWithIdentifiers removes pending notification requests
func (n NotificationCenter) RemovePendingNotificationRequestsWithIdentifiers(identifiers foundation.Array) {
	n.Send(objc.Sel("removePendingNotificationRequestsWithIdentifiers:"), identifiers)
}

// RemoveAllPendingNotificationRequests removes all pending notification requests
func (n NotificationCenter) RemoveAllPendingNotificationRequests() {
	n.Send(objc.Sel("removeAllPendingNotificationRequests"))
}

// NotificationRequest represents a notification request
type NotificationRequest struct {
	objc.Object
}

// NotificationRequestClass is the class instance for NotificationRequest
var NotificationRequestClass = objc.GetClass("UNNotificationRequest")

// NotificationRequestWithIdentifierContentTrigger creates a notification request
func NotificationRequestWithIdentifierContentTrigger(identifier foundation.String, content NotificationContent, trigger NotificationTrigger) NotificationRequest {
	return NotificationRequest{objc.Call[objc.Object](NotificationRequestClass, objc.Sel("requestWithIdentifier:content:trigger:"), identifier, content, trigger)}
}

// Identifier returns the identifier of the notification request
func (n NotificationRequest) Identifier() foundation.String {
	return objc.Call[foundation.String](n, objc.Sel("identifier"))
}

// Content returns the content of the notification request
func (n NotificationRequest) Content() NotificationContent {
	return NotificationContent{objc.Call[objc.Object](n, objc.Sel("content"))}
}

// Trigger returns the trigger of the notification request
func (n NotificationRequest) Trigger() NotificationTrigger {
	return NotificationTrigger{objc.Call[objc.Object](n, objc.Sel("trigger"))}
}

// NotificationContent represents the content of a notification
type NotificationContent struct {
	objc.Object
}

// NotificationContentClass is the class instance for NotificationContent
var NotificationContentClass = objc.GetClass("UNNotificationContent")

// MutableNotificationContent represents mutable notification content
type MutableNotificationContent struct {
	NotificationContent
}

// MutableNotificationContentClass is the class instance for MutableNotificationContent
var MutableNotificationContentClass = objc.GetClass("UNMutableNotificationContent")

// NewMutableNotificationContent creates a new mutable notification content
func NewMutableNotificationContent() MutableNotificationContent {
	return MutableNotificationContent{NotificationContent{objc.Call[objc.Object](MutableNotificationContentClass, objc.Sel("alloc")).Send(objc.Sel("init"))}}
}

// SetTitle sets the title of the notification content
func (n MutableNotificationContent) SetTitle(title foundation.String) {
	n.Send(objc.Sel("setTitle:"), title)
}

// SetSubtitle sets the subtitle of the notification content
func (n MutableNotificationContent) SetSubtitle(subtitle foundation.String) {
	n.Send(objc.Sel("setSubtitle:"), subtitle)
}

// SetBody sets the body of the notification content
func (n MutableNotificationContent) SetBody(body foundation.String) {
	n.Send(objc.Sel("setBody:"), body)
}

// SetBadge sets the badge of the notification content
func (n MutableNotificationContent) SetBadge(badge foundation.Number) {
	n.Send(objc.Sel("setBadge:"), badge)
}

// SetSound sets the sound of the notification content
func (n MutableNotificationContent) SetSound(sound NotificationSound) {
	n.Send(objc.Sel("setSound:"), sound)
}

// SetUserInfo sets the user info of the notification content
func (n MutableNotificationContent) SetUserInfo(userInfo foundation.Dictionary) {
	n.Send(objc.Sel("setUserInfo:"), userInfo)
}

// NotificationTrigger represents a notification trigger
type NotificationTrigger struct {
	objc.Object
}

// NotificationTriggerClass is the class instance for NotificationTrigger
var NotificationTriggerClass = objc.GetClass("UNNotificationTrigger")

// TimeIntervalNotificationTrigger represents a time interval notification trigger
type TimeIntervalNotificationTrigger struct {
	NotificationTrigger
}

// TimeIntervalNotificationTriggerClass is the class instance for TimeIntervalNotificationTrigger
var TimeIntervalNotificationTriggerClass = objc.GetClass("UNTimeIntervalNotificationTrigger")

// TriggerWithTimeIntervalRepeats creates a time interval notification trigger
func TriggerWithTimeIntervalRepeats(timeInterval float64, repeats bool) TimeIntervalNotificationTrigger {
	return TimeIntervalNotificationTrigger{NotificationTrigger{objc.Call[objc.Object](TimeIntervalNotificationTriggerClass, objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)}}
}

// CalendarNotificationTrigger represents a calendar notification trigger
type CalendarNotificationTrigger struct {
	NotificationTrigger
}

// CalendarNotificationTriggerClass is the class instance for CalendarNotificationTrigger
var CalendarNotificationTriggerClass = objc.GetClass("UNCalendarNotificationTrigger")

// TriggerWithDateMatchingComponentsRepeats creates a calendar notification trigger
func TriggerWithDateMatchingComponentsRepeats(dateComponents foundation.DateComponents, repeats bool) CalendarNotificationTrigger {
	return CalendarNotificationTrigger{NotificationTrigger{objc.Call[objc.Object](CalendarNotificationTriggerClass, objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)}}
}

// LocationNotificationTrigger represents a location notification trigger
type LocationNotificationTrigger struct {
	NotificationTrigger
}

// LocationNotificationTriggerClass is the class instance for LocationNotificationTrigger
var LocationNotificationTriggerClass = objc.GetClass("UNLocationNotificationTrigger")

// PushNotificationTrigger represents a push notification trigger
type PushNotificationTrigger struct {
	NotificationTrigger
}

// PushNotificationTriggerClass is the class instance for PushNotificationTrigger
var PushNotificationTriggerClass = objc.GetClass("UNPushNotificationTrigger")

// NotificationSound represents a notification sound
type NotificationSound struct {
	objc.Object
}

// NotificationSoundClass is the class instance for NotificationSound
var NotificationSoundClass = objc.GetClass("UNNotificationSound")

// DefaultSound returns the default notification sound
func DefaultSound() NotificationSound {
	return NotificationSound{objc.Call[objc.Object](NotificationSoundClass, objc.Sel("defaultSound"))}
}

// SoundNamed returns a notification sound with the specified name
func SoundNamed(name foundation.String) NotificationSound {
	return NotificationSound{objc.Call[objc.Object](NotificationSoundClass, objc.Sel("soundNamed:"), name)}
}

// AuthorizationOptions represents options for notification authorization
type AuthorizationOptions uint

const (
	AuthorizationOptionBadge   = 1 << 0
	AuthorizationOptionSound   = 1 << 1
	AuthorizationOptionAlert   = 1 << 2
	AuthorizationOptionCarPlay = 1 << 3
)

// NotificationPresentationOptions represents options for notification presentation
type NotificationPresentationOptions uint

const (
	NotificationPresentationOptionBadge    = 1 << 0
	NotificationPresentationOptionSound    = 1 << 1
	NotificationPresentationOptionAlert    = 1 << 2
	NotificationPresentationOptionList     = 1 << 3
	NotificationPresentationOptionBanner   = 1 << 4
)