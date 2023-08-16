//go:generate go run ../../generate/tools/genmod.go
package coremidi

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreMIDI
import "C"
