// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuitCommand] class.
var QuitCommandClass = _QuitCommandClass{objc.GetClass("NSQuitCommand")}

type _QuitCommandClass struct {
	objc.Class
}

// An interface definition for the [QuitCommand] class.
type IQuitCommand interface {
	IScriptCommand
	SaveOptions() SaveOptions
}

// A command that quits the specified app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsquitcommand?language=objc
type QuitCommand struct {
	ScriptCommand
}

func QuitCommandFrom(ptr unsafe.Pointer) QuitCommand {
	return QuitCommand{
		ScriptCommand: ScriptCommandFrom(ptr),
	}
}

func (qc _QuitCommandClass) Alloc() QuitCommand {
	rv := objc.Call[QuitCommand](qc, objc.Sel("alloc"))
	return rv
}

func QuitCommand_Alloc() QuitCommand {
	return QuitCommandClass.Alloc()
}

func (qc _QuitCommandClass) New() QuitCommand {
	rv := objc.Call[QuitCommand](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuitCommand() QuitCommand {
	return QuitCommandClass.New()
}

func (q_ QuitCommand) Init() QuitCommand {
	rv := objc.Call[QuitCommand](q_, objc.Sel("init"))
	return rv
}

func (q_ QuitCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) QuitCommand {
	rv := objc.Call[QuitCommand](q_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func QuitCommand_InitWithCommandDescription(commandDef IScriptCommandDescription) QuitCommand {
	return QuitCommandClass.Alloc().InitWithCommandDescription(commandDef)
}

// Returns a constant indicating how to deal with closing any modified documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsquitcommand/1407440-saveoptions?language=objc
func (q_ QuitCommand) SaveOptions() SaveOptions {
	rv := objc.Call[SaveOptions](q_, objc.Sel("saveOptions"))
	return rv
}
