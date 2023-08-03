//go:generate go run ../../generate/tools/genmod.go
package uniformtypeidentifiers

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework UniformTypeIdentifiers -framework Foundation
import "C"
