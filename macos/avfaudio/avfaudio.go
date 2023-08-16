//go:generate go run ../../generate/tools/genmod.go
package avfaudio

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AVFAudio
import "C"
