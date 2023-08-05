// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IWindowDelegate interface {
	ImplementsWindowWillPositionSheetUsingRect() bool
	// optional
	WindowWillPositionSheetUsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	ImplementsWindowWillBeginSheet() bool
	// optional
	WindowWillBeginSheet(notification foundation.Notification)
	ImplementsWindowDidEndSheet() bool
	// optional
	WindowDidEndSheet(notification foundation.Notification)
	ImplementsWindowWillResizeToSize() bool
	// optional
	WindowWillResizeToSize(sender Window, frameSize foundation.Size) foundation.Size
	ImplementsWindowDidResize() bool
	// optional
	WindowDidResize(notification foundation.Notification)
	ImplementsWindowWillStartLiveResize() bool
	// optional
	WindowWillStartLiveResize(notification foundation.Notification)
	ImplementsWindowDidEndLiveResize() bool
	// optional
	WindowDidEndLiveResize(notification foundation.Notification)
	ImplementsWindowWillMiniaturize() bool
	// optional
	WindowWillMiniaturize(notification foundation.Notification)
	ImplementsWindowDidMiniaturize() bool
	// optional
	WindowDidMiniaturize(notification foundation.Notification)
	ImplementsWindowDidDeminiaturize() bool
	// optional
	WindowDidDeminiaturize(notification foundation.Notification)
	ImplementsWindowWillUseStandardFrameDefaultFrame() bool
	// optional
	WindowWillUseStandardFrameDefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect
	ImplementsWindowShouldZoomToFrame() bool
	// optional
	WindowShouldZoomToFrame(window Window, newFrame foundation.Rect) bool
	ImplementsWindowWillUseFullScreenContentSize() bool
	// optional
	WindowWillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size
	ImplementsWindowWillUseFullScreenPresentationOptions() bool
	// optional
	WindowWillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	ImplementsWindowWillEnterFullScreen() bool
	// optional
	WindowWillEnterFullScreen(notification foundation.Notification)
	ImplementsWindowDidEnterFullScreen() bool
	// optional
	WindowDidEnterFullScreen(notification foundation.Notification)
	ImplementsWindowWillExitFullScreen() bool
	// optional
	WindowWillExitFullScreen(notification foundation.Notification)
	ImplementsWindowDidExitFullScreen() bool
	// optional
	WindowDidExitFullScreen(notification foundation.Notification)
	ImplementsCustomWindowsToEnterFullScreenForWindow() bool
	// optional
	CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow
	ImplementsCustomWindowsToEnterFullScreenForWindowOnScreen() bool
	// optional
	CustomWindowsToEnterFullScreenForWindowOnScreen(window Window, screen Screen) []IWindow
	ImplementsWindowStartCustomAnimationToEnterFullScreenWithDuration() bool
	// optional
	WindowStartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration() bool
	// optional
	WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(window Window, screen Screen, duration foundation.TimeInterval)
	ImplementsWindowDidFailToEnterFullScreen() bool
	// optional
	WindowDidFailToEnterFullScreen(window Window)
	ImplementsCustomWindowsToExitFullScreenForWindow() bool
	// optional
	CustomWindowsToExitFullScreenForWindow(window Window) []IWindow
	ImplementsWindowStartCustomAnimationToExitFullScreenWithDuration() bool
	// optional
	WindowStartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval)
	ImplementsWindowDidFailToExitFullScreen() bool
	// optional
	WindowDidFailToExitFullScreen(window Window)
	ImplementsWindowWillMove() bool
	// optional
	WindowWillMove(notification foundation.Notification)
	ImplementsWindowDidMove() bool
	// optional
	WindowDidMove(notification foundation.Notification)
	ImplementsWindowDidChangeScreen() bool
	// optional
	WindowDidChangeScreen(notification foundation.Notification)
	ImplementsWindowDidChangeScreenProfile() bool
	// optional
	WindowDidChangeScreenProfile(notification foundation.Notification)
	ImplementsWindowDidChangeBackingProperties() bool
	// optional
	WindowDidChangeBackingProperties(notification foundation.Notification)
	ImplementsWindowShouldClose() bool
	// optional
	WindowShouldClose(sender Window) bool
	ImplementsWindowWillClose() bool
	// optional
	WindowWillClose(notification foundation.Notification)
	ImplementsWindowDidBecomeKey() bool
	// optional
	WindowDidBecomeKey(notification foundation.Notification)
	ImplementsWindowDidResignKey() bool
	// optional
	WindowDidResignKey(notification foundation.Notification)
	ImplementsWindowDidBecomeMain() bool
	// optional
	WindowDidBecomeMain(notification foundation.Notification)
	ImplementsWindowDidResignMain() bool
	// optional
	WindowDidResignMain(notification foundation.Notification)
	ImplementsWindowWillReturnFieldEditorToObject() bool
	// optional
	WindowWillReturnFieldEditorToObject(sender Window, client objc.Object) objc.IObject
	ImplementsWindowDidUpdate() bool
	// optional
	WindowDidUpdate(notification foundation.Notification)
	ImplementsWindowDidExpose() bool
	// optional
	WindowDidExpose(notification foundation.Notification)
	ImplementsWindowDidChangeOcclusionState() bool
	// optional
	WindowDidChangeOcclusionState(notification foundation.Notification)
	ImplementsWindowShouldDragDocumentWithEventFromWithPasteboard() bool
	// optional
	WindowShouldDragDocumentWithEventFromWithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	ImplementsWindowWillReturnUndoManager() bool
	// optional
	WindowWillReturnUndoManager(window Window) foundation.IUndoManager
	ImplementsWindowShouldPopUpDocumentPathMenu() bool
	// optional
	WindowShouldPopUpDocumentPathMenu(window Window, menu Menu) bool
	ImplementsWindowWillEncodeRestorableState() bool
	// optional
	WindowWillEncodeRestorableState(window Window, state foundation.Coder)
	ImplementsWindowDidDecodeRestorableState() bool
	// optional
	WindowDidDecodeRestorableState(window Window, state foundation.Coder)
	ImplementsWindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize() bool
	// optional
	WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	ImplementsWindowWillEnterVersionBrowser() bool
	// optional
	WindowWillEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidEnterVersionBrowser() bool
	// optional
	WindowDidEnterVersionBrowser(notification foundation.Notification)
	ImplementsWindowWillExitVersionBrowser() bool
	// optional
	WindowWillExitVersionBrowser(notification foundation.Notification)
	ImplementsWindowDidExitVersionBrowser() bool
	// optional
	WindowDidExitVersionBrowser(notification foundation.Notification)
}

