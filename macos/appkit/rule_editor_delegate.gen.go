// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The NSRuleEditorDelegate protocol defines the optional methods implemented by delegates of NSRuleEditor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate?language=objc
type PRuleEditorDelegate interface {
	// optional
	RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool

	// optional
	RuleEditorRowsDidChange(notification foundation.Notification)
	HasRuleEditorRowsDidChange() bool
}

// A delegate implementation builder for the [PRuleEditorDelegate] protocol.
type RuleEditorDelegate struct {
	_RuleEditorNumberOfChildrenForCriterionWithRowType func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	_RuleEditorRowsDidChange                           func(notification foundation.Notification)
}

func (di *RuleEditorDelegate) HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool {
	return di._RuleEditorNumberOfChildrenForCriterionWithRowType != nil
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (di *RuleEditorDelegate) SetRuleEditorNumberOfChildrenForCriterionWithRowType(f func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int) {
	di._RuleEditorNumberOfChildrenForCriterionWithRowType = f
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (di *RuleEditorDelegate) RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int {
	return di._RuleEditorNumberOfChildrenForCriterionWithRowType(editor, criterion, rowType)
}
func (di *RuleEditorDelegate) HasRuleEditorRowsDidChange() bool {
	return di._RuleEditorRowsDidChange != nil
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (di *RuleEditorDelegate) SetRuleEditorRowsDidChange(f func(notification foundation.Notification)) {
	di._RuleEditorRowsDidChange = f
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (di *RuleEditorDelegate) RuleEditorRowsDidChange(notification foundation.Notification) {
	di._RuleEditorRowsDidChange(notification)
}

// A concrete type wrapper for the [PRuleEditorDelegate] protocol.
type RuleEditorDelegateWrapper struct {
	objc.Object
}

func (r_ RuleEditorDelegateWrapper) HasRuleEditorNumberOfChildrenForCriterionWithRowType() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"))
}

// Returns the number of child items of a given criterion or row type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1535329-ruleeditor?language=objc
func (r_ RuleEditorDelegateWrapper) RuleEditorNumberOfChildrenForCriterionWithRowType(editor IRuleEditor, criterion objc.IObject, rowType RuleEditorRowType) int {
	rv := objc.Call[int](r_, objc.Sel("ruleEditor:numberOfChildrenForCriterion:withRowType:"), objc.Ptr(editor), criterion, rowType)
	return rv
}

func (r_ RuleEditorDelegateWrapper) HasRuleEditorRowsDidChange() bool {
	return r_.RespondsToSelector(objc.Sel("ruleEditorRowsDidChange:"))
}

// Notifies the receiver that a rule editor’s rows changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsruleeditordelegate/1533292-ruleeditorrowsdidchange?language=objc
func (r_ RuleEditorDelegateWrapper) RuleEditorRowsDidChange(notification foundation.INotification) {
	objc.Call[objc.Void](r_, objc.Sel("ruleEditorRowsDidChange:"), objc.Ptr(notification))
}
