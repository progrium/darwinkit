// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextView] class.
var TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}

type _TextViewClass struct {
	objc.Class
}

// An interface definition for the [TextView] class.
type ITextView interface {
	IText
	DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool
	OrderFrontLinkPanel(sender objc.IObject)
	PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType
	ToggleAutomaticLinkDetection(sender objc.IObject)
	UseStandardLigatures(sender objc.IObject)
	ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.Range, replacementString string) bool
	DidChangeText()
	SetSelectedRangeAffinityStillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool)
	SetAlignmentRange(alignment TextAlignment, range_ foundation.Range)
	QuickLookPreviewableItemsInRanges(ranges []foundation.IValue) []objc.Object
	ToggleAutomaticTextCompletion(sender objc.IObject) objc.Object
	ToggleSmartInsertDelete(sender objc.IObject)
	ShowFindIndicatorForRange(charRange foundation.Range)
	SetSelectedRangesAffinityStillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool)
	CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.Range, index *int) []string
	BreakUndoCoalescing()
	ChangeColor(sender objc.IObject)
	OrderFrontSubstitutionsPanel(sender objc.IObject)
	DragImageForSelectionWithEventOrigin(event IEvent, origin foundation.PointPointer) Image
	StartSpeaking(sender objc.IObject)
	SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	UpdateQuickLookPreviewPanel()
	UpdateDragTypeRegistration()
	SetSpellingStateRange(value int, charRange foundation.Range)
	HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int)
	DrawViewBackgroundInRect(rect foundation.Rect)
	UseAllLigatures(sender objc.IObject)
	ToggleAutomaticSpellingCorrection(sender objc.IObject)
	TightenKerning(sender objc.IObject)
	ToggleContinuousSpellChecking(sender objc.IObject)
	ToggleAutomaticDataDetection(sender objc.IObject)
	SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range
	CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject)
	OrderFrontSpacingPanel(sender objc.IObject)
	UseStandardKerning(sender objc.IObject)
	ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool
	SetConstrainedFrameSize(desiredSize foundation.Size)
	OrderFrontTablePanel(sender objc.IObject)
	InvalidateTextContainerOrigin()
	LoosenKerning(sender objc.IObject)
	CheckTextInDocument(sender objc.IObject)
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	ToggleAutomaticTextReplacement(sender objc.IObject)
	OrderFrontSharingServicePicker(sender objc.IObject) objc.Object
	ToggleAutomaticDashSubstitution(sender objc.IObject)
	Outline(sender objc.IObject)
	PasteAsPlainText(sender objc.IObject)
	PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool
	SetLayoutOrientation(orientation TextLayoutOrientation)
	DragOperationForDraggingInfoType(dragInfo PDraggingInfo, type_ PasteboardType) DragOperation
	DragOperationForDraggingInfoObjectType(dragInfoObject objc.IObject, type_ PasteboardType) DragOperation
	CheckTextInSelection(sender objc.IObject)
	UpdateTouchBarItemIdentifiers()
	OrderFrontListPanel(sender objc.IObject)
	RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker)
	InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.Range, movement int, flag bool)
	Complete(sender objc.IObject)
	ClickedOnLinkAtIndex(link objc.IObject, charIndex uint)
	AlignJustified(sender objc.IObject)
	PasteAsRichText(sender objc.IObject)
	UpdateCandidates()
	ChangeLayoutOrientation(sender objc.IObject)
	ToggleQuickLookPreviewPanel(sender objc.IObject) objc.Object
	SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.Range, beforeString string, afterString string)
	WriteSelectionToPasteboardTypes(pboard IPasteboard, types []PasteboardType) bool
	TurnOffKerning(sender objc.IObject)
	ChangeDocumentBackgroundColor(sender objc.IObject)
	DrawInsertionPointInRectColorTurnedOn(rect foundation.Rect, color IColor, flag bool)
	StopSpeaking(sender objc.IObject)
	RaiseBaseline(sender objc.IObject)
	TurnOffLigatures(sender objc.IObject)
	UpdateTextTouchBarItems()
	ToggleAutomaticQuoteSubstitution(sender objc.IObject)
	LowerBaseline(sender objc.IObject)
	UpdateRuler()
	ReplaceTextContainer(newContainer ITextContainer)
	ChangeAttributes(sender objc.IObject)
	SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ foundation.Range)
	SetNeedsDisplayInRectAvoidAdditionalLayout(rect foundation.Rect, flag bool)
	PerformFindPanelAction(sender objc.IObject)
	UpdateFontPanel()
	SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range
	ToggleGrammarChecking(sender objc.IObject)
	CleanUpAfterDragOperation()
	CharacterIndexForInsertionAtPoint(point foundation.Point) uint
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	TextLayoutManager() TextLayoutManager
	DefaultParagraphStyle() ParagraphStyle
	SetDefaultParagraphStyle(value IParagraphStyle)
	IsAutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
	ShouldDrawInsertionPoint() bool
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	SpellCheckerDocumentTag() int
	RangeForUserParagraphAttributeChange() foundation.Range
	InsertionPointColor() Color
	SetInsertionPointColor(value IColor)
	EnabledTextCheckingTypes() foundation.TextCheckingTypes
	SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes)
	RangeForUserTextChange() foundation.Range
	SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	WritablePasteboardTypes() []PasteboardType
	TextContainerInset() foundation.Size
	SetTextContainerInset(value foundation.Size)
	IsCoalescingUndo() bool
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	RangesForUserParagraphAttributeChange() []foundation.Value
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	IsAutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	RangeForUserCompletion() foundation.Range
	LayoutManager() LayoutManager
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	RangesForUserCharacterAttributeChange() []foundation.Value
	RangeForUserCharacterAttributeChange() foundation.Range
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	SetRulerVisible(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	RangesForUserTextChange() []foundation.Value
	ReadablePasteboardTypes() []PasteboardType
	SelectedRanges() []foundation.Value
	SetSelectedRanges(value []foundation.IValue)
	TextStorage() TextStorage
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	SelectionAffinity() SelectionAffinity
	LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	IsGrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	AcceptableDragTypes() []PasteboardType
	TextContainerOrigin() foundation.Point
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	TextContentStorage() TextContentStorage
	SelectionGranularity() SelectionGranularity
	SetSelectionGranularity(value SelectionGranularity)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
}

