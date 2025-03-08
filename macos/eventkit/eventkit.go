//go:generate go run ../../generate/tools/genmod.go
package eventkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework EventKit
import "C"
