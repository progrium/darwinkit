// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var WindowClass = _WindowClass{objc.GetClass("NSWindow")}

type _WindowClass struct {
	objc.Class
}

type IWindow interface {
	IResponder
	ToggleFullScreen(sender objc.IObject)
	SetDynamicDepthLimit(flag bool)
	InvalidateShadow()
	AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool
	SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.RectEdge)
	ContentBorderThicknessForEdge(edge foundation.RectEdge) float64
	SetContentBorderThicknessForEdge(thickness float64, edge foundation.RectEdge)
	ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect
	FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect
	BeginSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	EndSheet(sheetWindow IWindow)
	EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse)
	SetFrameOrigin(point foundation.Point)
	SetFrameTopLeftPoint(point foundation.Point)
	ConstrainFrameRectToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect
	CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point
	SetFrameDisplay(frameRect foundation.Rect, flag bool)
	SetFrameDisplayAnimate(frameRect foundation.Rect, displayFlag bool, animateFlag bool)
	AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval
	PerformZoom(sender objc.IObject)
	Zoom(sender objc.IObject)
	SetContentSize(size foundation.Size)
	OrderOut(sender objc.IObject)
	OrderBack(sender objc.IObject)
	OrderFront(sender objc.IObject)
	OrderFrontRegardless()
	OrderWindowRelativeTo(place WindowOrderingMode, otherWin int)
	SetFrameUsingName(name WindowFrameAutosaveName) bool
	SetFrameUsingNameForce(name WindowFrameAutosaveName, force bool) bool
	SaveFrameUsingName(name WindowFrameAutosaveName)
	SetFrameAutosaveName(name WindowFrameAutosaveName) bool
	SetFrameFromString(string_ WindowPersistableFrameDescriptor)
	MakeKeyWindow()
	MakeKeyAndOrderFront(sender objc.IObject)
	BecomeKeyWindow()
	ResignKeyWindow()
	MakeMainWindow()
	BecomeMainWindow()
	ResignMainWindow()
	ToggleToolbarShown(sender objc.IObject)
	RunToolbarCustomizationPalette(sender objc.IObject)
	AddChildWindowOrdered(childWin IWindow, place WindowOrderingMode)
	RemoveChildWindow(childWin IWindow)
	EnableKeyEquivalentForDefaultButtonCell()
	DisableKeyEquivalentForDefaultButtonCell()
	FieldEditorForObject(createFlag bool, object objc.IObject) Text
	EndEditingFor(object objc.IObject)
	EnableCursorRects()
	DisableCursorRects()
	DiscardCursorRects()
	InvalidateCursorRectsForView(view IView)
	ResetCursorRects()
	StandardWindowButton(b WindowButton) Button
	AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController)
	InsertTitlebarAccessoryViewControllerAtIndex(childViewController ITitlebarAccessoryViewController, index int)
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode)
	MergeAllWindows(sender objc.IObject)
	SelectNextTab(sender objc.IObject)
	SelectPreviousTab(sender objc.IObject)
	MoveTabToNewWindow(sender objc.IObject)
	ToggleTabBar(sender objc.IObject)
	ToggleTabOverview(sender objc.IObject)
	NextEventMatchingMask(mask EventMask) Event
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	PostEventAtStart(event IEvent, flag bool)
	SendEvent(event IEvent)
	MakeFirstResponder(responder IResponder) bool
	SelectKeyViewPrecedingView(view IView)
	SelectKeyViewFollowingView(view IView)
	SelectPreviousKeyView(sender objc.IObject)
	SelectNextKeyView(sender objc.IObject)
	RecalculateKeyViewLoop()
	TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool))
	PerformWindowDragWithEvent(event IEvent)
	DisableSnapshotRestoration()
	EnableSnapshotRestoration()
	Display()
	DisplayIfNeeded()
	DisableScreenUpdatesUntilFlush()
	Update()
	DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	UnregisterDraggedTypes()
	BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectFromScreen(rect foundation.Rect) foundation.Rect
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	ConvertPointFromScreen(point foundation.Point) foundation.Point
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToScreen(rect foundation.Rect) foundation.Rect
	ConvertPointToBacking(point foundation.Point) foundation.Point
	ConvertPointToScreen(point foundation.Point) foundation.Point
	SetTitleWithRepresentedFilename(filename string)
	Center()
	PerformClose(sender objc.IObject)
	Close()
	PerformMiniaturize(sender objc.IObject)
	Miniaturize(sender objc.IObject)
	Deminiaturize(sender objc.IObject)
	Print(sender objc.IObject)
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	UpdateConstraintsIfNeeded()
	LayoutIfNeeded()
	VisualizeConstraints(constraints []ILayoutConstraint)
	AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute
	SetAnchorAttributeForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation)
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	SetIsMiniaturized(flag bool)
	SetIsVisible(flag bool)
	SetIsZoomed(flag bool)
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	InitWithWindowRef(windowRef unsafe.Pointer) Window
	Delegate() WindowDelegateWrapper
	SetDelegate(value IWindowDelegate)
	SetDelegate0(value objc.IObject)
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	ContentView() View
	SetContentView(value IView)
	StyleMask() WindowStyleMask
	SetStyleMask(value WindowStyleMask)
	WorksWhenModal() bool
	AlphaValue() float64
	SetAlphaValue(value float64)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	ColorSpace() ColorSpace
	SetColorSpace(value IColorSpace)
	CanHide() bool
	SetCanHide(value bool)
	IsOnActiveSpace() bool
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	CollectionBehavior() WindowCollectionBehavior
	SetCollectionBehavior(value WindowCollectionBehavior)
	IsOpaque() bool
	SetOpaque(value bool)
	HasShadow() bool
	SetHasShadow(value bool)
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	DepthLimit() WindowDepth
	SetDepthLimit(value WindowDepth)
	HasDynamicDepthLimit() bool
	WindowNumber() int
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	SharingType() WindowSharingType
	SetSharingType(value WindowSharingType)
	BackingType() BackingStoreType
	SetBackingType(value BackingStoreType)
	WindowController() WindowController
	SetWindowController(value IWindowController)
	AttachedSheet() Window
	IsSheet() bool
	SheetParent() Window
	Sheets() []Window
	Frame() foundation.Rect
	AspectRatio() foundation.Size
	SetAspectRatio(value foundation.Size)
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	IsZoomed() bool
	ResizeFlags() EventModifierFlags
	ResizeIncrements() foundation.Size
	SetResizeIncrements(value foundation.Size)
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	InLiveResize() bool
	ContentAspectRatio() foundation.Size
	SetContentAspectRatio(value foundation.Size)
	ContentMinSize() foundation.Size
	SetContentMinSize(value foundation.Size)
	ContentMaxSize() foundation.Size
	SetContentMaxSize(value foundation.Size)
	ContentResizeIncrements() foundation.Size
	SetContentResizeIncrements(value foundation.Size)
	ContentLayoutGuide() objc.Object
	ContentLayoutRect() foundation.Rect
	MaxFullScreenContentSize() foundation.Size
	SetMaxFullScreenContentSize(value foundation.Size)
	MinFullScreenContentSize() foundation.Size
	SetMinFullScreenContentSize(value foundation.Size)
	Level() WindowLevel
	SetLevel(value WindowLevel)
	IsVisible() bool
	OcclusionState() WindowOcclusionState
	FrameAutosaveName() WindowFrameAutosaveName
	StringWithSavedFrame() WindowPersistableFrameDescriptor
	IsKeyWindow() bool
	CanBecomeKeyWindow() bool
	IsMainWindow() bool
	CanBecomeMainWindow() bool
	Toolbar() Toolbar
	SetToolbar(value IToolbar)
	ChildWindows() []Window
	ParentWindow() Window
	SetParentWindow(value IWindow)
	DefaultButtonCell() ButtonCell
	SetDefaultButtonCell(value IButtonCell)
	IsExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)
	AreCursorRectsEnabled() bool
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	ToolbarStyle() WindowToolbarStyle
	SetToolbarStyle(value WindowToolbarStyle)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection
	TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController)
	Tab() WindowTab
	TabbingIdentifier() WindowTabbingIdentifier
	SetTabbingIdentifier(value WindowTabbingIdentifier)
	TabbingMode() WindowTabbingMode
	SetTabbingMode(value WindowTabbingMode)
	TabbedWindows() []Window
	TabGroup() WindowTabGroup
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)
	CurrentEvent() Event
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	FirstResponder() Responder
	KeyViewSelectionDirection() SelectionDirection
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	MouseLocationOutsideOfEventStream() foundation.Point
	IsRestorable() bool
	SetRestorable(value bool)
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)
	AnimationBehavior() WindowAnimationBehavior
	SetAnimationBehavior(value WindowAnimationBehavior)
	IsDocumentEdited() bool
	SetDocumentEdited(value bool)
	BackingScaleFactor() float64
	Title() string
	SetTitle(value string)
	Subtitle() string
	SetSubtitle(value string)
	TitleVisibility() WindowTitleVisibility
	SetTitleVisibility(value WindowTitleVisibility)
	RepresentedFilename() string
	SetRepresentedFilename(value string)
	RepresentedURL() foundation.URL
	SetRepresentedURL(value foundation.IURL)
	Screen() Screen
	DeepestScreen() Screen
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)
	IsMovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	IsMovable() bool
	SetMovable(value bool)
	IsReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)
	IsMiniaturized() bool
	MiniwindowImage() Image
	SetMiniwindowImage(value IImage)
	MiniwindowTitle() string
	SetMiniwindowTitle(value string)
	DockTile() DockTile
	HasCloseBox() bool
	HasTitleBar() bool
	IsModalPanel() bool
	IsFloatingPanel() bool
	IsZoomable() bool
	IsResizable() bool
	IsMiniaturizable() bool
	OrderedIndex() int
	SetOrderedIndex(value int)
	WindowRef() unsafe.Pointer
}

