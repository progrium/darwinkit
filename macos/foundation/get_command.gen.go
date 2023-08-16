// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GetCommand] class.
var GetCommandClass = _GetCommandClass{objc.GetClass("NSGetCommand")}

type _GetCommandClass struct {
	objc.Class
}

// An interface definition for the [GetCommand] class.
type IGetCommand interface {
	IScriptCommand
}

// A command that retrieves a value or object from a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsgetcommand?language=objc
type GetCommand struct {
	ScriptCommand
}

func GetCommandFrom(ptr unsafe.Pointer) GetCommand {
	return GetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (gc _GetCommandClass) Alloc() GetCommand {
	rv := objc.Call[GetCommand](gc, objc.Sel("alloc"))
	return rv
}

func GetCommand_Alloc() GetCommand {
	return GetCommandClass.Alloc()
}

func (gc _GetCommandClass) New() GetCommand {
	rv := objc.Call[GetCommand](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGetCommand() GetCommand {
	return GetCommandClass.New()
}

func (g_ GetCommand) Init() GetCommand {
	rv := objc.Call[GetCommand](g_, objc.Sel("init"))
	return rv
}

func (g_ GetCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) GetCommand {
	rv := objc.Call[GetCommand](g_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func GetCommand_InitWithCommandDescription(commandDef IScriptCommandDescription) GetCommand {
	return GetCommandClass.Alloc().InitWithCommandDescription(commandDef)
}
