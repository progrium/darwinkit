// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Window] class.
var WindowClass = _WindowClass{objc.GetClass("NSWindow")}

type _WindowClass struct {
	objc.Class
}

// An interface definition for the [Window] class.
type IWindow interface {
	IResponder
	AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode)
	SaveFrameUsingName(name WindowFrameAutosaveName)
	MakeKeyWindow()
	SendEvent(event IEvent)
	DisableKeyEquivalentForDefaultButtonCell()
	DiscardCursorRects()
	ConvertPointToBacking(point foundation.Point) foundation.Point
	SetFrameOrigin(point foundation.Point)
	EnableCursorRects()
	OrderFrontRegardless()
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	SetContentSize(size foundation.Size)
	AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool
	SetFrameFromString(string_ WindowPersistableFrameDescriptor)
	SelectKeyViewPrecedingView(view IView)
	DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool)
	MoveTabToNewWindow(sender objc.IObject) objc.Object
	OrderOut(sender objc.IObject)
	RecalculateKeyViewLoop()
	VisualizeConstraints(constraints []ILayoutConstraint)
	SetContentBorderThicknessForEdge(thickness float64, edge foundation.RectEdge)
	ConvertRectToScreen(rect foundation.Rect) foundation.Rect
	ResetCursorRects()
	AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute
	UpdateConstraintsIfNeeded()
	HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object
	Close()
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	ToggleTabBar(sender objc.IObject) objc.Object
	HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object
	RemoveChildWindow(childWin IWindow)
	MakeMainWindow()
	RunToolbarCustomizationPalette(sender objc.IObject)
	ResignMainWindow()
	BeginSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	DisableScreenUpdatesUntilFlush()
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	SetFrameDisplay(frameRect foundation.Rect, flag bool)
	AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval
	SelectKeyViewFollowingView(view IView)
	RegisterForDraggedTypes(newTypes []PasteboardType)
	SetAnchorAttributeForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation)
	ContentBorderThicknessForEdge(edge foundation.RectEdge) float64
	OrderBack(sender objc.IObject)
	TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool))
	InitWithWindowRef(windowRef unsafe.Pointer) Window
	InvalidateShadow()
	HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object
	Display()
	PerformMiniaturize(sender objc.IObject)
	DisableCursorRects()
	SetDynamicDepthLimit(flag bool)
	Miniaturize(sender objc.IObject)
	SelectNextTab(sender objc.IObject) objc.Object
	AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController)
	MakeFirstResponder(responder IResponder) bool
	SetIsZoomed(flag bool)
	BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse))
	ToggleFullScreen(sender objc.IObject)
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	OrderWindowRelativeTo(place WindowOrderingMode, otherWin int)
	Print(sender objc.IObject)
	FieldEditorForObject(createFlag bool, object objc.IObject) Text
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.RectEdge)
	SelectPreviousTab(sender objc.IObject) objc.Object
	SelectNextKeyView(sender objc.IObject)
	SetTitleWithRepresentedFilename(filename string)
	MakeKeyAndOrderFront(sender objc.IObject)
	SetFrameAutosaveName(name WindowFrameAutosaveName) bool
	SelectPreviousKeyView(sender objc.IObject)
	PostEventAtStart(event IEvent, flag bool)
	ToggleTabOverview(sender objc.IObject) objc.Object
	OrderFront(sender objc.IObject)
	ResignKeyWindow()
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	EndEditingFor(object objc.IObject)
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	BecomeKeyWindow()
	EnableSnapshotRestoration()
	DisableSnapshotRestoration()
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event
	Deminiaturize(sender objc.IObject)
	UnregisterDraggedTypes()
	SetFrameUsingName(name WindowFrameAutosaveName) bool
	InvalidateCursorRectsForView(view IView)
	ConvertPointToScreen(point foundation.Point) foundation.Point
	AddChildWindowOrdered(childWin IWindow, place WindowOrderingMode)
	PerformClose(sender objc.IObject)
	ConvertRectFromScreen(rect foundation.Rect) foundation.Rect
	SetIsVisible(flag bool)
	PerformWindowDragWithEvent(event IEvent)
	StandardWindowButton(b WindowButton) Button
	Zoom(sender objc.IObject)
	MergeAllWindows(sender objc.IObject) objc.Object
	Update()
	ConvertPointFromScreen(point foundation.Point) foundation.Point
	EnableKeyEquivalentForDefaultButtonCell()
	ToggleToolbarShown(sender objc.IObject)
	SetIsMiniaturized(flag bool)
	BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	ConstrainFrameRectToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect
	LayoutIfNeeded()
	EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse)
	BecomeMainWindow()
	InsertTitlebarAccessoryViewControllerAtIndex(childViewController ITitlebarAccessoryViewController, index int)
	PerformZoom(sender objc.IObject)
	SetFrameTopLeftPoint(point foundation.Point)
	DisplayIfNeeded()
	Center()
	IsRestorable() bool
	SetRestorable(value bool)
	MaxSize() foundation.Size
	SetMaxSize(value foundation.Size)
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	IsVisible() bool
	ParentWindow() Window
	SetParentWindow(value IWindow)
	MinFullScreenContentSize() foundation.Size
	SetMinFullScreenContentSize(value foundation.Size)
	Sheets() []Window
	IsExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)
	IsSheet() bool
	IsReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)
	SharingType() WindowSharingType
	SetSharingType(value WindowSharingType)
	HasShadow() bool
	SetHasShadow(value bool)
	StyleMask() WindowStyleMask
	SetStyleMask(value WindowStyleMask)
	Tab() WindowTab
	OrderedIndex() int
	SetOrderedIndex(value int)
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	Level() WindowLevel
	SetLevel(value WindowLevel)
	AnimationBehavior() WindowAnimationBehavior
	SetAnimationBehavior(value WindowAnimationBehavior)
	Screen() Screen
	IsModalPanel() bool
	MouseLocationOutsideOfEventStream() foundation.Point
	IsKeyWindow() bool
	ContentView() View
	SetContentView(value IView)
	TitleVisibility() WindowTitleVisibility
	SetTitleVisibility(value WindowTitleVisibility)
	ToolbarStyle() WindowToolbarStyle
	SetToolbarStyle(value WindowToolbarStyle)
	CurrentEvent() Event
	MiniwindowTitle() string
	SetMiniwindowTitle(value string)
	IsMainWindow() bool
	ChildWindows() []Window
	DefaultButtonCell() ButtonCell
	SetDefaultButtonCell(value IButtonCell)
	FirstResponder() Responder
	TabbedWindows() []Window
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	ContentAspectRatio() foundation.Size
	SetContentAspectRatio(value foundation.Size)
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)
	RestorationClass() objc.Class
	SetRestorationClass(value objc.IClass)
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	Subtitle() string
	SetSubtitle(value string)
	WorksWhenModal() bool
	IsMiniaturized() bool
	AlphaValue() float64
	SetAlphaValue(value float64)
	DeepestScreen() Screen
	ContentLayoutRect() foundation.Rect
	HasTitleBar() bool
	IsMiniaturizable() bool
	MinSize() foundation.Size
	SetMinSize(value foundation.Size)
	BackingType() BackingStoreType
	SetBackingType(value BackingStoreType)
	MiniwindowImage() Image
	SetMiniwindowImage(value IImage)
	Toolbar() Toolbar
	SetToolbar(value IToolbar)
	ColorSpace() ColorSpace
	SetColorSpace(value IColorSpace)
	DepthLimit() WindowDepth
	SetDepthLimit(value WindowDepth)
	Delegate() WindowDelegateWrapper
	SetDelegate(value PWindowDelegate)
	SetDelegateObject(valueObject objc.IObject)
	ResizeFlags() EventModifierFlags
	IsFloatingPanel() bool
	InLiveResize() bool
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController)
	ResizeIncrements() foundation.Size
	SetResizeIncrements(value foundation.Size)
	WindowRef() unsafe.Pointer
	CanHide() bool
	SetCanHide(value bool)
	IsOpaque() bool
	SetOpaque(value bool)
	RepresentedURL() foundation.URL
	SetRepresentedURL(value foundation.IURL)
	InitialFirstResponder() View
	SetInitialFirstResponder(value IView)
	IsZoomable() bool
	OcclusionState() WindowOcclusionState
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	AreCursorRectsEnabled() bool
	ContentViewController() ViewController
	SetContentViewController(value IViewController)
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
	ContentMinSize() foundation.Size
	SetContentMinSize(value foundation.Size)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	IsDocumentEdited() bool
	SetDocumentEdited(value bool)
	CollectionBehavior() WindowCollectionBehavior
	SetCollectionBehavior(value WindowCollectionBehavior)
	TabbingIdentifier() WindowTabbingIdentifier
	SetTabbingIdentifier(value WindowTabbingIdentifier)
	WindowController() WindowController
	SetWindowController(value IWindowController)
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	BackingScaleFactor() float64
	TabGroup() WindowTabGroup
	WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection
	TabbingMode() WindowTabbingMode
	SetTabbingMode(value WindowTabbingMode)
	RepresentedFilename() string
	SetRepresentedFilename(value string)
	AspectRatio() foundation.Size
	SetAspectRatio(value foundation.Size)
	WindowNumber() int
	Frame() foundation.Rect
	CanBecomeMainWindow() bool
	Title() string
	SetTitle(value string)
	HasCloseBox() bool
	HasDynamicDepthLimit() bool
	StringWithSavedFrame() WindowPersistableFrameDescriptor
	IsResizable() bool
	DockTile() DockTile
	IsOnActiveSpace() bool
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	AppearanceSource() objc.Object
	SetAppearanceSource(value objc.IObject)
	AttachedSheet() Window
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	ContentLayoutGuide() objc.Object
	KeyViewSelectionDirection() SelectionDirection
	IsMovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	IsMovable() bool
	SetMovable(value bool)
	IsZoomed() bool
	ContentMaxSize() foundation.Size
	SetContentMaxSize(value foundation.Size)
	CanBecomeKeyWindow() bool
	ContentResizeIncrements() foundation.Size
	SetContentResizeIncrements(value foundation.Size)
	FrameAutosaveName() WindowFrameAutosaveName
	MaxFullScreenContentSize() foundation.Size
	SetMaxFullScreenContentSize(value foundation.Size)
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)
	SheetParent() Window
}

