// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a table view uses to provide data to a table view and to allow the editing of the table view's data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdatasource?language=objc
type PTableViewDataSource interface {
	// optional
	NumberOfRowsInTableView(tableView TableView) int
	HasNumberOfRowsInTableView() bool

	// optional
	TableViewSortDescriptorsDidChange(tableView TableView, oldDescriptors []foundation.SortDescriptor)
	HasTableViewSortDescriptorsDidChange() bool
}

// A concrete type wrapper for the [PTableViewDataSource] protocol.
type TableViewDataSourceWrapper struct {
	objc.Object
}

func (t_ TableViewDataSourceWrapper) HasNumberOfRowsInTableView() bool {
	return t_.RespondsToSelector(objc.Sel("numberOfRowsInTableView:"))
}

// Returns the number of records managed for aTableView by the data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdatasource/1524583-numberofrowsintableview?language=objc
func (t_ TableViewDataSourceWrapper) NumberOfRowsInTableView(tableView ITableView) int {
	rv := objc.Call[int](t_, objc.Sel("numberOfRowsInTableView:"), objc.Ptr(tableView))
	return rv
}

func (t_ TableViewDataSourceWrapper) HasTableViewSortDescriptorsDidChange() bool {
	return t_.RespondsToSelector(objc.Sel("tableView:sortDescriptorsDidChange:"))
}

// Called by aTableView to indicate that sorting may need to be done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableviewdatasource/1532935-tableview?language=objc
func (t_ TableViewDataSourceWrapper) TableViewSortDescriptorsDidChange(tableView ITableView, oldDescriptors []foundation.ISortDescriptor) {
	objc.Call[objc.Void](t_, objc.Sel("tableView:sortDescriptorsDidChange:"), objc.Ptr(tableView), oldDescriptors)
}
