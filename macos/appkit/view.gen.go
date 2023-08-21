// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [View] class.
var ViewClass = _ViewClass{objc.GetClass("NSView")}

type _ViewClass struct {
	objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	IResponder
	RotateByAngle(angle float64)
	AdjustPageHeightNewTopBottomLimit(newBottom *float64, oldTop float64, oldBottom float64, bottomLimit float64)
	RemoveCursorRectCursor(rect foundation.Rect, object ICursor)
	GetRectsExposedDuringLiveResizeCount(exposedRects *foundation.Rect, count *int)
	RemoveToolTip(tag ToolTipTag)
	AddSubview(view IView)
	EnterFullScreenModeWithOptions(screen IScreen, options map[ViewFullScreenModeOptionKey]objc.IObject) bool
	ViewDidChangeEffectiveAppearance()
	DiscardCursorRects()
	ConvertPointToBacking(point foundation.Point) foundation.Point
	SetFrameOrigin(newOrigin foundation.Point)
	ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.IObject)
	DidCloseMenuWithEvent(menu IMenu, event IEvent)
	AddConstraint(constraint ILayoutConstraint)
	AddConstraints(constraints []ILayoutConstraint)
	AddTrackingRectOwnerUserDataAssumeInside(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer, flag bool) TrackingRectTag
	ConvertSizeToBacking(size foundation.Size) foundation.Size
	WillRemoveSubview(subview IView)
	AdjustScroll(newVisible foundation.Rect) foundation.Rect
	AddGestureRecognizer(gestureRecognizer IGestureRecognizer)
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep
	ShowDefinitionForAttributedStringAtPoint(attrString foundation.IAttributedString, textBaselineOrigin foundation.Point)
	NeedsToDrawRect(rect foundation.Rect) bool
	EndPage()
	HitTest(point foundation.Point) View
	DrawPageBorderWithSize(borderSize foundation.Size)
	AddTrackingArea(trackingArea ITrackingArea)
	ScrollRectToVisible(rect foundation.Rect) bool
	ResetCursorRects()
	ShouldDelayWindowOrderingForEvent(event IEvent) bool
	ConvertSizeFromBacking(size foundation.Size) foundation.Size
	WritePDFInsideRectToPasteboard(rect foundation.Rect, pasteboard IPasteboard)
	UpdateConstraintsForSubtreeIfNeeded()
	ConvertPointToLayer(point foundation.Point) foundation.Point
	RectForSmartMagnificationAtPointInRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect
	ScrollPoint(point foundation.Point)
	ViewDidMoveToSuperview()
	MenuForEvent(event IEvent) Menu
	DataWithEPSInsideRect(rect foundation.Rect) []byte
	AncestorSharedWithView(view IView) View
	UpdateLayer()
	AddCursorRectCursor(rect foundation.Rect, object ICursor)
	RemoveFromSuperview()
	EndDocument()
	TranslateRectsNeedingDisplayInRectBy(clipRect foundation.Rect, delta foundation.Size)
	RemoveLayoutGuide(guide ILayoutGuide)
	DisplayIfNeededInRect(rect foundation.Rect)
	Autoscroll(event IEvent) bool
	ConvertRectFromView(rect foundation.Rect, view IView) foundation.Rect
	AcceptsFirstMouse(event IEvent) bool
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	RemoveConstraint(constraint ILayoutConstraint)
	CenterScanRect(rect foundation.Rect) foundation.Rect
	ScrollClipViewToPoint(clipView IClipView, point foundation.Point)
	TranslateOriginToPoint(translation foundation.Point)
	DisplayRectIgnoringOpacity(rect foundation.Rect)
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	ViewDidUnhide()
	UpdateTrackingAreas()
	ViewDidChangeBackingProperties()
	RegisterForDraggedTypes(newTypes []PasteboardType)
	ViewDidHide()
	CacheDisplayInRectToBitmapImageRep(rect foundation.Rect, bitmapImageRep IBitmapImageRep)
	RemoveFromSuperviewWithoutNeedingDisplay()
	BeginDraggingSessionWithItemsEventSource(items []IDraggingItem, event IEvent, source PDraggingSource) DraggingSession
	BeginDraggingSessionWithItemsEventSourceObject(items []IDraggingItem, event IEvent, sourceObject objc.IObject) DraggingSession
	WillOpenMenuWithEvent(menu IMenu, event IEvent)
	Display()
	ResizeWithOldSuperviewSize(oldSize foundation.Size)
	AddToolTipRectOwnerUserData(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer) ToolTipTag
	SetContentHuggingPriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	ReflectScrolledClipView(clipView IClipView)
	SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect)
	UpdateConstraints()
	RemoveTrackingRect(tag TrackingRectTag)
	SetFrameSize(newSize foundation.Size)
	SetBoundsOrigin(newOrigin foundation.Point)
	LocationOfPrintRect(rect foundation.Rect) foundation.Point
	ConvertPointFromLayer(point foundation.Point) foundation.Point
	ConvertSizeFromLayer(size foundation.Size) foundation.Size
	ReplaceSubviewWith(oldView IView, newView IView)
	ConvertPointFromBacking(point foundation.Point) foundation.Point
	KnowsPageRange(range_ foundation.RangePointer) bool
	RulerViewLocationForPoint(ruler IRulerView, point foundation.Point) float64
	Print(sender objc.IObject)
	ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	PrepareContentInRect(rect foundation.Rect)
	RemoveConstraints(constraints []ILayoutConstraint)
	BeginDocument()
	ResizeSubviewsWithOldSize(oldSize foundation.Size)
	SetBoundsSize(newSize foundation.Size)
	DrawRect(dirtyRect foundation.Rect)
	DrawFocusRingMask()
	ConvertRectToLayer(rect foundation.Rect) foundation.Rect
	ScaleUnitSquareToSize(newUnitSize foundation.Size)
	BeginPageInRectAtPlacement(rect foundation.Rect, location foundation.Point)
	ViewWillMoveToSuperview(newSuperview IView)
	ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	Layout()
	PrepareForReuse()
	RemoveTrackingArea(trackingArea ITrackingArea)
	RemoveAllToolTips()
	DataWithPDFInsideRect(rect foundation.Rect) []byte
	InvalidateIntrinsicContentSize()
	SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	SortSubviewsUsingFunctionContext(compare func(arg0 View, arg1 View, arg2 unsafe.Pointer) foundation.ComparisonResult, context unsafe.Pointer)
	ViewDidEndLiveResize()
	ViewWillMoveToWindow(newWindow IWindow)
	DisplayIfNeededIgnoringOpacity()
	AlignmentRectForFrame(frame foundation.Rect) foundation.Rect
	UnregisterDraggedTypes()
	RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer)
	AdjustPageWidthNewLeftRightLimit(newRight *float64, oldLeft float64, oldRight float64, rightLimit float64)
	ViewWillStartLiveResize()
	RectForPage(page int) foundation.Rect
	ViewWillDraw()
	AddLayoutGuide(guide ILayoutGuide)
	NoteFocusRingMaskChanged()
	ViewWithTag(tag int) View
	WriteEPSInsideRectToPasteboard(rect foundation.Rect, pasteboard IPasteboard)
	ExerciseAmbiguityInLayout()
	ViewDidMoveToWindow()
	LayoutSubtreeIfNeeded()
	GetRectsBeingDrawnCount(rects *foundation.Rect, count *int)
	ConvertSizeToLayer(size foundation.Size) foundation.Size
	ConvertRectFromLayer(rect foundation.Rect) foundation.Rect
	DisplayRect(rect foundation.Rect)
	DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect)
	SetNeedsDisplayInRect(invalidRect foundation.Rect)
	MouseInRect(point foundation.Point, rect foundation.Rect) bool
	BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect
	DidAddSubview(subview IView)
	FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect
	ConvertSizeToView(size foundation.Size, view IView) foundation.Size
	IsDescendantOf(view IView) bool
	DisplayIfNeeded()
	MakeBackingLayer() quartzcore.Layer
	ConvertPointToView(point foundation.Point, view IView) foundation.Point
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	TopAnchor() LayoutYAxisAnchor
	LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy
	SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy)
	Layer() quartzcore.Layer
	SetLayer(value quartzcore.ILayer)
	PrintJobTitle() string
	LeftAnchor() LayoutXAxisAnchor
	VisibleRect() foundation.Rect
	InputContext() TextInputContext
	IsRotatedOrScaledFromBase() bool
	IsHidden() bool
	SetHidden(value bool)
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	NextKeyView() View
	SetNextKeyView(value IView)
	FirstBaselineOffsetFromTop() float64
	FittingSize() foundation.Size
	BoundsRotation() float64
	SetBoundsRotation(value float64)
	AlignmentRectInsets() foundation.EdgeInsets
	PageFooter() foundation.AttributedString
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	NeedsPanelToBecomeKey() bool
	ContentFilters() []coreimage.Filter
	SetContentFilters(value []coreimage.IFilter)
	NextValidKeyView() View
	ToolTip() string
	SetToolTip(value string)
	FrameCenterRotation() float64
	SetFrameCenterRotation(value float64)
	CenterYAnchor() LayoutYAxisAnchor
	PreservesContentDuringLiveResize() bool
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	FirstBaselineAnchor() LayoutYAxisAnchor
	LastBaselineOffsetFromBottom() float64
	PressureConfiguration() PressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	WantsUpdateLayer() bool
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	TrackingAreas() []TrackingArea
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	GestureRecognizers() []GestureRecognizer
	SetGestureRecognizers(value []IGestureRecognizer)
	EnclosingScrollView() ScrollView
	AlphaValue() float64
	SetAlphaValue(value float64)
	PageHeader() foundation.AttributedString
	SafeAreaLayoutGuide() LayoutGuide
	RightAnchor() LayoutXAxisAnchor
	IsRotatedFromBase() bool
	PreparedContentRect() foundation.Rect
	SetPreparedContentRect(value foundation.Rect)
	Bounds() foundation.Rect
	SetBounds(value foundation.Rect)
	BaselineOffsetFromBottom() float64
	SafeAreaInsets() foundation.EdgeInsets
	InLiveResize() bool
	IsDrawingFindIndicator() bool
	CandidateListTouchBarItem() CandidateListTouchBarItem
	LayoutMarginsGuide() LayoutGuide
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	BottomAnchor() LayoutYAxisAnchor
	OpaqueAncestor() View
	LayerContentsPlacement() ViewLayerContentsPlacement
	SetLayerContentsPlacement(value ViewLayerContentsPlacement)
	IsOpaque() bool
	PreviousValidKeyView() View
	EnclosingMenuItem() MenuItem
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	HeightAnchor() LayoutDimension
	LayoutGuides() []LayoutGuide
	WantsDefaultClipping() bool
	SafeAreaRect() foundation.Rect
	WantsLayer() bool
	SetWantsLayer(value bool)
	IntrinsicContentSize() foundation.Size
	AutoresizesSubviews() bool
	SetAutoresizesSubviews(value bool)
	TrailingAnchor() LayoutXAxisAnchor
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	PreviousKeyView() View
	LeadingAnchor() LayoutXAxisAnchor
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	FocusRingMaskBounds() foundation.Rect
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	CanBecomeKeyView() bool
	IsInFullScreenMode() bool
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	Tag() int
	Shadow() Shadow
	SetShadow(value IShadow)
	WidthAdjustLimit() float64
	AdditionalSafeAreaInsets() foundation.EdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.EdgeInsets)
	CompositingFilter() coreimage.Filter
	SetCompositingFilter(value coreimage.IFilter)
	RegisteredDraggedTypes() []PasteboardType
	IsHiddenOrHasHiddenAncestor() bool
	Window() Window
	IsVerticalContentSizeConstraintActive() bool
	SetVerticalContentSizeConstraintActive(value bool)
	AutoresizingMask() AutoresizingMaskOptions
	SetAutoresizingMask(value AutoresizingMaskOptions)
	BackgroundFilters() []coreimage.Filter
	SetBackgroundFilters(value []coreimage.IFilter)
	HasAmbiguousLayout() bool
	Constraints() []LayoutConstraint
	HeightAdjustLimit() float64
	IsHorizontalContentSizeConstraintActive() bool
	SetHorizontalContentSizeConstraintActive(value bool)
	AllowsVibrancy() bool
	Subviews() []View
	SetSubviews(value []IView)
	IsFlipped() bool
	RectPreservedDuringLiveResize() foundation.Rect
	MouseDownCanMoveWindow() bool
	CenterXAnchor() LayoutXAxisAnchor
	WantsRestingTouches() bool
	SetWantsRestingTouches(value bool)
	LastBaselineAnchor() LayoutYAxisAnchor
	WidthAnchor() LayoutDimension
	FrameRotation() float64
	SetFrameRotation(value float64)
	Superview() View
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
}

