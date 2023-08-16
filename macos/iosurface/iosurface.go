//go:generate go run ../../generate/tools/genmod.go
package iosurface

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework IOSurface
import "C"
