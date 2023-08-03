// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IRuleEditorDelegate interface {
	// required
	RuleEditorChildForCriterionWithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject
	// required
	RuleEditorDisplayValueForCriterionInRow(editor RuleEditor, criterion objc.Object, row int) objc.IObject
	// required
	RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	ImplementsRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool
	// optional
	RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject
	ImplementsRuleEditorRowsDidChange() bool
	// optional
	RuleEditorRowsDidChange(notification foundation.Notification)
}

type RuleEditorDelegate struct {
	_RuleEditorChildForCriterionWithRowType                    func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject
	_RuleEditorDisplayValueForCriterionInRow                   func(editor RuleEditor, criterion objc.Object, row int) objc.IObject
	_RuleEditorNumberOfChildrenForCriterionWithRowType         func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int
	_RuleEditorPredicatePartsForCriterionWithDisplayValueInRow func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject
	_RuleEditorRowsDidChange                                   func(notification foundation.Notification)
}

func (di *RuleEditorDelegate) SetRuleEditorChildForCriterionWithRowType(f func(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject) {
	di._RuleEditorChildForCriterionWithRowType = f
}

// required

func (di *RuleEditorDelegate) RuleEditorChildForCriterionWithRowType(editor RuleEditor, index int, criterion objc.Object, rowType RuleEditorRowType) objc.IObject {
	return di._RuleEditorChildForCriterionWithRowType(editor, index, criterion, rowType)
}

func (di *RuleEditorDelegate) SetRuleEditorDisplayValueForCriterionInRow(f func(editor RuleEditor, criterion objc.Object, row int) objc.IObject) {
	di._RuleEditorDisplayValueForCriterionInRow = f
}

// required

func (di *RuleEditorDelegate) RuleEditorDisplayValueForCriterionInRow(editor RuleEditor, criterion objc.Object, row int) objc.IObject {
	return di._RuleEditorDisplayValueForCriterionInRow(editor, criterion, row)
}

func (di *RuleEditorDelegate) SetRuleEditorNumberOfChildrenForCriterionWithRowType(f func(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int) {
	di._RuleEditorNumberOfChildrenForCriterionWithRowType = f
}

// required

func (di *RuleEditorDelegate) RuleEditorNumberOfChildrenForCriterionWithRowType(editor RuleEditor, criterion objc.Object, rowType RuleEditorRowType) int {
	return di._RuleEditorNumberOfChildrenForCriterionWithRowType(editor, criterion, rowType)
}
func (di *RuleEditorDelegate) ImplementsRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow != nil
}

func (di *RuleEditorDelegate) SetRuleEditorPredicatePartsForCriterionWithDisplayValueInRow(f func(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject) {
	di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow = f
}

func (di *RuleEditorDelegate) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor RuleEditor, criterion objc.Object, value objc.Object, row int) map[RuleEditorPredicatePartKey]objc.IObject {
	return di._RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor, criterion, value, row)
}
func (di *RuleEditorDelegate) ImplementsRuleEditorRowsDidChange() bool {
	return di._RuleEditorRowsDidChange != nil
}

func (di *RuleEditorDelegate) SetRuleEditorRowsDidChange(f func(notification foundation.Notification)) {
	di._RuleEditorRowsDidChange = f
}

func (di *RuleEditorDelegate) RuleEditorRowsDidChange(notification foundation.Notification) {
	di._RuleEditorRowsDidChange(notification)
}

type RuleEditorDelegateWrapper struct {
	objc.Object
}

func (r_ RuleEditorDelegateWrapper) RuleEditorChildForCriterionWithRowType(editor IRuleEditor, index int, criterion objc.IObject, rowType RuleEditorRowType) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("ruleEditor:child:forCriterion:withRowType:"), objc.ExtractPtr(editor), index, objc.ExtractPtr(criterion), rowType)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditorDisplayValueForCriterionInRow(editor IRuleEditor, criterion objc.IObject, row int) objc.Object {
	rv := objc.CallMethod[objc.Object](r_, objc.GetSelector("ruleEditor:displayValueForCriterion:inRow:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), row)
	return rv
}

func (r_ RuleEditorDelegateWrapper) RuleEditorNumberOfChildrenForCriterionWithRowType(editor IRuleEditor, criterion objc.IObject, rowType RuleEditorRowType) int {
	rv := objc.CallMethod[int](r_, objc.GetSelector("ruleEditor:numberOfChildrenForCriterion:withRowType:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), rowType)
	return rv
}

func (r_ RuleEditorDelegateWrapper) ImplementsRuleEditorPredicatePartsForCriterionWithDisplayValueInRow() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditorPredicatePartsForCriterionWithDisplayValueInRow(editor IRuleEditor, criterion objc.IObject, value objc.IObject, row int) map[RuleEditorPredicatePartKey]objc.Object {
	rv := objc.CallMethod[map[RuleEditorPredicatePartKey]objc.Object](r_, objc.GetSelector("ruleEditor:predicatePartsForCriterion:withDisplayValue:inRow:"), objc.ExtractPtr(editor), objc.ExtractPtr(criterion), objc.ExtractPtr(value), row)
	return rv
}

func (r_ RuleEditorDelegateWrapper) ImplementsRuleEditorRowsDidChange() bool {
	return r_.RespondsToSelector(objc.GetSelector("ruleEditorRowsDidChange:"))
}

func (r_ RuleEditorDelegateWrapper) RuleEditorRowsDidChange(notification foundation.INotification) {
	objc.CallMethod[objc.Void](r_, objc.GetSelector("ruleEditorRowsDidChange:"), objc.ExtractPtr(notification))
}
