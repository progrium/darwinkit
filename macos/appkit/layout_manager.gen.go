// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LayoutManager] class.
var LayoutManagerClass = _LayoutManagerClass{objc.GetClass("NSLayoutManager")}

type _LayoutManagerClass struct {
	objc.Class
}

// An interface definition for the [LayoutManager] class.
type ILayoutManager interface {
	objc.IObject
	DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool
	SetLocationForStartOfGlyphRange(location coregraphics.Point, glyphRange foundation.Range)
	SetTextContainerForGlyphRange(container ITextContainer, glyphRange foundation.Range)
	GlyphIndexForPointInTextContainer(point coregraphics.Point, container ITextContainer) uint
	LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) coregraphics.Rect
	TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range
	TextContainerChangedTextView(container ITextContainer)
	TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.Object
	SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect coregraphics.Rect, usedRect coregraphics.Rect, container ITextContainer)
	AddTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	SetBoundsRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool
	EnsureGlyphsForGlyphRange(glyphRange foundation.Range)
	EnsureLayoutForCharacterRange(charRange foundation.Range)
	SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint)
	EnsureLayoutForGlyphRange(glyphRange foundation.Range)
	GlyphIndexForCharacterAtIndex(charIndex uint) uint
	DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point)
	AttachmentSizeForGlyphAtIndex(glyphIndex uint) coregraphics.Size
	InvalidateDisplayForCharacterRange(charRange foundation.Range)
	AddTemporaryAttributeValueForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range)
	SetTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range)
	UsedRectForTextContainer(container ITextContainer) coregraphics.Rect
	RemoveTextContainerAtIndex(index uint)
	CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point coregraphics.Point, container ITextContainer, partialFraction *float64) uint
	InsertTextContainerAtIndex(container ITextContainer, index uint)
	AddTextContainer(container ITextContainer)
	DefaultBaselineOffsetForFont(theFont IFont) float64
	LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) coregraphics.Rect
	DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin coregraphics.Point)
	EnsureLayoutForTextContainer(container ITextContainer)
	ShowAttachmentCellInRectCharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint)
	InvalidateDisplayForGlyphRange(glyphRange foundation.Range)
	GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint
	GlyphRangeForTextContainer(container ITextContainer) foundation.Range
	TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) TextContainer
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange foundation.RangePointer) foundation.Range
	SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint)
	BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	FillBackgroundRectArrayCountForCharacterRangeColor(rectArray *coregraphics.Rect, rectCount uint, charRange foundation.Range, color IColor)
	ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef)
	FirstUnlaidGlyphIndex() uint
	CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph
	EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.Range, block func(rect coregraphics.Rect, usedRect coregraphics.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool))
	TextContainerChangedGeometry(container ITextContainer)
	ReplaceTextStorage(newTextStorage ITextStorage)
	SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range)
	EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect coregraphics.Rect, stop *bool))
	SetLayoutRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range)
	ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range)
	UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point)
	FirstUnlaidCharacterIndex() uint
	DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin coregraphics.Point)
	SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect coregraphics.Rect, glyphRange foundation.Range, usedRect coregraphics.Rect)
	NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool
	StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point)
	InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.Range, delta int, actualCharRange foundation.RangePointer)
	EnsureGlyphsForCharacterRange(charRange foundation.Range)
	GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds coregraphics.Rect, container ITextContainer) foundation.Range
	RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker
	SetAttachmentSizeForGlyphRange(attachmentSize coregraphics.Size, glyphRange foundation.Range)
	DefaultLineHeightForFont(theFont IFont) float64
	EnsureLayoutForBoundingRectInTextContainer(bounds coregraphics.Rect, container ITextContainer)
	PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange foundation.RangePointer) foundation.Range
	CharacterIndexForGlyphAtIndex(glyphIndex uint) uint
	LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect
	LocationForGlyphAtIndex(glyphIndex uint) coregraphics.Point
	GlyphRangeForBoundingRectInTextContainer(bounds coregraphics.Rect, container ITextContainer) foundation.Range
	TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.AttributedStringKey, location uint, range_ foundation.RangePointer, rangeLimit foundation.Range) objc.Object
	GetFirstUnlaidCharacterIndexGlyphIndex(charIndex *uint, glyphIndex *uint)
	RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View
	RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range
	BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.Range, container ITextContainer) coregraphics.Rect
	FractionOfDistanceThroughGlyphForPointInTextContainer(point coregraphics.Point, container ITextContainer) float64
	RemoveTemporaryAttributeForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range)
	DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point)
	IsValidGlyphIndex(glyphIndex uint) bool
	ShowsInvisibleCharacters() bool
	SetShowsInvisibleCharacters(value bool)
	ShowsControlCharacters() bool
	SetShowsControlCharacters(value bool)
	ExtraLineFragmentRect() coregraphics.Rect
	TextViewForBeginningOfSelection() TextView
	HasNonContiguousLayout() bool
	UsesDefaultHyphenation() bool
	SetUsesDefaultHyphenation(value bool)
	NumberOfGlyphs() uint
	Typesetter() Typesetter
	SetTypesetter(value ITypesetter)
	Delegate() LayoutManagerDelegateWrapper
	SetDelegate(value PLayoutManagerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	AllowsNonContiguousLayout() bool
	SetAllowsNonContiguousLayout(value bool)
	LimitsLayoutForSuspiciousContents() bool
	SetLimitsLayoutForSuspiciousContents(value bool)
	GlyphGenerator() GlyphGenerator
	SetGlyphGenerator(value IGlyphGenerator)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	TextContainers() []TextContainer
	FirstTextView() TextView
	BackgroundLayoutEnabled() bool
	SetBackgroundLayoutEnabled(value bool)
	ExtraLineFragmentUsedRect() coregraphics.Rect
	TextStorage() TextStorage
	SetTextStorage(value ITextStorage)
	DefaultAttachmentScaling() ImageScaling
	SetDefaultAttachmentScaling(value ImageScaling)
	ExtraLineFragmentTextContainer() TextContainer
}

