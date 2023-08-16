// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FetchRequest] class.
var FetchRequestClass = _FetchRequestClass{objc.GetClass("NSFetchRequest")}

type _FetchRequestClass struct {
	objc.Class
}

// An interface definition for the [FetchRequest] class.
type IFetchRequest interface {
	IPersistentStoreRequest
	Execute(error foundation.IError) []objc.Object
	EntityName() string
	ShouldRefreshRefetchedObjects() bool
	SetShouldRefreshRefetchedObjects(value bool)
	IncludesSubentities() bool
	SetIncludesSubentities(value bool)
	IncludesPendingChanges() bool
	SetIncludesPendingChanges(value bool)
	PropertiesToGroupBy() []objc.Object
	SetPropertiesToGroupBy(value []objc.IObject)
	Entity() EntityDescription
	SetEntity(value IEntityDescription)
	FetchOffset() uint
	SetFetchOffset(value uint)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	FetchBatchSize() uint
	SetFetchBatchSize(value uint)
	ReturnsDistinctResults() bool
	SetReturnsDistinctResults(value bool)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
	PropertiesToFetch() []objc.Object
	SetPropertiesToFetch(value []objc.IObject)
	IncludesPropertyValues() bool
	SetIncludesPropertyValues(value bool)
	ReturnsObjectsAsFaults() bool
	SetReturnsObjectsAsFaults(value bool)
	FetchLimit() uint
	SetFetchLimit(value uint)
	HavingPredicate() foundation.Predicate
	SetHavingPredicate(value foundation.IPredicate)
	ResultType() FetchRequestResultType
	SetResultType(value FetchRequestResultType)
	RelationshipKeyPathsForPrefetching() []string
	SetRelationshipKeyPathsForPrefetching(value []string)
}

// A description of search criteria used to retrieve data from a persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest?language=objc
type FetchRequest struct {
	PersistentStoreRequest
}

