// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

var TextViewClass = _TextViewClass{objc.GetClass("NSTextView")}

type _TextViewClass struct {
	objc.Class
}

type ITextView interface {
	IText
	ReplaceTextContainer(newContainer ITextContainer)
	InvalidateTextContainerOrigin()
	ChangeDocumentBackgroundColor(sender objc.IObject)
	SetNeedsDisplayInRectAvoidAdditionalLayout(rect foundation.Rect, flag bool)
	DrawInsertionPointInRectColorTurnedOn(rect foundation.Rect, color IColor, flag bool)
	DrawViewBackgroundInRect(rect foundation.Rect)
	SetConstrainedFrameSize(desiredSize foundation.Size)
	CleanUpAfterDragOperation()
	ShowFindIndicatorForRange(charRange foundation.Range)
	SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ foundation.Range)
	Outline(sender objc.IObject)
	ToggleAutomaticQuoteSubstitution(sender objc.IObject)
	ToggleAutomaticLinkDetection(sender objc.IObject)
	ToggleAutomaticTextCompletion(sender objc.IObject)
	SetSelectedRangeAffinityStillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool)
	SetSelectedRangesAffinityStillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool)
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	CharacterIndexForInsertionAtPoint(point foundation.Point) uint
	UpdateCandidates()
	PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType
	ReadSelectionFromPasteboard(pboard IPasteboard) bool
	ReadSelectionFromPasteboardType(pboard IPasteboard, type_ PasteboardType) bool
	WriteSelectionToPasteboardType(pboard IPasteboard, type_ PasteboardType) bool
	WriteSelectionToPasteboardTypes(pboard IPasteboard, types []PasteboardType) bool
	AlignJustified(sender objc.IObject)
	ChangeAttributes(sender objc.IObject)
	ChangeColor(sender objc.IObject)
	SetAlignmentRange(alignment TextAlignment, range_ foundation.Range)
	UseStandardKerning(sender objc.IObject)
	LowerBaseline(sender objc.IObject)
	RaiseBaseline(sender objc.IObject)
	TurnOffKerning(sender objc.IObject)
	LoosenKerning(sender objc.IObject)
	TightenKerning(sender objc.IObject)
	UseStandardLigatures(sender objc.IObject)
	TurnOffLigatures(sender objc.IObject)
	UseAllLigatures(sender objc.IObject)
	ClickedOnLinkAtIndex(link objc.IObject, charIndex uint)
	PasteAsPlainText(sender objc.IObject)
	PasteAsRichText(sender objc.IObject)
	BreakUndoCoalescing()
	UpdateFontPanel()
	UpdateRuler()
	UpdateDragTypeRegistration()
	SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range
	ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.Range, replacementString string) bool
	ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool
	DidChangeText()
	SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range
	SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string
	SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.Range, beforeString *foundation.String, afterString *foundation.String)
	ToggleSmartInsertDelete(sender objc.IObject)
	ToggleContinuousSpellChecking(sender objc.IObject)
	ToggleGrammarChecking(sender objc.IObject)
	SetSpellingStateRange(value int, charRange foundation.Range)
	OrderFrontSharingServicePicker(sender objc.IObject)
	DragImageForSelectionWithEventOrigin(event IEvent, origin *foundation.Point) Image
	DragOperationForDraggingInfoType(dragInfo IDraggingInfo, type_ PasteboardType) DragOperation
	DragOperationForDraggingInfo0Type(dragInfo objc.IObject, type_ PasteboardType) DragOperation
	DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool
	StartSpeaking(sender objc.IObject)
	StopSpeaking(sender objc.IObject)
	PerformFindPanelAction(sender objc.IObject)
	OrderFrontLinkPanel(sender objc.IObject)
	OrderFrontListPanel(sender objc.IObject)
	OrderFrontSpacingPanel(sender objc.IObject)
	OrderFrontTablePanel(sender objc.IObject)
	OrderFrontSubstitutionsPanel(sender objc.IObject)
	Complete(sender objc.IObject)
	CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.Range, index *int) []string
	InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.Range, movement int, flag bool)
	CheckTextInDocument(sender objc.IObject)
	CheckTextInSelection(sender objc.IObject)
	CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject)
	HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int)
	ToggleAutomaticDashSubstitution(sender objc.IObject)
	ToggleAutomaticDataDetection(sender objc.IObject)
	ToggleAutomaticSpellingCorrection(sender objc.IObject)
	ToggleAutomaticTextReplacement(sender objc.IObject)
	PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool
	UpdateQuickLookPreviewPanel()
	ToggleQuickLookPreviewPanel(sender objc.IObject)
	ChangeLayoutOrientation(sender objc.IObject)
	SetLayoutOrientation(orientation TextLayoutOrientation)
	UpdateTextTouchBarItems()
	UpdateTouchBarItemIdentifiers()
	TextContainer() TextContainer
	SetTextContainer(value ITextContainer)
	TextContainerInset() foundation.Size
	SetTextContainerInset(value foundation.Size)
	TextContainerOrigin() foundation.Point
	TextLayoutManager() TextLayoutManager
	LayoutManager() LayoutManager
	TextContentStorage() TextContentStorage
	TextStorage() TextStorage
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	ShouldDrawInsertionPoint() bool
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	DefaultParagraphStyle() ParagraphStyle
	SetDefaultParagraphStyle(value IParagraphStyle)
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	IsAutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	IsAutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	IsAutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)
	UsesRuler() bool
	SetUsesRuler(value bool)
	SetRulerVisible(value bool)
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)
	SelectedRanges() []foundation.Value
	SetSelectedRanges(value []foundation.IValue)
	SelectionAffinity() SelectionAffinity
	SelectionGranularity() SelectionGranularity
	SetSelectionGranularity(value SelectionGranularity)
	InsertionPointColor() Color
	SetInsertionPointColor(value IColor)
	SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object
	SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	ReadablePasteboardTypes() []PasteboardType
	WritablePasteboardTypes() []PasteboardType
	TypingAttributes() map[foundation.AttributedStringKey]objc.Object
	SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject)
	IsCoalescingUndo() bool
	AcceptableDragTypes() []PasteboardType
	RangeForUserCharacterAttributeChange() foundation.Range
	RangesForUserCharacterAttributeChange() []foundation.Value
	RangeForUserParagraphAttributeChange() foundation.Range
	RangesForUserParagraphAttributeChange() []foundation.Value
	RangeForUserTextChange() foundation.Range
	RangesForUserTextChange() []foundation.Value
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	IsContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	SpellCheckerDocumentTag() int
	IsGrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	RangeForUserCompletion() foundation.Range
	EnabledTextCheckingTypes() foundation.TextCheckingTypes
	SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes)
	IsAutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	IsAutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	IsAutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	IsAutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	IsIncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
}

