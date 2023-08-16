// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coredata"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ObjectController] class.
var ObjectControllerClass = _ObjectControllerClass{objc.GetClass("NSObjectController")}

type _ObjectControllerClass struct {
	objc.Class
}

// An interface definition for the [ObjectController] class.
type IObjectController interface {
	IController
	AddObject(object objc.IObject)
	PrepareContent()
	NewObject() objc.Object
	RemoveObject(object objc.IObject)
	Fetch(sender objc.IObject) objc.Object
	ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool
	ValidateUserInterfaceItemObject(itemObject objc.IObject) bool
	Remove(sender objc.IObject) objc.Object
	Add(sender objc.IObject) objc.Object
	FetchWithRequestMergeError(fetchRequest coredata.IFetchRequest, merge bool, error foundation.IError) bool
	DefaultFetchRequest() coredata.FetchRequest
	EntityName() string
	SetEntityName(value string)
	Content() objc.Object
	SetContent(value objc.IObject)
	CanAdd() bool
	ObjectClass() objc.Class
	SetObjectClass(value objc.IClass)
	IsEditable() bool
	SetEditable(value bool)
	FetchPredicate() foundation.Predicate
	SetFetchPredicate(value foundation.IPredicate)
	SelectedObjects() []objc.Object
	AutomaticallyPreparesContent() bool
	SetAutomaticallyPreparesContent(value bool)
	Selection() objc.Object
	ManagedObjectContext() coredata.ManagedObjectContext
	SetManagedObjectContext(value coredata.IManagedObjectContext)
	CanRemove() bool
	UsesLazyFetching() bool
	SetUsesLazyFetching(value bool)
}

// A controller that can manage an object's properties referenced by key-value paths. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller?language=objc
type ObjectController struct {
	Controller
}

func ObjectControllerFrom(ptr unsafe.Pointer) ObjectController {
	return ObjectController{
		Controller: ControllerFrom(ptr),
	}
}

func (o_ ObjectController) InitWithContent(content objc.IObject) ObjectController {
	rv := objc.Call[ObjectController](o_, objc.Sel("initWithContent:"), content)
	return rv
}

// Initializes and returns an NSObjectController object with the given content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1529422-initwithcontent?language=objc
func ObjectController_InitWithContent(content objc.IObject) ObjectController {
	return ObjectControllerClass.Alloc().InitWithContent(content)
}

func (oc _ObjectControllerClass) Alloc() ObjectController {
	rv := objc.Call[ObjectController](oc, objc.Sel("alloc"))
	return rv
}

func ObjectController_Alloc() ObjectController {
	return ObjectControllerClass.Alloc()
}

func (oc _ObjectControllerClass) New() ObjectController {
	rv := objc.Call[ObjectController](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewObjectController() ObjectController {
	return ObjectControllerClass.New()
}

func (o_ ObjectController) Init() ObjectController {
	rv := objc.Call[ObjectController](o_, objc.Sel("init"))
	return rv
}

// Sets the receiver’s content object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534093-addobject?language=objc
func (o_ ObjectController) AddObject(object objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("addObject:"), object)
}

// Typically overridden by subclasses that require additional control over the creation of new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534218-preparecontent?language=objc
func (o_ ObjectController) PrepareContent() {
	objc.Call[objc.Void](o_, objc.Sel("prepareContent"))
}

// Creates and returns a new object of the appropriate class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535921-newobject?language=objc
func (o_ ObjectController) NewObject() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("newObject"))
	rv.Autorelease()
	return rv
}

// Removes a given object from the receiver’s content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1532897-removeobject?language=objc
func (o_ ObjectController) RemoveObject(object objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("removeObject:"), object)
}

// Causes the receiver to fetch the data objects specified by the entity name and fetch predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1524554-fetch?language=objc
func (o_ ObjectController) Fetch(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("fetch:"), sender)
	return rv
}

// Returns whether the receiver can handle the action method for a user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534469-validateuserinterfaceitem?language=objc
func (o_ ObjectController) ValidateUserInterfaceItem(item PValidatedUserInterfaceItem) bool {
	po0 := objc.WrapAsProtocol("NSValidatedUserInterfaceItem", item)
	rv := objc.Call[bool](o_, objc.Sel("validateUserInterfaceItem:"), po0)
	return rv
}

// Returns whether the receiver can handle the action method for a user interface item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534469-validateuserinterfaceitem?language=objc
func (o_ ObjectController) ValidateUserInterfaceItemObject(itemObject objc.IObject) bool {
	rv := objc.Call[bool](o_, objc.Sel("validateUserInterfaceItem:"), objc.Ptr(itemObject))
	return rv
}

// Removes the receiver’s content object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1533713-remove?language=objc
func (o_ ObjectController) Remove(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("remove:"), sender)
	return rv
}

// Creates a new object and sets it as the receiver’s content object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1528376-add?language=objc
func (o_ ObjectController) Add(sender objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("add:"), sender)
	return rv
}