type WindowDelegate struct {
	_WindowWillPositionSheetUsingRect                                    func(window Window, sheet Window, rect foundation.Rect) foundation.Rect
	_WindowWillBeginSheet                                                func(notification foundation.Notification)
	_WindowDidEndSheet                                                   func(notification foundation.Notification)
	_WindowWillResizeToSize                                              func(sender Window, frameSize foundation.Size) foundation.Size
	_WindowDidResize                                                     func(notification foundation.Notification)
	_WindowWillStartLiveResize                                           func(notification foundation.Notification)
	_WindowDidEndLiveResize                                              func(notification foundation.Notification)
	_WindowWillMiniaturize                                               func(notification foundation.Notification)
	_WindowDidMiniaturize                                                func(notification foundation.Notification)
	_WindowDidDeminiaturize                                              func(notification foundation.Notification)
	_WindowWillUseStandardFrameDefaultFrame                              func(window Window, newFrame foundation.Rect) foundation.Rect
	_WindowShouldZoomToFrame                                             func(window Window, newFrame foundation.Rect) bool
	_WindowWillUseFullScreenContentSize                                  func(window Window, proposedSize foundation.Size) foundation.Size
	_WindowWillUseFullScreenPresentationOptions                          func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions
	_WindowWillEnterFullScreen                                           func(notification foundation.Notification)
	_WindowDidEnterFullScreen                                            func(notification foundation.Notification)
	_WindowWillExitFullScreen                                            func(notification foundation.Notification)
	_WindowDidExitFullScreen                                             func(notification foundation.Notification)
	_CustomWindowsToEnterFullScreenForWindow                             func(window Window) []IWindow
	_CustomWindowsToEnterFullScreenForWindowOnScreen                     func(window Window, screen Screen) []IWindow
	_WindowStartCustomAnimationToEnterFullScreenWithDuration             func(window Window, duration foundation.TimeInterval)
	_WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration     func(window Window, screen Screen, duration foundation.TimeInterval)
	_WindowDidFailToEnterFullScreen                                      func(window Window)
	_CustomWindowsToExitFullScreenForWindow                              func(window Window) []IWindow
	_WindowStartCustomAnimationToExitFullScreenWithDuration              func(window Window, duration foundation.TimeInterval)
	_WindowDidFailToExitFullScreen                                       func(window Window)
	_WindowWillMove                                                      func(notification foundation.Notification)
	_WindowDidMove                                                       func(notification foundation.Notification)
	_WindowDidChangeScreen                                               func(notification foundation.Notification)
	_WindowDidChangeScreenProfile                                        func(notification foundation.Notification)
	_WindowDidChangeBackingProperties                                    func(notification foundation.Notification)
	_WindowShouldClose                                                   func(sender Window) bool
	_WindowWillClose                                                     func(notification foundation.Notification)
	_WindowDidBecomeKey                                                  func(notification foundation.Notification)
	_WindowDidResignKey                                                  func(notification foundation.Notification)
	_WindowDidBecomeMain                                                 func(notification foundation.Notification)
	_WindowDidResignMain                                                 func(notification foundation.Notification)
	_WindowWillReturnFieldEditorToObject                                 func(sender Window, client objc.Object) objc.IObject
	_WindowDidUpdate                                                     func(notification foundation.Notification)
	_WindowDidExpose                                                     func(notification foundation.Notification)
	_WindowDidChangeOcclusionState                                       func(notification foundation.Notification)
	_WindowShouldDragDocumentWithEventFromWithPasteboard                 func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool
	_WindowWillReturnUndoManager                                         func(window Window) foundation.IUndoManager
	_WindowShouldPopUpDocumentPathMenu                                   func(window Window, menu Menu) bool
	_WindowWillEncodeRestorableState                                     func(window Window, state foundation.Coder)
	_WindowDidDecodeRestorableState                                      func(window Window, state foundation.Coder)
	_WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size
	_WindowWillEnterVersionBrowser                                       func(notification foundation.Notification)
	_WindowDidEnterVersionBrowser                                        func(notification foundation.Notification)
	_WindowWillExitVersionBrowser                                        func(notification foundation.Notification)
	_WindowDidExitVersionBrowser                                         func(notification foundation.Notification)
}

func (di *WindowDelegate) ImplementsWindowWillPositionSheetUsingRect() bool {
	return di._WindowWillPositionSheetUsingRect != nil
}

