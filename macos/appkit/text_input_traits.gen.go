// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits?language=objc
type PTextInputTraits interface {
	// optional
	SetTextCompletionType(value TextInputTraitType)
	HasSetTextCompletionType() bool

	// optional
	TextCompletionType() TextInputTraitType
	HasTextCompletionType() bool

	// optional
	SetSmartQuotesType(value TextInputTraitType)
	HasSetSmartQuotesType() bool

	// optional
	SmartQuotesType() TextInputTraitType
	HasSmartQuotesType() bool

	// optional
	SetGrammarCheckingType(value TextInputTraitType)
	HasSetGrammarCheckingType() bool

	// optional
	GrammarCheckingType() TextInputTraitType
	HasGrammarCheckingType() bool

	// optional
	SetSpellCheckingType(value TextInputTraitType)
	HasSetSpellCheckingType() bool

	// optional
	SpellCheckingType() TextInputTraitType
	HasSpellCheckingType() bool

	// optional
	SetAutocorrectionType(value TextInputTraitType)
	HasSetAutocorrectionType() bool

	// optional
	AutocorrectionType() TextInputTraitType
	HasAutocorrectionType() bool

	// optional
	SetSmartDashesType(value TextInputTraitType)
	HasSetSmartDashesType() bool

	// optional
	SmartDashesType() TextInputTraitType
	HasSmartDashesType() bool

	// optional
	SetSmartInsertDeleteType(value TextInputTraitType)
	HasSetSmartInsertDeleteType() bool

	// optional
	SmartInsertDeleteType() TextInputTraitType
	HasSmartInsertDeleteType() bool

	// optional
	SetLinkDetectionType(value TextInputTraitType)
	HasSetLinkDetectionType() bool

	// optional
	LinkDetectionType() TextInputTraitType
	HasLinkDetectionType() bool

	// optional
	SetTextReplacementType(value TextInputTraitType)
	HasSetTextReplacementType() bool

	// optional
	TextReplacementType() TextInputTraitType
	HasTextReplacementType() bool

	// optional
	SetDataDetectionType(value TextInputTraitType)
	HasSetDataDetectionType() bool

	// optional
	DataDetectionType() TextInputTraitType
	HasDataDetectionType() bool
}

// A concrete type wrapper for the [PTextInputTraits] protocol.
type TextInputTraitsWrapper struct {
	objc.Object
}

func (t_ TextInputTraitsWrapper) HasSetTextCompletionType() bool {
	return t_.RespondsToSelector(objc.Sel("setTextCompletionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242740-textcompletiontype?language=objc
func (t_ TextInputTraitsWrapper) SetTextCompletionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setTextCompletionType:"), value)
}

func (t_ TextInputTraitsWrapper) HasTextCompletionType() bool {
	return t_.RespondsToSelector(objc.Sel("textCompletionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242740-textcompletiontype?language=objc
func (t_ TextInputTraitsWrapper) TextCompletionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("textCompletionType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetSmartQuotesType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartQuotesType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242738-smartquotestype?language=objc
func (t_ TextInputTraitsWrapper) SetSmartQuotesType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartQuotesType:"), value)
}

func (t_ TextInputTraitsWrapper) HasSmartQuotesType() bool {
	return t_.RespondsToSelector(objc.Sel("smartQuotesType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242738-smartquotestype?language=objc
func (t_ TextInputTraitsWrapper) SmartQuotesType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartQuotesType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetGrammarCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("setGrammarCheckingType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242734-grammarcheckingtype?language=objc
func (t_ TextInputTraitsWrapper) SetGrammarCheckingType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setGrammarCheckingType:"), value)
}

func (t_ TextInputTraitsWrapper) HasGrammarCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("grammarCheckingType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242734-grammarcheckingtype?language=objc
func (t_ TextInputTraitsWrapper) GrammarCheckingType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("grammarCheckingType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetSpellCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("setSpellCheckingType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242739-spellcheckingtype?language=objc
func (t_ TextInputTraitsWrapper) SetSpellCheckingType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSpellCheckingType:"), value)
}

func (t_ TextInputTraitsWrapper) HasSpellCheckingType() bool {
	return t_.RespondsToSelector(objc.Sel("spellCheckingType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242739-spellcheckingtype?language=objc
func (t_ TextInputTraitsWrapper) SpellCheckingType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("spellCheckingType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetAutocorrectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setAutocorrectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242732-autocorrectiontype?language=objc
func (t_ TextInputTraitsWrapper) SetAutocorrectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setAutocorrectionType:"), value)
}

func (t_ TextInputTraitsWrapper) HasAutocorrectionType() bool {
	return t_.RespondsToSelector(objc.Sel("autocorrectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242732-autocorrectiontype?language=objc
func (t_ TextInputTraitsWrapper) AutocorrectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("autocorrectionType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetSmartDashesType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartDashesType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242736-smartdashestype?language=objc
func (t_ TextInputTraitsWrapper) SetSmartDashesType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartDashesType:"), value)
}

func (t_ TextInputTraitsWrapper) HasSmartDashesType() bool {
	return t_.RespondsToSelector(objc.Sel("smartDashesType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242736-smartdashestype?language=objc
func (t_ TextInputTraitsWrapper) SmartDashesType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartDashesType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetSmartInsertDeleteType() bool {
	return t_.RespondsToSelector(objc.Sel("setSmartInsertDeleteType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242737-smartinsertdeletetype?language=objc
func (t_ TextInputTraitsWrapper) SetSmartInsertDeleteType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setSmartInsertDeleteType:"), value)
}

func (t_ TextInputTraitsWrapper) HasSmartInsertDeleteType() bool {
	return t_.RespondsToSelector(objc.Sel("smartInsertDeleteType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242737-smartinsertdeletetype?language=objc
func (t_ TextInputTraitsWrapper) SmartInsertDeleteType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("smartInsertDeleteType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetLinkDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setLinkDetectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242735-linkdetectiontype?language=objc
func (t_ TextInputTraitsWrapper) SetLinkDetectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setLinkDetectionType:"), value)
}

func (t_ TextInputTraitsWrapper) HasLinkDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("linkDetectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242735-linkdetectiontype?language=objc
func (t_ TextInputTraitsWrapper) LinkDetectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("linkDetectionType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetTextReplacementType() bool {
	return t_.RespondsToSelector(objc.Sel("setTextReplacementType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242741-textreplacementtype?language=objc
func (t_ TextInputTraitsWrapper) SetTextReplacementType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setTextReplacementType:"), value)
}

func (t_ TextInputTraitsWrapper) HasTextReplacementType() bool {
	return t_.RespondsToSelector(objc.Sel("textReplacementType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242741-textreplacementtype?language=objc
func (t_ TextInputTraitsWrapper) TextReplacementType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("textReplacementType"))
	return rv
}

func (t_ TextInputTraitsWrapper) HasSetDataDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("setDataDetectionType:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242733-datadetectiontype?language=objc
func (t_ TextInputTraitsWrapper) SetDataDetectionType(value TextInputTraitType) {
	objc.Call[objc.Void](t_, objc.Sel("setDataDetectionType:"), value)
}

func (t_ TextInputTraitsWrapper) HasDataDetectionType() bool {
	return t_.RespondsToSelector(objc.Sel("dataDetectionType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextinputtraits/3242733-datadetectiontype?language=objc
func (t_ TextInputTraitsWrapper) DataDetectionType() TextInputTraitType {
	rv := objc.Call[TextInputTraitType](t_, objc.Sel("dataDetectionType"))
	return rv
}
