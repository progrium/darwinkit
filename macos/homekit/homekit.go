//go:generate go run ../../generate/tools/genmod.go
package homekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework HomeKit
import "C"