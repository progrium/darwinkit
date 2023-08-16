// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CloneCommand] class.
var CloneCommandClass = _CloneCommandClass{objc.GetClass("NSCloneCommand")}

type _CloneCommandClass struct {
	objc.Class
}

// An interface definition for the [CloneCommand] class.
type ICloneCommand interface {
	IScriptCommand
	KeySpecifier() ScriptObjectSpecifier
}

// A command that clones one or more scriptable objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclonecommand?language=objc
type CloneCommand struct {
	ScriptCommand
}

func CloneCommandFrom(ptr unsafe.Pointer) CloneCommand {
	return CloneCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (cc _CloneCommandClass) Alloc() CloneCommand {
	rv := objc.Call[CloneCommand](cc, objc.Sel("alloc"))
	return rv
}

func CloneCommand_Alloc() CloneCommand {
	return CloneCommandClass.Alloc()
}

func (cc _CloneCommandClass) New() CloneCommand {
	rv := objc.Call[CloneCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCloneCommand() CloneCommand {
	return CloneCommandClass.New()
}

func (c_ CloneCommand) Init() CloneCommand {
	rv := objc.Call[CloneCommand](c_, objc.Sel("init"))
	return rv
}

func (c_ CloneCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) CloneCommand {
	rv := objc.Call[CloneCommand](c_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func CloneCommand_InitWithCommandDescription(commandDef IScriptCommandDescription) CloneCommand {
	return CloneCommandClass.Alloc().InitWithCommandDescription(commandDef)
}

// Returns a specifier for the object or objects to be cloned. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsclonecommand/1407603-keyspecifier?language=objc
func (c_ CloneCommand) KeySpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](c_, objc.Sel("keySpecifier"))
	return rv
}
