package dispatch

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -l System
import "C"
import "unsafe"

type Data unsafe.Pointer
