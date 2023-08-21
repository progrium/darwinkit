// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CloseCommand] class.
var CloseCommandClass = _CloseCommandClass{objc.GetClass("NSCloseCommand")}

type _CloseCommandClass struct {
	objc.Class
}

// An interface definition for the [CloseCommand] class.
type ICloseCommand interface {
	IScriptCommand
	SaveOptions() SaveOptions
}

// A command that closes one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclosecommand?language=objc
type CloseCommand struct {
	ScriptCommand
}

func CloseCommandFrom(ptr unsafe.Pointer) CloseCommand {
	return CloseCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (cc _CloseCommandClass) Alloc() CloseCommand {
	rv := objc.Call[CloseCommand](cc, objc.Sel("alloc"))
	return rv
}

func CloseCommand_Alloc() CloseCommand {
	return CloseCommandClass.Alloc()
}

func (cc _CloseCommandClass) New() CloseCommand {
	rv := objc.Call[CloseCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCloseCommand() CloseCommand {
	return CloseCommandClass.New()
}

func (c_ CloseCommand) Init() CloseCommand {
	rv := objc.Call[CloseCommand](c_, objc.Sel("init"))
	return rv
}

func (c_ CloseCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) CloseCommand {
	rv := objc.Call[CloseCommand](c_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewCloseCommandWithCommandDescription(commandDef IScriptCommandDescription) CloseCommand {
	instance := CloseCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}

// Returns a constant indicating how to deal with closing any modified documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclosecommand/1415647-saveoptions?language=objc
func (c_ CloseCommand) SaveOptions() SaveOptions {
	rv := objc.Call[SaveOptions](c_, objc.Sel("saveOptions"))
	return rv
}