type TextView struct {
	Text
}

func MakeTextView(ptr unsafe.Pointer) TextView {
	return TextView{
		Text: MakeText(ptr),
	}
}

func (t_ TextView) InitWithFrameTextContainer(frameRect foundation.Rect, container ITextContainer) TextView {
	rv := objc.CallMethod[TextView](t_, objc.GetSelector("initWithFrame:textContainer:"), frameRect, objc.ExtractPtr(container))
	return rv
}

func TextView_InitWithFrameTextContainer(frameRect foundation.Rect, container ITextContainer) TextView {
	return TextViewClass.Alloc().InitWithFrameTextContainer(frameRect, container)
}

func (t_ TextView) InitWithFrame(frameRect foundation.Rect) TextView {
	rv := objc.CallMethod[TextView](t_, objc.GetSelector("initWithFrame:"), frameRect)
	return rv
}

func TextView_InitWithFrame(frameRect foundation.Rect) TextView {
	return TextViewClass.Alloc().InitWithFrame(frameRect)
}

func (tc _TextViewClass) FieldEditor() TextView {
	rv := objc.CallMethod[TextView](tc, objc.GetSelector("fieldEditor"))
	return rv
}

func TextView_FieldEditor() TextView {
	return TextViewClass.FieldEditor()
}

func (t_ TextView) Init() TextView {
	rv := objc.CallMethod[TextView](t_, objc.GetSelector("init"))
	return rv
}

func TextView_Init() TextView {
	return TextViewClass.Alloc().Init()
}

func (tc _TextViewClass) Alloc() TextView {
	rv := objc.CallMethod[TextView](tc, objc.GetSelector("alloc"))
	return rv
}

