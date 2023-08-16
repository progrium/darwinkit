//go:generate go run ../../generate/tools/genmod.go
package metal

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Metal
import "C"
