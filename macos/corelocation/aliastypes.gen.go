// AUTO-GENERATED CODE, DO NOT MODIFY

package corelocation

import (
	"github.com/progrium/macdriver/macos/foundation"
)

// A block to be called when a geocoding request is complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/corelocation/clgeocodecompletionhandler?language=objc
type GeocodeCompletionHandler = func(placemarks []Placemark, error foundation.Error)
