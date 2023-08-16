// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that manage your app’s life cycle and its interaction with common system services. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate?language=objc
type PApplicationDelegate interface {
	// optional
	ApplicationDidResignActive(notification foundation.Notification)
	HasApplicationDidResignActive() bool

	// optional
	ApplicationDidFinishLaunching(notification foundation.Notification)
	HasApplicationDidFinishLaunching() bool

	// optional
	ApplicationShouldTerminate(sender Application) ApplicationTerminateReply
	HasApplicationShouldTerminate() bool

	// optional
	ApplicationWillBecomeActive(notification foundation.Notification)
	HasApplicationWillBecomeActive() bool

	// optional
	ApplicationDockMenu(sender Application) IMenu
	HasApplicationDockMenu() bool

	// optional
	ApplicationWillFinishLaunching(notification foundation.Notification)
	HasApplicationWillFinishLaunching() bool

	// optional
	ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool
	HasApplicationShouldTerminateAfterLastWindowClosed() bool

	// optional
	ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification)
	HasApplicationProtectedDataWillBecomeUnavailable() bool

	// optional
	ApplicationDidHide(notification foundation.Notification)
	HasApplicationDidHide() bool

	// optional
	ApplicationWillResignActive(notification foundation.Notification)
	HasApplicationWillResignActive() bool

	// optional
	ApplicationDidBecomeActive(notification foundation.Notification)
	HasApplicationDidBecomeActive() bool

	// optional
	ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification)
	HasApplicationProtectedDataDidBecomeAvailable() bool

	// optional
	ApplicationWillHide(notification foundation.Notification)
	HasApplicationWillHide() bool

	// optional
	ApplicationWillUpdate(notification foundation.Notification)
	HasApplicationWillUpdate() bool

	// optional
	ApplicationShouldOpenUntitledFile(sender Application) bool
	HasApplicationShouldOpenUntitledFile() bool

	// optional
	ApplicationSupportsSecureRestorableState(app Application) bool
	HasApplicationSupportsSecureRestorableState() bool

	// optional
	ApplicationWillUnhide(notification foundation.Notification)
	HasApplicationWillUnhide() bool

	// optional
	ApplicationOpenUntitledFile(sender Application) bool
	HasApplicationOpenUntitledFile() bool

	// optional
	ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool
	HasApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool

	// optional
	ApplicationDidDecodeRestorableState(app Application, coder foundation.Coder)
	HasApplicationDidDecodeRestorableState() bool

	// optional
	ApplicationDidUpdate(notification foundation.Notification)
	HasApplicationDidUpdate() bool

	// optional
	ApplicationDidChangeOcclusionState(notification foundation.Notification)
	HasApplicationDidChangeOcclusionState() bool

	// optional
	ApplicationDidChangeScreenParameters(notification foundation.Notification)
	HasApplicationDidChangeScreenParameters() bool

	// optional
	ApplicationWillTerminate(notification foundation.Notification)
	HasApplicationWillTerminate() bool

	// optional
	ApplicationDidUnhide(notification foundation.Notification)
	HasApplicationDidUnhide() bool

	// optional
	ApplicationShouldHandleReopenHasVisibleWindows(sender Application, flag bool) bool
	HasApplicationShouldHandleReopenHasVisibleWindows() bool
}

