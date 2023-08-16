// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"github.com/progrium/macdriver/macos/foundation"
)

// A block that processes the outcome of a permissions request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckapplicationpermissionblock?language=objc
type ApplicationPermissionBlock = func(applicationPermissionStatus ApplicationPermissionStatus, error foundation.Error)