// An object that coordinates the layout and display of text characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager?language=objc
type LayoutManager struct {
	objc.Object
}

func LayoutManagerFrom(ptr unsafe.Pointer) LayoutManager {
	return LayoutManager{
		Object: objc.ObjectFrom(ptr),
	}
}

func (l_ LayoutManager) Init() LayoutManager {
	rv := objc.Call[LayoutManager](l_, objc.Sel("init"))
	return rv
}

func (lc _LayoutManagerClass) Alloc() LayoutManager {
	rv := objc.Call[LayoutManager](lc, objc.Sel("alloc"))
	return rv
}

func LayoutManager_Alloc() LayoutManager {
	return LayoutManagerClass.Alloc()
}

func (lc _LayoutManagerClass) New() LayoutManager {
	rv := objc.Call[LayoutManager](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLayoutManager() LayoutManager {
	return LayoutManagerClass.New()
}

// Indicates whether the glyph draws outside its line fragment rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403003-drawsoutsidelinefragmentforglyph?language=objc
func (l_ LayoutManager) DrawsOutsideLineFragmentForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Call[bool](l_, objc.Sel("drawsOutsideLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Sets the location for the first glyph in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402982-setlocation?language=objc
func (l_ LayoutManager) SetLocationForStartOfGlyphRange(location coregraphics.Point, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setLocation:forStartOfGlyphRange:"), location, glyphRange)
}

// Associates a text container with the specified range of glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403241-settextcontainer?language=objc
func (l_ LayoutManager) SetTextContainerForGlyphRange(container ITextContainer, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setTextContainer:forGlyphRange:"), objc.Ptr(container), glyphRange)
}

// Returns the index of the glyph at the specified location in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403112-glyphindexforpoint?language=objc
func (l_ LayoutManager) GlyphIndexForPointInTextContainer(point coregraphics.Point, container ITextContainer) uint {
	rv := objc.Call[uint](l_, objc.Sel("glyphIndexForPoint:inTextContainer:"), point, objc.Ptr(container))
	return rv
}

// Returns the usage rectangle for the line fragment and (optionally) returns the entire range of glyphs in that fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403193-linefragmentusedrectforglyphatin?language=objc
func (l_ LayoutManager) LineFragmentUsedRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("lineFragmentUsedRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

// Returns the range of truncated glyphs for a line fragment that contains the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403203-truncatedglyphrangeinlinefragmen?language=objc
func (l_ LayoutManager) TruncatedGlyphRangeInLineFragmentForGlyphAtIndex(glyphIndex uint) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("truncatedGlyphRangeInLineFragmentForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Updates the information necessary to manage text view objects for the specified text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403229-textcontainerchangedtextview?language=objc
func (l_ LayoutManager) TextContainerChangedTextView(container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("textContainerChangedTextView:"), objc.Ptr(container))
}

// Returns the dictionary of temporary attributes for the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403033-temporaryattributesatcharacterin?language=objc
func (l_ LayoutManager) TemporaryAttributesAtCharacterIndexEffectiveRange(charIndex uint, effectiveCharRange foundation.RangePointer) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](l_, objc.Sel("temporaryAttributesAtCharacterIndex:effectiveRange:"), charIndex, effectiveCharRange)
	return rv
}

// Sets the bounds and container for the extra line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403071-setextralinefragmentrect?language=objc
func (l_ LayoutManager) SetExtraLineFragmentRectUsedRectTextContainer(fragmentRect coregraphics.Rect, usedRect coregraphics.Rect, container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("setExtraLineFragmentRect:usedRect:textContainer:"), fragmentRect, usedRect, objc.Ptr(container))
}

// Appends one or more temporary attributes to the attributes dictionary of the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403250-addtemporaryattributes?language=objc
func (l_ LayoutManager) AddTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("addTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Sets the bounding rectangle that encloses the specified text block and glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402991-setboundsrect?language=objc
func (l_ LayoutManager) SetBoundsRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setBoundsRect:forTextBlock:glyphRange:"), rect, objc.Ptr(block), glyphRange)
}

// Indicates whether the first responder in the specified window is a text view for the layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403026-layoutmanagerownsfirstresponderi?language=objc
func (l_ LayoutManager) LayoutManagerOwnsFirstResponderInWindow(window IWindow) bool {
	rv := objc.Call[bool](l_, objc.Sel("layoutManagerOwnsFirstResponderInWindow:"), objc.Ptr(window))
	return rv
}

// Forces the layout manager to generate glyphs for the specified glyph range if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403219-ensureglyphsforglyphrange?language=objc
func (l_ LayoutManager) EnsureGlyphsForGlyphRange(glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("ensureGlyphsForGlyphRange:"), glyphRange)
}

// Forces the layout manager to perform layout for the specified character range if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402986-ensurelayoutforcharacterrange?language=objc
func (l_ LayoutManager) EnsureLayoutForCharacterRange(charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("ensureLayoutForCharacterRange:"), charRange)
}

