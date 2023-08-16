// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataQuery] class.
var MetadataQueryClass = _MetadataQueryClass{objc.GetClass("NSMetadataQuery")}

type _MetadataQueryClass struct {
	objc.Class
}

// An interface definition for the [MetadataQuery] class.
type IMetadataQuery interface {
	objc.IObject
	ResultAtIndex(idx uint) objc.Object
	StopQuery()
	StartQuery() bool
	IndexOfResult(result objc.IObject) uint
	EnumerateResultsUsingBlock(block func(result objc.Object, idx uint, stop *bool))
	DisableUpdates()
	ValueOfAttributeForResultAtIndex(attrName string, idx uint) objc.Object
	EnableUpdates()
	EnumerateResultsWithOptionsUsingBlock(opts EnumerationOptions, block func(result objc.Object, idx uint, stop *bool))
	ValueListAttributes() []string
	SetValueListAttributes(value []string)
	NotificationBatchingInterval() TimeInterval
	SetNotificationBatchingInterval(value TimeInterval)
	ValueLists() map[string][]MetadataQueryAttributeValueTuple
	GroupedResults() []MetadataQueryResultGroup
	SortDescriptors() []SortDescriptor
	SetSortDescriptors(value []ISortDescriptor)
	Results() []objc.Object
	IsStopped() bool
	SearchItems() []objc.Object
	SetSearchItems(value []objc.IObject)
	GroupingAttributes() []string
	SetGroupingAttributes(value []string)
	Delegate() MetadataQueryDelegateWrapper
	SetDelegate(value PMetadataQueryDelegate)
	SetDelegateObject(valueObject objc.IObject)
	Predicate() Predicate
	SetPredicate(value IPredicate)
	IsStarted() bool
	IsGathering() bool
	SearchScopes() []objc.Object
	SetSearchScopes(value []objc.IObject)
	ResultCount() uint
	OperationQueue() OperationQueue
	SetOperationQueue(value IOperationQueue)
}

// A query that you perform against Spotlight metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery?language=objc
type MetadataQuery struct {
	objc.Object
}

func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MetadataQueryClass) Alloc() MetadataQuery {
	rv := objc.Call[MetadataQuery](mc, objc.Sel("alloc"))
	return rv
}

func MetadataQuery_Alloc() MetadataQuery {
	return MetadataQueryClass.Alloc()
}

func (mc _MetadataQueryClass) New() MetadataQuery {
	rv := objc.Call[MetadataQuery](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataQuery() MetadataQuery {
	return MetadataQueryClass.New()
}

func (m_ MetadataQuery) Init() MetadataQuery {
	rv := objc.Call[MetadataQuery](m_, objc.Sel("init"))
	return rv
}

// Returns the query result at a specific index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1410162-resultatindex?language=objc
func (m_ MetadataQuery) ResultAtIndex(idx uint) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("resultAtIndex:"), idx)
	return rv
}

// Stops the receiver’s current query from gathering any further results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1408021-stopquery?language=objc
func (m_ MetadataQuery) StopQuery() {
	objc.Call[objc.Void](m_, objc.Sel("stopQuery"))
}

// Attempts to start the query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1407304-startquery?language=objc
func (m_ MetadataQuery) StartQuery() bool {
	rv := objc.Call[bool](m_, objc.Sel("startQuery"))
	return rv
}

// Returns the index of a query result object in the receiver’s results array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1410014-indexofresult?language=objc
func (m_ MetadataQuery) IndexOfResult(result objc.IObject) uint {
	rv := objc.Call[uint](m_, objc.Sel("indexOfResult:"), result)
	return rv
}

// Enumerates the current set of results using the given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1415856-enumerateresultsusingblock?language=objc
func (m_ MetadataQuery) EnumerateResultsUsingBlock(block func(result objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](m_, objc.Sel("enumerateResultsUsingBlock:"), block)
}

// Disables updates to the query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1416337-disableupdates?language=objc
func (m_ MetadataQuery) DisableUpdates() {
	objc.Call[objc.Void](m_, objc.Sel("disableUpdates"))
}

// Returns the value for the attribute name attrName at the index in the results specified by idx. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1417133-valueofattribute?language=objc
func (m_ MetadataQuery) ValueOfAttributeForResultAtIndex(attrName string, idx uint) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("valueOfAttribute:forResultAtIndex:"), attrName, idx)
	return rv
}

// Enables updates to the query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1416943-enableupdates?language=objc
func (m_ MetadataQuery) EnableUpdates() {
	objc.Call[objc.Void](m_, objc.Sel("enableUpdates"))
}

// Enumerates the current set of results using the given options and block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1415123-enumerateresultswithoptions?language=objc
func (m_ MetadataQuery) EnumerateResultsWithOptionsUsingBlock(opts EnumerationOptions, block func(result objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](m_, objc.Sel("enumerateResultsWithOptions:usingBlock:"), opts, block)
}

// An array of attributes whose values are gathered by the query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1407767-valuelistattributes?language=objc
func (m_ MetadataQuery) ValueListAttributes() []string {
	rv := objc.Call[[]string](m_, objc.Sel("valueListAttributes"))
	return rv
}

// An array of attributes whose values are gathered by the query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1407767-valuelistattributes?language=objc
func (m_ MetadataQuery) SetValueListAttributes(value []string) {
	objc.Call[objc.Void](m_, objc.Sel("setValueListAttributes:"), value)
}