// The infrastructure for drawing, printing, and handling events in an app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview?language=objc
type View struct {
	Responder
}

func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		Responder: ResponderFrom(ptr),
	}
}

func (v_ View) InitWithFrame(frameRect foundation.Rect) View {
	rv := objc.Call[View](v_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewViewWithFrame(frameRect foundation.Rect) View {
	instance := ViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

func (vc _ViewClass) Alloc() View {
	rv := objc.Call[View](vc, objc.Sel("alloc"))
	return rv
}

func View_Alloc() View {
	return ViewClass.Alloc()
}

func (vc _ViewClass) New() View {
	rv := objc.Call[View](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewView() View {
	return ViewClass.New()
}

func (v_ View) Init() View {
	rv := objc.Call[View](v_, objc.Sel("init"))
	return rv
}

// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483444-rotatebyangle?language=objc
func (v_ View) RotateByAngle(angle float64) {
	objc.Call[objc.Void](v_, objc.Sel("rotateByAngle:"), angle)
}

// Overridden by subclasses to adjust page height during automatic pagination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483258-adjustpageheightnew?language=objc
func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom *float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Call[objc.Void](v_, objc.Sel("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}

// Completely removes a cursor rectangle from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483676-removecursorrect?language=objc
func (v_ View) RemoveCursorRectCursor(rect foundation.Rect, object ICursor) {
	objc.Call[objc.Void](v_, objc.Sel("removeCursorRect:cursor:"), rect, objc.Ptr(object))
}

// Returns a list of rectangles indicating the newly exposed areas of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483221-getrectsexposedduringliveresize?language=objc
func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects *foundation.Rect, count *int) {
	objc.Call[objc.Void](v_, objc.Sel("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}

// Removes the tooltip identified by specified tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483522-removetooltip?language=objc
func (v_ View) RemoveToolTip(tag ToolTipTag) {
	objc.Call[objc.Void](v_, objc.Sel("removeToolTip:"), tag)
}

// Adds a view to the view’s subviews so it’s displayed above its siblings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483783-addsubview?language=objc
func (v_ View) AddSubview(view IView) {
	objc.Call[objc.Void](v_, objc.Sel("addSubview:"), objc.Ptr(view))
}

// Sets the view to full screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483780-enterfullscreenmode?language=objc
func (v_ View) EnterFullScreenModeWithOptions(screen IScreen, options map[ViewFullScreenModeOptionKey]objc.IObject) bool {
	rv := objc.Call[bool](v_, objc.Sel("enterFullScreenMode:withOptions:"), objc.Ptr(screen), options)
	return rv
}

// Informs the view that its effective appearance changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2977088-viewdidchangeeffectiveappearance?language=objc
func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidChangeEffectiveAppearance"))
}

// Invalidates all cursor rectangles set up using addCursorRect:cursor:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483733-discardcursorrects?language=objc
func (v_ View) DiscardCursorRects() {
	objc.Call[objc.Void](v_, objc.Sel("discardCursorRects"))
}

// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483803-convertpointtobacking?language=objc
func (v_ View) ConvertPointToBacking(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("convertPointToBacking:"), point)
	return rv
}

// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483283-setframeorigin?language=objc
func (v_ View) SetFrameOrigin(newOrigin foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("setFrameOrigin:"), newOrigin)
}

// Instructs the view to exit full screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483256-exitfullscreenmodewithoptions?language=objc
func (v_ View) ExitFullScreenModeWithOptions(options map[ViewFullScreenModeOptionKey]objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("exitFullScreenModeWithOptions:"), options)
}

// Called after a contextual menu that was displayed from the receiving view has been closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483770-didclosemenu?language=objc
func (v_ View) DidCloseMenuWithEvent(menu IMenu, event IEvent) {
	objc.Call[objc.Void](v_, objc.Sel("didCloseMenu:withEvent:"), objc.Ptr(menu), objc.Ptr(event))
}

// Adds a constraint on the layout of the receiving view or its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526969-addconstraint?language=objc
func (v_ View) AddConstraint(constraint ILayoutConstraint) {
	objc.Call[objc.Void](v_, objc.Sel("addConstraint:"), objc.Ptr(constraint))
}

// Adds multiple constraints on the layout of the receiving view or its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526931-addconstraints?language=objc
func (v_ View) AddConstraints(constraints []ILayoutConstraint) {
	objc.Call[objc.Void](v_, objc.Sel("addConstraints:"), constraints)
}

// Establishes  an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483668-addtrackingrect?language=objc
func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer, flag bool) TrackingRectTag {
	rv := objc.Call[TrackingRectTag](v_, objc.Sel("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return rv
}

// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483227-convertsizetobacking?language=objc
func (v_ View) ConvertSizeToBacking(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("convertSizeToBacking:"), size)
	return rv
}

// Overridden by subclasses to perform additional actions before subviews are removed from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483624-willremovesubview?language=objc
func (v_ View) WillRemoveSubview(subview IView) {
	objc.Call[objc.Void](v_, objc.Sel("willRemoveSubview:"), objc.Ptr(subview))
}

// Overridden by subclasses to modify a given rectangle, returning the altered rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483616-adjustscroll?language=objc
func (v_ View) AdjustScroll(newVisible foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("adjustScroll:"), newVisible)
	return rv
}

// Attaches a gesture recognizer to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483749-addgesturerecognizer?language=objc
func (v_ View) AddGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.Call[objc.Void](v_, objc.Sel("addGestureRecognizer:"), objc.Ptr(gestureRecognizer))
}

// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483819-convertrectfrombacking?language=objc
func (v_ View) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}

// Returns a bitmap-representation object suitable for caching the specified portion of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483440-bitmapimagerepforcachingdisplayi?language=objc
func (v_ View) BitmapImageRepForCachingDisplayInRect(rect foundation.Rect) BitmapImageRep {
	rv := objc.Call[BitmapImageRep](v_, objc.Sel("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}

// Shows a window displaying the definition of the attributed string at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483747-showdefinitionforattributedstrin?language=objc
func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString foundation.IAttributedString, textBaselineOrigin foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("showDefinitionForAttributedString:atPoint:"), objc.Ptr(attrString), textBaselineOrigin)
}

// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483570-needstodrawrect?language=objc
func (v_ View) NeedsToDrawRect(rect foundation.Rect) bool {
	rv := objc.Call[bool](v_, objc.Sel("needsToDrawRect:"), rect)
	return rv
}

// Writes the end of a conforming page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483549-endpage?language=objc
func (v_ View) EndPage() {
	objc.Call[objc.Void](v_, objc.Sel("endPage"))
}

// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or nil if that point lies completely outside the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483364-hittest?language=objc
func (v_ View) HitTest(point foundation.Point) View {
	rv := objc.Call[View](v_, objc.Sel("hitTest:"), point)
	return rv
}

// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483292-drawpageborderwithsize?language=objc
func (v_ View) DrawPageBorderWithSize(borderSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("drawPageBorderWithSize:"), borderSize)
}

// Adds a given tracking area to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483489-addtrackingarea?language=objc
func (v_ View) AddTrackingArea(trackingArea ITrackingArea) {
	objc.Call[objc.Void](v_, objc.Sel("addTrackingArea:"), objc.Ptr(trackingArea))
}

// Scrolls the view’s closest ancestor NSClipView object the minimum distance needed so a specified region of the view becomes visible in the clip view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483811-scrollrecttovisible?language=objc
func (v_ View) ScrollRectToVisible(rect foundation.Rect) bool {
	rv := objc.Call[bool](v_, objc.Sel("scrollRectToVisible:"), rect)
	return rv
}

// Overridden by subclasses to define their default cursor rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483448-resetcursorrects?language=objc
func (v_ View) ResetCursorRects() {
	objc.Call[objc.Void](v_, objc.Sel("resetCursorRects"))
}

// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483244-shoulddelaywindoworderingforeven?language=objc
func (v_ View) ShouldDelayWindowOrderingForEvent(event IEvent) bool {
	rv := objc.Call[bool](v_, objc.Sel("shouldDelayWindowOrderingForEvent:"), objc.Ptr(event))
	return rv
}

// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483319-convertsizefrombacking?language=objc
func (v_ View) ConvertSizeFromBacking(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("convertSizeFromBacking:"), size)
	return rv
}

// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483499-writepdfinsiderect?language=objc
func (v_ View) WritePDFInsideRectToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	objc.Call[objc.Void](v_, objc.Sel("writePDFInsideRect:toPasteboard:"), rect, objc.Ptr(pasteboard))
}

