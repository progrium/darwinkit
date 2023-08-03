// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type ITextViewDelegate interface {
	ITextDelegate
	ImplementsUndoManagerForTextView() bool
	// optional
	UndoManagerForTextView(view TextView) foundation.IUndoManager
	ImplementsTextViewWillDisplayToolTipForCharacterAtIndex() bool
	// optional
	TextViewWillDisplayToolTipForCharacterAtIndex(textView TextView, tooltip string, characterIndex uint) string
	ImplementsTextViewURLForContentsOfTextAttachmentAtIndex() bool
	// optional
	TextViewURLForContentsOfTextAttachmentAtIndex(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL
	ImplementsTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool
	// optional
	TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	ImplementsTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool
	// optional
	TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue
	ImplementsTextViewDidChangeSelection() bool
	// optional
	TextViewDidChangeSelection(notification foundation.Notification)
	ImplementsTextViewCandidatesForSelectedRange_() bool
	// optional
	TextViewCandidatesForSelectedRange_(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	ImplementsTextViewCandidatesForSelectedRange() bool
	// optional
	TextViewCandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject
	ImplementsTextViewShouldSelectCandidateAtIndex() bool
	// optional
	TextViewShouldSelectCandidateAtIndex(textView TextView, index uint) bool
	ImplementsTextViewShouldUpdateTouchBarItemIdentifiers() bool
	// optional
	TextViewShouldUpdateTouchBarItemIdentifiers(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	ImplementsTextViewShouldChangeTextInRangeReplacementString() bool
	// optional
	TextViewShouldChangeTextInRangeReplacementString(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	ImplementsTextViewShouldChangeTextInRangesReplacementStrings() bool
	// optional
	TextViewShouldChangeTextInRangesReplacementStrings(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	ImplementsTextViewShouldChangeTypingAttributesToAttributes() bool
	// optional
	TextViewShouldChangeTypingAttributesToAttributes(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject
	ImplementsTextViewDidChangeTypingAttributes() bool
	// optional
	TextViewDidChangeTypingAttributes(notification foundation.Notification)
	ImplementsTextViewClickedOnLinkAtIndex() bool
	// optional
	TextViewClickedOnLinkAtIndex(textView TextView, link objc.Object, charIndex uint) bool
	ImplementsTextViewShouldSetSpellingStateRange() bool
	// optional
	TextViewShouldSetSpellingStateRange(textView TextView, value int, affectedCharRange foundation.Range) int
	ImplementsTextViewWillCheckTextInRangeOptionsTypes() bool
	// optional
	TextViewWillCheckTextInRangeOptionsTypes(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	ImplementsTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool
	// optional
	TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
	ImplementsTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool
	// optional
	TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView TextView, words []string, charRange foundation.Range, index *int) []string
	ImplementsTextViewWillShowSharingServicePickerForItems() bool
	// optional
	TextViewWillShowSharingServicePickerForItems(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker
	ImplementsTextViewDoCommandBySelector() bool
	// optional
	TextViewDoCommandBySelector(textView TextView, commandSelector objc.Selector) bool
	ImplementsTextViewMenuForEventAtIndex() bool
	// optional
	TextViewMenuForEventAtIndex(view TextView, menu Menu, event Event, charIndex uint) IMenu
	ImplementsTextViewClickedOnLink() bool
	// optional
	// deprecated
	TextViewClickedOnLink(textView TextView, link objc.Object) bool
}

type TextViewDelegate struct {
	TextDelegate
	_UndoManagerForTextView                                             func(view TextView) foundation.IUndoManager
	_TextViewWillDisplayToolTipForCharacterAtIndex                      func(textView TextView, tooltip string, characterIndex uint) string
	_TextViewURLForContentsOfTextAttachmentAtIndex                      func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL
	_TextViewWillChangeSelectionFromCharacterRangeToCharacterRange      func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range
	_TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges    func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue
	_TextViewDidChangeSelection                                         func(notification foundation.Notification)
	_TextViewCandidatesForSelectedRange_                                func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult
	_TextViewCandidatesForSelectedRange                                 func(textView TextView, selectedRange foundation.Range) []objc.IObject
	_TextViewShouldSelectCandidateAtIndex                               func(textView TextView, index uint) bool
	_TextViewShouldUpdateTouchBarItemIdentifiers                        func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier
	_TextViewShouldChangeTextInRangeReplacementString                   func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool
	_TextViewShouldChangeTextInRangesReplacementStrings                 func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool
	_TextViewShouldChangeTypingAttributesToAttributes                   func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject
	_TextViewDidChangeTypingAttributes                                  func(notification foundation.Notification)
	_TextViewClickedOnLinkAtIndex                                       func(textView TextView, link objc.Object, charIndex uint) bool
	_TextViewShouldSetSpellingStateRange                                func(textView TextView, value int, affectedCharRange foundation.Range) int
	_TextViewWillCheckTextInRangeOptionsTypes                           func(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject
	_TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount func(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult
	_TextViewCompletionsForPartialWordRangeIndexOfSelectedItem          func(textView TextView, words []string, charRange foundation.Range, index *int) []string
	_TextViewWillShowSharingServicePickerForItems                       func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker
	_TextViewDoCommandBySelector                                        func(textView TextView, commandSelector objc.Selector) bool
	_TextViewMenuForEventAtIndex                                        func(view TextView, menu Menu, event Event, charIndex uint) IMenu
	_TextViewClickedOnLink                                              func(textView TextView, link objc.Object) bool
}

func (di *TextViewDelegate) ImplementsUndoManagerForTextView() bool {
	return di._UndoManagerForTextView != nil
}

func (di *TextViewDelegate) SetUndoManagerForTextView(f func(view TextView) foundation.IUndoManager) {
	di._UndoManagerForTextView = f
}

func (di *TextViewDelegate) UndoManagerForTextView(view TextView) foundation.IUndoManager {
	return di._UndoManagerForTextView(view)
}
func (di *TextViewDelegate) ImplementsTextViewWillDisplayToolTipForCharacterAtIndex() bool {
	return di._TextViewWillDisplayToolTipForCharacterAtIndex != nil
}

func (di *TextViewDelegate) SetTextViewWillDisplayToolTipForCharacterAtIndex(f func(textView TextView, tooltip string, characterIndex uint) string) {
	di._TextViewWillDisplayToolTipForCharacterAtIndex = f
}

func (di *TextViewDelegate) TextViewWillDisplayToolTipForCharacterAtIndex(textView TextView, tooltip string, characterIndex uint) string {
	return di._TextViewWillDisplayToolTipForCharacterAtIndex(textView, tooltip, characterIndex)
}
func (di *TextViewDelegate) ImplementsTextViewURLForContentsOfTextAttachmentAtIndex() bool {
	return di._TextViewURLForContentsOfTextAttachmentAtIndex != nil
}

func (di *TextViewDelegate) SetTextViewURLForContentsOfTextAttachmentAtIndex(f func(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL) {
	di._TextViewURLForContentsOfTextAttachmentAtIndex = f
}

func (di *TextViewDelegate) TextViewURLForContentsOfTextAttachmentAtIndex(textView TextView, textAttachment TextAttachment, charIndex uint) foundation.IURL {
	return di._TextViewURLForContentsOfTextAttachmentAtIndex(textView, textAttachment, charIndex)
}
func (di *TextViewDelegate) ImplementsTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool {
	return di._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange != nil
}

func (di *TextViewDelegate) SetTextViewWillChangeSelectionFromCharacterRangeToCharacterRange(f func(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range) {
	di._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange = f
}

func (di *TextViewDelegate) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView TextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	return di._TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView, oldSelectedCharRange, newSelectedCharRange)
}
func (di *TextViewDelegate) ImplementsTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool {
	return di._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges != nil
}

func (di *TextViewDelegate) SetTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(f func(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue) {
	di._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges = f
}

func (di *TextViewDelegate) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView TextView, oldSelectedCharRanges []foundation.Value, newSelectedCharRanges []foundation.Value) []foundation.IValue {
	return di._TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView, oldSelectedCharRanges, newSelectedCharRanges)
}
func (di *TextViewDelegate) ImplementsTextViewDidChangeSelection() bool {
	return di._TextViewDidChangeSelection != nil
}

func (di *TextViewDelegate) SetTextViewDidChangeSelection(f func(notification foundation.Notification)) {
	di._TextViewDidChangeSelection = f
}

func (di *TextViewDelegate) TextViewDidChangeSelection(notification foundation.Notification) {
	di._TextViewDidChangeSelection(notification)
}
func (di *TextViewDelegate) ImplementsTextViewCandidatesForSelectedRange_() bool {
	return di._TextViewCandidatesForSelectedRange_ != nil
}

func (di *TextViewDelegate) SetTextViewCandidatesForSelectedRange_(f func(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult) {
	di._TextViewCandidatesForSelectedRange_ = f
}

func (di *TextViewDelegate) TextViewCandidatesForSelectedRange_(textView TextView, candidates []foundation.TextCheckingResult, selectedRange foundation.Range) []foundation.ITextCheckingResult {
	return di._TextViewCandidatesForSelectedRange_(textView, candidates, selectedRange)
}
func (di *TextViewDelegate) ImplementsTextViewCandidatesForSelectedRange() bool {
	return di._TextViewCandidatesForSelectedRange != nil
}

func (di *TextViewDelegate) SetTextViewCandidatesForSelectedRange(f func(textView TextView, selectedRange foundation.Range) []objc.IObject) {
	di._TextViewCandidatesForSelectedRange = f
}

func (di *TextViewDelegate) TextViewCandidatesForSelectedRange(textView TextView, selectedRange foundation.Range) []objc.IObject {
	return di._TextViewCandidatesForSelectedRange(textView, selectedRange)
}
func (di *TextViewDelegate) ImplementsTextViewShouldSelectCandidateAtIndex() bool {
	return di._TextViewShouldSelectCandidateAtIndex != nil
}

func (di *TextViewDelegate) SetTextViewShouldSelectCandidateAtIndex(f func(textView TextView, index uint) bool) {
	di._TextViewShouldSelectCandidateAtIndex = f
}

func (di *TextViewDelegate) TextViewShouldSelectCandidateAtIndex(textView TextView, index uint) bool {
	return di._TextViewShouldSelectCandidateAtIndex(textView, index)
}
func (di *TextViewDelegate) ImplementsTextViewShouldUpdateTouchBarItemIdentifiers() bool {
	return di._TextViewShouldUpdateTouchBarItemIdentifiers != nil
}

func (di *TextViewDelegate) SetTextViewShouldUpdateTouchBarItemIdentifiers(f func(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier) {
	di._TextViewShouldUpdateTouchBarItemIdentifiers = f
}

func (di *TextViewDelegate) TextViewShouldUpdateTouchBarItemIdentifiers(textView TextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	return di._TextViewShouldUpdateTouchBarItemIdentifiers(textView, identifiers)
}
func (di *TextViewDelegate) ImplementsTextViewShouldChangeTextInRangeReplacementString() bool {
	return di._TextViewShouldChangeTextInRangeReplacementString != nil
}

func (di *TextViewDelegate) SetTextViewShouldChangeTextInRangeReplacementString(f func(textView TextView, affectedCharRange foundation.Range, replacementString string) bool) {
	di._TextViewShouldChangeTextInRangeReplacementString = f
}

func (di *TextViewDelegate) TextViewShouldChangeTextInRangeReplacementString(textView TextView, affectedCharRange foundation.Range, replacementString string) bool {
	return di._TextViewShouldChangeTextInRangeReplacementString(textView, affectedCharRange, replacementString)
}
func (di *TextViewDelegate) ImplementsTextViewShouldChangeTextInRangesReplacementStrings() bool {
	return di._TextViewShouldChangeTextInRangesReplacementStrings != nil
}

func (di *TextViewDelegate) SetTextViewShouldChangeTextInRangesReplacementStrings(f func(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool) {
	di._TextViewShouldChangeTextInRangesReplacementStrings = f
}

func (di *TextViewDelegate) TextViewShouldChangeTextInRangesReplacementStrings(textView TextView, affectedRanges []foundation.Value, replacementStrings []string) bool {
	return di._TextViewShouldChangeTextInRangesReplacementStrings(textView, affectedRanges, replacementStrings)
}
func (di *TextViewDelegate) ImplementsTextViewShouldChangeTypingAttributesToAttributes() bool {
	return di._TextViewShouldChangeTypingAttributesToAttributes != nil
}

func (di *TextViewDelegate) SetTextViewShouldChangeTypingAttributesToAttributes(f func(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject) {
	di._TextViewShouldChangeTypingAttributesToAttributes = f
}

func (di *TextViewDelegate) TextViewShouldChangeTypingAttributesToAttributes(textView TextView, oldTypingAttributes map[string]objc.Object, newTypingAttributes map[foundation.AttributedStringKey]objc.Object) map[foundation.AttributedStringKey]objc.IObject {
	return di._TextViewShouldChangeTypingAttributesToAttributes(textView, oldTypingAttributes, newTypingAttributes)
}
func (di *TextViewDelegate) ImplementsTextViewDidChangeTypingAttributes() bool {
	return di._TextViewDidChangeTypingAttributes != nil
}

func (di *TextViewDelegate) SetTextViewDidChangeTypingAttributes(f func(notification foundation.Notification)) {
	di._TextViewDidChangeTypingAttributes = f
}

func (di *TextViewDelegate) TextViewDidChangeTypingAttributes(notification foundation.Notification) {
	di._TextViewDidChangeTypingAttributes(notification)
}
func (di *TextViewDelegate) ImplementsTextViewClickedOnLinkAtIndex() bool {
	return di._TextViewClickedOnLinkAtIndex != nil
}

func (di *TextViewDelegate) SetTextViewClickedOnLinkAtIndex(f func(textView TextView, link objc.Object, charIndex uint) bool) {
	di._TextViewClickedOnLinkAtIndex = f
}

func (di *TextViewDelegate) TextViewClickedOnLinkAtIndex(textView TextView, link objc.Object, charIndex uint) bool {
	return di._TextViewClickedOnLinkAtIndex(textView, link, charIndex)
}
func (di *TextViewDelegate) ImplementsTextViewShouldSetSpellingStateRange() bool {
	return di._TextViewShouldSetSpellingStateRange != nil
}

func (di *TextViewDelegate) SetTextViewShouldSetSpellingStateRange(f func(textView TextView, value int, affectedCharRange foundation.Range) int) {
	di._TextViewShouldSetSpellingStateRange = f
}

func (di *TextViewDelegate) TextViewShouldSetSpellingStateRange(textView TextView, value int, affectedCharRange foundation.Range) int {
	return di._TextViewShouldSetSpellingStateRange(textView, value, affectedCharRange)
}
func (di *TextViewDelegate) ImplementsTextViewWillCheckTextInRangeOptionsTypes() bool {
	return di._TextViewWillCheckTextInRangeOptionsTypes != nil
}

func (di *TextViewDelegate) SetTextViewWillCheckTextInRangeOptionsTypes(f func(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject) {
	di._TextViewWillCheckTextInRangeOptionsTypes = f
}

func (di *TextViewDelegate) TextViewWillCheckTextInRangeOptionsTypes(view TextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.Object, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.IObject {
	return di._TextViewWillCheckTextInRangeOptionsTypes(view, range_, options, checkingTypes)
}
func (di *TextViewDelegate) ImplementsTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool {
	return di._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount != nil
}

func (di *TextViewDelegate) SetTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(f func(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult) {
	di._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount = f
}

func (di *TextViewDelegate) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view TextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.Object, results []foundation.TextCheckingResult, orthography foundation.Orthography, wordCount int) []foundation.ITextCheckingResult {
	return di._TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view, range_, checkingTypes, options, results, orthography, wordCount)
}
func (di *TextViewDelegate) ImplementsTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return di._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil
}

func (di *TextViewDelegate) SetTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

func (di *TextViewDelegate) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView, words, charRange, index)
}
func (di *TextViewDelegate) ImplementsTextViewWillShowSharingServicePickerForItems() bool {
	return di._TextViewWillShowSharingServicePickerForItems != nil
}

func (di *TextViewDelegate) SetTextViewWillShowSharingServicePickerForItems(f func(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker) {
	di._TextViewWillShowSharingServicePickerForItems = f
}

func (di *TextViewDelegate) TextViewWillShowSharingServicePickerForItems(textView TextView, servicePicker SharingServicePicker, items []objc.Object) ISharingServicePicker {
	return di._TextViewWillShowSharingServicePickerForItems(textView, servicePicker, items)
}
func (di *TextViewDelegate) ImplementsTextViewDoCommandBySelector() bool {
	return di._TextViewDoCommandBySelector != nil
}

func (di *TextViewDelegate) SetTextViewDoCommandBySelector(f func(textView TextView, commandSelector objc.Selector) bool) {
	di._TextViewDoCommandBySelector = f
}

func (di *TextViewDelegate) TextViewDoCommandBySelector(textView TextView, commandSelector objc.Selector) bool {
	return di._TextViewDoCommandBySelector(textView, commandSelector)
}
func (di *TextViewDelegate) ImplementsTextViewMenuForEventAtIndex() bool {
	return di._TextViewMenuForEventAtIndex != nil
}

func (di *TextViewDelegate) SetTextViewMenuForEventAtIndex(f func(view TextView, menu Menu, event Event, charIndex uint) IMenu) {
	di._TextViewMenuForEventAtIndex = f
}

func (di *TextViewDelegate) TextViewMenuForEventAtIndex(view TextView, menu Menu, event Event, charIndex uint) IMenu {
	return di._TextViewMenuForEventAtIndex(view, menu, event, charIndex)
}
func (di *TextViewDelegate) ImplementsTextViewClickedOnLink() bool {
	return di._TextViewClickedOnLink != nil
}

// deprecated
func (di *TextViewDelegate) SetTextViewClickedOnLink(f func(textView TextView, link objc.Object) bool) {
	di._TextViewClickedOnLink = f
}

// deprecated
func (di *TextViewDelegate) TextViewClickedOnLink(textView TextView, link objc.Object) bool {
	return di._TextViewClickedOnLink(textView, link)
}

type TextViewDelegateWrapper struct {
	TextDelegateWrapper
}

func (t_ TextViewDelegateWrapper) ImplementsUndoManagerForTextView() bool {
	return t_.RespondsToSelector(objc.GetSelector("undoManagerForTextView:"))
}

func (t_ TextViewDelegateWrapper) UndoManagerForTextView(view ITextView) foundation.UndoManager {
	rv := objc.CallMethod[foundation.UndoManager](t_, objc.GetSelector("undoManagerForTextView:"), objc.ExtractPtr(view))
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewWillDisplayToolTipForCharacterAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willDisplayToolTip:forCharacterAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextViewWillDisplayToolTipForCharacterAtIndex(textView ITextView, tooltip string, characterIndex uint) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("textView:willDisplayToolTip:forCharacterAtIndex:"), objc.ExtractPtr(textView), tooltip, characterIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewURLForContentsOfTextAttachmentAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:URLForContentsOfTextAttachment:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextViewURLForContentsOfTextAttachmentAtIndex(textView ITextView, textAttachment ITextAttachment, charIndex uint) foundation.URL {
	rv := objc.CallMethod[foundation.URL](t_, objc.GetSelector("textView:URLForContentsOfTextAttachment:atIndex:"), objc.ExtractPtr(textView), objc.ExtractPtr(textAttachment), charIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewWillChangeSelectionFromCharacterRangeToCharacterRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"))
}

func (t_ TextViewDelegateWrapper) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView ITextView, oldSelectedCharRange foundation.Range, newSelectedCharRange foundation.Range) foundation.Range {
	rv := objc.CallMethod[foundation.Range](t_, objc.GetSelector("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), objc.ExtractPtr(textView), oldSelectedCharRange, newSelectedCharRange)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewWillChangeSelectionFromCharacterRangesToCharacterRanges() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"))
}

func (t_ TextViewDelegateWrapper) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView ITextView, oldSelectedCharRanges []foundation.IValue, newSelectedCharRanges []foundation.IValue) []foundation.Value {
	rv := objc.CallMethod[[]foundation.Value](t_, objc.GetSelector("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), objc.ExtractPtr(textView), oldSelectedCharRanges, newSelectedCharRanges)
	// TODO: convert slice items...
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewDidChangeSelection() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeSelection:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeSelection(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textViewDidChangeSelection:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewCandidatesForSelectedRange_() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidates:forSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextViewCandidatesForSelectedRange_(textView ITextView, candidates []foundation.ITextCheckingResult, selectedRange foundation.Range) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textView:candidates:forSelectedRange:"), objc.ExtractPtr(textView), candidates, selectedRange)
	// TODO: convert slice items...
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewCandidatesForSelectedRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:candidatesForSelectedRange:"))
}

func (t_ TextViewDelegateWrapper) TextViewCandidatesForSelectedRange(textView ITextView, selectedRange foundation.Range) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("textView:candidatesForSelectedRange:"), objc.ExtractPtr(textView), selectedRange)
	// TODO: convert slice items...
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldSelectCandidateAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSelectCandidateAtIndex:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldSelectCandidateAtIndex(textView ITextView, index uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldSelectCandidateAtIndex:"), objc.ExtractPtr(textView), index)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldUpdateTouchBarItemIdentifiers() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldUpdateTouchBarItemIdentifiers:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldUpdateTouchBarItemIdentifiers(textView ITextView, identifiers []TouchBarItemIdentifier) []TouchBarItemIdentifier {
	rv := objc.CallMethod[[]TouchBarItemIdentifier](t_, objc.GetSelector("textView:shouldUpdateTouchBarItemIdentifiers:"), objc.ExtractPtr(textView), identifiers)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldChangeTextInRangeReplacementString() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRange:replacementString:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldChangeTextInRangeReplacementString(textView ITextView, affectedCharRange foundation.Range, replacementString string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldChangeTextInRange:replacementString:"), objc.ExtractPtr(textView), affectedCharRange, replacementString)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldChangeTextInRangesReplacementStrings() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTextInRanges:replacementStrings:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldChangeTextInRangesReplacementStrings(textView ITextView, affectedRanges []foundation.IValue, replacementStrings []string) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:shouldChangeTextInRanges:replacementStrings:"), objc.ExtractPtr(textView), affectedRanges, replacementStrings)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldChangeTypingAttributesToAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldChangeTypingAttributes:toAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldChangeTypingAttributesToAttributes(textView ITextView, oldTypingAttributes map[string]objc.IObject, newTypingAttributes map[foundation.AttributedStringKey]objc.IObject) map[foundation.AttributedStringKey]objc.Object {
	rv := objc.CallMethod[map[foundation.AttributedStringKey]objc.Object](t_, objc.GetSelector("textView:shouldChangeTypingAttributes:toAttributes:"), objc.ExtractPtr(textView), oldTypingAttributes, newTypingAttributes)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewDidChangeTypingAttributes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textViewDidChangeTypingAttributes:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidChangeTypingAttributes(notification foundation.INotification) {
	objc.CallMethod[objc.Void](t_, objc.GetSelector("textViewDidChangeTypingAttributes:"), objc.ExtractPtr(notification))
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewClickedOnLinkAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextViewClickedOnLinkAtIndex(textView ITextView, link objc.IObject, charIndex uint) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:clickedOnLink:atIndex:"), objc.ExtractPtr(textView), objc.ExtractPtr(link), charIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewShouldSetSpellingStateRange() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:shouldSetSpellingState:range:"))
}

func (t_ TextViewDelegateWrapper) TextViewShouldSetSpellingStateRange(textView ITextView, value int, affectedCharRange foundation.Range) int {
	rv := objc.CallMethod[int](t_, objc.GetSelector("textView:shouldSetSpellingState:range:"), objc.ExtractPtr(textView), value, affectedCharRange)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewWillCheckTextInRangeOptionsTypes() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willCheckTextInRange:options:types:"))
}

func (t_ TextViewDelegateWrapper) TextViewWillCheckTextInRangeOptionsTypes(view ITextView, range_ foundation.Range, options map[TextCheckingOptionKey]objc.IObject, checkingTypes *foundation.TextCheckingTypes) map[TextCheckingOptionKey]objc.Object {
	rv := objc.CallMethod[map[TextCheckingOptionKey]objc.Object](t_, objc.GetSelector("textView:willCheckTextInRange:options:types:"), objc.ExtractPtr(view), range_, options, checkingTypes)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"))
}

func (t_ TextViewDelegateWrapper) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view ITextView, range_ foundation.Range, checkingTypes foundation.TextCheckingTypes, options map[TextCheckingOptionKey]objc.IObject, results []foundation.ITextCheckingResult, orthography foundation.IOrthography, wordCount int) []foundation.TextCheckingResult {
	rv := objc.CallMethod[[]foundation.TextCheckingResult](t_, objc.GetSelector("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), objc.ExtractPtr(view), range_, checkingTypes, options, results, objc.ExtractPtr(orthography), wordCount)
	// TODO: convert slice items...
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

func (t_ TextViewDelegateWrapper) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView ITextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.CallMethod[[]string](t_, objc.GetSelector("textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.ExtractPtr(textView), words, charRange, index)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewWillShowSharingServicePickerForItems() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:willShowSharingServicePicker:forItems:"))
}

func (t_ TextViewDelegateWrapper) TextViewWillShowSharingServicePickerForItems(textView ITextView, servicePicker ISharingServicePicker, items []objc.IObject) SharingServicePicker {
	rv := objc.CallMethod[SharingServicePicker](t_, objc.GetSelector("textView:willShowSharingServicePicker:forItems:"), objc.ExtractPtr(textView), objc.ExtractPtr(servicePicker), items)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewDoCommandBySelector() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:doCommandBySelector:"))
}

func (t_ TextViewDelegateWrapper) TextViewDoCommandBySelector(textView ITextView, commandSelector objc.Selector) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("textView:doCommandBySelector:"), objc.ExtractPtr(textView), commandSelector)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewMenuForEventAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:menu:forEvent:atIndex:"))
}

func (t_ TextViewDelegateWrapper) TextViewMenuForEventAtIndex(view ITextView, menu IMenu, event IEvent, charIndex uint) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("textView:menu:forEvent:atIndex:"), objc.ExtractPtr(view), objc.ExtractPtr(menu), objc.ExtractPtr(event), charIndex)
	return rv
}

func (t_ TextViewDelegateWrapper) ImplementsTextViewClickedOnLink() bool {
	return t_.RespondsToSelector(objc.GetSelector("textView:clickedOnLink:"))
}