// A window that an app displays on the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow?language=objc
type Window struct {
	Responder
}

func WindowFrom(ptr unsafe.Pointer) Window {
	return Window{
		Responder: ResponderFrom(ptr),
	}
}

func (wc _WindowClass) WindowWithContentViewController(contentViewController IViewController) Window {
	rv := objc.Call[Window](wc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func Window_WindowWithContentViewController(contentViewController IViewController) Window {
	return WindowClass.WindowWithContentViewController(contentViewController)
}

func (w_ Window) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	rv := objc.Call[Window](w_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	instance := WindowClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

func (wc _WindowClass) Alloc() Window {
	rv := objc.Call[Window](wc, objc.Sel("alloc"))
	return rv
}

func Window_Alloc() Window {
	return WindowClass.Alloc()
}

func (wc _WindowClass) New() Window {
	rv := objc.Call[Window](wc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewWindow() Window {
	return WindowClass.New()
}

func (w_ Window) Init() Window {
	rv := objc.Call[Window](w_, objc.Sel("init"))
	return rv
}

// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1855947-addtabbedwindow?language=objc
func (w_ Window) AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode) {
	objc.Call[objc.Void](w_, objc.Sel("addTabbedWindow:ordered:"), objc.Ptr(window), ordered)
}

// Returns the frame rectangle used by a window with a given content rectangle and window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419372-framerectforcontentrect?language=objc
func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.Call[foundation.Rect](wc, objc.Sel("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}

// Returns the frame rectangle used by a window with a given content rectangle and window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419372-framerectforcontentrect?language=objc
func Window_FrameRectForContentRectStyleMask(cRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	return WindowClass.FrameRectForContentRectStyleMask(cRect, style)
}

// Saves the window’s frame rectangle in the user defaults system under a given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419290-saveframeusingname?language=objc
func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName) {
	objc.Call[objc.Void](w_, objc.Sel("saveFrameUsingName:"), name)
}

// Makes the window the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419368-makekeywindow?language=objc
func (w_ Window) MakeKeyWindow() {
	objc.Call[objc.Void](w_, objc.Sel("makeKeyWindow"))
}

// This action method dispatches mouse and keyboard events the global application object sends to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419228-sendevent?language=objc
func (w_ Window) SendEvent(event IEvent) {
	objc.Call[objc.Void](w_, objc.Sel("sendEvent:"), objc.Ptr(event))
}

// Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419242-disablekeyequivalentfordefaultbu?language=objc
func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.Call[objc.Void](w_, objc.Sel("disableKeyEquivalentForDefaultButtonCell"))
}

// Returns the window numbers for all visible windows satisfying the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419678-windownumberswithoptions?language=objc
func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := objc.Call[[]foundation.Number](wc, objc.Sel("windowNumbersWithOptions:"), options)
	return rv
}

// Returns the window numbers for all visible windows satisfying the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419678-windownumberswithoptions?language=objc
func Window_WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	return WindowClass.WindowNumbersWithOptions(options)
}

// Invalidates all cursor rectangles in the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419269-discardcursorrects?language=objc
func (w_ Window) DiscardCursorRects() {
	objc.Call[objc.Void](w_, objc.Sel("discardCursorRects"))
}

// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2967181-convertpointtobacking?language=objc
func (w_ Window) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("convertPointToBacking:"), point)
	return rv
}

// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419690-setframeorigin?language=objc
func (w_ Window) SetFrameOrigin(point foundation.Point) {
	objc.Call[objc.Void](w_, objc.Sel("setFrameOrigin:"), point)
}

// Reenables cursor rectangle management within the window after a disableCursorRects message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419202-enablecursorrects?language=objc
func (w_ Window) EnableCursorRects() {
	objc.Call[objc.Void](w_, objc.Sel("enableCursorRects"))
}

// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419444-orderfrontregardless?language=objc
func (w_ Window) OrderFrontRegardless() {
	objc.Call[objc.Void](w_, objc.Sel("orderFrontRegardless"))
}

// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419273-convertrectfrombacking?language=objc
func (w_ Window) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}

// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419100-setcontentsize?language=objc
func (w_ Window) SetContentSize(size foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setContentSize:"), size)
}

// Indicates whether the window calculates the thickness of a given border automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419356-autorecalculatescontentborderthi?language=objc
func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge foundation.RectEdge) bool {
	rv := objc.Call[bool](w_, objc.Sel("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}

// Sets the window’s frame rectangle from a given string representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419759-setframefromstring?language=objc
func (w_ Window) SetFrameFromString(string_ WindowPersistableFrameDescriptor) {
	objc.Call[objc.Void](w_, objc.Sel("setFrameFromString:"), string_)
}

// Gives key view status to the view that precedes the given view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419757-selectkeyviewprecedingview?language=objc
func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	objc.Call[objc.Void](w_, objc.Sel("selectKeyViewPrecedingView:"), objc.Ptr(view))
}

// Begins a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419224-dragimage?language=objc
func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, baseLocation foundation.Point, initialOffset foundation.Size, event IEvent, pboard IPasteboard, sourceObj objc.IObject, slideFlag bool) {
	objc.Call[objc.Void](w_, objc.Sel("dragImage:at:offset:event:pasteboard:source:slideBack:"), objc.Ptr(image), baseLocation, initialOffset, objc.Ptr(event), objc.Ptr(pboard), sourceObj, slideFlag)
}

// Moves the tab to a new containing window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644410-movetabtonewwindow?language=objc
func (w_ Window) MoveTabToNewWindow(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("moveTabToNewWindow:"), sender)
	return rv
}

// Removes the window from the screen list, which hides the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419660-orderout?language=objc
func (w_ Window) OrderOut(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("orderOut:"), sender)
}

// Marks the key view loop as “dirty” and in need of recalculation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419350-recalculatekeyviewloop?language=objc
func (w_ Window) RecalculateKeyViewLoop() {
	objc.Call[objc.Void](w_, objc.Sel("recalculateKeyViewLoop"))
}

// Displays a visual representation of the supplied constraints in the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526997-visualizeconstraints?language=objc
func (w_ Window) VisualizeConstraints(constraints []ILayoutConstraint) {
	objc.Call[objc.Void](w_, objc.Sel("visualizeConstraints:"), constraints)
}

// Specifies the thickness of a given border of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419541-setcontentborderthickness?language=objc
func (w_ Window) SetContentBorderThicknessForEdge(thickness float64, edge foundation.RectEdge) {
	objc.Call[objc.Void](w_, objc.Sel("setContentBorderThickness:forEdge:"), thickness, edge)
}

// Converts a rectangle to the screen coordinate system from the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419286-convertrecttoscreen?language=objc
func (w_ Window) ConvertRectToScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("convertRectToScreen:"), rect)
	return rv
}

// Clears the window’s cursor rectangles and the cursor rectangles of the NSView objects in its view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419464-resetcursorrects?language=objc
func (w_ Window) ResetCursorRects() {
	objc.Call[objc.Void](w_, objc.Sel("resetCursorRects"))
}

// Returns the part of the window that stays stationary during constraint-based layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526957-anchorattributefororientation?language=objc
func (w_ Window) AnchorAttributeForOrientation(orientation LayoutConstraintOrientation) LayoutAttribute {
	rv := objc.Call[LayoutAttribute](w_, objc.Sel("anchorAttributeForOrientation:"), orientation)
	return rv
}

// Updates the constraints based on changes to views in the window since the last layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526915-updateconstraintsifneeded?language=objc
func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.Call[objc.Void](w_, objc.Sel("updateConstraintsIfNeeded"))
}