func (di *WindowDelegate) SetWindowWillPositionSheetUsingRect(f func(window Window, sheet Window, rect foundation.Rect) foundation.Rect) {
	di._WindowWillPositionSheetUsingRect = f
}

func (di *WindowDelegate) WindowWillPositionSheetUsingRect(window Window, sheet Window, rect foundation.Rect) foundation.Rect {
	return di._WindowWillPositionSheetUsingRect(window, sheet, rect)
}
func (di *WindowDelegate) ImplementsWindowWillBeginSheet() bool {
	return di._WindowWillBeginSheet != nil
}

func (di *WindowDelegate) SetWindowWillBeginSheet(f func(notification foundation.Notification)) {
	di._WindowWillBeginSheet = f
}

func (di *WindowDelegate) WindowWillBeginSheet(notification foundation.Notification) {
	di._WindowWillBeginSheet(notification)
}
func (di *WindowDelegate) ImplementsWindowDidEndSheet() bool {
	return di._WindowDidEndSheet != nil
}

func (di *WindowDelegate) SetWindowDidEndSheet(f func(notification foundation.Notification)) {
	di._WindowDidEndSheet = f
}

func (di *WindowDelegate) WindowDidEndSheet(notification foundation.Notification) {
	di._WindowDidEndSheet(notification)
}
func (di *WindowDelegate) ImplementsWindowWillResizeToSize() bool {
	return di._WindowWillResizeToSize != nil
}

func (di *WindowDelegate) SetWindowWillResizeToSize(f func(sender Window, frameSize foundation.Size) foundation.Size) {
	di._WindowWillResizeToSize = f
}

func (di *WindowDelegate) WindowWillResizeToSize(sender Window, frameSize foundation.Size) foundation.Size {
	return di._WindowWillResizeToSize(sender, frameSize)
}
func (di *WindowDelegate) ImplementsWindowDidResize() bool {
	return di._WindowDidResize != nil
}

func (di *WindowDelegate) SetWindowDidResize(f func(notification foundation.Notification)) {
	di._WindowDidResize = f
}

func (di *WindowDelegate) WindowDidResize(notification foundation.Notification) {
	di._WindowDidResize(notification)
}
func (di *WindowDelegate) ImplementsWindowWillStartLiveResize() bool {
	return di._WindowWillStartLiveResize != nil
}

func (di *WindowDelegate) SetWindowWillStartLiveResize(f func(notification foundation.Notification)) {
	di._WindowWillStartLiveResize = f
}

func (di *WindowDelegate) WindowWillStartLiveResize(notification foundation.Notification) {
	di._WindowWillStartLiveResize(notification)
}
func (di *WindowDelegate) ImplementsWindowDidEndLiveResize() bool {
	return di._WindowDidEndLiveResize != nil
}

func (di *WindowDelegate) SetWindowDidEndLiveResize(f func(notification foundation.Notification)) {
	di._WindowDidEndLiveResize = f
}

func (di *WindowDelegate) WindowDidEndLiveResize(notification foundation.Notification) {
	di._WindowDidEndLiveResize(notification)
}
func (di *WindowDelegate) ImplementsWindowWillMiniaturize() bool {
	return di._WindowWillMiniaturize != nil
}

func (di *WindowDelegate) SetWindowWillMiniaturize(f func(notification foundation.Notification)) {
	di._WindowWillMiniaturize = f
}

func (di *WindowDelegate) WindowWillMiniaturize(notification foundation.Notification) {
	di._WindowWillMiniaturize(notification)
}
func (di *WindowDelegate) ImplementsWindowDidMiniaturize() bool {
	return di._WindowDidMiniaturize != nil
}

func (di *WindowDelegate) SetWindowDidMiniaturize(f func(notification foundation.Notification)) {
	di._WindowDidMiniaturize = f
}

func (di *WindowDelegate) WindowDidMiniaturize(notification foundation.Notification) {
	di._WindowDidMiniaturize(notification)
}
func (di *WindowDelegate) ImplementsWindowDidDeminiaturize() bool {
	return di._WindowDidDeminiaturize != nil
}

func (di *WindowDelegate) SetWindowDidDeminiaturize(f func(notification foundation.Notification)) {
	di._WindowDidDeminiaturize = f
}

func (di *WindowDelegate) WindowDidDeminiaturize(notification foundation.Notification) {
	di._WindowDidDeminiaturize(notification)
}
func (di *WindowDelegate) ImplementsWindowWillUseStandardFrameDefaultFrame() bool {
	return di._WindowWillUseStandardFrameDefaultFrame != nil
}

func (di *WindowDelegate) SetWindowWillUseStandardFrameDefaultFrame(f func(window Window, newFrame foundation.Rect) foundation.Rect) {
	di._WindowWillUseStandardFrameDefaultFrame = f
}

func (di *WindowDelegate) WindowWillUseStandardFrameDefaultFrame(window Window, newFrame foundation.Rect) foundation.Rect {
	return di._WindowWillUseStandardFrameDefaultFrame(window, newFrame)
}
func (di *WindowDelegate) ImplementsWindowShouldZoomToFrame() bool {
	return di._WindowShouldZoomToFrame != nil
}

func (di *WindowDelegate) SetWindowShouldZoomToFrame(f func(window Window, newFrame foundation.Rect) bool) {
	di._WindowShouldZoomToFrame = f
}

func (di *WindowDelegate) WindowShouldZoomToFrame(window Window, newFrame foundation.Rect) bool {
	return di._WindowShouldZoomToFrame(window, newFrame)
}
func (di *WindowDelegate) ImplementsWindowWillUseFullScreenContentSize() bool {
	return di._WindowWillUseFullScreenContentSize != nil
}

