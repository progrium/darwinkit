// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}

type _LayoutManagerClass struct {
	objc.Class
}

type ILayoutManager interface {
	objc.IObject
	ReplaceTextStorage(newTextStorage ITextStorage)
	AddTextContainer(container ITextContainer)
	InsertTextContainerAtIndex(container ITextContainer, index uint)
	RemoveTextContainerAtIndex(index uint)
	SetTextContainerForGlyphRange(container ITextContainer, glyphRange foundation.Range)
	TextContainerChangedGeometry(container ITextContainer)
	TextContainerChangedTextView(container ITextContainer)
	TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) TextContainer
	TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) TextContainer
	UsedRectForTextContainer(container ITextContainer) foundation.Rect
	InvalidateDisplayForCharacterRange(charRange foundation.Range)
	InvalidateDisplayForGlyphRange(glyphRange foundation.Range)
	InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.Range, delta int, actualCharRange *foundation.Range)
	InvalidateLayoutForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range)
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	EnsureGlyphsForCharacterRange(charRange foundation.Range)
	EnsureGlyphsForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForBoundingRectInTextContainer(bounds foundation.Rect, container ITextContainer)
	EnsureLayoutForCharacterRange(charRange foundation.Range)
	EnsureLayoutForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForTextContainer(container ITextContainer)
	GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange foundation.Range, glyphBuffer *coregraphics.Glyph, props *GlyphProperty, charIndexBuffer *uint, bidiLevelBuffer *byte) uint
	CGGlyphAtIndex(glyphIndex uint) coregraphics.Glyph
	CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph
	SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range)
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	IsValidGlyphIndex(glyphIndex uint) bool
	PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty
	SetAttachmentSizeForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint)
	SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container ITextContainer)
	SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect)
	SetLocationForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range)
	SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint)
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	FirstUnlaidCharacterIndex() uint
	FirstUnlaidGlyphIndex() uint
	GetFirstUnlaidCharacterIndexGlyphIndex(charIndex *uint, glyphIndex *uint)
	LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect
	LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect
	LocationForGlyphAtIndex(glyphIndex uint) foundation.Point
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range
	BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.Range, container ITextContainer) foundation.Rect
	CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point foundation.Point, container ITextContainer, partialFraction *float64) uint
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range
	EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect foundation.Rect, stop *bool))
	EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.Range, block func(rect foundation.Rect, usedRect foundation.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool))
	FractionOfDistanceThroughGlyphForPointInTextContainer(point foundation.Point, container ITextContainer) float64
	GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint
	GlyphIndexForPointInTextContainer(point foundation.Point, container ITextContainer) uint
	GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point foundation.Point, container ITextContainer, partialFraction *float64) uint
	GlyphRangeForBoundingRectInTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range
	GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range
	GlyphRangeForTextContainer(container ITextContainer) foundation.Range
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range
	DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin foundation.Point)
	DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	FillBackgroundRectArrayCountForCharacterRangeColor(rectArray *foundation.Rect, rectCount uint, charRange foundation.Range, color IColor)
	ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef)
	StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point)
	SetLayoutRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	SetBoundsRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	LayoutRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	BoundsRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect
	ShowAttachmentCellInRectCharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint)
	RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View
	RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker
	LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool
	DefaultLineHeightForFont(theFont IFont) float64
	DefaultBaselineOffsetForFont(theFont IFont) float64
	AddTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	AddTemporaryAttributeValueForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range)
	SetTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	RemoveTemporaryAttributeForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range)
	TemporaryAttributeAtCharacterIndexEffectiveRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range) objc.Object
	TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range, rangeLimit foundation.Range) objc.Object
	TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object
	TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ *foundation.Range, rangeLimit foundation.Range) map[foundation.AttributedStringKey]objc.Object
	Delegate() LayoutManagerDelegateWrapper
	SetDelegate(value ILayoutManagerDelegate)
	SetDelegate0(value objc.IObject)
	TextStorage() TextStorage
	SetTextStorage(value ITextStorage)
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	HasNonContiguousLayout() bool
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
	TextContainers() []TextContainer
	GlyphGenerator() GlyphGenerator
	SetGlyphGenerator(value IGlyphGenerator)
	NumberOfGlyphs() uint
	ExtraLineFragmentRect() foundation.Rect
	ExtraLineFragmentTextContainer() TextContainer
	ExtraLineFragmentUsedRect() foundation.Rect
	DefaultAttachmentScaling() ImageScaling
	SetDefaultAttachmentScaling(value ImageScaling)
	FirstTextView() TextView
	TextViewForBeginningOfSelection() TextView
	Typesetter() Typesetter
	SetTypesetter(value ITypesetter)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
}

