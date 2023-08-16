//go:generate go run ../../generate/tools/genmod.go
package corespotlight

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreSpotlight
import "C"