// Handles the AppleScript command to save the window (and its associated document, if any). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449591-handlesavescriptcommand?language=objc
func (w_ Window) HandleSaveScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("handleSaveScriptCommand:"), objc.Ptr(command))
	return rv
}

// Removes the window from the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419662-close?language=objc
func (w_ Window) Close() {
	objc.Call[objc.Void](w_, objc.Sel("close"))
}

// Returns EPS data that draws the region of the window within a given rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419128-datawithepsinsiderect?language=objc
func (w_ Window) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.Call[[]byte](w_, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}

// Shows or hides the tab bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644517-toggletabbar?language=objc
func (w_ Window) ToggleTabBar(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("toggleTabBar:"), sender)
	return rv
}

// Handles the AppleScript command to close the window (and its associated document, if any). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449581-handleclosescriptcommand?language=objc
func (w_ Window) HandleCloseScriptCommand(command foundation.ICloseCommand) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("handleCloseScriptCommand:"), objc.Ptr(command))
	return rv
}

// Detaches a given child window from the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419063-removechildwindow?language=objc
func (w_ Window) RemoveChildWindow(childWin IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("removeChildWindow:"), objc.Ptr(childWin))
}

// Makes the window the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419271-makemainwindow?language=objc
func (w_ Window) MakeMainWindow() {
	objc.Call[objc.Void](w_, objc.Sel("makeMainWindow"))
}

// Presents the toolbar customization user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419284-runtoolbarcustomizationpalette?language=objc
func (w_ Window) RunToolbarCustomizationPalette(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("runToolbarCustomizationPalette:"), sender)
}

// Resigns the window’s main window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419212-resignmainwindow?language=objc
func (w_ Window) ResignMainWindow() {
	objc.Call[objc.Void](w_, objc.Sel("resignMainWindow"))
}

// Starts a document-modal session and presents—or queues for presentation—a sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419653-beginsheet?language=objc
func (w_ Window) BeginSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.Call[objc.Void](w_, objc.Sel("beginSheet:completionHandler:"), objc.Ptr(sheetWindow), handler)
}

// Positions the window’s top-left to a given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419392-cascadetopleftfrompoint?language=objc
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}

// Forwards the message to the global application object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419676-discardeventsmatchingmask?language=objc
func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.Call[objc.Void](w_, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, objc.Ptr(lastEvent))
}

// Disables the window’s screen updates until the window is flushed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419483-disablescreenupdatesuntilflush?language=objc
func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.Call[objc.Void](w_, objc.Sel("disableScreenUpdatesUntilFlush"))
}

// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419260-convertrecttobacking?language=objc
func (w_ Window) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("convertRectToBacking:"), rect)
	return rv
}

// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419753-setframe?language=objc
func (w_ Window) SetFrameDisplay(frameRect foundation.Rect, flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setFrame:display:"), frameRect, flag)
}

// Specifies the duration of a smooth frame-size change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419655-animationresizetime?language=objc
func (w_ Window) AnimationResizeTime(newFrame foundation.Rect) foundation.TimeInterval {
	rv := objc.Call[foundation.TimeInterval](w_, objc.Sel("animationResizeTime:"), newFrame)
	return rv
}

// Gives key view status to the view that follows the given view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419633-selectkeyviewfollowingview?language=objc
func (w_ Window) SelectKeyViewFollowingView(view IView) {
	objc.Call[objc.Void](w_, objc.Sel("selectKeyViewFollowingView:"), objc.Ptr(view))
}

// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419140-registerfordraggedtypes?language=objc
func (w_ Window) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.Call[objc.Void](w_, objc.Sel("registerForDraggedTypes:"), newTypes)
}

// Sets the part of the window that stays stationary during constraint-based layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526985-setanchorattribute?language=objc
func (w_ Window) SetAnchorAttributeForOrientation(attr LayoutAttribute, orientation LayoutConstraintOrientation) {
	objc.Call[objc.Void](w_, objc.Sel("setAnchorAttribute:forOrientation:"), attr, orientation)
}

// Indicates the thickness of a given border of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419775-contentborderthicknessforedge?language=objc
func (w_ Window) ContentBorderThicknessForEdge(edge foundation.RectEdge) float64 {
	rv := objc.Call[float64](w_, objc.Sel("contentBorderThicknessForEdge:"), edge)
	return rv
}

// Moves the window to the back of its level in the screen list, without changing either the key window or the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419204-orderback?language=objc
func (w_ Window) OrderBack(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("orderBack:"), sender)
}

// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419727-trackeventsmatchingmask?language=objc
func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout foundation.TimeInterval, mode foundation.RunLoopMode, trackingHandler func(event Event, stop *bool)) {
	objc.Call[objc.Void](w_, objc.Sel("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}

// Returns a Cocoa window created from a Carbon window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419446-initwithwindowref?language=objc
func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) Window {
	rv := objc.Call[Window](w_, objc.Sel("initWithWindowRef:"), windowRef)
	return rv
}

// Invalidates the window shadow so that it is recomputed based on the current window shape. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419529-invalidateshadow?language=objc
func (w_ Window) InvalidateShadow() {
	objc.Call[objc.Void](w_, objc.Sel("invalidateShadow"))
}

// Handles the AppleScript command to print the contents of the window (or its associated document, if any). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449585-handleprintscriptcommand?language=objc
func (w_ Window) HandlePrintScriptCommand(command foundation.IScriptCommand) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("handlePrintScriptCommand:"), objc.Ptr(command))
	return rv
}

// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419358-display?language=objc
func (w_ Window) Display() {
	objc.Call[objc.Void](w_, objc.Sel("display"))
}

// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419749-performminiaturize?language=objc
func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("performMiniaturize:"), sender)
}

// Disables all cursor rectangle management within the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419639-disablecursorrects?language=objc
func (w_ Window) DisableCursorRects() {
	objc.Call[objc.Void](w_, objc.Sel("disableCursorRects"))
}

// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419473-setdynamicdepthlimit?language=objc
func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setDynamicDepthLimit:"), flag)
}

// Returns the content rectangle used by a window with a given frame rectangle and window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419586-contentrectforframerect?language=objc
func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	rv := objc.Call[foundation.Rect](wc, objc.Sel("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}

// Returns the content rectangle used by a window with a given frame rectangle and window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419586-contentrectforframerect?language=objc
func Window_ContentRectForFrameRectStyleMask(fRect foundation.Rect, style WindowStyleMask) foundation.Rect {
	return WindowClass.ContentRectForFrameRectStyleMask(fRect, style)
}

// Removes the window from the screen list and displays the minimized window in the Dock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419426-miniaturize?language=objc
func (w_ Window) Miniaturize(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("miniaturize:"), sender)
}

// Selects the next tab in the tab group in the trailing direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644693-selectnexttab?language=objc
func (w_ Window) SelectNextTab(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("selectNextTab:"), sender)
	return rv
}

// Adds the specified title bar accessory view controller to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419382-addtitlebaraccessoryviewcontroll?language=objc
func (w_ Window) AddTitlebarAccessoryViewController(childViewController ITitlebarAccessoryViewController) {
	objc.Call[objc.Void](w_, objc.Sel("addTitlebarAccessoryViewController:"), objc.Ptr(childViewController))
}

// Attempts to make a given responder the first responder for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419366-makefirstresponder?language=objc
func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := objc.Call[bool](w_, objc.Sel("makeFirstResponder:"), objc.Ptr(responder))
	return rv
}

// Sets the window’s zoomed state to the value you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449589-setiszoomed?language=objc
func (w_ Window) SetIsZoomed(flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setIsZoomed:"), flag)
}

// Starts a document-modal session and presents the specified critical sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419198-begincriticalsheet?language=objc
func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler func(returnCode ModalResponse)) {
	objc.Call[objc.Void](w_, objc.Sel("beginCriticalSheet:completionHandler:"), objc.Ptr(sheetWindow), handler)
}

// Takes the window into or out of fullscreen mode, [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419527-togglefullscreen?language=objc
func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("toggleFullScreen:"), sender)
}

// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2967179-convertpointfrombacking?language=objc
func (w_ Window) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("convertPointFromBacking:"), point)
	return rv
}

// Repositions the window’s window device in the window server’s screen list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419672-orderwindow?language=objc
func (w_ Window) OrderWindowRelativeTo(place WindowOrderingMode, otherWin int) {
	objc.Call[objc.Void](w_, objc.Sel("orderWindow:relativeTo:"), place, otherWin)
}

// Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419767-print?language=objc
func (w_ Window) Print(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("print:"), sender)
}

// Returns the window’s field editor, creating it if requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419647-fieldeditor?language=objc
func (w_ Window) FieldEditorForObject(createFlag bool, object objc.IObject) Text {
	rv := objc.Call[Text](w_, objc.Sel("fieldEditor:forObject:"), createFlag, object)
	return rv
}

