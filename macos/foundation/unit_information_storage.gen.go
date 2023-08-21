// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UnitInformationStorage] class.
var UnitInformationStorageClass = _UnitInformationStorageClass{objc.GetClass("NSUnitInformationStorage")}

type _UnitInformationStorageClass struct {
	objc.Class
}

// An interface definition for the [UnitInformationStorage] class.
type IUnitInformationStorage interface {
	IDimension
}

// A unit of measure for quantities of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage?language=objc
type UnitInformationStorage struct {
	Dimension
}

func UnitInformationStorageFrom(ptr unsafe.Pointer) UnitInformationStorage {
	return UnitInformationStorage{
		Dimension: DimensionFrom(ptr),
	}
}

func (uc _UnitInformationStorageClass) Alloc() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("alloc"))
	return rv
}

func UnitInformationStorage_Alloc() UnitInformationStorage {
	return UnitInformationStorageClass.Alloc()
}

func (uc _UnitInformationStorageClass) New() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUnitInformationStorage() UnitInformationStorage {
	return UnitInformationStorageClass.New()
}

func (u_ UnitInformationStorage) Init() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](u_, objc.Sel("init"))
	return rv
}

func (uc _UnitInformationStorageClass) BaseUnit() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("baseUnit"))
	return rv
}

// Returns the base unit. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1690740-baseunit?language=objc
func UnitInformationStorage_BaseUnit() UnitInformationStorage {
	return UnitInformationStorageClass.BaseUnit()
}

func (u_ UnitInformationStorage) InitWithSymbolConverter(symbol string, converter IUnitConverter) UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](u_, objc.Sel("initWithSymbol:converter:"), symbol, objc.Ptr(converter))
	return rv
}

// Initializes a dimensional unit with the symbol and unit converter you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdimension/1823633-initwithsymbol?language=objc
func NewUnitInformationStorageWithSymbolConverter(symbol string, converter IUnitConverter) UnitInformationStorage {
	instance := UnitInformationStorageClass.Alloc().InitWithSymbolConverter(symbol, converter)
	instance.Autorelease()
	return instance
}

func (u_ UnitInformationStorage) InitWithSymbol(symbol string) UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](u_, objc.Sel("initWithSymbol:"), symbol)
	return rv
}

// Initializes a new unit with the specified symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunit/1690760-initwithsymbol?language=objc
func NewUnitInformationStorageWithSymbol(symbol string) UnitInformationStorage {
	instance := UnitInformationStorageClass.Alloc().InitWithSymbol(symbol)
	instance.Autorelease()
	return instance
}

// The gibibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172551-gibibytes?language=objc
func (uc _UnitInformationStorageClass) Gibibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("gibibytes"))
	return rv
}

// The gibibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172551-gibibytes?language=objc
func UnitInformationStorage_Gibibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Gibibytes()
}

// The zebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172575-zebibits?language=objc
func (uc _UnitInformationStorageClass) Zebibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("zebibits"))
	return rv
}

// The zebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172575-zebibits?language=objc
func UnitInformationStorage_Zebibits() UnitInformationStorage {
	return UnitInformationStorageClass.Zebibits()
}

// The petabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172565-petabits?language=objc
func (uc _UnitInformationStorageClass) Petabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("petabits"))
	return rv
}

// The petabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172565-petabits?language=objc
func UnitInformationStorage_Petabits() UnitInformationStorage {
	return UnitInformationStorageClass.Petabits()
}

// The exbibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172549-exbibytes?language=objc
func (uc _UnitInformationStorageClass) Exbibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("exbibytes"))
	return rv
}

// The exbibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172549-exbibytes?language=objc
func UnitInformationStorage_Exbibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Exbibytes()
}

// The nibbles unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172562-nibbles?language=objc
func (uc _UnitInformationStorageClass) Nibbles() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("nibbles"))
	return rv
}

// The nibbles unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172562-nibbles?language=objc
func UnitInformationStorage_Nibbles() UnitInformationStorage {
	return UnitInformationStorageClass.Nibbles()
}

// The yobibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172571-yobibits?language=objc
func (uc _UnitInformationStorageClass) Yobibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("yobibits"))
	return rv
}

// The yobibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172571-yobibits?language=objc
func UnitInformationStorage_Yobibits() UnitInformationStorage {
	return UnitInformationStorageClass.Yobibits()
}

