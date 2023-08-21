// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DeleteCommand] class.
var DeleteCommandClass = _DeleteCommandClass{objc.GetClass("NSDeleteCommand")}

type _DeleteCommandClass struct {
	objc.Class
}

// An interface definition for the [DeleteCommand] class.
type IDeleteCommand interface {
	IScriptCommand
	KeySpecifier() ScriptObjectSpecifier
}

// A command that deletes a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdeletecommand?language=objc
type DeleteCommand struct {
	ScriptCommand
}

func DeleteCommandFrom(ptr unsafe.Pointer) DeleteCommand {
	return DeleteCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (dc _DeleteCommandClass) Alloc() DeleteCommand {
	rv := objc.Call[DeleteCommand](dc, objc.Sel("alloc"))
	return rv
}

func DeleteCommand_Alloc() DeleteCommand {
	return DeleteCommandClass.Alloc()
}

func (dc _DeleteCommandClass) New() DeleteCommand {
	rv := objc.Call[DeleteCommand](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDeleteCommand() DeleteCommand {
	return DeleteCommandClass.New()
}

func (d_ DeleteCommand) Init() DeleteCommand {
	rv := objc.Call[DeleteCommand](d_, objc.Sel("init"))
	return rv
}

func (d_ DeleteCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) DeleteCommand {
	rv := objc.Call[DeleteCommand](d_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewDeleteCommandWithCommandDescription(commandDef IScriptCommandDescription) DeleteCommand {
	instance := DeleteCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}

// Returns a specifier for the object or objects to be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsdeletecommand/1414705-keyspecifier?language=objc
func (d_ DeleteCommand) KeySpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](d_, objc.Sel("keySpecifier"))
	return rv
}
