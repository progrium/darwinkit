// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by the delegate of an NSAlert object to respond to a user's request for help. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalertdelegate?language=objc
type PAlertDelegate interface {
	// optional
	AlertShowHelp(alert Alert) bool
	HasAlertShowHelp() bool
}

// A delegate implementation builder for the [PAlertDelegate] protocol.
type AlertDelegate struct {
	_AlertShowHelp func(alert Alert) bool
}

func (di *AlertDelegate) HasAlertShowHelp() bool {
	return di._AlertShowHelp != nil
}

// Sent to the delegate when the user clicks the alert’s help button. The delegate causes help to be displayed for an alert, directly or indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalertdelegate/1526980-alertshowhelp?language=objc
func (di *AlertDelegate) SetAlertShowHelp(f func(alert Alert) bool) {
	di._AlertShowHelp = f
}

// Sent to the delegate when the user clicks the alert’s help button. The delegate causes help to be displayed for an alert, directly or indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalertdelegate/1526980-alertshowhelp?language=objc
func (di *AlertDelegate) AlertShowHelp(alert Alert) bool {
	return di._AlertShowHelp(alert)
}

// A concrete type wrapper for the [PAlertDelegate] protocol.
type AlertDelegateWrapper struct {
	objc.Object
}

func (a_ AlertDelegateWrapper) HasAlertShowHelp() bool {
	return a_.RespondsToSelector(objc.Sel("alertShowHelp:"))
}

// Sent to the delegate when the user clicks the alert’s help button. The delegate causes help to be displayed for an alert, directly or indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsalertdelegate/1526980-alertshowhelp?language=objc
func (a_ AlertDelegateWrapper) AlertShowHelp(alert IAlert) bool {
	rv := objc.Call[bool](a_, objc.Sel("alertShowHelp:"), objc.Ptr(alert))
	return rv
}
