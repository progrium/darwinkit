//go:generate go run ../../generate/tools/genmod.go
package gameplaykit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework GameplayKit
import "C"