func (di *WindowDelegate) SetWindowWillUseFullScreenContentSize(f func(window Window, proposedSize foundation.Size) foundation.Size) {
	di._WindowWillUseFullScreenContentSize = f
}

func (di *WindowDelegate) WindowWillUseFullScreenContentSize(window Window, proposedSize foundation.Size) foundation.Size {
	return di._WindowWillUseFullScreenContentSize(window, proposedSize)
}
func (di *WindowDelegate) ImplementsWindowWillUseFullScreenPresentationOptions() bool {
	return di._WindowWillUseFullScreenPresentationOptions != nil
}

func (di *WindowDelegate) SetWindowWillUseFullScreenPresentationOptions(f func(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions) {
	di._WindowWillUseFullScreenPresentationOptions = f
}

func (di *WindowDelegate) WindowWillUseFullScreenPresentationOptions(window Window, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	return di._WindowWillUseFullScreenPresentationOptions(window, proposedOptions)
}
func (di *WindowDelegate) ImplementsWindowWillEnterFullScreen() bool {
	return di._WindowWillEnterFullScreen != nil
}

func (di *WindowDelegate) SetWindowWillEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillEnterFullScreen = f
}

func (di *WindowDelegate) WindowWillEnterFullScreen(notification foundation.Notification) {
	di._WindowWillEnterFullScreen(notification)
}
func (di *WindowDelegate) ImplementsWindowDidEnterFullScreen() bool {
	return di._WindowDidEnterFullScreen != nil
}

func (di *WindowDelegate) SetWindowDidEnterFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidEnterFullScreen = f
}

func (di *WindowDelegate) WindowDidEnterFullScreen(notification foundation.Notification) {
	di._WindowDidEnterFullScreen(notification)
}
func (di *WindowDelegate) ImplementsWindowWillExitFullScreen() bool {
	return di._WindowWillExitFullScreen != nil
}

func (di *WindowDelegate) SetWindowWillExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowWillExitFullScreen = f
}

func (di *WindowDelegate) WindowWillExitFullScreen(notification foundation.Notification) {
	di._WindowWillExitFullScreen(notification)
}
func (di *WindowDelegate) ImplementsWindowDidExitFullScreen() bool {
	return di._WindowDidExitFullScreen != nil
}

func (di *WindowDelegate) SetWindowDidExitFullScreen(f func(notification foundation.Notification)) {
	di._WindowDidExitFullScreen = f
}

func (di *WindowDelegate) WindowDidExitFullScreen(notification foundation.Notification) {
	di._WindowDidExitFullScreen(notification)
}
func (di *WindowDelegate) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return di._CustomWindowsToEnterFullScreenForWindow != nil
}

func (di *WindowDelegate) SetCustomWindowsToEnterFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindow = f
}

func (di *WindowDelegate) CustomWindowsToEnterFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindow(window)
}
func (di *WindowDelegate) ImplementsCustomWindowsToEnterFullScreenForWindowOnScreen() bool {
	return di._CustomWindowsToEnterFullScreenForWindowOnScreen != nil
}

func (di *WindowDelegate) SetCustomWindowsToEnterFullScreenForWindowOnScreen(f func(window Window, screen Screen) []IWindow) {
	di._CustomWindowsToEnterFullScreenForWindowOnScreen = f
}

func (di *WindowDelegate) CustomWindowsToEnterFullScreenForWindowOnScreen(window Window, screen Screen) []IWindow {
	return di._CustomWindowsToEnterFullScreenForWindowOnScreen(window, screen)
}
func (di *WindowDelegate) ImplementsWindowStartCustomAnimationToEnterFullScreenWithDuration() bool {
	return di._WindowStartCustomAnimationToEnterFullScreenWithDuration != nil
}

func (di *WindowDelegate) SetWindowStartCustomAnimationToEnterFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._WindowStartCustomAnimationToEnterFullScreenWithDuration = f
}

func (di *WindowDelegate) WindowStartCustomAnimationToEnterFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._WindowStartCustomAnimationToEnterFullScreenWithDuration(window, duration)
}
func (di *WindowDelegate) ImplementsWindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration() bool {
	return di._WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration != nil
}

func (di *WindowDelegate) SetWindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(f func(window Window, screen Screen, duration foundation.TimeInterval)) {
	di._WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration = f
}

func (di *WindowDelegate) WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(window Window, screen Screen, duration foundation.TimeInterval) {
	di._WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(window, screen, duration)
}
func (di *WindowDelegate) ImplementsWindowDidFailToEnterFullScreen() bool {
	return di._WindowDidFailToEnterFullScreen != nil
}

func (di *WindowDelegate) SetWindowDidFailToEnterFullScreen(f func(window Window)) {
	di._WindowDidFailToEnterFullScreen = f
}

func (di *WindowDelegate) WindowDidFailToEnterFullScreen(window Window) {
	di._WindowDidFailToEnterFullScreen(window)
}
func (di *WindowDelegate) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return di._CustomWindowsToExitFullScreenForWindow != nil
}

func (di *WindowDelegate) SetCustomWindowsToExitFullScreenForWindow(f func(window Window) []IWindow) {
	di._CustomWindowsToExitFullScreenForWindow = f
}

func (di *WindowDelegate) CustomWindowsToExitFullScreenForWindow(window Window) []IWindow {
	return di._CustomWindowsToExitFullScreenForWindow(window)
}
func (di *WindowDelegate) ImplementsWindowStartCustomAnimationToExitFullScreenWithDuration() bool {
	return di._WindowStartCustomAnimationToExitFullScreenWithDuration != nil
}