type LayoutManager struct {
	objc.Object
}

func MakeLayoutManager(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{
		Object: objc.MakeObject(ptr),
	}
}

func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.CallMethod[LayoutManager](l_, objc.GetSelector("init"))
	return rv
}

func LayoutManager_Init() LayoutManager {
	return LayoutManagerClass.Alloc().Init()
}

func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.CallMethod[LayoutManager](lc, objc.GetSelector("alloc"))
	return rv
}

func LayoutManager_Alloc() LayoutManager {
	return LayoutManagerClass.Alloc()
}

func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.CallMethod[LayoutManager](lc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutManager() LayoutManager {
	return LayoutManagerClass.New()
}

func LayoutManager_New() LayoutManager {
	return LayoutManagerClass.New()
}

func (l_ LayoutManager) ReplaceTextStorage(newTextStorage ITextStorage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("replaceTextStorage:"), objc.ExtractPtr(newTextStorage))
}

func (l_ LayoutManager) AddTextContainer(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTextContainer:"), objc.ExtractPtr(container))
}

func (l_ LayoutManager) InsertTextContainerAtIndex(container ITextContainer, index uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("insertTextContainer:atIndex:"), objc.ExtractPtr(container), index)
}

func (l_ LayoutManager) RemoveTextContainerAtIndex(index uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeTextContainerAtIndex:"), index)
}

func (l_ LayoutManager) SetTextContainerForGlyphRange(container ITextContainer, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTextContainer:forGlyphRange:"), objc.ExtractPtr(container), glyphRange)
}

func (l_ LayoutManager) TextContainerChangedGeometry(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("textContainerChangedGeometry:"), objc.ExtractPtr(container))
}

func (l_ LayoutManager) TextContainerChangedTextView(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("textContainerChangedTextView:"), objc.ExtractPtr(container))
}

func (l_ LayoutManager) TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("textContainerForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) TextContainerForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("textContainerForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) UsedRectForTextContainer(container ITextContainer) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("usedRectForTextContainer:"), objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateDisplayForCharacterRange:"), charRange)
}

func (l_ LayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateDisplayForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.Range, delta int, actualCharRange *foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:"), charRange, delta, actualCharRange)
}

func (l_ LayoutManager) InvalidateLayoutForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("invalidateLayoutForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
}

func (l_ LayoutManager) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), objc.ExtractPtr(textStorage), editMask, newCharRange, delta, invalidatedCharRange)
}

func (l_ LayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureGlyphsForCharacterRange:"), charRange)
}

func (l_ LayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureGlyphsForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForBoundingRectInTextContainer(bounds foundation.Rect, container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForBoundingRect:inTextContainer:"), bounds, objc.ExtractPtr(container))
}

func (l_ LayoutManager) EnsureLayoutForCharacterRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForCharacterRange:"), charRange)
}

func (l_ LayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForGlyphRange:"), glyphRange)
}

func (l_ LayoutManager) EnsureLayoutForTextContainer(container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("ensureLayoutForTextContainer:"), objc.ExtractPtr(container))
}

