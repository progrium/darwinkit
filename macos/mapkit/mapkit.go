//go:generate go run ../../generate/tools/genmod.go
package mapkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework MapKit
import "C"
