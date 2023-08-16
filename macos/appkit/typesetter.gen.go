// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Typesetter] class.
var TypesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}

type _TypesetterClass struct {
	objc.Class
}

// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	objc.IObject
	SetLocationWithAdvancementsForStartOfGlyphRange(location foundation.Point, advancements *float64, glyphRange foundation.Range)
	ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool
	GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect foundation.RectPointer, lineFragmentUsedRect foundation.RectPointer, remainingRect foundation.RectPointer, startingGlyphIndex uint, proposedRect foundation.Rect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64)
	SubstituteFontForFont(originalFont IFont) Font
	SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.Range)
	SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.Range)
	HyphenCharacterForGlyphAtIndex(glyphIndex uint) objc.Object
	EndLineWithGlyphRange(lineGlyphRange foundation.Range)
	WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect foundation.RectPointer, glyphRange foundation.Range, usedRect foundation.RectPointer, baselineOffset *float64)
	ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool
	LayoutParagraphAtPoint(lineFragmentOrigin foundation.PointPointer) uint
	EndParagraph()
	HyphenationFactorForGlyphAtIndex(glyphIndex uint) float64
	BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange foundation.RangePointer) foundation.Range
	LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.Range)
	SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range)
	BeginLineWithGlyphAtIndex(glyphIndex uint)
	ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) TextTab
	BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64
	SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset float64)
	LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range
	SetAttachmentSizeForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range)
	ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64
	CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange foundation.RangePointer) foundation.Range
	SetBidiLevelsForGlyphRange(levels *uint8, glyphRange foundation.Range)
	BeginParagraph()
	BidiProcessingEnabled() bool
	SetBidiProcessingEnabled(value bool)
	AttributedString() foundation.AttributedString
	SetAttributedString(value foundation.IAttributedString)
	ParagraphGlyphRange() foundation.Range
	ParagraphSeparatorCharacterRange() foundation.Range
	LayoutManager() LayoutManager
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	UsesFontLeading() bool
	SetUsesFontLeading(value bool)
	TypesetterBehavior() TypesetterBehavior
	SetTypesetterBehavior(value TypesetterBehavior)
	TextContainers() []TextContainer
	HyphenationFactor() float64
	SetHyphenationFactor(value float64)
	ParagraphCharacterRange() foundation.Range
	CurrentParagraphStyle() ParagraphStyle
	CurrentTextContainer() TextContainer
	ParagraphSeparatorGlyphRange() foundation.Range
	AttributesForExtraLineFragment() map[foundation.AttributedStringKey]objc.Object
}

// An abstract class that performs various type layout tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter?language=objc
type Typesetter struct {
	objc.Object
}

func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{
		Object: objc.ObjectFrom(ptr),
	}
}

func (tc _TypesetterClass) Alloc() Typesetter {
	rv := objc.Call[Typesetter](tc, objc.Sel("alloc"))
	return rv
}

func Typesetter_Alloc() Typesetter {
	return TypesetterClass.Alloc()
}

