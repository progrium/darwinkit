// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextLayoutManager] class.
var TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}

type _TextLayoutManagerClass struct {
	objc.Class
}

// An interface definition for the [TextLayoutManager] class.
type ITextLayoutManager interface {
	objc.IObject
	TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment
	EnsureLayoutForRange(range_ ITextRange)
	RenderingAttributesForLinkAtLocation(link objc.IObject, location PTextLocation) map[foundation.AttributedStringKey]objc.Object
	RenderingAttributesForLinkAtLocationObject(link objc.IObject, locationObject objc.IObject) map[foundation.AttributedStringKey]objc.Object
	SetRenderingAttributesForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange)
	ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString)
	TextLayoutFragmentForLocation(location PTextLocation) TextLayoutFragment
	TextLayoutFragmentForLocationObject(locationObject objc.IObject) TextLayoutFragment
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool)
	RemoveRenderingAttributeForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange)
	EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location PTextLocation, options TextLayoutFragmentEnumerationOptions, block func(layoutFragment TextLayoutFragment) bool) TextLocationWrapper
	EnumerateTextLayoutFragmentsFromLocationObjectOptionsUsingBlock(locationObject objc.IObject, options TextLayoutFragmentEnumerationOptions, block func(layoutFragment TextLayoutFragment) bool) TextLocationWrapper
	EnumerateRenderingAttributesFromLocationReverseUsingBlock(location PTextLocation, reverse bool, block func(textLayoutManager TextLayoutManager, attributes map[foundation.AttributedStringKey]objc.Object, textRange TextRange) bool)
	EnumerateRenderingAttributesFromLocationObjectReverseUsingBlock(locationObject objc.IObject, reverse bool, block func(textLayoutManager TextLayoutManager, attributes map[foundation.AttributedStringKey]objc.Object, textRange TextRange) bool)
	AddRenderingAttributeValueForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange)
	EnsureLayoutForBounds(bounds coregraphics.Rect)
	ReplaceTextContentManager(textContentManager ITextContentManager)
	InvalidateLayoutForRange(range_ ITextRange)
	InvalidateRenderingAttributesForTextRange(textRange ITextRange)
	TextSelections() []TextSelection
	SetTextSelections(value []ITextSelection)
	TextSelectionNavigation() TextSelectionNavigation
	SetTextSelectionNavigation(value ITextSelectionNavigation)
	RenderingAttributesValidator() func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)
	SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment))
	TextViewportLayoutController() TextViewportLayoutController
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	UsageBoundsForTextContainer() coregraphics.Rect
	Delegate() TextLayoutManagerDelegateWrapper
	SetDelegate(value PTextLayoutManagerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	UsesHyphenation() bool
	SetUsesHyphenation(value bool)
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	TextContentManager() TextContentManager
}

// The primary class that you use to manage text layout and presentation for custom text displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager?language=objc
type TextLayoutManager struct {
	objc.Object
}

func TextLayoutManagerFrom(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("init"))
	return rv
}

func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](tc, objc.Sel("alloc"))
	return rv
}

func TextLayoutManager_Alloc() TextLayoutManager {
	return TextLayoutManagerClass.Alloc()
}

func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutManager() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

// Returns the text layout fragment at the position you specify in the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810013-textlayoutfragmentforposition?language=objc
func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](t_, objc.Sel("textLayoutFragmentForPosition:"), position)
	return rv
}

// Performs the layout for specified text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809992-ensurelayoutforrange?language=objc
func (t_ TextLayoutManager) EnsureLayoutForRange(range_ ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("ensureLayoutForRange:"), objc.Ptr(range_))
}

// Returns a dictionary of rendering attributes for rendering a link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810004-renderingattributesforlink?language=objc
func (t_ TextLayoutManager) RenderingAttributesForLinkAtLocation(link objc.IObject, location PTextLocation) map[foundation.AttributedStringKey]objc.Object {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("renderingAttributesForLink:atLocation:"), link, po1)
	return rv
}

// Returns a dictionary of rendering attributes for rendering a link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810004-renderingattributesforlink?language=objc
func (t_ TextLayoutManager) RenderingAttributesForLinkAtLocationObject(link objc.IObject, locationObject objc.IObject) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("renderingAttributesForLink:atLocation:"), link, objc.Ptr(locationObject))
	return rv
}

// Sets the rendering attributes for the range you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810009-setrenderingattributes?language=objc
func (t_ TextLayoutManager) SetRenderingAttributesForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("setRenderingAttributes:forTextRange:"), renderingAttributes, objc.Ptr(textRange))
}

// Replaces content at the location you specify with an attributed string you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810006-replacecontentsinrange?language=objc
func (t_ TextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("replaceContentsInRange:withAttributedString:"), objc.Ptr(range_), objc.Ptr(attributedString))
}

