//go:generate go run ../../generate/tools/genmod.go
package vision

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Vision
import "C"