func (l_ LayoutManager) GetGlyphsInRangeGlyphsPropertiesCharacterIndexesBidiLevels(glyphRange foundation.Range, glyphBuffer *coregraphics.Glyph, props *GlyphProperty, charIndexBuffer *uint, bidiLevelBuffer *byte) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getGlyphsInRange:glyphs:properties:characterIndexes:bidiLevels:"), glyphRange, glyphBuffer, props, charIndexBuffer, bidiLevelBuffer)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndex(glyphIndex uint) coregraphics.Glyph {
	rv := objc.CallMethod[coregraphics.Glyph](l_, objc.GetSelector("CGGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph {
	rv := objc.CallMethod[coregraphics.Glyph](l_, objc.GetSelector("CGGlyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}

func (l_ LayoutManager) SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setGlyphs:properties:characterIndexes:font:forGlyphRange:"), glyphs, props, charIndexes, objc.ExtractPtr(aFont), glyphRange)
}

func (l_ LayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("characterIndexForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForCharacterAtIndex:"), charIndex)
	return rv
}

func (l_ LayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("isValidGlyphIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	rv := objc.CallMethod[GlyphProperty](l_, objc.GetSelector("propertyForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) SetAttachmentSizeForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

func (l_ LayoutManager) SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDrawsOutsideLineFragment:forGlyphAtIndex:"), flag, glyphIndex)
}

func (l_ LayoutManager) SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect foundation.Rect, usedRect foundation.Rect, container ITextContainer) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setExtraLineFragmentRect:usedRect:textContainer:"), fragmentRect, usedRect, objc.ExtractPtr(container))
}

func (l_ LayoutManager) SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLineFragmentRect:forGlyphRange:usedRect:"), fragmentRect, glyphRange, usedRect)
}

func (l_ LayoutManager) SetLocationForStartOfGlyphRange(location foundation.Point, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLocation:forStartOfGlyphRange:"), location, glyphRange)
}

func (l_ LayoutManager) SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setNotShownAttribute:forGlyphAtIndex:"), flag, glyphIndex)
}

func (l_ LayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) foundation.Size {
	rv := objc.CallMethod[foundation.Size](l_, objc.GetSelector("attachmentSizeForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("drawsOutsideLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("firstUnlaidCharacterIndex"))
	return rv
}

func (l_ LayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("firstUnlaidGlyphIndex"))
	return rv
}

func (l_ LayoutManager) GetFirstUnlaidCharacterIndexGlyphIndex(charIndex *uint, glyphIndex *uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("getFirstUnlaidCharacterIndex:glyphIndex:"), charIndex, glyphIndex)
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRangeWithoutAdditionalLayout(glyphIndex uint, effectiveGlyphRange *foundation.Range, flag bool) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:withoutAdditionalLayout:"), glyphIndex, effectiveGlyphRange, flag)
	return rv
}

func (l_ LayoutManager) LocationForGlyphAtIndex(glyphIndex uint) foundation.Point {
	rv := objc.CallMethod[foundation.Point](l_, objc.GetSelector("locationForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("notShownAttributeForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("truncatedGlyphRangeInLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.Range, container ITextContainer) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundingRectForGlyphRange:inTextContainer:"), glyphRange, objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:"), point, objc.ExtractPtr(container), partialFraction)
	return rv
}

func (l_ LayoutManager) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}

func (l_ LayoutManager) EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect foundation.Rect, stop *bool)) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:"), glyphRange, selectedRange, objc.ExtractPtr(textContainer), block)
}

func (l_ LayoutManager) EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.Range, block func(rect foundation.Rect, usedRect foundation.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool)) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("enumerateLineFragmentsForGlyphRange:usingBlock:"), glyphRange, block)
}

func (l_ LayoutManager) FractionOfDistanceThroughGlyphForPointInTextContainer(point foundation.Point, container ITextContainer) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("fractionOfDistanceThroughGlyphForPoint:inTextContainer:"), point, objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:"), charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}

func (l_ LayoutManager) GlyphIndexForPointInTextContainer(point foundation.Point, container ITextContainer) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForPoint:inTextContainer:"), point, objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) GlyphIndexForPointInTextContainerFractionOfDistanceThroughGlyph(point foundation.Point, container ITextContainer, partialFraction *float64) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("glyphIndexForPoint:inTextContainer:fractionOfDistanceThroughGlyph:"), point, objc.ExtractPtr(container), partialFraction)
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRectInTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForBoundingRect:inTextContainer:"), bounds, objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds foundation.Rect, container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:"), bounds, objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) GlyphRangeForTextContainer(container ITextContainer) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForTextContainer:"), objc.ExtractPtr(container))
	return rv
}

func (l_ LayoutManager) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange *foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}

func (l_ LayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range {
	rv := objc.CallMethod[foundation.Range](l_, objc.GetSelector("rangeOfNominallySpacedGlyphsContainingIndex:"), glyphIndex)
	return rv
}

func (l_ LayoutManager) DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawBackgroundForGlyphRange:atPoint:"), glyphsToShow, origin)
}