type Window struct {
	Responder
}

func MakeWindow(ptr unsafe.Pointer) Window {
	return Window{
		Responder: MakeResponder(ptr),
	}
}

func (wc _WindowClass) WindowWithContentViewController(contentViewController IViewController) Window {
	rv := objc.CallMethod[Window](wc, objc.GetSelector("windowWithContentViewController:"), objc.ExtractPtr(contentViewController))
	return rv
}

func Window_WindowWithContentViewController(contentViewController IViewController) Window {
	return WindowClass.WindowWithContentViewController(contentViewController)
}

func (w_ Window) InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}

func Window_InitWithContentRectStyleMaskBackingDefer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	return WindowClass.Alloc().InitWithContentRectStyleMaskBackingDefer(contentRect, style, backingStoreType, flag)
}

func (w_ Window) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.ExtractPtr(screen))
	return rv
}

func Window_InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	return WindowClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
}

func (w_ Window) Init() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("init"))
	return rv
}

func Window_Init() Window {
	return WindowClass.Alloc().Init()
}

func (wc _WindowClass) Alloc() Window {
	rv := objc.CallMethod[Window](wc, objc.GetSelector("alloc"))
	return rv
}

func Window_Alloc() Window {
	return WindowClass.Alloc()
}

func (wc _WindowClass) New() Window {
	rv := objc.CallMethod[Window](wc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewWindow() Window {
	return WindowClass.New()
}

func Window_New() Window {
	return WindowClass.New()
}

func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("toggleFullScreen:"), objc.ExtractPtr(sender))
}

