// auto-generated code, do not modify
package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var DiffableDataSourceSnapshotClass = _DiffableDataSourceSnapshotClass{objc.GetClass("NSDiffableDataSourceSnapshot")}

type _DiffableDataSourceSnapshotClass struct {
	objc.Class
}

type IDiffableDataSourceSnapshot interface {
	objc.IObject
	AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objc.IObject, sectionIdentifier objc.IObject)
	AppendItemsWithIdentifiers(identifiers []objc.IObject)
	NumberOfItemsInSection(sectionIdentifier objc.IObject) int
	IndexOfItemIdentifier(itemIdentifier objc.IObject) int
	IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int
	ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object
	SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object
	InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject)
	InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject)
	InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject)
	InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject)
	DeleteAllItems()
	DeleteItemsWithIdentifiers(identifiers []objc.IObject)
	DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject)
	MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject)
	MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject)
	MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject)
	ReloadItemsWithIdentifiers(identifiers []objc.IObject)
	ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject)
	NumberOfItems() int
	NumberOfSections() int
	ItemIdentifiers() []objc.Object
	SectionIdentifiers() []objc.Object
}

type DiffableDataSourceSnapshot struct {
	objc.Object
}

func MakeDiffableDataSourceSnapshot(ptr unsafe.Pointer) DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshot{
		Object: objc.MakeObject(ptr),
	}
}

func (dc _DiffableDataSourceSnapshotClass) Alloc() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](dc, objc.GetSelector("alloc"))
	return rv
}

func DiffableDataSourceSnapshot_Alloc() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.Alloc()
}

func (dc _DiffableDataSourceSnapshotClass) New() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](dc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewDiffableDataSourceSnapshot() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.New()
}

func DiffableDataSourceSnapshot_New() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.New()
}

func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := objc.CallMethod[DiffableDataSourceSnapshot](d_, objc.GetSelector("init"))
	return rv
}

func DiffableDataSourceSnapshot_Init() DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshotClass.Alloc().Init()
}

func (d_ DiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("appendSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objc.IObject, sectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("appendItemsWithIdentifiers:intoSectionWithIdentifier:"), identifiers, objc.ExtractPtr(sectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("appendItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("numberOfItemsInSection:"), objc.ExtractPtr(sectionIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("indexOfItemIdentifier:"), objc.ExtractPtr(itemIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objc.IObject) int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("indexOfSectionIdentifier:"), objc.ExtractPtr(sectionIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objc.IObject) []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.GetSelector("itemIdentifiersInSectionWithIdentifier:"), objc.ExtractPtr(sectionIdentifier))
	// TODO: convert slice items...
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objc.IObject) objc.Object {
	rv := objc.CallMethod[objc.Object](d_, objc.GetSelector("sectionIdentifierForSectionContainingItemIdentifier:"), objc.ExtractPtr(itemIdentifier))
	return rv
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("insertItemsWithIdentifiers:afterItemWithIdentifier:"), identifiers, objc.ExtractPtr(itemIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objc.IObject, itemIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("insertItemsWithIdentifiers:beforeItemWithIdentifier:"), identifiers, objc.ExtractPtr(itemIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("insertSectionsWithIdentifiers:afterSectionWithIdentifier:"), sectionIdentifiers, objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("insertSectionsWithIdentifiers:beforeSectionWithIdentifier:"), sectionIdentifiers, objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) DeleteAllItems() {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("deleteAllItems"))
}

func (d_ DiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("deleteItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("deleteSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveItemWithIdentifier:afterItemWithIdentifier:"), objc.ExtractPtr(fromIdentifier), objc.ExtractPtr(toIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objc.IObject, toIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveItemWithIdentifier:beforeItemWithIdentifier:"), objc.ExtractPtr(fromIdentifier), objc.ExtractPtr(toIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveSectionWithIdentifier:afterSectionWithIdentifier:"), objc.ExtractPtr(fromSectionIdentifier), objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objc.IObject, toSectionIdentifier objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("moveSectionWithIdentifier:beforeSectionWithIdentifier:"), objc.ExtractPtr(fromSectionIdentifier), objc.ExtractPtr(toSectionIdentifier))
}

func (d_ DiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("reloadItemsWithIdentifiers:"), identifiers)
}

func (d_ DiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.IObject) {
	objc.CallMethod[objc.Void](d_, objc.GetSelector("reloadSectionsWithIdentifiers:"), sectionIdentifiers)
}

func (d_ DiffableDataSourceSnapshot) NumberOfItems() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("numberOfItems"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) NumberOfSections() int {
	rv := objc.CallMethod[int](d_, objc.GetSelector("numberOfSections"))
	return rv
}

func (d_ DiffableDataSourceSnapshot) ItemIdentifiers() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.GetSelector("itemIdentifiers"))
	// TODO: convert slice items...
	return rv
}

func (d_ DiffableDataSourceSnapshot) SectionIdentifiers() []objc.Object {
	rv := objc.CallMethod[[]objc.Object](d_, objc.GetSelector("sectionIdentifiers"))
	// TODO: convert slice items...
	return rv
}
