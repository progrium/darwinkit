// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient?language=objc
type PTextCheckingClient interface {
	// optional
	RemoveAnnotationRange(annotationName foundation.AttributedStringKey, range_ foundation.Range)
	HasRemoveAnnotationRange() bool

	// optional
	ViewForRangeFirstRectActualRange(range_ foundation.Range, firstRect foundation.RectPointer, actualRange foundation.RangePointer) IView
	HasViewForRangeFirstRectActualRange() bool

	// optional
	AddAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range)
	HasAddAnnotationsRange() bool

	// optional
	SetAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range)
	HasSetAnnotationsRange() bool

	// optional
	CandidateListTouchBarItem() ICandidateListTouchBarItem
	HasCandidateListTouchBarItem() bool

	// optional
	ReplaceCharactersInRangeWithAnnotatedString(range_ foundation.Range, annotatedString foundation.AttributedString)
	HasReplaceCharactersInRangeWithAnnotatedString() bool

	// optional
	AnnotatedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.IAttributedString
	HasAnnotatedSubstringForProposedRangeActualRange() bool

	// optional
	SelectAndShowRange(range_ foundation.Range)
	HasSelectAndShowRange() bool
}

// A concrete type wrapper for the [PTextCheckingClient] protocol.
type TextCheckingClientWrapper struct {
	objc.Object
}

func (t_ TextCheckingClientWrapper) HasRemoveAnnotationRange() bool {
	return t_.RespondsToSelector(objc.Sel("removeAnnotation:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242722-removeannotation?language=objc
func (t_ TextCheckingClientWrapper) RemoveAnnotationRange(annotationName foundation.AttributedStringKey, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("removeAnnotation:range:"), annotationName, range_)
}

func (t_ TextCheckingClientWrapper) HasViewForRangeFirstRectActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("viewForRange:firstRect:actualRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242726-viewforrange?language=objc
func (t_ TextCheckingClientWrapper) ViewForRangeFirstRectActualRange(range_ foundation.Range, firstRect foundation.RectPointer, actualRange foundation.RangePointer) View {
	rv := objc.Call[View](t_, objc.Sel("viewForRange:firstRect:actualRange:"), range_, firstRect, actualRange)
	return rv
}

func (t_ TextCheckingClientWrapper) HasAddAnnotationsRange() bool {
	return t_.RespondsToSelector(objc.Sel("addAnnotations:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242719-addannotations?language=objc
func (t_ TextCheckingClientWrapper) AddAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("addAnnotations:range:"), annotations, range_)
}

func (t_ TextCheckingClientWrapper) HasSetAnnotationsRange() bool {
	return t_.RespondsToSelector(objc.Sel("setAnnotations:range:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242725-setannotations?language=objc
func (t_ TextCheckingClientWrapper) SetAnnotationsRange(annotations map[foundation.AttributedStringKey]string, range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("setAnnotations:range:"), annotations, range_)
}

func (t_ TextCheckingClientWrapper) HasCandidateListTouchBarItem() bool {
	return t_.RespondsToSelector(objc.Sel("candidateListTouchBarItem"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242721-candidatelisttouchbaritem?language=objc
func (t_ TextCheckingClientWrapper) CandidateListTouchBarItem() CandidateListTouchBarItem {
	rv := objc.Call[CandidateListTouchBarItem](t_, objc.Sel("candidateListTouchBarItem"))
	return rv
}

func (t_ TextCheckingClientWrapper) HasReplaceCharactersInRangeWithAnnotatedString() bool {
	return t_.RespondsToSelector(objc.Sel("replaceCharactersInRange:withAnnotatedString:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242723-replacecharactersinrange?language=objc
func (t_ TextCheckingClientWrapper) ReplaceCharactersInRangeWithAnnotatedString(range_ foundation.Range, annotatedString foundation.IAttributedString) {
	objc.Call[objc.Void](t_, objc.Sel("replaceCharactersInRange:withAnnotatedString:"), range_, objc.Ptr(annotatedString))
}

func (t_ TextCheckingClientWrapper) HasAnnotatedSubstringForProposedRangeActualRange() bool {
	return t_.RespondsToSelector(objc.Sel("annotatedSubstringForProposedRange:actualRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242720-annotatedsubstringforproposedran?language=objc
func (t_ TextCheckingClientWrapper) AnnotatedSubstringForProposedRangeActualRange(range_ foundation.Range, actualRange foundation.RangePointer) foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](t_, objc.Sel("annotatedSubstringForProposedRange:actualRange:"), range_, actualRange)
	return rv
}

func (t_ TextCheckingClientWrapper) HasSelectAndShowRange() bool {
	return t_.RespondsToSelector(objc.Sel("selectAndShowRange:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingclient/3242724-selectandshowrange?language=objc
func (t_ TextCheckingClientWrapper) SelectAndShowRange(range_ foundation.Range) {
	objc.Call[objc.Void](t_, objc.Sel("selectAndShowRange:"), range_)
}