// Removes the view controller at the specified index from the window’s array of title bar accessory view controllers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419643-removetitlebaraccessoryviewcontr?language=objc
func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.Call[objc.Void](w_, objc.Sel("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}

// Specifies whether the window calculates the thickness of a given border automatically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419218-setautorecalculatescontentborder?language=objc
func (w_ Window) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge foundation.RectEdge) {
	objc.Call[objc.Void](w_, objc.Sel("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}

// Selects the previous tab in the tab group in the leading direction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644555-selectprevioustab?language=objc
func (w_ Window) SelectPreviousTab(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("selectPreviousTab:"), sender)
	return rv
}

// Searches for a candidate next key view and, if it finds one, tries to make it the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419715-selectnextkeyview?language=objc
func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("selectNextKeyView:"), sender)
}

// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419192-settitlewithrepresentedfilename?language=objc
func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	objc.Call[objc.Void](w_, objc.Sel("setTitleWithRepresentedFilename:"), filename)
}

// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419208-makekeyandorderfront?language=objc
func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("makeKeyAndOrderFront:"), sender)
}

// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419210-windownumberatpoint?language=objc
func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	rv := objc.Call[int](wc, objc.Sel("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419210-windownumberatpoint?language=objc
func Window_WindowNumberAtPointBelowWindowWithWindowNumber(point foundation.Point, windowNumber int) int {
	return WindowClass.WindowNumberAtPointBelowWindowWithWindowNumber(point, windowNumber)
}

// Sets the name AppKit uses to automatically save the window’s frame rectangle data in the defaults system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419509-setframeautosavename?language=objc
func (w_ Window) SetFrameAutosaveName(name WindowFrameAutosaveName) bool {
	rv := objc.Call[bool](w_, objc.Sel("setFrameAutosaveName:"), name)
	return rv
}

// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419110-selectpreviouskeyview?language=objc
func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("selectPreviousKeyView:"), sender)
}

// Forwards the message to the global application object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419376-postevent?language=objc
func (w_ Window) PostEventAtStart(event IEvent, flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("postEvent:atStart:"), objc.Ptr(event), flag)
}

// Shows or hides the tab overview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2870175-toggletaboverview?language=objc
func (w_ Window) ToggleTabOverview(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("toggleTabOverview:"), sender)
	return rv
}

// Moves the window to the front of its level in the screen list, without changing either the key window or the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419495-orderfront?language=objc
func (w_ Window) OrderFront(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("orderFront:"), sender)
}

// Resigns the window’s key window status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419047-resignkeywindow?language=objc
func (w_ Window) ResignKeyWindow() {
	objc.Call[objc.Void](w_, objc.Sel("resignKeyWindow"))
}

// Returns PDF data that draws the region of the window within a given rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419418-datawithpdfinsiderect?language=objc
func (w_ Window) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.Call[[]byte](w_, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}

// Forces the field editor to give up its first responder status and prepares it for its next assignment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419469-endeditingfor?language=objc
func (w_ Window) EndEditingFor(object objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("endEditingFor:"), object)
}

// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2138278-canrepresentdisplaygamut?language=objc
func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.Call[bool](w_, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

// Informs the window that it has become the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419338-becomekeywindow?language=objc
func (w_ Window) BecomeKeyWindow() {
	objc.Call[objc.Void](w_, objc.Sel("becomeKeyWindow"))
}

// Enables snapshot restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1525288-enablesnapshotrestoration?language=objc
func (w_ Window) EnableSnapshotRestoration() {
	objc.Call[objc.Void](w_, objc.Sel("enableSnapshotRestoration"))
}

// Disables snapshot restoration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526239-disablesnapshotrestoration?language=objc
func (w_ Window) DisableSnapshotRestoration() {
	objc.Call[objc.Void](w_, objc.Sel("disableSnapshotRestoration"))
}

// Forwards the message to the global application object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419721-nexteventmatchingmask?language=objc
func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration foundation.IDate, mode foundation.RunLoopMode, deqFlag bool) Event {
	rv := objc.Call[Event](w_, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, objc.Ptr(expiration), mode, deqFlag)
	return rv
}

// De-minimizes the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419334-deminiaturize?language=objc
func (w_ Window) Deminiaturize(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("deminiaturize:"), sender)
}

// Unregisters the window as a possible destination for dragging operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419456-unregisterdraggedtypes?language=objc
func (w_ Window) UnregisterDraggedTypes() {
	objc.Call[objc.Void](w_, objc.Sel("unregisterDraggedTypes"))
}

// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419723-setframeusingname?language=objc
func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName) bool {
	rv := objc.Call[bool](w_, objc.Sel("setFrameUsingName:"), name)
	return rv
}

// Removes the frame data stored under a given name from the application’s user defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419313-removeframeusingname?language=objc
func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName) {
	objc.Call[objc.Void](wc, objc.Sel("removeFrameUsingName:"), name)
}

// Removes the frame data stored under a given name from the application’s user defaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419313-removeframeusingname?language=objc
func Window_RemoveFrameUsingName(name WindowFrameAutosaveName) {
	WindowClass.RemoveFrameUsingName(name)
}

// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419601-invalidatecursorrectsforview?language=objc
func (w_ Window) InvalidateCursorRectsForView(view IView) {
	objc.Call[objc.Void](w_, objc.Sel("invalidateCursorRectsForView:"), objc.Ptr(view))
}

// Converts a point to the screen coordinate system from the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2967182-convertpointtoscreen?language=objc
func (w_ Window) ConvertPointToScreen(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("convertPointToScreen:"), point)
	return rv
}

// Adds a given window as a child window of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419152-addchildwindow?language=objc
func (w_ Window) AddChildWindowOrdered(childWin IWindow, place WindowOrderingMode) {
	objc.Call[objc.Void](w_, objc.Sel("addChildWindow:ordered:"), objc.Ptr(childWin), place)
}

// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419288-performclose?language=objc
func (w_ Window) PerformClose(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("performClose:"), sender)
}

// Converts a rectangle from the screen coordinate system to the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419603-convertrectfromscreen?language=objc
func (w_ Window) ConvertRectFromScreen(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("convertRectFromScreen:"), rect)
	return rv
}

// Sets the window’s visible state to the value you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449570-setisvisible?language=objc
func (w_ Window) SetIsVisible(flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setIsVisible:"), flag)
}

// Starts a window drag based on the specified mouse-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419386-performwindowdragwithevent?language=objc
func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	objc.Call[objc.Void](w_, objc.Sel("performWindowDragWithEvent:"), objc.Ptr(event))
}

// Returns the window button of a given window button kind in the window’s view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419491-standardwindowbutton?language=objc
func (w_ Window) StandardWindowButton(b WindowButton) Button {
	rv := objc.Call[Button](w_, objc.Sel("standardWindowButton:"), b)
	return rv
}

// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419513-zoom?language=objc
func (w_ Window) Zoom(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("zoom:"), sender)
}

// Merges all open windows into a single tabbed window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644639-mergeallwindows?language=objc
func (w_ Window) MergeAllWindows(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("mergeAllWindows:"), sender)
	return rv
}

// Updates the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419577-update?language=objc
func (w_ Window) Update() {
	objc.Call[objc.Void](w_, objc.Sel("update"))
}

// Converts a point from the screen coordinate system to the window’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2967180-convertpointfromscreen?language=objc
func (w_ Window) ConvertPointFromScreen(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("convertPointFromScreen:"), point)
	return rv
}

// Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419276-enablekeyequivalentfordefaultbut?language=objc
func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.Call[objc.Void](w_, objc.Sel("enableKeyEquivalentForDefaultButtonCell"))
}

// Toggles the visibility of the window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419554-toggletoolbarshown?language=objc
func (w_ Window) ToggleToolbarShown(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("toggleToolbarShown:"), sender)
}

// Sets the window’s miniaturized state to the value you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449566-setisminiaturized?language=objc
func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.Call[objc.Void](w_, objc.Sel("setIsMiniaturized:"), flag)
}

// Returns a backing store pixel-aligned rectangle in window coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419319-backingalignedrect?language=objc
func (w_ Window) BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}

// Modifies and returns a frame rectangle so that its top edge lies on a specific screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419779-constrainframerect?language=objc
func (w_ Window) ConstrainFrameRectToScreen(frameRect foundation.Rect, screen IScreen) foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("constrainFrameRect:toScreen:"), frameRect, objc.Ptr(screen))
	return rv
}

// Updates the layout of views in the window based on the current views and constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526910-layoutifneeded?language=objc
func (w_ Window) LayoutIfNeeded() {
	objc.Call[objc.Void](w_, objc.Sel("layoutIfNeeded"))
}

// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419294-minframewidthwithtitle?language=objc
func (wc _WindowClass) MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	rv := objc.Call[float64](wc, objc.Sel("minFrameWidthWithTitle:styleMask:"), title, style)
	return rv
}

// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419294-minframewidthwithtitle?language=objc
func Window_MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	return WindowClass.MinFrameWidthWithTitleStyleMask(title, style)
}

