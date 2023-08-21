// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CreateCommand] class.
var CreateCommandClass = _CreateCommandClass{objc.GetClass("NSCreateCommand")}

type _CreateCommandClass struct {
	objc.Class
}

// An interface definition for the [CreateCommand] class.
type ICreateCommand interface {
	IScriptCommand
	ResolvedKeyDictionary() map[string]objc.Object
	CreateClassDescription() ScriptClassDescription
}

// A command that creates a scriptable object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscreatecommand?language=objc
type CreateCommand struct {
	ScriptCommand
}

func CreateCommandFrom(ptr unsafe.Pointer) CreateCommand {
	return CreateCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (cc _CreateCommandClass) Alloc() CreateCommand {
	rv := objc.Call[CreateCommand](cc, objc.Sel("alloc"))
	return rv
}

func CreateCommand_Alloc() CreateCommand {
	return CreateCommandClass.Alloc()
}

func (cc _CreateCommandClass) New() CreateCommand {
	rv := objc.Call[CreateCommand](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCreateCommand() CreateCommand {
	return CreateCommandClass.New()
}

func (c_ CreateCommand) Init() CreateCommand {
	rv := objc.Call[CreateCommand](c_, objc.Sel("init"))
	return rv
}

func (c_ CreateCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) CreateCommand {
	rv := objc.Call[CreateCommand](c_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewCreateCommandWithCommandDescription(commandDef IScriptCommandDescription) CreateCommand {
	instance := CreateCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}

// Returns a dictionary that contains the properties that were specified in the make Apple event command that has been converted to this NSCreateCommand object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscreatecommand/1407639-resolvedkeydictionary?language=objc
func (c_ CreateCommand) ResolvedKeyDictionary() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("resolvedKeyDictionary"))
	return rv
}

// Returns the class description for the class that is to be created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscreatecommand/1413533-createclassdescription?language=objc
func (c_ CreateCommand) CreateClassDescription() ScriptClassDescription {
	rv := objc.Call[ScriptClassDescription](c_, objc.Sel("createClassDescription"))
	return rv
}
