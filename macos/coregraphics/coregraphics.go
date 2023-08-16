//go:generate go run ../../generate/tools/genmod.go
package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
import "C"
