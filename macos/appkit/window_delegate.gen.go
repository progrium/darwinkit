// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods that a window’s delegate can implement to respond to events, such as window resizing, moving, exposing, and minimizing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate?language=objc
type PWindowDelegate interface {
	// optional
	WindowDidFailToExitFullScreen(window Window)
	HasWindowDidFailToExitFullScreen() bool

	// optional
	WindowShouldZoomToFrame(window Window, newFrame foundation.Rect) bool
	HasWindowShouldZoomToFrame() bool

	// optional
	WindowDidEndSheet(notification foundation.Notification)
	HasWindowDidEndSheet() bool

	// optional
	WindowDidResize(notification foundation.Notification)
	HasWindowDidResize() bool

	// optional
	WindowWillEnterFullScreen(notification foundation.Notification)
	HasWindowWillEnterFullScreen() bool

	// optional
	WindowWillMove(notification foundation.Notification)
	HasWindowWillMove() bool

	// optional
	WindowWillStartLiveResize(notification foundation.Notification)
	HasWindowWillStartLiveResize() bool

	// optional
	WindowDidFailToEnterFullScreen(window Window)
	HasWindowDidFailToEnterFullScreen() bool

	// optional
	WindowWillResizeToSize(sender Window, frameSize foundation.Size) foundation.Size
	HasWindowWillResizeToSize() bool

	// optional
	WindowDidChangeOcclusionState(notification foundation.Notification)
	HasWindowDidChangeOcclusionState() bool

	// optional
	WindowDidChangeScreenProfile(notification foundation.Notification)
	HasWindowDidChangeScreenProfile() bool

	// optional
	WindowWillReturnUndoManager(window Window) foundation.IUndoManager
	HasWindowWillReturnUndoManager() bool

	// optional
	WindowWillBeginSheet(notification foundation.Notification)
	HasWindowWillBeginSheet() bool

	// optional
	WindowDidResignKey(notification foundation.Notification)
	HasWindowDidResignKey() bool

	// optional
	WindowShouldClose(sender Window) bool
	HasWindowShouldClose() bool

	// optional
	WindowDidExitVersionBrowser(notification foundation.Notification)
	HasWindowDidExitVersionBrowser() bool

	// optional
	WindowDidBecomeKey(notification foundation.Notification)
	HasWindowDidBecomeKey() bool

	// optional
	WindowDidResignMain(notification foundation.Notification)
	HasWindowDidResignMain() bool

	// optional
	WindowDidExitFullScreen(notification foundation.Notification)
	HasWindowDidExitFullScreen() bool

	// optional
	WindowDidBecomeMain(notification foundation.Notification)
	HasWindowDidBecomeMain() bool

	// optional
	WindowDidMiniaturize(notification foundation.Notification)
	HasWindowDidMiniaturize() bool

	// optional
	WindowDidExpose(notification foundation.Notification)
	HasWindowDidExpose() bool

	// optional
	WindowWillMiniaturize(notification foundation.Notification)
	HasWindowWillMiniaturize() bool

	// optional
	WindowDidChangeBackingProperties(notification foundation.Notification)
	HasWindowDidChangeBackingProperties() bool

	// optional
	WindowWillClose(notification foundation.Notification)
	HasWindowWillClose() bool

	// optional
	CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow
	HasCustomWindowsToEnterFullScreenForWindow() bool

	// optional
	WindowWillEnterVersionBrowser(notification foundation.Notification)
	HasWindowWillEnterVersionBrowser() bool

	// optional
	WindowDidEndLiveResize(notification foundation.Notification)
	HasWindowDidEndLiveResize() bool

	// optional
	WindowDidUpdate(notification foundation.Notification)
	HasWindowDidUpdate() bool

	// optional
	WindowWillExitVersionBrowser(notification foundation.Notification)
	HasWindowWillExitVersionBrowser() bool

	// optional
	WindowWillReturnFieldEditorToObject(sender Window, client objc.Object) objc.IObject
	HasWindowWillReturnFieldEditorToObject() bool

	// optional
	WindowDidDeminiaturize(notification foundation.Notification)
	HasWindowDidDeminiaturize() bool

	// optional
	CustomWindowsToExitFullScreenForWindow(window Window) []IWindow
	HasCustomWindowsToExitFullScreenForWindow() bool

	// optional
	WindowDidDecodeRestorableState(window Window, state foundation.Coder)
	HasWindowDidDecodeRestorableState() bool

	// optional
	WindowDidMove(notification foundation.Notification)
	HasWindowDidMove() bool

	// optional
	WindowDidChangeScreen(notification foundation.Notification)
	HasWindowDidChangeScreen() bool

	// optional
	WindowDidEnterFullScreen(notification foundation.Notification)
	HasWindowDidEnterFullScreen() bool

	// optional
	WindowWillExitFullScreen(notification foundation.Notification)
	HasWindowWillExitFullScreen() bool

	// optional
	WindowWillUseStandardFrameDefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect
	HasWindowWillUseStandardFrameDefaultFrame() bool

	// optional
	WindowDidEnterVersionBrowser(notification foundation.Notification)
	HasWindowDidEnterVersionBrowser() bool
}

