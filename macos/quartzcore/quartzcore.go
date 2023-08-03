//go:generate go run ../../generate/tools/genmod.go
package quartzcore

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework QuartzCore
// void importQuartzCoreProtocols();
import "C"

func init() {
	C.importQuartzCoreProtocols()
}
