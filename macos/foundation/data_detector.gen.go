// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DataDetector] class.
var DataDetectorClass = _DataDetectorClass{objc.GetClass("NSDataDetector")}

type _DataDetectorClass struct {
	objc.Class
}

// An interface definition for the [DataDetector] class.
type IDataDetector interface {
	IRegularExpression
	CheckingTypes() TextCheckingTypes
}

// A specialized regular expression object that matches natural language text for predefined data patterns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector?language=objc
type DataDetector struct {
	RegularExpression
}

func DataDetectorFrom(ptr unsafe.Pointer) DataDetector {
	return DataDetector{
		RegularExpression: RegularExpressionFrom(ptr),
	}
}

func (d_ DataDetector) InitWithTypesError(checkingTypes TextCheckingTypes, error IError) DataDetector {
	rv := objc.Call[DataDetector](d_, objc.Sel("initWithTypes:error:"), checkingTypes, objc.Ptr(error))
	return rv
}

// Initializes and returns a data detector instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/1409829-initwithtypes?language=objc
func NewDataDetectorWithTypesError(checkingTypes TextCheckingTypes, error IError) DataDetector {
	instance := DataDetectorClass.Alloc().InitWithTypesError(checkingTypes, error)
	instance.Autorelease()
	return instance
}

func (dc _DataDetectorClass) Alloc() DataDetector {
	rv := objc.Call[DataDetector](dc, objc.Sel("alloc"))
	return rv
}

func DataDetector_Alloc() DataDetector {
	return DataDetectorClass.Alloc()
}

func (dc _DataDetectorClass) New() DataDetector {
	rv := objc.Call[DataDetector](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDataDetector() DataDetector {
	return DataDetectorClass.New()
}

func (d_ DataDetector) Init() DataDetector {
	rv := objc.Call[DataDetector](d_, objc.Sel("init"))
	return rv
}

func (d_ DataDetector) InitWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) DataDetector {
	rv := objc.Call[DataDetector](d_, objc.Sel("initWithPattern:options:error:"), pattern, options, objc.Ptr(error))
	return rv
}

// Returns an initialized NSRegularExpression instance with the specified regular expression pattern and options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsregularexpression/1410900-initwithpattern?language=objc
func NewDataDetectorWithPatternOptionsError(pattern string, options RegularExpressionOptions, error IError) DataDetector {
	instance := DataDetectorClass.Alloc().InitWithPatternOptionsError(pattern, options, error)
	instance.Autorelease()
	return instance
}

// Creates and returns a new data detector instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/1557375-datadetectorwithtypes?language=objc
func (dc _DataDetectorClass) DataDetectorWithTypesError(checkingTypes TextCheckingTypes, error IError) DataDetector {
	rv := objc.Call[DataDetector](dc, objc.Sel("dataDetectorWithTypes:error:"), checkingTypes, objc.Ptr(error))
	return rv
}

// Creates and returns a new data detector instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/1557375-datadetectorwithtypes?language=objc
func DataDetector_DataDetectorWithTypesError(checkingTypes TextCheckingTypes, error IError) DataDetector {
	return DataDetectorClass.DataDetectorWithTypesError(checkingTypes, error)
}

// Returns the checking types for the data detector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdatadetector/1412372-checkingtypes?language=objc
func (d_ DataDetector) CheckingTypes() TextCheckingTypes {
	rv := objc.Call[TextCheckingTypes](d_, objc.Sel("checkingTypes"))
	return rv
}