func TextView_Alloc() TextView {
	return TextViewClass.Alloc()
}

func (tc _TextViewClass) New() TextView {
	rv := objc.CallMethod[TextView](tc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewTextView() TextView {
	return TextViewClass.New()
}

func TextView_New() TextView {
	return TextViewClass.New()
}

func (tc _TextViewClass) RegisterForServices() {
	objc.CallMethod[objc.Void](tc, objc.GetSelector("registerForServices"))
}

func TextView_RegisterForServices() {
	TextViewClass.RegisterForServices()
}

func (t_ TextView) ReplaceTextContainer(newContainer ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("replaceTextContainer:"), objc.ExtractPtr(newContainer))
}

func (t_ TextView) InvalidateTextContainerOrigin() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("invalidateTextContainerOrigin"))
}

func (t_ TextView) ChangeDocumentBackgroundColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("changeDocumentBackgroundColor:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetNeedsDisplayInRectAvoidAdditionalLayout(rect foundation.Rect, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}

func (t_ TextView) DrawInsertionPointInRectColorTurnedOn(rect foundation.Rect, color IColor, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawInsertionPointInRect:color:turnedOn:"), rect, objc.ExtractPtr(color), flag)
}

func (t_ TextView) DrawViewBackgroundInRect(rect foundation.Rect) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("drawViewBackgroundInRect:"), rect)
}

func (t_ TextView) SetConstrainedFrameSize(desiredSize foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setConstrainedFrameSize:"), desiredSize)
}

func (t_ TextView) CleanUpAfterDragOperation() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("cleanUpAfterDragOperation"))
}

func (t_ TextView) ShowFindIndicatorForRange(charRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("showFindIndicatorForRange:"), charRange)
}

func (tc _TextViewClass) ScrollableDocumentContentTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.GetSelector("scrollableDocumentContentTextView"))
	return rv
}

func TextView_ScrollableDocumentContentTextView() ScrollView {
	return TextViewClass.ScrollableDocumentContentTextView()
}

func (tc _TextViewClass) ScrollablePlainDocumentContentTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.GetSelector("scrollablePlainDocumentContentTextView"))
	return rv
}

func TextView_ScrollablePlainDocumentContentTextView() ScrollView {
	return TextViewClass.ScrollablePlainDocumentContentTextView()
}

func (tc _TextViewClass) ScrollableTextView() ScrollView {
	rv := objc.CallMethod[ScrollView](tc, objc.GetSelector("scrollableTextView"))
	return rv
}

func TextView_ScrollableTextView() ScrollView {
	return TextViewClass.ScrollableTextView()
}

func (t_ TextView) SetBaseWritingDirectionRange(writingDirection WritingDirection, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setBaseWritingDirection:range:"), writingDirection, range_)
}

func (t_ TextView) Outline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("outline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticQuoteSubstitution(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticQuoteSubstitution:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticLinkDetection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticLinkDetection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticTextCompletion(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticTextCompletion:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetSelectedRangeAffinityStillSelecting(charRange foundation.Range, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}

func (t_ TextView) SetSelectedRangesAffinityStillSelecting(ranges []foundation.IValue, affinity SelectionAffinity, stillSelectingFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedRanges:affinity:stillSelecting:"), ranges, affinity, stillSelectingFlag)
}

func (t_ TextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}

func (t_ TextView) CharacterIndexForInsertionAtPoint(point foundation.Point) uint {
	rv := objc.CallMethod[uint](t_, objc.GetSelector("characterIndexForInsertionAtPoint:"), point)
	return rv
}

func (t_ TextView) UpdateCandidates() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateCandidates"))
}

func (t_ TextView) PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []PasteboardType, allowedTypes []PasteboardType) PasteboardType {
	rv := objc.CallMethod[PasteboardType](t_, objc.GetSelector("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), availableTypes, allowedTypes)
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboard(pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("readSelectionFromPasteboard:"), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TextView) ReadSelectionFromPasteboardType(pboard IPasteboard, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("readSelectionFromPasteboard:type:"), objc.ExtractPtr(pboard), type_)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboardType(pboard IPasteboard, type_ PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("writeSelectionToPasteboard:type:"), objc.ExtractPtr(pboard), type_)
	return rv
}

func (t_ TextView) WriteSelectionToPasteboardTypes(pboard IPasteboard, types []PasteboardType) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("writeSelectionToPasteboard:types:"), objc.ExtractPtr(pboard), types)
	return rv
}