func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDynamicDepthLimit:"), flag)
}

func (w_ Window) InvalidateShadow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("invalidateShadow"))
}

func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}

func (w_ Window) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}

func (w_ Window) ContentBorderThicknessForEdge(edge foundation.RectEdge) float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("contentBorderThicknessForEdge:"), edge)
	return rv
}

func (w_ Window) SetContentBorderThicknessForEdge(thickness float64, edge foundation.RectEdge) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentBorderThickness:forEdge:"), thickness, edge)
}

func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := objc.CallMethod[[]foundation.Number](wc, objc.GetSelector("windowNumbersWithOptions:"), options)
	return rv
}

func Window_WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	return WindowClass.WindowNumbersWithOptions(options)
}

func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, objc.GetSelector("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}

func Window_ContentRectForFrameRectStyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	return WindowClass.ContentRectForFrameRectStyleMask(fRect, style)
}

func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](wc, objc.GetSelector("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}

func Window_FrameRectForContentRectStyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	return WindowClass.FrameRectForContentRectStyleMask(cRect, style)
}

func (wc _WindowClass) MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	rv := objc.CallMethod[float64](wc, objc.GetSelector("minFrameWidthWithTitle:styleMask:"), title, style)
	return rv
}

func Window_MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	return WindowClass.MinFrameWidthWithTitleStyleMask(title, style)
}

func (w_ Window) ContentRectForFrameRect(frameRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("contentRectForFrameRect:"), frameRect)
	return rv
}

func (w_ Window) FrameRectForContentRect(contentRect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("frameRectForContentRect:"), contentRect)
	return rv
}

func (w_ Window) BeginSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("beginSheet:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("beginCriticalSheet:completionHandler:"), objc.ExtractPtr(sheetWindow), handler)
}

func (w_ Window) EndSheet(sheetWindow IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("endSheet:"), objc.ExtractPtr(sheetWindow))
}

func (w_ Window) EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("endSheet:returnCode:"), objc.ExtractPtr(sheetWindow), returnCode)
}

func (w_ Window) SetFrameOrigin(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setFrameOrigin:"), point)
}

func (w_ Window) SetFrameTopLeftPoint(point foundation.Point) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setFrameTopLeftPoint:"), point)
}

func (w_ Window) ConstrainFrameRectToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("constrainFrameRect:toScreen:"), frameRect, objc.ExtractPtr(screen))
	return rv
}

func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}

func (w_ Window) SetFrameDisplay(frameRect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setFrame:display:"), frameRect, flag)
}

func (w_ Window) SetFrameDisplayAnimate(frameRect foundation.Rect, displayFlag bool, animateFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}

func (w_ Window) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	rv := objc.CallMethod[foundation.TimeInterval](w_, objc.GetSelector("animationResizeTime:"), newFrame)
	return rv
}

func (w_ Window) PerformZoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("performZoom:"), objc.ExtractPtr(sender))
}

func (w_ Window) Zoom(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("zoom:"), objc.ExtractPtr(sender))
}

func (w_ Window) SetContentSize(size foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentSize:"), size)
}

func (w_ Window) OrderOut(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("orderOut:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderBack(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("orderBack:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("orderFront:"), objc.ExtractPtr(sender))
}

func (w_ Window) OrderFrontRegardless() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("orderFrontRegardless"))
}

func (w_ Window) OrderWindowRelativeTo(place WindowOrderingMode, otherWin int) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("orderWindow:relativeTo:"), place, otherWin)
}

func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](wc, objc.GetSelector("removeFrameUsingName:"), name)
}

func Window_RemoveFrameUsingName(name WindowFrameAutosaveName) {
	WindowClass.RemoveFrameUsingName(name)
}

func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setFrameUsingName:"), name)
	return rv
}

func (w_ Window) SetFrameUsingNameForce(name WindowFrameAutosaveName, force bool) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setFrameUsingName:force:"), name, force)
	return rv
}

func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("saveFrameUsingName:"), name)
}

func (w_ Window) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("setFrameAutosaveName:"), name)
	return rv
}

func (w_ Window) SetFrameFromString(string_ WindowPersistableFrameDescriptor) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setFrameFromString:"), string_)
}

func (w_ Window) MakeKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("makeKeyWindow"))
}

func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("makeKeyAndOrderFront:"), objc.ExtractPtr(sender))
}

func (w_ Window) BecomeKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("becomeKeyWindow"))
}

func (w_ Window) ResignKeyWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("resignKeyWindow"))
}

