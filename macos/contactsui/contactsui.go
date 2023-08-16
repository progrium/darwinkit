//go:generate go run ../../generate/tools/genmod.go
package contactsui

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework ContactsUI
import "C"
