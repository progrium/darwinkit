//go:generate go run ../../generate/tools/genmod.go
package corevideo

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreVideo
import "C"
