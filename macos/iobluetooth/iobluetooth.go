//go:generate go run ../../generate/tools/genmod.go
package iobluetooth

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework IOBluetooth
import "C"
