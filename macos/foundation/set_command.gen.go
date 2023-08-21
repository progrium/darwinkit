// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SetCommand] class.
var SetCommandClass = _SetCommandClass{objc.GetClass("NSSetCommand")}

type _SetCommandClass struct {
	objc.Class
}

// An interface definition for the [SetCommand] class.
type ISetCommand interface {
	IScriptCommand
	KeySpecifier() ScriptObjectSpecifier
}

// A command that sets one or more attributes or relationships to one or more values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssetcommand?language=objc
type SetCommand struct {
	ScriptCommand
}

func SetCommandFrom(ptr unsafe.Pointer) SetCommand {
	return SetCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (sc _SetCommandClass) Alloc() SetCommand {
	rv := objc.Call[SetCommand](sc, objc.Sel("alloc"))
	return rv
}

func SetCommand_Alloc() SetCommand {
	return SetCommandClass.Alloc()
}

func (sc _SetCommandClass) New() SetCommand {
	rv := objc.Call[SetCommand](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSetCommand() SetCommand {
	return SetCommandClass.New()
}

func (s_ SetCommand) Init() SetCommand {
	rv := objc.Call[SetCommand](s_, objc.Sel("init"))
	return rv
}

func (s_ SetCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) SetCommand {
	rv := objc.Call[SetCommand](s_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewSetCommandWithCommandDescription(commandDef IScriptCommandDescription) SetCommand {
	instance := SetCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}

// Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the set AppleScript command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssetcommand/1415804-keyspecifier?language=objc
func (s_ SetCommand) KeySpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("keySpecifier"))
	return rv
}