func (t_ TextView) AlignJustified(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("alignJustified:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeAttributes(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("changeAttributes:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeColor(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("changeColor:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetAlignmentRange(alignment TextAlignment, range_ foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAlignment:range:"), alignment, range_)
}

func (t_ TextView) UseStandardKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("useStandardKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) LowerBaseline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("lowerBaseline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) RaiseBaseline(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("raiseBaseline:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TurnOffKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("turnOffKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) LoosenKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("loosenKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TightenKerning(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("tightenKerning:"), objc.ExtractPtr(sender))
}

func (t_ TextView) UseStandardLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("useStandardLigatures:"), objc.ExtractPtr(sender))
}

func (t_ TextView) TurnOffLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("turnOffLigatures:"), objc.ExtractPtr(sender))
}

func (t_ TextView) UseAllLigatures(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("useAllLigatures:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ClickedOnLinkAtIndex(link objc.IObject, charIndex uint) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("clickedOnLink:atIndex:"), objc.ExtractPtr(link), charIndex)
}

func (t_ TextView) PasteAsPlainText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("pasteAsPlainText:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PasteAsRichText(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("pasteAsRichText:"), objc.ExtractPtr(sender))
}

func (t_ TextView) BreakUndoCoalescing() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("breakUndoCoalescing"))
}

func (t_ TextView) UpdateFontPanel() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateFontPanel"))
}

func (t_ TextView) UpdateRuler() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateRuler"))
}

func (t_ TextView) UpdateDragTypeRegistration() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateDragTypeRegistration"))
}

func (t_ TextView) SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.Range, granularity SelectionGranularity) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return rv
}

func (t_ TextView) ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldChangeTextInRange:replacementString:"), affectedCharRange, replacementString)
	return rv
}

func (t_ TextView) ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldChangeTextInRanges:replacementStrings:"), affectedRanges, replacementStrings)
	return rv
}

func (t_ TextView) DidChangeText() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("didChangeText"))
}

func (t_ TextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return rv
}

func (t_ TextView) SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("smartInsertAfterStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.Range) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("smartInsertBeforeStringForString:replacingRange:"), pasteString, charRangeToReplace)
	return rv
}

func (t_ TextView) SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.Range, beforeString *foundation.String, afterString *foundation.String) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("smartInsertForString:replacingRange:beforeString:afterString:"), pasteString, charRangeToReplace, beforeString, afterString)
}

func (t_ TextView) ToggleSmartInsertDelete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleSmartInsertDelete:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleContinuousSpellChecking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleContinuousSpellChecking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleGrammarChecking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleGrammarChecking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetSpellingStateRange(value int, charRange foundation.Range) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSpellingState:range:"), value, charRange)
}

func (t_ TextView) OrderFrontSharingServicePicker(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontSharingServicePicker:"), objc.ExtractPtr(sender))
}

func (t_ TextView) DragImageForSelectionWithEventOrigin(event IEvent, origin *foundation.Point) Image {
	rv := objc.CallMethod[Image](t_, objc.GetSelector("dragImageForSelectionWithEvent:origin:"), objc.ExtractPtr(event), origin)
	return rv
}

func (t_ TextView) DragOperationForDraggingInfoType(dragInfo IDraggingInfo, type_ PasteboardType) DragOperation {
	po := objc.WrapAsProtocol("NSDraggingInfo", dragInfo)
	rv := objc.CallMethod[DragOperation](t_, objc.GetSelector("dragOperationForDraggingInfo:type:"), po, type_)
	return rv
}

func (t_ TextView) DragOperationForDraggingInfo0Type(dragInfo objc.IObject, type_ PasteboardType) DragOperation {
	rv := objc.CallMethod[DragOperation](t_, objc.GetSelector("dragOperationForDraggingInfo:type:"), objc.ExtractPtr(dragInfo), type_)
	return rv
}

func (t_ TextView) DragSelectionWithEventOffsetSlideBack(event IEvent, mouseOffset foundation.Size, slideBack bool) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("dragSelectionWithEvent:offset:slideBack:"), objc.ExtractPtr(event), mouseOffset, slideBack)
	return rv
}

