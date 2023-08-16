// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSDatePickerCell objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate?language=objc
type PDatePickerCellDelegate interface {
	// optional
	DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell DatePickerCell, proposedDateValue foundation.Date, proposedTimeInterval *foundation.TimeInterval)
	HasDatePickerCellValidateProposedDateValueTimeInterval() bool
}

// A delegate implementation builder for the [PDatePickerCellDelegate] protocol.
type DatePickerCellDelegate struct {
	_DatePickerCellValidateProposedDateValueTimeInterval func(datePickerCell DatePickerCell, proposedDateValue foundation.Date, proposedTimeInterval *foundation.TimeInterval)
}

func (di *DatePickerCellDelegate) HasDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return di._DatePickerCellValidateProposedDateValueTimeInterval != nil
}

// The delegate receives this message each time the user attempts to change the receiver‘s value, allowing the delegate the opportunity to override the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate/1459631-datepickercell?language=objc
func (di *DatePickerCellDelegate) SetDatePickerCellValidateProposedDateValueTimeInterval(f func(datePickerCell DatePickerCell, proposedDateValue foundation.Date, proposedTimeInterval *foundation.TimeInterval)) {
	di._DatePickerCellValidateProposedDateValueTimeInterval = f
}

// The delegate receives this message each time the user attempts to change the receiver‘s value, allowing the delegate the opportunity to override the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate/1459631-datepickercell?language=objc
func (di *DatePickerCellDelegate) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell DatePickerCell, proposedDateValue foundation.Date, proposedTimeInterval *foundation.TimeInterval) {
	di._DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell, proposedDateValue, proposedTimeInterval)
}

// A concrete type wrapper for the [PDatePickerCellDelegate] protocol.
type DatePickerCellDelegateWrapper struct {
	objc.Object
}

func (d_ DatePickerCellDelegateWrapper) HasDatePickerCellValidateProposedDateValueTimeInterval() bool {
	return d_.RespondsToSelector(objc.Sel("datePickerCell:validateProposedDateValue:timeInterval:"))
}

// The delegate receives this message each time the user attempts to change the receiver‘s value, allowing the delegate the opportunity to override the change. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdatepickercelldelegate/1459631-datepickercell?language=objc
func (d_ DatePickerCellDelegateWrapper) DatePickerCellValidateProposedDateValueTimeInterval(datePickerCell IDatePickerCell, proposedDateValue foundation.IDate, proposedTimeInterval *foundation.TimeInterval) {
	objc.Call[objc.Void](d_, objc.Sel("datePickerCell:validateProposedDateValue:timeInterval:"), objc.Ptr(datePickerCell), objc.Ptr(proposedDateValue), proposedTimeInterval)
}