func (tc _TypesetterClass) New() Typesetter {
	rv := objc.Call[Typesetter](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTypesetter() Typesetter {
	return TypesetterClass.New()
}

func (t_ Typesetter) Init() Typesetter {
	rv := objc.Call[Typesetter](t_, objc.Sel("init"))
	return rv
}

// Sets the location where the specified glyphs are laid out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533806-setlocation?language=objc
func (t_ Typesetter) SetLocationWithAdvancementsForStartOfGlyphRange(location foundation.Point, advancements *float64, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setLocation:withAdvancements:forStartOfGlyphRange:"), location, advancements, glyphRange)
}

// Returns whether the line being laid out should be broken by a word break at the specified character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1532732-shouldbreaklinebywordbeforechara?language=objc
func (t_ Typesetter) ShouldBreakLineByWordBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldBreakLineByWordBeforeCharacterAtIndex:"), charIndex)
	return rv
}

// Calculates line fragment rectangle, line fragment used rectangle, and remaining rectangle for a line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535809-getlinefragmentrect?language=objc
func (t_ Typesetter) GetLineFragmentRectUsedRectRemainingRectForStartingGlyphAtIndexProposedRectLineSpacingParagraphSpacingBeforeParagraphSpacingAfter(lineFragmentRect foundation.RectPointer, lineFragmentUsedRect foundation.RectPointer, remainingRect foundation.RectPointer, startingGlyphIndex uint, proposedRect foundation.Rect, lineSpacing float64, paragraphSpacingBefore float64, paragraphSpacingAfter float64) {
	objc.Call[objc.Void](t_, objc.Sel("getLineFragmentRect:usedRect:remainingRect:forStartingGlyphAtIndex:proposedRect:lineSpacing:paragraphSpacingBefore:paragraphSpacingAfter:"), lineFragmentRect, lineFragmentUsedRect, remainingRect, startingGlyphIndex, proposedRect, lineSpacing, paragraphSpacingBefore, paragraphSpacingAfter)
}

// Returns a screen font suitable for use in place of a given font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526766-substitutefontforfont?language=objc
func (t_ Typesetter) SubstituteFontForFont(originalFont IFont) Font {
	rv := objc.Call[Font](t_, objc.Sel("substituteFontForFont:"), objc.Ptr(originalFont))
	return rv
}

// Sets whether the specified glyphs are not shown. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530373-setnotshownattribute?language=objc
func (t_ Typesetter) SetNotShownAttributeForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setNotShownAttribute:forGlyphRange:"), flag, glyphRange)
}

// Sets whether to force the layout manager to invalidate the specified portion of the glyph cache when invalidating layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1529806-sethardinvalidation?language=objc
func (t_ Typesetter) SetHardInvalidationForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setHardInvalidation:forGlyphRange:"), flag, glyphRange)
}

// Returns the hyphen character to be inserted after the specified glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1531874-hyphencharacterforglyphatindex?language=objc
func (t_ Typesetter) HyphenCharacterForGlyphAtIndex(glyphIndex uint) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("hyphenCharacterForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Sets up layout parameters at the end of a line during typesetting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1525531-endlinewithglyphrange?language=objc
func (t_ Typesetter) EndLineWithGlyphRange(lineGlyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("endLineWithGlyphRange:"), lineGlyphRange)
}

// Called by the typesetter just prior to storing the actual line fragment rectangle location in the layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530261-willsetlinefragmentrect?language=objc
func (t_ Typesetter) WillSetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(lineRect foundation.RectPointer, glyphRange foundation.Range, usedRect foundation.RectPointer, baselineOffset *float64) {
	objc.Call[objc.Void](t_, objc.Sel("willSetLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), lineRect, glyphRange, usedRect, baselineOffset)
}

// Returns whether the line being laid out should be broken by hyphenating at the specified character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534525-shouldbreaklinebyhyphenatingbefo?language=objc
func (t_ Typesetter) ShouldBreakLineByHyphenatingBeforeCharacterAtIndex(charIndex uint) bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), charIndex)
	return rv
}

// Lays out glyphs in the current glyph range until the next paragraph separator is reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1528231-layoutparagraphatpoint?language=objc
func (t_ Typesetter) LayoutParagraphAtPoint(lineFragmentOrigin foundation.PointPointer) uint {
	rv := objc.Call[uint](t_, objc.Sel("layoutParagraphAtPoint:"), lineFragmentOrigin)
	return rv
}

// Sets up layout parameters at the end of a paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526802-endparagraph?language=objc
func (t_ Typesetter) EndParagraph() {
	objc.Call[objc.Void](t_, objc.Sel("endParagraph"))
}

// Returns the hyphenation factor in effect at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533345-hyphenationfactorforglyphatindex?language=objc
func (t_ Typesetter) HyphenationFactorForGlyphAtIndex(glyphIndex uint) float64 {
	rv := objc.Call[float64](t_, objc.Sel("hyphenationFactorForGlyphAtIndex:"), glyphIndex)
	return rv
}

// Returns the bounding rectangle for the specified control glyph with the specified parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535355-boundingboxforcontrolglyphatinde?language=objc
func (t_ Typesetter) BoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), glyphIndex, objc.Ptr(textContainer), proposedRect, glyphPosition, charIndex)
	return rv
}

