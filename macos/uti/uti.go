//go:generate go run ../../generate/tools/genmod.go
package uti

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework UniformTypeIdentifiers
import "C"
