//go:generate go run ../../generate/tools/genmod.go
package audiotoolbox

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AudioToolbox
import "C"
