// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ILayoutManagerDelegate interface {
	ImplementsLayoutManagerDidInvalidateLayout() bool
	// optional
	LayoutManagerDidInvalidateLayout(sender LayoutManager)
	ImplementsLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool
	// optional
	LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint
	ImplementsLayoutManagerShouldUseActionForControlCharacterAtIndex() bool
	// optional
	LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	ImplementsLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool
	// optional
	LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)
	ImplementsLayoutManagerTextContainerDidChangeGeometryFromSize() bool
	// optional
	LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)
	ImplementsLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool
	// optional
	LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool
	ImplementsLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool
	// optional
	LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool
	ImplementsLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool
	// optional
	LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool
	// optional
	LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool
	// optional
	LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	ImplementsLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool
	// optional
	LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	ImplementsLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool
	// optional
	LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject
}

type LayoutManagerDelegate struct {
	_LayoutManagerDidInvalidateLayout                                                                              func(sender LayoutManager)
	_LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange                                  func(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint
	_LayoutManagerShouldUseActionForControlCharacterAtIndex                                                        func(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction
	_LayoutManagerDidCompleteLayoutForTextContainerAtEnd                                                           func(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)
	_LayoutManagerTextContainerDidChangeGeometryFromSize                                                           func(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)
	_LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex                                               func(layoutManager LayoutManager, charIndex uint) bool
	_LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex                                                      func(layoutManager LayoutManager, charIndex uint) bool
	_LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect                                         func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect                                    func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect                                   func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64
	_LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex func(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect
	_LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange                     func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject
}

func (di *LayoutManagerDelegate) ImplementsLayoutManagerDidInvalidateLayout() bool {
	return di._LayoutManagerDidInvalidateLayout != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerDidInvalidateLayout(f func(sender LayoutManager)) {
	di._LayoutManagerDidInvalidateLayout = f
}

func (di *LayoutManagerDelegate) LayoutManagerDidInvalidateLayout(sender LayoutManager) {
	di._LayoutManagerDidInvalidateLayout(sender)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool {
	return di._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(f func(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint) {
	di._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange = f
}

func (di *LayoutManagerDelegate) LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager LayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont Font, glyphRange foundation.Range) uint {
	return di._LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager, glyphs, props, charIndexes, aFont, glyphRange)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerShouldUseActionForControlCharacterAtIndex() bool {
	return di._LayoutManagerShouldUseActionForControlCharacterAtIndex != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerShouldUseActionForControlCharacterAtIndex(f func(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction) {
	di._LayoutManagerShouldUseActionForControlCharacterAtIndex = f
}

func (di *LayoutManagerDelegate) LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager LayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	return di._LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager, action, charIndex)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool {
	return di._LayoutManagerDidCompleteLayoutForTextContainerAtEnd != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerDidCompleteLayoutForTextContainerAtEnd(f func(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool)) {
	di._LayoutManagerDidCompleteLayoutForTextContainerAtEnd = f
}

func (di *LayoutManagerDelegate) LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager LayoutManager, textContainer TextContainer, layoutFinishedFlag bool) {
	di._LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager, textContainer, layoutFinishedFlag)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerTextContainerDidChangeGeometryFromSize() bool {
	return di._LayoutManagerTextContainerDidChangeGeometryFromSize != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerTextContainerDidChangeGeometryFromSize(f func(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size)) {
	di._LayoutManagerTextContainerDidChangeGeometryFromSize = f
}

func (di *LayoutManagerDelegate) LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager LayoutManager, textContainer TextContainer, oldSize foundation.Size) {
	di._LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager, textContainer, oldSize)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return di._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(f func(layoutManager LayoutManager, charIndex uint) bool) {
	di._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex = f
}

func (di *LayoutManagerDelegate) LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	return di._LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager, charIndex)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return di._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(f func(layoutManager LayoutManager, charIndex uint) bool) {
	di._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex = f
}

func (di *LayoutManagerDelegate) LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager LayoutManager, charIndex uint) bool {
	return di._LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager, charIndex)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return di._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegate) LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return di._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegate) LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool {
	return di._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(f func(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64) {
	di._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect = f
}

func (di *LayoutManagerDelegate) LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager LayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	return di._LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager, glyphIndex, rect)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return di._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(f func(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect) {
	di._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex = f
}

func (di *LayoutManagerDelegate) LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager LayoutManager, glyphIndex uint, textContainer TextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	return di._LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager, glyphIndex, textContainer, proposedRect, glyphPosition, charIndex)
}
func (di *LayoutManagerDelegate) ImplementsLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange != nil
}

func (di *LayoutManagerDelegate) SetLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(f func(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject) {
	di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange = f
}

func (di *LayoutManagerDelegate) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager LayoutManager, attrs map[foundation.AttributedStringKey]objc.Object, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.IObject {
	return di._LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager, attrs, toScreen, charIndex, effectiveCharRange)
}

type LayoutManagerDelegateWrapper struct {
	objc.Object
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerDidInvalidateLayout() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManagerDidInvalidateLayout:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerDidInvalidateLayout(sender ILayoutManager) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManagerDidInvalidateLayout:"), objc.ExtractPtr(sender))
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldGenerateGlyphsPropertiesCharacterIndexesFontForGlyphRange(layoutManager ILayoutManager, glyphs *coregraphics.Glyph, props *GlyphProperty, charIndexes *uint, aFont IFont, glyphRange foundation.Range) uint {
	rv := objc.CallMethod[uint](l_, objc.GetSelector("layoutManager:shouldGenerateGlyphs:properties:characterIndexes:font:forGlyphRange:"), objc.ExtractPtr(layoutManager), glyphs, props, charIndexes, objc.ExtractPtr(aFont), glyphRange)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerShouldUseActionForControlCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldUseAction:forControlCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldUseActionForControlCharacterAtIndex(layoutManager ILayoutManager, action ControlCharacterAction, charIndex uint) ControlCharacterAction {
	rv := objc.CallMethod[ControlCharacterAction](l_, objc.GetSelector("layoutManager:shouldUseAction:forControlCharacterAtIndex:"), objc.ExtractPtr(layoutManager), action, charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerDidCompleteLayoutForTextContainerAtEnd() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:didCompleteLayoutForTextContainer:atEnd:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerDidCompleteLayoutForTextContainerAtEnd(layoutManager ILayoutManager, textContainer ITextContainer, layoutFinishedFlag bool) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManager:didCompleteLayoutForTextContainer:atEnd:"), objc.ExtractPtr(layoutManager), objc.ExtractPtr(textContainer), layoutFinishedFlag)
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerTextContainerDidChangeGeometryFromSize() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:textContainer:didChangeGeometryFromSize:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerTextContainerDidChangeGeometryFromSize(layoutManager ILayoutManager, textContainer ITextContainer, oldSize foundation.Size) {
	objc.CallMethod[objc.Void](l_, objc.GetSelector("layoutManager:textContainer:didChangeGeometryFromSize:"), objc.ExtractPtr(layoutManager), objc.ExtractPtr(textContainer), oldSize)
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldBreakLineByHyphenatingBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManager:shouldBreakLineByHyphenatingBeforeCharacterAtIndex:"), objc.ExtractPtr(layoutManager), charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldBreakLineByWordBeforeCharacterAtIndex(layoutManager ILayoutManager, charIndex uint) bool {
	rv := objc.CallMethod[bool](l_, objc.GetSelector("layoutManager:shouldBreakLineByWordBeforeCharacterAtIndex:"), objc.ExtractPtr(layoutManager), charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerLineSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:lineSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerParagraphSpacingAfterGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:paragraphSpacingAfterGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerParagraphSpacingBeforeGlyphAtIndexWithProposedLineFragmentRect(layoutManager ILayoutManager, glyphIndex uint, rect foundation.Rect) float64 {
	rv := objc.CallMethod[float64](l_, objc.GetSelector("layoutManager:paragraphSpacingBeforeGlyphAtIndex:withProposedLineFragmentRect:"), objc.ExtractPtr(layoutManager), glyphIndex, rect)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerBoundingBoxForControlGlyphAtIndexForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(layoutManager ILayoutManager, glyphIndex uint, textContainer ITextContainer, proposedRect foundation.Rect, glyphPosition foundation.Point, charIndex uint) foundation.Rect {
	rv := objc.CallMethod[foundation.Rect](l_, objc.GetSelector("layoutManager:boundingBoxForControlGlyphAtIndex:forTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), objc.ExtractPtr(layoutManager), glyphIndex, objc.ExtractPtr(textContainer), proposedRect, glyphPosition, charIndex)
	return rv
}

func (l_ LayoutManagerDelegateWrapper) ImplementsLayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange() bool {
	return l_.RespondsToSelector(objc.GetSelector("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"))
}

func (l_ LayoutManagerDelegateWrapper) LayoutManagerShouldUseTemporaryAttributesForDrawingToScreenAtCharacterIndexEffectiveRange(layoutManager ILayoutManager, attrs map[foundation.AttributedStringKey]objc.IObject, toScreen bool, charIndex uint, effectiveCharRange *foundation.Range) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](l_, objc.GetSelector("layoutManager:shouldUseTemporaryAttributes:forDrawingToScreen:atCharacterIndex:effectiveRange:"), objc.ExtractPtr(layoutManager), attrs, toScreen, charIndex, effectiveCharRange)
	return rv
}
