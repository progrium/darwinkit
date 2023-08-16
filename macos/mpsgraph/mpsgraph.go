//go:generate go run ../../generate/tools/genmod.go
package mpsgraph

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework MetalPerformanceShadersGraph
import "C"
