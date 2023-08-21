// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitMass] class.
var UnitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}

type _UnitMassClass struct {
	objc.Class
}

// An interface definition for the [UnitMass] class.
type IUnitMass interface {
	IDimension
}

// A unit of measure for mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass?language=objc
type UnitMass struct {
	Dimension
}

func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitMassClass) Alloc() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("alloc"))
	return rv
}

func UnitMass_Alloc() UnitMass {
	return UnitMassClass.Alloc()
}

func (uc _UnitMassClass) New() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitMass() UnitMass {
	return UnitMassClass.New()
}

func (u_ UnitMass) Init() UnitMass {
	rv := objc.Call[UnitMass](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitMassClass) BaseUnit() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitMass_BaseUnit() UnitMass {
	return UnitMassClass.BaseUnit()
}

func (u_ UnitMass) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitMass {
	rv := objc.Call[UnitMass](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitMassWithSymbolConverter(symbol string, converter IUnitConverter) UnitMass {
	instance := UnitMassClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitMass) InitWithSymbol(symbol string) UnitMass {
	rv := objc.Call[UnitMass](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitMassWithSymbol(symbol string) UnitMass {
	instance := UnitMassClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The micrograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856010-micrograms?language=objc
func (uc _UnitMassClass) Micrograms() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("micrograms"))
	return rv
}

// The micrograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856010-micrograms?language=objc
func UnitMass_Micrograms() UnitMass {
	return UnitMassClass.Micrograms()
}

// The milligrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856060-milligrams?language=objc
func (uc _UnitMassClass) Milligrams() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("milligrams"))
	return rv
}

// The milligrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856060-milligrams?language=objc
func UnitMass_Milligrams() UnitMass {
	return UnitMassClass.Milligrams()
}

// The grams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1855976-grams?language=objc
func (uc _UnitMassClass) Grams() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("grams"))
	return rv
}

// The grams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1855976-grams?language=objc
func UnitMass_Grams() UnitMass {
	return UnitMassClass.Grams()
}

// The carats unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856037-carats?language=objc
func (uc _UnitMassClass) Carats() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("carats"))
	return rv
}

// The carats unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856037-carats?language=objc
func UnitMass_Carats() UnitMass {
	return UnitMassClass.Carats()
}

// The short tons unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856081-shorttons?language=objc
func (uc _UnitMassClass) ShortTons() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("shortTons"))
	return rv
}

// The short tons unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856081-shorttons?language=objc
func UnitMass_ShortTons() UnitMass {
	return UnitMassClass.ShortTons()
}

// The nanograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856078-nanograms?language=objc
func (uc _UnitMassClass) Nanograms() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("nanograms"))
	return rv
}

// The nanograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856078-nanograms?language=objc
func UnitMass_Nanograms() UnitMass {
	return UnitMassClass.Nanograms()
}

// The centigrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856116-centigrams?language=objc
func (uc _UnitMassClass) Centigrams() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("centigrams"))
	return rv
}

// The centigrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856116-centigrams?language=objc
func UnitMass_Centigrams() UnitMass {
	return UnitMassClass.Centigrams()
}

// The kilograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1855996-kilograms?language=objc
func (uc _UnitMassClass) Kilograms() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("kilograms"))
	return rv
}

// The kilograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1855996-kilograms?language=objc
func UnitMass_Kilograms() UnitMass {
	return UnitMassClass.Kilograms()
}

// The ounces troy unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856003-ouncestroy?language=objc
func (uc _UnitMassClass) OuncesTroy() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("ouncesTroy"))
	return rv
}

// The ounces troy unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856003-ouncestroy?language=objc
func UnitMass_OuncesTroy() UnitMass {
	return UnitMassClass.OuncesTroy()
}

// The decigrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856063-decigrams?language=objc
func (uc _UnitMassClass) Decigrams() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("decigrams"))
	return rv
}

// The decigrams unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856063-decigrams?language=objc
func UnitMass_Decigrams() UnitMass {
	return UnitMassClass.Decigrams()
}

// The slugs unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856027-slugs?language=objc
func (uc _UnitMassClass) Slugs() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("slugs"))
	return rv
}

// The slugs unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856027-slugs?language=objc
func UnitMass_Slugs() UnitMass {
	return UnitMassClass.Slugs()
}

// The pounds unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856023-poundsmass?language=objc
func (uc _UnitMassClass) PoundsMass() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("poundsMass"))
	return rv
}

// The pounds unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856023-poundsmass?language=objc
func UnitMass_PoundsMass() UnitMass {
	return UnitMassClass.PoundsMass()
}

// The ounces unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856056-ounces?language=objc
func (uc _UnitMassClass) Ounces() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("ounces"))
	return rv
}

// The ounces unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856056-ounces?language=objc
func UnitMass_Ounces() UnitMass {
	return UnitMassClass.Ounces()
}

// The stone unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856033-stones?language=objc
func (uc _UnitMassClass) Stones() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("stones"))
	return rv
}

// The stone unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856033-stones?language=objc
func UnitMass_Stones() UnitMass {
	return UnitMassClass.Stones()
}

// The metric tons unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856076-metrictons?language=objc
func (uc _UnitMassClass) MetricTons() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("metricTons"))
	return rv
}

// The metric tons unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856076-metrictons?language=objc
func UnitMass_MetricTons() UnitMass {
	return UnitMassClass.MetricTons()
}

// The picograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856035-picograms?language=objc
func (uc _UnitMassClass) Picograms() UnitMass {
	rv := objc.Call[UnitMass](uc, objc.Sel("picograms"))
	return rv
}

// The picograms unit of mass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitmass/1856035-picograms?language=objc
func UnitMass_Picograms() UnitMass {
	return UnitMassClass.Picograms()
}
