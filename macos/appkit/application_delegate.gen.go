// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IApplicationDelegate interface {
	ImplementsApplicationWillFinishLaunching() bool
	// optional
	ApplicationWillFinishLaunching(notification foundation.Notification)
	ImplementsApplicationDidFinishLaunching() bool
	// optional
	ApplicationDidFinishLaunching(notification foundation.Notification)
	ImplementsApplicationWillBecomeActive() bool
	// optional
	ApplicationWillBecomeActive(notification foundation.Notification)
	ImplementsApplicationDidBecomeActive() bool
	// optional
	ApplicationDidBecomeActive(notification foundation.Notification)
	ImplementsApplicationWillResignActive() bool
	// optional
	ApplicationWillResignActive(notification foundation.Notification)
	ImplementsApplicationDidResignActive() bool
	// optional
	ApplicationDidResignActive(notification foundation.Notification)
	ImplementsApplicationShouldTerminate() bool
	// optional
	ApplicationShouldTerminate(sender Application) ApplicationTerminateReply
	ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool
	// optional
	ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool
	ImplementsApplicationWillTerminate() bool
	// optional
	ApplicationWillTerminate(notification foundation.Notification)
	ImplementsApplicationWillHide() bool
	// optional
	ApplicationWillHide(notification foundation.Notification)
	ImplementsApplicationDidHide() bool
	// optional
	ApplicationDidHide(notification foundation.Notification)
	ImplementsApplicationWillUnhide() bool
	// optional
	ApplicationWillUnhide(notification foundation.Notification)
	ImplementsApplicationDidUnhide() bool
	// optional
	ApplicationDidUnhide(notification foundation.Notification)
	ImplementsApplicationWillUpdate() bool
	// optional
	ApplicationWillUpdate(notification foundation.Notification)
	ImplementsApplicationDidUpdate() bool
	// optional
	ApplicationDidUpdate(notification foundation.Notification)
	ImplementsApplicationShouldHandleReopenHasVisibleWindows() bool
	// optional
	ApplicationShouldHandleReopenHasVisibleWindows(sender Application, flag bool) bool
	ImplementsApplicationDockMenu() bool
	// optional
	ApplicationDockMenu(sender Application) IMenu
	ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool
	// optional
	ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool
	ImplementsApplicationWillPresentError() bool
	// optional
	ApplicationWillPresentError(application Application, error foundation.Error) foundation.IError
	ImplementsApplicationDidChangeScreenParameters() bool
	// optional
	ApplicationDidChangeScreenParameters(notification foundation.Notification)
	ImplementsApplicationWillContinueUserActivityWithType() bool
	// optional
	ApplicationWillContinueUserActivityWithType(application Application, userActivityType string) bool
	ImplementsApplicationDidFailToContinueUserActivityWithTypeError() bool
	// optional
	ApplicationDidFailToContinueUserActivityWithTypeError(application Application, userActivityType string, error foundation.Error)
	ImplementsApplicationDidUpdateUserActivity() bool
	// optional
	ApplicationDidUpdateUserActivity(application Application, userActivity foundation.UserActivity)
	ImplementsApplicationDidRegisterForRemoteNotificationsWithDeviceToken() bool
	// optional
	ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application Application, deviceToken []byte)
	ImplementsApplicationDidFailToRegisterForRemoteNotificationsWithError() bool
	// optional
	ApplicationDidFailToRegisterForRemoteNotificationsWithError(application Application, error foundation.Error)
	ImplementsApplicationDidReceiveRemoteNotification() bool
	// optional
	ApplicationDidReceiveRemoteNotification(application Application, userInfo map[string]objc.Object)
	ImplementsApplicationOpenURLs() bool
	// optional
	ApplicationOpenURLs(application Application, urls []foundation.URL)
	ImplementsApplicationOpenFile() bool
	// optional
	ApplicationOpenFile(sender Application, filename string) bool
	ImplementsApplicationOpenFileWithoutUI() bool
	// optional
	ApplicationOpenFileWithoutUI(sender objc.Object, filename string) bool
	ImplementsApplicationOpenTempFile() bool
	// optional
	ApplicationOpenTempFile(sender Application, filename string) bool
	ImplementsApplicationOpenFiles() bool
	// optional
	ApplicationOpenFiles(sender Application, filenames []string)
	ImplementsApplicationShouldOpenUntitledFile() bool
	// optional
	ApplicationShouldOpenUntitledFile(sender Application) bool
	ImplementsApplicationOpenUntitledFile() bool
	// optional
	ApplicationOpenUntitledFile(sender Application) bool
	ImplementsApplicationPrintFile() bool
	// optional
	ApplicationPrintFile(sender Application, filename string) bool
	ImplementsApplicationPrintFilesWithSettingsShowPrintPanels() bool
	// optional
	ApplicationPrintFilesWithSettingsShowPrintPanels(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply
	ImplementsApplicationSupportsSecureRestorableState() bool
	// optional
	ApplicationSupportsSecureRestorableState(app Application) bool
	ImplementsApplicationProtectedDataDidBecomeAvailable() bool
	// optional
	ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification)
	ImplementsApplicationProtectedDataWillBecomeUnavailable() bool
	// optional
	ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification)
	ImplementsApplicationWillEncodeRestorableState() bool
	// optional
	ApplicationWillEncodeRestorableState(app Application, coder foundation.Coder)
	ImplementsApplicationDidDecodeRestorableState() bool
	// optional
	ApplicationDidDecodeRestorableState(app Application, coder foundation.Coder)
	ImplementsApplicationDidChangeOcclusionState() bool
	// optional
	ApplicationDidChangeOcclusionState(notification foundation.Notification)
	ImplementsApplicationDelegateHandlesKey() bool
	// optional
	ApplicationDelegateHandlesKey(sender Application, key string) bool
}

