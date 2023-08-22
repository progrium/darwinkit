//go:generate go run ../../generate/tools/genmod.go
package modelio

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework ModelIO
import "C"
