// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InflectionRuleExplicit] class.
var InflectionRuleExplicitClass = _InflectionRuleExplicitClass{objc.GetClass("NSInflectionRuleExplicit")}

type _InflectionRuleExplicitClass struct {
	objc.Class
}

// An interface definition for the [InflectionRuleExplicit] class.
type IInflectionRuleExplicit interface {
	IInflectionRule
	Morphology() Morphology
}

// An inflection rule that uses a morphology instance to determine how to inflect attribued strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionruleexplicit?language=objc
type InflectionRuleExplicit struct {
	InflectionRule
}

func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		InflectionRule: InflectionRuleFrom(ptr),
	}
}

func (i_ InflectionRuleExplicit) InitWithMorphology(morphology IMorphology) InflectionRuleExplicit {
	rv := objc.Call[InflectionRuleExplicit](i_, objc.Sel("initWithMorphology:"), objc.Ptr(morphology))
	return rv
}

// Creates an inflection rule with the given morphology. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionruleexplicit/3746914-initwithmorphology?language=objc
func NewInflectionRuleExplicitWithMorphology(morphology IMorphology) InflectionRuleExplicit {
	instance := InflectionRuleExplicitClass.Alloc().InitWithMorphology(morphology)
	instance.Autorelease()
	return instance
}

func (ic _InflectionRuleExplicitClass) Alloc() InflectionRuleExplicit {
	rv := objc.Call[InflectionRuleExplicit](ic, objc.Sel("alloc"))
	return rv
}

func InflectionRuleExplicit_Alloc() InflectionRuleExplicit {
	return InflectionRuleExplicitClass.Alloc()
}

func (ic _InflectionRuleExplicitClass) New() InflectionRuleExplicit {
	rv := objc.Call[InflectionRuleExplicit](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInflectionRuleExplicit() InflectionRuleExplicit {
	return InflectionRuleExplicitClass.New()
}

func (i_ InflectionRuleExplicit) Init() InflectionRuleExplicit {
	rv := objc.Call[InflectionRuleExplicit](i_, objc.Sel("init"))
	return rv
}

// The morphology used by this inflection rule. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsinflectionruleexplicit/3746915-morphology?language=objc
func (i_ InflectionRuleExplicit) Morphology() Morphology {
	rv := objc.Call[Morphology](i_, objc.Sel("morphology"))
	return rv
}