type ApplicationDelegate struct {
	_ApplicationWillFinishLaunching                              func(notification foundation.Notification)
	_ApplicationDidFinishLaunching                               func(notification foundation.Notification)
	_ApplicationWillBecomeActive                                 func(notification foundation.Notification)
	_ApplicationDidBecomeActive                                  func(notification foundation.Notification)
	_ApplicationWillResignActive                                 func(notification foundation.Notification)
	_ApplicationDidResignActive                                  func(notification foundation.Notification)
	_ApplicationShouldTerminate                                  func(sender Application) ApplicationTerminateReply
	_ApplicationShouldTerminateAfterLastWindowClosed             func(sender Application) bool
	_ApplicationWillTerminate                                    func(notification foundation.Notification)
	_ApplicationWillHide                                         func(notification foundation.Notification)
	_ApplicationDidHide                                          func(notification foundation.Notification)
	_ApplicationWillUnhide                                       func(notification foundation.Notification)
	_ApplicationDidUnhide                                        func(notification foundation.Notification)
	_ApplicationWillUpdate                                       func(notification foundation.Notification)
	_ApplicationDidUpdate                                        func(notification foundation.Notification)
	_ApplicationShouldHandleReopenHasVisibleWindows              func(sender Application, flag bool) bool
	_ApplicationDockMenu                                         func(sender Application) IMenu
	_ApplicationShouldAutomaticallyLocalizeKeyEquivalents        func(application Application) bool
	_ApplicationWillPresentError                                 func(application Application, error foundation.Error) foundation.IError
	_ApplicationDidChangeScreenParameters                        func(notification foundation.Notification)
	_ApplicationWillContinueUserActivityWithType                 func(application Application, userActivityType string) bool
	_ApplicationDidFailToContinueUserActivityWithTypeError       func(application Application, userActivityType string, error foundation.Error)
	_ApplicationDidUpdateUserActivity                            func(application Application, userActivity foundation.UserActivity)
	_ApplicationDidRegisterForRemoteNotificationsWithDeviceToken func(application Application, deviceToken []byte)
	_ApplicationDidFailToRegisterForRemoteNotificationsWithError func(application Application, error foundation.Error)
	_ApplicationDidReceiveRemoteNotification                     func(application Application, userInfo map[string]objc.Object)
	_ApplicationOpenURLs                                         func(application Application, urls []foundation.URL)
	_ApplicationOpenFile                                         func(sender Application, filename string) bool
	_ApplicationOpenFileWithoutUI                                func(sender objc.Object, filename string) bool
	_ApplicationOpenTempFile                                     func(sender Application, filename string) bool
	_ApplicationOpenFiles                                        func(sender Application, filenames []string)
	_ApplicationShouldOpenUntitledFile                           func(sender Application) bool
	_ApplicationOpenUntitledFile                                 func(sender Application) bool
	_ApplicationPrintFile                                        func(sender Application, filename string) bool
	_ApplicationPrintFilesWithSettingsShowPrintPanels            func(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply
	_ApplicationSupportsSecureRestorableState                    func(app Application) bool
	_ApplicationProtectedDataDidBecomeAvailable                  func(notification foundation.Notification)
	_ApplicationProtectedDataWillBecomeUnavailable               func(notification foundation.Notification)
	_ApplicationWillEncodeRestorableState                        func(app Application, coder foundation.Coder)
	_ApplicationDidDecodeRestorableState                         func(app Application, coder foundation.Coder)
	_ApplicationDidChangeOcclusionState                          func(notification foundation.Notification)
	_ApplicationDelegateHandlesKey                               func(sender Application, key string) bool
}

func (di *ApplicationDelegate) ImplementsApplicationWillFinishLaunching() bool {
	return di._ApplicationWillFinishLaunching != nil
}

func (di *ApplicationDelegate) SetApplicationWillFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationWillFinishLaunching = f
}

