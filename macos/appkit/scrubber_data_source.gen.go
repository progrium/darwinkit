// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A set of methods that a scrubber data source object implements to provide items to the scrubber from an associated data collection in your app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdatasource?language=objc
type PScrubberDataSource interface {
	// optional
	ScrubberViewForItemAtIndex(scrubber Scrubber, index int) IScrubberItemView
	HasScrubberViewForItemAtIndex() bool

	// optional
	NumberOfItemsForScrubber(scrubber Scrubber) int
	HasNumberOfItemsForScrubber() bool
}

// A concrete type wrapper for the [PScrubberDataSource] protocol.
type ScrubberDataSourceWrapper struct {
	objc.Object
}

func (s_ ScrubberDataSourceWrapper) HasScrubberViewForItemAtIndex() bool {
	return s_.RespondsToSelector(objc.Sel("scrubber:viewForItemAtIndex:"))
}

// Asks the data source object for the view the corresponds to the specified item in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdatasource/2544709-scrubber?language=objc
func (s_ ScrubberDataSourceWrapper) ScrubberViewForItemAtIndex(scrubber IScrubber, index int) ScrubberItemView {
	rv := objc.Call[ScrubberItemView](s_, objc.Sel("scrubber:viewForItemAtIndex:"), objc.Ptr(scrubber), index)
	return rv
}

func (s_ ScrubberDataSourceWrapper) HasNumberOfItemsForScrubber() bool {
	return s_.RespondsToSelector(objc.Sel("numberOfItemsForScrubber:"))
}

// Asks the data source for the number of items in the scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberdatasource/2544845-numberofitemsforscrubber?language=objc
func (s_ ScrubberDataSourceWrapper) NumberOfItemsForScrubber(scrubber IScrubber) int {
	rv := objc.Call[int](s_, objc.Sel("numberOfItemsForScrubber:"), objc.Ptr(scrubber))
	return rv
}