func (di *WindowDelegate) SetWindowStartCustomAnimationToExitFullScreenWithDuration(f func(window Window, duration foundation.TimeInterval)) {
	di._WindowStartCustomAnimationToExitFullScreenWithDuration = f
}

func (di *WindowDelegate) WindowStartCustomAnimationToExitFullScreenWithDuration(window Window, duration foundation.TimeInterval) {
	di._WindowStartCustomAnimationToExitFullScreenWithDuration(window, duration)
}
func (di *WindowDelegate) ImplementsWindowDidFailToExitFullScreen() bool {
	return di._WindowDidFailToExitFullScreen != nil
}

func (di *WindowDelegate) SetWindowDidFailToExitFullScreen(f func(window Window)) {
	di._WindowDidFailToExitFullScreen = f
}

func (di *WindowDelegate) WindowDidFailToExitFullScreen(window Window) {
	di._WindowDidFailToExitFullScreen(window)
}
func (di *WindowDelegate) ImplementsWindowWillMove() bool {
	return di._WindowWillMove != nil
}

func (di *WindowDelegate) SetWindowWillMove(f func(notification foundation.Notification)) {
	di._WindowWillMove = f
}

func (di *WindowDelegate) WindowWillMove(notification foundation.Notification) {
	di._WindowWillMove(notification)
}
func (di *WindowDelegate) ImplementsWindowDidMove() bool {
	return di._WindowDidMove != nil
}

func (di *WindowDelegate) SetWindowDidMove(f func(notification foundation.Notification)) {
	di._WindowDidMove = f
}

func (di *WindowDelegate) WindowDidMove(notification foundation.Notification) {
	di._WindowDidMove(notification)
}
func (di *WindowDelegate) ImplementsWindowDidChangeScreen() bool {
	return di._WindowDidChangeScreen != nil
}

func (di *WindowDelegate) SetWindowDidChangeScreen(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreen = f
}

func (di *WindowDelegate) WindowDidChangeScreen(notification foundation.Notification) {
	di._WindowDidChangeScreen(notification)
}
func (di *WindowDelegate) ImplementsWindowDidChangeScreenProfile() bool {
	return di._WindowDidChangeScreenProfile != nil
}

func (di *WindowDelegate) SetWindowDidChangeScreenProfile(f func(notification foundation.Notification)) {
	di._WindowDidChangeScreenProfile = f
}

func (di *WindowDelegate) WindowDidChangeScreenProfile(notification foundation.Notification) {
	di._WindowDidChangeScreenProfile(notification)
}
func (di *WindowDelegate) ImplementsWindowDidChangeBackingProperties() bool {
	return di._WindowDidChangeBackingProperties != nil
}

func (di *WindowDelegate) SetWindowDidChangeBackingProperties(f func(notification foundation.Notification)) {
	di._WindowDidChangeBackingProperties = f
}

func (di *WindowDelegate) WindowDidChangeBackingProperties(notification foundation.Notification) {
	di._WindowDidChangeBackingProperties(notification)
}
func (di *WindowDelegate) ImplementsWindowShouldClose() bool {
	return di._WindowShouldClose != nil
}

func (di *WindowDelegate) SetWindowShouldClose(f func(sender Window) bool) {
	di._WindowShouldClose = f
}

func (di *WindowDelegate) WindowShouldClose(sender Window) bool {
	return di._WindowShouldClose(sender)
}
func (di *WindowDelegate) ImplementsWindowWillClose() bool {
	return di._WindowWillClose != nil
}

func (di *WindowDelegate) SetWindowWillClose(f func(notification foundation.Notification)) {
	di._WindowWillClose = f
}

func (di *WindowDelegate) WindowWillClose(notification foundation.Notification) {
	di._WindowWillClose(notification)
}
func (di *WindowDelegate) ImplementsWindowDidBecomeKey() bool {
	return di._WindowDidBecomeKey != nil
}

func (di *WindowDelegate) SetWindowDidBecomeKey(f func(notification foundation.Notification)) {
	di._WindowDidBecomeKey = f
}

func (di *WindowDelegate) WindowDidBecomeKey(notification foundation.Notification) {
	di._WindowDidBecomeKey(notification)
}
func (di *WindowDelegate) ImplementsWindowDidResignKey() bool {
	return di._WindowDidResignKey != nil
}

func (di *WindowDelegate) SetWindowDidResignKey(f func(notification foundation.Notification)) {
	di._WindowDidResignKey = f
}

func (di *WindowDelegate) WindowDidResignKey(notification foundation.Notification) {
	di._WindowDidResignKey(notification)
}
func (di *WindowDelegate) ImplementsWindowDidBecomeMain() bool {
	return di._WindowDidBecomeMain != nil
}

func (di *WindowDelegate) SetWindowDidBecomeMain(f func(notification foundation.Notification)) {
	di._WindowDidBecomeMain = f
}

func (di *WindowDelegate) WindowDidBecomeMain(notification foundation.Notification) {
	di._WindowDidBecomeMain(notification)
}
func (di *WindowDelegate) ImplementsWindowDidResignMain() bool {
	return di._WindowDidResignMain != nil
}

func (di *WindowDelegate) SetWindowDidResignMain(f func(notification foundation.Notification)) {
	di._WindowDidResignMain = f
}

func (di *WindowDelegate) WindowDidResignMain(notification foundation.Notification) {
	di._WindowDidResignMain(notification)
}
func (di *WindowDelegate) ImplementsWindowWillReturnFieldEditorToObject() bool {
	return di._WindowWillReturnFieldEditorToObject != nil
}

