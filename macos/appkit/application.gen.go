// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var ApplicationClass = _ApplicationClass{objc.GetClass("NSApplication")}

type _ApplicationClass struct {
	objc.Class
}

type IApplication interface {
	IResponder
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	Run()
	FinishLaunching()
	Stop(sender objc.IObject)
	SendEvent(event IEvent)
	PostEventAtStart(event IEvent, flag bool)
	SendActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) bool
	TargetForAction(action objc.Selector) objc.Object
	TargetForActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object
	Terminate(sender objc.IObject)
	ReplyToApplicationShouldTerminate(shouldTerminate bool)
	ActivateIgnoringOtherApps(flag bool)
	Deactivate()
	DisableRelaunchOnLogin()
	EnableRelaunchOnLogin()
	RegisterForRemoteNotifications()
	UnregisterForRemoteNotifications()
	RegisterForRemoteNotificationTypes(types RemoteNotificationType)
	ToggleTouchBarCustomizationPalette(sender objc.IObject)
	RequestUserAttention(requestType RequestUserAttentionType) int
	CancelUserAttentionRequest(request int)
	ReplyToOpenOrPrint(reply ApplicationDelegateReply)
	RegisterUserInterfaceItemSearchHandler(handler IUserInterfaceItemSearching)
	RegisterUserInterfaceItemSearchHandler0(handler objc.IObject)
	SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool
	UnregisterUserInterfaceItemSearchHandler(handler IUserInterfaceItemSearching)
	UnregisterUserInterfaceItemSearchHandler0(handler objc.IObject)
	ShowHelp(sender objc.IObject)
	ActivateContextHelpMode(sender objc.IObject)
	HideOtherApplications(sender objc.IObject)
	UnhideAllApplications(sender objc.IObject)
	ReportException(exception foundation.IException)
	ActivationPolicy() ApplicationActivationPolicy
	SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool
	WindowWithWindowNumber(windowNum int) Window
	EnumerateWindowsWithOptionsUsingBlock(options WindowListOptions, block func(window Window, stop *bool))
	MiniaturizeAll(sender objc.IObject)
	Hide(sender objc.IObject)
	Unhide(sender objc.IObject)
	UnhideWithoutActivation()
	UpdateWindows()
	SetWindowsNeedUpdate(needUpdate bool)
	PreventWindowOrdering()
	ArrangeInFront(sender objc.IObject)
	ExtendStateRestoration()
	CompleteStateRestoration()
	RestoreWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) bool
	RunModalForWindow(window IWindow) ModalResponse
	StopModal()
	StopModalWithCode(returnCode ModalResponse)
	AbortModal()
	BeginModalSessionForWindow(window IWindow) ModalSession
	RunModalSession(session ModalSession) ModalResponse
	OrderFrontColorPanel(sender objc.IObject)
	OrderFrontStandardAboutPanel(sender objc.IObject)
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject)
	OrderFrontCharacterPalette(sender objc.IObject)
	RunPageLayout(sender objc.IObject)
	AddWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool)
	ChangeWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool)
	RemoveWindowsItem(win IWindow)
	UpdateWindowsItem(win IWindow)
	RegisterServicesMenuSendTypesReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType)
	EndModalSession(session ModalSession)
	Delegate() ApplicationDelegateWrapper
	SetDelegate(value IApplicationDelegate)
	SetDelegate0(value objc.IObject)
	CurrentEvent() Event
	IsRunning() bool
	IsActive() bool
	EnabledRemoteNotificationTypes() RemoteNotificationType
	IsRegisteredForRemoteNotifications() bool
	Appearance() Appearance
	SetAppearance(value IAppearance)
	EffectiveAppearance() Appearance
	CurrentSystemPresentationOptions() ApplicationPresentationOptions
	PresentationOptions() ApplicationPresentationOptions
	SetPresentationOptions(value ApplicationPresentationOptions)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	DockTile() DockTile
	ApplicationIconImage() Image
	SetApplicationIconImage(value IImage)
	HelpMenu() Menu
	SetHelpMenu(value IMenu)
	ServicesProvider() objc.Object
	SetServicesProvider(value objc.IObject)
	IsFullKeyboardAccessEnabled() bool
	OrderedDocuments() []Document
	OrderedWindows() []Window
	KeyWindow() Window
	MainWindow() Window
	Windows() []Window
	IsHidden() bool
	OcclusionState() ApplicationOcclusionState
	IsProtectedDataAvailable() bool
	ModalWindow() Window
	MainMenu() Menu
	SetMainMenu(value IMenu)
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	WindowsMenu() Menu
	SetWindowsMenu(value IMenu)
	ServicesMenu() Menu
	SetServicesMenu(value IMenu)
}

