//go:generate go run ../../generate/tools/genmod.go
package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
import "C"

func init() {
	// #import <stdlib.h>
	// void importFoundationProtocols();
	//C.importFoundationProtocols()
}
