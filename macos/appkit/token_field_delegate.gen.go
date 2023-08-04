// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ITokenFieldDelegate interface {
	ITextFieldDelegate
	ImplementsTokenFieldDisplayStringForRepresentedObject() bool
	// optional
	TokenFieldDisplayStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string
	ImplementsTokenFieldStyleForRepresentedObject() bool
	// optional
	TokenFieldStyleForRepresentedObject(tokenField TokenField, representedObject objc.Object) TokenStyle
	ImplementsTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool
	// optional
	TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	ImplementsTokenFieldEditingStringForRepresentedObject() bool
	// optional
	TokenFieldEditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string
	ImplementsTokenFieldRepresentedObjectForEditingString() bool
	// optional
	TokenFieldRepresentedObjectForEditingString(tokenField TokenField, editingString string) objc.IObject
	ImplementsTokenFieldShouldAddObjectsAtIndex() bool
	// optional
	TokenFieldShouldAddObjectsAtIndex(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject
	ImplementsTokenFieldReadFromPasteboard() bool
	// optional
	TokenFieldReadFromPasteboard(tokenField TokenField, pboard Pasteboard) []objc.IObject
	ImplementsTokenFieldWriteRepresentedObjectsToPasteboard() bool
	// optional
	TokenFieldWriteRepresentedObjectsToPasteboard(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool
	ImplementsTokenFieldHasMenuForRepresentedObject() bool
	// optional
	TokenFieldHasMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) bool
	ImplementsTokenFieldMenuForRepresentedObject() bool
	// optional
	TokenFieldMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) IMenu
}

type TokenFieldDelegate struct {
	TextFieldDelegate
	_TokenFieldDisplayStringForRepresentedObject                      func(tokenField TokenField, representedObject objc.Object) string
	_TokenFieldStyleForRepresentedObject                              func(tokenField TokenField, representedObject objc.Object) TokenStyle
	_TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem func(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject
	_TokenFieldEditingStringForRepresentedObject                      func(tokenField TokenField, representedObject objc.Object) string
	_TokenFieldRepresentedObjectForEditingString                      func(tokenField TokenField, editingString string) objc.IObject
	_TokenFieldShouldAddObjectsAtIndex                                func(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject
	_TokenFieldReadFromPasteboard                                     func(tokenField TokenField, pboard Pasteboard) []objc.IObject
	_TokenFieldWriteRepresentedObjectsToPasteboard                    func(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool
	_TokenFieldHasMenuForRepresentedObject                            func(tokenField TokenField, representedObject objc.Object) bool
	_TokenFieldMenuForRepresentedObject                               func(tokenField TokenField, representedObject objc.Object) IMenu
}

func (di *TokenFieldDelegate) ImplementsTokenFieldDisplayStringForRepresentedObject() bool {
	return di._TokenFieldDisplayStringForRepresentedObject != nil
}

func (di *TokenFieldDelegate) SetTokenFieldDisplayStringForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) string) {
	di._TokenFieldDisplayStringForRepresentedObject = f
}

func (di *TokenFieldDelegate) TokenFieldDisplayStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	return di._TokenFieldDisplayStringForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldStyleForRepresentedObject() bool {
	return di._TokenFieldStyleForRepresentedObject != nil
}

func (di *TokenFieldDelegate) SetTokenFieldStyleForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) TokenStyle) {
	di._TokenFieldStyleForRepresentedObject = f
}

func (di *TokenFieldDelegate) TokenFieldStyleForRepresentedObject(tokenField TokenField, representedObject objc.Object) TokenStyle {
	return di._TokenFieldStyleForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return di._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem != nil
}

func (di *TokenFieldDelegate) SetTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(f func(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject) {
	di._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem = f
}

func (di *TokenFieldDelegate) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField TokenField, substring string, tokenIndex int, selectedIndex *int) []objc.IObject {
	return di._TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField, substring, tokenIndex, selectedIndex)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldEditingStringForRepresentedObject() bool {
	return di._TokenFieldEditingStringForRepresentedObject != nil
}

func (di *TokenFieldDelegate) SetTokenFieldEditingStringForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) string) {
	di._TokenFieldEditingStringForRepresentedObject = f
}

