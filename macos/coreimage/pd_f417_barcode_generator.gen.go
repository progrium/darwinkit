// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a PDF417 barcode generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator?language=objc
type PPDF417BarcodeGenerator interface {
	// optional
	SetMinHeight(value float64)
	HasSetMinHeight() bool

	// optional
	MinHeight() float64
	HasMinHeight() bool

	// optional
	SetDataColumns(value float64)
	HasSetDataColumns() bool

	// optional
	DataColumns() float64
	HasDataColumns() bool

	// optional
	SetMinWidth(value float64)
	HasSetMinWidth() bool

	// optional
	MinWidth() float64
	HasMinWidth() bool

	// optional
	SetMaxHeight(value float64)
	HasSetMaxHeight() bool

	// optional
	MaxHeight() float64
	HasMaxHeight() bool

	// optional
	SetMaxWidth(value float64)
	HasSetMaxWidth() bool

	// optional
	MaxWidth() float64
	HasMaxWidth() bool

	// optional
	SetRows(value float64)
	HasSetRows() bool

	// optional
	Rows() float64
	HasRows() bool

	// optional
	SetPreferredAspectRatio(value float64)
	HasSetPreferredAspectRatio() bool

	// optional
	PreferredAspectRatio() float64
	HasPreferredAspectRatio() bool

	// optional
	SetAlwaysSpecifyCompaction(value float64)
	HasSetAlwaysSpecifyCompaction() bool

	// optional
	AlwaysSpecifyCompaction() float64
	HasAlwaysSpecifyCompaction() bool

	// optional
	SetCompactionMode(value float64)
	HasSetCompactionMode() bool

	// optional
	CompactionMode() float64
	HasCompactionMode() bool

	// optional
	SetCompactStyle(value float64)
	HasSetCompactStyle() bool

	// optional
	CompactStyle() float64
	HasCompactStyle() bool

	// optional
	SetMessage(value []byte)
	HasSetMessage() bool

	// optional
	Message() []byte
	HasMessage() bool

	// optional
	SetCorrectionLevel(value float64)
	HasSetCorrectionLevel() bool

	// optional
	CorrectionLevel() float64
	HasCorrectionLevel() bool
}

// A concrete type wrapper for the [PPDF417BarcodeGenerator] protocol.
type PDF417BarcodeGeneratorWrapper struct {
	objc.Object
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetMinHeight() bool {
	return p_.RespondsToSelector(objc.Sel("setMinHeight:"))
}

// The minimum height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228613-minheight?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetMinHeight(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMinHeight:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasMinHeight() bool {
	return p_.RespondsToSelector(objc.Sel("minHeight"))
}