func (t_ TextView) StartSpeaking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("startSpeaking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) StopSpeaking(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("stopSpeaking:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PerformFindPanelAction(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("performFindPanelAction:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontLinkPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontLinkPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontListPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontListPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontSpacingPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontSpacingPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontTablePanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontTablePanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) OrderFrontSubstitutionsPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("orderFrontSubstitutionsPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) Complete(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("complete:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return rv
}

func (t_ TextView) InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.Range, movement int, flag bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("insertCompletion:forPartialWordRange:movement:isFinal:"), word, charRange, movement, flag)
}

func (t_ TextView) CheckTextInDocument(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("checkTextInDocument:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CheckTextInSelection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("checkTextInSelection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) CheckTextInRangeTypesOptions(range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("checkTextInRange:types:options:"), range_, checkingTypes, options)
}

func (t_ TextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.ITextCheckingResult, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, orthography foundation.IOrthography, wordCount int) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), results, range_, checkingTypes, options, objc.ExtractPtr(orthography), wordCount)
}

func (t_ TextView) ToggleAutomaticDashSubstitution(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticDashSubstitution:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticDataDetection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticDataDetection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticSpellingCorrection(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticSpellingCorrection:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ToggleAutomaticTextReplacement(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleAutomaticTextReplacement:"), objc.ExtractPtr(sender))
}

func (t_ TextView) PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.Range, attributedString foundation.IAttributedString) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("performValidatedReplacementInRange:withAttributedString:"), range_, objc.ExtractPtr(attributedString))
	return rv
}

func (t_ TextView) UpdateQuickLookPreviewPanel() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateQuickLookPreviewPanel"))
}

func (t_ TextView) ToggleQuickLookPreviewPanel(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("toggleQuickLookPreviewPanel:"), objc.ExtractPtr(sender))
}

func (t_ TextView) ChangeLayoutOrientation(sender objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("changeLayoutOrientation:"), objc.ExtractPtr(sender))
}

func (t_ TextView) SetLayoutOrientation(orientation TextLayoutOrientation) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLayoutOrientation:"), orientation)
}

func (t_ TextView) UpdateTextTouchBarItems() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateTextTouchBarItems"))
}

func (t_ TextView) UpdateTouchBarItemIdentifiers() {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("updateTouchBarItemIdentifiers"))
}

func (tc _TextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.CallMethod[bool](tc, objc.GetSelector("stronglyReferencesTextStorage"))
	return rv
}

func TextView_StronglyReferencesTextStorage() bool {
	return TextViewClass.StronglyReferencesTextStorage()
}

func (t_ TextView) TextContainer() TextContainer {
	rv := objc.CallMethod[TextContainer](t_, objc.GetSelector("textContainer"))
	return rv
}

func (t_ TextView) SetTextContainer(value ITextContainer) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextContainer:"), objc.ExtractPtr(value))
}

func (t_ TextView) TextContainerInset() foundation.Size {
	rv := objc.CallMethod[foundation.Size](t_, objc.GetSelector("textContainerInset"))
	return rv
}

func (t_ TextView) SetTextContainerInset(value foundation.Size) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTextContainerInset:"), value)
}

func (t_ TextView) TextContainerOrigin() foundation.Point {
	rv := objc.CallMethod[foundation.Point](t_, objc.GetSelector("textContainerOrigin"))
	return rv
}

func (t_ TextView) TextLayoutManager() TextLayoutManager {
	rv := objc.CallMethod[TextLayoutManager](t_, objc.GetSelector("textLayoutManager"))
	return rv
}

func (t_ TextView) LayoutManager() LayoutManager {
	rv := objc.CallMethod[LayoutManager](t_, objc.GetSelector("layoutManager"))
	return rv
}

func (t_ TextView) TextContentStorage() TextContentStorage {
	rv := objc.CallMethod[TextContentStorage](t_, objc.GetSelector("textContentStorage"))
	return rv
}

func (t_ TextView) TextStorage() TextStorage {
	rv := objc.CallMethod[TextStorage](t_, objc.GetSelector("textStorage"))
	return rv
}

