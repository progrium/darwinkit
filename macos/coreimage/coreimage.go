//go:generate go run ../../generate/tools/genmod.go
package coreimage

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreImage
import "C"
