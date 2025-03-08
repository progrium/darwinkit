//go:generate go run ../../generate/tools/genmod.go
package gamekit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework GameKit
import "C"