// A delegate implementation builder for the [PWindowDelegate] protocol.
type WindowDelegate struct {
	_WindowDidFailToExitFullScreen           func(window Window)
	_WindowShouldZoomToFrame                 func(window Window, newFrame foundation.Rect) bool
	_WindowDidEndSheet                       func(notification foundation.Notification)
	_WindowDidResize                         func(notification foundation.Notification)
	_WindowWillEnterFullScreen               func(notification foundation.Notification)
	_WindowWillMove                          func(notification foundation.Notification)
	_WindowWillStartLiveResize               func(notification foundation.Notification)
	_WindowDidFailToEnterFullScreen          func(window Window)
	_WindowWillResizeToSize                  func(sender Window, frameSize foundation.Size) foundation.Size
	_WindowDidChangeOcclusionState           func(notification foundation.Notification)
	_WindowDidChangeScreenProfile            func(notification foundation.Notification)
	_WindowWillReturnUndoManager             func(window Window) foundation.IUndoManager
	_WindowWillBeginSheet                    func(notification foundation.Notification)
	_WindowDidResignKey                      func(notification foundation.Notification)
	_WindowShouldClose                       func(sender Window) bool
	_WindowDidExitVersionBrowser             func(notification foundation.Notification)
	_WindowDidBecomeKey                      func(notification foundation.Notification)
	_WindowDidResignMain                     func(notification foundation.Notification)
	_WindowDidExitFullScreen                 func(notification foundation.Notification)
	_WindowDidBecomeMain                     func(notification foundation.Notification)
	_WindowDidMiniaturize                    func(notification foundation.Notification)
	_WindowDidExpose                         func(notification foundation.Notification)
	_WindowWillMiniaturize                   func(notification foundation.Notification)
	_WindowDidChangeBackingProperties        func(notification foundation.Notification)
	_WindowWillClose                         func(notification foundation.Notification)
	_CustomWindowsToEnterFullScreenForWindow func(window Window) []IWindow
	_WindowWillEnterVersionBrowser           func(notification foundation.Notification)
	_WindowDidEndLiveResize                  func(notification foundation.Notification)
	_WindowDidUpdate                         func(notification foundation.Notification)
	_WindowWillExitVersionBrowser            func(notification foundation.Notification)
	_WindowWillReturnFieldEditorToObject     func(sender Window, client objc.Object) objc.IObject
	_WindowDidDeminiaturize                  func(notification foundation.Notification)
	_CustomWindowsToExitFullScreenForWindow  func(window Window) []IWindow
	_WindowDidDecodeRestorableState          func(window Window, state foundation.Coder)
	_WindowDidMove                           func(notification foundation.Notification)
	_WindowDidChangeScreen                   func(notification foundation.Notification)
	_WindowDidEnterFullScreen                func(notification foundation.Notification)
	_WindowWillExitFullScreen                func(notification foundation.Notification)
	_WindowWillUseStandardFrameDefaultFrame  func(window Window, newFrame foundation.Rect) foundation.Rect
	_WindowDidEnterVersionBrowser            func(notification foundation.Notification)
}

func (di *WindowDelegate) HasWindowDidFailToExitFullScreen() bool {
	return di._WindowDidFailToExitFullScreen != nil
}

// Called if the window failed to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419573-windowdidfailtoexitfullscreen?language=objc
func (di *WindowDelegate) SetWindowDidFailToExitFullScreen(f func(window Window)) {
	di._WindowDidFailToExitFullScreen = f
}

// Called if the window failed to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419573-windowdidfailtoexitfullscreen?language=objc
func (di *WindowDelegate) WindowDidFailToExitFullScreen(window Window) {
	di._WindowDidFailToExitFullScreen(window)
}
func (di *WindowDelegate) HasWindowShouldZoomToFrame() bool {
	return di._WindowShouldZoomToFrame != nil
}

// Asks the delegate whether the specified window should zoom to the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419533-windowshouldzoom?language=objc
func (di *WindowDelegate) SetWindowShouldZoomToFrame(f func(window Window, newFrame foundation.Rect) bool) {
	di._WindowShouldZoomToFrame = f
}

// Asks the delegate whether the specified window should zoom to the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419533-windowshouldzoom?language=objc
func (di *WindowDelegate) WindowShouldZoomToFrame(window Window, newFrame foundation.Rect) bool {
	return di._WindowShouldZoomToFrame(window, newFrame)
}
func (di *WindowDelegate) HasWindowDidEndSheet() bool {
	return di._WindowDidEndSheet != nil
}

// Tells the delegate that the window has closed a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419773-windowdidendsheet?language=objc
func (di *WindowDelegate) SetWindowDidEndSheet(f func(notification foundation.Notification)) {
	di._WindowDidEndSheet = f
}

// Tells the delegate that the window has closed a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419773-windowdidendsheet?language=objc
func (di *WindowDelegate) WindowDidEndSheet(notification foundation.Notification) {
	di._WindowDidEndSheet(notification)
}
func (di *WindowDelegate) HasWindowDidResize() bool {
	return di._WindowDidResize != nil
}

// Tells the delegate that the window has been resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419567-windowdidresize?language=objc
func (di *WindowDelegate) SetWindowDidResize(f func(notification foundation.Notification)) {
	di._WindowDidResize = f
}

// Tells the delegate that the window has been resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419567-windowdidresize?language=objc
func (di *WindowDelegate) WindowDidResize(notification foundation.Notification) {
	di._WindowDidResize(notification)
}
func (di *WindowDelegate) HasWindowWillEnterFullScreen() bool {
	return di._WindowWillEnterFullScreen != nil
}

// The window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419563-windowwillenterfullscreen?language=objc
func (di *WindowDelegate) SetWindowWillEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillEnterFullScreen = f
}