// Updates the constraints for the receiving view and its subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526939-updateconstraintsforsubtreeifnee?language=objc
func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Call[objc.Void](v_, objc.Sel("updateConstraintsForSubtreeIfNeeded"))
}

// Convert the size from the view’s interior coordinate system to the layer's interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483315-convertpointtolayer?language=objc
func (v_ View) ConvertPointToLayer(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("convertPointToLayer:"), point)
	return rv
}

// Returns the appropriate rectangle to use when magnifying around the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483305-rectforsmartmagnificationatpoint?language=objc
func (v_ View) RectForSmartMagnificationAtPointInRect(location foundation.Point, visibleRect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}

// Scrolls the view’s closest ancestor NSClipView object so a point in the view lies at the origin of the clip view's bounds rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483394-scrollpoint?language=objc
func (v_ View) ScrollPoint(point foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("scrollPoint:"), point)
}

// Informs the view that its superview has changed (possibly to nil). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483568-viewdidmovetosuperview?language=objc
func (v_ View) ViewDidMoveToSuperview() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidMoveToSuperview"))
}

// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483231-menuforevent?language=objc
func (v_ View) MenuForEvent(event IEvent) Menu {
	rv := objc.Call[Menu](v_, objc.Sel("menuForEvent:"), objc.Ptr(event))
	return rv
}

// Returns EPS data that draws the region of the view within a specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483735-datawithepsinsiderect?language=objc
func (v_ View) DataWithEPSInsideRect(rect foundation.Rect) []byte {
	rv := objc.Call[[]byte](v_, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}

// Returns the closest ancestor shared by the view and another specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483353-ancestorsharedwithview?language=objc
func (v_ View) AncestorSharedWithView(view IView) View {
	rv := objc.Call[View](v_, objc.Sel("ancestorSharedWithView:"), objc.Ptr(view))
	return rv
}

// Updates the view’s content by modifying its underlying layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483580-updatelayer?language=objc
func (v_ View) UpdateLayer() {
	objc.Call[objc.Void](v_, objc.Sel("updateLayer"))
}

// Establishes  the cursor to be used when the mouse pointer lies within a specified region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483540-addcursorrect?language=objc
func (v_ View) AddCursorRectCursor(rect foundation.Rect, object ICursor) {
	objc.Call[objc.Void](v_, objc.Sel("addCursorRect:cursor:"), rect, objc.Ptr(object))
}

// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483265-removefromsuperview?language=objc
func (v_ View) RemoveFromSuperview() {
	objc.Call[objc.Void](v_, objc.Sel("removeFromSuperview"))
}

// This method is invoked at the end of the printing session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483610-enddocument?language=objc
func (v_ View) EndDocument() {
	objc.Call[objc.Void](v_, objc.Sel("endDocument"))
}

// Translates the display rectangles by the specified delta. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483731-translaterectsneedingdisplayinre?language=objc
func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect foundation.Rect, delta foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}

// Removes the provided layout guide from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1527086-removelayoutguide?language=objc
func (v_ View) RemoveLayoutGuide(guide ILayoutGuide) {
	objc.Call[objc.Void](v_, objc.Sel("removeLayoutGuide:"), objc.Ptr(guide))
}

// Acts as displayIfNeeded, confining drawing to a specified region of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483813-displayifneededinrect?language=objc
func (v_ View) DisplayIfNeededInRect(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("displayIfNeededInRect:"), rect)
}

// Scrolls the view’s closest ancestor NSClipView object proportionally to the distance of an event that occurs outside of it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483471-autoscroll?language=objc
func (v_ View) Autoscroll(event IEvent) bool {
	rv := objc.Call[bool](v_, objc.Sel("autoscroll:"), objc.Ptr(event))
	return rv
}

// Converts a rectangle from the coordinate system of another view to that of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483785-convertrect?language=objc
func (v_ View) ConvertRectFromView(rect foundation.Rect, view IView) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("convertRect:fromView:"), rect, objc.Ptr(view))
	return rv
}

// Overridden by subclasses to return YES if the view should be sent a mouseDown: message for an initial mouse-down event, NO if not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483410-acceptsfirstmouse?language=objc
func (v_ View) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Call[bool](v_, objc.Sel("acceptsFirstMouse:"), objc.Ptr(event))
	return rv
}

// Returns the constraints impacting the layout of the view for a given orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1525968-constraintsaffectinglayoutforori?language=objc
func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](v_, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}

// Removes the specified constraint from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1524333-removeconstraint?language=objc
func (v_ View) RemoveConstraint(constraint ILayoutConstraint) {
	objc.Call[objc.Void](v_, objc.Sel("removeConstraint:"), objc.Ptr(constraint))
}

// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483725-centerscanrect?language=objc
func (v_ View) CenterScanRect(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("centerScanRect:"), rect)
	return rv
}

// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1531337-scrollclipview?language=objc
func (v_ View) ScrollClipViewToPoint(clipView IClipView, point foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("scrollClipView:toPoint:"), objc.Ptr(clipView), point)
}

// Translates the view’s coordinate system so that its origin moves to a new location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483385-translateorigintopoint?language=objc
func (v_ View) TranslateOriginToPoint(translation foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("translateOriginToPoint:"), translation)
}

// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483699-displayrectignoringopacity?language=objc
func (v_ View) DisplayRectIgnoringOpacity(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("displayRectIgnoringOpacity:"), rect)
}

// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483648-convertrecttobacking?language=objc
func (v_ View) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("convertRectToBacking:"), rect)
	return rv
}

// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483275-viewdidunhide?language=objc
func (v_ View) ViewDidUnhide() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidUnhide"))
}

// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483719-updatetrackingareas?language=objc
func (v_ View) UpdateTrackingAreas() {
	objc.Call[objc.Void](v_, objc.Sel("updateTrackingAreas"))
}

// Responds when the view’s backing store properties change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483742-viewdidchangebackingproperties?language=objc
func (v_ View) ViewDidChangeBackingProperties() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidChangeBackingProperties"))
}

// Registers the pasteboard types that the view will accept as the destination of an image-dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483578-registerfordraggedtypes?language=objc
func (v_ View) RegisterForDraggedTypes(newTypes []PasteboardType) {
	objc.Call[objc.Void](v_, objc.Sel("registerForDraggedTypes:"), newTypes)
}

// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483596-viewdidhide?language=objc
func (v_ View) ViewDidHide() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidHide"))
}

// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483552-cachedisplayinrect?language=objc
func (v_ View) CacheDisplayInRectToBitmapImageRep(rect foundation.Rect, bitmapImageRep IBitmapImageRep) {
	objc.Call[objc.Void](v_, objc.Sel("cacheDisplayInRect:toBitmapImageRep:"), rect, objc.Ptr(bitmapImageRep))
}

// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483644-removefromsuperviewwithoutneedin?language=objc
func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Call[objc.Void](v_, objc.Sel("removeFromSuperviewWithoutNeedingDisplay"))
}

// Initiates a dragging session with a group of dragging items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483791-begindraggingsessionwithitems?language=objc
func (v_ View) BeginDraggingSessionWithItemsEventSource(items []IDraggingItem, event IEvent, source PDraggingSource) DraggingSession {
	po2 := objc.WrapAsProtocol("NSDraggingSource", source)
	rv := objc.Call[DraggingSession](v_, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, objc.Ptr(event), po2)
	return rv
}

// Initiates a dragging session with a group of dragging items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483791-begindraggingsessionwithitems?language=objc
func (v_ View) BeginDraggingSessionWithItemsEventSourceObject(items []IDraggingItem, event IEvent, sourceObject objc.IObject) DraggingSession {
	rv := objc.Call[DraggingSession](v_, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, objc.Ptr(event), objc.Ptr(sourceObject))
	return rv
}

