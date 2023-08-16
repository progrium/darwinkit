//go:generate go run ../../generate/tools/genmod.go
package fileprovider

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework FileProvider
import "C"
