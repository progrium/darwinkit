// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MoveCommand] class.
var MoveCommandClass = _MoveCommandClass{objc.GetClass("NSMoveCommand")}

type _MoveCommandClass struct {
	objc.Class
}

// An interface definition for the [MoveCommand] class.
type IMoveCommand interface {
	IScriptCommand
	KeySpecifier() ScriptObjectSpecifier
}

// A command that moves one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmovecommand?language=objc
type MoveCommand struct {
	ScriptCommand
}

func MoveCommandFrom(ptr unsafe.Pointer) MoveCommand {
	return MoveCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (mc _MoveCommandClass) Alloc() MoveCommand {
	rv := objc.Call[MoveCommand](mc, objc.Sel("alloc"))
	return rv
}

func MoveCommand_Alloc() MoveCommand {
	return MoveCommandClass.Alloc()
}

func (mc _MoveCommandClass) New() MoveCommand {
	rv := objc.Call[MoveCommand](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMoveCommand() MoveCommand {
	return MoveCommandClass.New()
}

func (m_ MoveCommand) Init() MoveCommand {
	rv := objc.Call[MoveCommand](m_, objc.Sel("init"))
	return rv
}

func (m_ MoveCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) MoveCommand {
	rv := objc.Call[MoveCommand](m_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func MoveCommand_InitWithCommandDescription(commandDef IScriptCommandDescription) MoveCommand {
	return MoveCommandClass.Alloc().InitWithCommandDescription(commandDef)
}

// Returns a specifier for the object or objects to be moved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmovecommand/1413005-keyspecifier?language=objc
func (m_ MoveCommand) KeySpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](m_, objc.Sel("keySpecifier"))
	return rv
}