// Ends a document-modal session and dismisses the specified sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419497-endsheet?language=objc
func (w_ Window) EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse) {
	objc.Call[objc.Void](w_, objc.Sel("endSheet:returnCode:"), objc.Ptr(sheetWindow), returnCode)
}

// Informs the window that it has become the main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419084-becomemainwindow?language=objc
func (w_ Window) BecomeMainWindow() {
	objc.Call[objc.Void](w_, objc.Sel("becomeMainWindow"))
}

// Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419275-inserttitlebaraccessoryviewcontr?language=objc
func (w_ Window) InsertTitlebarAccessoryViewControllerAtIndex(childViewController ITitlebarAccessoryViewController, index int) {
	objc.Call[objc.Void](w_, objc.Sel("insertTitlebarAccessoryViewController:atIndex:"), objc.Ptr(childViewController), index)
}

// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419450-performzoom?language=objc
func (w_ Window) PerformZoom(sender objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("performZoom:"), sender)
}

// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419658-setframetopleftpoint?language=objc
func (w_ Window) SetFrameTopLeftPoint(point foundation.Point) {
	objc.Call[objc.Void](w_, objc.Sel("setFrameTopLeftPoint:"), point)
}

// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419096-displayifneeded?language=objc
func (w_ Window) DisplayIfNeeded() {
	objc.Call[objc.Void](w_, objc.Sel("displayIfNeeded"))
}

// Sets the window’s location to the center of the screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419090-center?language=objc
func (w_ Window) Center() {
	objc.Call[objc.Void](w_, objc.Sel("center"))
}

// A Boolean value indicating whether the window configuration is preserved between application launches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc
func (w_ Window) IsRestorable() bool {
	rv := objc.Call[bool](w_, objc.Sel("isRestorable"))
	return rv
}

// A Boolean value indicating whether the window configuration is preserved between application launches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526255-restorable?language=objc
func (w_ Window) SetRestorable(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setRestorable:"), value)
}

// The maximum size to which the window’s frame (including its title bar) can be sized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc
func (w_ Window) MaxSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("maxSize"))
	return rv
}

// The maximum size to which the window’s frame (including its title bar) can be sized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419595-maxsize?language=objc
func (w_ Window) SetMaxSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setMaxSize:"), value)
}

// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc
func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.Call[bool](w_, objc.Sel("hidesOnDeactivate"))
	return rv
}

// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419777-hidesondeactivate?language=objc
func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setHidesOnDeactivate:"), value)
}

// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419132-visible?language=objc
func (w_ Window) IsVisible() bool {
	rv := objc.Call[bool](w_, objc.Sel("isVisible"))
	return rv
}

// The parent window to which the window is attached as a child. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc
func (w_ Window) ParentWindow() Window {
	rv := objc.Call[Window](w_, objc.Sel("parentWindow"))
	return rv
}

// The parent window to which the window is attached as a child. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419695-parentwindow?language=objc
func (w_ Window) SetParentWindow(value IWindow) {
	objc.Call[objc.Void](w_, objc.Sel("setParentWindow:"), objc.Ptr(value))
}

// A minimum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc
func (w_ Window) MinFullScreenContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("minFullScreenContentSize"))
	return rv
}

// A minimum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419627-minfullscreencontentsize?language=objc
func (w_ Window) SetMinFullScreenContentSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setMinFullScreenContentSize:"), value)
}

// An array of the sheets currently attached to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419765-sheets?language=objc
func (w_ Window) Sheets() []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("sheets"))
	return rv
}

// A Boolean value that indicates whether the window is excluded from the application’s Windows menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc
func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := objc.Call[bool](w_, objc.Sel("isExcludedFromWindowsMenu"))
	return rv
}

// A Boolean value that indicates whether the window is excluded from the application’s Windows menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419175-excludedfromwindowsmenu?language=objc
func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setExcludedFromWindowsMenu:"), value)
}

// A Boolean value that indicates whether the window has ever run as a modal sheet. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419364-sheet?language=objc
func (w_ Window) IsSheet() bool {
	rv := objc.Call[bool](w_, objc.Sel("isSheet"))
	return rv
}

// A Boolean value that indicates whether the window is released when it receives the close message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc
func (w_ Window) IsReleasedWhenClosed() bool {
	rv := objc.Call[bool](w_, objc.Sel("isReleasedWhenClosed"))
	return rv
}

// A Boolean value that indicates whether the window is released when it receives the close message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419062-releasedwhenclosed?language=objc
func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setReleasedWhenClosed:"), value)
}

// A Boolean value that indicates the level of access other processes have to the window’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419729-sharingtype?language=objc
func (w_ Window) SharingType() WindowSharingType {
	rv := objc.Call[WindowSharingType](w_, objc.Sel("sharingType"))
	return rv
}

// A Boolean value that indicates the level of access other processes have to the window’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419729-sharingtype?language=objc
func (w_ Window) SetSharingType(value WindowSharingType) {
	objc.Call[objc.Void](w_, objc.Sel("setSharingType:"), value)
}

// A Boolean value that indicates whether the window has a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (w_ Window) HasShadow() bool {
	rv := objc.Call[bool](w_, objc.Sel("hasShadow"))
	return rv
}

// A Boolean value that indicates whether the window has a shadow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419234-hasshadow?language=objc
func (w_ Window) SetHasShadow(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setHasShadow:"), value)
}

// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc
func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.Call[WindowStyleMask](w_, objc.Sel("styleMask"))
	return rv
}

// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419078-stylemask?language=objc
func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.Call[objc.Void](w_, objc.Sel("setStyleMask:"), value)
}

// An object that represents information about a window when it displays as a tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2870102-tab?language=objc
func (w_ Window) Tab() WindowTab {
	rv := objc.Call[WindowTab](w_, objc.Sel("tab"))
	return rv
}

// The zero-based position of the window, based on its order from front to back among all visible application windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc
func (w_ Window) OrderedIndex() int {
	rv := objc.Call[int](w_, objc.Sel("orderedIndex"))
	return rv
}

// The zero-based position of the window, based on its order from front to back among all visible application windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449577-orderedindex?language=objc
func (w_ Window) SetOrderedIndex(value int) {
	objc.Call[objc.Void](w_, objc.Sel("setOrderedIndex:"), value)
}

// A Boolean value that indicates whether any of the window’s views need to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc
func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.Call[bool](w_, objc.Sel("viewsNeedDisplay"))
	return rv
}

// A Boolean value that indicates whether any of the window’s views need to be displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419609-viewsneeddisplay?language=objc
func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setViewsNeedDisplay:"), value)
}

// The window level of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc
func (w_ Window) Level() WindowLevel {
	rv := objc.Call[WindowLevel](w_, objc.Sel("level"))
	return rv
}

// The window level of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419511-level?language=objc
func (w_ Window) SetLevel(value WindowLevel) {
	objc.Call[objc.Void](w_, objc.Sel("setLevel:"), value)
}

// The window’s automatic animation behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419763-animationbehavior?language=objc
func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := objc.Call[WindowAnimationBehavior](w_, objc.Sel("animationBehavior"))
	return rv
}

// The window’s automatic animation behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419763-animationbehavior?language=objc
func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	objc.Call[objc.Void](w_, objc.Sel("setAnimationBehavior:"), value)
}

// The screen the window is on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419232-screen?language=objc
func (w_ Window) Screen() Screen {
	rv := objc.Call[Screen](w_, objc.Sel("screen"))
	return rv
}

// A Boolean value that indicates whether the window is a modal panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449576-modalpanel?language=objc
func (w_ Window) IsModalPanel() bool {
	rv := objc.Call[bool](w_, objc.Sel("isModalPanel"))
	return rv
}

// The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419280-mouselocationoutsideofeventstrea?language=objc
func (w_ Window) MouseLocationOutsideOfEventStream() foundation.Point {
	rv := objc.Call[foundation.Point](w_, objc.Sel("mouseLocationOutsideOfEventStream"))
	return rv
}

// A Boolean value that indicates whether the window is the key window for the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419735-keywindow?language=objc
func (w_ Window) IsKeyWindow() bool {
	rv := objc.Call[bool](w_, objc.Sel("isKeyWindow"))
	return rv
}

// The window’s content view, the highest accessible view object in the window’s view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc
func (w_ Window) ContentView() View {
	rv := objc.Call[View](w_, objc.Sel("contentView"))
	return rv
}

// The window’s content view, the highest accessible view object in the window’s view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419160-contentview?language=objc
func (w_ Window) SetContentView(value IView) {
	objc.Call[objc.Void](w_, objc.Sel("setContentView:"), objc.Ptr(value))
}

// A value that indicates the visibility of the window’s title and title bar buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc
func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := objc.Call[WindowTitleVisibility](w_, objc.Sel("titleVisibility"))
	return rv
}

// A value that indicates the visibility of the window’s title and title bar buttons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419635-titlevisibility?language=objc
func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	objc.Call[objc.Void](w_, objc.Sel("setTitleVisibility:"), value)
}

