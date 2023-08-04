// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ITokenFieldCellDelegate interface {
	ImplementsTokenFieldCellDisplayStringForRepresentedObject() bool
	// optional
	TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	ImplementsTokenFieldCellStyleForRepresentedObject() bool
	// optional
	TokenFieldCellStyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle
	ImplementsTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool
	// optional
	TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	ImplementsTokenFieldCellEditingStringForRepresentedObject() bool
	// optional
	TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	ImplementsTokenFieldCellRepresentedObjectForEditingString() bool
	// optional
	TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject
	ImplementsTokenFieldCellShouldAddObjectsAtIndex() bool
	// optional
	TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject
	ImplementsTokenFieldCellReadFromPasteboard() bool
	// optional
	TokenFieldCellReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject
	ImplementsTokenFieldCellWriteRepresentedObjectsToPasteboard() bool
	// optional
	TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool
	ImplementsTokenFieldCellHasMenuForRepresentedObject() bool
	// optional
	TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool
	ImplementsTokenFieldCellMenuForRepresentedObject() bool
	// optional
	TokenFieldCellMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu
}

type TokenFieldCellDelegate struct {
	_TokenFieldCellDisplayStringForRepresentedObject                      func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCellStyleForRepresentedObject                              func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle
	_TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	_TokenFieldCellEditingStringForRepresentedObject                      func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string
	_TokenFieldCellRepresentedObjectForEditingString                      func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject
	_TokenFieldCellShouldAddObjectsAtIndex                                func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject
	_TokenFieldCellReadFromPasteboard                                     func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject
	_TokenFieldCellWriteRepresentedObjectsToPasteboard                    func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool
	_TokenFieldCellHasMenuForRepresentedObject                            func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool
	_TokenFieldCellMenuForRepresentedObject                               func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu
}

func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellDisplayStringForRepresentedObject() bool {
	return di._TokenFieldCellDisplayStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellDisplayStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCellDisplayStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellStyleForRepresentedObject() bool {
	return di._TokenFieldCellStyleForRepresentedObject != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellStyleForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle) {
	di._TokenFieldCellStyleForRepresentedObject = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellStyleForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) TokenStyle {
	return di._TokenFieldCellStyleForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return di._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(f func(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	di._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell TokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	return di._TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell, substring, tokenIndex, selectedIndex)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellEditingStringForRepresentedObject() bool {
	return di._TokenFieldCellEditingStringForRepresentedObject != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellEditingStringForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) string) {
	di._TokenFieldCellEditingStringForRepresentedObject = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) string {
	return di._TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellRepresentedObjectForEditingString() bool {
	return di._TokenFieldCellRepresentedObjectForEditingString != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellRepresentedObjectForEditingString(f func(tokenFieldCell TokenFieldCell, editingString string) objc.IObject) {
	di._TokenFieldCellRepresentedObjectForEditingString = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell TokenFieldCell, editingString string) objc.IObject {
	return di._TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell, editingString)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellShouldAddObjectsAtIndex() bool {
	return di._TokenFieldCellShouldAddObjectsAtIndex != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellShouldAddObjectsAtIndex(f func(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject) {
	di._TokenFieldCellShouldAddObjectsAtIndex = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell TokenFieldCell, tokens []objc.Object, index uint) []objc.IObject {
	return di._TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell, tokens, index)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellReadFromPasteboard() bool {
	return di._TokenFieldCellReadFromPasteboard != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellReadFromPasteboard(f func(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject) {
	di._TokenFieldCellReadFromPasteboard = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellReadFromPasteboard(tokenFieldCell TokenFieldCell, pboard Pasteboard) []objc.IObject {
	return di._TokenFieldCellReadFromPasteboard(tokenFieldCell, pboard)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellWriteRepresentedObjectsToPasteboard() bool {
	return di._TokenFieldCellWriteRepresentedObjectsToPasteboard != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellWriteRepresentedObjectsToPasteboard(f func(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool) {
	di._TokenFieldCellWriteRepresentedObjectsToPasteboard = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell TokenFieldCell, objects []objc.Object, pboard Pasteboard) bool {
	return di._TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell, objects, pboard)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellHasMenuForRepresentedObject() bool {
	return di._TokenFieldCellHasMenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellHasMenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool) {
	di._TokenFieldCellHasMenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) bool {
	return di._TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell, representedObject)
}
func (di *TokenFieldCellDelegate) ImplementsTokenFieldCellMenuForRepresentedObject() bool {
	return di._TokenFieldCellMenuForRepresentedObject != nil
}

func (di *TokenFieldCellDelegate) SetTokenFieldCellMenuForRepresentedObject(f func(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu) {
	di._TokenFieldCellMenuForRepresentedObject = f
}

func (di *TokenFieldCellDelegate) TokenFieldCellMenuForRepresentedObject(tokenFieldCell TokenFieldCell, representedObject objc.Object) IMenu {
	return di._TokenFieldCellMenuForRepresentedObject(tokenFieldCell, representedObject)
}

type TokenFieldCellDelegateWrapper struct {
	objc.Object
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellDisplayStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellDisplayStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellStyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellStyleForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenFieldCell:styleForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenFieldCell ITokenFieldCell, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenFieldCell), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellEditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellEditingStringForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenFieldCell:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellRepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellRepresentedObjectForEditingString(tokenFieldCell ITokenFieldCell, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenFieldCell:representedObjectForEditingString:"), objc.ExtractPtr(tokenFieldCell), editingString)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellShouldAddObjectsAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellShouldAddObjectsAtIndex(tokenFieldCell ITokenFieldCell, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenFieldCell), tokens, index)
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:readFromPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellReadFromPasteboard(tokenFieldCell ITokenFieldCell, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenFieldCell:readFromPasteboard:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellWriteRepresentedObjectsToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellWriteRepresentedObjectsToPasteboard(tokenFieldCell ITokenFieldCell, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenFieldCell), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellHasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellHasMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenFieldCell:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldCellDelegateWrapper) ImplementsTokenFieldCellMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"))
}

func (t_ TokenFieldCellDelegateWrapper) TokenFieldCellMenuForRepresentedObject(tokenFieldCell ITokenFieldCell, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenFieldCell:menuForRepresentedObject:"), objc.ExtractPtr(tokenFieldCell), objc.ExtractPtr(representedObject))
	return rv
}