func (l_ LayoutManager) DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawGlyphsForGlyphRange:atPoint:"), glyphsToShow, origin)
}

func (l_ LayoutManager) DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) FillBackgroundRectArrayCountForCharacterRangeColor(rectArray *foundation.Rect, rectCount uint, charRange foundation.Range, color IColor) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("fillBackgroundRectArray:count:forCharacterRange:color:"), rectArray, rectCount, charRange, objc.ExtractPtr(color))
}

func (l_ LayoutManager) ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:"), glyphs, positions, glyphCount, objc.ExtractPtr(font), textMatrix, attributes, CGContext)
}

func (l_ LayoutManager) StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect foundation.Rect, lineGlyphRange foundation.Range, containerOrigin foundation.Point) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}

func (l_ LayoutManager) SetLayoutRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLayoutRect:forTextBlock:glyphRange:"), rect, objc.ExtractPtr(block), glyphRange)
}

func (l_ LayoutManager) LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutRectForTextBlock:glyphRange:"), objc.ExtractPtr(block), glyphRange)
	return rv
}

func (l_ LayoutManager) SetBoundsRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBoundsRect:forTextBlock:glyphRange:"), rect, objc.ExtractPtr(block), glyphRange)
}

func (l_ LayoutManager) BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundsRectForTextBlock:glyphRange:"), objc.ExtractPtr(block), glyphRange)
	return rv
}

func (l_ LayoutManager) LayoutRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutRectForTextBlock:atIndex:effectiveRange:"), objc.ExtractPtr(block), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) BoundsRectForTextBlockAtIndexEffectiveRange(block ITextBlock, glyphIndex uint, effectiveGlyphRange *foundation.Range) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("boundsRectForTextBlock:atIndex:effectiveRange:"), objc.ExtractPtr(block), glyphIndex, effectiveGlyphRange)
	return rv
}

func (l_ LayoutManager) ShowAttachmentCellInRectCharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("showAttachmentCell:inRect:characterIndex:"), objc.ExtractPtr(cell), rect, attachmentIndex)
}

func (l_ LayoutManager) RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View {
	rv := objc.CallMethod[View](l_, objc.GetSelector("rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:"), objc.ExtractPtr(view), objc.ExtractPtr(style), objc.ExtractPtr(ruler), isEnabled)
	return rv
}

func (l_ LayoutManager) RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker {
	rv := objc.CallMethod[[]RulerMarker](l_, objc.GetSelector("rulerMarkersForTextView:paragraphStyle:ruler:"), objc.ExtractPtr(view), objc.ExtractPtr(style), objc.ExtractPtr(ruler))
	return rv
}

func (l_ LayoutManager) LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManagerOwnsFirstResponderInWindow:"), objc.ExtractPtr(window))
	return rv
}

func (l_ LayoutManager) DefaultLineHeightForFont(theFont IFont) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("defaultLineHeightForFont:"), objc.ExtractPtr(theFont))
	return rv
}

func (l_ LayoutManager) DefaultBaselineOffsetForFont(theFont IFont) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("defaultBaselineOffsetForFont:"), objc.ExtractPtr(theFont))
	return rv
}

func (l_ LayoutManager) AddTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

func (l_ LayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("addTemporaryAttribute:value:forCharacterRange:"), attrName, objc.ExtractPtr(value), charRange)
}

func (l_ LayoutManager) SetTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

func (l_ LayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

func (l_ LayoutManager) TemporaryAttributeAtCharacterIndexEffectiveRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("temporaryAttribute:atCharacterIndex:effectiveRange:"), attrName, location, range_)
	return rv
}

