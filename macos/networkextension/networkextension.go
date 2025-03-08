//go:generate go run ../../generate/tools/genmod.go
package networkextension

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework NetworkExtension
import "C"