// Returns the text layout fragment from the document at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810012-textlayoutfragmentforlocation?language=objc
func (t_ TextLayoutManager) TextLayoutFragmentForLocation(location PTextLocation) TextLayoutFragment {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextLayoutFragment](t_, objc.Sel("textLayoutFragmentForLocation:"), po0)
	return rv
}

// Returns the text layout fragment from the document at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810012-textlayoutfragmentforlocation?language=objc
func (t_ TextLayoutManager) TextLayoutFragmentForLocationObject(locationObject objc.IObject) TextLayoutFragment {
	rv := objc.Call[TextLayoutFragment](t_, objc.Sel("textLayoutFragmentForLocation:"), objc.Ptr(locationObject))
	return rv
}

// Enumerates text segments of a specific type and in the text range you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809995-enumeratetextsegmentsinrange?language=objc
func (t_ TextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool) {
	objc.Call[objc.Void](t_, objc.Sel("enumerateTextSegmentsInRange:type:options:usingBlock:"), objc.Ptr(textRange), type_, options, block)
}

// Removes the rendering attribute from the specified text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810003-removerenderingattribute?language=objc
func (t_ TextLayoutManager) RemoveRenderingAttributeForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("removeRenderingAttribute:forTextRange:"), renderingAttribute, objc.Ptr(textRange))
}

// Enumerates the text layout fragments starting at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809994-enumeratetextlayoutfragmentsfrom?language=objc
func (t_ TextLayoutManager) EnumerateTextLayoutFragmentsFromLocationOptionsUsingBlock(location PTextLocation, options TextLayoutFragmentEnumerationOptions, block func(layoutFragment TextLayoutFragment) bool) TextLocationWrapper {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("enumerateTextLayoutFragmentsFromLocation:options:usingBlock:"), po0, options, block)
	return rv
}

// Enumerates the text layout fragments starting at the specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809994-enumeratetextlayoutfragmentsfrom?language=objc
func (t_ TextLayoutManager) EnumerateTextLayoutFragmentsFromLocationObjectOptionsUsingBlock(locationObject objc.IObject, options TextLayoutFragmentEnumerationOptions, block func(layoutFragment TextLayoutFragment) bool) TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("enumerateTextLayoutFragmentsFromLocation:options:usingBlock:"), objc.Ptr(locationObject), options, block)
	return rv
}

// Enumerates the rendering attributes from a location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809993-enumeraterenderingattributesfrom?language=objc
func (t_ TextLayoutManager) EnumerateRenderingAttributesFromLocationReverseUsingBlock(location PTextLocation, reverse bool, block func(textLayoutManager TextLayoutManager, attributes map[foundation.AttributedStringKey]objc.Object, textRange TextRange) bool) {
	po0 := objc.WrapAsProtocol("NSTextLocation", location)
	objc.Call[objc.Void](t_, objc.Sel("enumerateRenderingAttributesFromLocation:reverse:usingBlock:"), po0, reverse, block)
}

// Enumerates the rendering attributes from a location you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809993-enumeraterenderingattributesfrom?language=objc
func (t_ TextLayoutManager) EnumerateRenderingAttributesFromLocationObjectReverseUsingBlock(locationObject objc.IObject, reverse bool, block func(textLayoutManager TextLayoutManager, attributes map[foundation.AttributedStringKey]objc.Object, textRange TextRange) bool) {
	objc.Call[objc.Void](t_, objc.Sel("enumerateRenderingAttributesFromLocation:reverse:usingBlock:"), objc.Ptr(locationObject), reverse, block)
}

// Sets the rendering attribute for the value and range you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809989-addrenderingattribute?language=objc
func (t_ TextLayoutManager) AddRenderingAttributeValueForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("addRenderingAttribute:value:forTextRange:"), renderingAttribute, value, objc.Ptr(textRange))
}

// Performs the layout for filling the bounds you specify inside the last text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809991-ensurelayoutforbounds?language=objc
func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds coregraphics.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("ensureLayoutForBounds:"), bounds)
}

// Replaces the current text content manager with a new one you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810008-replacetextcontentmanager?language=objc
func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	objc.Call[objc.Void](t_, objc.Sel("replaceTextContentManager:"), objc.Ptr(textContentManager))
}

// Invalidates the layout information for specified text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809998-invalidatelayoutforrange?language=objc
func (t_ TextLayoutManager) InvalidateLayoutForRange(range_ ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("invalidateLayoutForRange:"), objc.Ptr(range_))
}

// Invalidates the rendering attributes of the specified text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809999-invalidaterenderingattributesfor?language=objc
func (t_ TextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange ITextRange) {
	objc.Call[objc.Void](t_, objc.Sel("invalidateRenderingAttributesForTextRange:"), objc.Ptr(textRange))
}