// The zettabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172577-zettabits?language=objc
func (uc _UnitInformationStorageClass) Zettabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("zettabits"))
	return rv
}

// The zettabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172577-zettabits?language=objc
func UnitInformationStorage_Zettabits() UnitInformationStorage {
	return UnitInformationStorageClass.Zettabits()
}

// The bits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172544-bits?language=objc
func (uc _UnitInformationStorageClass) Bits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("bits"))
	return rv
}

// The bits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172544-bits?language=objc
func UnitInformationStorage_Bits() UnitInformationStorage {
	return UnitInformationStorageClass.Bits()
}

// The gigabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172552-gigabits?language=objc
func (uc _UnitInformationStorageClass) Gigabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("gigabits"))
	return rv
}

// The gigabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172552-gigabits?language=objc
func UnitInformationStorage_Gigabits() UnitInformationStorage {
	return UnitInformationStorageClass.Gigabits()
}

// The kibibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172555-kibibytes?language=objc
func (uc _UnitInformationStorageClass) Kibibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("kibibytes"))
	return rv
}

// The kibibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172555-kibibytes?language=objc
func UnitInformationStorage_Kibibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Kibibytes()
}

// The mebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172559-mebibytes?language=objc
func (uc _UnitInformationStorageClass) Mebibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("mebibytes"))
	return rv
}

// The mebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172559-mebibytes?language=objc
func UnitInformationStorage_Mebibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Mebibytes()
}

// The kibibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172554-kibibits?language=objc
func (uc _UnitInformationStorageClass) Kibibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("kibibits"))
	return rv
}

// The kibibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172554-kibibits?language=objc
func UnitInformationStorage_Kibibits() UnitInformationStorage {
	return UnitInformationStorageClass.Kibibits()
}

// The terrabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172570-terabytes?language=objc
func (uc _UnitInformationStorageClass) Terabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("terabytes"))
	return rv
}

// The terrabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172570-terabytes?language=objc
func UnitInformationStorage_Terabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Terabytes()
}

// The tebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172568-tebibytes?language=objc
func (uc _UnitInformationStorageClass) Tebibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("tebibytes"))
	return rv
}

// The tebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172568-tebibytes?language=objc
func UnitInformationStorage_Tebibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Tebibytes()
}

// The pebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172563-pebibits?language=objc
func (uc _UnitInformationStorageClass) Pebibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("pebibits"))
	return rv
}

// The pebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172563-pebibits?language=objc
func UnitInformationStorage_Pebibits() UnitInformationStorage {
	return UnitInformationStorageClass.Pebibits()
}

// The yobibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172572-yobibytes?language=objc
func (uc _UnitInformationStorageClass) Yobibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("yobibytes"))
	return rv
}

// The yobibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172572-yobibytes?language=objc
func UnitInformationStorage_Yobibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Yobibytes()
}

// The petabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172566-petabytes?language=objc
func (uc _UnitInformationStorageClass) Petabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("petabytes"))
	return rv
}

// The petabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172566-petabytes?language=objc
func UnitInformationStorage_Petabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Petabytes()
}

// The gibibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172550-gibibits?language=objc
func (uc _UnitInformationStorageClass) Gibibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("gibibits"))
	return rv
}

// The gibibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172550-gibibits?language=objc
func UnitInformationStorage_Gibibits() UnitInformationStorage {
	return UnitInformationStorageClass.Gibibits()
}

// The kilobytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172557-kilobytes?language=objc
func (uc _UnitInformationStorageClass) Kilobytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("kilobytes"))
	return rv
}

// The kilobytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172557-kilobytes?language=objc
func UnitInformationStorage_Kilobytes() UnitInformationStorage {
	return UnitInformationStorageClass.Kilobytes()
}

// The exabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172547-exabytes?language=objc
func (uc _UnitInformationStorageClass) Exabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("exabytes"))
	return rv
}

// The exabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172547-exabytes?language=objc
func UnitInformationStorage_Exabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Exabytes()
}

// The megabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172561-megabytes?language=objc
func (uc _UnitInformationStorageClass) Megabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("megabytes"))
	return rv
}

// The megabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172561-megabytes?language=objc
func UnitInformationStorage_Megabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Megabytes()
}

