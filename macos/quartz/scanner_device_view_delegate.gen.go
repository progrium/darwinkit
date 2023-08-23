// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKScannerDeviceViewDelegate protocol defines the delegate protocol that the IKScannerDeviceView delegate must conform to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate?language=objc
type PScannerDeviceViewDelegate interface {
	// optional
	ScannerDeviceViewDidScanToURLError(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)
	HasScannerDeviceViewDidScanToURLError() bool
}

// A delegate implementation builder for the [PScannerDeviceViewDelegate] protocol.
type ScannerDeviceViewDelegate struct {
	_ScannerDeviceViewDidScanToURLError func(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)
}

func (di *ScannerDeviceViewDelegate) HasScannerDeviceViewDidScanToURLError() bool {
	return di._ScannerDeviceViewDidScanToURLError != nil
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) SetScannerDeviceViewDidScanToURLError(f func(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error)) {
	di._ScannerDeviceViewDidScanToURLError = f
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (di *ScannerDeviceViewDelegate) ScannerDeviceViewDidScanToURLError(scannerDeviceView ScannerDeviceView, url foundation.URL, error foundation.Error) {
	di._ScannerDeviceViewDidScanToURLError(scannerDeviceView, url, error)
}

// A concrete type wrapper for the [PScannerDeviceViewDelegate] protocol.
type ScannerDeviceViewDelegateWrapper struct {
	objc.Object
}

func (s_ ScannerDeviceViewDelegateWrapper) HasScannerDeviceViewDidScanToURLError() bool {
	return s_.RespondsToSelector(objc.Sel("scannerDeviceView:didScanToURL:error:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceviewdelegate/1504768-scannerdeviceview?language=objc
func (s_ ScannerDeviceViewDelegateWrapper) ScannerDeviceViewDidScanToURLError(scannerDeviceView IScannerDeviceView, url foundation.IURL, error foundation.IError) {
	objc.Call[objc.Void](s_, objc.Sel("scannerDeviceView:didScanToURL:error:"), objc.Ptr(scannerDeviceView), objc.Ptr(url), objc.Ptr(error))
}