type Application struct {
	Responder
}

func MakeApplication(ptr unsafe.Pointer) Application {
	return Application{
		Responder: MakeResponder(ptr),
	}
}

func (a_ Application) Init() Application {
	rv := objc.CallMethod[Application](a_, objc.GetSelector("init"))
	return rv
}

func Application_Init() Application {
	return ApplicationClass.Alloc().Init()
}

func (ac _ApplicationClass) Alloc() Application {
	rv := objc.CallMethod[Application](ac, objc.GetSelector("alloc"))
	return rv
}

func Application_Alloc() Application {
	return ApplicationClass.Alloc()
}

func (ac _ApplicationClass) New() Application {
	rv := objc.CallMethod[Application](ac, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewApplication() Application {
	return ApplicationClass.New()
}

func Application_New() Application {
	return ApplicationClass.New()
}

func (a_ Application) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.CallMethod[Event](a_, objc.GetSelector("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, objc.ExtractPtr(expiration), mode, deqFlag)
	return rv
}

func (a_ Application) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("discardEventsMatchingMask:beforeEvent:"), mask, objc.ExtractPtr(lastEvent))
}

func (a_ Application) Run() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("run"))
}

func (a_ Application) FinishLaunching() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("finishLaunching"))
}

func (a_ Application) Stop(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stop:"), objc.ExtractPtr(sender))
}

func (a_ Application) SendEvent(event IEvent) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("sendEvent:"), objc.ExtractPtr(event))
}

func (a_ Application) PostEventAtStart(event IEvent, flag bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("postEvent:atStart:"), objc.ExtractPtr(event), flag)
}

func (a_ Application) SendActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("sendAction:to:from:"), action, objc.ExtractPtr(target), objc.ExtractPtr(sender))
	return rv
}

func (a_ Application) TargetForAction(action objc.Selector) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("targetForAction:"), action)
	return rv
}

func (a_ Application) TargetForActionToFrom(action objc.Selector, target objc.IObject, sender objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("targetForAction:to:from:"), action, objc.ExtractPtr(target), objc.ExtractPtr(sender))
	return rv
}

func (a_ Application) Terminate(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("terminate:"), objc.ExtractPtr(sender))
}

func (a_ Application) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("replyToApplicationShouldTerminate:"), shouldTerminate)
}

func (a_ Application) ActivateIgnoringOtherApps(flag bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("activateIgnoringOtherApps:"), flag)
}

func (a_ Application) Deactivate() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("deactivate"))
}

func (a_ Application) DisableRelaunchOnLogin() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("disableRelaunchOnLogin"))
}

func (a_ Application) EnableRelaunchOnLogin() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("enableRelaunchOnLogin"))
}

func (a_ Application) RegisterForRemoteNotifications() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("registerForRemoteNotifications"))
}

func (a_ Application) UnregisterForRemoteNotifications() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unregisterForRemoteNotifications"))
}

func (a_ Application) RegisterForRemoteNotificationTypes(types RemoteNotificationType) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("registerForRemoteNotificationTypes:"), types)
}

func (a_ Application) ToggleTouchBarCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("toggleTouchBarCustomizationPalette:"), objc.ExtractPtr(sender))
}

func (a_ Application) RequestUserAttention(requestType RequestUserAttentionType) int {
	rv := objc.CallMethod[int](a_, objc.GetSelector("requestUserAttention:"), requestType)
	return rv
}

func (a_ Application) CancelUserAttentionRequest(request int) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("cancelUserAttentionRequest:"), request)
}