// Called just before a contextual menu for a view is opened on screen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483429-willopenmenu?language=objc
func (v_ View) WillOpenMenuWithEvent(menu IMenu, event IEvent) {
	objc.Call[objc.Void](v_, objc.Sel("willOpenMenu:withEvent:"), objc.Ptr(menu), objc.Ptr(event))
}

// Displays the view and all its subviews if possible, invoking each of the NSView methods lockFocus, drawRect:, and unlockFocus as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483487-display?language=objc
func (v_ View) Display() {
	objc.Call[objc.Void](v_, objc.Sel("display"))
}

// Informs the view that the bounds size of its superview has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483697-resizewitholdsuperviewsize?language=objc
func (v_ View) ResizeWithOldSuperviewSize(oldSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("resizeWithOldSuperviewSize:"), oldSize)
}

// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483229-addtooltiprect?language=objc
func (v_ View) AddToolTipRectOwnerUserData(rect foundation.Rect, owner objc.IObject, data unsafe.Pointer) ToolTipTag {
	rv := objc.Call[ToolTipTag](v_, objc.Sel("addToolTipRect:owner:userData:"), rect, owner, data)
	return rv
}

// Sets the priority with which a view resists being made larger than its intrinsic size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526937-setcontenthuggingpriority?language=objc
func (v_ View) SetContentHuggingPriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Call[objc.Void](v_, objc.Sel("setContentHuggingPriority:forOrientation:"), priority, orientation)
}

// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1534216-reflectscrolledclipview?language=objc
func (v_ View) ReflectScrolledClipView(clipView IClipView) {
	objc.Call[objc.Void](v_, objc.Sel("reflectScrolledClipView:"), objc.Ptr(clipView))
}

// Invalidates the area around the focus ring. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483556-setkeyboardfocusringneedsdisplay?language=objc
func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}

// Update constraints for the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526891-updateconstraints?language=objc
func (v_ View) UpdateConstraints() {
	objc.Call[objc.Void](v_, objc.Sel("updateConstraints"))
}

// Removes the tracking rectangle identified by a tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483729-removetrackingrect?language=objc
func (v_ View) RemoveTrackingRect(tag TrackingRectTag) {
	objc.Call[objc.Void](v_, objc.Sel("removeTrackingRect:"), tag)
}

// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483530-setframesize?language=objc
func (v_ View) SetFrameSize(newSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("setFrameSize:"), newSize)
}

// Sets the origin of the view’s bounds rectangle to a specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483345-setboundsorigin?language=objc
func (v_ View) SetBoundsOrigin(newOrigin foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("setBoundsOrigin:"), newOrigin)
}

// Invoked by print: to determine the location of the region of the view being printed on the physical page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483223-locationofprintrect?language=objc
func (v_ View) LocationOfPrintRect(rect foundation.Rect) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("locationOfPrintRect:"), rect)
	return rv
}

// Convert the point from the layer's interior coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483554-convertpointfromlayer?language=objc
func (v_ View) ConvertPointFromLayer(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("convertPointFromLayer:"), point)
	return rv
}

// Convert the size from the layer's interior coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483479-convertsizefromlayer?language=objc
func (v_ View) ConvertSizeFromLayer(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("convertSizeFromLayer:"), size)
	return rv
}

// Replaces one of the view’s subviews with another view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483632-replacesubview?language=objc
func (v_ View) ReplaceSubviewWith(oldView IView, newView IView) {
	objc.Call[objc.Void](v_, objc.Sel("replaceSubview:with:"), objc.Ptr(oldView), objc.Ptr(newView))
}

// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483456-convertpointfrombacking?language=objc
func (v_ View) ConvertPointFromBacking(point foundation.Point) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("convertPointFromBacking:"), point)
	return rv
}

// Returns YES if the view handles page boundaries, NO otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483774-knowspagerange?language=objc
func (v_ View) KnowsPageRange(range_ foundation.RangePointer) bool {
	rv := objc.Call[bool](v_, objc.Sel("knowsPageRange:"), range_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1535261-rulerview?language=objc
func (v_ View) RulerViewLocationForPoint(ruler IRulerView, point foundation.Point) float64 {
	rv := objc.Call[float64](v_, objc.Sel("rulerView:locationForPoint:"), objc.Ptr(ruler), point)
	return rv
}

// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483705-print?language=objc
func (v_ View) Print(sender objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("print:"), sender)
}

// Returns the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526991-contentcompressionresistanceprio?language=objc
func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Call[LayoutPriority](v_, objc.Sel("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}

// Prepares the overdraw region for drawing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483427-preparecontentinrect?language=objc
func (v_ View) PrepareContentInRect(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("prepareContentInRect:"), rect)
}

// Removes the specified constraints from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526932-removeconstraints?language=objc
func (v_ View) RemoveConstraints(constraints []ILayoutConstraint) {
	objc.Call[objc.Void](v_, objc.Sel("removeConstraints:"), constraints)
}

// Invoked at the beginning of the printing session, this method sets up the current graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483423-begindocument?language=objc
func (v_ View) BeginDocument() {
	objc.Call[objc.Void](v_, objc.Sel("beginDocument"))
}

// Informs the view’s subviews that the view’s bounds rectangle size has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483495-resizesubviewswitholdsize?language=objc
func (v_ View) ResizeSubviewsWithOldSize(oldSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("resizeSubviewsWithOldSize:"), oldSize)
}

// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483399-setboundssize?language=objc
func (v_ View) SetBoundsSize(newSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("setBoundsSize:"), newSize)
}

// Overridden by subclasses to draw the view’s image within the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483686-drawrect?language=objc
func (v_ View) DrawRect(dirtyRect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("drawRect:"), dirtyRect)
}

// Draws the focus ring mask for the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483335-drawfocusringmask?language=objc
func (v_ View) DrawFocusRingMask() {
	objc.Call[objc.Void](v_, objc.Sel("drawFocusRingMask"))
}

// Convert the size from the view’s interior coordinate system to the layer's interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483776-convertrecttolayer?language=objc
func (v_ View) ConvertRectToLayer(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("convertRectToLayer:"), rect)
	return rv
}

// Scales the view’s coordinate system so that the unit square scales to the specified dimensions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483721-scaleunitsquaretosize?language=objc
func (v_ View) ScaleUnitSquareToSize(newUnitSize foundation.Size) {
	objc.Call[objc.Void](v_, objc.Sel("scaleUnitSquareToSize:"), newUnitSize)
}

// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483438-beginpageinrect?language=objc
func (v_ View) BeginPageInRectAtPlacement(rect foundation.Rect, location foundation.Point) {
	objc.Call[objc.Void](v_, objc.Sel("beginPageInRect:atPlacement:"), rect, location)
}

// Informs the view that its superview is about to change to the specified superview (which may be nil). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483545-viewwillmovetosuperview?language=objc
func (v_ View) ViewWillMoveToSuperview(newSuperview IView) {
	objc.Call[objc.Void](v_, objc.Sel("viewWillMoveToSuperview:"), objc.Ptr(newSuperview))
}

