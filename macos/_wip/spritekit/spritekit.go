//go:generate go run ../../generate/tools/genmod.go
package spritekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework SpriteKit
import "C"