// A delegate implementation builder for the [PApplicationDelegate] protocol.
type ApplicationDelegate struct {
	_ApplicationDidResignActive                           func(notification foundation.Notification)
	_ApplicationDidFinishLaunching                        func(notification foundation.Notification)
	_ApplicationShouldTerminate                           func(sender Application) ApplicationTerminateReply
	_ApplicationWillBecomeActive                          func(notification foundation.Notification)
	_ApplicationDockMenu                                  func(sender Application) IMenu
	_ApplicationWillFinishLaunching                       func(notification foundation.Notification)
	_ApplicationShouldTerminateAfterLastWindowClosed      func(sender Application) bool
	_ApplicationProtectedDataWillBecomeUnavailable        func(notification foundation.Notification)
	_ApplicationDidHide                                   func(notification foundation.Notification)
	_ApplicationWillResignActive                          func(notification foundation.Notification)
	_ApplicationDidBecomeActive                           func(notification foundation.Notification)
	_ApplicationProtectedDataDidBecomeAvailable           func(notification foundation.Notification)
	_ApplicationWillHide                                  func(notification foundation.Notification)
	_ApplicationWillUpdate                                func(notification foundation.Notification)
	_ApplicationShouldOpenUntitledFile                    func(sender Application) bool
	_ApplicationSupportsSecureRestorableState             func(app Application) bool
	_ApplicationWillUnhide                                func(notification foundation.Notification)
	_ApplicationOpenUntitledFile                          func(sender Application) bool
	_ApplicationShouldAutomaticallyLocalizeKeyEquivalents func(application Application) bool
	_ApplicationDidDecodeRestorableState                  func(app Application, coder foundation.Coder)
	_ApplicationDidUpdate                                 func(notification foundation.Notification)
	_ApplicationDidChangeOcclusionState                   func(notification foundation.Notification)
	_ApplicationDidChangeScreenParameters                 func(notification foundation.Notification)
	_ApplicationWillTerminate                             func(notification foundation.Notification)
	_ApplicationDidUnhide                                 func(notification foundation.Notification)
	_ApplicationShouldHandleReopenHasVisibleWindows       func(sender Application, flag bool) bool
}

func (di *ApplicationDelegate) HasApplicationDidResignActive() bool {
	return di._ApplicationDidResignActive != nil
}

// Tells the delegate that the app is no longer active and doesn’t have focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428636-applicationdidresignactive?language=objc
func (di *ApplicationDelegate) SetApplicationDidResignActive(f func(notification foundation.Notification)) {
	di._ApplicationDidResignActive = f
}

// Tells the delegate that the app is no longer active and doesn’t have focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428636-applicationdidresignactive?language=objc
func (di *ApplicationDelegate) ApplicationDidResignActive(notification foundation.Notification) {
	di._ApplicationDidResignActive(notification)
}
func (di *ApplicationDelegate) HasApplicationDidFinishLaunching() bool {
	return di._ApplicationDidFinishLaunching != nil
}

// Tells the delegate that the app’s initialization is complete but it hasn’t received its first event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428385-applicationdidfinishlaunching?language=objc
func (di *ApplicationDelegate) SetApplicationDidFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationDidFinishLaunching = f
}

// Tells the delegate that the app’s initialization is complete but it hasn’t received its first event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428385-applicationdidfinishlaunching?language=objc
func (di *ApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	di._ApplicationDidFinishLaunching(notification)
}
func (di *ApplicationDelegate) HasApplicationShouldTerminate() bool {
	return di._ApplicationShouldTerminate != nil
}

// Returns a value that indicates if the app should terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428642-applicationshouldterminate?language=objc
func (di *ApplicationDelegate) SetApplicationShouldTerminate(f func(sender Application) ApplicationTerminateReply) {
	di._ApplicationShouldTerminate = f
}

// Returns a value that indicates if the app should terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428642-applicationshouldterminate?language=objc
func (di *ApplicationDelegate) ApplicationShouldTerminate(sender Application) ApplicationTerminateReply {
	return di._ApplicationShouldTerminate(sender)
}
func (di *ApplicationDelegate) HasApplicationWillBecomeActive() bool {
	return di._ApplicationWillBecomeActive != nil
}

// Tells the delegate that the app is about to become active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428699-applicationwillbecomeactive?language=objc
func (di *ApplicationDelegate) SetApplicationWillBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationWillBecomeActive = f
}

// Tells the delegate that the app is about to become active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428699-applicationwillbecomeactive?language=objc
func (di *ApplicationDelegate) ApplicationWillBecomeActive(notification foundation.Notification) {
	di._ApplicationWillBecomeActive(notification)
}
func (di *ApplicationDelegate) HasApplicationDockMenu() bool {
	return di._ApplicationDockMenu != nil
}

// Returns the app’s dock menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428564-applicationdockmenu?language=objc
func (di *ApplicationDelegate) SetApplicationDockMenu(f func(sender Application) IMenu) {
	di._ApplicationDockMenu = f
}

