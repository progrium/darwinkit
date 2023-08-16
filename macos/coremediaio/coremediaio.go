//go:generate go run ../../generate/tools/genmod.go
package coremediaio

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreMediaIO
import "C"
