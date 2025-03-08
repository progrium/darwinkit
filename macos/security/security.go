//go:generate go run ../../generate/tools/genmod.go
package security

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Security
import "C"