// Returns the app’s dock menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428564-applicationdockmenu?language=objc
func (di *ApplicationDelegate) ApplicationDockMenu(sender Application) IMenu {
	return di._ApplicationDockMenu(sender)
}
func (di *ApplicationDelegate) HasApplicationWillFinishLaunching() bool {
	return di._ApplicationWillFinishLaunching != nil
}

// Tells the delegate that the app’s initialization is about to complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428623-applicationwillfinishlaunching?language=objc
func (di *ApplicationDelegate) SetApplicationWillFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationWillFinishLaunching = f
}

// Tells the delegate that the app’s initialization is about to complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428623-applicationwillfinishlaunching?language=objc
func (di *ApplicationDelegate) ApplicationWillFinishLaunching(notification foundation.Notification) {
	di._ApplicationWillFinishLaunching(notification)
}
func (di *ApplicationDelegate) HasApplicationShouldTerminateAfterLastWindowClosed() bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed != nil
}

// Returns a Boolean value that indicates if the app terminates once the last window closes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428381-applicationshouldterminateafterl?language=objc
func (di *ApplicationDelegate) SetApplicationShouldTerminateAfterLastWindowClosed(f func(sender Application) bool) {
	di._ApplicationShouldTerminateAfterLastWindowClosed = f
}

// Returns a Boolean value that indicates if the app terminates once the last window closes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428381-applicationshouldterminateafterl?language=objc
func (di *ApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed(sender)
}
func (di *ApplicationDelegate) HasApplicationProtectedDataWillBecomeUnavailable() bool {
	return di._ApplicationProtectedDataWillBecomeUnavailable != nil
}

// Tells the delegate that protected data is about to become unavailable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752994-applicationprotecteddatawillbeco?language=objc
func (di *ApplicationDelegate) SetApplicationProtectedDataWillBecomeUnavailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataWillBecomeUnavailable = f
}

// Tells the delegate that protected data is about to become unavailable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752994-applicationprotecteddatawillbeco?language=objc
func (di *ApplicationDelegate) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification) {
	di._ApplicationProtectedDataWillBecomeUnavailable(notification)
}
func (di *ApplicationDelegate) HasApplicationDidHide() bool {
	return di._ApplicationDidHide != nil
}

// Tells the delegate that the app is now hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428552-applicationdidhide?language=objc
func (di *ApplicationDelegate) SetApplicationDidHide(f func(notification foundation.Notification)) {
	di._ApplicationDidHide = f
}

// Tells the delegate that the app is now hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428552-applicationdidhide?language=objc
func (di *ApplicationDelegate) ApplicationDidHide(notification foundation.Notification) {
	di._ApplicationDidHide(notification)
}
func (di *ApplicationDelegate) HasApplicationWillResignActive() bool {
	return di._ApplicationWillResignActive != nil
}

// Tells the delegate that the app is about to become inactive and will lose focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428539-applicationwillresignactive?language=objc
func (di *ApplicationDelegate) SetApplicationWillResignActive(f func(notification foundation.Notification)) {
	di._ApplicationWillResignActive = f
}

// Tells the delegate that the app is about to become inactive and will lose focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428539-applicationwillresignactive?language=objc
func (di *ApplicationDelegate) ApplicationWillResignActive(notification foundation.Notification) {
	di._ApplicationWillResignActive(notification)
}
func (di *ApplicationDelegate) HasApplicationDidBecomeActive() bool {
	return di._ApplicationDidBecomeActive != nil
}

// Tells the delegate that the app is now active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428577-applicationdidbecomeactive?language=objc
func (di *ApplicationDelegate) SetApplicationDidBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationDidBecomeActive = f
}

// Tells the delegate that the app is now active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428577-applicationdidbecomeactive?language=objc
func (di *ApplicationDelegate) ApplicationDidBecomeActive(notification foundation.Notification) {
	di._ApplicationDidBecomeActive(notification)
}
func (di *ApplicationDelegate) HasApplicationProtectedDataDidBecomeAvailable() bool {
	return di._ApplicationProtectedDataDidBecomeAvailable != nil
}

