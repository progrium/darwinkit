//go:generate go run ../../generate/tools/genmod.go
package coreservices

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreServices
import "C"