// Returns the range for the glyphs mapped to the characters of the text store in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1531461-glyphrangeforcharacterrange?language=objc
func (t_ Typesetter) GlyphRangeForCharacterRangeActualCharacterRange(charRange foundation.Range, actualCharRange foundation.RangePointer) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("glyphRangeForCharacterRange:actualCharacterRange:"), charRange, actualCharRange)
	return rv
}

// Returns the line spacing in effect following the specified glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1532355-linespacingafterglyphatindex?language=objc
func (t_ Typesetter) LineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.Call[float64](t_, objc.Sel("lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Sets whether the specified glyphs exceed the bounds of the line fragment in which they are laid out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526312-setdrawsoutsidelinefragment?language=objc
func (t_ Typesetter) SetDrawsOutsideLineFragmentForGlyphRange(flag bool, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setDrawsOutsideLineFragment:forGlyphRange:"), flag, glyphRange)
}

// Sets the current glyph range being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530393-setparagraphglyphrange?language=objc
func (t_ Typesetter) SetParagraphGlyphRangeSeparatorGlyphRange(paragraphRange foundation.Range, paragraphSeparatorRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setParagraphGlyphRange:separatorGlyphRange:"), paragraphRange, paragraphSeparatorRange)
}

// Sets up layout parameters at the beginning of a line during typesetting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534305-beginlinewithglyphatindex?language=objc
func (t_ Typesetter) BeginLineWithGlyphAtIndex(glyphIndex uint) {
	objc.Call[objc.Void](t_, objc.Sel("beginLineWithGlyphAtIndex:"), glyphIndex)
}

// Returns the number of points of space—added before a paragraph—that is in effect before the specified glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1525163-paragraphspacingbeforeglyphatind?language=objc
func (t_ Typesetter) ParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.Call[float64](t_, objc.Sel("paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Returns the text tab next closest to a given glyph location within the given parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1524483-texttabforglyphlocation?language=objc
func (t_ Typesetter) TextTabForGlyphLocationWritingDirectionMaxLocation(glyphLocation float64, direction WritingDirection, maxLocation float64) TextTab {
	rv := objc.Call[TextTab](t_, objc.Sel("textTabForGlyphLocation:writingDirection:maxLocation:"), glyphLocation, direction, maxLocation)
	return rv
}

// Returns the interglyph spacing in the specified range when sent to a printer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530764-printingadjustmentinlayoutmanage?language=objc
func (tc _TypesetterClass) PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr ILayoutManager, nominallySpacedGlyphsRange foundation.Range, packedGlyphs *uint8, packedGlyphsCount uint) foundation.Size {
	rv := objc.Call[foundation.Size](tc, objc.Sel("printingAdjustmentInLayoutManager:forNominallySpacedGlyphRange:packedGlyphs:count:"), objc.Ptr(layoutMgr), nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
	return rv
}

// Returns the interglyph spacing in the specified range when sent to a printer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530764-printingadjustmentinlayoutmanage?language=objc
func Typesetter_PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr ILayoutManager, nominallySpacedGlyphsRange foundation.Range, packedGlyphs *uint8, packedGlyphsCount uint) foundation.Size {
	return TypesetterClass.PrintingAdjustmentInLayoutManagerForNominallySpacedGlyphRangePackedGlyphsCount(layoutMgr, nominallySpacedGlyphsRange, packedGlyphs, packedGlyphsCount)
}

// Returns the distance from the bottom of the line fragment rectangle in which the glyph resides to the glyph baseline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535394-baselineoffsetinlayoutmanager?language=objc
func (t_ Typesetter) BaselineOffsetInLayoutManagerGlyphIndex(layoutMgr ILayoutManager, glyphIndex uint) float64 {
	rv := objc.Call[float64](t_, objc.Sel("baselineOffsetInLayoutManager:glyphIndex:"), objc.Ptr(layoutMgr), glyphIndex)
	return rv
}

// Sets the line fragment rectangle where the specified glyphs are laid out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534015-setlinefragmentrect?language=objc
func (t_ Typesetter) SetLineFragmentRectForGlyphRangeUsedRectBaselineOffset(fragmentRect foundation.Rect, glyphRange foundation.Range, usedRect foundation.Rect, baselineOffset float64) {
	objc.Call[objc.Void](t_, objc.Sel("setLineFragmentRect:forGlyphRange:usedRect:baselineOffset:"), fragmentRect, glyphRange, usedRect, baselineOffset)
}

// Lays out characters in the given character range for the specified layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535876-layoutcharactersinrange?language=objc
func (t_ Typesetter) LayoutCharactersInRangeForLayoutManagerMaximumNumberOfLineFragments(characterRange foundation.Range, layoutManager ILayoutManager, maxNumLines uint) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("layoutCharactersInRange:forLayoutManager:maximumNumberOfLineFragments:"), characterRange, objc.Ptr(layoutManager), maxNumLines)
	return rv
}

// Returns a shared instance of a reentrant typesetter that implements typesetting with the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530659-sharedsystemtypesetterforbehavio?language=objc
func (tc _TypesetterClass) SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.Object {
	rv := objc.Call[objc.Object](tc, objc.Sel("sharedSystemTypesetterForBehavior:"), behavior)
	return rv
}

// Returns a shared instance of a reentrant typesetter that implements typesetting with the specified behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1530659-sharedsystemtypesetterforbehavio?language=objc
func Typesetter_SharedSystemTypesetterForBehavior(behavior TypesetterBehavior) objc.Object {
	return TypesetterClass.SharedSystemTypesetterForBehavior(behavior)
}

// Sets the size the specified glyphs (assumed to be attachments) will be asked to draw themselves at. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1532282-setattachmentsize?language=objc
func (t_ Typesetter) SetAttachmentSizeForGlyphRange(attachmentSize foundation.Size, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setAttachmentSize:forGlyphRange:"), attachmentSize, glyphRange)
}

// Returns the paragraph spacing that is in effect after the specified glyph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534661-paragraphspacingafterglyphatinde?language=objc
func (t_ Typesetter) ParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.Call[float64](t_, objc.Sel("paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), glyphIndex, rect)
	return rv
}

// Returns the range for the characters in the receiver’s text store that are mapped to the specified glyphs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1524391-characterrangeforglyphrange?language=objc
func (t_ Typesetter) CharacterRangeForGlyphRangeActualGlyphRange(glyphRange foundation.Range, actualGlyphRange foundation.RangePointer) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("characterRangeForGlyphRange:actualGlyphRange:"), glyphRange, actualGlyphRange)
	return rv
}