// A view that draws text and handles user interactions with that text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview?language=objc
type TextView struct {
	Text
}

func TextViewFrom(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: TextFrom(ptr),
	}
}

func (t_ TextView) InitWithFrameTextContainer(frameRect foundation.Rect, container ITextContainer) TextView {
	rv := objc.Call[TextView](t_, objc.Sel("initWithFrame:textContainer:"), frameRect, objc.Ptr(container))
	return rv
}

// Initializes a text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449347-initwithframe?language=objc
func NewTextViewWithFrameTextContainer(frameRect foundation.Rect, container ITextContainer) TextView {
	instance := TextViewClass.Alloc().InitWithFrameTextContainer(frameRect, container)
	instance.Autorelease()
	return instance
}

func (tc _TextViewClass) Alloc() TextView {
	rv := objc.Call[TextView](tc, objc.Sel("alloc"))
	return rv
}

func TextView_Alloc() TextView {
	return TextViewClass.Alloc()
}

func (tc _TextViewClass) New() TextView {
	rv := objc.Call[TextView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextView() TextView {
	return TextViewClass.New()
}

func (t_ TextView) Init() TextView {
	rv := objc.Call[TextView](t_, objc.Sel("init"))
	return rv
}

func (t_ TextView) InitWithFrame(frameRect foundation.Rect) TextView {
	rv := objc.Call[TextView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstext/1525191-initwithframe?language=objc
func NewTextViewWithFrame(frameRect foundation.Rect) TextView {
	instance := TextViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Begins dragging the current selected text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449413-dragselectionwithevent?language=objc
func (t_ TextView) DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool {
	rv := objc.Call[bool](t_, objc.Sel("dragSelectionWithEvent:offset:slideBack:"), objc.Ptr(event), mouseOffset, slideBack)
	return rv
}

// Brings forward a panel allowing the user to manipulate links in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449238-orderfrontlinkpanel?language=objc
func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontLinkPanel:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990527-scrollabledocumentcontenttextvie?language=objc
func (tc _TextViewClass) ScrollableDocumentContentTextView() ScrollView {
	rv := objc.Call[ScrollView](tc, objc.Sel("scrollableDocumentContentTextView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990527-scrollabledocumentcontenttextvie?language=objc
func TextView_ScrollableDocumentContentTextView() ScrollView {
	return TextViewClass.ScrollableDocumentContentTextView()
}

// Returns whatever type on the pasteboard would be most preferred for copying data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449131-preferredpasteboardtypefromarray?language=objc
func (t_ TextView) PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType {
	rv := objc.Call[PasteboardType](t_, objc.Sel("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), availableTypes, allowedTypes)
	return rv
}

// Changes the state of automatic link detection from enabled to disabled and vice versa. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449353-toggleautomaticlinkdetection?language=objc
func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticLinkDetection:"), sender)
}

// Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449144-usestandardligatures?language=objc
func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("useStandardLigatures:"), sender)
}

// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449532-shouldchangetextinrange?language=objc
func (t_ TextView) ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldChangeTextInRange:replacementString:"), affectedCharRange, replacementString)
	return rv
}

// Sends out necessary notifications when a text change completes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449296-didchangetext?language=objc
func (t_ TextView) DidChangeText() {
	objc.Call[objc.Void](t_, objc.Sel("didChangeText"))
}

// Sets the selection to a range of characters in response to user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449462-setselectedrange?language=objc
func (t_ TextView) SetSelectedRangeAffinityStillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}

// Sets the alignment of the paragraphs containing characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449536-setalignment?language=objc
func (t_ TextView) SetAlignmentRange(alignment TextAlignment, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setAlignment:range:"), alignment, range_)
}

// Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449426-quicklookpreviewableitemsinrange?language=objc
func (t_ TextView) QuickLookPreviewableItemsInRanges(ranges []foundation.IValue) []objc.Object {
	rv := objc.Call[[]objc.Object](t_, objc.Sel("quickLookPreviewableItemsInRanges:"), ranges)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544841-toggleautomatictextcompletion?language=objc
func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("toggleAutomaticTextCompletion:"), sender)
	return rv
}

// Changes the state of smart insert and delete from enabled to disabled and vice versa. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449273-togglesmartinsertdelete?language=objc
func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleSmartInsertDelete:"), sender)
}

// Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449540-showfindindicatorforrange?language=objc
func (t_ TextView) ShowFindIndicatorForRange(charRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("showFindIndicatorForRange:"), charRange)
}

// Sets the selection to the characters in an array of ranges in response to user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449208-setselectedranges?language=objc
func (t_ TextView) SetSelectedRangesAffinityStillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedRanges:affinity:stillSelecting:"), ranges, affinity, stillSelectingFlag)
}

// Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449405-completionsforpartialwordrange?language=objc
func (t_ TextView) CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.Range, index *int) []string {
	rv := objc.Call[[]string](t_, objc.Sel("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return rv
}

// Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449384-breakundocoalescing?language=objc
func (t_ TextView) BreakUndoCoalescing() {
	objc.Call[objc.Void](t_, objc.Sel("breakUndoCoalescing"))
}

// Sets the color of the selected text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449282-changecolor?language=objc
func (t_ TextView) ChangeColor(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeColor:"), sender)
}

// Brings forward a panel allowing the user to specify string substitutions in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449327-orderfrontsubstitutionspanel?language=objc
func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}

// Returns an appropriate drag image for the drag initiated by the specified event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449423-dragimageforselectionwithevent?language=objc
func (t_ TextView) DragImageForSelectionWithEventOrigin(event IEvent, origin foundation.PointPointer) Image {
	rv := objc.Call[Image](t_, objc.Sel("dragImageForSelectionWithEvent:origin:"), objc.Ptr(event), origin)
	return rv
}

// Speaks the selected text, or all text if no selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449519-startspeaking?language=objc
func (t_ TextView) StartSpeaking(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("startSpeaking:"), sender)
}

// Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449467-smartinsertafterstringforstring?language=objc
func (t_ TextView) SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.Call[string](t_, objc.Sel("smartInsertAfterStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

// Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449242-smartinsertbeforestringforstring?language=objc
func (t_ TextView) SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.Call[string](t_, objc.Sel("smartInsertBeforeStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

// Notifies the QuickLook panel that an update may be required. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449409-updatequicklookpreviewpanel?language=objc
func (t_ TextView) UpdateQuickLookPreviewPanel() {
	objc.Call[objc.Void](t_, objc.Sel("updateQuickLookPreviewPanel"))
}

// Updates the acceptable drag types of all text views associated with the receiver's layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449181-updatedragtyperegistration?language=objc
func (t_ TextView) UpdateDragTypeRegistration() {
	objc.Call[objc.Void](t_, objc.Sel("updateDragTypeRegistration"))
}

// Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449493-setspellingstate?language=objc
func (t_ TextView) SetSpellingStateRange(value int, charRange foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setSpellingState:range:"), value, charRange)
}

// Handles the text checking results returned by the text view [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449473-handletextcheckingresults?language=objc
func (t_ TextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int) {
	objc.Call[objc.Void](t_, objc.Sel("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, objc.Ptr(orthography), wordCount)
}

// Draws the background of the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449135-drawviewbackgroundinrect?language=objc
func (t_ TextView) DrawViewBackgroundInRect(rect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawViewBackgroundInRect:"), rect)
}

// Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449213-useallligatures?language=objc
func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("useAllLigatures:"), sender)
}

// Toggles the state of the automatic spelling correction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449178-toggleautomaticspellingcorrectio?language=objc
func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticSpellingCorrection:"), sender)
}

// Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449137-tightenkerning?language=objc
func (t_ TextView) TightenKerning(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("tightenKerning:"), sender)
}

// Toggles whether continuous spell checking is enabled for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449471-togglecontinuousspellchecking?language=objc
func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleContinuousSpellChecking:"), sender)
}

// Toggles the state of the automatic data detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449499-toggleautomaticdatadetection?language=objc
func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticDataDetection:"), sender)
}

// Returns an adjusted selected range based on the selection granularity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449188-selectionrangeforproposedrange?language=objc
func (t_ TextView) SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return rv
}

// Check and replace the text in the range using the specified checking types and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449240-checktextinrange?language=objc
func (t_ TextView) CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}

// Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449438-orderfrontspacingpanel?language=objc
func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontSpacingPanel:"), sender)
}

// Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449491-usestandardkerning?language=objc
func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("useStandardKerning:"), sender)
}

// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449311-shouldchangetextinranges?language=objc
func (t_ TextView) ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldChangeTextInRanges:replacementStrings:"), affectedRanges, replacementStrings)
	return rv
}

// Attempts to set the frame size as if by user action. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449230-setconstrainedframesize?language=objc
func (t_ TextView) SetConstrainedFrameSize(desiredSize foundation.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setConstrainedFrameSize:"), desiredSize)
}

// Brings forward a panel allowing the user to manipulate text tables in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449442-orderfronttablepanel?language=objc
func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontTablePanel:"), sender)
}

// Invalidates the calculated origin of the text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449546-invalidatetextcontainerorigin?language=objc
func (t_ TextView) InvalidateTextContainerOrigin() {
	objc.Call[objc.Void](t_, objc.Sel("invalidateTextContainerOrigin"))
}

// Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449183-loosenkerning?language=objc
func (t_ TextView) LoosenKerning(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("loosenKerning:"), sender)
}

// Performs the default text checking on the entire document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449440-checktextindocument?language=objc
func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInDocument:"), sender)
}

// Reads the text view’s preferred type of data from the specified pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449469-readselectionfrompasteboard?language=objc
func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.Call[bool](t_, objc.Sel("readSelectionFromPasteboard:"), objc.Ptr(pboard))
	return rv
}

// Toggles the state of the automatic text replacement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449200-toggleautomatictextreplacement?language=objc
func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticTextReplacement:"), sender)
}

// Creates and displays a new instance of the sharing service picker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449150-orderfrontsharingservicepicker?language=objc
func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("orderFrontSharingServicePicker:"), sender)
	return rv
}

// Toggles the state of the automatic dash substitution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449305-toggleautomaticdashsubstitution?language=objc
func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticDashSubstitution:"), sender)
}

// Adds the outline attribute to the selected text attributes if absent; removes the attribute if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449386-outline?language=objc
func (t_ TextView) Outline(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("outline:"), sender)
}

// Inserts the contents of the pasteboard into the receiver’s text as plain text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449250-pasteasplaintext?language=objc
func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("pasteAsPlainText:"), sender)
}

// Replaces text in the range you specify with the attributed string you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990526-performvalidatedreplacementinran?language=objc
func (t_ TextView) PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool {
	rv := objc.Call[bool](t_, objc.Sel("performValidatedReplacementInRange:withAttributedString:"), range_, objc.Ptr(attributedString))
	return rv
}

// Changes the receiver's layout orientation and invalidates the contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449483-setlayoutorientation?language=objc
func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutOrientation:"), orientation)
}

// Returns the type of drag operation that should be performed if the image were released now. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449176-dragoperationfordragginginfo?language=objc
func (t_ TextView) DragOperationForDraggingInfoType(dragInfo PDraggingInfo, type_ PasteboardType) DragOperation {
	po0 := objc.WrapAsProtocol("NSDraggingInfo", dragInfo)
	rv := objc.Call[DragOperation](t_, objc.Sel("dragOperationForDraggingInfo:type:"), po0, type_)
	return rv
}

// Returns the type of drag operation that should be performed if the image were released now. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449176-dragoperationfordragginginfo?language=objc
func (t_ TextView) DragOperationForDraggingInfoObjectType(dragInfoObject objc.IObject, type_ PasteboardType) DragOperation {
	rv := objc.Call[DragOperation](t_, objc.Sel("dragOperationForDraggingInfo:type:"), objc.Ptr(dragInfoObject), type_)
	return rv
}

// Performs the default text checking on the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449382-checktextinselection?language=objc
func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("checkTextInSelection:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544834-updatetouchbaritemidentifiers?language=objc
func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	objc.Call[objc.Void](t_, objc.Sel("updateTouchBarItemIdentifiers"))
}

// Brings forward a panel allowing the user to manipulate text lists in the text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449349-orderfrontlistpanel?language=objc
func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("orderFrontListPanel:"), sender)
}

// Modifies the paragraph style of the paragraphs containing the selection to accommodate a new marker. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449523-rulerview?language=objc
func (t_ TextView) RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Call[objc.Void](t_, objc.Sel("rulerView:didAddMarker:"), objc.Ptr(ruler), objc.Ptr(marker))
}

// Inserts the selected completion into the text at the appropriate location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449448-insertcompletion?language=objc
func (t_ TextView) InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.Range, movement int, flag bool) {
	objc.Call[objc.Void](t_, objc.Sel("insertCompletion:forPartialWordRange:movement:isFinal:"), word, charRange, movement, flag)
}

// Invokes completion in a text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449359-complete?language=objc
func (t_ TextView) Complete(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("complete:"), sender)
}

// Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449497-clickedonlink?language=objc
func (t_ TextView) ClickedOnLinkAtIndex(link objc.IObject, charIndex uint) {
	objc.Call[objc.Void](t_, objc.Sel("clickedOnLink:atIndex:"), link, charIndex)
}

// Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449515-alignjustified?language=objc
func (t_ TextView) AlignJustified(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("alignJustified:"), sender)
}

// This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449395-pasteasrichtext?language=objc
func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("pasteAsRichText:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544833-updatecandidates?language=objc
func (t_ TextView) UpdateCandidates() {
	objc.Call[objc.Void](t_, objc.Sel("updateCandidates"))
}

// An action method that sets the layout orientation of the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449286-changelayoutorientation?language=objc
func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeLayoutOrientation:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990528-scrollableplaindocumentcontentte?language=objc
func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() ScrollView {
	rv := objc.Call[ScrollView](tc, objc.Sel("scrollablePlainDocumentContentTextView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990528-scrollableplaindocumentcontentte?language=objc
func TextView_ScrollablePlainDocumentContentTextView() ScrollView {
	return TextViewClass.ScrollablePlainDocumentContentTextView()
}

// An action message that toggles the visibility state of the Quick Look preview panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449415-togglequicklookpreviewpanel?language=objc
func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("toggleQuickLookPreviewPanel:"), sender)
	return rv
}

// Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449544-smartinsertforstring?language=objc
func (t_ TextView) SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.Range, beforeString string, afterString string) {
	objc.Call[objc.Void](t_, objc.Sel("smartInsertForString:replacingRange:beforeString:afterString:"), pasteString, charRangeToReplace, beforeString, afterString)
}

// Writes the current selection to the specified pasteboard under each given type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449277-writeselectiontopasteboard?language=objc
func (t_ TextView) WriteSelectionToPasteboardTypes(pboard IPasteboard, types []PasteboardType) bool {
	rv := objc.Call[bool](t_, objc.Sel("writeSelectionToPasteboard:types:"), objc.Ptr(pboard), types)
	return rv
}

// Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449464-turnoffkerning?language=objc
func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("turnOffKerning:"), sender)
}

// An action method used to set the background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449475-changedocumentbackgroundcolor?language=objc
func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeDocumentBackgroundColor:"), sender)
}

// Draws or erases the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449232-drawinsertionpointinrect?language=objc
func (t_ TextView) DrawInsertionPointInRectColorTurnedOn(rect foundation.Rect, color IColor, flag bool) {
	objc.Call[objc.Void](t_, objc.Sel("drawInsertionPointInRect:color:turnedOn:"), rect, objc.Ptr(color), flag)
}

// Stops the speaking of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449172-stopspeaking?language=objc
func (t_ TextView) StopSpeaking(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("stopSpeaking:"), sender)
}

// Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449198-raisebaseline?language=objc
func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("raiseBaseline:"), sender)
}

// Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449436-turnoffligatures?language=objc
func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("turnOffLigatures:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544676-updatetexttouchbaritems?language=objc
func (t_ TextView) UpdateTextTouchBarItems() {
	objc.Call[objc.Void](t_, objc.Sel("updateTextTouchBarItems"))
}

// Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449444-toggleautomaticquotesubstitution?language=objc
func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleAutomaticQuoteSubstitution:"), sender)
}

// Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449289-lowerbaseline?language=objc
func (t_ TextView) LowerBaseline(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("lowerBaseline:"), sender)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990529-scrollabletextview?language=objc
func (tc _TextViewClass) ScrollableTextView() ScrollView {
	rv := objc.Call[ScrollView](tc, objc.Sel("scrollableTextView"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2990529-scrollabletextview?language=objc
func TextView_ScrollableTextView() ScrollView {
	return TextViewClass.ScrollableTextView()
}

// Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449323-updateruler?language=objc
func (t_ TextView) UpdateRuler() {
	objc.Call[objc.Void](t_, objc.Sel("updateRuler"))
}

// Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449479-replacetextcontainer?language=objc
func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	objc.Call[objc.Void](t_, objc.Sel("replaceTextContainer:"), objc.Ptr(newContainer))
}

// Changes the attributes of the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449216-changeattributes?language=objc
func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("changeAttributes:"), sender)
}

// Sets the base writing direction of a range of text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449146-setbasewritingdirection?language=objc
func (t_ TextView) SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}

// Marks the receiver as requiring display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449279-setneedsdisplayinrect?language=objc
func (t_ TextView) SetNeedsDisplayInRectAvoidAdditionalLayout(rect foundation.Rect, flag bool) {
	objc.Call[objc.Void](t_, objc.Sel("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}

// Performs a find panel action specified by the sender's tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449525-performfindpanelaction?language=objc
func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("performFindPanelAction:"), sender)
}

// Updates the Font panel to contain the font attributes of the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449401-updatefontpanel?language=objc
func (t_ TextView) UpdateFontPanel() {
	objc.Call[objc.Void](t_, objc.Sel("updateFontPanel"))
}

// Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449428-smartdeleterangeforproposedrange?language=objc
func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return rv
}