// The style that determines the appearance and location of the toolbar in relation to the title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3608199-toolbarstyle?language=objc
func (w_ Window) ToolbarStyle() WindowToolbarStyle {
	rv := objc.Call[WindowToolbarStyle](w_, objc.Sel("toolbarStyle"))
	return rv
}

// The style that determines the appearance and location of the toolbar in relation to the title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3608199-toolbarstyle?language=objc
func (w_ Window) SetToolbarStyle(value WindowToolbarStyle) {
	objc.Call[objc.Void](w_, objc.Sel("setToolbarStyle:"), value)
}

// The event currently being processed by the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419298-currentevent?language=objc
func (w_ Window) CurrentEvent() Event {
	rv := objc.Call[Event](w_, objc.Sel("currentEvent"))
	return rv
}

// The title displayed in the window’s minimized window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc
func (w_ Window) MiniwindowTitle() string {
	rv := objc.Call[string](w_, objc.Sel("miniwindowTitle"))
	return rv
}

// The title displayed in the window’s minimized window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419571-miniwindowtitle?language=objc
func (w_ Window) SetMiniwindowTitle(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setMiniwindowTitle:"), value)
}

// A Boolean value that indicates whether the window is the application’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419130-mainwindow?language=objc
func (w_ Window) IsMainWindow() bool {
	rv := objc.Call[bool](w_, objc.Sel("isMainWindow"))
	return rv
}

// An array of the window’s attached child windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419236-childwindows?language=objc
func (w_ Window) ChildWindows() []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("childWindows"))
	return rv
}

// The button cell that performs as if clicked when the window receives a Return (or Enter) key event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419328-defaultbuttoncell?language=objc
func (w_ Window) DefaultButtonCell() ButtonCell {
	rv := objc.Call[ButtonCell](w_, objc.Sel("defaultButtonCell"))
	return rv
}

// The button cell that performs as if clicked when the window receives a Return (or Enter) key event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419328-defaultbuttoncell?language=objc
func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	objc.Call[objc.Void](w_, objc.Sel("setDefaultButtonCell:"), objc.Ptr(value))
}

// A value that indicates the user’s preference for window tabbing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646658-usertabbingpreference?language=objc
func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.Call[WindowUserTabbingPreference](wc, objc.Sel("userTabbingPreference"))
	return rv
}

// A value that indicates the user’s preference for window tabbing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646658-usertabbingpreference?language=objc
func Window_UserTabbingPreference() WindowUserTabbingPreference {
	return WindowClass.UserTabbingPreference()
}

// The window’s first responder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419440-firstresponder?language=objc
func (w_ Window) FirstResponder() Responder {
	rv := objc.Call[Responder](w_, objc.Sel("firstResponder"))
	return rv
}

// An array of windows that display as tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1792044-tabbedwindows?language=objc
func (w_ Window) TabbedWindows() []Window {
	rv := objc.Call[[]Window](w_, objc.Sel("tabbedWindows"))
	return rv
}

// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc
func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.Call[bool](w_, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}

// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419588-preservescontentduringliveresize?language=objc
func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setPreservesContentDuringLiveResize:"), value)
}

// Returns the default depth limit for instances of NSWindow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419049-defaultdepthlimit?language=objc
func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := objc.Call[WindowDepth](wc, objc.Sel("defaultDepthLimit"))
	return rv
}

// Returns the default depth limit for instances of NSWindow. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419049-defaultdepthlimit?language=objc
func Window_DefaultDepthLimit() WindowDepth {
	return WindowClass.DefaultDepthLimit()
}

// The window’s content aspect ratio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc
func (w_ Window) ContentAspectRatio() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("contentAspectRatio"))
	return rv
}

// The window’s content aspect ratio. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419148-contentaspectratio?language=objc
func (w_ Window) SetContentAspectRatio(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setContentAspectRatio:"), value)
}

// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc
func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.Call[bool](w_, objc.Sel("displaysWhenScreenProfileChanges"))
	return rv
}

// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419430-displayswhenscreenprofilechanges?language=objc
func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setDisplaysWhenScreenProfileChanges:"), value)
}

// The restoration class associated with the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526241-restorationclass?language=objc
func (w_ Window) RestorationClass() objc.Class {
	rv := objc.Call[objc.Class](w_, objc.Sel("restorationClass"))
	return rv
}

// The restoration class associated with the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1526241-restorationclass?language=objc
func (w_ Window) SetRestorationClass(value objc.IClass) {
	objc.Call[objc.Void](w_, objc.Sel("setRestorationClass:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the window accepts mouse-moved events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc
func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.Call[bool](w_, objc.Sel("acceptsMouseMovedEvents"))
	return rv
}

// A Boolean value that indicates whether the window accepts mouse-moved events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419340-acceptsmousemovedevents?language=objc
func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAcceptsMouseMovedEvents:"), value)
}

// A secondary line of text that appears in the title bar of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc
func (w_ Window) Subtitle() string {
	rv := objc.Call[string](w_, objc.Sel("subtitle"))
	return rv
}

// A secondary line of text that appears in the title bar of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3608198-subtitle?language=objc
func (w_ Window) SetSubtitle(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setSubtitle:"), value)
}

// A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419220-workswhenmodal?language=objc
func (w_ Window) WorksWhenModal() bool {
	rv := objc.Call[bool](w_, objc.Sel("worksWhenModal"))
	return rv
}

// A Boolean value that indicates whether the window is minimized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419699-miniaturized?language=objc
func (w_ Window) IsMiniaturized() bool {
	rv := objc.Call[bool](w_, objc.Sel("isMiniaturized"))
	return rv
}

// The window’s alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc
func (w_ Window) AlphaValue() float64 {
	rv := objc.Call[float64](w_, objc.Sel("alphaValue"))
	return rv
}

// The window’s alpha value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419186-alphavalue?language=objc
func (w_ Window) SetAlphaValue(value float64) {
	objc.Call[objc.Void](w_, objc.Sel("setAlphaValue:"), value)
}

// The deepest screen the window is on (it may be split over several screens). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419080-deepestscreen?language=objc
func (w_ Window) DeepestScreen() Screen {
	rv := objc.Call[Screen](w_, objc.Sel("deepestScreen"))
	return rv
}

// The area inside the window that is for non-obscured content, in window coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419124-contentlayoutrect?language=objc
func (w_ Window) ContentLayoutRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("contentLayoutRect"))
	return rv
}

// A Boolean value that indicates if the window has a title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449568-hastitlebar?language=objc
func (w_ Window) HasTitleBar() bool {
	rv := objc.Call[bool](w_, objc.Sel("hasTitleBar"))
	return rv
}

// A Boolean value that indicates whether the window can minimize. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449583-miniaturizable?language=objc
func (w_ Window) IsMiniaturizable() bool {
	rv := objc.Call[bool](w_, objc.Sel("isMiniaturizable"))
	return rv
}

// The minimum size to which the window’s frame (including its title bar) can be sized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc
func (w_ Window) MinSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("minSize"))
	return rv
}

// The minimum size to which the window’s frame (including its title bar) can be sized. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419206-minsize?language=objc
func (w_ Window) SetMinSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setMinSize:"), value)
}

// The window’s backing store type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc
func (w_ Window) BackingType() BackingStoreType {
	rv := objc.Call[BackingStoreType](w_, objc.Sel("backingType"))
	return rv
}

// The window’s backing store type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419599-backingtype?language=objc
func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.Call[objc.Void](w_, objc.Sel("setBackingType:"), value)
}

// The custom miniaturized window image of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc
func (w_ Window) MiniwindowImage() Image {
	rv := objc.Call[Image](w_, objc.Sel("miniwindowImage"))
	return rv
}

// The custom miniaturized window image of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419185-miniwindowimage?language=objc
func (w_ Window) SetMiniwindowImage(value IImage) {
	objc.Call[objc.Void](w_, objc.Sel("setMiniwindowImage:"), objc.Ptr(value))
}

// The window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419731-toolbar?language=objc
func (w_ Window) Toolbar() Toolbar {
	rv := objc.Call[Toolbar](w_, objc.Sel("toolbar"))
	return rv
}

// The window’s toolbar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419731-toolbar?language=objc
func (w_ Window) SetToolbar(value IToolbar) {
	objc.Call[objc.Void](w_, objc.Sel("setToolbar:"), objc.Ptr(value))
}

// The window’s color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419569-colorspace?language=objc
func (w_ Window) ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](w_, objc.Sel("colorSpace"))
	return rv
}

// The window’s color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419569-colorspace?language=objc
func (w_ Window) SetColorSpace(value IColorSpace) {
	objc.Call[objc.Void](w_, objc.Sel("setColorSpace:"), objc.Ptr(value))
}

// The depth limit of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419613-depthlimit?language=objc
func (w_ Window) DepthLimit() WindowDepth {
	rv := objc.Call[WindowDepth](w_, objc.Sel("depthLimit"))
	return rv
}

// The depth limit of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419613-depthlimit?language=objc
func (w_ Window) SetDepthLimit(value WindowDepth) {
	objc.Call[objc.Void](w_, objc.Sel("setDepthLimit:"), value)
}

