// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// Methods that responder subclasses implement to support key binding commands, such as inserting tabs and newlines, or moving the insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding?language=objc
type PStandardKeyBindingResponding interface {
	// optional
	InsertParagraphSeparator(sender objc.Object)
	HasInsertParagraphSeparator() bool

	// optional
	MoveLeftAndModifySelection(sender objc.Object)
	HasMoveLeftAndModifySelection() bool

	// optional
	InsertContainerBreak(sender objc.Object)
	HasInsertContainerBreak() bool

	// optional
	CapitalizeWord(sender objc.Object)
	HasCapitalizeWord() bool

	// optional
	InsertLineBreak(sender objc.Object)
	HasInsertLineBreak() bool

	// optional
	MoveToRightEndOfLine(sender objc.Object)
	HasMoveToRightEndOfLine() bool

	// optional
	MoveRight(sender objc.Object)
	HasMoveRight() bool

	// optional
	InsertText(insertString objc.Object)
	HasInsertText() bool

	// optional
	InsertSingleQuoteIgnoringSubstitution(sender objc.Object)
	HasInsertSingleQuoteIgnoringSubstitution() bool

	// optional
	MoveWordBackward(sender objc.Object)
	HasMoveWordBackward() bool

	// optional
	MoveToEndOfDocument(sender objc.Object)
	HasMoveToEndOfDocument() bool

	// optional
	SelectWord(sender objc.Object)
	HasSelectWord() bool

	// optional
	MoveToEndOfParagraph(sender objc.Object)
	HasMoveToEndOfParagraph() bool

	// optional
	ScrollToEndOfDocument(sender objc.Object)
	HasScrollToEndOfDocument() bool

	// optional
	DeleteForward(sender objc.Object)
	HasDeleteForward() bool

	// optional
	MoveForwardAndModifySelection(sender objc.Object)
	HasMoveForwardAndModifySelection() bool

	// optional
	DeleteToEndOfLine(sender objc.Object)
	HasDeleteToEndOfLine() bool

	// optional
	MoveToEndOfLine(sender objc.Object)
	HasMoveToEndOfLine() bool

	// optional
	QuickLookPreviewItems(sender objc.Object)
	HasQuickLookPreviewItems() bool

	// optional
	MakeTextWritingDirectionRightToLeft(sender objc.Object)
	HasMakeTextWritingDirectionRightToLeft() bool

	// optional
	SelectParagraph(sender objc.Object)
	HasSelectParagraph() bool

	// optional
	MoveToBeginningOfDocumentAndModifySelection(sender objc.Object)
	HasMoveToBeginningOfDocumentAndModifySelection() bool

	// optional
	ScrollLineDown(sender objc.Object)
	HasScrollLineDown() bool

	// optional
	DeleteToBeginningOfLine(sender objc.Object)
	HasDeleteToBeginningOfLine() bool

	// optional
	MoveUpAndModifySelection(sender objc.Object)
	HasMoveUpAndModifySelection() bool

	// optional
	DeleteWordForward(sender objc.Object)
	HasDeleteWordForward() bool

	// optional
	DeleteToMark(sender objc.Object)
	HasDeleteToMark() bool

	// optional
	MoveWordForward(sender objc.Object)
	HasMoveWordForward() bool

	// optional
	InsertNewline(sender objc.Object)
	HasInsertNewline() bool

	// optional
	DeleteBackwardByDecomposingPreviousCharacter(sender objc.Object)
	HasDeleteBackwardByDecomposingPreviousCharacter() bool

	// optional
	MoveDownAndModifySelection(sender objc.Object)
	HasMoveDownAndModifySelection() bool

	// optional
	MoveBackwardAndModifySelection(sender objc.Object)
	HasMoveBackwardAndModifySelection() bool

	// optional
	MoveWordLeftAndModifySelection(sender objc.Object)
	HasMoveWordLeftAndModifySelection() bool

	// optional
	MoveUp(sender objc.Object)
	HasMoveUp() bool

	// optional
	MoveToBeginningOfLine(sender objc.Object)
	HasMoveToBeginningOfLine() bool

	// optional
	PageDown(sender objc.Object)
	HasPageDown() bool

	// optional
	InsertTabIgnoringFieldEditor(sender objc.Object)
	HasInsertTabIgnoringFieldEditor() bool

	// optional
	SelectAll(sender objc.Object)
	HasSelectAll() bool

	// optional
	Transpose(sender objc.Object)
	HasTranspose() bool

	// optional
	DeleteToEndOfParagraph(sender objc.Object)
	HasDeleteToEndOfParagraph() bool

	// optional
	MoveBackward(sender objc.Object)
	HasMoveBackward() bool

	// optional
	MakeBaseWritingDirectionRightToLeft(sender objc.Object)
	HasMakeBaseWritingDirectionRightToLeft() bool

	// optional
	ScrollToBeginningOfDocument(sender objc.Object)
	HasScrollToBeginningOfDocument() bool

	// optional
	MoveWordLeft(sender objc.Object)
	HasMoveWordLeft() bool

	// optional
	SelectToMark(sender objc.Object)
	HasSelectToMark() bool

	// optional
	DeleteToBeginningOfParagraph(sender objc.Object)
	HasDeleteToBeginningOfParagraph() bool

	// optional
	MoveLeft(sender objc.Object)
	HasMoveLeft() bool

	// optional
	SelectLine(sender objc.Object)
	HasSelectLine() bool

	// optional
	MakeTextWritingDirectionNatural(sender objc.Object)
	HasMakeTextWritingDirectionNatural() bool

	// optional
	MoveToBeginningOfParagraph(sender objc.Object)
	HasMoveToBeginningOfParagraph() bool

	// optional
	MoveToBeginningOfParagraphAndModifySelection(sender objc.Object)
	HasMoveToBeginningOfParagraphAndModifySelection() bool

	// optional
	TransposeWords(sender objc.Object)
	HasTransposeWords() bool

	// optional
	MakeBaseWritingDirectionLeftToRight(sender objc.Object)
	HasMakeBaseWritingDirectionLeftToRight() bool

	// optional
	Complete(sender objc.Object)
	HasComplete() bool

	// optional
	MoveParagraphBackwardAndModifySelection(sender objc.Object)
	HasMoveParagraphBackwardAndModifySelection() bool

	// optional
	PageUp(sender objc.Object)
	HasPageUp() bool

	// optional
	MoveWordRightAndModifySelection(sender objc.Object)
	HasMoveWordRightAndModifySelection() bool

	// optional
	InsertDoubleQuoteIgnoringSubstitution(sender objc.Object)
	HasInsertDoubleQuoteIgnoringSubstitution() bool

	// optional
	DeleteWordBackward(sender objc.Object)
	HasDeleteWordBackward() bool

	// optional
	UppercaseWord(sender objc.Object)
	HasUppercaseWord() bool

	// optional
	CenterSelectionInVisibleArea(sender objc.Object)
	HasCenterSelectionInVisibleArea() bool

	// optional
	Yank(sender objc.Object)
	HasYank() bool

	// optional
	MoveToRightEndOfLineAndModifySelection(sender objc.Object)
	HasMoveToRightEndOfLineAndModifySelection() bool

	// optional
	CancelOperation(sender objc.Object)
	HasCancelOperation() bool

	// optional
	MoveRightAndModifySelection(sender objc.Object)
	HasMoveRightAndModifySelection() bool

	// optional
	InsertNewlineIgnoringFieldEditor(sender objc.Object)
	HasInsertNewlineIgnoringFieldEditor() bool

	// optional
	MoveWordRight(sender objc.Object)
	HasMoveWordRight() bool

	// optional
	MoveForward(sender objc.Object)
	HasMoveForward() bool

	// optional
	ChangeCaseOfLetter(sender objc.Object)
	HasChangeCaseOfLetter() bool

	// optional
	ScrollPageUp(sender objc.Object)
	HasScrollPageUp() bool

	// optional
	InsertBacktab(sender objc.Object)
	HasInsertBacktab() bool

	// optional
	InsertTab(sender objc.Object)
	HasInsertTab() bool

	// optional
	MoveDown(sender objc.Object)
	HasMoveDown() bool

	// optional
	PageDownAndModifySelection(sender objc.Object)
	HasPageDownAndModifySelection() bool

	// optional
	MoveToLeftEndOfLineAndModifySelection(sender objc.Object)
	HasMoveToLeftEndOfLineAndModifySelection() bool

	// optional
	ScrollLineUp(sender objc.Object)
	HasScrollLineUp() bool

	// optional
	DoCommandBySelector(selector objc.Selector)
	HasDoCommandBySelector() bool

	// optional
	SwapWithMark(sender objc.Object)
	HasSwapWithMark() bool

	// optional
	MoveToLeftEndOfLine(sender objc.Object)
	HasMoveToLeftEndOfLine() bool

	// optional
	DeleteBackward(sender objc.Object)
	HasDeleteBackward() bool

	// optional
	MakeTextWritingDirectionLeftToRight(sender objc.Object)
	HasMakeTextWritingDirectionLeftToRight() bool

	// optional
	MoveWordBackwardAndModifySelection(sender objc.Object)
	HasMoveWordBackwardAndModifySelection() bool

	// optional
	ScrollPageDown(sender objc.Object)
	HasScrollPageDown() bool

	// optional
	SetMark(sender objc.Object)
	HasSetMark() bool

	// optional
	MoveWordForwardAndModifySelection(sender objc.Object)
	HasMoveWordForwardAndModifySelection() bool

	// optional
	PageUpAndModifySelection(sender objc.Object)
	HasPageUpAndModifySelection() bool

	// optional
	MoveToBeginningOfDocument(sender objc.Object)
	HasMoveToBeginningOfDocument() bool

	// optional
	MoveToBeginningOfLineAndModifySelection(sender objc.Object)
	HasMoveToBeginningOfLineAndModifySelection() bool

	// optional
	LowercaseWord(sender objc.Object)
	HasLowercaseWord() bool

	// optional
	MoveToEndOfDocumentAndModifySelection(sender objc.Object)
	HasMoveToEndOfDocumentAndModifySelection() bool

	// optional
	MoveToEndOfLineAndModifySelection(sender objc.Object)
	HasMoveToEndOfLineAndModifySelection() bool

	// optional
	Indent(sender objc.Object)
	HasIndent() bool

	// optional
	MakeBaseWritingDirectionNatural(sender objc.Object)
	HasMakeBaseWritingDirectionNatural() bool

	// optional
	MoveParagraphForwardAndModifySelection(sender objc.Object)
	HasMoveParagraphForwardAndModifySelection() bool

	// optional
	MoveToEndOfParagraphAndModifySelection(sender objc.Object)
	HasMoveToEndOfParagraphAndModifySelection() bool
}

