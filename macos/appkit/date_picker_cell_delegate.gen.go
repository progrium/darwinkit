// AUTO-GENERATED CODE, DO NOT MODIFY
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IDatePickerCellDelegate interface {
	ImplementsDatePickerCellValidateProposedDateValueTimeInterval() bool
	// optional
	DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

type DatePickerCellDelegate struct {
	_DatePickerCellValidateProposedDateValueTimeInterval func(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

func (di *DatePickerCellDelegate) ImplementsDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return di._DatePickerCellValidateProposedDateValueTimeInterval != nil
}

func (di *DatePickerCellDelegate) SetDatePickerCellValidateProposedDateValueTimeInterval(f func(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval)) {
	di._DatePickerCellValidateProposedDateValueTimeInterval = f
}

func (di *DatePickerCellDelegate) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell DatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	di._DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell, proposedDateValue, proposedTimeInterval)
}

type DatePickerCellDelegateWrapper struct {
	objc.Object
}

func (d_ DatePickerCellDelegateWrapper) ImplementsDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return d_.RespondsToSelector(objc.GetSelector("datePickerCell:validateProposedDateValue:timeInterval:"))
}

func (d_ DatePickerCellDelegateWrapper) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue *foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("datePickerCell:validateProposedDateValue:timeInterval:"), objc.ExtractPtr(datePickerCell), unsafe.Pointer(proposedDateValue), proposedTimeInterval)
}
