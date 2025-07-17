//go:generate go run ../../generate/tools/genmod.go
package virtualization

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Virtualization
import "C"