// The window’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc
func (w_ Window) Delegate() WindowDelegateWrapper {
	rv := objc.Call[WindowDelegateWrapper](w_, objc.Sel("delegate"))
	return rv
}

// The window’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc
func (w_ Window) SetDelegate(value PWindowDelegate) {
	po0 := objc.WrapAsProtocol("NSWindowDelegate", value)
	objc.SetAssociatedObject(w_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](w_, objc.Sel("setDelegate:"), po0)
}

// The window’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419060-delegate?language=objc
func (w_ Window) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The flags field of the event record for the mouse-down event that initiated the resizing session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419302-resizeflags?language=objc
func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](w_, objc.Sel("resizeFlags"))
	return rv
}

// A Boolean value that indicates whether the window is a floating panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449579-floatingpanel?language=objc
func (w_ Window) IsFloatingPanel() bool {
	rv := objc.Call[bool](w_, objc.Sel("isFloatingPanel"))
	return rv
}

// A Boolean value that indicates whether the window is being resized by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419378-inliveresize?language=objc
func (w_ Window) InLiveResize() bool {
	rv := objc.Call[bool](w_, objc.Sel("inLiveResize"))
	return rv
}

// The color of the window’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc
func (w_ Window) BackgroundColor() Color {
	rv := objc.Call[Color](w_, objc.Sel("backgroundColor"))
	return rv
}

// The color of the window’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419751-backgroundcolor?language=objc
func (w_ Window) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](w_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the title bar draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc
func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.Call[bool](w_, objc.Sel("titlebarAppearsTransparent"))
	return rv
}

// A Boolean value that indicates whether the title bar draws its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419167-titlebarappearstransparent?language=objc
func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setTitlebarAppearsTransparent:"), value)
}

// An array of title bar accessory view controllers that are currently added to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc
func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := objc.Call[[]TitlebarAccessoryViewController](w_, objc.Sel("titlebarAccessoryViewControllers"))
	return rv
}

// An array of title bar accessory view controllers that are currently added to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419547-titlebaraccessoryviewcontrollers?language=objc
func (w_ Window) SetTitlebarAccessoryViewControllers(value []ITitlebarAccessoryViewController) {
	objc.Call[objc.Void](w_, objc.Sel("setTitlebarAccessoryViewControllers:"), value)
}

// The window’s resizing increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc
func (w_ Window) ResizeIncrements() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("resizeIncrements"))
	return rv
}

// The window’s resizing increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419390-resizeincrements?language=objc
func (w_ Window) SetResizeIncrements(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setResizeIncrements:"), value)
}

// The Carbon window reference associated with the window, creating one if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419485-windowref?language=objc
func (w_ Window) WindowRef() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](w_, objc.Sel("windowRef"))
	return rv
}

// A Boolean value that indicates whether the window can hide when its application becomes hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc
func (w_ Window) CanHide() bool {
	rv := objc.Call[bool](w_, objc.Sel("canHide"))
	return rv
}

// A Boolean value that indicates whether the window can hide when its application becomes hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419725-canhide?language=objc
func (w_ Window) SetCanHide(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setCanHide:"), value)
}

// A Boolean value that indicates whether the window is opaque. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc
func (w_ Window) IsOpaque() bool {
	rv := objc.Call[bool](w_, objc.Sel("isOpaque"))
	return rv
}

// A Boolean value that indicates whether the window is opaque. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419086-opaque?language=objc
func (w_ Window) SetOpaque(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setOpaque:"), value)
}

// The URL of the file the window represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc
func (w_ Window) RepresentedURL() foundation.URL {
	rv := objc.Call[foundation.URL](w_, objc.Sel("representedURL"))
	return rv
}

// The URL of the file the window represents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419066-representedurl?language=objc
func (w_ Window) SetRepresentedURL(value foundation.IURL) {
	objc.Call[objc.Void](w_, objc.Sel("setRepresentedURL:"), objc.Ptr(value))
}

// The view that’s made first responder (also called the key view) the first time the window is placed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc
func (w_ Window) InitialFirstResponder() View {
	rv := objc.Call[View](w_, objc.Sel("initialFirstResponder"))
	return rv
}

// The view that’s made first responder (also called the key view) the first time the window is placed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419479-initialfirstresponder?language=objc
func (w_ Window) SetInitialFirstResponder(value IView) {
	objc.Call[objc.Void](w_, objc.Sel("setInitialFirstResponder:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the window allows zooming. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449587-zoomable?language=objc
func (w_ Window) IsZoomable() bool {
	rv := objc.Call[bool](w_, objc.Sel("isZoomable"))
	return rv
}

// The occlusion state of the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419321-occlusionstate?language=objc
func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := objc.Call[WindowOcclusionState](w_, objc.Sel("occlusionState"))
	return rv
}

// A Boolean value that indicates whether the window allows multithreaded view drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc
func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsConcurrentViewDrawing"))
	return rv
}

// A Boolean value that indicates whether the window allows multithreaded view drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419300-allowsconcurrentviewdrawing?language=objc
func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsConcurrentViewDrawing:"), value)
}

// A Boolean value that indicates whether the window can be displayed at the login window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc
func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.Call[bool](w_, objc.Sel("canBecomeVisibleWithoutLogin"))
	return rv
}

// A Boolean value that indicates whether the window can be displayed at the login window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419179-canbecomevisiblewithoutlogin?language=objc
func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setCanBecomeVisibleWithoutLogin:"), value)
}

// A Boolean value that indicates whether the window’s cursor rectangles are enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419668-arecursorrectsenabled?language=objc
func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.Call[bool](w_, objc.Sel("areCursorRectsEnabled"))
	return rv
}

// The main content view controller for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc
func (w_ Window) ContentViewController() ViewController {
	rv := objc.Call[ViewController](w_, objc.Sel("contentViewController"))
	return rv
}

// The main content view controller for the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419615-contentviewcontroller?language=objc
func (w_ Window) SetContentViewController(value IViewController) {
	objc.Call[objc.Void](w_, objc.Sel("setContentViewController:"), objc.Ptr(value))
}

// A dictionary containing information about the window’s resolution, such as color, depth, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419741-devicedescription?language=objc
func (w_ Window) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	rv := objc.Call[map[DeviceDescriptionKey]objc.Object](w_, objc.Sel("deviceDescription"))
	return rv
}

// The minimum size of the window’s content view in the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc
func (w_ Window) ContentMinSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("contentMinSize"))
	return rv
}

// The minimum size of the window’s content view in the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419670-contentminsize?language=objc
func (w_ Window) SetContentMinSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setContentMinSize:"), value)
}

// The type of separator that the app displays between the title bar and content of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3622489-titlebarseparatorstyle?language=objc
func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.Call[TitlebarSeparatorStyle](w_, objc.Sel("titlebarSeparatorStyle"))
	return rv
}

// The type of separator that the app displays between the title bar and content of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/3622489-titlebarseparatorstyle?language=objc
func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.Call[objc.Void](w_, objc.Sel("setTitlebarSeparatorStyle:"), value)
}

// A Boolean value that indicates whether the window’s document has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc
func (w_ Window) IsDocumentEdited() bool {
	rv := objc.Call[bool](w_, objc.Sel("isDocumentEdited"))
	return rv
}

// A Boolean value that indicates whether the window’s document has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419311-documentedited?language=objc
func (w_ Window) SetDocumentEdited(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setDocumentEdited:"), value)
}

// A value that identifies the window’s behavior in window collections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc
func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := objc.Call[WindowCollectionBehavior](w_, objc.Sel("collectionBehavior"))
	return rv
}

// A value that identifies the window’s behavior in window collections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419471-collectionbehavior?language=objc
func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	objc.Call[objc.Void](w_, objc.Sel("setCollectionBehavior:"), value)
}

// A value that allows a group of related windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644704-tabbingidentifier?language=objc
func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier {
	rv := objc.Call[WindowTabbingIdentifier](w_, objc.Sel("tabbingIdentifier"))
	return rv
}

// A value that allows a group of related windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644704-tabbingidentifier?language=objc
func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier) {
	objc.Call[objc.Void](w_, objc.Sel("setTabbingIdentifier:"), value)
}

// The window’s window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419092-windowcontroller?language=objc
func (w_ Window) WindowController() WindowController {
	rv := objc.Call[WindowController](w_, objc.Sel("windowController"))
	return rv
}

// The window’s window controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419092-windowcontroller?language=objc
func (w_ Window) SetWindowController(value IWindowController) {
	objc.Call[objc.Void](w_, objc.Sel("setWindowController:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the window prevents application termination when modal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc
func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.Call[bool](w_, objc.Sel("preventsApplicationTerminationWhenModal"))
	return rv
}

// A Boolean value that indicates whether the window prevents application termination when modal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419743-preventsapplicationterminationwh?language=objc
func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setPreventsApplicationTerminationWhenModal:"), value)
}

// The backing scale factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419459-backingscalefactor?language=objc
func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.Call[float64](w_, objc.Sel("backingScaleFactor"))
	return rv
}

