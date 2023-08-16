//go:generate go run ../../generate/tools/genmod.go
package corelocation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreLocation
import "C"