// Returns the priority with which a view resists being made larger than its intrinsic size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526979-contenthuggingpriorityfororienta?language=objc
func (v_ View) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Call[LayoutPriority](v_, objc.Sel("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}

// Perform layout in concert with the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526146-layout?language=objc
func (v_ View) Layout() {
	objc.Call[objc.Void](v_, objc.Sel("layout"))
}

// Restores the view to an initial state so that it can be reused. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483626-prepareforreuse?language=objc
func (v_ View) PrepareForReuse() {
	objc.Call[objc.Void](v_, objc.Sel("prepareForReuse"))
}

// Removes a given tracking area from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483634-removetrackingarea?language=objc
func (v_ View) RemoveTrackingArea(trackingArea ITrackingArea) {
	objc.Call[objc.Void](v_, objc.Sel("removeTrackingArea:"), objc.Ptr(trackingArea))
}

// Removes all tooltips assigned to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483801-removealltooltips?language=objc
func (v_ View) RemoveAllToolTips() {
	objc.Call[objc.Void](v_, objc.Sel("removeAllToolTips"))
}

// Returns PDF data that draws the region of the view within a specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483797-datawithpdfinsiderect?language=objc
func (v_ View) DataWithPDFInsideRect(rect foundation.Rect) []byte {
	rv := objc.Call[[]byte](v_, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}

// Invalidates the view’s intrinsic content size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526864-invalidateintrinsiccontentsize?language=objc
func (v_ View) InvalidateIntrinsicContentSize() {
	objc.Call[objc.Void](v_, objc.Sel("invalidateIntrinsicContentSize"))
}

// Sets the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1524974-setcontentcompressionresistancep?language=objc
func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Call[objc.Void](v_, objc.Sel("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}

// Orders the view's immediate subviews using the specified comparator function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483805-sortsubviewsusingfunction?language=objc
func (v_ View) SortSubviewsUsingFunctionContext(compare func(arg0 View, arg1 View, arg2 unsafe.Pointer) foundation.ComparisonResult, context unsafe.Pointer) {
	objc.Call[objc.Void](v_, objc.Sel("sortSubviewsUsingFunction:context:"), compare, context)
}

// Informs the view of the end of a live resize—the user has finished resizing the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483543-viewdidendliveresize?language=objc
func (v_ View) ViewDidEndLiveResize() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidEndLiveResize"))
}

// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be nil). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483415-viewwillmovetowindow?language=objc
func (v_ View) ViewWillMoveToWindow(newWindow IWindow) {
	objc.Call[objc.Void](v_, objc.Sel("viewWillMoveToWindow:"), objc.Ptr(newWindow))
}

// Acts as displayIfNeeded, except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483526-displayifneededignoringopacity?language=objc
func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.Call[objc.Void](v_, objc.Sel("displayIfNeededIgnoringOpacity"))
}

// Returns the view’s alignment rectangle for a given frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526905-alignmentrectforframe?language=objc
func (v_ View) AlignmentRectForFrame(frame foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("alignmentRectForFrame:"), frame)
	return rv
}

// Unregisters the view as a possible destination in a dragging session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483602-unregisterdraggedtypes?language=objc
func (v_ View) UnregisterDraggedTypes() {
	objc.Call[objc.Void](v_, objc.Sel("unregisterDraggedTypes"))
}

// Detaches a gesture recognizer from the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483789-removegesturerecognizer?language=objc
func (v_ View) RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.Call[objc.Void](v_, objc.Sel("removeGestureRecognizer:"), objc.Ptr(gestureRecognizer))
}

// Overridden by subclasses to adjust page width during automatic pagination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483622-adjustpagewidthnew?language=objc
func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight *float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Call[objc.Void](v_, objc.Sel("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}

// Informs the view of the start of a live resize—the user has started resizing the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483620-viewwillstartliveresize?language=objc
func (v_ View) ViewWillStartLiveResize() {
	objc.Call[objc.Void](v_, objc.Sel("viewWillStartLiveResize"))
}

// Implemented by subclasses to determine the portion of the view to be printed for the specified page number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483252-rectforpage?language=objc
func (v_ View) RectForPage(page int) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("rectForPage:"), page)
	return rv
}

// Informs the view that it’s required to draw content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483351-viewwilldraw?language=objc
func (v_ View) ViewWillDraw() {
	objc.Call[objc.Void](v_, objc.Sel("viewWillDraw"))
}

// Adds the provided layout guide to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1530406-addlayoutguide?language=objc
func (v_ View) AddLayoutGuide(guide ILayoutGuide) {
	objc.Call[objc.Void](v_, objc.Sel("addLayoutGuide:"), objc.Ptr(guide))
}

// Invoked to notify the view that the focus ring mask requires updating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483809-notefocusringmaskchanged?language=objc
func (v_ View) NoteFocusRingMaskChanged() {
	objc.Call[objc.Void](v_, objc.Sel("noteFocusRingMaskChanged"))
}

// Returns the view’s nearest descendant (including itself) with a specific tag, or nil if no subview has that tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483294-viewwithtag?language=objc
func (v_ View) ViewWithTag(tag int) View {
	rv := objc.Call[View](v_, objc.Sel("viewWithTag:"), tag)
	return rv
}

// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483235-writeepsinsiderect?language=objc
func (v_ View) WriteEPSInsideRectToPasteboard(rect foundation.Rect, pasteboard IPasteboard) {
	objc.Call[objc.Void](v_, objc.Sel("writeEPSInsideRect:toPasteboard:"), rect, objc.Ptr(pasteboard))
}

// Randomly changes the frame of a view with an ambiguous layout between the different valid values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526934-exerciseambiguityinlayout?language=objc
func (v_ View) ExerciseAmbiguityInLayout() {
	objc.Call[objc.Void](v_, objc.Sel("exerciseAmbiguityInLayout"))
}

// Informs the view that it has been added to a new view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483329-viewdidmovetowindow?language=objc
func (v_ View) ViewDidMoveToWindow() {
	objc.Call[objc.Void](v_, objc.Sel("viewDidMoveToWindow"))
}

// Updates the layout of the receiving view and its subviews based on the current views and constraints. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526871-layoutsubtreeifneeded?language=objc
func (v_ View) LayoutSubtreeIfNeeded() {
	objc.Call[objc.Void](v_, objc.Sel("layoutSubtreeIfNeeded"))
}

// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in drawRect:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483772-getrectsbeingdrawn?language=objc
func (v_ View) GetRectsBeingDrawnCount(rects *foundation.Rect, count *int) {
	objc.Call[objc.Void](v_, objc.Sel("getRectsBeingDrawn:count:"), rects, count)
}

// Convert the size from the view’s interior coordinate system to the layer's interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483701-convertsizetolayer?language=objc
func (v_ View) ConvertSizeToLayer(size foundation.Size) foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("convertSizeToLayer:"), size)
	return rv
}

// Convert the rectangle from the layer's interior coordinate system to the view’s interior coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483404-convertrectfromlayer?language=objc
func (v_ View) ConvertRectFromLayer(rect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("convertRectFromLayer:"), rect)
	return rv
}

// Acts as display, but confining drawing to a rectangular region of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483518-displayrect?language=objc
func (v_ View) DisplayRect(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("displayRect:"), rect)
}

// Acts as displayIfNeeded, but confining drawing to aRect and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483481-displayifneededinrectignoringopa?language=objc
func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("displayIfNeededInRectIgnoringOpacity:"), rect)
}

// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483475-setneedsdisplayinrect?language=objc
func (v_ View) SetNeedsDisplayInRect(invalidRect foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("setNeedsDisplayInRect:"), invalidRect)
}

// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483237-mouse?language=objc
func (v_ View) MouseInRect(point foundation.Point, rect foundation.Rect) bool {
	rv := objc.Call[bool](v_, objc.Sel("mouse:inRect:"), point, rect)
	return rv
}

// Returns a backing store pixel-aligned rectangle in local view coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483321-backingalignedrect?language=objc
func (v_ View) BackingAlignedRectOptions(rect foundation.Rect, options foundation.AlignmentOptions) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}

// Overridden by subclasses to perform additional actions when subviews are added to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483454-didaddsubview?language=objc
func (v_ View) DidAddSubview(subview IView) {
	objc.Call[objc.Void](v_, objc.Sel("didAddSubview:"), objc.Ptr(subview))
}

// Returns the view’s frame for a given alignment rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1525584-frameforalignmentrect?language=objc
func (v_ View) FrameForAlignmentRect(alignmentRect foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("frameForAlignmentRect:"), alignmentRect)
	return rv
}

// Converts a size from the view’s coordinate system to that of another view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483744-convertsize?language=objc
func (v_ View) ConvertSizeToView(size foundation.Size, view IView) foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("convertSize:toView:"), size, objc.Ptr(view))
	return rv
}

// Returns a Boolean value that indicates whether the view is a subview of the specified view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483219-isdescendantof?language=objc
func (v_ View) IsDescendantOf(view IView) bool {
	rv := objc.Call[bool](v_, objc.Sel("isDescendantOf:"), objc.Ptr(view))
	return rv
}

// Displays the view and all its subviews if any part of the view has been marked as needing display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483566-displayifneeded?language=objc
func (v_ View) DisplayIfNeeded() {
	objc.Call[objc.Void](v_, objc.Sel("displayIfNeeded"))
}

// Creates the view’s backing layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483687-makebackinglayer?language=objc
func (v_ View) MakeBackingLayer() quartzcore.Layer {
	rv := objc.Call[quartzcore.Layer](v_, objc.Sel("makeBackingLayer"))
	return rv
}

// Converts a point from the view’s coordinate system to that of a given view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483406-convertpoint?language=objc
func (v_ View) ConvertPointToView(point foundation.Point, view IView) foundation.Point {
	rv := objc.Call[foundation.Point](v_, objc.Sel("convertPoint:toView:"), point, objc.Ptr(view))
	return rv
}

// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc
func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := objc.Call[bool](v_, objc.Sel("layerUsesCoreImageFilters"))
	return rv
}

// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483576-layerusescoreimagefilters?language=objc
func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setLayerUsesCoreImageFilters:"), value)
}

// A layout anchor representing the top edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526859-topanchor?language=objc
func (v_ View) TopAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](v_, objc.Sel("topAnchor"))
	return rv
}

// The contents redraw policy for the view’s layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483514-layercontentsredrawpolicy?language=objc
func (v_ View) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	rv := objc.Call[ViewLayerContentsRedrawPolicy](v_, objc.Sel("layerContentsRedrawPolicy"))
	return rv
}

// The contents redraw policy for the view’s layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483514-layercontentsredrawpolicy?language=objc
func (v_ View) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	objc.Call[objc.Void](v_, objc.Sel("setLayerContentsRedrawPolicy:"), value)
}

// The Core Animation layer that the view uses as its backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc
func (v_ View) Layer() quartzcore.Layer {
	rv := objc.Call[quartzcore.Layer](v_, objc.Sel("layer"))
	return rv
}

// The Core Animation layer that the view uses as its backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483298-layer?language=objc
func (v_ View) SetLayer(value quartzcore.ILayer) {
	objc.Call[objc.Void](v_, objc.Sel("setLayer:"), objc.Ptr(value))
}

// The view’s print job title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483753-printjobtitle?language=objc
func (v_ View) PrintJobTitle() string {
	rv := objc.Call[string](v_, objc.Sel("printJobTitle"))
	return rv
}

// Overridden by subclasses to return the default pop-up menu for instances of the receiving class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483417-defaultmenu?language=objc
func (vc _ViewClass) DefaultMenu() Menu {
	rv := objc.Call[Menu](vc, objc.Sel("defaultMenu"))
	return rv
}

// Overridden by subclasses to return the default pop-up menu for instances of the receiving class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483417-defaultmenu?language=objc
func View_DefaultMenu() Menu {
	return ViewClass.DefaultMenu()
}

// A layout anchor representing the left edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526951-leftanchor?language=objc
func (v_ View) LeftAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](v_, objc.Sel("leftAnchor"))
	return rv
}