// Tells the delegate that protected data is now available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752993-applicationprotecteddatadidbecom?language=objc
func (di *ApplicationDelegate) SetApplicationProtectedDataDidBecomeAvailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataDidBecomeAvailable = f
}

// Tells the delegate that protected data is now available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752993-applicationprotecteddatadidbecom?language=objc
func (di *ApplicationDelegate) ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification) {
	di._ApplicationProtectedDataDidBecomeAvailable(notification)
}
func (di *ApplicationDelegate) HasApplicationWillHide() bool {
	return di._ApplicationWillHide != nil
}

// Tells the delegate that the app is about to be hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428478-applicationwillhide?language=objc
func (di *ApplicationDelegate) SetApplicationWillHide(f func(notification foundation.Notification)) {
	di._ApplicationWillHide = f
}

// Tells the delegate that the app is about to be hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428478-applicationwillhide?language=objc
func (di *ApplicationDelegate) ApplicationWillHide(notification foundation.Notification) {
	di._ApplicationWillHide(notification)
}
func (di *ApplicationDelegate) HasApplicationWillUpdate() bool {
	return di._ApplicationWillUpdate != nil
}

// Tells the delegate that the app is about to update its windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428774-applicationwillupdate?language=objc
func (di *ApplicationDelegate) SetApplicationWillUpdate(f func(notification foundation.Notification)) {
	di._ApplicationWillUpdate = f
}

// Tells the delegate that the app is about to update its windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428774-applicationwillupdate?language=objc
func (di *ApplicationDelegate) ApplicationWillUpdate(notification foundation.Notification) {
	di._ApplicationWillUpdate(notification)
}
func (di *ApplicationDelegate) HasApplicationShouldOpenUntitledFile() bool {
	return di._ApplicationShouldOpenUntitledFile != nil
}

// Returns a Boolean value that indicates if the app can open an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428444-applicationshouldopenuntitledfil?language=objc
func (di *ApplicationDelegate) SetApplicationShouldOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationShouldOpenUntitledFile = f
}

// Returns a Boolean value that indicates if the app can open an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428444-applicationshouldopenuntitledfil?language=objc
func (di *ApplicationDelegate) ApplicationShouldOpenUntitledFile(sender Application) bool {
	return di._ApplicationShouldOpenUntitledFile(sender)
}
func (di *ApplicationDelegate) HasApplicationSupportsSecureRestorableState() bool {
	return di._ApplicationSupportsSecureRestorableState != nil
}

// Returns a Boolean value that indicates if the app supports secure state restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3762521-applicationsupportssecurerestora?language=objc
func (di *ApplicationDelegate) SetApplicationSupportsSecureRestorableState(f func(app Application) bool) {
	di._ApplicationSupportsSecureRestorableState = f
}

// Returns a Boolean value that indicates if the app supports secure state restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3762521-applicationsupportssecurerestora?language=objc
func (di *ApplicationDelegate) ApplicationSupportsSecureRestorableState(app Application) bool {
	return di._ApplicationSupportsSecureRestorableState(app)
}
func (di *ApplicationDelegate) HasApplicationWillUnhide() bool {
	return di._ApplicationWillUnhide != nil
}

// Tells the delegate that the app is about to become visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428585-applicationwillunhide?language=objc
func (di *ApplicationDelegate) SetApplicationWillUnhide(f func(notification foundation.Notification)) {
	di._ApplicationWillUnhide = f
}

// Tells the delegate that the app is about to become visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428585-applicationwillunhide?language=objc
func (di *ApplicationDelegate) ApplicationWillUnhide(notification foundation.Notification) {
	di._ApplicationWillUnhide(notification)
}
func (di *ApplicationDelegate) HasApplicationOpenUntitledFile() bool {
	return di._ApplicationOpenUntitledFile != nil
}

// Returns a Boolean value that indicates if the app opens an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428491-applicationopenuntitledfile?language=objc
func (di *ApplicationDelegate) SetApplicationOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationOpenUntitledFile = f
}