// The window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419563-windowwillenterfullscreen?language=objc
func (di *WindowDelegate) WindowWillEnterFullScreen(notification foundation.Notification) {
	di._WindowWillEnterFullScreen(notification)
}
func (di *WindowDelegate) HasWindowWillMove() bool {
	return di._WindowWillMove != nil
}

// Tells the delegate that the window is about to move. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419336-windowwillmove?language=objc
func (di *WindowDelegate) SetWindowWillMove(f func(notification foundation.Notification)) {
	di._WindowWillMove = f
}

// Tells the delegate that the window is about to move. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419336-windowwillmove?language=objc
func (di *WindowDelegate) WindowWillMove(notification foundation.Notification) {
	di._WindowWillMove(notification)
}
func (di *WindowDelegate) HasWindowWillStartLiveResize() bool {
	return di._WindowWillStartLiveResize != nil
}

// Tells the delegate that the window is about to be live resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419555-windowwillstartliveresize?language=objc
func (di *WindowDelegate) SetWindowWillStartLiveResize(f func(notification foundation.Notification)) {
	di._WindowWillStartLiveResize = f
}

// Tells the delegate that the window is about to be live resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419555-windowwillstartliveresize?language=objc
func (di *WindowDelegate) WindowWillStartLiveResize(notification foundation.Notification) {
	di._WindowWillStartLiveResize(notification)
}
func (di *WindowDelegate) HasWindowDidFailToEnterFullScreen() bool {
	return di._WindowDidFailToEnterFullScreen != nil
}

// Called if the window failed to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419591-windowdidfailtoenterfullscreen?language=objc
func (di *WindowDelegate) SetWindowDidFailToEnterFullScreen(f func(window Window)) {
	di._WindowDidFailToEnterFullScreen = f
}

// Called if the window failed to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419591-windowdidfailtoenterfullscreen?language=objc
func (di *WindowDelegate) WindowDidFailToEnterFullScreen(window Window) {
	di._WindowDidFailToEnterFullScreen(window)
}
func (di *WindowDelegate) HasWindowWillResizeToSize() bool {
	return di._WindowWillResizeToSize != nil
}

// Tells the delegate that the window is being resized (whether by the user or through one of the setFrame... methods other than setFrame:display:). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419292-windowwillresize?language=objc
func (di *WindowDelegate) SetWindowWillResizeToSize(f func(sender Window, frameSize foundation.Size) foundation.Size) {
	di._WindowWillResizeToSize = f
}

// Tells the delegate that the window is being resized (whether by the user or through one of the setFrame... methods other than setFrame:display:). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419292-windowwillresize?language=objc
func (di *WindowDelegate) WindowWillResizeToSize(sender Window, frameSize foundation.Size) foundation.Size {
	return di._WindowWillResizeToSize(sender, frameSize)
}
func (di *WindowDelegate) HasWindowDidChangeOcclusionState() bool {
	return di._WindowDidChangeOcclusionState != nil
}

// Tells the delegate that the window changed its occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419424-windowdidchangeocclusionstate?language=objc
func (di *WindowDelegate) SetWindowDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._WindowDidChangeOcclusionState = f
}

// Tells the delegate that the window changed its occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419424-windowdidchangeocclusionstate?language=objc
func (di *WindowDelegate) WindowDidChangeOcclusionState(notification foundation.Notification) {
	di._WindowDidChangeOcclusionState(notification)
}
func (di *WindowDelegate) HasWindowDidChangeScreenProfile() bool {
	return di._WindowDidChangeScreenProfile != nil
}

// Tells the delegate that the window has changed screen display profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419581-windowdidchangescreenprofile?language=objc
func (di *WindowDelegate) SetWindowDidChangeScreenProfile(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreenProfile = f
}

// Tells the delegate that the window has changed screen display profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419581-windowdidchangescreenprofile?language=objc
func (di *WindowDelegate) WindowDidChangeScreenProfile(notification foundation.Notification) {
	di._WindowDidChangeScreenProfile(notification)
}
func (di *WindowDelegate) HasWindowWillReturnUndoManager() bool {
	return di._WindowWillReturnUndoManager != nil
}

// Tells the delegate that the window’s undo manager has been requested. Returns the appropriate undo manager for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419745-windowwillreturnundomanager?language=objc
func (di *WindowDelegate) SetWindowWillReturnUndoManager(f func(window Window) foundation.IUndoManager) {
	di._WindowWillReturnUndoManager = f
}

// Tells the delegate that the window’s undo manager has been requested. Returns the appropriate undo manager for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419745-windowwillreturnundomanager?language=objc
func (di *WindowDelegate) WindowWillReturnUndoManager(window Window) foundation.IUndoManager {
	return di._WindowWillReturnUndoManager(window)
}
func (di *WindowDelegate) HasWindowWillBeginSheet() bool {
	return di._WindowWillBeginSheet != nil
}

// Notifies the delegate that the window is about to open a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419408-windowwillbeginsheet?language=objc
func (di *WindowDelegate) SetWindowWillBeginSheet(f func(notification foundation.Notification)) {
	di._WindowWillBeginSheet = f
}

// Notifies the delegate that the window is about to open a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419408-windowwillbeginsheet?language=objc
func (di *WindowDelegate) WindowWillBeginSheet(notification foundation.Notification) {
	di._WindowWillBeginSheet(notification)
}
func (di *WindowDelegate) HasWindowDidResignKey() bool {
	return di._WindowDidResignKey != nil
}