// Sets the visibility of the glyph at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403078-setnotshownattribute?language=objc
func (l_ LayoutManager) SetNotShownAttributeForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Call[objc.Void](l_, objc.Sel("setNotShownAttribute:forGlyphAtIndex:"), flag, glyphIndex)
}

// Forces the layout manager to perform layout for the specified glyph range if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402939-ensurelayoutforglyphrange?language=objc
func (l_ LayoutManager) EnsureLayoutForGlyphRange(glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("ensureLayoutForGlyphRange:"), glyphRange)
}

// Returns the index of the first glyph of the character at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403001-glyphindexforcharacteratindex?language=objc
func (l_ LayoutManager) GlyphIndexForCharacterAtIndex(charIndex uint) uint {
	rv := objc.Call[uint](l_, objc.Sel("glyphIndexForCharacterAtIndex:"), charIndex)
	return rv
}

// Draws a strikethrough for the specified glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403169-drawstrikethroughforglyphrange?language=objc
func (l_ LayoutManager) DrawStrikethroughForGlyphRangeStrikethroughTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, baselineOffset float64, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("drawStrikethroughForGlyphRange:strikethroughType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

// Returns the size of the attachment glyph at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403099-attachmentsizeforglyphatindex?language=objc
func (l_ LayoutManager) AttachmentSizeForGlyphAtIndex(glyphIndex uint) coregraphics.Size {
	rv := objc.Call[coregraphics.Size](l_, objc.Sel("attachmentSizeForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Invalidates display for the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402924-invalidatedisplayforcharacterran?language=objc
func (l_ LayoutManager) InvalidateDisplayForCharacterRange(charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("invalidateDisplayForCharacterRange:"), charRange)
}

// Adds a temporary attribute to the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403005-addtemporaryattribute?language=objc
func (l_ LayoutManager) AddTemporaryAttributeValueForCharacterRange(attrName foundation.AttributedStringKey, value objc.IObject, charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("addTemporaryAttribute:value:forCharacterRange:"), attrName, value, charRange)
}

// Sets one or more temporary attributes for the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403037-settemporaryattributes?language=objc
func (l_ LayoutManager) SetTemporaryAttributesForCharacterRange(attrs map[foundation.AttributedStringKey]objc.IObject, charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setTemporaryAttributes:forCharacterRange:"), attrs, charRange)
}

// Returns the bounding rectangle for the glyphs in the specified text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402980-usedrectfortextcontainer?language=objc
func (l_ LayoutManager) UsedRectForTextContainer(container ITextContainer) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("usedRectForTextContainer:"), objc.Ptr(container))
	return rv
}

// Removes the text container at the specified index and invalidates the layout as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403017-removetextcontaineratindex?language=objc
func (l_ LayoutManager) RemoveTextContainerAtIndex(index uint) {
	objc.Call[objc.Void](l_, objc.Sel("removeTextContainerAtIndex:"), index)
}

// Returns the index of the character that lies beneath the specified point using the specified container's coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403028-characterindexforpoint?language=objc
func (l_ LayoutManager) CharacterIndexForPointInTextContainerFractionOfDistanceBetweenInsertionPoints(point coregraphics.Point, container ITextContainer, partialFraction *float64) uint {
	rv := objc.Call[uint](l_, objc.Sel("characterIndexForPoint:inTextContainer:fractionOfDistanceBetweenInsertionPoints:"), point, objc.Ptr(container), partialFraction)
	return rv
}

// Inserts a text container at the specified index in the list of text containers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403010-inserttextcontainer?language=objc
func (l_ LayoutManager) InsertTextContainerAtIndex(container ITextContainer, index uint) {
	objc.Call[objc.Void](l_, objc.Sel("insertTextContainer:atIndex:"), objc.Ptr(container), index)
}

// Appends the specified text container to the series of text containers where the layout manager arranges text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402946-addtextcontainer?language=objc
func (l_ LayoutManager) AddTextContainer(container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("addTextContainer:"), objc.Ptr(container))
}

// Returns the default baseline offset that the layout manager's typesetter uses for the specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403058-defaultbaselineoffsetforfont?language=objc
func (l_ LayoutManager) DefaultBaselineOffsetForFont(theFont IFont) float64 {
	rv := objc.Call[float64](l_, objc.Sel("defaultBaselineOffsetForFont:"), objc.Ptr(theFont))
	return rv
}

// Returns the rectangle for the line fragment where the glyph lies and (optionally), by reference, the entire range of glyphs in that fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403140-linefragmentrectforglyphatindex?language=objc
func (l_ LayoutManager) LineFragmentRectForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("lineFragmentRectForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

// Draws background marks for the specified glyphs, which must lie completely within a single text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402949-drawbackgroundforglyphrange?language=objc
func (l_ LayoutManager) DrawBackgroundForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("drawBackgroundForGlyphRange:atPoint:"), glyphsToShow, origin)
}

// Forces the layout manager to perform layout for the specified text container if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402967-ensurelayoutfortextcontainer?language=objc
func (l_ LayoutManager) EnsureLayoutForTextContainer(container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("ensureLayoutForTextContainer:"), objc.Ptr(container))
}

