// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an accordion fold transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition?language=objc
type PAccordionFoldTransition interface {
	// optional
	SetBottomHeight(value float64)
	HasSetBottomHeight() bool

	// optional
	BottomHeight() float64
	HasBottomHeight() bool

	// optional
	SetFoldShadowAmount(value float64)
	HasSetFoldShadowAmount() bool

	// optional
	FoldShadowAmount() float64
	HasFoldShadowAmount() bool

	// optional
	SetNumberOfFolds(value float64)
	HasSetNumberOfFolds() bool

	// optional
	NumberOfFolds() float64
	HasNumberOfFolds() bool
}

// A concrete type wrapper for the [PAccordionFoldTransition] protocol.
type AccordionFoldTransitionWrapper struct {
	objc.Object
}

func (a_ AccordionFoldTransitionWrapper) HasSetBottomHeight() bool {
	return a_.RespondsToSelector(objc.Sel("setBottomHeight:"))
}

// The height of the accordion-fold part of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228050-bottomheight?language=objc
func (a_ AccordionFoldTransitionWrapper) SetBottomHeight(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setBottomHeight:"), value)
}

func (a_ AccordionFoldTransitionWrapper) HasBottomHeight() bool {
	return a_.RespondsToSelector(objc.Sel("bottomHeight"))
}

// The height of the accordion-fold part of the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228050-bottomheight?language=objc
func (a_ AccordionFoldTransitionWrapper) BottomHeight() float64 {
	rv := objc.Call[float64](a_, objc.Sel("bottomHeight"))
	return rv
}

func (a_ AccordionFoldTransitionWrapper) HasSetFoldShadowAmount() bool {
	return a_.RespondsToSelector(objc.Sel("setFoldShadowAmount:"))
}

// A value that specifies the intensity of the shadow in the transtion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228051-foldshadowamount?language=objc
func (a_ AccordionFoldTransitionWrapper) SetFoldShadowAmount(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setFoldShadowAmount:"), value)
}

func (a_ AccordionFoldTransitionWrapper) HasFoldShadowAmount() bool {
	return a_.RespondsToSelector(objc.Sel("foldShadowAmount"))
}

// A value that specifies the intensity of the shadow in the transtion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228051-foldshadowamount?language=objc
func (a_ AccordionFoldTransitionWrapper) FoldShadowAmount() float64 {
	rv := objc.Call[float64](a_, objc.Sel("foldShadowAmount"))
	return rv
}

func (a_ AccordionFoldTransitionWrapper) HasSetNumberOfFolds() bool {
	return a_.RespondsToSelector(objc.Sel("setNumberOfFolds:"))
}

// The number of folds used in the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228052-numberoffolds?language=objc
func (a_ AccordionFoldTransitionWrapper) SetNumberOfFolds(value float64) {
	objc.Call[objc.Void](a_, objc.Sel("setNumberOfFolds:"), value)
}

func (a_ AccordionFoldTransitionWrapper) HasNumberOfFolds() bool {
	return a_.RespondsToSelector(objc.Sel("numberOfFolds"))
}

// The number of folds used in the transition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaccordionfoldtransition/3228052-numberoffolds?language=objc
func (a_ AccordionFoldTransitionWrapper) NumberOfFolds() float64 {
	rv := objc.Call[float64](a_, objc.Sel("numberOfFolds"))
	return rv
}
