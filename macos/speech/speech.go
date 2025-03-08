//go:generate go run ../../generate/tools/genmod.go
package speech

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Speech
import "C"
