//go:generate go run ../../generate/tools/genmod.go
package metalkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework MetalKit
import "C"
