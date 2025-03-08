//go:generate go run ../../generate/tools/genmod.go
package pdfkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework PDFKit
import "C"
