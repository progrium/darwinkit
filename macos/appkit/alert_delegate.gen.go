// auto-generated code, do not modify
package appkit

import (
	"github.com/progrium/macdriver/objc"
)

type IAlertDelegate interface {
	ImplementsAlertShowHelp() bool
	// optional
	AlertShowHelp(alert Alert) bool
}

type AlertDelegate struct {
	_AlertShowHelp func(alert Alert) bool
}

func (di *AlertDelegate) ImplementsAlertShowHelp() bool {
	return di._AlertShowHelp != nil
}

func (di *AlertDelegate) SetAlertShowHelp(f func(alert Alert) bool) {
	di._AlertShowHelp = f
}

func (di *AlertDelegate) AlertShowHelp(alert Alert) bool {
	return di._AlertShowHelp(alert)
}

type AlertDelegateWrapper struct {
	objc.Object
}

func (a_ AlertDelegateWrapper) ImplementsAlertShowHelp() bool {
	return a_.RespondsToSelector(objc.GetSelector("alertShowHelp:"))
}

func (a_ AlertDelegateWrapper) AlertShowHelp(alert IAlert) bool {
	rv := objc.CallMethod[bool](a_, objc.GetSelector("alertShowHelp:"), objc.ExtractPtr(alert))
	return rv
}
