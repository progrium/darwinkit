// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptCommand] class.
var ScriptCommandClass = _ScriptCommandClass{objc.GetClass("NSScriptCommand")}

type _ScriptCommandClass struct {
	objc.Class
}

// An interface definition for the [ScriptCommand] class.
type IScriptCommand interface {
	objc.IObject
	ExecuteCommand() objc.Object
	PerformDefaultImplementation() objc.Object
	ResumeExecutionWithResult(result objc.IObject)
	SuspendExecution()
	Arguments() map[string]objc.Object
	SetArguments(value map[string]objc.IObject)
	ReceiversSpecifier() ScriptObjectSpecifier
	SetReceiversSpecifier(value IScriptObjectSpecifier)
	ScriptErrorString() string
	SetScriptErrorString(value string)
	DirectParameter() objc.Object
	SetDirectParameter(value objc.IObject)
	CommandDescription() ScriptCommandDescription
	EvaluatedArguments() map[string]objc.Object
	EvaluatedReceivers() objc.Object
	ScriptErrorNumber() int
	SetScriptErrorNumber(value int)
	IsWellFormed() bool
	AppleEvent() AppleEventDescriptor
	ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor
	SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor)
	ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor
	SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor)
}

// A self-contained scripting statement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand?language=objc
type ScriptCommand struct {
	objc.Object
}

func ScriptCommandFrom(ptr unsafe.Pointer) ScriptCommand {
	return ScriptCommand{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScriptCommand) InitWithCommandDescription(commandDef IScriptCommandDescription) ScriptCommand {
	rv := objc.Call[ScriptCommand](s_, objc.Sel("initWithCommandDescription:"), objc.Ptr(commandDef))
	return rv
}

// Returns an a script command object initialized from the passed command description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413516-initwithcommanddescription?language=objc
func NewScriptCommandWithCommandDescription(commandDef IScriptCommandDescription) ScriptCommand {
	instance := ScriptCommandClass.Alloc().InitWithCommandDescription(commandDef)
	instance.Autorelease()
	return instance
}

func (sc _ScriptCommandClass) Alloc() ScriptCommand {
	rv := objc.Call[ScriptCommand](sc, objc.Sel("alloc"))
	return rv
}

func ScriptCommand_Alloc() ScriptCommand {
	return ScriptCommandClass.Alloc()
}

func (sc _ScriptCommandClass) New() ScriptCommand {
	rv := objc.Call[ScriptCommand](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCommand() ScriptCommand {
	return ScriptCommandClass.New()
}

func (s_ ScriptCommand) Init() ScriptCommand {
	rv := objc.Call[ScriptCommand](s_, objc.Sel("init"))
	return rv
}

// Executes the command if it is valid and returns the result, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1414780-executecommand?language=objc
func (s_ ScriptCommand) ExecuteCommand() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("executeCommand"))
	return rv
}

// If a command is being executed in the current thread by Cocoa scripting's built-in Apple event handling, return the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1408255-currentcommand?language=objc
func (sc _ScriptCommandClass) CurrentCommand() ScriptCommand {
	rv := objc.Call[ScriptCommand](sc, objc.Sel("currentCommand"))
	return rv
}

// If a command is being executed in the current thread by Cocoa scripting's built-in Apple event handling, return the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1408255-currentcommand?language=objc
func ScriptCommand_CurrentCommand() ScriptCommand {
	return ScriptCommandClass.CurrentCommand()
}

// Overridden by subclasses to provide a default implementation for the command represented by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413379-performdefaultimplementation?language=objc
func (s_ ScriptCommand) PerformDefaultImplementation() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("performDefaultImplementation"))
	return rv
}

// If a successful, unmatched, invocation of suspendExecution has been made, resume the execution of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1408227-resumeexecutionwithresult?language=objc
func (s_ ScriptCommand) ResumeExecutionWithResult(result objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("resumeExecutionWithResult:"), result)
}

// Suspends the execution of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1417785-suspendexecution?language=objc
func (s_ ScriptCommand) SuspendExecution() {
	objc.Call[objc.Void](s_, objc.Sel("suspendExecution"))
}

// Sets the arguments of the command to args. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1414071-arguments?language=objc
func (s_ ScriptCommand) Arguments() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](s_, objc.Sel("arguments"))
	return rv
}

// Sets the arguments of the command to args. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1414071-arguments?language=objc
func (s_ ScriptCommand) SetArguments(value map[string]objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setArguments:"), value)
}