func (w_ Window) MakeMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("makeMainWindow"))
}

func (w_ Window) BecomeMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("becomeMainWindow"))
}

func (w_ Window) ResignMainWindow() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("resignMainWindow"))
}

func (w_ Window) ToggleToolbarShown(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("toggleToolbarShown:"), objc.ExtractPtr(sender))
}

func (w_ Window) RunToolbarCustomizationPalette(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("runToolbarCustomizationPalette:"), objc.ExtractPtr(sender))
}

func (w_ Window) AddChildWindowOrdered(childWin IWindow, place WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("addChildWindow:ordered:"), objc.ExtractPtr(childWin), place)
}

func (w_ Window) RemoveChildWindow(childWin IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("removeChildWindow:"), objc.ExtractPtr(childWin))
}

func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("enableKeyEquivalentForDefaultButtonCell"))
}

func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("disableKeyEquivalentForDefaultButtonCell"))
}

func (w_ Window) FieldEditorForObject(createFlag bool, object objc.IObject) Text {
	rv := objc.CallMethod[Text](w_, objc.GetSelector("fieldEditor:forObject:"), createFlag, objc.ExtractPtr(object))
	return rv
}

func (w_ Window) EndEditingFor(object objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("endEditingFor:"), objc.ExtractPtr(object))
}

func (w_ Window) EnableCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("enableCursorRects"))
}

func (w_ Window) DisableCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("disableCursorRects"))
}

func (w_ Window) DiscardCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("discardCursorRects"))
}

func (w_ Window) InvalidateCursorRectsForView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("invalidateCursorRectsForView:"), objc.ExtractPtr(view))
}

func (w_ Window) ResetCursorRects() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("resetCursorRects"))
}

func (wc _WindowClass) StandardWindowButtonForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	rv := objc.CallMethod[Button](wc, objc.GetSelector("standardWindowButton:forStyleMask:"), b, styleMask)
	return rv
}

func Window_StandardWindowButtonForStyleMask(b WindowButton, styleMask WindowStyleMask) Button {
	return WindowClass.StandardWindowButtonForStyleMask(b, styleMask)
}

func (w_ Window) StandardWindowButton(b WindowButton) Button {
	rv := objc.CallMethod[Button](w_, objc.GetSelector("standardWindowButton:"), b)
	return rv
}

func (w_ Window) AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("addTitlebarAccessoryViewController:"), objc.ExtractPtr(childViewController))
}

func (w_ Window) InsertTitlebarAccessoryViewControllerAtIndex(childViewController ITitlebarAccessoryViewController, index int) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("insertTitlebarAccessoryViewController:atIndex:"), objc.ExtractPtr(childViewController), index)
}

func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}

func (w_ Window) AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("addTabbedWindow:ordered:"), objc.ExtractPtr(window), ordered)
}

func (w_ Window) MergeAllWindows(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("mergeAllWindows:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectNextTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectNextTab:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectPreviousTab(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectPreviousTab:"), objc.ExtractPtr(sender))
}

func (w_ Window) MoveTabToNewWindow(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("moveTabToNewWindow:"), objc.ExtractPtr(sender))
}

func (w_ Window) ToggleTabBar(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("toggleTabBar:"), objc.ExtractPtr(sender))
}

func (w_ Window) ToggleTabOverview(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("toggleTabOverview:"), objc.ExtractPtr(sender))
}

func (w_ Window) NextEventMatchingMask(mask EventMask) Event {
	rv := objc.CallMethod[Event](w_, objc.GetSelector("nextEventMatchingMask:"), mask)
	return rv
}

func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.CallMethod[Event](w_, objc.GetSelector("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, objc.ExtractPtr(expiration), mode, deqFlag)
	return rv
}

func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("discardEventsMatchingMask:beforeEvent:"), mask, objc.ExtractPtr(lastEvent))
}

func (w_ Window) PostEventAtStart(event IEvent, flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("postEvent:atStart:"), objc.ExtractPtr(event), flag)
}

func (w_ Window) SendEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("sendEvent:"), objc.ExtractPtr(event))
}

func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("makeFirstResponder:"), objc.ExtractPtr(responder))
	return rv
}

func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectKeyViewPrecedingView:"), objc.ExtractPtr(view))
}

func (w_ Window) SelectKeyViewFollowingView(view IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectKeyViewFollowingView:"), objc.ExtractPtr(view))
}

func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectPreviousKeyView:"), objc.ExtractPtr(sender))
}

func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("selectNextKeyView:"), objc.ExtractPtr(sender))
}

func (w_ Window) RecalculateKeyViewLoop() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("recalculateKeyViewLoop"))
}

func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	rv := objc.CallMethod[int](wc, objc.GetSelector("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

func Window_WindowNumberAtPointBelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	return WindowClass.WindowNumberAtPointBelowWindowWithWindowNumber(point, windowNumber)
}

func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool)) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}

func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("performWindowDragWithEvent:"), objc.ExtractPtr(event))
}

func (w_ Window) DisableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("disableSnapshotRestoration"))
}

func (w_ Window) EnableSnapshotRestoration() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("enableSnapshotRestoration"))
}

func (w_ Window) Display() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("display"))
}

func (w_ Window) DisplayIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("displayIfNeeded"))
}

func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("disableScreenUpdatesUntilFlush"))
}

