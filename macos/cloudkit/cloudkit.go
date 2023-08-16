//go:generate go run ../../generate/tools/genmod.go
package cloudkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CloudKit
import "C"
