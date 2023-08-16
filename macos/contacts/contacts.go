//go:generate go run ../../generate/tools/genmod.go
package contacts

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Contacts
import "C"
