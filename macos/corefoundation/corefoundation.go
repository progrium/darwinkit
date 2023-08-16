//go:generate go run ../../generate/tools/genmod.go
package corefoundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreFoundation
import "C"
