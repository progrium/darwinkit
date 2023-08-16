// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a candidate list item delegate uses to enable selection state and list visibility. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate?language=objc
type PCandidateListTouchBarItemDelegate interface {
	// optional
	CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int)
	HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool
}

// A delegate implementation builder for the [PCandidateListTouchBarItemDelegate] protocol.
type CandidateListTouchBarItemDelegate struct {
	_CandidateListTouchBarItemBeginSelectingCandidateAtIndex func(anItem CandidateListTouchBarItem, index int)
}

func (di *CandidateListTouchBarItemDelegate) HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool {
	return di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex != nil
}

// Tells the delegate that the user has started touching one of the candidates in the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate/2544747-candidatelisttouchbaritem?language=objc
func (di *CandidateListTouchBarItemDelegate) SetCandidateListTouchBarItemBeginSelectingCandidateAtIndex(f func(anItem CandidateListTouchBarItem, index int)) {
	di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex = f
}

// Tells the delegate that the user has started touching one of the candidates in the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate/2544747-candidatelisttouchbaritem?language=objc
func (di *CandidateListTouchBarItemDelegate) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem CandidateListTouchBarItem, index int) {
	di._CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem, index)
}

// A concrete type wrapper for the [PCandidateListTouchBarItemDelegate] protocol.
type CandidateListTouchBarItemDelegateWrapper struct {
	objc.Object
}

func (c_ CandidateListTouchBarItemDelegateWrapper) HasCandidateListTouchBarItemBeginSelectingCandidateAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"))
}

// Tells the delegate that the user has started touching one of the candidates in the candidate list item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscandidatelisttouchbaritemdelegate/2544747-candidatelisttouchbaritem?language=objc
func (c_ CandidateListTouchBarItemDelegateWrapper) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem ICandidateListTouchBarItem, index int) {
	objc.Call[objc.Void](c_, objc.Sel("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), objc.Ptr(anItem), index)
}
