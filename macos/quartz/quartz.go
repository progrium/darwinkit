//go:generate go run ../../generate/tools/genmod.go
package quartz

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Quartz
import "C"
