//go:generate go run ../../generate/tools/genmod.go
package coremedia

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreMedia
import "C"