func (di *ApplicationDelegate) ApplicationWillFinishLaunching(notification foundation.Notification) {
	di._ApplicationWillFinishLaunching(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidFinishLaunching() bool {
	return di._ApplicationDidFinishLaunching != nil
}

func (di *ApplicationDelegate) SetApplicationDidFinishLaunching(f func(notification foundation.Notification)) {
	di._ApplicationDidFinishLaunching = f
}

func (di *ApplicationDelegate) ApplicationDidFinishLaunching(notification foundation.Notification) {
	di._ApplicationDidFinishLaunching(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillBecomeActive() bool {
	return di._ApplicationWillBecomeActive != nil
}

func (di *ApplicationDelegate) SetApplicationWillBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationWillBecomeActive = f
}

func (di *ApplicationDelegate) ApplicationWillBecomeActive(notification foundation.Notification) {
	di._ApplicationWillBecomeActive(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidBecomeActive() bool {
	return di._ApplicationDidBecomeActive != nil
}

func (di *ApplicationDelegate) SetApplicationDidBecomeActive(f func(notification foundation.Notification)) {
	di._ApplicationDidBecomeActive = f
}

func (di *ApplicationDelegate) ApplicationDidBecomeActive(notification foundation.Notification) {
	di._ApplicationDidBecomeActive(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillResignActive() bool {
	return di._ApplicationWillResignActive != nil
}

func (di *ApplicationDelegate) SetApplicationWillResignActive(f func(notification foundation.Notification)) {
	di._ApplicationWillResignActive = f
}

func (di *ApplicationDelegate) ApplicationWillResignActive(notification foundation.Notification) {
	di._ApplicationWillResignActive(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidResignActive() bool {
	return di._ApplicationDidResignActive != nil
}

func (di *ApplicationDelegate) SetApplicationDidResignActive(f func(notification foundation.Notification)) {
	di._ApplicationDidResignActive = f
}

func (di *ApplicationDelegate) ApplicationDidResignActive(notification foundation.Notification) {
	di._ApplicationDidResignActive(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationShouldTerminate() bool {
	return di._ApplicationShouldTerminate != nil
}

func (di *ApplicationDelegate) SetApplicationShouldTerminate(f func(sender Application) ApplicationTerminateReply) {
	di._ApplicationShouldTerminate = f
}

func (di *ApplicationDelegate) ApplicationShouldTerminate(sender Application) ApplicationTerminateReply {
	return di._ApplicationShouldTerminate(sender)
}
func (di *ApplicationDelegate) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed != nil
}

func (di *ApplicationDelegate) SetApplicationShouldTerminateAfterLastWindowClosed(f func(sender Application) bool) {
	di._ApplicationShouldTerminateAfterLastWindowClosed = f
}

func (di *ApplicationDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender Application) bool {
	return di._ApplicationShouldTerminateAfterLastWindowClosed(sender)
}
func (di *ApplicationDelegate) ImplementsApplicationWillTerminate() bool {
	return di._ApplicationWillTerminate != nil
}

func (di *ApplicationDelegate) SetApplicationWillTerminate(f func(notification foundation.Notification)) {
	di._ApplicationWillTerminate = f
}

func (di *ApplicationDelegate) ApplicationWillTerminate(notification foundation.Notification) {
	di._ApplicationWillTerminate(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillHide() bool {
	return di._ApplicationWillHide != nil
}

func (di *ApplicationDelegate) SetApplicationWillHide(f func(notification foundation.Notification)) {
	di._ApplicationWillHide = f
}

func (di *ApplicationDelegate) ApplicationWillHide(notification foundation.Notification) {
	di._ApplicationWillHide(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidHide() bool {
	return di._ApplicationDidHide != nil
}

func (di *ApplicationDelegate) SetApplicationDidHide(f func(notification foundation.Notification)) {
	di._ApplicationDidHide = f
}

func (di *ApplicationDelegate) ApplicationDidHide(notification foundation.Notification) {
	di._ApplicationDidHide(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillUnhide() bool {
	return di._ApplicationWillUnhide != nil
}

func (di *ApplicationDelegate) SetApplicationWillUnhide(f func(notification foundation.Notification)) {
	di._ApplicationWillUnhide = f
}

func (di *ApplicationDelegate) ApplicationWillUnhide(notification foundation.Notification) {
	di._ApplicationWillUnhide(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidUnhide() bool {
	return di._ApplicationDidUnhide != nil
}

func (di *ApplicationDelegate) SetApplicationDidUnhide(f func(notification foundation.Notification)) {
	di._ApplicationDidUnhide = f
}

func (di *ApplicationDelegate) ApplicationDidUnhide(notification foundation.Notification) {
	di._ApplicationDidUnhide(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillUpdate() bool {
	return di._ApplicationWillUpdate != nil
}

func (di *ApplicationDelegate) SetApplicationWillUpdate(f func(notification foundation.Notification)) {
	di._ApplicationWillUpdate = f
}

func (di *ApplicationDelegate) ApplicationWillUpdate(notification foundation.Notification) {
	di._ApplicationWillUpdate(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDidUpdate() bool {
	return di._ApplicationDidUpdate != nil
}

func (di *ApplicationDelegate) SetApplicationDidUpdate(f func(notification foundation.Notification)) {
	di._ApplicationDidUpdate = f
}

func (di *ApplicationDelegate) ApplicationDidUpdate(notification foundation.Notification) {
	di._ApplicationDidUpdate(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationShouldHandleReopenHasVisibleWindows() bool {
	return di._ApplicationShouldHandleReopenHasVisibleWindows != nil
}

func (di *ApplicationDelegate) SetApplicationShouldHandleReopenHasVisibleWindows(f func(sender Application, flag bool) bool) {
	di._ApplicationShouldHandleReopenHasVisibleWindows = f
}

func (di *ApplicationDelegate) ApplicationShouldHandleReopenHasVisibleWindows(sender Application, flag bool) bool {
	return di._ApplicationShouldHandleReopenHasVisibleWindows(sender, flag)
}
func (di *ApplicationDelegate) ImplementsApplicationDockMenu() bool {
	return di._ApplicationDockMenu != nil
}

func (di *ApplicationDelegate) SetApplicationDockMenu(f func(sender Application) IMenu) {
	di._ApplicationDockMenu = f
}

func (di *ApplicationDelegate) ApplicationDockMenu(sender Application) IMenu {
	return di._ApplicationDockMenu(sender)
}
func (di *ApplicationDelegate) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents != nil
}

func (di *ApplicationDelegate) SetApplicationShouldAutomaticallyLocalizeKeyEquivalents(f func(application Application) bool) {
	di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents = f
}

func (di *ApplicationDelegate) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application Application) bool {
	return di._ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application)
}
func (di *ApplicationDelegate) ImplementsApplicationWillPresentError() bool {
	return di._ApplicationWillPresentError != nil
}

func (di *ApplicationDelegate) SetApplicationWillPresentError(f func(application Application, error foundation.Error) foundation.IError) {
	di._ApplicationWillPresentError = f
}

func (di *ApplicationDelegate) ApplicationWillPresentError(application Application, error foundation.Error) foundation.IError {
	return di._ApplicationWillPresentError(application, error)
}
func (di *ApplicationDelegate) ImplementsApplicationDidChangeScreenParameters() bool {
	return di._ApplicationDidChangeScreenParameters != nil
}

func (di *ApplicationDelegate) SetApplicationDidChangeScreenParameters(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeScreenParameters = f
}

func (di *ApplicationDelegate) ApplicationDidChangeScreenParameters(notification foundation.Notification) {
	di._ApplicationDidChangeScreenParameters(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillContinueUserActivityWithType() bool {
	return di._ApplicationWillContinueUserActivityWithType != nil
}

func (di *ApplicationDelegate) SetApplicationWillContinueUserActivityWithType(f func(application Application, userActivityType string) bool) {
	di._ApplicationWillContinueUserActivityWithType = f
}

func (di *ApplicationDelegate) ApplicationWillContinueUserActivityWithType(application Application, userActivityType string) bool {
	return di._ApplicationWillContinueUserActivityWithType(application, userActivityType)
}
func (di *ApplicationDelegate) ImplementsApplicationDidFailToContinueUserActivityWithTypeError() bool {
	return di._ApplicationDidFailToContinueUserActivityWithTypeError != nil
}

func (di *ApplicationDelegate) SetApplicationDidFailToContinueUserActivityWithTypeError(f func(application Application, userActivityType string, error foundation.Error)) {
	di._ApplicationDidFailToContinueUserActivityWithTypeError = f
}

func (di *ApplicationDelegate) ApplicationDidFailToContinueUserActivityWithTypeError(application Application, userActivityType string, error foundation.Error) {
	di._ApplicationDidFailToContinueUserActivityWithTypeError(application, userActivityType, error)
}
func (di *ApplicationDelegate) ImplementsApplicationDidUpdateUserActivity() bool {
	return di._ApplicationDidUpdateUserActivity != nil
}

func (di *ApplicationDelegate) SetApplicationDidUpdateUserActivity(f func(application Application, userActivity foundation.UserActivity)) {
	di._ApplicationDidUpdateUserActivity = f
}

func (di *ApplicationDelegate) ApplicationDidUpdateUserActivity(application Application, userActivity foundation.UserActivity) {
	di._ApplicationDidUpdateUserActivity(application, userActivity)
}
func (di *ApplicationDelegate) ImplementsApplicationDidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return di._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken != nil
}

func (di *ApplicationDelegate) SetApplicationDidRegisterForRemoteNotificationsWithDeviceToken(f func(application Application, deviceToken []byte)) {
	di._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken = f
}

func (di *ApplicationDelegate) ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application Application, deviceToken []byte) {
	di._ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application, deviceToken)
}
func (di *ApplicationDelegate) ImplementsApplicationDidFailToRegisterForRemoteNotificationsWithError() bool {
	return di._ApplicationDidFailToRegisterForRemoteNotificationsWithError != nil
}

func (di *ApplicationDelegate) SetApplicationDidFailToRegisterForRemoteNotificationsWithError(f func(application Application, error foundation.Error)) {
	di._ApplicationDidFailToRegisterForRemoteNotificationsWithError = f
}

func (di *ApplicationDelegate) ApplicationDidFailToRegisterForRemoteNotificationsWithError(application Application, error foundation.Error) {
	di._ApplicationDidFailToRegisterForRemoteNotificationsWithError(application, error)
}
func (di *ApplicationDelegate) ImplementsApplicationDidReceiveRemoteNotification() bool {
	return di._ApplicationDidReceiveRemoteNotification != nil
}

func (di *ApplicationDelegate) SetApplicationDidReceiveRemoteNotification(f func(application Application, userInfo map[string]objc.Object)) {
	di._ApplicationDidReceiveRemoteNotification = f
}

func (di *ApplicationDelegate) ApplicationDidReceiveRemoteNotification(application Application, userInfo map[string]objc.Object) {
	di._ApplicationDidReceiveRemoteNotification(application, userInfo)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenURLs() bool {
	return di._ApplicationOpenURLs != nil
}

func (di *ApplicationDelegate) SetApplicationOpenURLs(f func(application Application, urls []foundation.URL)) {
	di._ApplicationOpenURLs = f
}

func (di *ApplicationDelegate) ApplicationOpenURLs(application Application, urls []foundation.URL) {
	di._ApplicationOpenURLs(application, urls)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenFile() bool {
	return di._ApplicationOpenFile != nil
}

func (di *ApplicationDelegate) SetApplicationOpenFile(f func(sender Application, filename string) bool) {
	di._ApplicationOpenFile = f
}

func (di *ApplicationDelegate) ApplicationOpenFile(sender Application, filename string) bool {
	return di._ApplicationOpenFile(sender, filename)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenFileWithoutUI() bool {
	return di._ApplicationOpenFileWithoutUI != nil
}

func (di *ApplicationDelegate) SetApplicationOpenFileWithoutUI(f func(sender objc.Object, filename string) bool) {
	di._ApplicationOpenFileWithoutUI = f
}

func (di *ApplicationDelegate) ApplicationOpenFileWithoutUI(sender objc.Object, filename string) bool {
	return di._ApplicationOpenFileWithoutUI(sender, filename)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenTempFile() bool {
	return di._ApplicationOpenTempFile != nil
}

func (di *ApplicationDelegate) SetApplicationOpenTempFile(f func(sender Application, filename string) bool) {
	di._ApplicationOpenTempFile = f
}

func (di *ApplicationDelegate) ApplicationOpenTempFile(sender Application, filename string) bool {
	return di._ApplicationOpenTempFile(sender, filename)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenFiles() bool {
	return di._ApplicationOpenFiles != nil
}

func (di *ApplicationDelegate) SetApplicationOpenFiles(f func(sender Application, filenames []string)) {
	di._ApplicationOpenFiles = f
}

func (di *ApplicationDelegate) ApplicationOpenFiles(sender Application, filenames []string) {
	di._ApplicationOpenFiles(sender, filenames)
}
func (di *ApplicationDelegate) ImplementsApplicationShouldOpenUntitledFile() bool {
	return di._ApplicationShouldOpenUntitledFile != nil
}

func (di *ApplicationDelegate) SetApplicationShouldOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationShouldOpenUntitledFile = f
}

func (di *ApplicationDelegate) ApplicationShouldOpenUntitledFile(sender Application) bool {
	return di._ApplicationShouldOpenUntitledFile(sender)
}
func (di *ApplicationDelegate) ImplementsApplicationOpenUntitledFile() bool {
	return di._ApplicationOpenUntitledFile != nil
}

func (di *ApplicationDelegate) SetApplicationOpenUntitledFile(f func(sender Application) bool) {
	di._ApplicationOpenUntitledFile = f
}

func (di *ApplicationDelegate) ApplicationOpenUntitledFile(sender Application) bool {
	return di._ApplicationOpenUntitledFile(sender)
}
func (di *ApplicationDelegate) ImplementsApplicationPrintFile() bool {
	return di._ApplicationPrintFile != nil
}

func (di *ApplicationDelegate) SetApplicationPrintFile(f func(sender Application, filename string) bool) {
	di._ApplicationPrintFile = f
}

func (di *ApplicationDelegate) ApplicationPrintFile(sender Application, filename string) bool {
	return di._ApplicationPrintFile(sender, filename)
}
func (di *ApplicationDelegate) ImplementsApplicationPrintFilesWithSettingsShowPrintPanels() bool {
	return di._ApplicationPrintFilesWithSettingsShowPrintPanels != nil
}

func (di *ApplicationDelegate) SetApplicationPrintFilesWithSettingsShowPrintPanels(f func(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply) {
	di._ApplicationPrintFilesWithSettingsShowPrintPanels = f
}

func (di *ApplicationDelegate) ApplicationPrintFilesWithSettingsShowPrintPanels(application Application, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.Object, showPrintPanels bool) ApplicationPrintReply {
	return di._ApplicationPrintFilesWithSettingsShowPrintPanels(application, fileNames, printSettings, showPrintPanels)
}
func (di *ApplicationDelegate) ImplementsApplicationSupportsSecureRestorableState() bool {
	return di._ApplicationSupportsSecureRestorableState != nil
}

func (di *ApplicationDelegate) SetApplicationSupportsSecureRestorableState(f func(app Application) bool) {
	di._ApplicationSupportsSecureRestorableState = f
}

func (di *ApplicationDelegate) ApplicationSupportsSecureRestorableState(app Application) bool {
	return di._ApplicationSupportsSecureRestorableState(app)
}
func (di *ApplicationDelegate) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return di._ApplicationProtectedDataDidBecomeAvailable != nil
}

func (di *ApplicationDelegate) SetApplicationProtectedDataDidBecomeAvailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataDidBecomeAvailable = f
}

func (di *ApplicationDelegate) ApplicationProtectedDataDidBecomeAvailable(notification foundation.Notification) {
	di._ApplicationProtectedDataDidBecomeAvailable(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return di._ApplicationProtectedDataWillBecomeUnavailable != nil
}

func (di *ApplicationDelegate) SetApplicationProtectedDataWillBecomeUnavailable(f func(notification foundation.Notification)) {
	di._ApplicationProtectedDataWillBecomeUnavailable = f
}

func (di *ApplicationDelegate) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.Notification) {
	di._ApplicationProtectedDataWillBecomeUnavailable(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationWillEncodeRestorableState() bool {
	return di._ApplicationWillEncodeRestorableState != nil
}

func (di *ApplicationDelegate) SetApplicationWillEncodeRestorableState(f func(app Application, coder foundation.Coder)) {
	di._ApplicationWillEncodeRestorableState = f
}

func (di *ApplicationDelegate) ApplicationWillEncodeRestorableState(app Application, coder foundation.Coder) {
	di._ApplicationWillEncodeRestorableState(app, coder)
}
func (di *ApplicationDelegate) ImplementsApplicationDidDecodeRestorableState() bool {
	return di._ApplicationDidDecodeRestorableState != nil
}

func (di *ApplicationDelegate) SetApplicationDidDecodeRestorableState(f func(app Application, coder foundation.Coder)) {
	di._ApplicationDidDecodeRestorableState = f
}

func (di *ApplicationDelegate) ApplicationDidDecodeRestorableState(app Application, coder foundation.Coder) {
	di._ApplicationDidDecodeRestorableState(app, coder)
}
func (di *ApplicationDelegate) ImplementsApplicationDidChangeOcclusionState() bool {
	return di._ApplicationDidChangeOcclusionState != nil
}

func (di *ApplicationDelegate) SetApplicationDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._ApplicationDidChangeOcclusionState = f
}

func (di *ApplicationDelegate) ApplicationDidChangeOcclusionState(notification foundation.Notification) {
	di._ApplicationDidChangeOcclusionState(notification)
}
func (di *ApplicationDelegate) ImplementsApplicationDelegateHandlesKey() bool {
	return di._ApplicationDelegateHandlesKey != nil
}

func (di *ApplicationDelegate) SetApplicationDelegateHandlesKey(f func(sender Application, key string) bool) {
	di._ApplicationDelegateHandlesKey = f
}

func (di *ApplicationDelegate) ApplicationDelegateHandlesKey(sender Application, key string) bool {
	return di._ApplicationDelegateHandlesKey(sender, key)
}

type ApplicationDelegateWrapper struct {
	objc.Object
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillFinishLaunching() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillFinishLaunching:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillFinishLaunching(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidFinishLaunching() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidFinishLaunching:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFinishLaunching(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidFinishLaunching:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidBecomeActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidBecomeActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidBecomeActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidBecomeActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillResignActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidResignActive() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidResignActive:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidResignActive(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidResignActive:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationShouldTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminate(sender IApplication) ApplicationTerminateReply {
	rv := objc.CallMethod[ApplicationTerminateReply](a_, objc.GetSelector("applicationShouldTerminate:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationShouldTerminateAfterLastWindowClosed() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldTerminateAfterLastWindowClosed(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldTerminateAfterLastWindowClosed:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillTerminate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillTerminate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillTerminate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillTerminate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillHide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidHide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidHide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidHide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidHide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUnhide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidUnhide() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUnhide:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUnhide(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUnhide:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationWillUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationWillUpdate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidUpdate() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidUpdate:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidUpdate:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationShouldHandleReopenHasVisibleWindows() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldHandleReopenHasVisibleWindows(sender IApplication, flag bool) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldHandleReopen:hasVisibleWindows:"), objc.ExtractPtr(sender), flag)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDockMenu() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDockMenu:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDockMenu(sender IApplication) Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("applicationDockMenu:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationShouldAutomaticallyLocalizeKeyEquivalents() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), objc.ExtractPtr(application))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillPresentError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willPresentError:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillPresentError(application IApplication, error foundation.IError) foundation.Error {
	rv := objc.CallMethod[foundation.Error](a_, objc.GetSelector("application:willPresentError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidChangeScreenParameters() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeScreenParameters:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeScreenParameters(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeScreenParameters:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillContinueUserActivityWithType() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willContinueUserActivityWithType:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillContinueUserActivityWithType(application IApplication, userActivityType string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:willContinueUserActivityWithType:"), objc.ExtractPtr(application), userActivityType)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidFailToContinueUserActivityWithTypeError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFailToContinueUserActivityWithTypeError(application IApplication, userActivityType string, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToContinueUserActivityWithType:error:"), objc.ExtractPtr(application), userActivityType, objc.ExtractPtr(error))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidUpdateUserActivity() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didUpdateUserActivity:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidUpdateUserActivity(application IApplication, userActivity foundation.IUserActivity) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didUpdateUserActivity:"), objc.ExtractPtr(application), objc.ExtractPtr(userActivity))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidRegisterForRemoteNotificationsWithDeviceToken() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application IApplication, deviceToken []byte) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didRegisterForRemoteNotificationsWithDeviceToken:"), objc.ExtractPtr(application), deviceToken)
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidFailToRegisterForRemoteNotificationsWithError() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidFailToRegisterForRemoteNotificationsWithError(application IApplication, error foundation.IError) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didFailToRegisterForRemoteNotificationsWithError:"), objc.ExtractPtr(application), objc.ExtractPtr(error))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidReceiveRemoteNotification() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didReceiveRemoteNotification:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidReceiveRemoteNotification(application IApplication, userInfo map[string]objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didReceiveRemoteNotification:"), objc.ExtractPtr(application), userInfo)
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenURLs() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openURLs:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenURLs(application IApplication, urls []foundation.IURL) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openURLs:"), objc.ExtractPtr(application), urls)
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenFileWithoutUI() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFileWithoutUI:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenFileWithoutUI(sender objc.IObject, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openFileWithoutUI:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenTempFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openTempFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenTempFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:openTempFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenFiles() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:openFiles:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenFiles(sender IApplication, filenames []string) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:openFiles:"), objc.ExtractPtr(sender), filenames)
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationShouldOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationShouldOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationShouldOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationShouldOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationOpenUntitledFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationOpenUntitledFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationOpenUntitledFile(sender IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationOpenUntitledFile:"), objc.ExtractPtr(sender))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationPrintFile() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFile:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationPrintFile(sender IApplication, filename string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:printFile:"), objc.ExtractPtr(sender), filename)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationPrintFilesWithSettingsShowPrintPanels() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationPrintFilesWithSettingsShowPrintPanels(application IApplication, fileNames []string, printSettings map[PrintInfoAttributeKey]objc.IObject, showPrintPanels bool) ApplicationPrintReply {
	rv := objc.CallMethod[ApplicationPrintReply](a_, objc.GetSelector("application:printFiles:withSettings:showPrintPanels:"), objc.ExtractPtr(application), fileNames, printSettings, showPrintPanels)
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationSupportsSecureRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationSupportsSecureRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationSupportsSecureRestorableState(app IApplication) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("applicationSupportsSecureRestorableState:"), objc.ExtractPtr(app))
	return rv
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationProtectedDataDidBecomeAvailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataDidBecomeAvailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataDidBecomeAvailable:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationProtectedDataWillBecomeUnavailable() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationProtectedDataWillBecomeUnavailable:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationWillEncodeRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:willEncodeRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationWillEncodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:willEncodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidDecodeRestorableState() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:didDecodeRestorableState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidDecodeRestorableState(app IApplication, coder foundation.ICoder) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("application:didDecodeRestorableState:"), objc.ExtractPtr(app), objc.ExtractPtr(coder))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDidChangeOcclusionState() bool {
	return a_.RespondsToSelector(objc.GetSelector("applicationDidChangeOcclusionState:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDidChangeOcclusionState(notification foundation.INotification) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("applicationDidChangeOcclusionState:"), objc.ExtractPtr(notification))
}

func (a_ ApplicationDelegateWrapper) ImplementsApplicationDelegateHandlesKey() bool {
	return a_.RespondsToSelector(objc.GetSelector("application:delegateHandlesKey:"))
}

func (a_ ApplicationDelegateWrapper) ApplicationDelegateHandlesKey(sender IApplication, key string) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("application:delegateHandlesKey:"), objc.ExtractPtr(sender), key)
	return rv
}
