// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RemoteCommand] class.
var RemoteCommandClass = _RemoteCommandClass{objc.GetClass("MPRemoteCommand")}

type _RemoteCommandClass struct {
	objc.Class
}

// An interface definition for the [RemoteCommand] class.
type IRemoteCommand interface {
	objc.IObject
	AddTargetWithHandler(handler func(event RemoteCommandEvent) RemoteCommandHandlerStatus) objc.Object
	RemoveTarget(target objc.IObject)
	AddTargetAction(target objc.IObject, action objc.Selector)
	IsEnabled() bool
	SetEnabled(value bool)
}

// An object that responds to remote command events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand?language=objc
type RemoteCommand struct {
	objc.Object
}

func RemoteCommandFrom(ptr unsafe.Pointer) RemoteCommand {
	return RemoteCommand{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RemoteCommandClass) Alloc() RemoteCommand {
	rv := objc.Call[RemoteCommand](rc, objc.Sel("alloc"))
	return rv
}

func RemoteCommand_Alloc() RemoteCommand {
	return RemoteCommandClass.Alloc()
}

func (rc _RemoteCommandClass) New() RemoteCommand {
	rv := objc.Call[RemoteCommand](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRemoteCommand() RemoteCommand {
	return RemoteCommandClass.New()
}

func (r_ RemoteCommand) Init() RemoteCommand {
	rv := objc.Call[RemoteCommand](r_, objc.Sel("init"))
	return rv
}

// Adds a block to be called when an event is received. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/1622910-addtargetwithhandler?language=objc
func (r_ RemoteCommand) AddTargetWithHandler(handler func(event RemoteCommandEvent) RemoteCommandHandlerStatus) objc.Object {
	rv := objc.Call[objc.Object](r_, objc.Sel("addTargetWithHandler:"), handler)
	return rv
}

// Removes a target from the remote command object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/1622903-removetarget?language=objc
func (r_ RemoteCommand) RemoveTarget(target objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("removeTarget:"), target)
}

// Adds a target object to be called when an event is received. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/1622895-addtarget?language=objc
func (r_ RemoteCommand) AddTargetAction(target objc.IObject, action objc.Selector) {
	objc.Call[objc.Void](r_, objc.Sel("addTarget:action:"), target, action)
}

// A Boolean value that indicates whether a user can interact with the displayed element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/1622908-enabled?language=objc
func (r_ RemoteCommand) IsEnabled() bool {
	rv := objc.Call[bool](r_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether a user can interact with the displayed element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpremotecommand/1622908-enabled?language=objc
func (r_ RemoteCommand) SetEnabled(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setEnabled:"), value)
}