// A concrete type wrapper for the [PStandardKeyBindingResponding] protocol.
type StandardKeyBindingRespondingWrapper struct {
	objc.Object
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertParagraphSeparator() bool {
	return s_.RespondsToSelector(objc.Sel("insertParagraphSeparator:"))
}

// Inserts a paragraph separator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005219-insertparagraphseparator?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertParagraphSeparator(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertParagraphSeparator:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveLeftAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveLeftAndModifySelection:"))
}

// Extends the selection to include the content to the left of the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005238-moveleftandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveLeftAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveLeftAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertContainerBreak() bool {
	return s_.RespondsToSelector(objc.Sel("insertContainerBreak:"))
}

// Inserts a container break, such as a new page break. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005214-insertcontainerbreak?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertContainerBreak(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertContainerBreak:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasCapitalizeWord() bool {
	return s_.RespondsToSelector(objc.Sel("capitalizeWord:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005197-capitalizeword?language=objc
func (s_ StandardKeyBindingRespondingWrapper) CapitalizeWord(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("capitalizeWord:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertLineBreak() bool {
	return s_.RespondsToSelector(objc.Sel("insertLineBreak:"))
}

// Inserts a line break character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005216-insertlinebreak?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertLineBreak(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertLineBreak:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToRightEndOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("moveToRightEndOfLine:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005257-movetorightendofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToRightEndOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToRightEndOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveRight() bool {
	return s_.RespondsToSelector(objc.Sel("moveRight:"))
}

// Moves the insertion pointer right in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005241-moveright?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveRight(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveRight:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertText() bool {
	return s_.RespondsToSelector(objc.Sel("insertText:"))
}

// Inserts the text you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005223-inserttext?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertText(insertString objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertText:"), insertString)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertSingleQuoteIgnoringSubstitution() bool {
	return s_.RespondsToSelector(objc.Sel("insertSingleQuoteIgnoringSubstitution:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005220-insertsinglequoteignoringsubstit?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertSingleQuoteIgnoringSubstitution(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertSingleQuoteIgnoringSubstitution:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordBackward() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordBackward:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005261-movewordbackward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordBackward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordBackward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfDocument() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfDocument:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005249-movetoendofdocument?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfDocument(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfDocument:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSelectWord() bool {
	return s_.RespondsToSelector(objc.Sel("selectWord:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005284-selectword?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SelectWord(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("selectWord:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfParagraph() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfParagraph:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005253-movetoendofparagraph?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfParagraph(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfParagraph:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollToEndOfDocument() bool {
	return s_.RespondsToSelector(objc.Sel("scrollToEndOfDocument:"))
}

// Scrolls the content to the end of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005279-scrolltoendofdocument?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollToEndOfDocument(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollToEndOfDocument:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteForward() bool {
	return s_.RespondsToSelector(objc.Sel("deleteForward:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005203-deleteforward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteForward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteForward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveForwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveForwardAndModifySelection:"))
}

// Extends the selection to include the content after the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005236-moveforwardandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveForwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveForwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteToEndOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("deleteToEndOfLine:"))
}

// Deletes content from the insertion point to the end of the current line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005206-deletetoendofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteToEndOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteToEndOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfLine:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005251-movetoendofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasQuickLookPreviewItems() bool {
	return s_.RespondsToSelector(objc.Sel("quickLookPreviewItems:"))
}

// Invokes QuickLook to preview the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005273-quicklookpreviewitems?language=objc
func (s_ StandardKeyBindingRespondingWrapper) QuickLookPreviewItems(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("quickLookPreviewItems:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeTextWritingDirectionRightToLeft() bool {
	return s_.RespondsToSelector(objc.Sel("makeTextWritingDirectionRightToLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005230-maketextwritingdirectionrighttol?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeTextWritingDirectionRightToLeft(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeTextWritingDirectionRightToLeft:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSelectParagraph() bool {
	return s_.RespondsToSelector(objc.Sel("selectParagraph:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005282-selectparagraph?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SelectParagraph(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("selectParagraph:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfDocumentAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfDocumentAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005244-movetobeginningofdocumentandmodi?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfDocumentAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfDocumentAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollLineDown() bool {
	return s_.RespondsToSelector(objc.Sel("scrollLineDown:"))
}

// Scrolls the content down by a line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005274-scrolllinedown?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollLineDown(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollLineDown:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteToBeginningOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("deleteToBeginningOfLine:"))
}

// Deletes content from the insertion point to the beginning of the current line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005204-deletetobeginningofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteToBeginningOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteToBeginningOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveUpAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveUpAndModifySelection:"))
}

// Extends the selection to include the content above the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005260-moveupandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveUpAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveUpAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteWordForward() bool {
	return s_.RespondsToSelector(objc.Sel("deleteWordForward:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005210-deletewordforward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteWordForward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteWordForward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteToMark() bool {
	return s_.RespondsToSelector(objc.Sel("deleteToMark:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005208-deletetomark?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteToMark(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteToMark:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordForward() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordForward:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005263-movewordforward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordForward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordForward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertNewline() bool {
	return s_.RespondsToSelector(objc.Sel("insertNewline:"))
}

// Inserts a newline character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005217-insertnewline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertNewline(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertNewline:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteBackwardByDecomposingPreviousCharacter() bool {
	return s_.RespondsToSelector(objc.Sel("deleteBackwardByDecomposingPreviousCharacter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005202-deletebackwardbydecomposingprevi?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteBackwardByDecomposingPreviousCharacter(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteBackwardByDecomposingPreviousCharacter:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveDownAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveDownAndModifySelection:"))
}

// Extends the selection to include the content below the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005234-movedownandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveDownAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveDownAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveBackwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveBackwardAndModifySelection:"))
}

// Extends the selection to include the content before the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005232-movebackwardandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveBackwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveBackwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordLeftAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordLeftAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005266-movewordleftandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordLeftAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordLeftAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveUp() bool {
	return s_.RespondsToSelector(objc.Sel("moveUp:"))
}

// Moves the insertion pointer up in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005259-moveup?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveUp(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveUp:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfLine:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005245-movetobeginningofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasPageDown() bool {
	return s_.RespondsToSelector(objc.Sel("pageDown:"))
}

// Moves the visible content region down by a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005269-pagedown?language=objc
func (s_ StandardKeyBindingRespondingWrapper) PageDown(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("pageDown:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertTabIgnoringFieldEditor() bool {
	return s_.RespondsToSelector(objc.Sel("insertTabIgnoringFieldEditor:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005222-inserttabignoringfieldeditor?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertTabIgnoringFieldEditor(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertTabIgnoringFieldEditor:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSelectAll() bool {
	return s_.RespondsToSelector(objc.Sel("selectAll:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005280-selectall?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SelectAll(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("selectAll:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasTranspose() bool {
	return s_.RespondsToSelector(objc.Sel("transpose:"))
}

// Transposes the content around the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005287-transpose?language=objc
func (s_ StandardKeyBindingRespondingWrapper) Transpose(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("transpose:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteToEndOfParagraph() bool {
	return s_.RespondsToSelector(objc.Sel("deleteToEndOfParagraph:"))
}

// Deletes content from the insertion point to the end of the current paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005207-deletetoendofparagraph?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteToEndOfParagraph(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteToEndOfParagraph:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveBackward() bool {
	return s_.RespondsToSelector(objc.Sel("moveBackward:"))
}

// Moves the insertion pointer backward in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005231-movebackward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveBackward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveBackward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeBaseWritingDirectionRightToLeft() bool {
	return s_.RespondsToSelector(objc.Sel("makeBaseWritingDirectionRightToLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005227-makebasewritingdirectionrighttol?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeBaseWritingDirectionRightToLeft(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeBaseWritingDirectionRightToLeft:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollToBeginningOfDocument() bool {
	return s_.RespondsToSelector(objc.Sel("scrollToBeginningOfDocument:"))
}

// Scrolls the content to the beginning of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005278-scrolltobeginningofdocument?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollToBeginningOfDocument(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollToBeginningOfDocument:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordLeft() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005265-movewordleft?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordLeft(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordLeft:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSelectToMark() bool {
	return s_.RespondsToSelector(objc.Sel("selectToMark:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005283-selecttomark?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SelectToMark(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("selectToMark:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteToBeginningOfParagraph() bool {
	return s_.RespondsToSelector(objc.Sel("deleteToBeginningOfParagraph:"))
}

// Deletes content from the insertion point to the beginning of the current paragraph. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005205-deletetobeginningofparagraph?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteToBeginningOfParagraph(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteToBeginningOfParagraph:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveLeft() bool {
	return s_.RespondsToSelector(objc.Sel("moveLeft:"))
}

// Moves the insertion pointer left in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005237-moveleft?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveLeft(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveLeft:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSelectLine() bool {
	return s_.RespondsToSelector(objc.Sel("selectLine:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005281-selectline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SelectLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("selectLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeTextWritingDirectionNatural() bool {
	return s_.RespondsToSelector(objc.Sel("makeTextWritingDirectionNatural:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005229-maketextwritingdirectionnatural?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeTextWritingDirectionNatural(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeTextWritingDirectionNatural:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfParagraph() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfParagraph:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005247-movetobeginningofparagraph?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfParagraph(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfParagraph:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfParagraphAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfParagraphAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005248-movetobeginningofparagraphandmod?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfParagraphAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfParagraphAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasTransposeWords() bool {
	return s_.RespondsToSelector(objc.Sel("transposeWords:"))
}

// Transposes the words around the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005288-transposewords?language=objc
func (s_ StandardKeyBindingRespondingWrapper) TransposeWords(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("transposeWords:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeBaseWritingDirectionLeftToRight() bool {
	return s_.RespondsToSelector(objc.Sel("makeBaseWritingDirectionLeftToRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005225-makebasewritingdirectionlefttori?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeBaseWritingDirectionLeftToRight(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeBaseWritingDirectionLeftToRight:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasComplete() bool {
	return s_.RespondsToSelector(objc.Sel("complete:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005200-complete?language=objc
func (s_ StandardKeyBindingRespondingWrapper) Complete(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("complete:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveParagraphBackwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveParagraphBackwardAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005239-moveparagraphbackwardandmodifyse?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveParagraphBackwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveParagraphBackwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasPageUp() bool {
	return s_.RespondsToSelector(objc.Sel("pageUp:"))
}

// Moves the visible content region up by a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005271-pageup?language=objc
func (s_ StandardKeyBindingRespondingWrapper) PageUp(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("pageUp:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordRightAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordRightAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005268-movewordrightandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordRightAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordRightAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertDoubleQuoteIgnoringSubstitution() bool {
	return s_.RespondsToSelector(objc.Sel("insertDoubleQuoteIgnoringSubstitution:"))
}

// Inserts a double quotation mark without substituting a curly quotation mark. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005215-insertdoublequoteignoringsubstit?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertDoubleQuoteIgnoringSubstitution(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertDoubleQuoteIgnoringSubstitution:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteWordBackward() bool {
	return s_.RespondsToSelector(objc.Sel("deleteWordBackward:"))
}

// Deletes the word preceding the current insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005209-deletewordbackward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteWordBackward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteWordBackward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasUppercaseWord() bool {
	return s_.RespondsToSelector(objc.Sel("uppercaseWord:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005289-uppercaseword?language=objc
func (s_ StandardKeyBindingRespondingWrapper) UppercaseWord(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("uppercaseWord:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasCenterSelectionInVisibleArea() bool {
	return s_.RespondsToSelector(objc.Sel("centerSelectionInVisibleArea:"))
}

// Moves the visible content region so the current selection is visually centered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005198-centerselectioninvisiblearea?language=objc
func (s_ StandardKeyBindingRespondingWrapper) CenterSelectionInVisibleArea(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("centerSelectionInVisibleArea:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasYank() bool {
	return s_.RespondsToSelector(objc.Sel("yank:"))
}

// Deletes the current selection, placing it in a temporary buffer, such as the Clipboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005290-yank?language=objc
func (s_ StandardKeyBindingRespondingWrapper) Yank(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("yank:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToRightEndOfLineAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToRightEndOfLineAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005258-movetorightendoflineandmodifysel?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToRightEndOfLineAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToRightEndOfLineAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasCancelOperation() bool {
	return s_.RespondsToSelector(objc.Sel("cancelOperation:"))
}

// Cancels the current operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005196-canceloperation?language=objc
func (s_ StandardKeyBindingRespondingWrapper) CancelOperation(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("cancelOperation:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveRightAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveRightAndModifySelection:"))
}

// Extends the selection to include the content to the right of the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005242-moverightandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveRightAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveRightAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertNewlineIgnoringFieldEditor() bool {
	return s_.RespondsToSelector(objc.Sel("insertNewlineIgnoringFieldEditor:"))
}

// Inserts a newline character without invoking the field editor’s normal handling to end editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005218-insertnewlineignoringfieldeditor?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertNewlineIgnoringFieldEditor(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertNewlineIgnoringFieldEditor:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordRight() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005267-movewordright?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordRight(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordRight:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveForward() bool {
	return s_.RespondsToSelector(objc.Sel("moveForward:"))
}

// Moves the insertion pointer forward in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005235-moveforward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveForward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveForward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasChangeCaseOfLetter() bool {
	return s_.RespondsToSelector(objc.Sel("changeCaseOfLetter:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005199-changecaseofletter?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ChangeCaseOfLetter(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("changeCaseOfLetter:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollPageUp() bool {
	return s_.RespondsToSelector(objc.Sel("scrollPageUp:"))
}

// Scrolls the content up by a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005277-scrollpageup?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollPageUp(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollPageUp:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertBacktab() bool {
	return s_.RespondsToSelector(objc.Sel("insertBacktab:"))
}

// Inserts a backtab character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005213-insertbacktab?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertBacktab(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertBacktab:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasInsertTab() bool {
	return s_.RespondsToSelector(objc.Sel("insertTab:"))
}

// Inserts a tab character. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005221-inserttab?language=objc
func (s_ StandardKeyBindingRespondingWrapper) InsertTab(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("insertTab:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveDown() bool {
	return s_.RespondsToSelector(objc.Sel("moveDown:"))
}

// Moves the insertion pointer down in the current content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005233-movedown?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveDown(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveDown:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasPageDownAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("pageDownAndModifySelection:"))
}

// Moves the visible content region down by a page, and extends the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005270-pagedownandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) PageDownAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("pageDownAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToLeftEndOfLineAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToLeftEndOfLineAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005256-movetoleftendoflineandmodifysele?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToLeftEndOfLineAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToLeftEndOfLineAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollLineUp() bool {
	return s_.RespondsToSelector(objc.Sel("scrollLineUp:"))
}

// Scrolls the content up by a line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005275-scrolllineup?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollLineUp(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollLineUp:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDoCommandBySelector() bool {
	return s_.RespondsToSelector(objc.Sel("doCommandBySelector:"))
}

// Performs the given selector if possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005211-docommandbyselector?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DoCommandBySelector(selector objc.Selector) {
	objc.Call[objc.Void](s_, objc.Sel("doCommandBySelector:"), selector)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSwapWithMark() bool {
	return s_.RespondsToSelector(objc.Sel("swapWithMark:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005286-swapwithmark?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SwapWithMark(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("swapWithMark:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToLeftEndOfLine() bool {
	return s_.RespondsToSelector(objc.Sel("moveToLeftEndOfLine:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005255-movetoleftendofline?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToLeftEndOfLine(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToLeftEndOfLine:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasDeleteBackward() bool {
	return s_.RespondsToSelector(objc.Sel("deleteBackward:"))
}

// Deletes content moving backward from the current insertion point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005201-deletebackward?language=objc
func (s_ StandardKeyBindingRespondingWrapper) DeleteBackward(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("deleteBackward:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeTextWritingDirectionLeftToRight() bool {
	return s_.RespondsToSelector(objc.Sel("makeTextWritingDirectionLeftToRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005228-maketextwritingdirectionlefttori?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeTextWritingDirectionLeftToRight(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeTextWritingDirectionLeftToRight:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordBackwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordBackwardAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005262-movewordbackwardandmodifyselecti?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordBackwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordBackwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasScrollPageDown() bool {
	return s_.RespondsToSelector(objc.Sel("scrollPageDown:"))
}

// Scrolls the content down by a page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005276-scrollpagedown?language=objc
func (s_ StandardKeyBindingRespondingWrapper) ScrollPageDown(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("scrollPageDown:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasSetMark() bool {
	return s_.RespondsToSelector(objc.Sel("setMark:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005285-setmark?language=objc
func (s_ StandardKeyBindingRespondingWrapper) SetMark(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setMark:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveWordForwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveWordForwardAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005264-movewordforwardandmodifyselectio?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveWordForwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveWordForwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasPageUpAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("pageUpAndModifySelection:"))
}

// Moves the visible content region up by a page, and extends the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005272-pageupandmodifyselection?language=objc
func (s_ StandardKeyBindingRespondingWrapper) PageUpAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("pageUpAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfDocument() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfDocument:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005243-movetobeginningofdocument?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfDocument(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfDocument:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToBeginningOfLineAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToBeginningOfLineAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005246-movetobeginningoflineandmodifyse?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToBeginningOfLineAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToBeginningOfLineAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasLowercaseWord() bool {
	return s_.RespondsToSelector(objc.Sel("lowercaseWord:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005224-lowercaseword?language=objc
func (s_ StandardKeyBindingRespondingWrapper) LowercaseWord(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("lowercaseWord:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfDocumentAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfDocumentAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005250-movetoendofdocumentandmodifysele?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfDocumentAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfDocumentAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfLineAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfLineAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005252-movetoendoflineandmodifyselectio?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfLineAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfLineAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasIndent() bool {
	return s_.RespondsToSelector(objc.Sel("indent:"))
}

// Indents the content at the current selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005212-indent?language=objc
func (s_ StandardKeyBindingRespondingWrapper) Indent(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("indent:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMakeBaseWritingDirectionNatural() bool {
	return s_.RespondsToSelector(objc.Sel("makeBaseWritingDirectionNatural:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005226-makebasewritingdirectionnatural?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MakeBaseWritingDirectionNatural(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("makeBaseWritingDirectionNatural:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveParagraphForwardAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveParagraphForwardAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005240-moveparagraphforwardandmodifysel?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveParagraphForwardAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveParagraphForwardAndModifySelection:"), sender)
}

func (s_ StandardKeyBindingRespondingWrapper) HasMoveToEndOfParagraphAndModifySelection() bool {
	return s_.RespondsToSelector(objc.Sel("moveToEndOfParagraphAndModifySelection:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstandardkeybindingresponding/3005254-movetoendofparagraphandmodifysel?language=objc
func (s_ StandardKeyBindingRespondingWrapper) MoveToEndOfParagraphAndModifySelection(sender objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("moveToEndOfParagraphAndModifySelection:"), sender)
}
