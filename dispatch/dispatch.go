// Execute code concurrently on multicore hardware by submitting work to dispatch queues managed by the system.
//
// [Apple Documentation]
//
// [AppleDocumentation]: https://developer.apple.com/documentation/dispatch?language=objc
package dispatch

// #cgo CFLAGS: -x objective-c
import "C"
import "unsafe"

type Data unsafe.Pointer
