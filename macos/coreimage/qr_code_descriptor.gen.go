// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QRCodeDescriptor] class.
var QRCodeDescriptorClass = _QRCodeDescriptorClass{objc.GetClass("CIQRCodeDescriptor")}

type _QRCodeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [QRCodeDescriptor] class.
type IQRCodeDescriptor interface {
	IBarcodeDescriptor
	SymbolVersion() int
	MaskPattern() uint8
	ErrorCorrectionLevel() QRCodeErrorCorrectionLevel
	ErrorCorrectedPayload() []byte
}

// A concrete subclass of CIBarcodeDescriptor that represents a square QR code symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor?language=objc
type QRCodeDescriptor struct {
	BarcodeDescriptor
}

func QRCodeDescriptorFrom(ptr unsafe.Pointer) QRCodeDescriptor {
	return QRCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

func (q_ QRCodeDescriptor) InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload []byte, symbolVersion int, maskPattern uint8, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](q_, objc.Sel("initWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}

// Initializes a descriptor that can be used as input to the CIBarcodeGenerator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875180-initwithpayload?language=objc
func QRCodeDescriptor_InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload []byte, symbolVersion int, maskPattern uint8, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	return QRCodeDescriptorClass.Alloc().InitWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
}

func (qc _QRCodeDescriptorClass) DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload []byte, symbolVersion int, maskPattern uint8, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](qc, objc.Sel("descriptorWithPayload:symbolVersion:maskPattern:errorCorrectionLevel:"), errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
	return rv
}

// Creates a QR code descriptor encoding the given payload and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875169-descriptorwithpayload?language=objc
func QRCodeDescriptor_DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload []byte, symbolVersion int, maskPattern uint8, errorCorrectionLevel QRCodeErrorCorrectionLevel) QRCodeDescriptor {
	return QRCodeDescriptorClass.DescriptorWithPayloadSymbolVersionMaskPatternErrorCorrectionLevel(errorCorrectedPayload, symbolVersion, maskPattern, errorCorrectionLevel)
}

func (qc _QRCodeDescriptorClass) Alloc() QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](qc, objc.Sel("alloc"))
	return rv
}

func QRCodeDescriptor_Alloc() QRCodeDescriptor {
	return QRCodeDescriptorClass.Alloc()
}

func (qc _QRCodeDescriptorClass) New() QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQRCodeDescriptor() QRCodeDescriptor {
	return QRCodeDescriptorClass.New()
}

func (q_ QRCodeDescriptor) Init() QRCodeDescriptor {
	rv := objc.Call[QRCodeDescriptor](q_, objc.Sel("init"))
	return rv
}

// The version of the QR code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875193-symbolversion?language=objc
func (q_ QRCodeDescriptor) SymbolVersion() int {
	rv := objc.Call[int](q_, objc.Sel("symbolVersion"))
	return rv
}

// The QR code's mask pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875191-maskpattern?language=objc
func (q_ QRCodeDescriptor) MaskPattern() uint8 {
	rv := objc.Call[uint8](q_, objc.Sel("maskPattern"))
	return rv
}

// The QR code error correction level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875196-errorcorrectionlevel?language=objc
func (q_ QRCodeDescriptor) ErrorCorrectionLevel() QRCodeErrorCorrectionLevel {
	rv := objc.Call[QRCodeErrorCorrectionLevel](q_, objc.Sel("errorCorrectionLevel"))
	return rv
}

// The error-corrected payload containing the data encoded in the QR code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodedescriptor/2875167-errorcorrectedpayload?language=objc
func (q_ QRCodeDescriptor) ErrorCorrectedPayload() []byte {
	rv := objc.Call[[]byte](q_, objc.Sel("errorCorrectedPayload"))
	return rv
}