// The portion of the view that is not clipped by its superviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483446-visiblerect?language=objc
func (v_ View) VisibleRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("visibleRect"))
	return rv
}

// The text input context object for the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483323-inputcontext?language=objc
func (v_ View) InputContext() TextInputContext {
	rv := objc.Call[TextInputContext](v_, objc.Sel("inputContext"))
	return rv
}

// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483390-rotatedorscaledfrombase?language=objc
func (v_ View) IsRotatedOrScaledFromBase() bool {
	rv := objc.Call[bool](v_, objc.Sel("isRotatedOrScaledFromBase"))
	return rv
}

// A Boolean value indicating whether the view is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc
func (v_ View) IsHidden() bool {
	rv := objc.Call[bool](v_, objc.Sel("isHidden"))
	return rv
}

// A Boolean value indicating whether the view is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483369-hidden?language=objc
func (v_ View) SetHidden(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setHidden:"), value)
}

// A Boolean value indicating whether the view needs a layout pass before it can be drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc
func (v_ View) NeedsLayout() bool {
	rv := objc.Call[bool](v_, objc.Sel("needsLayout"))
	return rv
}

// A Boolean value indicating whether the view needs a layout pass before it can be drawn. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526912-needslayout?language=objc
func (v_ View) SetNeedsLayout(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setNeedsLayout:"), value)
}

// The view object that follows the current view in the key view loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc
func (v_ View) NextKeyView() View {
	rv := objc.Call[View](v_, objc.Sel("nextKeyView"))
	return rv
}

// The view object that follows the current view in the key view loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483465-nextkeyview?language=objc
func (v_ View) SetNextKeyView(value IView) {
	objc.Call[objc.Void](v_, objc.Sel("setNextKeyView:"), objc.Ptr(value))
}

// The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526963-firstbaselineoffsetfromtop?language=objc
func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := objc.Call[float64](v_, objc.Sel("firstBaselineOffsetFromTop"))
	return rv
}

// The minimum size of the view that satisfies the constraints it holds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526904-fittingsize?language=objc
func (v_ View) FittingSize() foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("fittingSize"))
	return rv
}

// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc
func (v_ View) BoundsRotation() float64 {
	rv := objc.Call[float64](v_, objc.Sel("boundsRotation"))
	return rv
}

// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483746-boundsrotation?language=objc
func (v_ View) SetBoundsRotation(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setBoundsRotation:"), value)
}

// The insets (in points) from the view’s frame that define its content rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526870-alignmentrectinsets?language=objc
func (v_ View) AlignmentRectInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](v_, objc.Sel("alignmentRectInsets"))
	return rv
}

// A default footer string that includes the current page number and page count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483355-pagefooter?language=objc
func (v_ View) PageFooter() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](v_, objc.Sel("pageFooter"))
	return rv
}

// A Boolean value indicating whether the view posts notifications when its frame rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc
func (v_ View) PostsFrameChangedNotifications() bool {
	rv := objc.Call[bool](v_, objc.Sel("postsFrameChangedNotifications"))
	return rv
}

// A Boolean value indicating whether the view posts notifications when its frame rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483524-postsframechangednotifications?language=objc
func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setPostsFrameChangedNotifications:"), value)
}

// A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483512-needspaneltobecomekey?language=objc
func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := objc.Call[bool](v_, objc.Sel("needsPanelToBecomeKey"))
	return rv
}

// An array of Core Image filters to apply to the contents of the view and its sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc
func (v_ View) ContentFilters() []coreimage.Filter {
	rv := objc.Call[[]coreimage.Filter](v_, objc.Sel("contentFilters"))
	return rv
}

// An array of Core Image filters to apply to the contents of the view and its sublayers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483703-contentfilters?language=objc
func (v_ View) SetContentFilters(value []coreimage.IFilter) {
	objc.Call[objc.Void](v_, objc.Sel("setContentFilters:"), value)
}

// The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483572-nextvalidkeyview?language=objc
func (v_ View) NextValidKeyView() View {
	rv := objc.Call[View](v_, objc.Sel("nextValidKeyView"))
	return rv
}

// The text for the view’s tooltip. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc
func (v_ View) ToolTip() string {
	rv := objc.Call[string](v_, objc.Sel("toolTip"))
	return rv
}

// The text for the view’s tooltip. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483541-tooltip?language=objc
func (v_ View) SetToolTip(value string) {
	objc.Call[objc.Void](v_, objc.Sel("setToolTip:"), value)
}

// The rotation angle of the view around the center of its layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc
func (v_ View) FrameCenterRotation() float64 {
	rv := objc.Call[float64](v_, objc.Sel("frameCenterRotation"))
	return rv
}

// The rotation angle of the view around the center of its layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483367-framecenterrotation?language=objc
func (v_ View) SetFrameCenterRotation(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setFrameCenterRotation:"), value)
}

// A layout anchor representing the vertical center of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526935-centeryanchor?language=objc
func (v_ View) CenterYAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](v_, objc.Sel("centerYAnchor"))
	return rv
}

// A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483795-preservescontentduringliveresize?language=objc
func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := objc.Call[bool](v_, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}

// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc
func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.Call[bool](v_, objc.Sel("translatesAutoresizingMaskIntoConstraints"))
	return rv
}

// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526961-translatesautoresizingmaskintoco?language=objc
func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}

// A layout anchor representing the baseline for the topmost line of text in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526900-firstbaselineanchor?language=objc
func (v_ View) FirstBaselineAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](v_, objc.Sel("firstBaselineAnchor"))
	return rv
}

// The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1525942-lastbaselineoffsetfrombottom?language=objc
func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := objc.Call[float64](v_, objc.Sel("lastBaselineOffsetFromBottom"))
	return rv
}

// A Boolean value that indicates whether views support responsive scrolling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2870005-compatiblewithresponsivescrollin?language=objc
func (vc _ViewClass) CompatibleWithResponsiveScrolling() bool {
	rv := objc.Call[bool](vc, objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}

// A Boolean value that indicates whether views support responsive scrolling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2870005-compatiblewithresponsivescrollin?language=objc
func View_CompatibleWithResponsiveScrolling() bool {
	return ViewClass.CompatibleWithResponsiveScrolling()
}

// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1426890-pressureconfiguration?language=objc
func (v_ View) PressureConfiguration() PressureConfiguration {
	rv := objc.Call[PressureConfiguration](v_, objc.Sel("pressureConfiguration"))
	return rv
}

// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1426890-pressureconfiguration?language=objc
func (v_ View) SetPressureConfiguration(value IPressureConfiguration) {
	objc.Call[objc.Void](v_, objc.Sel("setPressureConfiguration:"), objc.Ptr(value))
}

// A Boolean value indicating which drawing path the view takes when updating its contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483461-wantsupdatelayer?language=objc
func (v_ View) WantsUpdateLayer() bool {
	rv := objc.Call[bool](v_, objc.Sel("wantsUpdateLayer"))
	return rv
}

// The type of focus ring drawn around the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483261-focusringtype?language=objc
func (v_ View) FocusRingType() FocusRingType {
	rv := objc.Call[FocusRingType](v_, objc.Sel("focusRingType"))
	return rv
}

// The type of focus ring drawn around the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483261-focusringtype?language=objc
func (v_ View) SetFocusRingType(value FocusRingType) {
	objc.Call[objc.Void](v_, objc.Sel("setFocusRingType:"), value)
}

// An array of the view’s tracking areas. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483333-trackingareas?language=objc
func (v_ View) TrackingAreas() []TrackingArea {
	rv := objc.Call[[]TrackingArea](v_, objc.Sel("trackingAreas"))
	return rv
}

// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc
func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := objc.Call[bool](v_, objc.Sel("postsBoundsChangedNotifications"))
	return rv
}

// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483239-postsboundschangednotifications?language=objc
func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setPostsBoundsChangedNotifications:"), value)
}

// The gesture recognize objects currently attached to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc
func (v_ View) GestureRecognizers() []GestureRecognizer {
	rv := objc.Call[[]GestureRecognizer](v_, objc.Sel("gestureRecognizers"))
	return rv
}

// The gesture recognize objects currently attached to the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483658-gesturerecognizers?language=objc
func (v_ View) SetGestureRecognizers(value []IGestureRecognizer) {
	objc.Call[objc.Void](v_, objc.Sel("setGestureRecognizers:"), value)
}

// The currently focused view object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483662-focusview?language=objc
func (vc _ViewClass) FocusView() View {
	rv := objc.Call[View](vc, objc.Sel("focusView"))
	return rv
}

// The currently focused view object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483662-focusview?language=objc
func View_FocusView() View {
	return ViewClass.FocusView()
}

// The nearest ancestor scroll view that contains the current view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483654-enclosingscrollview?language=objc
func (v_ View) EnclosingScrollView() ScrollView {
	rv := objc.Call[ScrollView](v_, objc.Sel("enclosingScrollView"))
	return rv
}

// The opacity of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc
func (v_ View) AlphaValue() float64 {
	rv := objc.Call[float64](v_, objc.Sel("alphaValue"))
	return rv
}

// The opacity of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483560-alphavalue?language=objc
func (v_ View) SetAlphaValue(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setAlphaValue:"), value)
}

// A default header string that includes the print job title and date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483674-pageheader?language=objc
func (v_ View) PageHeader() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](v_, objc.Sel("pageHeader"))
	return rv
}

// The layout guide you use to position content inside your view’s safe area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3553228-safearealayoutguide?language=objc
func (v_ View) SafeAreaLayoutGuide() LayoutGuide {
	rv := objc.Call[LayoutGuide](v_, objc.Sel("safeAreaLayoutGuide"))
	return rv
}

