//go:generate go run ../../generate/tools/genmod.go
package javascriptcore

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework JavaScriptCore
import "C"