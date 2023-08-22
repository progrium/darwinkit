//go:generate go run ../../generate/tools/genmod.go
package sysconfig

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework SystemConfiguration
import "C"
