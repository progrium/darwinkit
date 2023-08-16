// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [UserDefaultsController] class.
var UserDefaultsControllerClass = _UserDefaultsControllerClass{objc.GetClass("NSUserDefaultsController")}

type _UserDefaultsControllerClass struct {
	objc.Class
}

// An interface definition for the [UserDefaultsController] class.
type IUserDefaultsController interface {
	IController
	Save(sender objc.IObject) objc.Object
	Revert(sender objc.IObject) objc.Object
	RevertToInitialValues(sender objc.IObject) objc.Object
	Defaults() foundation.UserDefaults
	HasUnappliedChanges() bool
	AppliesImmediately() bool
	SetAppliesImmediately(value bool)
	InitialValues() map[string]objc.Object
	SetInitialValues(value map[string]objc.IObject)
	Values() objc.Object
}

// A controller that accesses user preference information for your app from the user’s defaults database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller?language=objc
type UserDefaultsController struct {
	Controller
}

func UserDefaultsControllerFrom(ptr unsafe.Pointer) UserDefaultsController {
	return UserDefaultsController{
		Controller: ControllerFrom(ptr),
	}
}

func (u_ UserDefaultsController) InitWithDefaultsInitialValues(defaults foundation.IUserDefaults, initialValues map[string]objc.IObject) UserDefaultsController {
	rv := objc.Call[UserDefaultsController](u_, objc.Sel("initWithDefaults:initialValues:"), objc.Ptr(defaults), initialValues)
	return rv
}

// Returns an initialized NSUserDefaultsController object using the NSUserDefaults instance specified in defaults and the initial default values contained in the initialValues dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388184-initwithdefaults?language=objc
func UserDefaultsController_InitWithDefaultsInitialValues(defaults foundation.IUserDefaults, initialValues map[string]objc.IObject) UserDefaultsController {
	return UserDefaultsControllerClass.Alloc().InitWithDefaultsInitialValues(defaults, initialValues)
}

func (uc _UserDefaultsControllerClass) Alloc() UserDefaultsController {
	rv := objc.Call[UserDefaultsController](uc, objc.Sel("alloc"))
	return rv
}

func UserDefaultsController_Alloc() UserDefaultsController {
	return UserDefaultsControllerClass.Alloc()
}

func (uc _UserDefaultsControllerClass) New() UserDefaultsController {
	rv := objc.Call[UserDefaultsController](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewUserDefaultsController() UserDefaultsController {
	return UserDefaultsControllerClass.New()
}

func (u_ UserDefaultsController) Init() UserDefaultsController {
	rv := objc.Call[UserDefaultsController](u_, objc.Sel("init"))
	return rv
}

// Saves the values of the receiver’s user default properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388180-save?language=objc
func (u_ UserDefaultsController) Save(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("save:"), sender)
	return rv
}

// Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388178-revert?language=objc
func (u_ UserDefaultsController) Revert(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("revert:"), sender)
	return rv
}

// Causes the receiver to discard all edits and replace the values of all the user default properties with any corresponding values in the initialValues dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388170-reverttoinitialvalues?language=objc
func (u_ UserDefaultsController) RevertToInitialValues(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("revertToInitialValues:"), sender)
	return rv
}

// Returns the instance of NSUserDefaults in use by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388182-defaults?language=objc
func (u_ UserDefaultsController) Defaults() foundation.UserDefaults {
	rv := objc.Call[foundation.UserDefaults](u_, objc.Sel("defaults"))
	return rv
}

// Returns whether the receiver has user default values that have not been saved to NSUserDefaults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388186-hasunappliedchanges?language=objc
func (u_ UserDefaultsController) HasUnappliedChanges() bool {
	rv := objc.Call[bool](u_, objc.Sel("hasUnappliedChanges"))
	return rv
}

// Returns whether any changes made to bound user default properties are saved immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388191-appliesimmediately?language=objc
func (u_ UserDefaultsController) AppliesImmediately() bool {
	rv := objc.Call[bool](u_, objc.Sel("appliesImmediately"))
	return rv
}

// Returns whether any changes made to bound user default properties are saved immediately. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388191-appliesimmediately?language=objc
func (u_ UserDefaultsController) SetAppliesImmediately(value bool) {
	objc.Call[objc.Void](u_, objc.Sel("setAppliesImmediately:"), value)
}

// Returns a dictionary containing the receiver’s initial default values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388188-initialvalues?language=objc
func (u_ UserDefaultsController) InitialValues() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](u_, objc.Sel("initialValues"))
	return rv
}

// Returns a dictionary containing the receiver’s initial default values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388188-initialvalues?language=objc
func (u_ UserDefaultsController) SetInitialValues(value map[string]objc.IObject) {
	objc.Call[objc.Void](u_, objc.Sel("setInitialValues:"), value)
}

// Returns a key value coding compliant object that is used to access the user default properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388176-values?language=objc
func (u_ UserDefaultsController) Values() objc.Object {
	rv := objc.Call[objc.Object](u_, objc.Sel("values"))
	return rv
}

// Returns the shared instance of NSUserDefaultsController, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388190-shareduserdefaultscontroller?language=objc
func (uc _UserDefaultsControllerClass) SharedUserDefaultsController() UserDefaultsController {
	rv := objc.Call[UserDefaultsController](uc, objc.Sel("sharedUserDefaultsController"))
	return rv
}

// Returns the shared instance of NSUserDefaultsController, creating it if necessary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuserdefaultscontroller/1388190-shareduserdefaultscontroller?language=objc
func UserDefaultsController_SharedUserDefaultsController() UserDefaultsController {
	return UserDefaultsControllerClass.SharedUserDefaultsController()
}