// Draws an attachment cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402965-showattachmentcell?language=objc
func (l_ LayoutManager) ShowAttachmentCellInRectCharacterIndex(cell ICell, rect foundation.Rect, attachmentIndex uint) {
	objc.Call[objc.Void](l_, objc.Sel("showAttachmentCell:inRect:characterIndex:"), objc.Ptr(cell), rect, attachmentIndex)
}

// Invalidates a range of glyphs, requiring new layout information, and updates the appropriate regions of any text views that display those glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403118-invalidatedisplayforglyphrange?language=objc
func (l_ LayoutManager) InvalidateDisplayForGlyphRange(glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("invalidateDisplayForGlyphRange:"), glyphRange)
}

// Returns insertion points in bulk for a specified line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403012-getlinefragmentinsertionpointsfo?language=objc
func (l_ LayoutManager) GetLineFragmentInsertionPointsForCharacterAtIndexAlternatePositionsInDisplayOrderPositionsCharacterIndexes(charIndex uint, aFlag bool, dFlag bool, positions *float64, charIndexes *uint) uint {
	rv := objc.Call[uint](l_, objc.Sel("getLineFragmentInsertionPointsForCharacterAtIndex:alternatePositions:inDisplayOrder:positions:characterIndexes:"), charIndex, aFlag, dFlag, positions, charIndexes)
	return rv
}

// Returns the range of glyphs lying within the specified text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403041-glyphrangefortextcontainer?language=objc
func (l_ LayoutManager) GlyphRangeForTextContainer(container ITextContainer) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("glyphRangeForTextContainer:"), objc.Ptr(container))
	return rv
}

// Returns the text container that manages the layout for the specified glyph, causing layout to happen as necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403110-textcontainerforglyphatindex?language=objc
func (l_ LayoutManager) TextContainerForGlyphAtIndexEffectiveRange(glyphIndex uint, effectiveGlyphRange foundation.RangePointer) TextContainer {
	rv := objc.Call[TextContainer](l_, objc.Sel("textContainerForGlyphAtIndex:effectiveRange:"), glyphIndex, effectiveGlyphRange)
	return rv
}

// Returns the range of glyphs that the specified range of characters generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402999-glyphrangeforcharacterrange?language=objc
func (l_ LayoutManager) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange foundation.RangePointer) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}

// Indicates whether the specified glyph exceeds the bounds of the line fragment for its layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402964-setdrawsoutsidelinefragment?language=objc
func (l_ LayoutManager) SetDrawsOutsideLineFragmentForGlyphAtIndex(flag bool, glyphIndex uint) {
	objc.Call[objc.Void](l_, objc.Sel("setDrawsOutsideLineFragment:forGlyphAtIndex:"), flag, glyphIndex)
}

// Returns the bounding rectangle that encloses the specified text block and glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403138-boundsrectfortextblock?language=objc
func (l_ LayoutManager) BoundsRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](l_, objc.Sel("boundsRectForTextBlock:glyphRange:"), objc.Ptr(block), glyphRange)
	return rv
}

// Fills background rectangles with a color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403161-fillbackgroundrectarray?language=objc
func (l_ LayoutManager) FillBackgroundRectArrayCountForCharacterRangeColor(rectArray *coregraphics.Rect, rectCount uint, charRange foundation.Range, color IColor) {
	objc.Call[objc.Void](l_, objc.Sel("fillBackgroundRectArray:count:forCharacterRange:color:"), rectArray, rectCount, charRange, objc.Ptr(color))
}

// Renders the glyphs at the specified positions, using the specified attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/3180379-showcgglyphs?language=objc
func (l_ LayoutManager) ShowCGGlyphsPositionsCountFontTextMatrixAttributesInContext(glyphs *coregraphics.Glyph, positions *coregraphics.Point, glyphCount int, font IFont, textMatrix coregraphics.AffineTransform, attributes map[foundation.AttributedStringKey]objc.IObject, CGContext coregraphics.ContextRef) {
	objc.Call[objc.Void](l_, objc.Sel("showCGGlyphs:positions:count:font:textMatrix:attributes:inContext:"), glyphs, positions, glyphCount, objc.Ptr(font), textMatrix, attributes, CGContext)
}

// Returns the index for the first glyph in the layout manager that isn’t in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403245-firstunlaidglyphindex?language=objc
func (l_ LayoutManager) FirstUnlaidGlyphIndex() uint {
	rv := objc.Call[uint](l_, objc.Sel("firstUnlaidGlyphIndex"))
	return rv
}

// Returns the glyph at the specified index along with information about whether the glyph index is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403132-cgglyphatindex?language=objc
func (l_ LayoutManager) CGGlyphAtIndexIsValidIndex(glyphIndex uint, isValidIndex *bool) coregraphics.Glyph {
	rv := objc.Call[coregraphics.Glyph](l_, objc.Sel("CGGlyphAtIndex:isValidIndex:"), glyphIndex, isValidIndex)
	return rv
}

// Enumerates line fragments intersecting with the specified glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403160-enumeratelinefragmentsforglyphra?language=objc
func (l_ LayoutManager) EnumerateLineFragmentsForGlyphRangeUsingBlock(glyphRange foundation.Range, block func(rect coregraphics.Rect, usedRect coregraphics.Rect, textContainer TextContainer, glyphRange foundation.Range, stop *bool)) {
	objc.Call[objc.Void](l_, objc.Sel("enumerateLineFragmentsForGlyphRange:usingBlock:"), glyphRange, block)
}