// Registers send and return types for the Services facility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449507-registerforservices?language=objc
func (tc _TextViewClass) RegisterForServices() {
	objc.Call[objc.Void](tc, objc.Sel("registerForServices"))
}

// Registers send and return types for the Services facility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449507-registerforservices?language=objc
func TextView_RegisterForServices() {
	TextViewClass.RegisterForServices()
}

// Changes the state of grammar checking from enabled to disabled and vice versa. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449393-togglegrammarchecking?language=objc
func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("toggleGrammarChecking:"), sender)
}

// Releases the drag information still existing after the dragging session has completed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449202-cleanupafterdragoperation?language=objc
func (t_ TextView) CleanUpAfterDragOperation() {
	objc.Call[objc.Void](t_, objc.Sel("cleanUpAfterDragOperation"))
}

// Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449505-characterindexforinsertionatpoin?language=objc
func (t_ TextView) CharacterIndexForInsertionAtPoint(point foundation.Point) uint {
	rv := objc.Call[uint](t_, objc.Sel("characterIndexForInsertionAtPoint:"), point)
	return rv
}

// Updates the insertion point’s location and optionally restarts the blinking cursor timer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449268-updateinsertionpointstateandrest?language=objc
func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.Call[objc.Void](t_, objc.Sel("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}

// The manager that lays out text for the receiver’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/3824763-textlayoutmanager?language=objc
func (t_ TextView) TextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("textLayoutManager"))
	return rv
}

// The receiver’s default paragraph style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449271-defaultparagraphstyle?language=objc
func (t_ TextView) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.Call[ParagraphStyle](t_, objc.Sel("defaultParagraphStyle"))
	return rv
}

// The receiver’s default paragraph style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449271-defaultparagraphstyle?language=objc
func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setDefaultParagraphStyle:"), objc.Ptr(value))
}

// A Boolean value that enables or disables automatic link detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc
func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}

// A Boolean value that enables or disables automatic link detection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449170-automaticlinkdetectionenabled?language=objc
func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticLinkDetectionEnabled:"), value)
}

// Indicates whether image attachments should permit editing of their images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc
func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsImageEditing"))
	return rv
}

// Indicates whether image attachments should permit editing of their images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449425-allowsimageediting?language=objc
func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsImageEditing:"), value)
}

// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc
func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}

// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/3237223-usesadaptivecolormappingfordarka?language=objc
func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}

// A Boolean value that indicates whether automatic text replacement is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc
func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}

// A Boolean value that indicates whether automatic text replacement is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449210-automatictextreplacementenabled?language=objc
func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticTextReplacementEnabled:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc
func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449357-usesrolloverbuttonforselection?language=objc
func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}

// A Boolean value that determines whether the receiver should draw its insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449152-shoulddrawinsertionpoint?language=objc
func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Call[bool](t_, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2269433-stronglyreferencestextstorage?language=objc
func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.Call[bool](tc, objc.Sel("stronglyReferencesTextStorage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2269433-stronglyreferencestextstorage?language=objc
func TextView_StronglyReferencesTextStorage() bool {
	return TextViewClass.StronglyReferencesTextStorage()
}

// A Boolean value that enables and disables automatic quotation mark substitution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc
func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}

// A Boolean value that enables and disables automatic quotation mark substitution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449258-automaticquotesubstitutionenable?language=objc
func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticQuoteSubstitutionEnabled:"), value)
}

// A tag identifying the text view's text as a document for the spell checker server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449513-spellcheckerdocumenttag?language=objc
func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.Call[int](t_, objc.Sel("spellCheckerDocumentTag"))
	return rv
}

// The range of characters affected by an action method that changes paragraph (not character) attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449252-rangeforuserparagraphattributech?language=objc
func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("rangeForUserParagraphAttributeChange"))
	return rv
}

// The color of the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc
func (t_ TextView) InsertionPointColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("insertionPointColor"))
	return rv
}

// The color of the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449309-insertionpointcolor?language=objc
func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setInsertionPointColor:"), objc.Ptr(value))
}

// The default text checking types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449529-enabledtextcheckingtypes?language=objc
func (t_ TextView) EnabledTextCheckingTypes() foundation.TextCheckingTypes {
	rv := objc.Call[foundation.TextCheckingTypes](t_, objc.Sel("enabledTextCheckingTypes"))
	return rv
}

// The default text checking types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449529-enabledtextcheckingtypes?language=objc
func (t_ TextView) SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes) {
	objc.Call[objc.Void](t_, objc.Sel("setEnabledTextCheckingTypes:"), value)
}

// The range of characters affected by a method that changes characters (as opposed to attributes). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449315-rangeforusertextchange?language=objc
func (t_ TextView) RangeForUserTextChange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("rangeForUserTextChange"))
	return rv
}

// The attributes used to indicate the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc
func (t_ TextView) SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("selectedTextAttributes"))
	return rv
}

