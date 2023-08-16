//go:generate go run ../../generate/tools/genmod.go
package quartzcore

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework QuartzCore
import "C"

func init() {
	// void importQuartzCoreProtocols();
	// C.importQuartzCoreProtocols()
}