// The interval at which notification of updated results occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411884-notificationbatchinginterval?language=objc
func (m_ MetadataQuery) NotificationBatchingInterval() TimeInterval {
	rv := objc.Call[TimeInterval](m_, objc.Sel("notificationBatchingInterval"))
	return rv
}

// The interval at which notification of updated results occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411884-notificationbatchinginterval?language=objc
func (m_ MetadataQuery) SetNotificationBatchingInterval(value TimeInterval) {
	objc.Call[objc.Void](m_, objc.Sel("setNotificationBatchingInterval:"), value)
}

// A dictionary containing the value lists generated by the query. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1418401-valuelists?language=objc
func (m_ MetadataQuery) ValueLists() map[string][]MetadataQueryAttributeValueTuple {
	rv := objc.Call[map[string][]MetadataQueryAttributeValueTuple](m_, objc.Sel("valueLists"))
	return rv
}

// An array containing hierarchical groups of query results. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1416579-groupedresults?language=objc
func (m_ MetadataQuery) GroupedResults() []MetadataQueryResultGroup {
	rv := objc.Call[[]MetadataQueryResultGroup](m_, objc.Sel("groupedResults"))
	return rv
}

// An array of sort descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411847-sortdescriptors?language=objc
func (m_ MetadataQuery) SortDescriptors() []SortDescriptor {
	rv := objc.Call[[]SortDescriptor](m_, objc.Sel("sortDescriptors"))
	return rv
}

// An array of sort descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411847-sortdescriptors?language=objc
func (m_ MetadataQuery) SetSortDescriptors(value []ISortDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("setSortDescriptors:"), value)
}

// An array containing the query’s results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1408872-results?language=objc
func (m_ MetadataQuery) Results() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("results"))
	return rv
}

// A Boolean value that indicates whether the query has stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411941-stopped?language=objc
func (m_ MetadataQuery) IsStopped() bool {
	rv := objc.Call[bool](m_, objc.Sel("isStopped"))
	return rv
}

// An array of objects that define the query’s scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411307-searchitems?language=objc
func (m_ MetadataQuery) SearchItems() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("searchItems"))
	return rv
}

// An array of objects that define the query’s scope. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411307-searchitems?language=objc
func (m_ MetadataQuery) SetSearchItems(value []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setSearchItems:"), value)
}

// An array of grouping attributes. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1409191-groupingattributes?language=objc
func (m_ MetadataQuery) GroupingAttributes() []string {
	rv := objc.Call[[]string](m_, objc.Sel("groupingAttributes"))
	return rv
}

// An array of grouping attributes. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1409191-groupingattributes?language=objc
func (m_ MetadataQuery) SetGroupingAttributes(value []string) {
	objc.Call[objc.Void](m_, objc.Sel("setGroupingAttributes:"), value)
}

// The query’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1413181-delegate?language=objc
func (m_ MetadataQuery) Delegate() MetadataQueryDelegateWrapper {
	rv := objc.Call[MetadataQueryDelegateWrapper](m_, objc.Sel("delegate"))
	return rv
}

// The query’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1413181-delegate?language=objc
func (m_ MetadataQuery) SetDelegate(value PMetadataQueryDelegate) {
	po0 := objc.WrapAsProtocol("NSMetadataQueryDelegate", value)
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), po0)
}

// The query’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1413181-delegate?language=objc
func (m_ MetadataQuery) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The predicate used to filter query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411478-predicate?language=objc
func (m_ MetadataQuery) Predicate() Predicate {
	rv := objc.Call[Predicate](m_, objc.Sel("predicate"))
	return rv
}

// The predicate used to filter query results. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1411478-predicate?language=objc
func (m_ MetadataQuery) SetPredicate(value IPredicate) {
	objc.Call[objc.Void](m_, objc.Sel("setPredicate:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the query has started. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1416780-started?language=objc
func (m_ MetadataQuery) IsStarted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isStarted"))
	return rv
}

// A Boolean value that indicates whether the receiver is in the initial gathering phase of the query. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1407539-gathering?language=objc
func (m_ MetadataQuery) IsGathering() bool {
	rv := objc.Call[bool](m_, objc.Sel("isGathering"))
	return rv
}

// An array containing the search scopes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1412155-searchscopes?language=objc
func (m_ MetadataQuery) SearchScopes() []objc.Object {
	rv := objc.Call[[]objc.Object](m_, objc.Sel("searchScopes"))
	return rv
}

// An array containing the search scopes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1412155-searchscopes?language=objc
func (m_ MetadataQuery) SetSearchScopes(value []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setSearchScopes:"), value)
}

// The number of results returned by the query. (read-only) [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1418315-resultcount?language=objc
func (m_ MetadataQuery) ResultCount() uint {
	rv := objc.Call[uint](m_, objc.Sel("resultCount"))
	return rv
}

// The queue on which query result notifications are posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1410953-operationqueue?language=objc
func (m_ MetadataQuery) OperationQueue() OperationQueue {
	rv := objc.Call[OperationQueue](m_, objc.Sel("operationQueue"))
	return rv
}

// The queue on which query result notifications are posted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquery/1410953-operationqueue?language=objc
func (m_ MetadataQuery) SetOperationQueue(value IOperationQueue) {
	objc.Call[objc.Void](m_, objc.Sel("setOperationQueue:"), objc.Ptr(value))
}