func (w_ Window) Update() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("update"))
}

func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("dragImage:at:offset:event:pasteboard:source:slideBack:"), objc.ExtractPtr(image), baseLocation, initialOffset, objc.ExtractPtr(event), objc.ExtractPtr(pboard), objc.ExtractPtr(sourceObj), slideFlag)
}

func (w_ Window) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("registerForDraggedTypes:"), newTypes)
}

func (w_ Window) UnregisterDraggedTypes() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("unregisterDraggedTypes"))
}

func (w_ Window) BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("backingAlignedRect:options:"), rect, options)
	return rv
}

func (w_ Window) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("convertRectFromBacking:"), rect)
	return rv
}

func (w_ Window) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("convertRectFromScreen:"), rect)
	return rv
}

func (w_ Window) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("convertPointFromBacking:"), point)
	return rv
}

func (w_ Window) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("convertPointFromScreen:"), point)
	return rv
}

func (w_ Window) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("convertRectToBacking:"), rect)
	return rv
}

func (w_ Window) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("convertRectToScreen:"), rect)
	return rv
}

func (w_ Window) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("convertPointToBacking:"), point)
	return rv
}

func (w_ Window) ConvertPointToScreen(point foundation.Point) foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("convertPointToScreen:"), point)
	return rv
}

func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitleWithRepresentedFilename:"), filename)
}

func (w_ Window) Center() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("center"))
}

func (w_ Window) PerformClose(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("performClose:"), objc.ExtractPtr(sender))
}

func (w_ Window) Close() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("close"))
}

func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("performMiniaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Miniaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("miniaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Deminiaturize(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("deminiaturize:"), objc.ExtractPtr(sender))
}

func (w_ Window) Print(sender objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("print:"), objc.ExtractPtr(sender))
}

func (w_ Window) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, objc.GetSelector("dataWithEPSInsideRect:"), rect)
	return rv
}

func (w_ Window) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.CallMethod[[]byte](w_, objc.GetSelector("dataWithPDFInsideRect:"), rect)
	return rv
}

func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("updateConstraintsIfNeeded"))
}

func (w_ Window) LayoutIfNeeded() {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("layoutIfNeeded"))
}

func (w_ Window) VisualizeConstraints(constraints []ILayoutConstraint) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("visualizeConstraints:"), constraints)
}

func (w_ Window) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	rv := objc.CallMethod[LayoutAttribute](w_, objc.GetSelector("anchorAttributeForOrientation:"), orientation)
	return rv
}

func (w_ Window) SetAnchorAttributeForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAnchorAttribute:forOrientation:"), attr, orientation)
}

func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setIsMiniaturized:"), flag)
}

func (w_ Window) SetIsVisible(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setIsVisible:"), flag)
}

func (w_ Window) SetIsZoomed(flag bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setIsZoomed:"), flag)
}

func (w_ Window) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("handleCloseScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (w_ Window) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("handlePrintScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (w_ Window) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("handleSaveScriptCommand:"), objc.ExtractPtr(command))
	return rv
}

func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("initWithWindowRef:"), windowRef)
	return rv
}

func (w_ Window) Delegate() WindowDelegateWrapper {
	rv := objc.CallMethod[WindowDelegateWrapper](w_, objc.GetSelector("delegate"))
	return rv
}

func (w_ Window) SetDelegate(value IWindowDelegate) {
	po := objc.WrapAsProtocol("NSWindowDelegate", value)
	objc.SetAssociatedObject(w_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDelegate:"), po)
}

func (w_ Window) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (w_ Window) ContentViewController() ViewController {
	rv := objc.CallMethod[ViewController](w_, objc.GetSelector("contentViewController"))
	return rv
}

func (w_ Window) SetContentViewController(value IViewController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentViewController:"), objc.ExtractPtr(value))
}

func (w_ Window) ContentView() View {
	rv := objc.CallMethod[View](w_, objc.GetSelector("contentView"))
	return rv
}

func (w_ Window) SetContentView(value IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentView:"), objc.ExtractPtr(value))
}

func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.CallMethod[WindowStyleMask](w_, objc.GetSelector("styleMask"))
	return rv
}

func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setStyleMask:"), value)
}

func (w_ Window) WorksWhenModal() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("worksWhenModal"))
	return rv
}

func (w_ Window) AlphaValue() float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("alphaValue"))
	return rv
}

func (w_ Window) SetAlphaValue(value float64) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAlphaValue:"), value)
}

func (w_ Window) BackgroundColor() Color {
	rv := objc.CallMethod[Color](w_, objc.GetSelector("backgroundColor"))
	return rv
}

func (w_ Window) SetBackgroundColor(value IColor) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setBackgroundColor:"), objc.ExtractPtr(value))
}

func (w_ Window) ColorSpace() ColorSpace {
	rv := objc.CallMethod[ColorSpace](w_, objc.GetSelector("colorSpace"))
	return rv
}

func (w_ Window) SetColorSpace(value IColorSpace) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setColorSpace:"), objc.ExtractPtr(value))
}

func (w_ Window) CanHide() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canHide"))
	return rv
}

func (w_ Window) SetCanHide(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setCanHide:"), value)
}

func (w_ Window) IsOnActiveSpace() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isOnActiveSpace"))
	return rv
}

func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hidesOnDeactivate"))
	return rv
}