func (t_ TextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsDocumentBackgroundColorChange"))
	return rv
}

func (t_ TextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsDocumentBackgroundColorChange:"), value)
}

func (t_ TextView) ShouldDrawInsertionPoint() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("shouldDrawInsertionPoint"))
	return rv
}

func (t_ TextView) AllowedInputSourceLocales() []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("allowedInputSourceLocales"))
	return rv
}

func (t_ TextView) SetAllowedInputSourceLocales(value []string) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowedInputSourceLocales:"), value)
}

func (t_ TextView) AllowsUndo() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsUndo"))
	return rv
}

func (t_ TextView) SetAllowsUndo(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsUndo:"), value)
}

func (t_ TextView) DefaultParagraphStyle() ParagraphStyle {
	rv := objc.CallMethod[ParagraphStyle](t_, objc.GetSelector("defaultParagraphStyle"))
	return rv
}

func (t_ TextView) SetDefaultParagraphStyle(value IParagraphStyle) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDefaultParagraphStyle:"), objc.ExtractPtr(value))
}

func (t_ TextView) AllowsImageEditing() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsImageEditing"))
	return rv
}

func (t_ TextView) SetAllowsImageEditing(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsImageEditing:"), value)
}

func (t_ TextView) IsAutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticQuoteSubstitutionEnabled:"), value)
}

func (t_ TextView) IsAutomaticLinkDetectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticLinkDetectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticLinkDetectionEnabled:"), value)
}

func (t_ TextView) DisplaysLinkToolTips() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("displaysLinkToolTips"))
	return rv
}

func (t_ TextView) SetDisplaysLinkToolTips(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setDisplaysLinkToolTips:"), value)
}

func (t_ TextView) IsAutomaticTextCompletionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticTextCompletionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticTextCompletionEnabled:"), value)
}

func (t_ TextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}

func (t_ TextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}

func (t_ TextView) UsesRolloverButtonForSelection() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesRolloverButtonForSelection"))
	return rv
}

func (t_ TextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesRolloverButtonForSelection:"), value)
}

func (t_ TextView) UsesRuler() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesRuler"))
	return rv
}

func (t_ TextView) SetUsesRuler(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesRuler:"), value)
}

func (t_ TextView) SetRulerVisible(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setRulerVisible:"), value)
}

func (t_ TextView) UsesInspectorBar() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesInspectorBar"))
	return rv
}

func (t_ TextView) SetUsesInspectorBar(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesInspectorBar:"), value)
}

func (t_ TextView) SelectedRanges() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("selectedRanges"))
	// TODO: convert slice items...
	return rv
}

func (t_ TextView) SetSelectedRanges(value []foundation.IValue) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedRanges:"), value)
}

func (t_ TextView) SelectionAffinity() SelectionAffinity {
	rv := objc.CallMethod[SelectionAffinity](t_, objc.GetSelector("selectionAffinity"))
	return rv
}

func (t_ TextView) SelectionGranularity() SelectionGranularity {
	rv := objc.CallMethod[SelectionGranularity](t_, objc.GetSelector("selectionGranularity"))
	return rv
}

func (t_ TextView) SetSelectionGranularity(value SelectionGranularity) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectionGranularity:"), value)
}

func (t_ TextView) InsertionPointColor() Color {
	rv := objc.CallMethod[Color](t_, objc.GetSelector("insertionPointColor"))
	return rv
}

func (t_ TextView) SetInsertionPointColor(value IColor) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setInsertionPointColor:"), objc.ExtractPtr(value))
}

func (t_ TextView) SelectedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("selectedTextAttributes"))
	return rv
}

func (t_ TextView) SetSelectedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSelectedTextAttributes:"), value)
}

func (t_ TextView) MarkedTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("markedTextAttributes"))
	return rv
}

func (t_ TextView) SetMarkedTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setMarkedTextAttributes:"), value)
}

func (t_ TextView) LinkTextAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("linkTextAttributes"))
	return rv
}

func (t_ TextView) SetLinkTextAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setLinkTextAttributes:"), value)
}

func (t_ TextView) ReadablePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.GetSelector("readablePasteboardTypes"))
	return rv
}

func (t_ TextView) WritablePasteboardTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.GetSelector("writablePasteboardTypes"))
	return rv
}