func (di *WindowDelegate) SetWindowWillReturnFieldEditorToObject(f func(sender Window, client objc.Object) objc.IObject) {
	di._WindowWillReturnFieldEditorToObject = f
}

func (di *WindowDelegate) WindowWillReturnFieldEditorToObject(sender Window, client objc.Object) objc.IObject {
	return di._WindowWillReturnFieldEditorToObject(sender, client)
}
func (di *WindowDelegate) ImplementsWindowDidUpdate() bool {
	return di._WindowDidUpdate != nil
}

func (di *WindowDelegate) SetWindowDidUpdate(f func(notification foundation.Notification)) {
	di._WindowDidUpdate = f
}

func (di *WindowDelegate) WindowDidUpdate(notification foundation.Notification) {
	di._WindowDidUpdate(notification)
}
func (di *WindowDelegate) ImplementsWindowDidExpose() bool {
	return di._WindowDidExpose != nil
}

func (di *WindowDelegate) SetWindowDidExpose(f func(notification foundation.Notification)) {
	di._WindowDidExpose = f
}

func (di *WindowDelegate) WindowDidExpose(notification foundation.Notification) {
	di._WindowDidExpose(notification)
}
func (di *WindowDelegate) ImplementsWindowDidChangeOcclusionState() bool {
	return di._WindowDidChangeOcclusionState != nil
}

func (di *WindowDelegate) SetWindowDidChangeOcclusionState(f func(notification foundation.Notification)) {
	di._WindowDidChangeOcclusionState = f
}

func (di *WindowDelegate) WindowDidChangeOcclusionState(notification foundation.Notification) {
	di._WindowDidChangeOcclusionState(notification)
}
func (di *WindowDelegate) ImplementsWindowShouldDragDocumentWithEventFromWithPasteboard() bool {
	return di._WindowShouldDragDocumentWithEventFromWithPasteboard != nil
}

func (di *WindowDelegate) SetWindowShouldDragDocumentWithEventFromWithPasteboard(f func(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool) {
	di._WindowShouldDragDocumentWithEventFromWithPasteboard = f
}

func (di *WindowDelegate) WindowShouldDragDocumentWithEventFromWithPasteboard(window Window, event Event, dragImageLocation foundation.Point, pasteboard Pasteboard) bool {
	return di._WindowShouldDragDocumentWithEventFromWithPasteboard(window, event, dragImageLocation, pasteboard)
}
func (di *WindowDelegate) ImplementsWindowWillReturnUndoManager() bool {
	return di._WindowWillReturnUndoManager != nil
}

func (di *WindowDelegate) SetWindowWillReturnUndoManager(f func(window Window) foundation.IUndoManager) {
	di._WindowWillReturnUndoManager = f
}

func (di *WindowDelegate) WindowWillReturnUndoManager(window Window) foundation.IUndoManager {
	return di._WindowWillReturnUndoManager(window)
}
func (di *WindowDelegate) ImplementsWindowShouldPopUpDocumentPathMenu() bool {
	return di._WindowShouldPopUpDocumentPathMenu != nil
}

func (di *WindowDelegate) SetWindowShouldPopUpDocumentPathMenu(f func(window Window, menu Menu) bool) {
	di._WindowShouldPopUpDocumentPathMenu = f
}

func (di *WindowDelegate) WindowShouldPopUpDocumentPathMenu(window Window, menu Menu) bool {
	return di._WindowShouldPopUpDocumentPathMenu(window, menu)
}
func (di *WindowDelegate) ImplementsWindowWillEncodeRestorableState() bool {
	return di._WindowWillEncodeRestorableState != nil
}

func (di *WindowDelegate) SetWindowWillEncodeRestorableState(f func(window Window, state foundation.Coder)) {
	di._WindowWillEncodeRestorableState = f
}

func (di *WindowDelegate) WindowWillEncodeRestorableState(window Window, state foundation.Coder) {
	di._WindowWillEncodeRestorableState(window, state)
}
func (di *WindowDelegate) ImplementsWindowDidDecodeRestorableState() bool {
	return di._WindowDidDecodeRestorableState != nil
}

func (di *WindowDelegate) SetWindowDidDecodeRestorableState(f func(window Window, state foundation.Coder)) {
	di._WindowDidDecodeRestorableState = f
}

func (di *WindowDelegate) WindowDidDecodeRestorableState(window Window, state foundation.Coder) {
	di._WindowDidDecodeRestorableState(window, state)
}
func (di *WindowDelegate) ImplementsWindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize() bool {
	return di._WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize != nil
}

func (di *WindowDelegate) SetWindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(f func(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size) {
	di._WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize = f
}

func (di *WindowDelegate) WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(window Window, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	return di._WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(window, maxPreferredFrameSize, maxAllowedFrameSize)
}
func (di *WindowDelegate) ImplementsWindowWillEnterVersionBrowser() bool {
	return di._WindowWillEnterVersionBrowser != nil
}

func (di *WindowDelegate) SetWindowWillEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillEnterVersionBrowser = f
}

func (di *WindowDelegate) WindowWillEnterVersionBrowser(notification foundation.Notification) {
	di._WindowWillEnterVersionBrowser(notification)
}
func (di *WindowDelegate) ImplementsWindowDidEnterVersionBrowser() bool {
	return di._WindowDidEnterVersionBrowser != nil
}

func (di *WindowDelegate) SetWindowDidEnterVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidEnterVersionBrowser = f
}