// Tells the delegate that the window has resigned key window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419711-windowdidresignkey?language=objc
func (di *WindowDelegate) SetWindowDidResignKey(f func(notification foundation.Notification)) {
	di._WindowDidResignKey = f
}

// Tells the delegate that the window has resigned key window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419711-windowdidresignkey?language=objc
func (di *WindowDelegate) WindowDidResignKey(notification foundation.Notification) {
	di._WindowDidResignKey(notification)
}
func (di *WindowDelegate) HasWindowShouldClose() bool {
	return di._WindowShouldClose != nil
}

// Tells the delegate that the user has attempted to close a window or the window has received a performClose: message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419380-windowshouldclose?language=objc
func (di *WindowDelegate) SetWindowShouldClose(f func(sender Window) bool) {
	di._WindowShouldClose = f
}

// Tells the delegate that the user has attempted to close a window or the window has received a performClose: message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419380-windowshouldclose?language=objc
func (di *WindowDelegate) WindowShouldClose(sender Window) bool {
	return di._WindowShouldClose(sender)
}
func (di *WindowDelegate) HasWindowDidExitVersionBrowser() bool {
	return di._WindowDidExitVersionBrowser != nil
}

// Tells the delegate that the window has left version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419501-windowdidexitversionbrowser?language=objc
func (di *WindowDelegate) SetWindowDidExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidExitVersionBrowser = f
}

// Tells the delegate that the window has left version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419501-windowdidexitversionbrowser?language=objc
func (di *WindowDelegate) WindowDidExitVersionBrowser(notification foundation.Notification) {
	di._WindowDidExitVersionBrowser(notification)
}
func (di *WindowDelegate) HasWindowDidBecomeKey() bool {
	return di._WindowDidBecomeKey != nil
}

// Tells the delegate that the window has become the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419737-windowdidbecomekey?language=objc
func (di *WindowDelegate) SetWindowDidBecomeKey(f func(notification foundation.Notification)) {
	di._WindowDidBecomeKey = f
}

// Tells the delegate that the window has become the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419737-windowdidbecomekey?language=objc
func (di *WindowDelegate) WindowDidBecomeKey(notification foundation.Notification) {
	di._WindowDidBecomeKey(notification)
}
func (di *WindowDelegate) HasWindowDidResignMain() bool {
	return di._WindowDidResignMain != nil
}

// Tells the delegate that the window has resigned main window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419584-windowdidresignmain?language=objc
func (di *WindowDelegate) SetWindowDidResignMain(f func(notification foundation.Notification)) {
	di._WindowDidResignMain = f
}

// Tells the delegate that the window has resigned main window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419584-windowdidresignmain?language=objc
func (di *WindowDelegate) WindowDidResignMain(notification foundation.Notification) {
	di._WindowDidResignMain(notification)
}
func (di *WindowDelegate) HasWindowDidExitFullScreen() bool {
	return di._WindowDidExitFullScreen != nil
}

// The window has left full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419146-windowdidexitfullscreen?language=objc
func (di *WindowDelegate) SetWindowDidExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidExitFullScreen = f
}

// The window has left full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419146-windowdidexitfullscreen?language=objc
func (di *WindowDelegate) WindowDidExitFullScreen(notification foundation.Notification) {
	di._WindowDidExitFullScreen(notification)
}
func (di *WindowDelegate) HasWindowDidBecomeMain() bool {
	return di._WindowDidBecomeMain != nil
}

// Tells the delegate that the window has become main. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419190-windowdidbecomemain?language=objc
func (di *WindowDelegate) SetWindowDidBecomeMain(f func(notification foundation.Notification)) {
	di._WindowDidBecomeMain = f
}

// Tells the delegate that the window has become main. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419190-windowdidbecomemain?language=objc
func (di *WindowDelegate) WindowDidBecomeMain(notification foundation.Notification) {
	di._WindowDidBecomeMain(notification)
}
func (di *WindowDelegate) HasWindowDidMiniaturize() bool {
	return di._WindowDidMiniaturize != nil
}

// Tells the delegate that the window has been minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419621-windowdidminiaturize?language=objc
func (di *WindowDelegate) SetWindowDidMiniaturize(f func(notification foundation.Notification)) {
	di._WindowDidMiniaturize = f
}

// Tells the delegate that the window has been minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419621-windowdidminiaturize?language=objc
func (di *WindowDelegate) WindowDidMiniaturize(notification foundation.Notification) {
	di._WindowDidMiniaturize(notification)
}
func (di *WindowDelegate) HasWindowDidExpose() bool {
	return di._WindowDidExpose != nil
}

// Tells the delegate that the window has been exposed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419258-windowdidexpose?language=objc
func (di *WindowDelegate) SetWindowDidExpose(f func(notification foundation.Notification)) {
	di._WindowDidExpose = f
}

// Tells the delegate that the window has been exposed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419258-windowdidexpose?language=objc
func (di *WindowDelegate) WindowDidExpose(notification foundation.Notification) {
	di._WindowDidExpose(notification)
}
func (di *WindowDelegate) HasWindowWillMiniaturize() bool {
	return di._WindowWillMiniaturize != nil
}

// Tells the delegate that the window is about to be minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419461-windowwillminiaturize?language=objc
func (di *WindowDelegate) SetWindowWillMiniaturize(f func(notification foundation.Notification)) {
	di._WindowWillMiniaturize = f
}