// Invalidates the layout information, and possibly glyphs, for the specified text container and all subsequent text container objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403091-textcontainerchangedgeometry?language=objc
func (l_ LayoutManager) TextContainerChangedGeometry(container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("textContainerChangedGeometry:"), objc.Ptr(container))
}

// Replaces the layout manager's current text storage object with the specified object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403227-replacetextstorage?language=objc
func (l_ LayoutManager) ReplaceTextStorage(newTextStorage ITextStorage) {
	objc.Call[objc.Void](l_, objc.Sel("replaceTextStorage:"), objc.Ptr(newTextStorage))
}

// Stores the initial glyphs and glyph properties for a character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403030-setglyphs?language=objc
func (l_ LayoutManager) SetGlyphsPropertiesCharacterIndexesFontForGlyphRange(glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setGlyphs:properties:characterIndexes:font:forGlyphRange:"), glyphs, props, charIndexes, objc.Ptr(aFont), glyphRange)
}

// Enumerates enclosing rectangles for the specified glyph range in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403021-enumerateenclosingrectsforglyphr?language=objc
func (l_ LayoutManager) EnumerateEnclosingRectsForGlyphRangeWithinSelectedGlyphRangeInTextContainerUsingBlock(glyphRange foundation.Range, selectedRange foundation.Range, textContainer ITextContainer, block func(rect coregraphics.Rect, stop *bool)) {
	objc.Call[objc.Void](l_, objc.Sel("enumerateEnclosingRectsForGlyphRange:withinSelectedGlyphRange:inTextContainer:usingBlock:"), glyphRange, selectedRange, objc.Ptr(textContainer), block)
}

// Sets the layout rectangle that encloses the specified text block and glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402929-setlayoutrect?language=objc
func (l_ LayoutManager) SetLayoutRectForTextBlockGlyphRange(rect foundation.Rect, block ITextBlock, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setLayoutRect:forTextBlock:glyphRange:"), rect, objc.Ptr(block), glyphRange)
}

// Notifies the layout manager when an edit action changes the contents of its text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403065-processeditingfortextstorage?language=objc
func (l_ LayoutManager) ProcessEditingForTextStorageEditedRangeChangeInLengthInvalidatedRange(textStorage ITextStorage, editMask TextStorageEditActions, newCharRange foundation.Range, delta int, invalidatedCharRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("processEditingForTextStorage:edited:range:changeInLength:invalidatedRange:"), objc.Ptr(textStorage), editMask, newCharRange, delta, invalidatedCharRange)
}

// Calculates subranges to underline for the specified glyphs and draws the underlining as appropriate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403114-underlineglyphrange?language=objc
func (l_ LayoutManager) UnderlineGlyphRangeUnderlineTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("underlineGlyphRange:underlineType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, lineRect, lineGlyphRange, containerOrigin)
}

// Returns the index for the first character in the layout manager that isn’t in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403067-firstunlaidcharacterindex?language=objc
func (l_ LayoutManager) FirstUnlaidCharacterIndex() uint {
	rv := objc.Call[uint](l_, objc.Sel("firstUnlaidCharacterIndex"))
	return rv
}

// Draws the specified glyphs, which must lie completely within a single text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403158-drawglyphsforglyphrange?language=objc
func (l_ LayoutManager) DrawGlyphsForGlyphRangeAtPoint(glyphsToShow foundation.Range, origin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("drawGlyphsForGlyphRange:atPoint:"), glyphsToShow, origin)
}

