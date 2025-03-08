//go:generate go run ../../generate/tools/genmod.go
package usernotifications

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework UserNotifications
import "C"