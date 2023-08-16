//go:generate go run ../../generate/tools/genmod.go
package coredata

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreData
import "C"