// Returns a Boolean value that indicates if the app opens an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428491-applicationopenuntitledfile?language=objc
func (di *ApplicationDelegate) ApplicationOpenUntitledFile(sender Application) bool {
	return di._ApplicationOpenUntitledFile(sender)
}
func (di *ApplicationDelegate) HasApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents != nil
}

// Returns a Boolean value that tells the system whether to remap menu shortcuts to support localized keyboards. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3787553-applicationshouldautomaticallylo?language=objc
func (di *ApplicationDelegate) SetApplicationShouldAutomaticallyLocalizeKeyEquivalents(f func(application Application) bool) {
	di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents = f
}

// Returns a Boolean value that tells the system whether to remap menu shortcuts to support localized keyboards. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3787553-applicationshouldautomaticallylo?language=objc
func (di *ApplicationDelegate) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application)
}
func (di *ApplicationDelegate) HasApplicationDidDecodeRestorableState() bool {
	return di._ApplicationDidDecodeRestorableState != nil
}

// Tells the delegate when the app finished decoding its restorable state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428693-application?language=objc
func (di *ApplicationDelegate) SetApplicationDidDecodeRestorableState(f func(app Application, coder foundation.Coder)) {
	di._ApplicationDidDecodeRestorableState = f
}

// Tells the delegate when the app finished decoding its restorable state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428693-application?language=objc
func (di *ApplicationDelegate) ApplicationDidDecodeRestorableState(app Application, coder foundation.Coder) {
	di._ApplicationDidDecodeRestorableState(app, coder)
}
func (di *ApplicationDelegate) HasApplicationDidUpdate() bool {
	return di._ApplicationDidUpdate != nil
}

// Tells the delegate that the app’s windows did update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428589-applicationdidupdate?language=objc
func (di *ApplicationDelegate) SetApplicationDidUpdate(f func(notification foundation.Notification)) {
	di._ApplicationDidUpdate = f
}

// Tells the delegate that the app’s windows did update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428589-applicationdidupdate?language=objc
func (di *ApplicationDelegate) ApplicationDidUpdate(notification foundation.Notification) {
	di._ApplicationDidUpdate(notification)
}
func (di *ApplicationDelegate) HasApplicationDidChangeOcclusionState() bool {
	return di._ApplicationDidChangeOcclusionState != nil
}

// Tells the delegate about changes to the app’s occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428362-applicationdidchangeocclusionsta?language=objc
func (di *ApplicationDelegate) SetApplicationDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeOcclusionState = f
}

// Tells the delegate about changes to the app’s occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428362-applicationdidchangeocclusionsta?language=objc
func (di *ApplicationDelegate) ApplicationDidChangeOcclusionState(notification foundation.Notification) {
	di._ApplicationDidChangeOcclusionState(notification)
}
func (di *ApplicationDelegate) HasApplicationDidChangeScreenParameters() bool {
	return di._ApplicationDidChangeScreenParameters != nil
}

// Tells the delegate about changes to the configuration of any attached displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428424-applicationdidchangescreenparame?language=objc
func (di *ApplicationDelegate) SetApplicationDidChangeScreenParameters(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeScreenParameters = f
}

// Tells the delegate about changes to the configuration of any attached displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428424-applicationdidchangescreenparame?language=objc
func (di *ApplicationDelegate) ApplicationDidChangeScreenParameters(notification foundation.Notification) {
	di._ApplicationDidChangeScreenParameters(notification)
}
func (di *ApplicationDelegate) HasApplicationWillTerminate() bool {
	return di._ApplicationWillTerminate != nil
}

// Tells the delegate that the app is about to terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428522-applicationwillterminate?language=objc
func (di *ApplicationDelegate) SetApplicationWillTerminate(f func(notification foundation.Notification)) {
	di._ApplicationWillTerminate = f
}

// Tells the delegate that the app is about to terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428522-applicationwillterminate?language=objc
func (di *ApplicationDelegate) ApplicationWillTerminate(notification foundation.Notification) {
	di._ApplicationWillTerminate(notification)
}
func (di *ApplicationDelegate) HasApplicationDidUnhide() bool {
	return di._ApplicationDidUnhide != nil
}

