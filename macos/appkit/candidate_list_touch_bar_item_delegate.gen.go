// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type ICandidateListTouchBarItemDelegate interface {
	ImplementsCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool
	// optional
	CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int)
	ImplementsCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex() bool
	// optional
	CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem CandidateListTouchBarItem, previousIndex int, index int)
	ImplementsCandidateListTouchBarItemEndSelectingCandidateAtIndex() bool
	// optional
	CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int)
	ImplementsCandidateListTouchBarItemChangedCandidateListVisibility() bool
	// optional
	CandidateListTouchBarItemChangedCandidateListVisibility(anItem CandidateListTouchBarItem, isVisible bool)
}

type CandidateListTouchBarItemDelegate struct {
	_CandidateListTouchBarItemBeginSelectingCandidateAtIndex             func(anItem CandidateListTouchBarItem, index int)
	_CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex func(anItem CandidateListTouchBarItem, previousIndex int, index int)
	_CandidateListTouchBarItemEndSelectingCandidateAtIndex               func(anItem CandidateListTouchBarItem, index int)
	_CandidateListTouchBarItemChangedCandidateListVisibility             func(anItem CandidateListTouchBarItem, isVisible bool)
}

func (di *CandidateListTouchBarItemDelegate) ImplementsCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool {
	return di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex != nil
}

func (di *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemBeginSelectingCandidateAtIndex(f func(anItem CandidateListTouchBarItem, index int)) {
	di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex = f
}

func (di *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem, index)
}
func (di *CandidateListTouchBarItemDelegate) ImplementsCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex() bool {
	return di._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex != nil
}

func (di *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(f func(anItem CandidateListTouchBarItem, previousIndex int, index int)) {
	di._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex = f
}

func (di *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem CandidateListTouchBarItem, previousIndex int, index int) {
	di._CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem, previousIndex, index)
}
func (di *CandidateListTouchBarItemDelegate) ImplementsCandidateListTouchBarItemEndSelectingCandidateAtIndex() bool {
	return di._CandidateListTouchBarItemEndSelectingCandidateAtIndex != nil
}

func (di *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemEndSelectingCandidateAtIndex(f func(anItem CandidateListTouchBarItem, index int)) {
	di._CandidateListTouchBarItemEndSelectingCandidateAtIndex = f
}

func (di *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	di._CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem, index)
}
func (di *CandidateListTouchBarItemDelegate) ImplementsCandidateListTouchBarItemChangedCandidateListVisibility() bool {
	return di._CandidateListTouchBarItemChangedCandidateListVisibility != nil
}

func (di *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemChangedCandidateListVisibility(f func(anItem CandidateListTouchBarItem, isVisible bool)) {
	di._CandidateListTouchBarItemChangedCandidateListVisibility = f
}

func (di *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemChangedCandidateListVisibility(anItem CandidateListTouchBarItem, isVisible bool) {
	di._CandidateListTouchBarItemChangedCandidateListVisibility(anItem, isVisible)
}

type CandidateListTouchBarItemDelegateWrapper struct {
	objc.Object
}

func (c_ CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), objc.ExtractPtr(anItem), index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem ICandidateListTouchBarItem, previousIndex int, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), objc.ExtractPtr(anItem), previousIndex, index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItemEndSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), objc.ExtractPtr(anItem), index)
}

func (c_ CandidateListTouchBarItemDelegateWrapper) ImplementsCandidateListTouchBarItemChangedCandidateListVisibility() bool {
	return c_.RespondsToSelector(objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"))
}

func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItemChangedCandidateListVisibility(anItem ICandidateListTouchBarItem, isVisible bool) {
	objc.CallMethod[objc.Void](c_, objc.GetSelector("candidateListTouchBarItem:changedCandidateListVisibility:"), objc.ExtractPtr(anItem), isVisible)
}