// Tells the delegate that the window is about to be minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419461-windowwillminiaturize?language=objc
func (di *WindowDelegate) WindowWillMiniaturize(notification foundation.Notification) {
	di._WindowWillMiniaturize(notification)
}
func (di *WindowDelegate) HasWindowDidChangeBackingProperties() bool {
	return di._WindowDidChangeBackingProperties != nil
}

// Tells the delegate that the window backing properties changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419517-windowdidchangebackingproperties?language=objc
func (di *WindowDelegate) SetWindowDidChangeBackingProperties(f func(notification foundation.Notification)) {
	di._WindowDidChangeBackingProperties = f
}

// Tells the delegate that the window backing properties changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419517-windowdidchangebackingproperties?language=objc
func (di *WindowDelegate) WindowDidChangeBackingProperties(notification foundation.Notification) {
	di._WindowDidChangeBackingProperties(notification)
}
func (di *WindowDelegate) HasWindowWillClose() bool {
	return di._WindowWillClose != nil
}

// Tells the delegate that the window is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419605-windowwillclose?language=objc
func (di *WindowDelegate) SetWindowWillClose(f func(notification foundation.Notification)) {
	di._WindowWillClose = f
}

// Tells the delegate that the window is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419605-windowwillclose?language=objc
func (di *WindowDelegate) WindowWillClose(notification foundation.Notification) {
	di._WindowWillClose(notification)
}
func (di *WindowDelegate) HasCustomWindowsToEnterFullScreenForWindow() bool {
	return di._CustomWindowsToEnterFullScreenForWindow != nil
}

// Called when the window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419521-customwindowstoenterfullscreenfo?language=objc
func (di *WindowDelegate) SetCustomWindowsToEnterFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow = f
}

// Called when the window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419521-customwindowstoenterfullscreenfo?language=objc
func (di *WindowDelegate) CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow(window)
}
func (di *WindowDelegate) HasWindowWillEnterVersionBrowser() bool {
	return di._WindowWillEnterVersionBrowser != nil
}

// Tells the delegate the window is about to enter version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419463-windowwillenterversionbrowser?language=objc
func (di *WindowDelegate) SetWindowWillEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillEnterVersionBrowser = f
}

// Tells the delegate the window is about to enter version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419463-windowwillenterversionbrowser?language=objc
func (di *WindowDelegate) WindowWillEnterVersionBrowser(notification foundation.Notification) {
	di._WindowWillEnterVersionBrowser(notification)
}
func (di *WindowDelegate) HasWindowDidEndLiveResize() bool {
	return di._WindowDidEndLiveResize != nil
}

// Tells the delegate that a live resize operation on the window has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419150-windowdidendliveresize?language=objc
func (di *WindowDelegate) SetWindowDidEndLiveResize(f func(notification foundation.Notification)) {
	di._WindowDidEndLiveResize = f
}

// Tells the delegate that a live resize operation on the window has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419150-windowdidendliveresize?language=objc
func (di *WindowDelegate) WindowDidEndLiveResize(notification foundation.Notification) {
	di._WindowDidEndLiveResize(notification)
}
func (di *WindowDelegate) HasWindowDidUpdate() bool {
	return di._WindowDidUpdate != nil
}

// Tells the delegate that the window received an update message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419493-windowdidupdate?language=objc
func (di *WindowDelegate) SetWindowDidUpdate(f func(notification foundation.Notification)) {
	di._WindowDidUpdate = f
}

// Tells the delegate that the window received an update message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419493-windowdidupdate?language=objc
func (di *WindowDelegate) WindowDidUpdate(notification foundation.Notification) {
	di._WindowDidUpdate(notification)
}
func (di *WindowDelegate) HasWindowWillExitVersionBrowser() bool {
	return di._WindowWillExitVersionBrowser != nil
}

// Tells the delegate that the window is about to leave version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419252-windowwillexitversionbrowser?language=objc
func (di *WindowDelegate) SetWindowWillExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillExitVersionBrowser = f
}

// Tells the delegate that the window is about to leave version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419252-windowwillexitversionbrowser?language=objc
func (di *WindowDelegate) WindowWillExitVersionBrowser(notification foundation.Notification) {
	di._WindowWillExitVersionBrowser(notification)
}
func (di *WindowDelegate) HasWindowWillReturnFieldEditorToObject() bool {
	return di._WindowWillReturnFieldEditorToObject != nil
}

// Tells the delegate that the field editor for a text-displaying object has been requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419416-windowwillreturnfieldeditor?language=objc
func (di *WindowDelegate) SetWindowWillReturnFieldEditorToObject(f func(sender Window, client objc.Object) objc.IObject) {
	di._WindowWillReturnFieldEditorToObject = f
}

// Tells the delegate that the field editor for a text-displaying object has been requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419416-windowwillreturnfieldeditor?language=objc
func (di *WindowDelegate) WindowWillReturnFieldEditorToObject(sender Window, client objc.Object) objc.IObject {
	return di._WindowWillReturnFieldEditorToObject(sender, client)
}
func (di *WindowDelegate) HasWindowDidDeminiaturize() bool {
	return di._WindowDidDeminiaturize != nil
}

// Tells the delegate that the window has been deminimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419296-windowdiddeminiaturize?language=objc
func (di *WindowDelegate) SetWindowDidDeminiaturize(f func(notification foundation.Notification)) {
	di._WindowDidDeminiaturize = f
}

