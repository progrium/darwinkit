//go:generate go run ../../generate/tools/genmod.go
package avkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AVKit
import "C"
