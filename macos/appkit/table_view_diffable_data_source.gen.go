// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableViewDiffableDataSource] class.
var TableViewDiffableDataSourceClass = _TableViewDiffableDataSourceClass{objc.GetClass("NSTableViewDiffableDataSource")}

type _TableViewDiffableDataSourceClass struct {
	objc.Class
}

// An interface definition for the [TableViewDiffableDataSource] class.
type ITableViewDiffableDataSource interface {
	objc.IObject
	RowForItemIdentifier(identifier objc.IObject) int
	ItemIdentifierForRow(row int) objc.Object
	RowForSectionIdentifier(identifier objc.IObject) int
	SectionIdentifierForRow(row int) objc.Object
	RowViewProvider() TableViewDiffableDataSourceRowProvider
	SetRowViewProvider(value TableViewDiffableDataSourceRowProvider)
	SectionHeaderViewProvider() TableViewDiffableDataSourceSectionHeaderViewProvider
	SetSectionHeaderViewProvider(value TableViewDiffableDataSourceSectionHeaderViewProvider)
	DefaultRowAnimation() TableViewAnimationOptions
	SetDefaultRowAnimation(value TableViewAnimationOptions)
}

// The object you use to manage data and provide items for a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource?language=objc
type TableViewDiffableDataSource struct {
	objc.Object
}

func TableViewDiffableDataSourceFrom(ptr unsafe.Pointer) TableViewDiffableDataSource {
	return TableViewDiffableDataSource{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TableViewDiffableDataSource) InitWithTableViewCellProvider(tableView ITableView, cellProvider TableViewDiffableDataSourceCellProvider) TableViewDiffableDataSource {
	rv := objc.Call[TableViewDiffableDataSource](t_, objc.Sel("initWithTableView:cellProvider:"), objc.Ptr(tableView), cellProvider)
	return rv
}

// Creates a diffable data source with the specified cell provider, and connects it to the specified table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553215-initwithtableview?language=objc
func TableViewDiffableDataSource_InitWithTableViewCellProvider(tableView ITableView, cellProvider TableViewDiffableDataSourceCellProvider) TableViewDiffableDataSource {
	return TableViewDiffableDataSourceClass.Alloc().InitWithTableViewCellProvider(tableView, cellProvider)
}

func (tc _TableViewDiffableDataSourceClass) Alloc() TableViewDiffableDataSource {
	rv := objc.Call[TableViewDiffableDataSource](tc, objc.Sel("alloc"))
	return rv
}

func TableViewDiffableDataSource_Alloc() TableViewDiffableDataSource {
	return TableViewDiffableDataSourceClass.Alloc()
}

func (tc _TableViewDiffableDataSourceClass) New() TableViewDiffableDataSource {
	rv := objc.Call[TableViewDiffableDataSource](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableViewDiffableDataSource() TableViewDiffableDataSource {
	return TableViewDiffableDataSourceClass.New()
}

func (t_ TableViewDiffableDataSource) Init() TableViewDiffableDataSource {
	rv := objc.Call[TableViewDiffableDataSource](t_, objc.Sel("init"))
	return rv
}

// Returns a row for the item with the specified identifier in the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553217-rowforitemidentifier?language=objc
func (t_ TableViewDiffableDataSource) RowForItemIdentifier(identifier objc.IObject) int {
	rv := objc.Call[int](t_, objc.Sel("rowForItemIdentifier:"), objc.Ptr(identifier))
	return rv
}

// Returns an identifier for the item at the specified row in the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553216-itemidentifierforrow?language=objc
func (t_ TableViewDiffableDataSource) ItemIdentifierForRow(row int) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("itemIdentifierForRow:"), row)
	return rv
}

// Returns a row for the section with the specified identifier in the table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553218-rowforsectionidentifier?language=objc
func (t_ TableViewDiffableDataSource) RowForSectionIdentifier(identifier objc.IObject) int {
	rv := objc.Call[int](t_, objc.Sel("rowForSectionIdentifier:"), objc.Ptr(identifier))
	return rv
}

// Returns the identifier of the section containing the specified row in the snapshot. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553221-sectionidentifierforrow?language=objc
func (t_ TableViewDiffableDataSource) SectionIdentifierForRow(row int) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("sectionIdentifierForRow:"), row)
	return rv
}

// The closure that configures and returns the table view’s row views from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553219-rowviewprovider?language=objc
func (t_ TableViewDiffableDataSource) RowViewProvider() TableViewDiffableDataSourceRowProvider {
	rv := objc.Call[TableViewDiffableDataSourceRowProvider](t_, objc.Sel("rowViewProvider"))
	return rv
}

// The closure that configures and returns the table view’s row views from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553219-rowviewprovider?language=objc
func (t_ TableViewDiffableDataSource) SetRowViewProvider(value TableViewDiffableDataSourceRowProvider) {
	objc.Call[objc.Void](t_, objc.Sel("setRowViewProvider:"), value)
}

// The closure that configures and returns the table view’s section header views from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553220-sectionheaderviewprovider?language=objc
func (t_ TableViewDiffableDataSource) SectionHeaderViewProvider() TableViewDiffableDataSourceSectionHeaderViewProvider {
	rv := objc.Call[TableViewDiffableDataSourceSectionHeaderViewProvider](t_, objc.Sel("sectionHeaderViewProvider"))
	return rv
}

// The closure that configures and returns the table view’s section header views from the diffable data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553220-sectionheaderviewprovider?language=objc
func (t_ TableViewDiffableDataSource) SetSectionHeaderViewProvider(value TableViewDiffableDataSourceSectionHeaderViewProvider) {
	objc.Call[objc.Void](t_, objc.Sel("setSectionHeaderViewProvider:"), value)
}

// The default animation the UI uses to show differences between rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553214-defaultrowanimation?language=objc
func (t_ TableViewDiffableDataSource) DefaultRowAnimation() TableViewAnimationOptions {
	rv := objc.Call[TableViewAnimationOptions](t_, objc.Sel("defaultRowAnimation"))
	return rv
}

// The default animation the UI uses to show differences between rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdiffabledatasource/3553214-defaultrowanimation?language=objc
func (t_ TableViewDiffableDataSource) SetDefaultRowAnimation(value TableViewAnimationOptions) {
	objc.Call[objc.Void](t_, objc.Sel("setDefaultRowAnimation:"), value)
}