func (a_ Application) ReplyToOpenOrPrint(reply ApplicationDelegateReply) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("replyToOpenOrPrint:"), reply)
}

func (a_ Application) RegisterUserInterfaceItemSearchHandler(handler IUserInterfaceItemSearching) {
	po := objc.WrapAsProtocol("NSUserInterfaceItemSearching", handler)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("registerUserInterfaceItemSearchHandler:"), po)
}

func (a_ Application) RegisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("registerUserInterfaceItemSearchHandler:"), objc.ExtractPtr(handler))
}

func (a_ Application) SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.Range, foundRange *foundation.Range) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("searchString:inUserInterfaceItemString:searchRange:foundRange:"), searchString, stringToSearch, searchRange, foundRange)
	return rv
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler(handler IUserInterfaceItemSearching) {
	po := objc.WrapAsProtocol("NSUserInterfaceItemSearching", handler)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unregisterUserInterfaceItemSearchHandler:"), po)
}

func (a_ Application) UnregisterUserInterfaceItemSearchHandler0(handler objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unregisterUserInterfaceItemSearchHandler:"), objc.ExtractPtr(handler))
}

func (a_ Application) ShowHelp(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("showHelp:"), objc.ExtractPtr(sender))
}

func (a_ Application) ActivateContextHelpMode(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("activateContextHelpMode:"), objc.ExtractPtr(sender))
}

func (a_ Application) HideOtherApplications(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("hideOtherApplications:"), objc.ExtractPtr(sender))
}

func (a_ Application) UnhideAllApplications(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unhideAllApplications:"), objc.ExtractPtr(sender))
}

func (ac _ApplicationClass) DetachDrawingThreadToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	objc.CallMethod[objc.Void](ac, objc.GetSelector("detachDrawingThread:toTarget:withObject:"), selector, objc.ExtractPtr(target), objc.ExtractPtr(argument))
}

func Application_DetachDrawingThreadToTargetWithObject(selector objc.Selector, target objc.IObject, argument objc.IObject) {
	ApplicationClass.DetachDrawingThreadToTargetWithObject(selector, target, argument)
}

func (a_ Application) ReportException(exception foundation.IException) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("reportException:"), objc.ExtractPtr(exception))
}

func (a_ Application) ActivationPolicy() ApplicationActivationPolicy {
	rv := objc.CallMethod[ApplicationActivationPolicy](a_, objc.GetSelector("activationPolicy"))
	return rv
}

func (a_ Application) SetActivationPolicy(activationPolicy ApplicationActivationPolicy) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("setActivationPolicy:"), activationPolicy)
	return rv
}

func (a_ Application) WindowWithWindowNumber(windowNum int) Window {
	rv := objc.CallMethod[Window](a_, objc.GetSelector("windowWithWindowNumber:"), windowNum)
	return rv
}

func (a_ Application) EnumerateWindowsWithOptionsUsingBlock(options WindowListOptions, block func(window Window, stop *bool)) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("enumerateWindowsWithOptions:usingBlock:"), options, block)
}

func (a_ Application) MiniaturizeAll(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("miniaturizeAll:"), objc.ExtractPtr(sender))
}

func (a_ Application) Hide(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("hide:"), objc.ExtractPtr(sender))
}

func (a_ Application) Unhide(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unhide:"), objc.ExtractPtr(sender))
}

func (a_ Application) UnhideWithoutActivation() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("unhideWithoutActivation"))
}

func (a_ Application) UpdateWindows() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("updateWindows"))
}

func (a_ Application) SetWindowsNeedUpdate(needUpdate bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setWindowsNeedUpdate:"), needUpdate)
}

func (a_ Application) PreventWindowOrdering() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("preventWindowOrdering"))
}

func (a_ Application) ArrangeInFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("arrangeInFront:"), objc.ExtractPtr(sender))
}

func (a_ Application) ExtendStateRestoration() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("extendStateRestoration"))
}

func (a_ Application) CompleteStateRestoration() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("completeStateRestoration"))
}

func (a_ Application) RestoreWindowWithIdentifierStateCompletionHandler(identifier UserInterfaceItemIdentifier, state foundation.ICoder, completionHandler func(param1 Window, param2 foundation.Error)) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("restoreWindowWithIdentifier:state:completionHandler:"), identifier, objc.ExtractPtr(state), completionHandler)
	return rv
}