// The mebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172558-mebibits?language=objc
func (uc _UnitInformationStorageClass) Mebibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("mebibits"))
	return rv
}

// The mebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172558-mebibits?language=objc
func UnitInformationStorage_Mebibits() UnitInformationStorage {
	return UnitInformationStorageClass.Mebibits()
}

// The pebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172564-pebibytes?language=objc
func (uc _UnitInformationStorageClass) Pebibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("pebibytes"))
	return rv
}

// The pebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172564-pebibytes?language=objc
func UnitInformationStorage_Pebibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Pebibytes()
}

// The terabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172569-terabits?language=objc
func (uc _UnitInformationStorageClass) Terabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("terabits"))
	return rv
}

// The terabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172569-terabits?language=objc
func UnitInformationStorage_Terabits() UnitInformationStorage {
	return UnitInformationStorageClass.Terabits()
}

// The bytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172545-bytes?language=objc
func (uc _UnitInformationStorageClass) Bytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("bytes"))
	return rv
}

// The bytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172545-bytes?language=objc
func UnitInformationStorage_Bytes() UnitInformationStorage {
	return UnitInformationStorageClass.Bytes()
}

// The yottabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172574-yottabytes?language=objc
func (uc _UnitInformationStorageClass) Yottabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("yottabytes"))
	return rv
}

// The yottabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172574-yottabytes?language=objc
func UnitInformationStorage_Yottabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Yottabytes()
}

// The yottabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172573-yottabits?language=objc
func (uc _UnitInformationStorageClass) Yottabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("yottabits"))
	return rv
}

// The yottabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172573-yottabits?language=objc
func UnitInformationStorage_Yottabits() UnitInformationStorage {
	return UnitInformationStorageClass.Yottabits()
}

// The gigabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172553-gigabytes?language=objc
func (uc _UnitInformationStorageClass) Gigabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("gigabytes"))
	return rv
}

// The gigabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172553-gigabytes?language=objc
func UnitInformationStorage_Gigabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Gigabytes()
}

// The zebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172576-zebibytes?language=objc
func (uc _UnitInformationStorageClass) Zebibytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("zebibytes"))
	return rv
}

// The zebibytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172576-zebibytes?language=objc
func UnitInformationStorage_Zebibytes() UnitInformationStorage {
	return UnitInformationStorageClass.Zebibytes()
}

// The kilobits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172556-kilobits?language=objc
func (uc _UnitInformationStorageClass) Kilobits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("kilobits"))
	return rv
}

// The kilobits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172556-kilobits?language=objc
func UnitInformationStorage_Kilobits() UnitInformationStorage {
	return UnitInformationStorageClass.Kilobits()
}

// The tebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172567-tebibits?language=objc
func (uc _UnitInformationStorageClass) Tebibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("tebibits"))
	return rv
}

// The tebibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172567-tebibits?language=objc
func UnitInformationStorage_Tebibits() UnitInformationStorage {
	return UnitInformationStorageClass.Tebibits()
}

// The megabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172560-megabits?language=objc
func (uc _UnitInformationStorageClass) Megabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("megabits"))
	return rv
}

// The megabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172560-megabits?language=objc
func UnitInformationStorage_Megabits() UnitInformationStorage {
	return UnitInformationStorageClass.Megabits()
}

// The exbibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172548-exbibits?language=objc
func (uc _UnitInformationStorageClass) Exbibits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("exbibits"))
	return rv
}

// The exbibits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172548-exbibits?language=objc
func UnitInformationStorage_Exbibits() UnitInformationStorage {
	return UnitInformationStorageClass.Exbibits()
}

// The exabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172546-exabits?language=objc
func (uc _UnitInformationStorageClass) Exabits() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("exabits"))
	return rv
}

// The exabits unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172546-exabits?language=objc
func UnitInformationStorage_Exabits() UnitInformationStorage {
	return UnitInformationStorageClass.Exabits()
}

// The zettabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172578-zettabytes?language=objc
func (uc _UnitInformationStorageClass) Zettabytes() UnitInformationStorage {
	rv := objc.Call[UnitInformationStorage](uc, objc.Sel("zettabytes"))
	return rv
}

// The zettabytes unit of information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsunitinformationstorage/3172578-zettabytes?language=objc
func UnitInformationStorage_Zettabytes() UnitInformationStorage {
	return UnitInformationStorageClass.Zettabytes()
}
