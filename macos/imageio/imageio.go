//go:generate go run ../../generate/tools/genmod.go
package imageio

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework ImageIO
import "C"