// Sets the direction of the specified glyphs for bidirectional text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1532079-setbidilevels?language=objc
func (t_ Typesetter) SetBidiLevelsForGlyphRange(levels *uint8, glyphRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setBidiLevels:forGlyphRange:"), levels, glyphRange)
}

// Sets up layout parameters at the beginning of a paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526595-beginparagraph?language=objc
func (t_ Typesetter) BeginParagraph() {
	objc.Call[objc.Void](t_, objc.Sel("beginParagraph"))
}

// Returns whether bidirectional text processing is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533588-bidiprocessingenabled?language=objc
func (t_ Typesetter) BidiProcessingEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("bidiProcessingEnabled"))
	return rv
}

// Returns whether bidirectional text processing is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533588-bidiprocessingenabled?language=objc
func (t_ Typesetter) SetBidiProcessingEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setBidiProcessingEnabled:"), value)
}

// Returns the text backing store, usually an instance of NSTextStorage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1524704-attributedstring?language=objc
func (t_ Typesetter) AttributedString() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("attributedString"))
	return rv
}

// Returns the text backing store, usually an instance of NSTextStorage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1524704-attributedstring?language=objc
func (t_ Typesetter) SetAttributedString(value foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("setAttributedString:"), objc.Ptr(value))
}

// Returns the glyph range currently being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1528301-paragraphglyphrange?language=objc
func (t_ Typesetter) ParagraphGlyphRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("paragraphGlyphRange"))
	return rv
}

// Returns the current paragraph separator character range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1531746-paragraphseparatorcharacterrange?language=objc
func (t_ Typesetter) ParagraphSeparatorCharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("paragraphSeparatorCharacterRange"))
	return rv
}

// Returns the layout manager for the text being typeset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533958-layoutmanager?language=objc
func (t_ Typesetter) LayoutManager() LayoutManager {
	rv := objc.Call[LayoutManager](t_, objc.Sel("layoutManager"))
	return rv
}