// Tells the delegate that the window has been deminimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419296-windowdiddeminiaturize?language=objc
func (di *WindowDelegate) WindowDidDeminiaturize(notification foundation.Notification) {
	di._WindowDidDeminiaturize(notification)
}
func (di *WindowDelegate) HasCustomWindowsToExitFullScreenForWindow() bool {
	return di._CustomWindowsToExitFullScreenForWindow != nil
}

// Called when the window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419070-customwindowstoexitfullscreenfor?language=objc
func (di *WindowDelegate) SetCustomWindowsToExitFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToExitFullScreenForWindow = f
}

// Called when the window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419070-customwindowstoexitfullscreenfor?language=objc
func (di *WindowDelegate) CustomWindowsToExitFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToExitFullScreenForWindow(window)
}
func (di *WindowDelegate) HasWindowDidDecodeRestorableState() bool {
	return di._WindowDidDecodeRestorableState != nil
}

// Tells the delegate the window is has extracted its restorable state from a given archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419475-window?language=objc
func (di *WindowDelegate) SetWindowDidDecodeRestorableState(f func(window Window, state foundation.Coder)) {
	di._WindowDidDecodeRestorableState = f
}

// Tells the delegate the window is has extracted its restorable state from a given archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419475-window?language=objc
func (di *WindowDelegate) WindowDidDecodeRestorableState(window Window, state foundation.Coder) {
	di._WindowDidDecodeRestorableState(window, state)
}
func (di *WindowDelegate) HasWindowDidMove() bool {
	return di._WindowDidMove != nil
}

// Tells the delegate that the window has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419674-windowdidmove?language=objc
func (di *WindowDelegate) SetWindowDidMove(f func(notification foundation.Notification)) {
	di._WindowDidMove = f
}

// Tells the delegate that the window has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419674-windowdidmove?language=objc
func (di *WindowDelegate) WindowDidMove(notification foundation.Notification) {
	di._WindowDidMove(notification)
}
func (di *WindowDelegate) HasWindowDidChangeScreen() bool {
	return di._WindowDidChangeScreen != nil
}

// Tells the delegate that the window has changed screens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419267-windowdidchangescreen?language=objc
func (di *WindowDelegate) SetWindowDidChangeScreen(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreen = f
}

// Tells the delegate that the window has changed screens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419267-windowdidchangescreen?language=objc
func (di *WindowDelegate) WindowDidChangeScreen(notification foundation.Notification) {
	di._WindowDidChangeScreen(notification)
}
func (di *WindowDelegate) HasWindowDidEnterFullScreen() bool {
	return di._WindowDidEnterFullScreen != nil
}

// The window has entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419116-windowdidenterfullscreen?language=objc
func (di *WindowDelegate) SetWindowDidEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidEnterFullScreen = f
}

// The window has entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419116-windowdidenterfullscreen?language=objc
func (di *WindowDelegate) WindowDidEnterFullScreen(notification foundation.Notification) {
	di._WindowDidEnterFullScreen(notification)
}
func (di *WindowDelegate) HasWindowWillExitFullScreen() bool {
	return di._WindowWillExitFullScreen != nil
}

// The window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419332-windowwillexitfullscreen?language=objc
func (di *WindowDelegate) SetWindowWillExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillExitFullScreen = f
}

// The window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419332-windowwillexitfullscreen?language=objc
func (di *WindowDelegate) WindowWillExitFullScreen(notification foundation.Notification) {
	di._WindowWillExitFullScreen(notification)
}
func (di *WindowDelegate) HasWindowWillUseStandardFrameDefaultFrame() bool {
	return di._WindowWillUseStandardFrameDefaultFrame != nil
}

// Called by NSWindow’s zoom: method while determining the frame a window may be zoomed to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419684-windowwillusestandardframe?language=objc
func (di *WindowDelegate) SetWindowWillUseStandardFrameDefaultFrame(f func(window Window, newFrame foundation.Rect) foundation.Rect) {
	di._WindowWillUseStandardFrameDefaultFrame = f
}

// Called by NSWindow’s zoom: method while determining the frame a window may be zoomed to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419684-windowwillusestandardframe?language=objc
func (di *WindowDelegate) WindowWillUseStandardFrameDefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect {
	return di._WindowWillUseStandardFrameDefaultFrame(window, newFrame)
}
func (di *WindowDelegate) HasWindowDidEnterVersionBrowser() bool {
	return di._WindowDidEnterVersionBrowser != nil
}

// Tells the delegate that the window has entered version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419064-windowdidenterversionbrowser?language=objc
func (di *WindowDelegate) SetWindowDidEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidEnterVersionBrowser = f
}

// Tells the delegate that the window has entered version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419064-windowdidenterversionbrowser?language=objc
func (di *WindowDelegate) WindowDidEnterVersionBrowser(notification foundation.Notification) {
	di._WindowDidEnterVersionBrowser(notification)
}

// A concrete type wrapper for the [PWindowDelegate] protocol.
type WindowDelegateWrapper struct {
	objc.Object
}

func (w_ WindowDelegateWrapper) HasWindowDidFailToExitFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidFailToExitFullScreen:"))
}

// Called if the window failed to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419573-windowdidfailtoexitfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowDidFailToExitFullScreen(window IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidFailToExitFullScreen:"), objc.Ptr(window))
}

func (w_ WindowDelegateWrapper) HasWindowShouldZoomToFrame() bool {
	return w_.RespondsToSelector(objc.Sel("windowShouldZoom:toFrame:"))
}