func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setHidesOnDeactivate:"), value)
}

func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := objc.CallMethod[WindowCollectionBehavior](w_, objc.GetSelector("collectionBehavior"))
	return rv
}

func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setCollectionBehavior:"), value)
}

func (w_ Window) IsOpaque() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isOpaque"))
	return rv
}

func (w_ Window) SetOpaque(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setOpaque:"), value)
}

func (w_ Window) HasShadow() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hasShadow"))
	return rv
}

func (w_ Window) SetHasShadow(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setHasShadow:"), value)
}

func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("preventsApplicationTerminationWhenModal"))
	return rv
}

func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setPreventsApplicationTerminationWhenModal:"), value)
}

func (w_ Window) DepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](w_, objc.GetSelector("depthLimit"))
	return rv
}

func (w_ Window) SetDepthLimit(value WindowDepth) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDepthLimit:"), value)
}

func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hasDynamicDepthLimit"))
	return rv
}

func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := objc.CallMethod[WindowDepth](wc, objc.GetSelector("defaultDepthLimit"))
	return rv
}

func Window_DefaultDepthLimit() WindowDepth {
	return WindowClass.DefaultDepthLimit()
}

func (w_ Window) WindowNumber() int {
	rv := objc.CallMethod[int](w_, objc.GetSelector("windowNumber"))
	return rv
}

func (w_ Window) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.CallMethod[map[DeviceDescriptionKey]objc.Object](w_, objc.GetSelector("deviceDescription"))
	return rv
}

func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canBecomeVisibleWithoutLogin"))
	return rv
}

func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setCanBecomeVisibleWithoutLogin:"), value)
}

func (w_ Window) SharingType() WindowSharingType {
	rv := objc.CallMethod[WindowSharingType](w_, objc.GetSelector("sharingType"))
	return rv
}

func (w_ Window) SetSharingType(value WindowSharingType) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setSharingType:"), value)
}

func (w_ Window) BackingType() BackingStoreType {
	rv := objc.CallMethod[BackingStoreType](w_, objc.GetSelector("backingType"))
	return rv
}

func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setBackingType:"), value)
}

func (w_ Window) WindowController() WindowController {
	rv := objc.CallMethod[WindowController](w_, objc.GetSelector("windowController"))
	return rv
}

func (w_ Window) SetWindowController(value IWindowController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setWindowController:"), objc.ExtractPtr(value))
}

func (w_ Window) AttachedSheet() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("attachedSheet"))
	return rv
}

func (w_ Window) IsSheet() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isSheet"))
	return rv
}

func (w_ Window) SheetParent() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("sheetParent"))
	return rv
}

func (w_ Window) Sheets() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("sheets"))
	return rv
}

func (w_ Window) Frame() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("frame"))
	return rv
}

func (w_ Window) AspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("aspectRatio"))
	return rv
}

func (w_ Window) SetAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAspectRatio:"), value)
}

func (w_ Window) MinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("minSize"))
	return rv
}

func (w_ Window) SetMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMinSize:"), value)
}

func (w_ Window) MaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("maxSize"))
	return rv
}

func (w_ Window) SetMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMaxSize:"), value)
}

func (w_ Window) IsZoomed() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isZoomed"))
	return rv
}

func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.CallMethod[EventModifierFlags](w_, objc.GetSelector("resizeFlags"))
	return rv
}

func (w_ Window) ResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("resizeIncrements"))
	return rv
}

func (w_ Window) SetResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setResizeIncrements:"), value)
}

func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("preservesContentDuringLiveResize"))
	return rv
}

func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setPreservesContentDuringLiveResize:"), value)
}

func (w_ Window) InLiveResize() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("inLiveResize"))
	return rv
}

func (w_ Window) ContentAspectRatio() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("contentAspectRatio"))
	return rv
}

func (w_ Window) SetContentAspectRatio(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentAspectRatio:"), value)
}

func (w_ Window) ContentMinSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("contentMinSize"))
	return rv
}

func (w_ Window) SetContentMinSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentMinSize:"), value)
}

func (w_ Window) ContentMaxSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("contentMaxSize"))
	return rv
}

func (w_ Window) SetContentMaxSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentMaxSize:"), value)
}

func (w_ Window) ContentResizeIncrements() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("contentResizeIncrements"))
	return rv
}

func (w_ Window) SetContentResizeIncrements(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setContentResizeIncrements:"), value)
}

func (w_ Window) ContentLayoutGuide() objc.Object {
	rv := objc.CallMethod[objc.Object](w_, objc.GetSelector("contentLayoutGuide"))
	return rv
}

func (w_ Window) ContentLayoutRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](w_, objc.GetSelector("contentLayoutRect"))
	return rv
}

func (w_ Window) MaxFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("maxFullScreenContentSize"))
	return rv
}

func (w_ Window) SetMaxFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMaxFullScreenContentSize:"), value)
}

func (w_ Window) MinFullScreenContentSize() foundation.Size {
	rv := objc.CallMethod[foundation.Size](w_, objc.GetSelector("minFullScreenContentSize"))
	return rv
}

func (w_ Window) SetMinFullScreenContentSize(value foundation.Size) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMinFullScreenContentSize:"), value)
}

func (w_ Window) Level() WindowLevel {
	rv := objc.CallMethod[WindowLevel](w_, objc.GetSelector("level"))
	return rv
}

