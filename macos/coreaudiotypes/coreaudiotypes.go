//go:generate go run ../../generate/tools/genmod.go
package coreaudiotypes

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreAudio
import "C"