// Asks the delegate whether the specified window should zoom to the specified frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419533-windowshouldzoom?language=objc
func (w_ WindowDelegateWrapper) WindowShouldZoomToFrame(window IWindow, newFrame foundation.Rect) bool {
	rv := objc.Call[bool](w_, objc.Sel("windowShouldZoom:toFrame:"), objc.Ptr(window), newFrame)
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidEndSheet() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidEndSheet:"))
}

// Tells the delegate that the window has closed a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419773-windowdidendsheet?language=objc
func (w_ WindowDelegateWrapper) WindowDidEndSheet(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidEndSheet:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidResize() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidResize:"))
}

// Tells the delegate that the window has been resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419567-windowdidresize?language=objc
func (w_ WindowDelegateWrapper) WindowDidResize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidResize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillEnterFullScreen:"))
}

// The window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419563-windowwillenterfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowWillEnterFullScreen(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillEnterFullScreen:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillMove() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillMove:"))
}

// Tells the delegate that the window is about to move. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419336-windowwillmove?language=objc
func (w_ WindowDelegateWrapper) WindowWillMove(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillMove:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillStartLiveResize() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillStartLiveResize:"))
}

// Tells the delegate that the window is about to be live resized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419555-windowwillstartliveresize?language=objc
func (w_ WindowDelegateWrapper) WindowWillStartLiveResize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillStartLiveResize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidFailToEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidFailToEnterFullScreen:"))
}

// Called if the window failed to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419591-windowdidfailtoenterfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowDidFailToEnterFullScreen(window IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidFailToEnterFullScreen:"), objc.Ptr(window))
}

func (w_ WindowDelegateWrapper) HasWindowWillResizeToSize() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillResize:toSize:"))
}

// Tells the delegate that the window is being resized (whether by the user or through one of the setFrame... methods other than setFrame:display:). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419292-windowwillresize?language=objc
func (w_ WindowDelegateWrapper) WindowWillResizeToSize(sender IWindow, frameSize foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("windowWillResize:toSize:"), objc.Ptr(sender), frameSize)
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidChangeOcclusionState() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidChangeOcclusionState:"))
}

// Tells the delegate that the window changed its occlusion state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419424-windowdidchangeocclusionstate?language=objc
func (w_ WindowDelegateWrapper) WindowDidChangeOcclusionState(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidChangeOcclusionState:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidChangeScreenProfile() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidChangeScreenProfile:"))
}

// Tells the delegate that the window has changed screen display profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419581-windowdidchangescreenprofile?language=objc
func (w_ WindowDelegateWrapper) WindowDidChangeScreenProfile(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidChangeScreenProfile:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillReturnUndoManager() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillReturnUndoManager:"))
}

// Tells the delegate that the window’s undo manager has been requested. Returns the appropriate undo manager for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419745-windowwillreturnundomanager?language=objc
func (w_ WindowDelegateWrapper) WindowWillReturnUndoManager(window IWindow) foundation.UndoManager {
	rv := objc.Call[foundation.UndoManager](w_, objc.Sel("windowWillReturnUndoManager:"), objc.Ptr(window))
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowWillBeginSheet() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillBeginSheet:"))
}

// Notifies the delegate that the window is about to open a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419408-windowwillbeginsheet?language=objc
func (w_ WindowDelegateWrapper) WindowWillBeginSheet(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillBeginSheet:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidResignKey() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidResignKey:"))
}

// Tells the delegate that the window has resigned key window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419711-windowdidresignkey?language=objc
func (w_ WindowDelegateWrapper) WindowDidResignKey(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidResignKey:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowShouldClose() bool {
	return w_.RespondsToSelector(objc.Sel("windowShouldClose:"))
}

// Tells the delegate that the user has attempted to close a window or the window has received a performClose: message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419380-windowshouldclose?language=objc
func (w_ WindowDelegateWrapper) WindowShouldClose(sender IWindow) bool {
	rv := objc.Call[bool](w_, objc.Sel("windowShouldClose:"), objc.Ptr(sender))
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidExitVersionBrowser:"))
}

// Tells the delegate that the window has left version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419501-windowdidexitversionbrowser?language=objc
func (w_ WindowDelegateWrapper) WindowDidExitVersionBrowser(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidExitVersionBrowser:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidBecomeKey() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidBecomeKey:"))
}

// Tells the delegate that the window has become the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419737-windowdidbecomekey?language=objc
func (w_ WindowDelegateWrapper) WindowDidBecomeKey(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidBecomeKey:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidResignMain() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidResignMain:"))
}

// Tells the delegate that the window has resigned main window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419584-windowdidresignmain?language=objc
func (w_ WindowDelegateWrapper) WindowDidResignMain(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidResignMain:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidExitFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidExitFullScreen:"))
}

// The window has left full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419146-windowdidexitfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowDidExitFullScreen(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidExitFullScreen:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidBecomeMain() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidBecomeMain:"))
}

// Tells the delegate that the window has become main. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419190-windowdidbecomemain?language=objc
func (w_ WindowDelegateWrapper) WindowDidBecomeMain(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidBecomeMain:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidMiniaturize() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidMiniaturize:"))
}

// Tells the delegate that the window has been minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419621-windowdidminiaturize?language=objc
func (w_ WindowDelegateWrapper) WindowDidMiniaturize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidMiniaturize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidExpose() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidExpose:"))
}

// Tells the delegate that the window has been exposed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419258-windowdidexpose?language=objc
func (w_ WindowDelegateWrapper) WindowDidExpose(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidExpose:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillMiniaturize() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillMiniaturize:"))
}

