//go:generate go run ../../generate/tools/genmod.go
package coreaudiokit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreAudioKit
import "C"