// Tells the delegate that the app is now visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428755-applicationdidunhide?language=objc
func (di *ApplicationDelegate) SetApplicationDidUnhide(f func(notification foundation.Notification)) {
	di._ApplicationDidUnhide = f
}

// Tells the delegate that the app is now visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428755-applicationdidunhide?language=objc
func (di *ApplicationDelegate) ApplicationDidUnhide(notification foundation.Notification) {
	di._ApplicationDidUnhide(notification)
}
func (di *ApplicationDelegate) HasApplicationShouldHandleReopenHasVisibleWindows() bool {
	return di._ApplicationShouldHandleReopenHasVisibleWindows != nil
}

// Returns a Boolean value that indicates if the app responds to reopen AppleEvents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428638-applicationshouldhandlereopen?language=objc
func (di *ApplicationDelegate) SetApplicationShouldHandleReopenHasVisibleWindows(f func(sender Application, flag bool) bool) {
	di._ApplicationShouldHandleReopenHasVisibleWindows = f
}

// Returns a Boolean value that indicates if the app responds to reopen AppleEvents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428638-applicationshouldhandlereopen?language=objc
func (di *ApplicationDelegate) ApplicationShouldHandleReopenHasVisibleWindows(sender Application, flag bool) bool {
	return di._ApplicationShouldHandleReopenHasVisibleWindows(sender, flag)
}

// A concrete type wrapper for the [PApplicationDelegate] protocol.
type ApplicationDelegateWrapper struct {
	objc.Object
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidResignActive() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidResignActive:"))
}

// Tells the delegate that the app is no longer active and doesn’t have focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428636-applicationdidresignactive?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidResignActive(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidResignActive:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidFinishLaunching() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidFinishLaunching:"))
}

// Tells the delegate that the app’s initialization is complete but it hasn’t received its first event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428385-applicationdidfinishlaunching?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidFinishLaunching(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidFinishLaunching:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationShouldTerminate() bool {
	return a_.RespondsToSelector(objc.Sel("applicationShouldTerminate:"))
}

// Returns a value that indicates if the app should terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428642-applicationshouldterminate?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	rv := objc.Call[ApplicationTerminateReply](a_, objc.Sel("applicationShouldTerminate:"), objc.Ptr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillBecomeActive() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillBecomeActive:"))
}

// Tells the delegate that the app is about to become active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428699-applicationwillbecomeactive?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillBecomeActive(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillBecomeActive:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDockMenu() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDockMenu:"))
}

// Returns the app’s dock menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428564-applicationdockmenu?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDockMenu(sender IApplication) Menu {
	rv := objc.Call[Menu](a_, objc.Sel("applicationDockMenu:"), objc.Ptr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillFinishLaunching() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillFinishLaunching:"))
}

// Tells the delegate that the app’s initialization is about to complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428623-applicationwillfinishlaunching?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillFinishLaunching(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillFinishLaunching:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationShouldTerminateAfterLastWindowClosed() bool {
	return a_.RespondsToSelector(objc.Sel("applicationShouldTerminateAfterLastWindowClosed:"))
}

// Returns a Boolean value that indicates if the app terminates once the last window closes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428381-applicationshouldterminateafterl?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationShouldTerminateAfterLastWindowClosed:"), objc.Ptr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationProtectedDataWillBecomeUnavailable() bool {
	return a_.RespondsToSelector(objc.Sel("applicationProtectedDataWillBecomeUnavailable:"))
}

// Tells the delegate that protected data is about to become unavailable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752994-applicationprotecteddatawillbeco?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationProtectedDataWillBecomeUnavailable:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidHide() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidHide:"))
}

// Tells the delegate that the app is now hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428552-applicationdidhide?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidHide(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidHide:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillResignActive() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillResignActive:"))
}

// Tells the delegate that the app is about to become inactive and will lose focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428539-applicationwillresignactive?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillResignActive(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillResignActive:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidBecomeActive() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidBecomeActive:"))
}

// Tells the delegate that the app is now active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428577-applicationdidbecomeactive?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidBecomeActive(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidBecomeActive:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationProtectedDataDidBecomeAvailable() bool {
	return a_.RespondsToSelector(objc.Sel("applicationProtectedDataDidBecomeAvailable:"))
}