// Sets the object specifier to receiversSpec that, when evaluated, indicates the receiver or receivers of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1417016-receiversspecifier?language=objc
func (s_ ScriptCommand) ReceiversSpecifier() ScriptObjectSpecifier {
	rv := objc.Call[ScriptObjectSpecifier](s_, objc.Sel("receiversSpecifier"))
	return rv
}

// Sets the object specifier to receiversSpec that, when evaluated, indicates the receiver or receivers of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1417016-receiversspecifier?language=objc
func (s_ ScriptCommand) SetReceiversSpecifier(value IScriptObjectSpecifier) {
	objc.Call[objc.Void](s_, objc.Sel("setReceiversSpecifier:"), objc.Ptr(value))
}

// Sets a script error string that is associated with execution of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1414596-scripterrorstring?language=objc
func (s_ ScriptCommand) ScriptErrorString() string {
	rv := objc.Call[string](s_, objc.Sel("scriptErrorString"))
	return rv
}

// Sets a script error string that is associated with execution of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1414596-scripterrorstring?language=objc
func (s_ ScriptCommand) SetScriptErrorString(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setScriptErrorString:"), value)
}

// Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1410675-directparameter?language=objc
func (s_ ScriptCommand) DirectParameter() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("directParameter"))
	return rv
}

// Sets the object that corresponds to the direct parameter of the Apple event from which the receiver derives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1410675-directparameter?language=objc
func (s_ ScriptCommand) SetDirectParameter(value objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDirectParameter:"), value)
}

// Returns the command description for the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1407452-commanddescription?language=objc
func (s_ ScriptCommand) CommandDescription() ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](s_, objc.Sel("commandDescription"))
	return rv
}

// Returns a dictionary containing the arguments of the command, evaluated from object specifiers to objects if necessary. The keys in the dictionary are the argument names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413335-evaluatedarguments?language=objc
func (s_ ScriptCommand) EvaluatedArguments() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](s_, objc.Sel("evaluatedArguments"))
	return rv
}

// Returns the object or objects to which the command is to be sent (called both the “receivers” or “targets” of script commands). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1411257-evaluatedreceivers?language=objc
func (s_ ScriptCommand) EvaluatedReceivers() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("evaluatedReceivers"))
	return rv
}

// Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1411484-scripterrornumber?language=objc
func (s_ ScriptCommand) ScriptErrorNumber() int {
	rv := objc.Call[int](s_, objc.Sel("scriptErrorNumber"))
	return rv
}

// Sets a script error number that is associated with the execution of the command and is returned in the reply Apple event, if a reply was requested by the sender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1411484-scripterrornumber?language=objc
func (s_ ScriptCommand) SetScriptErrorNumber(value int) {
	objc.Call[objc.Void](s_, objc.Sel("setScriptErrorNumber:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1413090-wellformed?language=objc
func (s_ ScriptCommand) IsWellFormed() bool {
	rv := objc.Call[bool](s_, objc.Sel("isWellFormed"))
	return rv
}

// If the receiver was constructed by Cocoa scripting's built-in Apple event handling, returns the Apple event descriptor from which it was constructed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1415626-appleevent?language=objc
func (s_ ScriptCommand) AppleEvent() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](s_, objc.Sel("appleEvent"))
	return rv
}

// Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1417217-scripterroroffendingobjectdescri?language=objc
func (s_ ScriptCommand) ScriptErrorOffendingObjectDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](s_, objc.Sel("scriptErrorOffendingObjectDescriptor"))
	return rv
}

// Sets a descriptor for an object that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1417217-scripterroroffendingobjectdescri?language=objc
func (s_ ScriptCommand) SetScriptErrorOffendingObjectDescriptor(value IAppleEventDescriptor) {
	objc.Call[objc.Void](s_, objc.Sel("setScriptErrorOffendingObjectDescriptor:"), objc.Ptr(value))
}

// Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1411714-scripterrorexpectedtypedescripto?language=objc
func (s_ ScriptCommand) ScriptErrorExpectedTypeDescriptor() AppleEventDescriptor {
	rv := objc.Call[AppleEventDescriptor](s_, objc.Sel("scriptErrorExpectedTypeDescriptor"))
	return rv
}

// Sets a descriptor for the expected type that will be put in the reply Apple event if the sender requested a reply, execution of the receiver completes, and an error number was set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommand/1411714-scripterrorexpectedtypedescripto?language=objc
func (s_ ScriptCommand) SetScriptErrorExpectedTypeDescriptor(value IAppleEventDescriptor) {
	objc.Call[objc.Void](s_, objc.Sel("setScriptErrorExpectedTypeDescriptor:"), objc.Ptr(value))
}
