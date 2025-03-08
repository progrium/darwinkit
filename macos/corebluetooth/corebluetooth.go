//go:generate go run ../../generate/tools/genmod.go
package corebluetooth

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreBluetooth
import "C"