func (t_ TextView) TypingAttributes() map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("typingAttributes"))
	return rv
}

func (t_ TextView) SetTypingAttributes(value map[foundation.AttributedStringKey]objc.IObject) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setTypingAttributes:"), value)
}

func (t_ TextView) IsCoalescingUndo() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isCoalescingUndo"))
	return rv
}

func (t_ TextView) AcceptableDragTypes() []PasteboardType {
	rv := objc.CallMethod[[]PasteboardType](t_, objc.GetSelector("acceptableDragTypes"))
	return rv
}

func (t_ TextView) RangeForUserCharacterAttributeChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rangeForUserCharacterAttributeChange"))
	return rv
}

func (t_ TextView) RangesForUserCharacterAttributeChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("rangesForUserCharacterAttributeChange"))
	// TODO: convert slice items...
	return rv
}

func (t_ TextView) RangeForUserParagraphAttributeChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rangeForUserParagraphAttributeChange"))
	return rv
}

func (t_ TextView) RangesForUserParagraphAttributeChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("rangesForUserParagraphAttributeChange"))
	// TODO: convert slice items...
	return rv
}

func (t_ TextView) RangeForUserTextChange() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rangeForUserTextChange"))
	return rv
}

func (t_ TextView) RangesForUserTextChange() []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("rangesForUserTextChange"))
	// TODO: convert slice items...
	return rv
}

func (t_ TextView) SmartInsertDeleteEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("smartInsertDeleteEnabled"))
	return rv
}

func (t_ TextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setSmartInsertDeleteEnabled:"), value)
}

func (t_ TextView) IsContinuousSpellCheckingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isContinuousSpellCheckingEnabled"))
	return rv
}

func (t_ TextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setContinuousSpellCheckingEnabled:"), value)
}

func (t_ TextView) SpellCheckerDocumentTag() int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("spellCheckerDocumentTag"))
	return rv
}

func (t_ TextView) IsGrammarCheckingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isGrammarCheckingEnabled"))
	return rv
}

func (t_ TextView) SetGrammarCheckingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setGrammarCheckingEnabled:"), value)
}

func (t_ TextView) AcceptsGlyphInfo() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("acceptsGlyphInfo"))
	return rv
}

func (t_ TextView) SetAcceptsGlyphInfo(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAcceptsGlyphInfo:"), value)
}

func (t_ TextView) UsesFindPanel() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesFindPanel"))
	return rv
}

func (t_ TextView) SetUsesFindPanel(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesFindPanel:"), value)
}

func (t_ TextView) RangeForUserCompletion() foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("rangeForUserCompletion"))
	return rv
}

func (t_ TextView) EnabledTextCheckingTypes() foundation.TextCheckingTypes {
	rv := objc.CallMethod[foundation.TextCheckingTypes](t_, objc.GetSelector("enabledTextCheckingTypes"))
	return rv
}

func (t_ TextView) SetEnabledTextCheckingTypes(value foundation.TextCheckingTypes) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setEnabledTextCheckingTypes:"), value)
}

func (t_ TextView) IsAutomaticDashSubstitutionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticDashSubstitutionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticDashSubstitutionEnabled:"), value)
}

func (t_ TextView) IsAutomaticDataDetectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticDataDetectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticDataDetectionEnabled:"), value)
}

func (t_ TextView) IsAutomaticSpellingCorrectionEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticSpellingCorrectionEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticSpellingCorrectionEnabled:"), value)
}

func (t_ TextView) IsAutomaticTextReplacementEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isAutomaticTextReplacementEnabled"))
	return rv
}

func (t_ TextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAutomaticTextReplacementEnabled:"), value)
}

func (t_ TextView) UsesFindBar() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("usesFindBar"))
	return rv
}

func (t_ TextView) SetUsesFindBar(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setUsesFindBar:"), value)
}

func (t_ TextView) IsIncrementalSearchingEnabled() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("isIncrementalSearchingEnabled"))
	return rv
}

func (t_ TextView) SetIncrementalSearchingEnabled(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setIncrementalSearchingEnabled:"), value)
}

func (t_ TextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("allowsCharacterPickerTouchBarItem"))
	return rv
}

func (t_ TextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("setAllowsCharacterPickerTouchBarItem:"), value)
}