// A group of windows that display together as a tab group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2879189-tabgroup?language=objc
func (w_ Window) TabGroup() WindowTabGroup {
	rv := objc.Call[WindowTabGroup](w_, objc.Sel("tabGroup"))
	return rv
}

// The direction the window’s title bar lays text out, either left to right or right to left. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644535-windowtitlebarlayoutdirection?language=objc
func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](w_, objc.Sel("windowTitlebarLayoutDirection"))
	return rv
}

// A value that indicates when a window displays tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644729-tabbingmode?language=objc
func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := objc.Call[WindowTabbingMode](w_, objc.Sel("tabbingMode"))
	return rv
}

// A value that indicates when a window displays tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1644729-tabbingmode?language=objc
func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	objc.Call[objc.Void](w_, objc.Sel("setTabbingMode:"), value)
}

// The path to the file of the window’s represented file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc
func (w_ Window) RepresentedFilename() string {
	rv := objc.Call[string](w_, objc.Sel("representedFilename"))
	return rv
}

// The path to the file of the window’s represented file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419631-representedfilename?language=objc
func (w_ Window) SetRepresentedFilename(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setRepresentedFilename:"), value)
}

// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc
func (w_ Window) AspectRatio() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("aspectRatio"))
	return rv
}

// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419507-aspectratio?language=objc
func (w_ Window) SetAspectRatio(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setAspectRatio:"), value)
}

// The window number of the window’s window device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419068-windownumber?language=objc
func (w_ Window) WindowNumber() int {
	rv := objc.Call[int](w_, objc.Sel("windowNumber"))
	return rv
}

// The window’s frame rectangle in screen coordinates, including the title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419697-frame?language=objc
func (w_ Window) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](w_, objc.Sel("frame"))
	return rv
}

// A Boolean value that indicates whether the window can become the application’s main window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419162-canbecomemainwindow?language=objc
func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.Call[bool](w_, objc.Sel("canBecomeMainWindow"))
	return rv
}

// The string that appears in the title bar of the window or the path to the represented file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc
func (w_ Window) Title() string {
	rv := objc.Call[string](w_, objc.Sel("title"))
	return rv
}

// The string that appears in the title bar of the window or the path to the represented file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419404-title?language=objc
func (w_ Window) SetTitle(value string) {
	objc.Call[objc.Void](w_, objc.Sel("setTitle:"), value)
}

// A Boolean value that indicates if the window has a close box. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449574-hasclosebox?language=objc
func (w_ Window) HasCloseBox() bool {
	rv := objc.Call[bool](w_, objc.Sel("hasCloseBox"))
	return rv
}

// A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419330-hasdynamicdepthlimit?language=objc
func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.Call[bool](w_, objc.Sel("hasDynamicDepthLimit"))
	return rv
}

// A string representation of the window’s frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419515-stringwithsavedframe?language=objc
func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor {
	rv := objc.Call[WindowPersistableFrameDescriptor](w_, objc.Sel("stringWithSavedFrame"))
	return rv
}

// A Boolean value that indicates if the user can resize the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1449572-resizable?language=objc
func (w_ Window) IsResizable() bool {
	rv := objc.Call[bool](w_, objc.Sel("isResizable"))
	return rv
}

// The application’s Dock tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419088-docktile?language=objc
func (w_ Window) DockTile() DockTile {
	rv := objc.Call[DockTile](w_, objc.Sel("dockTile"))
	return rv
}

// A Boolean value that indicates whether the window is on the currently active space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419707-onactivespace?language=objc
func (w_ Window) IsOnActiveSpace() bool {
	rv := objc.Call[bool](w_, objc.Sel("isOnActiveSpace"))
	return rv
}

// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc
func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Call[bool](w_, objc.Sel("autorecalculatesKeyViewLoop"))
	return rv
}

// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419214-autorecalculateskeyviewloop?language=objc
func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAutorecalculatesKeyViewLoop:"), value)
}

// An object that the window inherits its appearance from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2998855-appearancesource?language=objc
func (w_ Window) AppearanceSource() objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("appearanceSource"))
	return rv
}

// An object that the window inherits its appearance from. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/2998855-appearancesource?language=objc
func (w_ Window) SetAppearanceSource(value objc.IObject) {
	objc.Call[objc.Void](w_, objc.Sel("setAppearanceSource:"), objc.Ptr(value))
}

// The sheet attached to the window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419467-attachedsheet?language=objc
func (w_ Window) AttachedSheet() Window {
	rv := objc.Call[Window](w_, objc.Sel("attachedSheet"))
	return rv
}

// A Boolean value that indicates whether the window is transparent to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc
func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.Call[bool](w_, objc.Sel("ignoresMouseEvents"))
	return rv
}

// A Boolean value that indicates whether the window is transparent to mouse events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419354-ignoresmouseevents?language=objc
func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setIgnoresMouseEvents:"), value)
}

// A Boolean value that indicates whether the app can automatically organize windows into tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Call[bool](wc, objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}

// A Boolean value that indicates whether the app can automatically organize windows into tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func Window_AllowsAutomaticWindowTabbing() bool {
	return WindowClass.AllowsAutomaticWindowTabbing()
}

// A Boolean value that indicates whether the app can automatically organize windows into tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func (wc _WindowClass) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.Call[objc.Void](wc, objc.Sel("setAllowsAutomaticWindowTabbing:"), value)
}

// A Boolean value that indicates whether the app can automatically organize windows into tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1646657-allowsautomaticwindowtabbing?language=objc
func Window_SetAllowsAutomaticWindowTabbing(value bool) {
	WindowClass.SetAllowsAutomaticWindowTabbing(value)
}

// A value used by Auto Layout constraints to automatically bind to the value of contentLayoutRect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419094-contentlayoutguide?language=objc
func (w_ Window) ContentLayoutGuide() objc.Object {
	rv := objc.Call[objc.Object](w_, objc.Sel("contentLayoutGuide"))
	return rv
}

// The direction the window is currently using to change the key view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419158-keyviewselectiondirection?language=objc
func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := objc.Call[SelectionDirection](w_, objc.Sel("keyViewSelectionDirection"))
	return rv
}

// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc
func (w_ Window) IsMovableByWindowBackground() bool {
	rv := objc.Call[bool](w_, objc.Sel("isMovableByWindowBackground"))
	return rv
}

// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419072-movablebywindowbackground?language=objc
func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setMovableByWindowBackground:"), value)
}

// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc
func (w_ Window) IsMovable() bool {
	rv := objc.Call[bool](w_, objc.Sel("isMovable"))
	return rv
}

// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419579-movable?language=objc
func (w_ Window) SetMovable(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setMovable:"), value)
}

// A Boolean value that indicates whether the window is in a zoomed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419398-zoomed?language=objc
func (w_ Window) IsZoomed() bool {
	rv := objc.Call[bool](w_, objc.Sel("isZoomed"))
	return rv
}

// The maximum size of the window’s content view in the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc
func (w_ Window) ContentMaxSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("contentMaxSize"))
	return rv
}

// The maximum size of the window’s content view in the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419154-contentmaxsize?language=objc
func (w_ Window) SetContentMaxSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setContentMaxSize:"), value)
}

// A Boolean value that indicates whether the window can become the key window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419543-canbecomekeywindow?language=objc
func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.Call[bool](w_, objc.Sel("canBecomeKeyWindow"))
	return rv
}

// The window’s content-view resizing increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc
func (w_ Window) ContentResizeIncrements() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("contentResizeIncrements"))
	return rv
}

// The window’s content-view resizing increments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419649-contentresizeincrements?language=objc
func (w_ Window) SetContentResizeIncrements(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setContentResizeIncrements:"), value)
}

// The name used to automatically save the window’s frame rectangle data in the defaults system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419362-frameautosavename?language=objc
func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName {
	rv := objc.Call[WindowFrameAutosaveName](w_, objc.Sel("frameAutosaveName"))
	return rv
}

// A maximum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc
func (w_ Window) MaxFullScreenContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](w_, objc.Sel("maxFullScreenContentSize"))
	return rv
}

// A maximum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419438-maxfullscreencontentsize?language=objc
func (w_ Window) SetMaxFullScreenContentSize(value foundation.Size) {
	objc.Call[objc.Void](w_, objc.Sel("setMaxFullScreenContentSize:"), value)
}

// A Boolean value that indicates whether the window can display tooltips even when the application is in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc
func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.Call[bool](w_, objc.Sel("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}

// A Boolean value that indicates whether the window can display tooltips even when the application is in the background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419138-allowstooltipswhenapplicationisi?language=objc
func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.Call[objc.Void](w_, objc.Sel("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}

// The window to which the sheet is attached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419052-sheetparent?language=objc
func (w_ Window) SheetParent() Window {
	rv := objc.Call[Window](w_, objc.Sel("sheetParent"))
	return rv
}
