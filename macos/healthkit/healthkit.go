//go:generate go run ../../generate/tools/genmod.go
package healthkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework HealthKit
import "C"