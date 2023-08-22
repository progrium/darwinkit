//go:generate go run ../../generate/tools/genmod.go
package mediaplayer

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework MediaPlayer
import "C"