func (w_ Window) SetLevel(value WindowLevel) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setLevel:"), value)
}

func (w_ Window) IsVisible() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isVisible"))
	return rv
}

func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := objc.CallMethod[WindowOcclusionState](w_, objc.GetSelector("occlusionState"))
	return rv
}

func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.CallMethod[WindowFrameAutosaveName](w_, objc.GetSelector("frameAutosaveName"))
	return rv
}

func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	rv := objc.CallMethod[WindowPersistableFrameDescriptor](w_, objc.GetSelector("stringWithSavedFrame"))
	return rv
}

func (w_ Window) IsKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isKeyWindow"))
	return rv
}

func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canBecomeKeyWindow"))
	return rv
}

func (w_ Window) IsMainWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isMainWindow"))
	return rv
}

func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("canBecomeMainWindow"))
	return rv
}

func (w_ Window) Toolbar() Toolbar {
	rv := objc.CallMethod[Toolbar](w_, objc.GetSelector("toolbar"))
	return rv
}

func (w_ Window) SetToolbar(value IToolbar) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setToolbar:"), objc.ExtractPtr(value))
}

func (w_ Window) ChildWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("childWindows"))
	return rv
}

func (w_ Window) ParentWindow() Window {
	rv := objc.CallMethod[Window](w_, objc.GetSelector("parentWindow"))
	return rv
}

func (w_ Window) SetParentWindow(value IWindow) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setParentWindow:"), objc.ExtractPtr(value))
}

func (w_ Window) DefaultButtonCell() ButtonCell {
	rv := objc.CallMethod[ButtonCell](w_, objc.GetSelector("defaultButtonCell"))
	return rv
}

func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDefaultButtonCell:"), objc.ExtractPtr(value))
}

func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isExcludedFromWindowsMenu"))
	return rv
}

func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setExcludedFromWindowsMenu:"), value)
}

func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("areCursorRectsEnabled"))
	return rv
}

func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("titlebarAppearsTransparent"))
	return rv
}

func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitlebarAppearsTransparent:"), value)
}

func (w_ Window) ToolbarStyle() WindowToolbarStyle {
	rv := objc.CallMethod[WindowToolbarStyle](w_, objc.GetSelector("toolbarStyle"))
	return rv
}

func (w_ Window) SetToolbarStyle(value WindowToolbarStyle) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setToolbarStyle:"), value)
}

func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.CallMethod[TitlebarSeparatorStyle](w_, objc.GetSelector("titlebarSeparatorStyle"))
	return rv
}

func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitlebarSeparatorStyle:"), value)
}

func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.CallMethod[UserInterfaceLayoutDirection](w_, objc.GetSelector("windowTitlebarLayoutDirection"))
	return rv
}

func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := objc.CallMethod[[]TitlebarAccessoryViewController](w_, objc.GetSelector("titlebarAccessoryViewControllers"))
	return rv
}

func (w_ Window) SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitlebarAccessoryViewControllers:"), value)
}

func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.CallMethod[bool](wc, objc.GetSelector("allowsAutomaticWindowTabbing"))
	return rv
}

func Window_AllowsAutomaticWindowTabbing() bool {
	return WindowClass.AllowsAutomaticWindowTabbing()
}

func (wc _WindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.CallMethod[objc.Void](wc, objc.GetSelector("setAllowsAutomaticWindowTabbing:"), value)
}

func Window_SetAllowsAutomaticWindowTabbing(value bool) {
	WindowClass.SetAllowsAutomaticWindowTabbing(value)
}

func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.CallMethod[WindowUserTabbingPreference](wc, objc.GetSelector("userTabbingPreference"))
	return rv
}

func Window_UserTabbingPreference() WindowUserTabbingPreference {
	return WindowClass.UserTabbingPreference()
}

func (w_ Window) Tab() WindowTab {
	rv := objc.CallMethod[WindowTab](w_, objc.GetSelector("tab"))
	return rv
}

func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier {
	rv := objc.CallMethod[WindowTabbingIdentifier](w_, objc.GetSelector("tabbingIdentifier"))
	return rv
}

func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTabbingIdentifier:"), value)
}

func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := objc.CallMethod[WindowTabbingMode](w_, objc.GetSelector("tabbingMode"))
	return rv
}

func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTabbingMode:"), value)
}

func (w_ Window) TabbedWindows() []Window {
	rv := objc.CallMethod[[]Window](w_, objc.GetSelector("tabbedWindows"))
	return rv
}

func (w_ Window) TabGroup() WindowTabGroup {
	rv := objc.CallMethod[WindowTabGroup](w_, objc.GetSelector("tabGroup"))
	return rv
}

func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}

func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}

func (w_ Window) CurrentEvent() Event {
	rv := objc.CallMethod[Event](w_, objc.GetSelector("currentEvent"))
	return rv
}

func (w_ Window) InitialFirstResponder() View {
	rv := objc.CallMethod[View](w_, objc.GetSelector("initialFirstResponder"))
	return rv
}

func (w_ Window) SetInitialFirstResponder(value IView) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setInitialFirstResponder:"), objc.ExtractPtr(value))
}

func (w_ Window) FirstResponder() Responder {
	rv := objc.CallMethod[Responder](w_, objc.GetSelector("firstResponder"))
	return rv
}

func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := objc.CallMethod[SelectionDirection](w_, objc.GetSelector("keyViewSelectionDirection"))
	return rv
}