// Associates the line fragment bounds for the specified range of glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402935-setlinefragmentrect?language=objc
func (l_ LayoutManager) SetLineFragmentRectForGlyphRangeUsedRect(fragmentRect coregraphics.Rect, glyphRange foundation.Range, usedRect coregraphics.Rect) {
	objc.Call[objc.Void](l_, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:"), fragmentRect, glyphRange, usedRect)
}

// Indicates whether the glyph at the specified index has a visible representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402931-notshownattributeforglyphatindex?language=objc
func (l_ LayoutManager) NotShownAttributeForGlyphAtIndex(glyphIndex uint) bool {
	rv := objc.Call[bool](l_, objc.Sel("notShownAttributeForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Calculates and draws strikethrough for the specified glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403009-strikethroughglyphrange?language=objc
func (l_ LayoutManager) StrikethroughGlyphRangeStrikethroughTypeLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, strikethroughVal UnderlineStyle, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("strikethroughGlyphRange:strikethroughType:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, strikethroughVal, lineRect, lineGlyphRange, containerOrigin)
}

// Invalidates and adjusts the glyphs in the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403171-invalidateglyphsforcharacterrang?language=objc
func (l_ LayoutManager) InvalidateGlyphsForCharacterRangeChangeInLengthActualCharacterRange(charRange foundation.Range, delta int, actualCharRange foundation.RangePointer) {
	objc.Call[objc.Void](l_, objc.Sel("invalidateGlyphsForCharacterRange:changeInLength:actualCharacterRange:"), charRange, delta, actualCharRange)
}

// Forces the layout manager to generate glyphs for the specified character range if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403189-ensureglyphsforcharacterrange?language=objc
func (l_ LayoutManager) EnsureGlyphsForCharacterRange(charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("ensureGlyphsForCharacterRange:"), charRange)
}

// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403183-glyphrangeforboundingrectwithout?language=objc
func (l_ LayoutManager) GlyphRangeForBoundingRectWithoutAdditionalLayoutInTextContainer(bounds coregraphics.Rect, container ITextContainer) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("glyphRangeForBoundingRectWithoutAdditionalLayout:inTextContainer:"), bounds, objc.Ptr(container))
	return rv
}

// Returns an array of text ruler objects for the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403163-rulermarkersfortextview?language=objc
func (l_ LayoutManager) RulerMarkersForTextViewParagraphStyleRuler(view ITextView, style IParagraphStyle, ruler IRulerView) []RulerMarker {
	rv := objc.Call[[]RulerMarker](l_, objc.Sel("rulerMarkersForTextView:paragraphStyle:ruler:"), objc.Ptr(view), objc.Ptr(style), objc.Ptr(ruler))
	return rv
}

// Sets the size to use when drawing a glyph that represents an attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403047-setattachmentsize?language=objc
func (l_ LayoutManager) SetAttachmentSizeForGlyphRange(attachmentSize coregraphics.Size, glyphRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

// Returns the default line height for a line of text that uses a specified font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403007-defaultlineheightforfont?language=objc
func (l_ LayoutManager) DefaultLineHeightForFont(theFont IFont) float64 {
	rv := objc.Call[float64](l_, objc.Sel("defaultLineHeightForFont:"), objc.Ptr(theFont))
	return rv
}

// Forces the layout manager to perform layout for the specified area in the specified text container if it hasn’t already. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402962-ensurelayoutforboundingrect?language=objc
func (l_ LayoutManager) EnsureLayoutForBoundingRectInTextContainer(bounds coregraphics.Rect, container ITextContainer) {
	objc.Call[objc.Void](l_, objc.Sel("ensureLayoutForBoundingRect:inTextContainer:"), bounds, objc.Ptr(container))
}

// Returns the glyph property of the glyph at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403014-propertyforglyphatindex?language=objc
func (l_ LayoutManager) PropertyForGlyphAtIndex(glyphIndex uint) GlyphProperty {
	rv := objc.Call[GlyphProperty](l_, objc.Sel("propertyForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the range of characters that correspond to the glyphs in the specified glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403225-characterrangeforglyphrange?language=objc
func (l_ LayoutManager) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange foundation.RangePointer) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}

// Returns the index in the text storage for the first character of the specified glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402944-characterindexforglyphatindex?language=objc
func (l_ LayoutManager) CharacterIndexForGlyphAtIndex(glyphIndex uint) uint {
	rv := objc.Call[uint](l_, objc.Sel("characterIndexForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the rectangle for the layout of the specified text block and glyph range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403201-layoutrectfortextblock?language=objc
func (l_ LayoutManager) LayoutRectForTextBlockGlyphRange(block ITextBlock, glyphRange foundation.Range) foundation.Rect {
	rv := objc.Call[foundation.Rect](l_, objc.Sel("layoutRectForTextBlock:glyphRange:"), objc.Ptr(block), glyphRange)
	return rv
}

// Returns the location for the specified glyph within its line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403239-locationforglyphatindex?language=objc
func (l_ LayoutManager) LocationForGlyphAtIndex(glyphIndex uint) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("locationForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the smallest contiguous range for glyphs lying wholly or partially within the specified rectangle of the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403053-glyphrangeforboundingrect?language=objc
func (l_ LayoutManager) GlyphRangeForBoundingRectInTextContainer(bounds coregraphics.Rect, container ITextContainer) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("glyphRangeForBoundingRect:inTextContainer:"), bounds, objc.Ptr(container))
	return rv
}

// Returns the value for the temporary attribute of a character, and the maximum range it applies to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403223-temporaryattribute?language=objc
func (l_ LayoutManager) TemporaryAttributeAtCharacterIndexLongestEffectiveRangeInRange(attrName foundation.AttributedStringKey, location uint, range_ foundation.RangePointer, rangeLimit foundation.Range) objc.Object {
	rv := objc.Call[objc.Object](l_, objc.Sel("temporaryAttribute:atCharacterIndex:longestEffectiveRange:inRange:"), attrName, location, range_, rangeLimit)
	return rv
}

// Returns the indexes for the first character and glyph that have invalid layout information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403187-getfirstunlaidcharacterindex?language=objc
func (l_ LayoutManager) GetFirstUnlaidCharacterIndexGlyphIndex(charIndex *uint, glyphIndex *uint) {
	objc.Call[objc.Void](l_, objc.Sel("getFirstUnlaidCharacterIndex:glyphIndex:"), charIndex, glyphIndex)
}

// Returns the accessory view that the text system uses for its ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403130-ruleraccessoryviewfortextview?language=objc
func (l_ LayoutManager) RulerAccessoryViewForTextViewParagraphStyleRulerEnabled(view ITextView, style IParagraphStyle, ruler IRulerView, isEnabled bool) View {
	rv := objc.Call[View](l_, objc.Sel("rulerAccessoryViewForTextView:paragraphStyle:ruler:enabled:"), objc.Ptr(view), objc.Ptr(style), objc.Ptr(ruler), isEnabled)
	return rv
}

// Returns the range of displayable glyphs that surround the glyph at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403152-rangeofnominallyspacedglyphscont?language=objc
func (l_ LayoutManager) RangeOfNominallySpacedGlyphsContainingIndex(glyphIndex uint) foundation.Range {
	rv := objc.Call[foundation.Range](l_, objc.Sel("rangeOfNominallySpacedGlyphsContainingIndex:"), glyphIndex)
	return rv
}

// Returns the bounding rectangle for the specified glyphs in a container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403255-boundingrectforglyphrange?language=objc
func (l_ LayoutManager) BoundingRectForGlyphRangeInTextContainer(glyphRange foundation.Range, container ITextContainer) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("boundingRectForGlyphRange:inTextContainer:"), glyphRange, objc.Ptr(container))
	return rv
}

// Returns the fraction of the distance between the glyph at the specified point and the next glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403097-fractionofdistancethroughglyphfo?language=objc
func (l_ LayoutManager) FractionOfDistanceThroughGlyphForPointInTextContainer(point coregraphics.Point, container ITextContainer) float64 {
	rv := objc.Call[float64](l_, objc.Sel("fractionOfDistanceThroughGlyphForPoint:inTextContainer:"), point, objc.Ptr(container))
	return rv
}

// Removes a temporary attribute from the list of attributes for the specified character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403087-removetemporaryattribute?language=objc
func (l_ LayoutManager) RemoveTemporaryAttributeForCharacterRange(attrName foundation.AttributedStringKey, charRange foundation.Range) {
	objc.Call[objc.Void](l_, objc.Sel("removeTemporaryAttribute:forCharacterRange:"), attrName, charRange)
}

// Draws underlining for the glyphs in a specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403079-drawunderlineforglyphrange?language=objc
func (l_ LayoutManager) DrawUnderlineForGlyphRangeUnderlineTypeBaselineOffsetLineFragmentRectLineFragmentGlyphRangeContainerOrigin(glyphRange foundation.Range, underlineVal UnderlineStyle, baselineOffset float64, lineRect coregraphics.Rect, lineGlyphRange foundation.Range, containerOrigin coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("drawUnderlineForGlyphRange:underlineType:baselineOffset:lineFragmentRect:lineFragmentGlyphRange:containerOrigin:"), glyphRange, underlineVal, baselineOffset, lineRect, lineGlyphRange, containerOrigin)
}

// Indicates whether the specified index refers to a valid glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402950-isvalidglyphindex?language=objc
func (l_ LayoutManager) IsValidGlyphIndex(glyphIndex uint) bool {
	rv := objc.Call[bool](l_, objc.Sel("isValidGlyphIndex:"), glyphIndex)
	return rv
}

// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc
func (l_ LayoutManager) ShowsInvisibleCharacters() bool {
	rv := objc.Call[bool](l_, objc.Sel("showsInvisibleCharacters"))
	return rv
}

// A Boolean value that indicates whether to substitute visible glyphs for whitespace and other typically invisible characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403254-showsinvisiblecharacters?language=objc
func (l_ LayoutManager) SetShowsInvisibleCharacters(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setShowsInvisibleCharacters:"), value)
}

// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc
func (l_ LayoutManager) ShowsControlCharacters() bool {
	rv := objc.Call[bool](l_, objc.Sel("showsControlCharacters"))
	return rv
}

// A Boolean value that indicates whether the layout manager substitutes visible glyphs for control characters in the layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402912-showscontrolcharacters?language=objc
func (l_ LayoutManager) SetShowsControlCharacters(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setShowsControlCharacters:"), value)
}

// The rectangle for the extra line fragment at the end of a document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403175-extralinefragmentrect?language=objc
func (l_ LayoutManager) ExtraLineFragmentRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("extraLineFragmentRect"))
	return rv
}