// Tells the delegate that protected data is now available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3752993-applicationprotecteddatadidbecom?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataDidBecomeAvailable(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationProtectedDataDidBecomeAvailable:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillHide() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillHide:"))
}

// Tells the delegate that the app is about to be hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428478-applicationwillhide?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillHide(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillHide:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillUpdate() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillUpdate:"))
}

// Tells the delegate that the app is about to update its windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428774-applicationwillupdate?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillUpdate(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillUpdate:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationShouldOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.Sel("applicationShouldOpenUntitledFile:"))
}

// Returns a Boolean value that indicates if the app can open an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428444-applicationshouldopenuntitledfil?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationShouldOpenUntitledFile:"), objc.Ptr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationSupportsSecureRestorableState() bool {
	return a_.RespondsToSelector(objc.Sel("applicationSupportsSecureRestorableState:"))
}

// Returns a Boolean value that indicates if the app supports secure state restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3762521-applicationsupportssecurerestora?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationSupportsSecureRestorableState:"), objc.Ptr(app))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillUnhide() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillUnhide:"))
}

// Tells the delegate that the app is about to become visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428585-applicationwillunhide?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillUnhide(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillUnhide:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.Sel("applicationOpenUntitledFile:"))
}

// Returns a Boolean value that indicates if the app opens an untitled file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428491-applicationopenuntitledfile?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationOpenUntitledFile(sender IApplication) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationOpenUntitledFile:"), objc.Ptr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return a_.RespondsToSelector(objc.Sel("applicationShouldAutomaticallyLocalizeKeyEquivalents:"))
}

// Returns a Boolean value that tells the system whether to remap menu shortcuts to support localized keyboards. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/3787553-applicationshouldautomaticallylo?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), objc.Ptr(application))
	return rv
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidDecodeRestorableState() bool {
	return a_.RespondsToSelector(objc.Sel("application:didDecodeRestorableState:"))
}

// Tells the delegate when the app finished decoding its restorable state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428693-application?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidDecodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.Call[objc.Void](a_, objc.Sel("application:didDecodeRestorableState:"), objc.Ptr(app), objc.Ptr(coder))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidUpdate() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidUpdate:"))
}

// Tells the delegate that the app’s windows did update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428589-applicationdidupdate?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidUpdate(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidUpdate:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidChangeOcclusionState() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidChangeOcclusionState:"))
}

// Tells the delegate about changes to the app’s occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428362-applicationdidchangeocclusionsta?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidChangeOcclusionState(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidChangeOcclusionState:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidChangeScreenParameters() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidChangeScreenParameters:"))
}

// Tells the delegate about changes to the configuration of any attached displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428424-applicationdidchangescreenparame?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidChangeScreenParameters(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidChangeScreenParameters:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationWillTerminate() bool {
	return a_.RespondsToSelector(objc.Sel("applicationWillTerminate:"))
}

// Tells the delegate that the app is about to terminate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428522-applicationwillterminate?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationWillTerminate(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationWillTerminate:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationDidUnhide() bool {
	return a_.RespondsToSelector(objc.Sel("applicationDidUnhide:"))
}

// Tells the delegate that the app is now visible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428755-applicationdidunhide?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationDidUnhide(notification foundation.INotification) {
	objc.Call[objc.Void](a_, objc.Sel("applicationDidUnhide:"), objc.Ptr(notification))
}

func (a_ ApplicationDelegateWrapper) HasApplicationShouldHandleReopenHasVisibleWindows() bool {
	return a_.RespondsToSelector(objc.Sel("applicationShouldHandleReopen:hasVisibleWindows:"))
}

// Returns a Boolean value that indicates if the app responds to reopen AppleEvents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate/1428638-applicationshouldhandlereopen?language=objc
func (a_ ApplicationDelegateWrapper) ApplicationShouldHandleReopenHasVisibleWindows(sender IApplication, flag bool) bool {
	rv := objc.Call[bool](a_, objc.Sel("applicationShouldHandleReopen:hasVisibleWindows:"), objc.Ptr(sender), flag)
	return rv
}
