//go:generate go run ../../generate/tools/genmod.go
package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
// void importWebKitProtocols();
import "C"

func init() {
	C.importWebKitProtocols()
}