func (a_ Application) RunModalForWindow(window IWindow) ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, objc.GetSelector("runModalForWindow:"), objc.ExtractPtr(window))
	return rv
}

func (a_ Application) StopModal() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopModal"))
}

func (a_ Application) StopModalWithCode(returnCode ModalResponse) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("stopModalWithCode:"), returnCode)
}

func (a_ Application) AbortModal() {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("abortModal"))
}

func (a_ Application) BeginModalSessionForWindow(window IWindow) ModalSession {
	rv := objc.CallMethod[ModalSession](a_, objc.GetSelector("beginModalSessionForWindow:"), objc.ExtractPtr(window))
	return rv
}

func (a_ Application) RunModalSession(session ModalSession) ModalResponse {
	rv := objc.CallMethod[ModalResponse](a_, objc.GetSelector("runModalSession:"), session)
	return rv
}

func (a_ Application) OrderFrontColorPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("orderFrontColorPanel:"), objc.ExtractPtr(sender))
}

func (a_ Application) OrderFrontStandardAboutPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("orderFrontStandardAboutPanel:"), objc.ExtractPtr(sender))
}

func (a_ Application) OrderFrontStandardAboutPanelWithOptions(optionsDictionary map[AboutPanelOptionKey]objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("orderFrontStandardAboutPanelWithOptions:"), optionsDictionary)
}

func (a_ Application) OrderFrontCharacterPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("orderFrontCharacterPalette:"), objc.ExtractPtr(sender))
}

func (a_ Application) RunPageLayout(sender objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("runPageLayout:"), objc.ExtractPtr(sender))
}

func (a_ Application) AddWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("addWindowsItem:title:filename:"), objc.ExtractPtr(win), string_, isFilename)
}

func (a_ Application) ChangeWindowsItemTitleFilename(win IWindow, string_ string, isFilename bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("changeWindowsItem:title:filename:"), objc.ExtractPtr(win), string_, isFilename)
}

func (a_ Application) RemoveWindowsItem(win IWindow) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("removeWindowsItem:"), objc.ExtractPtr(win))
}

func (a_ Application) UpdateWindowsItem(win IWindow) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("updateWindowsItem:"), objc.ExtractPtr(win))
}

func (a_ Application) RegisterServicesMenuSendTypesReturnTypes(sendTypes []PasteboardType, returnTypes []PasteboardType) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("registerServicesMenuSendTypes:returnTypes:"), sendTypes, returnTypes)
}

func (a_ Application) EndModalSession(session ModalSession) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("endModalSession:"), session)
}

func (ac _ApplicationClass) SharedApplication() Application {
	rv := objc.CallMethod[Application](ac, objc.GetSelector("sharedApplication"))
	return rv
}

func Application_SharedApplication() Application {
	return ApplicationClass.SharedApplication()
}

func (a_ Application) Delegate() ApplicationDelegateWrapper {
	rv := objc.CallMethod[ApplicationDelegateWrapper](a_, objc.GetSelector("delegate"))
	return rv
}

func (a_ Application) SetDelegate(value IApplicationDelegate) {
	po := objc.WrapAsProtocol("NSApplicationDelegate", value)
	objc.SetAssociatedObject(a_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), po)
}

func (a_ Application) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (a_ Application) CurrentEvent() Event {
	rv := objc.CallMethod[Event](a_, objc.GetSelector("currentEvent"))
	return rv
}

func (a_ Application) IsRunning() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isRunning"))
	return rv
}

func (a_ Application) IsActive() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isActive"))
	return rv
}

func (a_ Application) EnabledRemoteNotificationTypes() RemoteNotificationType {
	rv := objc.CallMethod[RemoteNotificationType](a_, objc.GetSelector("enabledRemoteNotificationTypes"))
	return rv
}

func (a_ Application) IsRegisteredForRemoteNotifications() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isRegisteredForRemoteNotifications"))
	return rv
}

func (a_ Application) Appearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("appearance"))
	return rv
}

