//go:generate go run ../../generate/tools/genmod.go
package coreml

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreML
import "C"