// Returns the current line fragment padding, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1531129-linefragmentpadding?language=objc
func (t_ Typesetter) LineFragmentPadding() float64 {
	rv := objc.Call[float64](t_, objc.Sel("lineFragmentPadding"))
	return rv
}

// Returns the current line fragment padding, in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1531129-linefragmentpadding?language=objc
func (t_ Typesetter) SetLineFragmentPadding(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setLineFragmentPadding:"), value)
}

// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526716-usesfontleading?language=objc
func (t_ Typesetter) UsesFontLeading() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesFontLeading"))
	return rv
}

// Returns whether the typesetter uses the leading (or line gap) value specified in the font metric information of the current font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526716-usesfontleading?language=objc
func (t_ Typesetter) SetUsesFontLeading(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesFontLeading:"), value)
}

// Returns the current typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1528244-typesetterbehavior?language=objc
func (t_ Typesetter) TypesetterBehavior() TypesetterBehavior {
	rv := objc.Call[TypesetterBehavior](t_, objc.Sel("typesetterBehavior"))
	return rv
}

// Returns the current typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1528244-typesetterbehavior?language=objc
func (t_ Typesetter) SetTypesetterBehavior(value TypesetterBehavior) {
	objc.Call[objc.Void](t_, objc.Sel("setTypesetterBehavior:"), value)
}

// Returns an array containing the text containers belonging to the current layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1526310-textcontainers?language=objc
func (t_ Typesetter) TextContainers() []TextContainer {
	rv := objc.Call[[]TextContainer](t_, objc.Sel("textContainers"))
	return rv
}

// Returns a shared instance of a reentrant typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534814-sharedsystemtypesetter?language=objc
func (tc _TypesetterClass) SharedSystemTypesetter() Typesetter {
	rv := objc.Call[Typesetter](tc, objc.Sel("sharedSystemTypesetter"))
	return rv
}

// Returns a shared instance of a reentrant typesetter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534814-sharedsystemtypesetter?language=objc
func Typesetter_SharedSystemTypesetter() Typesetter {
	return TypesetterClass.SharedSystemTypesetter()
}

// Returns the current hyphenation factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535877-hyphenationfactor?language=objc
func (t_ Typesetter) HyphenationFactor() float64 {
	rv := objc.Call[float64](t_, objc.Sel("hyphenationFactor"))
	return rv
}

// Returns the current hyphenation factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1535877-hyphenationfactor?language=objc
func (t_ Typesetter) SetHyphenationFactor(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setHyphenationFactor:"), value)
}

// Returns the character range currently being processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1533422-paragraphcharacterrange?language=objc
func (t_ Typesetter) ParagraphCharacterRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("paragraphCharacterRange"))
	return rv
}

// Returns the paragraph style object for the text being typeset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1528893-currentparagraphstyle?language=objc
func (t_ Typesetter) CurrentParagraphStyle() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](t_, objc.Sel("currentParagraphStyle"))
	return rv
}

// Returns the text container for the text being typeset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534527-currenttextcontainer?language=objc
func (t_ Typesetter) CurrentTextContainer() TextContainer {
	rv := objc.Call[TextContainer](t_, objc.Sel("currentTextContainer"))
	return rv
}

// Returns the default typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534626-defaulttypesetterbehavior?language=objc
func (tc _TypesetterClass) DefaultTypesetterBehavior() TypesetterBehavior {
	rv := objc.Call[TypesetterBehavior](tc, objc.Sel("defaultTypesetterBehavior"))
	return rv
}

// Returns the default typesetter behavior. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534626-defaulttypesetterbehavior?language=objc
func Typesetter_DefaultTypesetterBehavior() TypesetterBehavior {
	return TypesetterClass.DefaultTypesetterBehavior()
}

// Returns the current paragraph separator range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534165-paragraphseparatorglyphrange?language=objc
func (t_ Typesetter) ParagraphSeparatorGlyphRange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("paragraphSeparatorGlyphRange"))
	return rv
}

// Returns the attributes used to lay out the extra line fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstypesetter/1534922-attributesforextralinefragment?language=objc
func (t_ Typesetter) AttributesForExtraLineFragment() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("attributesForExtraLineFragment"))
	return rv
}