// The text view that contains the first glyph in the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403089-textviewforbeginningofselection?language=objc
func (l_ LayoutManager) TextViewForBeginningOfSelection() TextView {
	rv := objc.Call[TextView](l_, objc.Sel("textViewForBeginningOfSelection"))
	return rv
}

// A Boolean value that indicates whether the layout manager currently has any areas of noncontiguous layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403207-hasnoncontiguouslayout?language=objc
func (l_ LayoutManager) HasNonContiguousLayout() bool {
	rv := objc.Call[bool](l_, objc.Sel("hasNonContiguousLayout"))
	return rv
}

// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc
func (l_ LayoutManager) UsesDefaultHyphenation() bool {
	rv := objc.Call[bool](l_, objc.Sel("usesDefaultHyphenation"))
	return rv
}

// A Boolean value that indicates whether the layout manager uses the default hyphenation rules to wrap lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/3180380-usesdefaulthyphenation?language=objc
func (l_ LayoutManager) SetUsesDefaultHyphenation(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setUsesDefaultHyphenation:"), value)
}

// The number of glyphs in the layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402937-numberofglyphs?language=objc
func (l_ LayoutManager) NumberOfGlyphs() uint {
	rv := objc.Call[uint](l_, objc.Sel("numberOfGlyphs"))
	return rv
}

// The current typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403205-typesetter?language=objc
func (l_ LayoutManager) Typesetter() Typesetter {
	rv := objc.Call[Typesetter](l_, objc.Sel("typesetter"))
	return rv
}

// The current typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403205-typesetter?language=objc
func (l_ LayoutManager) SetTypesetter(value ITypesetter) {
	objc.Call[objc.Void](l_, objc.Sel("setTypesetter:"), objc.Ptr(value))
}

// The layout manager’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc
func (l_ LayoutManager) Delegate() LayoutManagerDelegateWrapper {
	rv := objc.Call[LayoutManagerDelegateWrapper](l_, objc.Sel("delegate"))
	return rv
}

