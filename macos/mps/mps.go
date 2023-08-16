//go:generate go run ../../generate/tools/genmod.go
package mps

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework MetalPerformanceShaders
import "C"