func FetchRequestFrom(ptr unsafe.Pointer) FetchRequest {
	return FetchRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

func (fc _FetchRequestClass) FetchRequestWithEntityName(entityName string) FetchRequest {
	rv := objc.Call[FetchRequest](fc, objc.Sel("fetchRequestWithEntityName:"), entityName)
	return rv
}

// Returns a fetch request configured with a given entity name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1563437-fetchrequestwithentityname?language=objc
func FetchRequest_FetchRequestWithEntityName(entityName string) FetchRequest {
	return FetchRequestClass.FetchRequestWithEntityName(entityName)
}

func (f_ FetchRequest) Init() FetchRequest {
	rv := objc.Call[FetchRequest](f_, objc.Sel("init"))
	return rv
}

func (f_ FetchRequest) InitWithEntityName(entityName string) FetchRequest {
	rv := objc.Call[FetchRequest](f_, objc.Sel("initWithEntityName:"), entityName)
	return rv
}

// Initializes a fetch request configured with a given entity name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506802-initwithentityname?language=objc
func FetchRequest_InitWithEntityName(entityName string) FetchRequest {
	return FetchRequestClass.Alloc().InitWithEntityName(entityName)
}

func (fc _FetchRequestClass) Alloc() FetchRequest {
	rv := objc.Call[FetchRequest](fc, objc.Sel("alloc"))
	return rv
}

func FetchRequest_Alloc() FetchRequest {
	return FetchRequestClass.Alloc()
}

func (fc _FetchRequestClass) New() FetchRequest {
	rv := objc.Call[FetchRequest](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFetchRequest() FetchRequest {
	return FetchRequestClass.New()
}

// Executes the fetch request against the managed object context that is associated with the current queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1640594-execute?language=objc
func (f_ FetchRequest) Execute(error foundation.IError) []objc.Object {
	rv := objc.Call[[]objc.Object](f_, objc.Sel("execute:"), objc.Ptr(error))
	return rv
}

// The name of the entity the request is configured to fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506233-entityname?language=objc
func (f_ FetchRequest) EntityName() string {
	rv := objc.Call[string](f_, objc.Sel("entityName"))
	return rv
}

// A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506440-shouldrefreshrefetchedobjects?language=objc
func (f_ FetchRequest) ShouldRefreshRefetchedObjects() bool {
	rv := objc.Call[bool](f_, objc.Sel("shouldRefreshRefetchedObjects"))
	return rv
}

// A Boolean value that indicates whether the property values of fetched objects will be updated with the current values in the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506440-shouldrefreshrefetchedobjects?language=objc
func (f_ FetchRequest) SetShouldRefreshRefetchedObjects(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setShouldRefreshRefetchedObjects:"), value)
}

// A Boolean value that indicates whether the fetch request includes subentities in the results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506366-includessubentities?language=objc
func (f_ FetchRequest) IncludesSubentities() bool {
	rv := objc.Call[bool](f_, objc.Sel("includesSubentities"))
	return rv
}

// A Boolean value that indicates whether the fetch request includes subentities in the results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506366-includessubentities?language=objc
func (f_ FetchRequest) SetIncludesSubentities(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setIncludesSubentities:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506724-includespendingchanges?language=objc
func (f_ FetchRequest) IncludesPendingChanges() bool {
	rv := objc.Call[bool](f_, objc.Sel("includesPendingChanges"))
	return rv
}

// A Boolean value that indicates whether, when the fetch is executed, it matches against currently unsaved changes in the managed object context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506724-includespendingchanges?language=objc
func (f_ FetchRequest) SetIncludesPendingChanges(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setIncludesPendingChanges:"), value)
}

// An array of objects that indicates how data should be grouped before a select statement is run in a SQL database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506191-propertiestogroupby?language=objc
func (f_ FetchRequest) PropertiesToGroupBy() []objc.Object {
	rv := objc.Call[[]objc.Object](f_, objc.Sel("propertiesToGroupBy"))
	return rv
}

// An array of objects that indicates how data should be grouped before a select statement is run in a SQL database. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506191-propertiestogroupby?language=objc
func (f_ FetchRequest) SetPropertiesToGroupBy(value []objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setPropertiesToGroupBy:"), value)
}

// The entity specified for the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506979-entity?language=objc
func (f_ FetchRequest) Entity() EntityDescription {
	rv := objc.Call[EntityDescription](f_, objc.Sel("entity"))
	return rv
}

// The entity specified for the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506979-entity?language=objc
func (f_ FetchRequest) SetEntity(value IEntityDescription) {
	objc.Call[objc.Void](f_, objc.Sel("setEntity:"), objc.Ptr(value))
}

// The fetch offset of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506770-fetchoffset?language=objc
func (f_ FetchRequest) FetchOffset() uint {
	rv := objc.Call[uint](f_, objc.Sel("fetchOffset"))
	return rv
}

// The fetch offset of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506770-fetchoffset?language=objc
func (f_ FetchRequest) SetFetchOffset(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchOffset:"), value)
}

// The sort descriptors of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506262-sortdescriptors?language=objc
func (f_ FetchRequest) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Call[[]foundation.SortDescriptor](f_, objc.Sel("sortDescriptors"))
	return rv
}

// The sort descriptors of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506262-sortdescriptors?language=objc
func (f_ FetchRequest) SetSortDescriptors(value []foundation.ISortDescriptor) {
	objc.Call[objc.Void](f_, objc.Sel("setSortDescriptors:"), value)
}

// The batch size of the objects specified in the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506558-fetchbatchsize?language=objc
func (f_ FetchRequest) FetchBatchSize() uint {
	rv := objc.Call[uint](f_, objc.Sel("fetchBatchSize"))
	return rv
}

// The batch size of the objects specified in the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506558-fetchbatchsize?language=objc
func (f_ FetchRequest) SetFetchBatchSize(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchBatchSize:"), value)
}

// A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by propertiesToFetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506344-returnsdistinctresults?language=objc
func (f_ FetchRequest) ReturnsDistinctResults() bool {
	rv := objc.Call[bool](f_, objc.Sel("returnsDistinctResults"))
	return rv
}

// A Boolean value that indicates whether the fetch request returns only distinct values for the fields specified by propertiesToFetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506344-returnsdistinctresults?language=objc
func (f_ FetchRequest) SetReturnsDistinctResults(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setReturnsDistinctResults:"), value)
}

// The predicate of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506638-predicate?language=objc
func (f_ FetchRequest) Predicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](f_, objc.Sel("predicate"))
	return rv
}

// The predicate of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506638-predicate?language=objc
func (f_ FetchRequest) SetPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](f_, objc.Sel("setPredicate:"), objc.Ptr(value))
}

// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506851-propertiestofetch?language=objc
func (f_ FetchRequest) PropertiesToFetch() []objc.Object {
	rv := objc.Call[[]objc.Object](f_, objc.Sel("propertiesToFetch"))
	return rv
}

// A collection of either property descriptions or string property names that specify which properties should be returned by the fetch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506851-propertiestofetch?language=objc
func (f_ FetchRequest) SetPropertiesToFetch(value []objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("setPropertiesToFetch:"), value)
}

// A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506387-includespropertyvalues?language=objc
func (f_ FetchRequest) IncludesPropertyValues() bool {
	rv := objc.Call[bool](f_, objc.Sel("includesPropertyValues"))
	return rv
}

// A Boolean value that indicates whether, when the fetch is executed, property data is obtained from the persistent store. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506387-includespropertyvalues?language=objc
func (f_ FetchRequest) SetIncludesPropertyValues(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setIncludesPropertyValues:"), value)
}

// A Boolean value that indicates whether the objects resulting from a fetch request are faults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506756-returnsobjectsasfaults?language=objc
func (f_ FetchRequest) ReturnsObjectsAsFaults() bool {
	rv := objc.Call[bool](f_, objc.Sel("returnsObjectsAsFaults"))
	return rv
}

// A Boolean value that indicates whether the objects resulting from a fetch request are faults. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506756-returnsobjectsasfaults?language=objc
func (f_ FetchRequest) SetReturnsObjectsAsFaults(value bool) {
	objc.Call[objc.Void](f_, objc.Sel("setReturnsObjectsAsFaults:"), value)
}

// The fetch limit of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506622-fetchlimit?language=objc
func (f_ FetchRequest) FetchLimit() uint {
	rv := objc.Call[uint](f_, objc.Sel("fetchLimit"))
	return rv
}

// The fetch limit of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506622-fetchlimit?language=objc
func (f_ FetchRequest) SetFetchLimit(value uint) {
	objc.Call[objc.Void](f_, objc.Sel("setFetchLimit:"), value)
}

// The predicate used to filter rows being returned by a query containing a GROUP BY directive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506429-havingpredicate?language=objc
func (f_ FetchRequest) HavingPredicate() foundation.Predicate {
	rv := objc.Call[foundation.Predicate](f_, objc.Sel("havingPredicate"))
	return rv
}

// The predicate used to filter rows being returned by a query containing a GROUP BY directive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506429-havingpredicate?language=objc
func (f_ FetchRequest) SetHavingPredicate(value foundation.IPredicate) {
	objc.Call[objc.Void](f_, objc.Sel("setHavingPredicate:"), objc.Ptr(value))
}

// The result type of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506189-resulttype?language=objc
func (f_ FetchRequest) ResultType() FetchRequestResultType {
	rv := objc.Call[FetchRequestResultType](f_, objc.Sel("resultType"))
	return rv
}

// The result type of the fetch request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506189-resulttype?language=objc
func (f_ FetchRequest) SetResultType(value FetchRequestResultType) {
	objc.Call[objc.Void](f_, objc.Sel("setResultType:"), value)
}

// The relationship key paths to prefetch along with the entity for the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506813-relationshipkeypathsforprefetchi?language=objc
func (f_ FetchRequest) RelationshipKeyPathsForPrefetching() []string {
	rv := objc.Call[[]string](f_, objc.Sel("relationshipKeyPathsForPrefetching"))
	return rv
}

// The relationship key paths to prefetch along with the entity for the request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsfetchrequest/1506813-relationshipkeypathsforprefetchi?language=objc
func (f_ FetchRequest) SetRelationshipKeyPathsForPrefetching(value []string) {
	objc.Call[objc.Void](f_, objc.Sel("setRelationshipKeyPathsForPrefetching:"), value)
}