// Subclasses should override this method to customize a fetch request, for example to specify fetch limits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1531782-fetchwithrequest?language=objc
func (o_ ObjectController) FetchWithRequestMergeError(fetchRequest coredata.IFetchRequest, merge bool, error foundation.IError) bool {
	rv := objc.Call[bool](o_, objc.Sel("fetchWithRequest:merge:error:"), objc.Ptr(fetchRequest), merge, objc.Ptr(error))
	return rv
}

// Returns the default fetch request used by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1528145-defaultfetchrequest?language=objc
func (o_ ObjectController) DefaultFetchRequest() coredata.FetchRequest {
	rv := objc.Call[coredata.FetchRequest](o_, objc.Sel("defaultFetchRequest"))
	return rv
}

// The entity name used by the receiver to create new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535467-entityname?language=objc
func (o_ ObjectController) EntityName() string {
	rv := objc.Call[string](o_, objc.Sel("entityName"))
	return rv
}

// The entity name used by the receiver to create new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535467-entityname?language=objc
func (o_ ObjectController) SetEntityName(value string) {
	objc.Call[objc.Void](o_, objc.Sel("setEntityName:"), value)
}

// The receiver’s content object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1530826-content?language=objc
func (o_ ObjectController) Content() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("content"))
	return rv
}

// The receiver’s content object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1530826-content?language=objc
func (o_ ObjectController) SetContent(value objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("setContent:"), value)
}

// A Boolean value that indicates whether an object can be added to the receiver using add:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1528497-canadd?language=objc
func (o_ ObjectController) CanAdd() bool {
	rv := objc.Call[bool](o_, objc.Sel("canAdd"))
	return rv
}

// The object class to use when creating new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535459-objectclass?language=objc
func (o_ ObjectController) ObjectClass() objc.Class {
	rv := objc.Call[objc.Class](o_, objc.Sel("objectClass"))
	return rv
}

// The object class to use when creating new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535459-objectclass?language=objc
func (o_ ObjectController) SetObjectClass(value objc.IClass) {
	objc.Call[objc.Void](o_, objc.Sel("setObjectClass:"), objc.Ptr(value))
}

// A Boolean that indicates whether the receiver allows adding and removing objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534699-editable?language=objc
func (o_ ObjectController) IsEditable() bool {
	rv := objc.Call[bool](o_, objc.Sel("isEditable"))
	return rv
}

// A Boolean that indicates whether the receiver allows adding and removing objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534699-editable?language=objc
func (o_ ObjectController) SetEditable(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setEditable:"), value)
}

// The receiver’s fetch predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1533545-fetchpredicate?language=objc
func (o_ ObjectController) FetchPredicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](o_, objc.Sel("fetchPredicate"))
	return rv
}

// The receiver’s fetch predicate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1533545-fetchpredicate?language=objc
func (o_ ObjectController) SetFetchPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](o_, objc.Sel("setFetchPredicate:"), objc.Ptr(value))
}

// An array of all objects to be affected by editing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1535397-selectedobjects?language=objc
func (o_ ObjectController) SelectedObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](o_, objc.Sel("selectedObjects"))
	return rv
}

// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534767-automaticallypreparescontent?language=objc
func (o_ ObjectController) AutomaticallyPreparesContent() bool {
	rv := objc.Call[bool](o_, objc.Sel("automaticallyPreparesContent"))
	return rv
}

// A Boolean that shows whether the receiver automatically creates and inserts new content objects automatically when loading from a nib file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1534767-automaticallypreparescontent?language=objc
func (o_ ObjectController) SetAutomaticallyPreparesContent(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAutomaticallyPreparesContent:"), value)
}

// A proxy object representing the receiver’s selection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1527403-selection?language=objc
func (o_ ObjectController) Selection() objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("selection"))
	return rv
}

// The receiver’s managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1530382-managedobjectcontext?language=objc
func (o_ ObjectController) ManagedObjectContext() coredata.ManagedObjectContext {
	rv := objc.Call[coredata.ManagedObjectContext](o_, objc.Sel("managedObjectContext"))
	return rv
}

// The receiver’s managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1530382-managedobjectcontext?language=objc
func (o_ ObjectController) SetManagedObjectContext(value coredata.IManagedObjectContext) {
	objc.Call[objc.Void](o_, objc.Sel("setManagedObjectContext:"), objc.Ptr(value))
}

// A Boolean value that indicates whether an object can be removed from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1532378-canremove?language=objc
func (o_ ObjectController) CanRemove() bool {
	rv := objc.Call[bool](o_, objc.Sel("canRemove"))
	return rv
}

// A Boolean that indicates whether the receiver uses lazy fetching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1531411-useslazyfetching?language=objc
func (o_ ObjectController) UsesLazyFetching() bool {
	rv := objc.Call[bool](o_, objc.Sel("usesLazyFetching"))
	return rv
}

// A Boolean that indicates whether the receiver uses lazy fetching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsobjectcontroller/1531411-useslazyfetching?language=objc
func (o_ ObjectController) SetUsesLazyFetching(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setUsesLazyFetching:"), value)
}