func (di *TokenFieldDelegate) TokenFieldEditingStringForRepresentedObject(tokenField TokenField, representedObject objc.Object) string {
	return di._TokenFieldEditingStringForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldRepresentedObjectForEditingString() bool {
	return di._TokenFieldRepresentedObjectForEditingString != nil
}

func (di *TokenFieldDelegate) SetTokenFieldRepresentedObjectForEditingString(f func(tokenField TokenField, editingString string) objc.IObject) {
	di._TokenFieldRepresentedObjectForEditingString = f
}

func (di *TokenFieldDelegate) TokenFieldRepresentedObjectForEditingString(tokenField TokenField, editingString string) objc.IObject {
	return di._TokenFieldRepresentedObjectForEditingString(tokenField, editingString)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldShouldAddObjectsAtIndex() bool {
	return di._TokenFieldShouldAddObjectsAtIndex != nil
}

func (di *TokenFieldDelegate) SetTokenFieldShouldAddObjectsAtIndex(f func(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject) {
	di._TokenFieldShouldAddObjectsAtIndex = f
}

func (di *TokenFieldDelegate) TokenFieldShouldAddObjectsAtIndex(tokenField TokenField, tokens []objc.Object, index uint) []objc.IObject {
	return di._TokenFieldShouldAddObjectsAtIndex(tokenField, tokens, index)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldReadFromPasteboard() bool {
	return di._TokenFieldReadFromPasteboard != nil
}

func (di *TokenFieldDelegate) SetTokenFieldReadFromPasteboard(f func(tokenField TokenField, pboard Pasteboard) []objc.IObject) {
	di._TokenFieldReadFromPasteboard = f
}

func (di *TokenFieldDelegate) TokenFieldReadFromPasteboard(tokenField TokenField, pboard Pasteboard) []objc.IObject {
	return di._TokenFieldReadFromPasteboard(tokenField, pboard)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldWriteRepresentedObjectsToPasteboard() bool {
	return di._TokenFieldWriteRepresentedObjectsToPasteboard != nil
}

func (di *TokenFieldDelegate) SetTokenFieldWriteRepresentedObjectsToPasteboard(f func(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool) {
	di._TokenFieldWriteRepresentedObjectsToPasteboard = f
}

func (di *TokenFieldDelegate) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField TokenField, objects []objc.Object, pboard Pasteboard) bool {
	return di._TokenFieldWriteRepresentedObjectsToPasteboard(tokenField, objects, pboard)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldHasMenuForRepresentedObject() bool {
	return di._TokenFieldHasMenuForRepresentedObject != nil
}

func (di *TokenFieldDelegate) SetTokenFieldHasMenuForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) bool) {
	di._TokenFieldHasMenuForRepresentedObject = f
}

func (di *TokenFieldDelegate) TokenFieldHasMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) bool {
	return di._TokenFieldHasMenuForRepresentedObject(tokenField, representedObject)
}
func (di *TokenFieldDelegate) ImplementsTokenFieldMenuForRepresentedObject() bool {
	return di._TokenFieldMenuForRepresentedObject != nil
}

func (di *TokenFieldDelegate) SetTokenFieldMenuForRepresentedObject(f func(tokenField TokenField, representedObject objc.Object) IMenu) {
	di._TokenFieldMenuForRepresentedObject = f
}

func (di *TokenFieldDelegate) TokenFieldMenuForRepresentedObject(tokenField TokenField, representedObject objc.Object) IMenu {
	return di._TokenFieldMenuForRepresentedObject(tokenField, representedObject)
}

type TokenFieldDelegateWrapper struct {
	TextFieldDelegateWrapper
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldDisplayStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:displayStringForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldDisplayStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:displayStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldStyleForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:styleForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldStyleForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) TokenStyle {
	rv := objc.CallMethod[TokenStyle](t_, objc.GetSelector("tokenField:styleForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldCompletionsForSubstringIndexOfTokenIndexOfSelectedItem(tokenField ITokenField, substring string, tokenIndex int, selectedIndex *int) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:completionsForSubstring:indexOfToken:indexOfSelectedItem:"), objc.ExtractPtr(tokenField), substring, tokenIndex, selectedIndex)
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldEditingStringForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:editingStringForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldEditingStringForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) string {
	rv := objc.CallMethod[string](t_, objc.GetSelector("tokenField:editingStringForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldRepresentedObjectForEditingString() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:representedObjectForEditingString:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldRepresentedObjectForEditingString(tokenField ITokenField, editingString string) objc.Object {
	rv := objc.CallMethod[objc.Object](t_, objc.GetSelector("tokenField:representedObjectForEditingString:"), objc.ExtractPtr(tokenField), editingString)
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldShouldAddObjectsAtIndex() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:shouldAddObjects:atIndex:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldShouldAddObjectsAtIndex(tokenField ITokenField, tokens []objc.IObject, index uint) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:shouldAddObjects:atIndex:"), objc.ExtractPtr(tokenField), tokens, index)
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldReadFromPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:readFromPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldReadFromPasteboard(tokenField ITokenField, pboard IPasteboard) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](t_, objc.GetSelector("tokenField:readFromPasteboard:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldWriteRepresentedObjectsToPasteboard() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldWriteRepresentedObjectsToPasteboard(tokenField ITokenField, objects []objc.IObject, pboard IPasteboard) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:writeRepresentedObjects:toPasteboard:"), objc.ExtractPtr(tokenField), objects, objc.ExtractPtr(pboard))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldHasMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:hasMenuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldHasMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) bool {
	rv := objc.CallMethod[bool](t_, objc.GetSelector("tokenField:hasMenuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}

func (t_ TokenFieldDelegateWrapper) ImplementsTokenFieldMenuForRepresentedObject() bool {
	return t_.RespondsToSelector(objc.GetSelector("tokenField:menuForRepresentedObject:"))
}

func (t_ TokenFieldDelegateWrapper) TokenFieldMenuForRepresentedObject(tokenField ITokenField, representedObject objc.IObject) Menu {
	rv := objc.CallMethod[Menu](t_, objc.GetSelector("tokenField:menuForRepresentedObject:"), objc.ExtractPtr(tokenField), objc.ExtractPtr(representedObject))
	return rv
}
