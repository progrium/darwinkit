//go:generate go run ../../generate/tools/genmod.go
package screencapturekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework ScreenCaptureKit
import "C"
