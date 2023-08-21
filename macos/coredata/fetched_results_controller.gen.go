// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchedResultsController] class.
var FetchedResultsControllerClass = _FetchedResultsControllerClass{objc.GetClass("NSFetchedResultsController")}

type _FetchedResultsControllerClass struct {
	objc.Class
}

// An interface definition for the [FetchedResultsController] class.
type IFetchedResultsController interface {
	objc.IObject
	PerformFetch(error foundation.IError) bool
	SectionIndexTitleForSectionName(sectionName string) string
	IndexPathForObject(object objc.IObject) foundation.IndexPath
	SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int
	ObjectAtIndexPath(indexPath foundation.IIndexPath) objc.Object
	CacheName() string
	FetchRequest() FetchRequest
	SectionNameKeyPath() string
	Delegate() FetchedResultsControllerDelegateWrapper
	SetDelegate(value PFetchedResultsControllerDelegate)
	SetDelegateObject(valueObject objc.IObject)
	SectionIndexTitles() []string
	FetchedObjects() []objc.Object
	Sections() []FetchedResultsSectionInfoWrapper
	ManagedObjectContext() ManagedObjectContext
}

// A controller that you use to  manage the results of a Core Data fetch request and to display data to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller?language=objc
type FetchedResultsController struct {
	objc.Object
}

func FetchedResultsControllerFrom(ptr unsafe.Pointer) FetchedResultsController {
	return FetchedResultsController{
		Object: objc.ObjectFrom(ptr),
	}
}

func (f_ FetchedResultsController) InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest IFetchRequest, context IManagedObjectContext, sectionNameKeyPath string, name string) FetchedResultsController {
	rv := objc.Call[FetchedResultsController](f_, objc.Sel("initWithFetchRequest:managedObjectContext:sectionNameKeyPath:cacheName:"), objc.Ptr(fetchRequest), objc.Ptr(context), sectionNameKeyPath, name)
	return rv
}

// Returns a fetch request controller initialized using the given arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622282-initwithfetchrequest?language=objc
func NewFetchedResultsControllerWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest IFetchRequest, context IManagedObjectContext, sectionNameKeyPath string, name string) FetchedResultsController {
	instance := FetchedResultsControllerClass.Alloc().InitWithFetchRequestManagedObjectContextSectionNameKeyPathCacheName(fetchRequest, context, sectionNameKeyPath, name)
	instance.Autorelease()
	return instance
}

func (fc _FetchedResultsControllerClass) Alloc() FetchedResultsController {
	rv := objc.Call[FetchedResultsController](fc, objc.Sel("alloc"))
	return rv
}

func FetchedResultsController_Alloc() FetchedResultsController {
	return FetchedResultsControllerClass.Alloc()
}

func (fc _FetchedResultsControllerClass) New() FetchedResultsController {
	rv := objc.Call[FetchedResultsController](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchedResultsController() FetchedResultsController {
	return FetchedResultsControllerClass.New()
}

func (f_ FetchedResultsController) Init() FetchedResultsController {
	rv := objc.Call[FetchedResultsController](f_, objc.Sel("init"))
	return rv
}

// Executes the controller’s fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622305-performfetch?language=objc
func (f_ FetchedResultsController) PerformFetch(error foundation.IError) bool {
	rv := objc.Call[bool](f_, objc.Sel("performFetch:"), objc.Ptr(error))
	return rv
}

// Returns the corresponding section index entry for a given section name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622308-sectionindextitleforsectionname?language=objc
func (f_ FetchedResultsController) SectionIndexTitleForSectionName(sectionName string) string {
	rv := objc.Call[string](f_, objc.Sel("sectionIndexTitleForSectionName:"), sectionName)
	return rv
}

// Deletes the cached section information with the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622283-deletecachewithname?language=objc
func (fc _FetchedResultsControllerClass) DeleteCacheWithName(name string) {
	objc.Call[objc.Void](fc, objc.Sel("deleteCacheWithName:"), name)
}

// Deletes the cached section information with the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622283-deletecachewithname?language=objc
func FetchedResultsController_DeleteCacheWithName(name string) {
	FetchedResultsControllerClass.DeleteCacheWithName(name)
}

// Returns the index path of a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622306-indexpathforobject?language=objc
func (f_ FetchedResultsController) IndexPathForObject(object objc.IObject) foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](f_, objc.Sel("indexPathForObject:"), objc.Ptr(object))
	return rv
}

// Returns the section number for a given section title and index in the section index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622284-sectionforsectionindextitle?language=objc
func (f_ FetchedResultsController) SectionForSectionIndexTitleAtIndex(title string, sectionIndex int) int {
	rv := objc.Call[int](f_, objc.Sel("sectionForSectionIndexTitle:atIndex:"), title, sectionIndex)
	return rv
}

// Returns the object at the given index path in the fetch results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622281-objectatindexpath?language=objc
func (f_ FetchedResultsController) ObjectAtIndexPath(indexPath foundation.IIndexPath) objc.Object {
	rv := objc.Call[objc.Object](f_, objc.Sel("objectAtIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// The name of the file used to cache section information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622280-cachename?language=objc
func (f_ FetchedResultsController) CacheName() string {
	rv := objc.Call[string](f_, objc.Sel("cacheName"))
	return rv
}

// The fetch request used to do the fetching. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622287-fetchrequest?language=objc
func (f_ FetchedResultsController) FetchRequest() FetchRequest {
	rv := objc.Call[FetchRequest](f_, objc.Sel("fetchRequest"))
	return rv
}

// The key path of the attribute that determines which section the fetched entity belongs to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622285-sectionnamekeypath?language=objc
func (f_ FetchedResultsController) SectionNameKeyPath() string {
	rv := objc.Call[string](f_, objc.Sel("sectionNameKeyPath"))
	return rv
}

// The object that is notified when the fetched results changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622291-delegate?language=objc
func (f_ FetchedResultsController) Delegate() FetchedResultsControllerDelegateWrapper {
	rv := objc.Call[FetchedResultsControllerDelegateWrapper](f_, objc.Sel("delegate"))
	return rv
}

// The object that is notified when the fetched results changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622291-delegate?language=objc
func (f_ FetchedResultsController) SetDelegate(value PFetchedResultsControllerDelegate) {
	po0 := objc.WrapAsProtocol("NSFetchedResultsControllerDelegate", value)
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), po0)
}

// The object that is notified when the fetched results changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622291-delegate?language=objc
func (f_ FetchedResultsController) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The array of section index titles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622299-sectionindextitles?language=objc
func (f_ FetchedResultsController) SectionIndexTitles() []string {
	rv := objc.Call[[]string](f_, objc.Sel("sectionIndexTitles"))
	return rv
}

// The results of the fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622278-fetchedobjects?language=objc
func (f_ FetchedResultsController) FetchedObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](f_, objc.Sel("fetchedObjects"))
	return rv
}

// The sections for the fetch results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622303-sections?language=objc
func (f_ FetchedResultsController) Sections() []FetchedResultsSectionInfoWrapper {
	rv := objc.Call[[]FetchedResultsSectionInfoWrapper](f_, objc.Sel("sections"))
	return rv
}

// The managed object context used to fetch objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchedresultscontroller/1622304-managedobjectcontext?language=objc
func (f_ FetchedResultsController) ManagedObjectContext() ManagedObjectContext {
	rv := objc.Call[ManagedObjectContext](f_, objc.Sel("managedObjectContext"))
	return rv
}