func (di *WindowDelegate) WindowDidEnterVersionBrowser(notification foundation.Notification) {
	di._WindowDidEnterVersionBrowser(notification)
}
func (di *WindowDelegate) ImplementsWindowWillExitVersionBrowser() bool {
	return di._WindowWillExitVersionBrowser != nil
}

func (di *WindowDelegate) SetWindowWillExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowWillExitVersionBrowser = f
}

func (di *WindowDelegate) WindowWillExitVersionBrowser(notification foundation.Notification) {
	di._WindowWillExitVersionBrowser(notification)
}
func (di *WindowDelegate) ImplementsWindowDidExitVersionBrowser() bool {
	return di._WindowDidExitVersionBrowser != nil
}

func (di *WindowDelegate) SetWindowDidExitVersionBrowser(f func(notification foundation.Notification)) {
	di._WindowDidExitVersionBrowser = f
}

func (di *WindowDelegate) WindowDidExitVersionBrowser(notification foundation.Notification) {
	di._WindowDidExitVersionBrowser(notification)
}

type WindowDelegateWrapper struct {
	objc.Object
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillPositionSheetUsingRect() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willPositionSheet:usingRect:"))
}

func (w_ WindowDelegateWrapper) WindowWillPositionSheetUsingRect(window IWindow, sheet IWindow, rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("window:willPositionSheet:usingRect:"), objc.ExtractPtr(window), objc.ExtractPtr(sheet), rect)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillBeginSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillBeginSheet:"))
}

func (w_ WindowDelegateWrapper) WindowWillBeginSheet(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillBeginSheet:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidEndSheet() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndSheet:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndSheet(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidEndSheet:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillResizeToSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillResize:toSize:"))
}

func (w_ WindowDelegateWrapper) WindowWillResizeToSize(sender IWindow, frameSize foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("windowWillResize:toSize:"), objc.ExtractPtr(sender), frameSize)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidResize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillStartLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillStartLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowWillStartLiveResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillStartLiveResize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidEndLiveResize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEndLiveResize:"))
}

func (w_ WindowDelegateWrapper) WindowDidEndLiveResize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidEndLiveResize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowWillMiniaturize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillMiniaturize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidMiniaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMiniaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidMiniaturize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidMiniaturize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidDeminiaturize() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidDeminiaturize:"))
}

func (w_ WindowDelegateWrapper) WindowDidDeminiaturize(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidDeminiaturize:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillUseStandardFrameDefaultFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillUseStandardFrame:defaultFrame:"))
}

func (w_ WindowDelegateWrapper) WindowWillUseStandardFrameDefaultFrame(window IWindow, newFrame foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("windowWillUseStandardFrame:defaultFrame:"), objc.ExtractPtr(window), newFrame)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowShouldZoomToFrame() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldZoom:toFrame:"))
}

func (w_ WindowDelegateWrapper) WindowShouldZoomToFrame(window IWindow, newFrame foundation.Rect) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("windowShouldZoom:toFrame:"), objc.ExtractPtr(window), newFrame)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillUseFullScreenContentSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenContentSize:"))
}

func (w_ WindowDelegateWrapper) WindowWillUseFullScreenContentSize(window IWindow, proposedSize foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("window:willUseFullScreenContentSize:"), objc.ExtractPtr(window), proposedSize)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillUseFullScreenPresentationOptions() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willUseFullScreenPresentationOptions:"))
}

func (w_ WindowDelegateWrapper) WindowWillUseFullScreenPresentationOptions(window IWindow, proposedOptions ApplicationPresentationOptions) ApplicationPresentationOptions {
	rv := objc.CallMethod[ApplicationPresentationOptions](w_, objc.GetSelector("window:willUseFullScreenPresentationOptions:"), objc.ExtractPtr(window), proposedOptions)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterFullScreen(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillEnterFullScreen:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterFullScreen(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidEnterFullScreen:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitFullScreen(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillExitFullScreen:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitFullScreen(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidExitFullScreen:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindow(window IWindow) []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("customWindowsToEnterFullScreenForWindow:"), objc.ExtractPtr(window))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsCustomWindowsToEnterFullScreenForWindowOnScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToEnterFullScreenForWindow:onScreen:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToEnterFullScreenForWindowOnScreen(window IWindow, screen IScreen) []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("customWindowsToEnterFullScreenForWindow:onScreen:"), objc.ExtractPtr(window), objc.ExtractPtr(screen))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowStartCustomAnimationToEnterFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) WindowStartCustomAnimationToEnterFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("window:startCustomAnimationToEnterFullScreenWithDuration:"), objc.ExtractPtr(window), duration)
}

func (w_ WindowDelegateWrapper) ImplementsWindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"))
}

func (w_ WindowDelegateWrapper) WindowStartCustomAnimationToEnterFullScreenOnScreenWithDuration(window IWindow, screen IScreen, duration foundation.TimeInterval) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("window:startCustomAnimationToEnterFullScreenOnScreen:withDuration:"), objc.ExtractPtr(window), objc.ExtractPtr(screen), duration)
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidFailToEnterFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToEnterFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToEnterFullScreen(window IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidFailToEnterFullScreen:"), objc.ExtractPtr(window))
}

func (w_ WindowDelegateWrapper) ImplementsCustomWindowsToExitFullScreenForWindow() bool {
	return w_.RespondsToSelector(objc.GetSelector("customWindowsToExitFullScreenForWindow:"))
}