// The attributes used to indicate the selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449270-selectedtextattributes?language=objc
func (t_ TextView) SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedTextAttributes:"), value)
}

// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc
func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}

// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544655-automatictextcompletionenabled?language=objc
func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}

// The pasteboard types that can be provided from the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449222-writablepasteboardtypes?language=objc
func (t_ TextView) WritablePasteboardTypes() []PasteboardType {
	rv := objc.Call[[]PasteboardType](t_, objc.Sel("writablePasteboardTypes"))
	return rv
}

// The empty space the receiver leaves around its associated text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc
func (t_ TextView) TextContainerInset() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("textContainerInset"))
	return rv
}

// The empty space the receiver leaves around its associated text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449168-textcontainerinset?language=objc
func (t_ TextView) SetTextContainerInset(value foundation.Size) {
	objc.Call[objc.Void](t_, objc.Sel("setTextContainerInset:"), value)
}

// A Boolean value that indicates whether undo coalescing is in progress. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449368-coalescingundo?language=objc
func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.Call[bool](t_, objc.Sel("isCoalescingUndo"))
	return rv
}

// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc
func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := objc.Call[[]string](t_, objc.Sel("allowedInputSourceLocales"))
	return rv
}

// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449370-allowedinputsourcelocales?language=objc
func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowedInputSourceLocales:"), value)
}

// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449161-rangesforuserparagraphattributec?language=objc
func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("rangesForUserParagraphAttributeChange"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc
func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/2544680-allowscharacterpickertouchbarite?language=objc
func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}

// The receiver’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc
func (t_ TextView) TextContainer() TextContainer {
	rv := objc.Call[TextContainer](t_, objc.Sel("textContainer"))
	return rv
}

// The receiver’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449364-textcontainer?language=objc
func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.Call[objc.Void](t_, objc.Sel("setTextContainer:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the receiver accepts the glyph info attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc
func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.Call[bool](t_, objc.Sel("acceptsGlyphInfo"))
	return rv
}

// A Boolean value that indicates whether the receiver accepts the glyph info attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449163-acceptsglyphinfo?language=objc
func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAcceptsGlyphInfo:"), value)
}

// A Boolean value that indicates whether the receiver allows for a find panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc
func (t_ TextView) UsesFindPanel() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesFindPanel"))
	return rv
}

// A Boolean value that indicates whether the receiver allows for a find panel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449293-usesfindpanel?language=objc
func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesFindPanel:"), value)
}

// A Boolean value that indicates whether automatic dash substitution is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc
func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}

// A Boolean value that indicates whether automatic dash substitution is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449403-automaticdashsubstitutionenabled?language=objc
func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticDashSubstitutionEnabled:"), value)
}

// The partial range from the most recent beginning of a word up to the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449329-rangeforusercompletion?language=objc
func (t_ TextView) RangeForUserCompletion() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("rangeForUserCompletion"))
	return rv
}

// The layout manager that lays out text for the receiver’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449148-layoutmanager?language=objc
func (t_ TextView) LayoutManager() LayoutManager {
	rv := objc.Call[LayoutManager](t_, objc.Sel("layoutManager"))
	return rv
}

// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc
func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}

// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449236-smartinsertdeleteenabled?language=objc
func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}

// A Boolean value that indicates whether incremental searching is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc
func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}

// A Boolean value that indicates whether incremental searching is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449458-incrementalsearchingenabled?language=objc
func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setIncrementalSearchingEnabled:"), value)
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc
func (t_ TextView) UsesRuler() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesRuler"))
	return rv
}

// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449218-usesruler?language=objc
func (t_ TextView) SetUsesRuler(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesRuler:"), value)
}

// A Boolean value that indicates whether automatic data detection is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc
func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}

// A Boolean value that indicates whether automatic data detection is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449192-automaticdatadetectionenabled?language=objc
func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticDataDetectionEnabled:"), value)
}

// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449503-rangesforusercharacterattributec?language=objc
func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("rangesForUserCharacterAttributeChange"))
	return rv
}

// The range of characters affected by an action method that changes character (not paragraph) attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449392-rangeforusercharacterattributech?language=objc
func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.Range {
	rv := objc.Call[foundation.Range](t_, objc.Sel("rangeForUserCharacterAttributeChange"))
	return rv
}

// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc
func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.Call[bool](t_, objc.Sel("displaysLinkToolTips"))
	return rv
}

// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449204-displayslinktooltips?language=objc
func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDisplaysLinkToolTips:"), value)
}

// A Boolean value that controls whether the scroll view enclosing text views sharing the receiver’s layout manager displays the ruler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449406-rulervisible?language=objc
func (t_ TextView) SetRulerVisible(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setRulerVisible:"), value)
}

// A Boolean value that indicates whether the receiver has continuous spell checking enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc
func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}

// A Boolean value that indicates whether the receiver has continuous spell checking enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449430-continuousspellcheckingenabled?language=objc
func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setContinuousSpellCheckingEnabled:"), value)
}

// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449434-rangesforusertextchange?language=objc
func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("rangesForUserTextChange"))
	return rv
}

// The types this text view can read immediately from the pasteboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449361-readablepasteboardtypes?language=objc
func (t_ TextView) ReadablePasteboardTypes() []PasteboardType {
	rv := objc.Call[[]PasteboardType](t_, objc.Sel("readablePasteboardTypes"))
	return rv
}

// An array containing the ranges of characters selected in the receiver’s layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc
func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := objc.Call[[]foundation.Value](t_, objc.Sel("selectedRanges"))
	return rv
}

// An array containing the ranges of characters selected in the receiver’s layout manager. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449129-selectedranges?language=objc
func (t_ TextView) SetSelectedRanges(value []foundation.IValue) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectedRanges:"), value)
}

// The receiver’s text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449196-textstorage?language=objc
func (t_ TextView) TextStorage() TextStorage {
	rv := objc.Call[TextStorage](t_, objc.Sel("textStorage"))
	return rv
}

// A Boolean value that indicates whether the receiver allows its background color to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc
func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}

// A Boolean value that indicates whether the receiver allows its background color to change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449397-allowsdocumentbackgroundcolorcha?language=objc
func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}

// The receiver’s typing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc
func (t_ TextView) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("typingAttributes"))
	return rv
}

// The receiver’s typing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449487-typingattributes?language=objc
func (t_ TextView) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setTypingAttributes:"), value)
}

// The preferred direction of selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449291-selectionaffinity?language=objc
func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := objc.Call[SelectionAffinity](t_, objc.Sel("selectionAffinity"))
	return rv
}

// The attributes used to draw the onscreen presentation of link text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc
func (t_ TextView) LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("linkTextAttributes"))
	return rv
}

// The attributes used to draw the onscreen presentation of link text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449452-linktextattributes?language=objc
func (t_ TextView) SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setLinkTextAttributes:"), value)
}

// A Boolean value that indicates whether automatic spelling correction is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc
func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}

// A Boolean value that indicates whether automatic spelling correction is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449254-automaticspellingcorrectionenabl?language=objc
func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticSpellingCorrectionEnabled:"), value)
}

// Enables and disables grammar checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc
func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.Call[bool](t_, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}

// Enables and disables grammar checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449166-grammarcheckingenabled?language=objc
func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setGrammarCheckingEnabled:"), value)
}

// A Boolean value that indicates whether the receiver allows undo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc
func (t_ TextView) AllowsUndo() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsUndo"))
	return rv
}

// A Boolean value that indicates whether the receiver allows undo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449450-allowsundo?language=objc
func (t_ TextView) SetAllowsUndo(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsUndo:"), value)
}

// The attributes used to draw marked text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc
func (t_ TextView) MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.Call[map[foundation.AttributedStringKey]objc.Object](t_, objc.Sel("markedTextAttributes"))
	return rv
}

// The attributes used to draw marked text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449179-markedtextattributes?language=objc
func (t_ TextView) SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setMarkedTextAttributes:"), value)
}

// The data types that the receiver accepts as the destination view of a dragging operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449234-acceptabledragtypes?language=objc
func (t_ TextView) AcceptableDragTypes() []PasteboardType {
	rv := objc.Call[[]PasteboardType](t_, objc.Sel("acceptableDragTypes"))
	return rv
}

// The origin of the receiver’s text container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449477-textcontainerorigin?language=objc
func (t_ TextView) TextContainerOrigin() foundation.Point {
	rv := objc.Call[foundation.Point](t_, objc.Sel("textContainerOrigin"))
	return rv
}

// A Boolean value that indicates whether to use the find bar for this text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc
func (t_ TextView) UsesFindBar() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesFindBar"))
	return rv
}

// A Boolean value that indicates whether to use the find bar for this text view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449456-usesfindbar?language=objc
func (t_ TextView) SetUsesFindBar(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesFindBar:"), value)
}

// The receiver’s text storage object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/3824762-textcontentstorage?language=objc
func (t_ TextView) TextContentStorage() TextContentStorage {
	rv := objc.Call[TextContentStorage](t_, objc.Sel("textContentStorage"))
	return rv
}

// The selection granularity for subsequent extension of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449165-selectiongranularity?language=objc
func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := objc.Call[SelectionGranularity](t_, objc.Sel("selectionGranularity"))
	return rv
}

// The selection granularity for subsequent extension of a selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449165-selectiongranularity?language=objc
func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectionGranularity:"), value)
}

// A Boolean value that indicates whether this text view uses the inspector bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc
func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.Call[bool](t_, objc.Sel("usesInspectorBar"))
	return rv
}

// A Boolean value that indicates whether this text view uses the inspector bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextview/1449407-usesinspectorbar?language=objc
func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setUsesInspectorBar:"), value)
}
