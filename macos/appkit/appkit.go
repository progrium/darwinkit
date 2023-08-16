//go:generate go run ../../generate/tools/genmod.go
package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// void importAppKitProtocols();
import "C"

func init() {
	C.importAppKitProtocols()
}