// The minimum height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228613-minheight?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) MinHeight() float64 {
	rv := objc.Call[float64](p_, objc.Sel("minHeight"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetDataColumns() bool {
	return p_.RespondsToSelector(objc.Sel("setDataColumns:"))
}

// The number of data columns in the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228609-datacolumns?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetDataColumns(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setDataColumns:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasDataColumns() bool {
	return p_.RespondsToSelector(objc.Sel("dataColumns"))
}

// The number of data columns in the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228609-datacolumns?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) DataColumns() float64 {
	rv := objc.Call[float64](p_, objc.Sel("dataColumns"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetMinWidth() bool {
	return p_.RespondsToSelector(objc.Sel("setMinWidth:"))
}

// The minimum width, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228614-minwidth?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetMinWidth(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMinWidth:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasMinWidth() bool {
	return p_.RespondsToSelector(objc.Sel("minWidth"))
}

// The minimum width, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228614-minwidth?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) MinWidth() float64 {
	rv := objc.Call[float64](p_, objc.Sel("minWidth"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetMaxHeight() bool {
	return p_.RespondsToSelector(objc.Sel("setMaxHeight:"))
}

// The maximum height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228610-maxheight?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetMaxHeight(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMaxHeight:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasMaxHeight() bool {
	return p_.RespondsToSelector(objc.Sel("maxHeight"))
}

// The maximum height, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228610-maxheight?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) MaxHeight() float64 {
	rv := objc.Call[float64](p_, objc.Sel("maxHeight"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetMaxWidth() bool {
	return p_.RespondsToSelector(objc.Sel("setMaxWidth:"))
}

// The maximum width, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228611-maxwidth?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetMaxWidth(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setMaxWidth:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasMaxWidth() bool {
	return p_.RespondsToSelector(objc.Sel("maxWidth"))
}

// The maximum width, in pixels, of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228611-maxwidth?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) MaxWidth() float64 {
	rv := objc.Call[float64](p_, objc.Sel("maxWidth"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetRows() bool {
	return p_.RespondsToSelector(objc.Sel("setRows:"))
}

// The number of rows in the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228616-rows?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetRows(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRows:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasRows() bool {
	return p_.RespondsToSelector(objc.Sel("rows"))
}

// The number of rows in the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228616-rows?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) Rows() float64 {
	rv := objc.Call[float64](p_, objc.Sel("rows"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetPreferredAspectRatio() bool {
	return p_.RespondsToSelector(objc.Sel("setPreferredAspectRatio:"))
}

// The preferred aspect ratio of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228615-preferredaspectratio?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetPreferredAspectRatio(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setPreferredAspectRatio:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasPreferredAspectRatio() bool {
	return p_.RespondsToSelector(objc.Sel("preferredAspectRatio"))
}

// The preferred aspect ratio of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228615-preferredaspectratio?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) PreferredAspectRatio() float64 {
	rv := objc.Call[float64](p_, objc.Sel("preferredAspectRatio"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetAlwaysSpecifyCompaction() bool {
	return p_.RespondsToSelector(objc.Sel("setAlwaysSpecifyCompaction:"))
}

// A Boolean value specifying whether to force compaction style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228605-alwaysspecifycompaction?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetAlwaysSpecifyCompaction(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAlwaysSpecifyCompaction:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasAlwaysSpecifyCompaction() bool {
	return p_.RespondsToSelector(objc.Sel("alwaysSpecifyCompaction"))
}

// A Boolean value specifying whether to force compaction style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228605-alwaysspecifycompaction?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) AlwaysSpecifyCompaction() float64 {
	rv := objc.Call[float64](p_, objc.Sel("alwaysSpecifyCompaction"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetCompactionMode() bool {
	return p_.RespondsToSelector(objc.Sel("setCompactionMode:"))
}

// The compaction mode of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228607-compactionmode?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetCompactionMode(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setCompactionMode:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasCompactionMode() bool {
	return p_.RespondsToSelector(objc.Sel("compactionMode"))
}

// The compaction mode of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228607-compactionmode?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) CompactionMode() float64 {
	rv := objc.Call[float64](p_, objc.Sel("compactionMode"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetCompactStyle() bool {
	return p_.RespondsToSelector(objc.Sel("setCompactStyle:"))
}

// A Boolean value specifying whether to force compact style Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228606-compactstyle?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetCompactStyle(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setCompactStyle:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasCompactStyle() bool {
	return p_.RespondsToSelector(objc.Sel("compactStyle"))
}

// A Boolean value specifying whether to force compact style Aztec code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228606-compactstyle?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) CompactStyle() float64 {
	rv := objc.Call[float64](p_, objc.Sel("compactStyle"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetMessage() bool {
	return p_.RespondsToSelector(objc.Sel("setMessage:"))
}

// The message to encode in the PDF417 barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228612-message?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetMessage(value []byte) {
	objc.Call[objc.Void](p_, objc.Sel("setMessage:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasMessage() bool {
	return p_.RespondsToSelector(objc.Sel("message"))
}

// The message to encode in the PDF417 barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228612-message?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) Message() []byte {
	rv := objc.Call[[]byte](p_, objc.Sel("message"))
	return rv
}

func (p_ PDF417BarcodeGeneratorWrapper) HasSetCorrectionLevel() bool {
	return p_.RespondsToSelector(objc.Sel("setCorrectionLevel:"))
}

// The correction level ratio of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228608-correctionlevel?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) SetCorrectionLevel(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setCorrectionLevel:"), value)
}

func (p_ PDF417BarcodeGeneratorWrapper) HasCorrectionLevel() bool {
	return p_.RespondsToSelector(objc.Sel("correctionLevel"))
}

// The correction level ratio of the generated barcode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipdf417barcodegenerator/3228608-correctionlevel?language=objc
func (p_ PDF417BarcodeGeneratorWrapper) CorrectionLevel() float64 {
	rv := objc.Call[float64](p_, objc.Sel("correctionLevel"))
	return rv
}
