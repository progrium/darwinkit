//go:generate go run ../../generate/tools/genmod.go
package avfoundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AVFoundation
import "C"