// The layout manager’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc
func (l_ LayoutManager) SetDelegate(value PLayoutManagerDelegate) {
	po0 := objc.WrapAsProtocol("NSLayoutManagerDelegate", value)
	objc.SetAssociatedObject(l_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), po0)
}

// The layout manager’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402920-delegate?language=objc
func (l_ LayoutManager) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// A Boolean value that indicates whether the layout manager allows noncontiguous layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc
func (l_ LayoutManager) AllowsNonContiguousLayout() bool {
	rv := objc.Call[bool](l_, objc.Sel("allowsNonContiguousLayout"))
	return rv
}

// A Boolean value that indicates whether the layout manager allows noncontiguous layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403197-allowsnoncontiguouslayout?language=objc
func (l_ LayoutManager) SetAllowsNonContiguousLayout(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setAllowsNonContiguousLayout:"), value)
}

// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc
func (l_ LayoutManager) LimitsLayoutForSuspiciousContents() bool {
	rv := objc.Call[bool](l_, objc.Sel("limitsLayoutForSuspiciousContents"))
	return rv
}

// A Boolean value that indicates whether the layout manager avoids laying out unusually long or suspicious input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/3021179-limitslayoutforsuspiciouscontent?language=objc
func (l_ LayoutManager) SetLimitsLayoutForSuspiciousContents(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setLimitsLayoutForSuspiciousContents:"), value)
}

// The glyph generator that the layout manager uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403069-glyphgenerator?language=objc
func (l_ LayoutManager) GlyphGenerator() GlyphGenerator {
	rv := objc.Call[GlyphGenerator](l_, objc.Sel("glyphGenerator"))
	return rv
}

// The glyph generator that the layout manager uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403069-glyphgenerator?language=objc
func (l_ LayoutManager) SetGlyphGenerator(value IGlyphGenerator) {
	objc.Call[objc.Void](l_, objc.Sel("setGlyphGenerator:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the layout manager uses the leading of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc
func (l_ LayoutManager) UsesFontLeading() bool {
	rv := objc.Call[bool](l_, objc.Sel("usesFontLeading"))
	return rv
}

// A Boolean value that indicates whether the layout manager uses the leading of the font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403156-usesfontleading?language=objc
func (l_ LayoutManager) SetUsesFontLeading(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setUsesFontLeading:"), value)
}

// The default typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403199-typesetterbehavior?language=objc
func (l_ LayoutManager) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Call[TypesetterBehavior](l_, objc.Sel("typesetterBehavior"))
	return rv
}

// The default typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403199-typesetterbehavior?language=objc
func (l_ LayoutManager) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Call[objc.Void](l_, objc.Sel("setTypesetterBehavior:"), value)
}

// The current text containers of the layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403144-textcontainers?language=objc
func (l_ LayoutManager) TextContainers() []TextContainer {
	rv := objc.Call[[]TextContainer](l_, objc.Sel("textContainers"))
	return rv
}

// The first text view in the layout manager’s series of text views. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402995-firsttextview?language=objc
func (l_ LayoutManager) FirstTextView() TextView {
	rv := objc.Call[TextView](l_, objc.Sel("firstTextView"))
	return rv
}

// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc
func (l_ LayoutManager) BackgroundLayoutEnabled() bool {
	rv := objc.Call[bool](l_, objc.Sel("backgroundLayoutEnabled"))
	return rv
}

// A Boolean value that indicates whether the layout manager generates glyphs and lays them out when the app's run loop is idle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1402952-backgroundlayoutenabled?language=objc
func (l_ LayoutManager) SetBackgroundLayoutEnabled(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setBackgroundLayoutEnabled:"), value)
}

// The rectangle that encloses the insertion point in the extra line fragment rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1402988-extralinefragmentusedrect?language=objc
func (l_ LayoutManager) ExtraLineFragmentUsedRect() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](l_, objc.Sel("extraLineFragmentUsedRect"))
	return rv
}

// The text storage object that contains the content to lay out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403015-textstorage?language=objc
func (l_ LayoutManager) TextStorage() TextStorage {
	rv := objc.Call[TextStorage](l_, objc.Sel("textStorage"))
	return rv
}

// The text storage object that contains the content to lay out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403015-textstorage?language=objc
func (l_ LayoutManager) SetTextStorage(value ITextStorage) {
	objc.Call[objc.Void](l_, objc.Sel("setTextStorage:"), objc.Ptr(value))
}

// The default amount of scaling to apply when an attachment image is too large to fit in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403045-defaultattachmentscaling?language=objc
func (l_ LayoutManager) DefaultAttachmentScaling() ImageScaling {
	rv := objc.Call[ImageScaling](l_, objc.Sel("defaultAttachmentScaling"))
	return rv
}

// The default amount of scaling to apply when an attachment image is too large to fit in a text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslayoutmanager/1403045-defaultattachmentscaling?language=objc
func (l_ LayoutManager) SetDefaultAttachmentScaling(value ImageScaling) {
	objc.Call[objc.Void](l_, objc.Sel("setDefaultAttachmentScaling:"), value)
}

// The text container for the extra line fragment rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nslayoutmanager/1403165-extralinefragmenttextcontainer?language=objc
func (l_ LayoutManager) ExtraLineFragmentTextContainer() TextContainer {
	rv := objc.Call[TextContainer](l_, objc.Sel("extraLineFragmentTextContainer"))
	return rv
}
