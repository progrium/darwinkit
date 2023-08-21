// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CountCommand] class.
var CountCommandClass = _CountCommandClass{objc.GetClass("NSCountCommand")}

type _CountCommandClass struct {
	objc.Class
}

// An interface definition for the [CountCommand] class.
type ICountCommand interface {
	IScriptCommand
}

// A command that counts the number of objects of a specified class in the specified object container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscountcommand?language=objc
type CountCommand struct {
	ScriptCommand
}

func CountCommandFrom(ptr unsafe.Pointer) CountCommand {
	return CountCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (cc _CountCommandClass) Alloc() CountCommand {
	rv := objc.Call[CountCommand](cc, objc.Sel("alloc"))
	return rv
}

func CountCommand_Alloc() CountCommand {
	return CountCommandClass.Alloc()
}

func (cc _CountCommandClass) New() CountCommand {
	rv := objc.Call[CountCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCountCommand() CountCommand {
	return CountCommandClass.New()
}

func (c_ CountCommand) Init() CountCommand {
	rv := objc.Call[CountCommand](c_, objc.Sel("init"))
	return rv
}

func (c_ CountCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) CountCommand {
	rv := objc.Call[CountCommand](c_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewCountCommandWithCommandDescription(commandDef IScriptCommandDescription) CountCommand {
	instance := CountCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}