// A layout anchor representing the right edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1524466-rightanchor?language=objc
func (v_ View) RightAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](v_, objc.Sel("rightAnchor"))
	return rv
}

// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483709-rotatedfrombase?language=objc
func (v_ View) IsRotatedFromBase() bool {
	rv := objc.Call[bool](v_, objc.Sel("isRotatedFromBase"))
	return rv
}

// The portion of the view that has been rendered and is available for responsive scrolling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc
func (v_ View) PreparedContentRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("preparedContentRect"))
	return rv
}

// The portion of the view that has been rendered and is available for responsive scrolling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483215-preparedcontentrect?language=objc
func (v_ View) SetPreparedContentRect(value foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("setPreparedContentRect:"), value)
}

// The view’s bounds rectangle, which expresses its location and size in its own coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc
func (v_ View) Bounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("bounds"))
	return rv
}

// The view’s bounds rectangle, which expresses its location and size in its own coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483817-bounds?language=objc
func (v_ View) SetBounds(value foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("setBounds:"), value)
}

// The distance (in points) between the bottom of the view’s alignment rectangle and its baseline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526949-baselineoffsetfrombottom?language=objc
func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := objc.Call[float64](v_, objc.Sel("baselineOffsetFromBottom"))
	return rv
}

// The distances from the edges of your view that define the safe area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3553227-safeareainsets?language=objc
func (v_ View) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](v_, objc.Sel("safeAreaInsets"))
	return rv
}

// A Boolean value indicating whether the view is being rendered as part of a live resizing operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483267-inliveresize?language=objc
func (v_ View) InLiveResize() bool {
	rv := objc.Call[bool](v_, objc.Sel("inLiveResize"))
	return rv
}

// A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483317-drawingfindindicator?language=objc
func (v_ View) IsDrawingFindIndicator() bool {
	rv := objc.Call[bool](v_, objc.Sel("isDrawingFindIndicator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2544729-candidatelisttouchbaritem?language=objc
func (v_ View) CandidateListTouchBarItem() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](v_, objc.Sel("candidateListTouchBarItem"))
	return rv
}

// A layout guide that provides the recommended amount of padding for content inside of a view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3622483-layoutmarginsguide?language=objc
func (v_ View) LayoutMarginsGuide() LayoutGuide {
	rv := objc.Call[LayoutGuide](v_, objc.Sel("layoutMarginsGuide"))
	return rv
}

// The layout direction for content in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483254-userinterfacelayoutdirection?language=objc
func (v_ View) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Call[UserInterfaceLayoutDirection](v_, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}

// The layout direction for content in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483254-userinterfacelayoutdirection?language=objc
func (v_ View) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Call[objc.Void](v_, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// A layout anchor representing the bottom edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526974-bottomanchor?language=objc
func (v_ View) BottomAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](v_, objc.Sel("bottomAnchor"))
	return rv
}

// The view’s closest opaque ancestor, which might be the view itself. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483383-opaqueancestor?language=objc
func (v_ View) OpaqueAncestor() View {
	rv := objc.Call[View](v_, objc.Sel("opaqueAncestor"))
	return rv
}

// The current layer contents placement policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483375-layercontentsplacement?language=objc
func (v_ View) LayerContentsPlacement() ViewLayerContentsPlacement {
	rv := objc.Call[ViewLayerContentsPlacement](v_, objc.Sel("layerContentsPlacement"))
	return rv
}

// The current layer contents placement policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483375-layercontentsplacement?language=objc
func (v_ View) SetLayerContentsPlacement(value ViewLayerContentsPlacement) {
	objc.Call[objc.Void](v_, objc.Sel("setLayerContentsPlacement:"), value)
}

// A Boolean value indicating whether the view fills its frame rectangle with opaque content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483558-opaque?language=objc
func (v_ View) IsOpaque() bool {
	rv := objc.Call[bool](v_, objc.Sel("isOpaque"))
	return rv
}

// The closest view object in the key view loop that precedes the current view and accepts first responder status. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483371-previousvalidkeyview?language=objc
func (v_ View) PreviousValidKeyView() View {
	rv := objc.Call[View](v_, objc.Sel("previousValidKeyView"))
	return rv
}

// The menu item containing the view or any of its superviews in the view hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1514865-enclosingmenuitem?language=objc
func (v_ View) EnclosingMenuItem() MenuItem {
	rv := objc.Call[MenuItem](v_, objc.Sel("enclosingMenuItem"))
	return rv
}

// A Boolean value indicating whether the view’s constraints need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc
func (v_ View) NeedsUpdateConstraints() bool {
	rv := objc.Call[bool](v_, objc.Sel("needsUpdateConstraints"))
	return rv
}

// A Boolean value indicating whether the view’s constraints need to be updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526856-needsupdateconstraints?language=objc
func (v_ View) SetNeedsUpdateConstraints(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setNeedsUpdateConstraints:"), value)
}

// A layout anchor representing the height of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526942-heightanchor?language=objc
func (v_ View) HeightAnchor() LayoutDimension {
	rv := objc.Call[LayoutDimension](v_, objc.Sel("heightAnchor"))
	return rv
}

// The array of layout guide objects owned by this view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1534395-layoutguides?language=objc
func (v_ View) LayoutGuides() []LayoutGuide {
	rv := objc.Call[[]LayoutGuide](v_, objc.Sel("layoutGuides"))
	return rv
}

// A Boolean value indicating whether AppKit’s default clipping behavior is in effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483365-wantsdefaultclipping?language=objc
func (v_ View) WantsDefaultClipping() bool {
	rv := objc.Call[bool](v_, objc.Sel("wantsDefaultClipping"))
	return rv
}

// A rectangle in the view’s coordinate system that contains the unobscured portion of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3553229-safearearect?language=objc
func (v_ View) SafeAreaRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("safeAreaRect"))
	return rv
}

// Returns the default focus ring type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483589-defaultfocusringtype?language=objc
func (vc _ViewClass) DefaultFocusRingType() FocusRingType {
	rv := objc.Call[FocusRingType](vc, objc.Sel("defaultFocusRingType"))
	return rv
}

// Returns the default focus ring type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483589-defaultfocusringtype?language=objc
func View_DefaultFocusRingType() FocusRingType {
	return ViewClass.DefaultFocusRingType()
}

// A Boolean value indicating whether the view uses a layer as its backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc
func (v_ View) WantsLayer() bool {
	rv := objc.Call[bool](v_, objc.Sel("wantsLayer"))
	return rv
}

// A Boolean value indicating whether the view uses a layer as its backing store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483695-wantslayer?language=objc
func (v_ View) SetWantsLayer(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setWantsLayer:"), value)
}

// The natural size for the receiving view, considering only properties of the view itself. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526996-intrinsiccontentsize?language=objc
func (v_ View) IntrinsicContentSize() foundation.Size {
	rv := objc.Call[foundation.Size](v_, objc.Sel("intrinsicContentSize"))
	return rv
}

// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc
func (v_ View) AutoresizesSubviews() bool {
	rv := objc.Call[bool](v_, objc.Sel("autoresizesSubviews"))
	return rv
}

// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483358-autoresizessubviews?language=objc
func (v_ View) SetAutoresizesSubviews(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setAutoresizesSubviews:"), value)
}

// A layout anchor representing the trailing edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526983-trailinganchor?language=objc
func (v_ View) TrailingAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](v_, objc.Sel("trailingAnchor"))
	return rv
}

// A Boolean value indicating whether the view incorporates content from its subviews into its own layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc
func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := objc.Call[bool](v_, objc.Sel("canDrawSubviewsIntoLayer"))
	return rv
}

// A Boolean value indicating whether the view incorporates content from its subviews into its own layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483347-candrawsubviewsintolayer?language=objc
func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setCanDrawSubviewsIntoLayer:"), value)
}

// The view object preceding the current view in the key view loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483646-previouskeyview?language=objc
func (v_ View) PreviousKeyView() View {
	rv := objc.Call[View](v_, objc.Sel("previousKeyView"))
	return rv
}

// A layout anchor representing the leading edge of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1524264-leadinganchor?language=objc
func (v_ View) LeadingAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](v_, objc.Sel("leadingAnchor"))
	return rv
}

// The types of touch interactions the view allows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2544839-allowedtouchtypes?language=objc
func (v_ View) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Call[TouchTypeMask](v_, objc.Sel("allowedTouchTypes"))
	return rv
}

// The types of touch interactions the view allows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/2544839-allowedtouchtypes?language=objc
func (v_ View) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Call[objc.Void](v_, objc.Sel("setAllowedTouchTypes:"), value)
}

// The focus ring mask bounds, specified in the view’s coordinate space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483287-focusringmaskbounds?language=objc
func (v_ View) FocusRingMaskBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("focusRingMaskBounds"))
	return rv
}

// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc
func (v_ View) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("frame"))
	return rv
}

// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483713-frame?language=objc
func (v_ View) SetFrame(value foundation.Rect) {
	objc.Call[objc.Void](v_, objc.Sel("setFrame:"), value)
}

// A Boolean value indicating whether the view can become key view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483759-canbecomekeyview?language=objc
func (v_ View) CanBecomeKeyView() bool {
	rv := objc.Call[bool](v_, objc.Sel("canBecomeKeyView"))
	return rv
}

// A Boolean value indicating whether the view is in full screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483337-infullscreenmode?language=objc
func (v_ View) IsInFullScreenMode() bool {
	rv := objc.Call[bool](v_, objc.Sel("isInFullScreenMode"))
	return rv
}

