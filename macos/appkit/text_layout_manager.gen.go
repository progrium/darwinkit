// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextLayoutManagerClass = _TextLayoutManagerClass{objc.GetClass("NSTextLayoutManager")}

type _TextLayoutManagerClass struct {
	objc.Class
}

type ITextLayoutManager interface {
	objc.IObject
	EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool)
	ReplaceTextContentManager(textContentManager ITextContentManager)
	ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString)
	ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []ITextElement)
	AddRenderingAttributeValueForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange)
	InvalidateRenderingAttributesForTextRange(textRange ITextRange)
	RemoveRenderingAttributeForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange)
	SetRenderingAttributesForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange)
	InvalidateLayoutForRange(range_ ITextRange)
	TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment
	EnsureLayoutForBounds(bounds coregraphics.Rect)
	EnsureLayoutForRange(range_ ITextRange)
	LayoutQueue() foundation.OperationQueue
	SetLayoutQueue(value foundation.IOperationQueue)
	RenderingAttributesValidator() func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)
	SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment))
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	UsesHyphenation() bool
	SetUsesHyphenation(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	TextContentManager() TextContentManager
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	TextSelectionNavigation() TextSelectionNavigation
	SetTextSelectionNavigation(value ITextSelectionNavigation)
	TextSelections() []TextSelection
	SetTextSelections(value []ITextSelection)
	UsageBoundsForTextContainer() coregraphics.Rect
	TextViewportLayoutController() TextViewportLayoutController
}

type TextLayoutManager struct {
	objc.Object
}

func MakeTextLayoutManager(ptr unsafe.Pointer) TextLayoutManager {
	return TextLayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (t_ TextLayoutManager) Init() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("init"))
	return rv
}

func TextLayoutManager_Init() TextLayoutManager {
	return TextLayoutManagerClass.Alloc().Init()
}

func (tc _TextLayoutManagerClass) Alloc() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](tc, objc.GetSelector("alloc"))
	return rv
}

func TextLayoutManager_Alloc() TextLayoutManager {
	return TextLayoutManagerClass.Alloc()
}

func (tc _TextLayoutManagerClass) New() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextLayoutManager() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

func TextLayoutManager_New() TextLayoutManager {
	return TextLayoutManagerClass.New()
}

func (t_ TextLayoutManager) EnumerateTextSegmentsInRangeTypeOptionsUsingBlock(textRange ITextRange, type_ TextLayoutManagerSegmentType, options TextLayoutManagerSegmentOptions, block func(textSegmentRange TextRange, textSegmentFrame coregraphics.Rect, baselinePosition float64, textContainer TextContainer) bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("enumerateTextSegmentsInRange:type:options:usingBlock:"), objc.ExtractPtr(textRange), type_, options, block)
}

func (t_ TextLayoutManager) ReplaceTextContentManager(textContentManager ITextContentManager) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceTextContentManager:"), objc.ExtractPtr(textContentManager))
}

func (t_ TextLayoutManager) ReplaceContentsInRangeWithAttributedString(range_ ITextRange, attributedString foundation.IAttributedString) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceContentsInRange:withAttributedString:"), objc.ExtractPtr(range_), objc.ExtractPtr(attributedString))
}

func (t_ TextLayoutManager) ReplaceContentsInRangeWithTextElements(range_ ITextRange, textElements []ITextElement) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceContentsInRange:withTextElements:"), objc.ExtractPtr(range_), textElements)
}

func (t_ TextLayoutManager) AddRenderingAttributeValueForTextRange(renderingAttribute foundation.AttributedStringKey, value objc.IObject, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("addRenderingAttribute:value:forTextRange:"), renderingAttribute, objc.ExtractPtr(value), objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) InvalidateRenderingAttributesForTextRange(textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateRenderingAttributesForTextRange:"), objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) RemoveRenderingAttributeForTextRange(renderingAttribute foundation.AttributedStringKey, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("removeRenderingAttribute:forTextRange:"), renderingAttribute, objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) SetRenderingAttributesForTextRange(renderingAttributes map[foundation.AttributedStringKey]objc.IObject, textRange ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRenderingAttributes:forTextRange:"), renderingAttributes, objc.ExtractPtr(textRange))
}