// Returns the default set of attributes for rendering a link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810002-linkrenderingattributes?language=objc
func (tc _TextLayoutManagerClass) LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](tc, objc.Sel("linkRenderingAttributes"))
	return rv
}

// Returns the default set of attributes for rendering a link. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810002-linkrenderingattributes?language=objc
func TextLayoutManager_LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	return TextLayoutManagerClass.LinkRenderingAttributes()
}

// An array of text selections associated by the text layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810015-textselections?language=objc
func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := objc.Call[[]TextSelection](t_, objc.Sel("textSelections"))
	return rv
}

// An array of text selections associated by the text layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810015-textselections?language=objc
func (t_ TextLayoutManager) SetTextSelections(value []ITextSelection) {
	objc.Call[objc.Void](t_, objc.Sel("setTextSelections:"), value)
}

// Returns a text selection manager configured to have the text layout manager as its data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810014-textselectionnavigation?language=objc
func (t_ TextLayoutManager) TextSelectionNavigation() TextSelectionNavigation {
	rv := objc.Call[TextSelectionNavigation](t_, objc.Sel("textSelectionNavigation"))
	return rv
}

// Returns a text selection manager configured to have the text layout manager as its data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810014-textselectionnavigation?language=objc
func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	objc.Call[objc.Void](t_, objc.Sel("setTextSelectionNavigation:"), objc.Ptr(value))
}

// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810005-renderingattributesvalidator?language=objc
func (t_ TextLayoutManager) RenderingAttributesValidator() func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment) {
	rv := objc.Call[func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)](t_, objc.Sel("renderingAttributesValidator"))
	return rv
}

// A callback block that the framework invokes whenever the text layout manager needs to validate the rendering attributes for the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810005-renderingattributesvalidator?language=objc
func (t_ TextLayoutManager) SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)) {
	objc.Call[objc.Void](t_, objc.Sel("setRenderingAttributesValidator:"), value)
}

// Returns text viewport layout controller associated with the layout manager’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810016-textviewportlayoutcontroller?language=objc
func (t_ TextLayoutManager) TextViewportLayoutController() TextViewportLayoutController {
	rv := objc.Call[TextViewportLayoutController](t_, objc.Sel("textViewportLayoutController"))
	return rv
}

// The queue that the framework dispatches layout operations on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810000-layoutqueue?language=objc
func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := objc.Call[foundation.OperationQueue](t_, objc.Sel("layoutQueue"))
	return rv
}

// The queue that the framework dispatches layout operations on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810000-layoutqueue?language=objc
func (t_ TextLayoutManager) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutQueue:"), objc.Ptr(value))
}

// Returns the usage bounds for the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810017-usageboundsfortextcontainer?language=objc
func (t_ TextLayoutManager) UsageBoundsForTextContainer() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("usageBoundsForTextContainer"))
	return rv
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809990-delegate?language=objc
func (t_ TextLayoutManager) Delegate() TextLayoutManagerDelegateWrapper {
	rv := objc.Call[TextLayoutManagerDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809990-delegate?language=objc
func (t_ TextLayoutManager) SetDelegate(value PTextLayoutManagerDelegate) {
	po0 := objc.WrapAsProtocol("NSTextLayoutManagerDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The delegate for the text layout manager object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3809990-delegate?language=objc
func (t_ TextLayoutManager) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810001-limitslayoutforsuspiciouscontent?language=objc
func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Call[bool](t_, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}

// A Boolean value that controls internal security analysis for malicious inputs and activates defensive behaviors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810001-limitslayoutforsuspiciouscontent?language=objc
func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}

// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810019-useshyphenation?language=objc
func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesHyphenation"))
	return rv
}

// A Boolean values that controls whether the text layout manager attempts to hyphenate when wrapping lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810019-useshyphenation?language=objc
func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesHyphenation:"), value)
}

// The text container object that provides geometric information for the layout destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810010-textcontainer?language=objc
func (t_ TextLayoutManager) TextContainer() TextContainer {
	rv := objc.Call[TextContainer](t_, objc.Sel("textContainer"))
	return rv
}

// The text container object that provides geometric information for the layout destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810010-textcontainer?language=objc
func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.Call[objc.Void](t_, objc.Sel("setTextContainer:"), objc.Ptr(value))
}

// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810018-usesfontleading?language=objc
func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesFontLeading"))
	return rv
}

// A Boolean value that controls whether the framework uses the leading information specified by the font when laying out text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810018-usesfontleading?language=objc
func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesFontLeading:"), value)
}

// Returns the text content manager associated with this text layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextlayoutmanager/3810011-textcontentmanager?language=objc
func (t_ TextLayoutManager) TextContentManager() TextContentManager {
	rv := objc.Call[TextContentManager](t_, objc.Sel("textContentManager"))
	return rv
}