// A Boolean value indicating whether the view can draw its contents on a background thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc
func (v_ View) CanDrawConcurrently() bool {
	rv := objc.Call[bool](v_, objc.Sel("canDrawConcurrently"))
	return rv
}

// A Boolean value indicating whether the view can draw its contents on a background thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483425-candrawconcurrently?language=objc
func (v_ View) SetCanDrawConcurrently(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setCanDrawConcurrently:"), value)
}

// The view’s tag, which is an integer that you use to identify the view within your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483248-tag?language=objc
func (v_ View) Tag() int {
	rv := objc.Call[int](v_, objc.Sel("tag"))
	return rv
}

// The shadow displayed underneath the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483263-shadow?language=objc
func (v_ View) Shadow() Shadow {
	rv := objc.Call[Shadow](v_, objc.Sel("shadow"))
	return rv
}

// The shadow displayed underneath the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483263-shadow?language=objc
func (v_ View) SetShadow(value IShadow) {
	objc.Call[objc.Void](v_, objc.Sel("setShadow:"), objc.Ptr(value))
}

// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483392-widthadjustlimit?language=objc
func (v_ View) WidthAdjustLimit() float64 {
	rv := objc.Call[float64](v_, objc.Sel("widthAdjustLimit"))
	return rv
}

// Custom insets that you specify to modify your view’s safe area [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3553226-additionalsafeareainsets?language=objc
func (v_ View) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](v_, objc.Sel("additionalSafeAreaInsets"))
	return rv
}

// Custom insets that you specify to modify your view’s safe area [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3553226-additionalsafeareainsets?language=objc
func (v_ View) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](v_, objc.Sel("setAdditionalSafeAreaInsets:"), value)
}

// The Core Image filter used to composite the view’s contents with its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483516-compositingfilter?language=objc
func (v_ View) CompositingFilter() coreimage.Filter {
	rv := objc.Call[coreimage.Filter](v_, objc.Sel("compositingFilter"))
	return rv
}

// The Core Image filter used to composite the view’s contents with its background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483516-compositingfilter?language=objc
func (v_ View) SetCompositingFilter(value coreimage.IFilter) {
	objc.Call[objc.Void](v_, objc.Sel("setCompositingFilter:"), objc.Ptr(value))
}

// The array of pasteboard drag types that the view can accept. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483564-registereddraggedtypes?language=objc
func (v_ View) RegisteredDraggedTypes() []PasteboardType {
	rv := objc.Call[[]PasteboardType](v_, objc.Sel("registeredDraggedTypes"))
	return rv
}

// A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483473-hiddenorhashiddenancestor?language=objc
func (v_ View) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Call[bool](v_, objc.Sel("isHiddenOrHasHiddenAncestor"))
	return rv
}

// The view’s window object, if it is installed in a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483301-window?language=objc
func (v_ View) Window() Window {
	rv := objc.Call[Window](v_, objc.Sel("window"))
	return rv
}

// A Boolean value that indicates whether the view’s vertical size constraints are active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc
func (v_ View) IsVerticalContentSizeConstraintActive() bool {
	rv := objc.Call[bool](v_, objc.Sel("isVerticalContentSizeConstraintActive"))
	return rv
}

// A Boolean value that indicates whether the view’s vertical size constraints are active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3353054-verticalcontentsizeconstraintact?language=objc
func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setVerticalContentSizeConstraintActive:"), value)
}

// The options that determine how the view is resized relative to its superview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483281-autoresizingmask?language=objc
func (v_ View) AutoresizingMask() AutoresizingMaskOptions {
	rv := objc.Call[AutoresizingMaskOptions](v_, objc.Sel("autoresizingMask"))
	return rv
}

// The options that determine how the view is resized relative to its superview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483281-autoresizingmask?language=objc
func (v_ View) SetAutoresizingMask(value AutoresizingMaskOptions) {
	objc.Call[objc.Void](v_, objc.Sel("setAutoresizingMask:"), value)
}

// An array of Core Image filters to apply to the view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc
func (v_ View) BackgroundFilters() []coreimage.Filter {
	rv := objc.Call[[]coreimage.Filter](v_, objc.Sel("backgroundFilters"))
	return rv
}

// An array of Core Image filters to apply to the view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483689-backgroundfilters?language=objc
func (v_ View) SetBackgroundFilters(value []coreimage.IFilter) {
	objc.Call[objc.Void](v_, objc.Sel("setBackgroundFilters:"), value)
}

// A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526907-hasambiguouslayout?language=objc
func (v_ View) HasAmbiguousLayout() bool {
	rv := objc.Call[bool](v_, objc.Sel("hasAmbiguousLayout"))
	return rv
}

// Returns the constraints held by the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526917-constraints?language=objc
func (v_ View) Constraints() []LayoutConstraint {
	rv := objc.Call[[]LayoutConstraint](v_, objc.Sel("constraints"))
	return rv
}

// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483691-heightadjustlimit?language=objc
func (v_ View) HeightAdjustLimit() float64 {
	rv := objc.Call[float64](v_, objc.Sel("heightAdjustLimit"))
	return rv
}

// A Boolean value that indicates whether the view’s horizontal size constraints are active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc
func (v_ View) IsHorizontalContentSizeConstraintActive() bool {
	rv := objc.Call[bool](v_, objc.Sel("isHorizontalContentSizeConstraintActive"))
	return rv
}

// A Boolean value that indicates whether the view’s horizontal size constraints are active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/3353053-horizontalcontentsizeconstrainta?language=objc
func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setHorizontalContentSizeConstraintActive:"), value)
}

// A Boolean value indicating whether the view ensures it is vibrant on top of other content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483793-allowsvibrancy?language=objc
func (v_ View) AllowsVibrancy() bool {
	rv := objc.Call[bool](v_, objc.Sel("allowsVibrancy"))
	return rv
}

// The array of views embedded in the current view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc
func (v_ View) Subviews() []View {
	rv := objc.Call[[]View](v_, objc.Sel("subviews"))
	return rv
}

// The array of views embedded in the current view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483539-subviews?language=objc
func (v_ View) SetSubviews(value []IView) {
	objc.Call[objc.Void](v_, objc.Sel("setSubviews:"), value)
}

// A Boolean value indicating whether the view uses a flipped coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483532-flipped?language=objc
func (v_ View) IsFlipped() bool {
	rv := objc.Call[bool](v_, objc.Sel("isFlipped"))
	return rv
}

// The rectangle identifying the portion of your view that did not change during a live resize operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483528-rectpreservedduringliveresize?language=objc
func (v_ View) RectPreservedDuringLiveResize() foundation.Rect {
	rv := objc.Call[foundation.Rect](v_, objc.Sel("rectPreservedDuringLiveResize"))
	return rv
}

// A Boolean value indicating whether the view can pass mouse down events through to its superviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483666-mousedowncanmovewindow?language=objc
func (v_ View) MouseDownCanMoveWindow() bool {
	rv := objc.Call[bool](v_, objc.Sel("mouseDownCanMoveWindow"))
	return rv
}

// A layout anchor representing the horizontal center of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526924-centerxanchor?language=objc
func (v_ View) CenterXAnchor() LayoutXAxisAnchor {
	rv := objc.Call[LayoutXAxisAnchor](v_, objc.Sel("centerXAnchor"))
	return rv
}

// A Boolean value indicating whether the view wants resting touches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc
func (v_ View) WantsRestingTouches() bool {
	rv := objc.Call[bool](v_, objc.Sel("wantsRestingTouches"))
	return rv
}

// A Boolean value indicating whether the view wants resting touches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483594-wantsrestingtouches?language=objc
func (v_ View) SetWantsRestingTouches(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setWantsRestingTouches:"), value)
}

// A layout anchor representing the baseline for the bottommost line of text in the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526959-lastbaselineanchor?language=objc
func (v_ View) LastBaselineAnchor() LayoutYAxisAnchor {
	rv := objc.Call[LayoutYAxisAnchor](v_, objc.Sel("lastBaselineAnchor"))
	return rv
}

// A layout anchor representing the width of the view’s frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526638-widthanchor?language=objc
func (v_ View) WidthAnchor() LayoutDimension {
	rv := objc.Call[LayoutDimension](v_, objc.Sel("widthAnchor"))
	return rv
}

// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc
func (v_ View) FrameRotation() float64 {
	rv := objc.Call[float64](v_, objc.Sel("frameRotation"))
	return rv
}

// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483412-framerotation?language=objc
func (v_ View) SetFrameRotation(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setFrameRotation:"), value)
}

// The view that is the parent of the current view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483737-superview?language=objc
func (v_ View) Superview() View {
	rv := objc.Call[View](v_, objc.Sel("superview"))
	return rv
}

// Returns a Boolean value indicating whether the view depends on the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526926-requiresconstraintbasedlayout?language=objc
func (vc _ViewClass) RequiresConstraintBasedLayout() bool {
	rv := objc.Call[bool](vc, objc.Sel("requiresConstraintBasedLayout"))
	return rv
}

// Returns a Boolean value indicating whether the view depends on the constraint-based layout system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1526926-requiresconstraintbasedlayout?language=objc
func View_RequiresConstraintBasedLayout() bool {
	return ViewClass.RequiresConstraintBasedLayout()
}

// A Boolean value that determines whether the view needs to be redrawn before being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc
func (v_ View) NeedsDisplay() bool {
	rv := objc.Call[bool](v_, objc.Sel("needsDisplay"))
	return rv
}

// A Boolean value that determines whether the view needs to be redrawn before being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483360-needsdisplay?language=objc
func (v_ View) SetNeedsDisplay(value bool) {
	objc.Call[objc.Void](v_, objc.Sel("setNeedsDisplay:"), value)
}
