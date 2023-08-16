//go:generate go run ../../generate/tools/genmod.go
package scenekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework SceneKit
import "C"