func (t_ TextLayoutManager) InvalidateLayoutForRange(range_ ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateLayoutForRange:"), objc.ExtractPtr(range_))
}

func (t_ TextLayoutManager) TextLayoutFragmentForPosition(position coregraphics.Point) TextLayoutFragment {
	rv := objc.CallMethod[TextLayoutFragment](t_, objc.GetSelector("textLayoutFragmentForPosition:"), position)
	return rv
}

func (t_ TextLayoutManager) EnsureLayoutForBounds(bounds coregraphics.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("ensureLayoutForBounds:"), bounds)
}

func (t_ TextLayoutManager) EnsureLayoutForRange(range_ ITextRange) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("ensureLayoutForRange:"), objc.ExtractPtr(range_))
}

func (t_ TextLayoutManager) LayoutQueue() foundation.OperationQueue {
	rv := objc.CallMethod[foundation.OperationQueue](t_, objc.GetSelector("layoutQueue"))
	return rv
}

func (t_ TextLayoutManager) SetLayoutQueue(value foundation.IOperationQueue) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutQueue:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) RenderingAttributesValidator() func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment) {
	rv := objc.CallMethod[func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)](t_, objc.GetSelector("renderingAttributesValidator"))
	return rv
}

func (t_ TextLayoutManager) SetRenderingAttributesValidator(value func(textLayoutManager TextLayoutManager, textLayoutFragment TextLayoutFragment)) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRenderingAttributesValidator:"), value)
}

func (t_ TextLayoutManager) UsesFontLeading() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesFontLeading"))
	return rv
}

func (t_ TextLayoutManager) SetUsesFontLeading(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesFontLeading:"), value)
}

func (t_ TextLayoutManager) UsesHyphenation() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesHyphenation"))
	return rv
}

func (t_ TextLayoutManager) SetUsesHyphenation(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesHyphenation:"), value)
}

func (t_ TextLayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("limitsLayoutForSuspiciousContents"))
	return rv
}

func (t_ TextLayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLimitsLayoutForSuspiciousContents:"), value)
}

func (t_ TextLayoutManager) TextContentManager() TextContentManager {
	rv := objc.CallMethod[TextContentManager](t_, objc.GetSelector("textContentManager"))
	return rv
}

func (t_ TextLayoutManager) TextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("textContainer"))
	return rv
}

func (t_ TextLayoutManager) SetTextContainer(value ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextContainer:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) TextSelectionNavigation() TextSelectionNavigation {
	rv := objc.CallMethod[TextSelectionNavigation](t_, objc.GetSelector("textSelectionNavigation"))
	return rv
}

func (t_ TextLayoutManager) SetTextSelectionNavigation(value ITextSelectionNavigation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextSelectionNavigation:"), objc.ExtractPtr(value))
}

func (t_ TextLayoutManager) TextSelections() []TextSelection {
	rv := objc.CallMethod[[]TextSelection](t_, objc.GetSelector("textSelections"))
	return rv
}

func (t_ TextLayoutManager) SetTextSelections(value []ITextSelection) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextSelections:"), value)
}

func (t_ TextLayoutManager) UsageBoundsForTextContainer() coregraphics.Rect {
	rv := objc.CallMethod[coregraphics.Rect](t_, objc.GetSelector("usageBoundsForTextContainer"))
	return rv
}

func (tc _TextLayoutManagerClass) LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](tc, objc.GetSelector("linkRenderingAttributes"))
	return rv
}

func TextLayoutManager_LinkRenderingAttributes() map[foundation.AttributedStringKey]objc.Object {
	return TextLayoutManagerClass.LinkRenderingAttributes()
}

func (t_ TextLayoutManager) TextViewportLayoutController() TextViewportLayoutController {
	rv := objc.CallMethod[TextViewportLayoutController](t_, objc.GetSelector("textViewportLayoutController"))
	return rv
}
