package helper

import "runtime"
import "reflect"
import "strings"
import "fmt"
import "github.com/progrium/macdriver/objc"
import "github.com/progrium/macdriver/cocoa"


// Adds the ActionHelper class to the objective-c system
func init() {
	newClass := objc.NewClass("ActionHelper", "NSObject")
	objc.RegisterClass(newClass)
}


// Creates a unique method name using a control and a Go function
func createActionName(control interface{}, callback func(sender interface{})) string {
	actionName := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	actionName = fmt.Sprintf("%s%s", actionName, control)
	
	// remove any characters that should not be in a method's name
	actionName = strings.ReplaceAll(actionName, " ", "") // remove space
	actionName = strings.ReplaceAll(actionName, ":", "") // remove middle colon
	actionName = strings.ReplaceAll(actionName, "<", "") // remove <
	actionName = strings.ReplaceAll(actionName, ">", "") // remove >
	actionName = actionName + ":"						 // add colon at end
	return actionName
}


// Sets a control's action and target for use with a Go function
// Input 1: control - NSControl subclass
// Input 2: callback - a Go function that has one argument of interface{}
func SetAction(control cocoa.NSControlRef, callback func(sender interface{})) {
	actionClass := objc.Get("ActionHelper")
	actionName := createActionName(control, callback)
	actionClass.AddMethod(actionName, func(sender objc.Object){
		callback(control)
	})
	theControl := cocoa.NSControl_FromRef(control)
	actionObj := actionClass.Alloc()
	theControl.SetTarget(actionObj)
	theControl.SetAction(objc.Sel(actionName))
}
