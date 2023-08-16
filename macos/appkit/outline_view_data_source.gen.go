// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of methods that an outline view calls to retrieve data and information about it from the data source delegate, and—optionally—to update data values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource?language=objc
type POutlineViewDataSource interface {
	// optional
	OutlineViewSortDescriptorsDidChange(outlineView OutlineView, oldDescriptors []foundation.SortDescriptor)
	HasOutlineViewSortDescriptorsDidChange() bool
}

// A concrete type wrapper for the [POutlineViewDataSource] protocol.
type OutlineViewDataSourceWrapper struct {
	objc.Object
}

func (o_ OutlineViewDataSourceWrapper) HasOutlineViewSortDescriptorsDidChange() bool {
	return o_.RespondsToSelector(objc.Sel("outlineView:sortDescriptorsDidChange:"))
}

// Invoked by an outline view to notify the data source that the descriptors changed and the data may need to be resorted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineviewdatasource/1535892-outlineview?language=objc
func (o_ OutlineViewDataSourceWrapper) OutlineViewSortDescriptorsDidChange(outlineView IOutlineView, oldDescriptors []foundation.ISortDescriptor) {
	objc.Call[objc.Void](o_, objc.Sel("outlineView:sortDescriptorsDidChange:"), objc.Ptr(outlineView), oldDescriptors)
}