// Tells the delegate that the window is about to be minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419461-windowwillminiaturize?language=objc
func (w_ WindowDelegateWrapper) WindowWillMiniaturize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillMiniaturize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidChangeBackingProperties() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidChangeBackingProperties:"))
}

// Tells the delegate that the window backing properties changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419517-windowdidchangebackingproperties?language=objc
func (w_ WindowDelegateWrapper) WindowDidChangeBackingProperties(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidChangeBackingProperties:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillClose() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillClose:"))
}

// Tells the delegate that the window is about to close. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419605-windowwillclose?language=objc
func (w_ WindowDelegateWrapper) WindowWillClose(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillClose:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasCustomWindowsToEnterFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.Sel("customWindowsToEnterFullScreenForWindow:"))
}

// Called when the window is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419521-customwindowstoenterfullscreenfo?language=objc
func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow(window IWindow) []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("customWindowsToEnterFullScreenForWindow:"), objc.Ptr(window))
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowWillEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillEnterVersionBrowser:"))
}

// Tells the delegate the window is about to enter version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419463-windowwillenterversionbrowser?language=objc
func (w_ WindowDelegateWrapper) WindowWillEnterVersionBrowser(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillEnterVersionBrowser:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidEndLiveResize() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidEndLiveResize:"))
}

// Tells the delegate that a live resize operation on the window has ended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419150-windowdidendliveresize?language=objc
func (w_ WindowDelegateWrapper) WindowDidEndLiveResize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidEndLiveResize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidUpdate() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidUpdate:"))
}

// Tells the delegate that the window received an update message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419493-windowdidupdate?language=objc
func (w_ WindowDelegateWrapper) WindowDidUpdate(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidUpdate:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillExitVersionBrowser:"))
}

// Tells the delegate that the window is about to leave version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419252-windowwillexitversionbrowser?language=objc
func (w_ WindowDelegateWrapper) WindowWillExitVersionBrowser(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillExitVersionBrowser:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillReturnFieldEditorToObject() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillReturnFieldEditor:toObject:"))
}

// Tells the delegate that the field editor for a text-displaying object has been requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419416-windowwillreturnfieldeditor?language=objc
func (w_ WindowDelegateWrapper) WindowWillReturnFieldEditorToObject(sender IWindow, client objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("windowWillReturnFieldEditor:toObject:"), objc.Ptr(sender), client)
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidDeminiaturize() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidDeminiaturize:"))
}

// Tells the delegate that the window has been deminimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419296-windowdiddeminiaturize?language=objc
func (w_ WindowDelegateWrapper) WindowDidDeminiaturize(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidDeminiaturize:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasCustomWindowsToExitFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.Sel("customWindowsToExitFullScreenForWindow:"))
}

// Called when the window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419070-customwindowstoexitfullscreenfor?language=objc
func (w_ WindowDelegateWrapper) CustomWindowsToExitFullScreenForWindow(window IWindow) []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("customWindowsToExitFullScreenForWindow:"), objc.Ptr(window))
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidDecodeRestorableState() bool {
	return w_.RespondsToSelector(objc.Sel("window:didDecodeRestorableState:"))
}

// Tells the delegate the window is has extracted its restorable state from a given archiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419475-window?language=objc
func (w_ WindowDelegateWrapper) WindowDidDecodeRestorableState(window IWindow, state foundation.ICoder) {
	objc.Call[objc.Void](w_, objc.Sel("window:didDecodeRestorableState:"), objc.Ptr(window), objc.Ptr(state))
}

func (w_ WindowDelegateWrapper) HasWindowDidMove() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidMove:"))
}

// Tells the delegate that the window has moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419674-windowdidmove?language=objc
func (w_ WindowDelegateWrapper) WindowDidMove(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidMove:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidChangeScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidChangeScreen:"))
}

// Tells the delegate that the window has changed screens. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419267-windowdidchangescreen?language=objc
func (w_ WindowDelegateWrapper) WindowDidChangeScreen(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidChangeScreen:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowDidEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidEnterFullScreen:"))
}

// The window has entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419116-windowdidenterfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowDidEnterFullScreen(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidEnterFullScreen:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillExitFullScreen() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillExitFullScreen:"))
}

// The window is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419332-windowwillexitfullscreen?language=objc
func (w_ WindowDelegateWrapper) WindowWillExitFullScreen(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowWillExitFullScreen:"), objc.Ptr(notification))
}

func (w_ WindowDelegateWrapper) HasWindowWillUseStandardFrameDefaultFrame() bool {
	return w_.RespondsToSelector(objc.Sel("windowWillUseStandardFrame:defaultFrame:"))
}

// Called by NSWindow’s zoom: method while determining the frame a window may be zoomed to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419684-windowwillusestandardframe?language=objc
func (w_ WindowDelegateWrapper) WindowWillUseStandardFrameDefaultFrame(window IWindow, newFrame foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("windowWillUseStandardFrame:defaultFrame:"), objc.Ptr(window), newFrame)
	return rv
}

func (w_ WindowDelegateWrapper) HasWindowDidEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.Sel("windowDidEnterVersionBrowser:"))
}

// Tells the delegate that the window has entered version browsing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindowdelegate/1419064-windowdidenterversionbrowser?language=objc
func (w_ WindowDelegateWrapper) WindowDidEnterVersionBrowser(notification foundation.INotification) {
	objc.Call[objc.Void](w_, objc.Sel("windowDidEnterVersionBrowser:"), objc.Ptr(notification))
}