func (l_ LayoutManager) TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.AttributedStringKey, location uint, range_ *foundation.Range, rangeLimit foundation.Range) objc.Object {
	rv := objc.CallMethod[objc.Object](l_, objc.GetSelector("temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("temporaryAttributesAtCharacterIndex:effectiveRange:"), charIndex, effectiveCharRange)
	return rv
}

func (l_ LayoutManager) TemporaryAttributesAtCharacterIndexLongestEffectiveRangeInRange(location uint, range_ *foundation.Range, rangeLimit foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("temporaryAttributesAtCharacterIndex:longestEffectiveRange:inRange:"), location, range_, rangeLimit)
	return rv
}

func (l_ LayoutManager) Delegate() LayoutManagerDelegateWrapper {
	rv := objc.CallMethod[LayoutManagerDelegateWrapper](l_, objc.GetSelector("delegate"))
	return rv
}

func (l_ LayoutManager) SetDelegate(value ILayoutManagerDelegate) {
	po := objc.WrapAsProtocol("NSLayoutManagerDelegate", value)
	objc.SetAssociatedObject(l_, objc.AssociationKey("setDelegate"), po, objc.ASSOCIATION_RETAIN)
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), po)
}

func (l_ LayoutManager) SetDelegate0(value objc.IObject) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDelegate:"), objc.ExtractPtr(value))
}

func (l_ LayoutManager) TextStorage() TextStorage {
	rv := objc.CallMethod[TextStorage](l_, objc.GetSelector("textStorage"))
	return rv
}

func (l_ LayoutManager) SetTextStorage(value ITextStorage) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTextStorage:"), objc.ExtractPtr(value))
}

func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("allowsNonContiguousLayout"))
	return rv
}

func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setAllowsNonContiguousLayout:"), value)
}

func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("hasNonContiguousLayout"))
	return rv
}

func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("showsInvisibleCharacters"))
	return rv
}

func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShowsInvisibleCharacters:"), value)
}

func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("showsControlCharacters"))
	return rv
}

func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setShowsControlCharacters:"), value)
}

func (l_ LayoutManager) UsesFontLeading() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesFontLeading"))
	return rv
}

func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setUsesFontLeading:"), value)
}

func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("backgroundLayoutEnabled"))
	return rv
}

func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setBackgroundLayoutEnabled:"), value)
}

func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("limitsLayoutForSuspiciousContents"))
	return rv
}

func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setLimitsLayoutForSuspiciousContents:"), value)
}

func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("usesDefaultHyphenation"))
	return rv
}

func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setUsesDefaultHyphenation:"), value)
}

func (l_ LayoutManager) TextContainers() []TextContainer {
	rv := objc.CallMethod[[]TextContainer](l_, objc.GetSelector("textContainers"))
	return rv
}

func (l_ LayoutManager) GlyphGenerator() GlyphGenerator {
	rv := objc.CallMethod[GlyphGenerator](l_, objc.GetSelector("glyphGenerator"))
	return rv
}

func (l_ LayoutManager) SetGlyphGenerator(value IGlyphGenerator) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setGlyphGenerator:"), objc.ExtractPtr(value))
}

func (l_ LayoutManager) NumberOfGlyphs() uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("numberOfGlyphs"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("extraLineFragmentRect"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentTextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](l_, objc.GetSelector("extraLineFragmentTextContainer"))
	return rv
}

func (l_ LayoutManager) ExtraLineFragmentUsedRect() foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("extraLineFragmentUsedRect"))
	return rv
}

func (l_ LayoutManager) DefaultAttachmentScaling() ImageScaling {
	rv := objc.CallMethod[ImageScaling](l_, objc.GetSelector("defaultAttachmentScaling"))
	return rv
}

func (l_ LayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setDefaultAttachmentScaling:"), value)
}

func (l_ LayoutManager) FirstTextView() TextView {
	rv := objc.CallMethod[TextView](l_, objc.GetSelector("firstTextView"))
	return rv
}

func (l_ LayoutManager) TextViewForBeginningOfSelection() TextView {
	rv := objc.CallMethod[TextView](l_, objc.GetSelector("textViewForBeginningOfSelection"))
	return rv
}

func (l_ LayoutManager) Typesetter() Typesetter {
	rv := objc.CallMethod[Typesetter](l_, objc.GetSelector("typesetter"))
	return rv
}

func (l_ LayoutManager) SetTypesetter(value ITypesetter) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTypesetter:"), objc.ExtractPtr(value))
}

func (l_ LayoutManager) TypesetterBehavior() TypesetterBehavior {
	rv := objc.CallMethod[TypesetterBehavior](l_, objc.GetSelector("typesetterBehavior"))
	return rv
}

func (l_ LayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("setTypesetterBehavior:"), value)
}
