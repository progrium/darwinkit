// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DiffableDataSourceSnapshot] class.
var DiffableDataSourceSnapshotClass = _DiffableDataSourceSnapshotClass{objc.GetClass("NSDiffableDataSourceSnapshot")}

type _DiffableDataSourceSnapshotClass struct {
	objc.Class
}

// An interface definition for the [DiffableDataSourceSnapshot] class.
type IDiffableDataSourceSnapshot interface {
	objc.IObject
	SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object
	ReloadItemsWithIdentifiers(identifiers []objc.IObject)
	AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject)
	IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int
	ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	NumberOfItemsInSection(sectionIdentifier objc.IObject) int
	DeleteItemsWithIdentifiers(identifiers []objc.IObject)
	DeleteAllItems()
	InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject)
	MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject)
	IndexOfItemIdentifier(itemIdentifier objc.IObject) int
	ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object
	InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject)
	AppendItemsWithIdentifiers(identifiers []objc.IObject)
	DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	SectionIdentifiers() []objc.Object
	NumberOfItems() int
	ItemIdentifiers() []objc.Object
	NumberOfSections() int
}

// A representation of the state of the data in a view at a specific point in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot?language=objc
type DiffableDataSourceSnapshot struct {
	objc.Object
}

func DiffableDataSourceSnapshotFrom(ptr unsafe.Pointer) DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshot{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DiffableDataSourceSnapshotClass) Alloc() DiffableDataSourceSnapshot {
	rv := objc.Call[DiffableDataSourceSnapshot](dc, objc.Sel("alloc"))
	return rv
}

func DiffableDataSourceSnapshot_Alloc() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.Alloc()
}

func (dc _DiffableDataSourceSnapshotClass) New() DiffableDataSourceSnapshot {
	rv := objc.Call[DiffableDataSourceSnapshot](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDiffableDataSourceSnapshot() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.New()
}

func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := objc.Call[DiffableDataSourceSnapshot](d_, objc.Sel("init"))
	return rv
}

// Returns the identifier of the section containing the specified item in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182918-sectionidentifierforsectionconta?language=objc
func (d_ DiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](d_, objc.Sel("sectionIdentifierForSectionContainingItemIdentifier:"), objc.Ptr(itemIdentifier))
	return rv
}

// Reloads the data within the specified items in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182916-reloaditemswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("reloadItemsWithIdentifiers:"), identifiers)
}

// Adds the sections with the specified identifiers to the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182897-appendsectionswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("appendSectionsWithIdentifiers:"), sectionIdentifiers)
}

// Moves the section from its current position in the snapshot to the position immediately after the specified section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182911-movesectionwithidentifier?language=objc
func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("moveSectionWithIdentifier:afterSectionWithIdentifier:"), objc.Ptr(fromSectionIdentifier), objc.Ptr(toSectionIdentifier))
}

// Returns the index of the section of the snapshot with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182902-indexofsectionidentifier?language=objc
func (d_ DiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int {
	rv := objc.Call[int](d_, objc.Sel("indexOfSectionIdentifier:"), objc.Ptr(sectionIdentifier))
	return rv
}

// Reloads the data within the specified sections of the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182917-reloadsectionswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("reloadSectionsWithIdentifiers:"), sectionIdentifiers)
}

// Returns the number of items in the specified section of the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182914-numberofitemsinsection?language=objc
func (d_ DiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objc.IObject) int {
	rv := objc.Call[int](d_, objc.Sel("numberOfItemsInSection:"), objc.Ptr(sectionIdentifier))
	return rv
}

// Deletes the items with the specified identifiers from the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182899-deleteitemswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("deleteItemsWithIdentifiers:"), identifiers)
}

// Deletes all of the items from the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182898-deleteallitems?language=objc
func (d_ DiffableDataSourceSnapshot) DeleteAllItems() {
	objc.Call[objc.Void](d_, objc.Sel("deleteAllItems"))
}

// Inserts the provided sections immediately after the section with the specified identifier in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182905-insertsectionswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("insertSectionsWithIdentifiers:afterSectionWithIdentifier:"), sectionIdentifiers, objc.Ptr(toSectionIdentifier))
}

// Moves the item from its current position in the snapshot to the position immediately after the specified item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182909-moveitemwithidentifier?language=objc
func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("moveItemWithIdentifier:afterItemWithIdentifier:"), objc.Ptr(fromIdentifier), objc.Ptr(toIdentifier))
}

// Returns the index of the item in the snapshot with the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182901-indexofitemidentifier?language=objc
func (d_ DiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objc.IObject) int {
	rv := objc.Call[int](d_, objc.Sel("indexOfItemIdentifier:"), objc.Ptr(itemIdentifier))
	return rv
}

// Returns the identifiers of all of the items in the specified section of the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182908-itemidentifiersinsectionwithiden?language=objc
func (d_ DiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("itemIdentifiersInSectionWithIdentifier:"), objc.Ptr(sectionIdentifier))
	return rv
}

// Inserts the provided items immediately after the item with the specified identifier in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182903-insertitemswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("insertItemsWithIdentifiers:afterItemWithIdentifier:"), identifiers, objc.Ptr(itemIdentifier))
}

// Adds the items with the specified identifiers to the last section of the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182895-appenditemswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("appendItemsWithIdentifiers:"), identifiers)
}

// Deletes the sections with the specified identifiers from the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182900-deletesectionswithidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("deleteSectionsWithIdentifiers:"), sectionIdentifiers)
}

// The identifiers of all of the sections in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182919-sectionidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) SectionIdentifiers() []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("sectionIdentifiers"))
	return rv
}

// The number of items in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182913-numberofitems?language=objc
func (d_ DiffableDataSourceSnapshot) NumberOfItems() int {
	rv := objc.Call[int](d_, objc.Sel("numberOfItems"))
	return rv
}

// The identifiers of all of the items in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182907-itemidentifiers?language=objc
func (d_ DiffableDataSourceSnapshot) ItemIdentifiers() []objc.Object {
	rv := objc.Call[[]objc.Object](d_, objc.Sel("itemIdentifiers"))
	return rv
}

// The number of sections in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdiffabledatasourcesnapshot/3182915-numberofsections?language=objc
func (d_ DiffableDataSourceSnapshot) NumberOfSections() int {
	rv := objc.Call[int](d_, objc.Sel("numberOfSections"))
	return rv
}