func (a_ Application) SetAppearance(value IAppearance) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAppearance:"), objc.ExtractPtr(value))
}

func (a_ Application) EffectiveAppearance() Appearance {
	rv := objc.CallMethod[Appearance](a_, objc.GetSelector("effectiveAppearance"))
	return rv
}

func (a_ Application) CurrentSystemPresentationOptions() ApplicationPresentationOptions {
	rv := objc.CallMethod[ApplicationPresentationOptions](a_, objc.GetSelector("currentSystemPresentationOptions"))
	return rv
}

func (a_ Application) PresentationOptions() ApplicationPresentationOptions {
	rv := objc.CallMethod[ApplicationPresentationOptions](a_, objc.GetSelector("presentationOptions"))
	return rv
}

func (a_ Application) SetPresentationOptions(value ApplicationPresentationOptions) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setPresentationOptions:"), value)
}

func (a_ Application) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](a_, objc.GetSelector("userInterfaceLayoutDirection"))
	return rv
}

func (a_ Application) DockTile() DockTile {
	rv := objc.CallMethod[DockTile](a_, objc.GetSelector("dockTile"))
	return rv
}

func (a_ Application) ApplicationIconImage() Image {
	rv := objc.CallMethod[Image](a_, objc.GetSelector("applicationIconImage"))
	return rv
}

func (a_ Application) SetApplicationIconImage(value IImage) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setApplicationIconImage:"), objc.ExtractPtr(value))
}

func (a_ Application) HelpMenu() Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("helpMenu"))
	return rv
}

func (a_ Application) SetHelpMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setHelpMenu:"), objc.ExtractPtr(value))
}

func (a_ Application) ServicesProvider() objc.Object {
	rv := objc.CallMethod[objc.Object](a_, objc.GetSelector("servicesProvider"))
	return rv
}

func (a_ Application) SetServicesProvider(value objc.IObject) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setServicesProvider:"), objc.ExtractPtr(value))
}

func (a_ Application) IsFullKeyboardAccessEnabled() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isFullKeyboardAccessEnabled"))
	return rv
}

func (a_ Application) OrderedDocuments() []Document {
	rv := objc.CallMethod[[]Document](a_, objc.GetSelector("orderedDocuments"))
	return rv
}

func (a_ Application) OrderedWindows() []Window {
	rv := objc.CallMethod[[]Window](a_, objc.GetSelector("orderedWindows"))
	return rv
}

func (a_ Application) KeyWindow() Window {
	rv := objc.CallMethod[Window](a_, objc.GetSelector("keyWindow"))
	return rv
}

func (a_ Application) MainWindow() Window {
	rv := objc.CallMethod[Window](a_, objc.GetSelector("mainWindow"))
	return rv
}

func (a_ Application) Windows() []Window {
	rv := objc.CallMethod[[]Window](a_, objc.GetSelector("windows"))
	return rv
}

func (a_ Application) IsHidden() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isHidden"))
	return rv
}

func (a_ Application) OcclusionState() ApplicationOcclusionState {
	rv := objc.CallMethod[ApplicationOcclusionState](a_, objc.GetSelector("occlusionState"))
	return rv
}

func (a_ Application) IsProtectedDataAvailable() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isProtectedDataAvailable"))
	return rv
}

func (a_ Application) ModalWindow() Window {
	rv := objc.CallMethod[Window](a_, objc.GetSelector("modalWindow"))
	return rv
}

func (a_ Application) MainMenu() Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("mainMenu"))
	return rv
}

func (a_ Application) SetMainMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setMainMenu:"), objc.ExtractPtr(value))
}

func (a_ Application) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}

func (a_ Application) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

func (a_ Application) WindowsMenu() Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("windowsMenu"))
	return rv
}

func (a_ Application) SetWindowsMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setWindowsMenu:"), objc.ExtractPtr(value))
}

func (a_ Application) ServicesMenu() Menu {
	rv := objc.CallMethod[Menu](a_, objc.GetSelector("servicesMenu"))
	return rv
}

func (a_ Application) SetServicesMenu(value IMenu) {
	objc.CallMethod[objc.Void](a_, objc.GetSelector("setServicesMenu:"), objc.ExtractPtr(value))
}