func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("autorecalculatesKeyViewLoop"))
	return rv
}

func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAutorecalculatesKeyViewLoop:"), value)
}

func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("acceptsMouseMovedEvents"))
	return rv
}

func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAcceptsMouseMovedEvents:"), value)
}

func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("ignoresMouseEvents"))
	return rv
}

func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setIgnoresMouseEvents:"), value)
}

func (w_ Window) MouseLocationOutsideOfEventStream() foundation.Point {
	rv := objc.CallMethod[foundation.Point](w_, objc.GetSelector("mouseLocationOutsideOfEventStream"))
	return rv
}

func (w_ Window) IsRestorable() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isRestorable"))
	return rv
}

func (w_ Window) SetRestorable(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setRestorable:"), value)
}

func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("viewsNeedDisplay"))
	return rv
}

func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setViewsNeedDisplay:"), value)
}

func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("allowsConcurrentViewDrawing"))
	return rv
}

func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAllowsConcurrentViewDrawing:"), value)
}

func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := objc.CallMethod[WindowAnimationBehavior](w_, objc.GetSelector("animationBehavior"))
	return rv
}

func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setAnimationBehavior:"), value)
}

func (w_ Window) IsDocumentEdited() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isDocumentEdited"))
	return rv
}

func (w_ Window) SetDocumentEdited(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDocumentEdited:"), value)
}

func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.CallMethod[float64](w_, objc.GetSelector("backingScaleFactor"))
	return rv
}

func (w_ Window) Title() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("title"))
	return rv
}

func (w_ Window) SetTitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitle:"), value)
}

func (w_ Window) Subtitle() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("subtitle"))
	return rv
}

func (w_ Window) SetSubtitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setSubtitle:"), value)
}

func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := objc.CallMethod[WindowTitleVisibility](w_, objc.GetSelector("titleVisibility"))
	return rv
}

func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setTitleVisibility:"), value)
}

func (w_ Window) RepresentedFilename() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("representedFilename"))
	return rv
}

func (w_ Window) SetRepresentedFilename(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setRepresentedFilename:"), value)
}

func (w_ Window) RepresentedURL() foundation.URL {
	rv := objc.CallMethod[foundation.URL](w_, objc.GetSelector("representedURL"))
	return rv
}

func (w_ Window) SetRepresentedURL(value foundation.IURL) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setRepresentedURL:"), objc.ExtractPtr(value))
}

func (w_ Window) Screen() Screen {
	rv := objc.CallMethod[Screen](w_, objc.GetSelector("screen"))
	return rv
}

func (w_ Window) DeepestScreen() Screen {
	rv := objc.CallMethod[Screen](w_, objc.GetSelector("deepestScreen"))
	return rv
}

func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("displaysWhenScreenProfileChanges"))
	return rv
}

func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setDisplaysWhenScreenProfileChanges:"), value)
}

func (w_ Window) IsMovableByWindowBackground() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isMovableByWindowBackground"))
	return rv
}

func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMovableByWindowBackground:"), value)
}

func (w_ Window) IsMovable() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isMovable"))
	return rv
}

func (w_ Window) SetMovable(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMovable:"), value)
}

func (w_ Window) IsReleasedWhenClosed() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isReleasedWhenClosed"))
	return rv
}

func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setReleasedWhenClosed:"), value)
}

func (w_ Window) IsMiniaturized() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isMiniaturized"))
	return rv
}

func (w_ Window) MiniwindowImage() Image {
	rv := objc.CallMethod[Image](w_, objc.GetSelector("miniwindowImage"))
	return rv
}

func (w_ Window) SetMiniwindowImage(value IImage) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMiniwindowImage:"), objc.ExtractPtr(value))
}

func (w_ Window) MiniwindowTitle() string {
	rv := objc.CallMethod[string](w_, objc.GetSelector("miniwindowTitle"))
	return rv
}

func (w_ Window) SetMiniwindowTitle(value string) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setMiniwindowTitle:"), value)
}

func (w_ Window) DockTile() DockTile {
	rv := objc.CallMethod[DockTile](w_, objc.GetSelector("dockTile"))
	return rv
}

func (w_ Window) HasCloseBox() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hasCloseBox"))
	return rv
}

func (w_ Window) HasTitleBar() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("hasTitleBar"))
	return rv
}

func (w_ Window) IsModalPanel() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isModalPanel"))
	return rv
}

func (w_ Window) IsFloatingPanel() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isFloatingPanel"))
	return rv
}

func (w_ Window) IsZoomable() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isZoomable"))
	return rv
}

func (w_ Window) IsResizable() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isResizable"))
	return rv
}

func (w_ Window) IsMiniaturizable() bool {
	rv := objc.CallMethod[bool](w_, objc.GetSelector("isMiniaturizable"))
	return rv
}

func (w_ Window) OrderedIndex() int {
	rv := objc.CallMethod[int](w_, objc.GetSelector("orderedIndex"))
	return rv
}

func (w_ Window) SetOrderedIndex(value int) {
	objc.CallMethod[objc.Void](w_, objc.GetSelector("setOrderedIndex:"), value)
}

func (w_ Window) WindowRef() unsafe.Pointer {
	rv := objc.CallMethod[unsafe.Pointer](w_, objc.GetSelector("windowRef"))
	return rv
}