func (w_ WindowDelegateWrapper) CustomWindowsToExitFullScreenForWindow(window IWindow) []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("customWindowsToExitFullScreenForWindow:"), objc.ExtractPtr(window))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowStartCustomAnimationToExitFullScreenWithDuration() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:startCustomAnimationToExitFullScreenWithDuration:"))
}

func (w_ WindowDelegateWrapper) WindowStartCustomAnimationToExitFullScreenWithDuration(window IWindow, duration foundation.TimeInterval) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("window:startCustomAnimationToExitFullScreenWithDuration:"), objc.ExtractPtr(window), duration)
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidFailToExitFullScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidFailToExitFullScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidFailToExitFullScreen(window IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidFailToExitFullScreen:"), objc.ExtractPtr(window))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillMove:"))
}

func (w_ WindowDelegateWrapper) WindowWillMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillMove:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidMove() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidMove:"))
}

func (w_ WindowDelegateWrapper) WindowDidMove(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidMove:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidChangeScreen() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreen:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreen(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidChangeScreen:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidChangeScreenProfile() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeScreenProfile:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeScreenProfile(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidChangeScreenProfile:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidChangeBackingProperties() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeBackingProperties:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeBackingProperties(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidChangeBackingProperties:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowShouldClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowShouldClose:"))
}

func (w_ WindowDelegateWrapper) WindowShouldClose(sender IWindow) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("windowShouldClose:"), objc.ExtractPtr(sender))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillClose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillClose:"))
}

func (w_ WindowDelegateWrapper) WindowWillClose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillClose:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidBecomeKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeKey(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidBecomeKey:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidResignKey() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignKey:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignKey(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidResignKey:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidBecomeMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidBecomeMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidBecomeMain(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidBecomeMain:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidResignMain() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidResignMain:"))
}

func (w_ WindowDelegateWrapper) WindowDidResignMain(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidResignMain:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillReturnFieldEditorToObject() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnFieldEditor:toObject:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnFieldEditorToObject(sender IWindow, client objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("windowWillReturnFieldEditor:toObject:"), objc.ExtractPtr(sender), objc.ExtractPtr(client))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidUpdate() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidUpdate:"))
}

func (w_ WindowDelegateWrapper) WindowDidUpdate(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidUpdate:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidExpose() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExpose:"))
}

func (w_ WindowDelegateWrapper) WindowDidExpose(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidExpose:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidChangeOcclusionState() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidChangeOcclusionState:"))
}

func (w_ WindowDelegateWrapper) WindowDidChangeOcclusionState(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidChangeOcclusionState:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowShouldDragDocumentWithEventFromWithPasteboard() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldDragDocumentWithEvent:from:withPasteboard:"))
}

func (w_ WindowDelegateWrapper) WindowShouldDragDocumentWithEventFromWithPasteboard(window IWindow, event IEvent, dragImageLocation foundation.Point, pasteboard IPasteboard) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("window:shouldDragDocumentWithEvent:from:withPasteboard:"), objc.ExtractPtr(window), objc.ExtractPtr(event), dragImageLocation, objc.ExtractPtr(pasteboard))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillReturnUndoManager() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillReturnUndoManager:"))
}

func (w_ WindowDelegateWrapper) WindowWillReturnUndoManager(window IWindow) foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](w_, objc.GetSelector("windowWillReturnUndoManager:"), objc.ExtractPtr(window))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowShouldPopUpDocumentPathMenu() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:shouldPopUpDocumentPathMenu:"))
}

func (w_ WindowDelegateWrapper) WindowShouldPopUpDocumentPathMenu(window IWindow, menu IMenu) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("window:shouldPopUpDocumentPathMenu:"), objc.ExtractPtr(window), objc.ExtractPtr(menu))
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillEncodeRestorableState() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willEncodeRestorableState:"))
}

func (w_ WindowDelegateWrapper) WindowWillEncodeRestorableState(window IWindow, state foundation.ICoder) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("window:willEncodeRestorableState:"), objc.ExtractPtr(window), objc.ExtractPtr(state))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidDecodeRestorableState() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:didDecodeRestorableState:"))
}

func (w_ WindowDelegateWrapper) WindowDidDecodeRestorableState(window IWindow, state foundation.ICoder) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("window:didDecodeRestorableState:"), objc.ExtractPtr(window), objc.ExtractPtr(state))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize() bool {
	return w_.RespondsToSelector(objc.GetSelector("window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:"))
}

func (w_ WindowDelegateWrapper) WindowWillResizeForVersionBrowserWithMaxPreferredSizeMaxAllowedSize(window IWindow, maxPreferredFrameSize foundation.Size, maxAllowedFrameSize foundation.Size) foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("window:willResizeForVersionBrowserWithMaxPreferredSize:maxAllowedSize:"), objc.ExtractPtr(window), maxPreferredFrameSize, maxAllowedFrameSize)
	return rv
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillEnterVersionBrowser(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillEnterVersionBrowser:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidEnterVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidEnterVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidEnterVersionBrowser(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidEnterVersionBrowser:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowWillExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowWillExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowWillExitVersionBrowser(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowWillExitVersionBrowser:"), objc.ExtractPtr(notification))
}

func (w_ WindowDelegateWrapper) ImplementsWindowDidExitVersionBrowser() bool {
	return w_.RespondsToSelector(objc.GetSelector("windowDidExitVersionBrowser:"))
}

func (w_ WindowDelegateWrapper) WindowDidExitVersionBrowser(notification foundation.INotification) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("windowDidExitVersionBrowser:"), objc.ExtractPtr(notification))
}
