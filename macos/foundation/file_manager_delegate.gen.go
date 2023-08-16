// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface a file manager's delegate uses to intervene during operations or if an error occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate?language=objc
type PFileManagerDelegate interface {
	// optional
	FileManagerShouldRemoveItemAtURL(fileManager FileManager, URL URL) bool
	HasFileManagerShouldRemoveItemAtURL() bool
}

// A delegate implementation builder for the [PFileManagerDelegate] protocol.
type FileManagerDelegate struct {
	_FileManagerShouldRemoveItemAtURL func(fileManager FileManager, URL URL) bool
}

func (di *FileManagerDelegate) HasFileManagerShouldRemoveItemAtURL() bool {
	return di._FileManagerShouldRemoveItemAtURL != nil
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (di *FileManagerDelegate) SetFileManagerShouldRemoveItemAtURL(f func(fileManager FileManager, URL URL) bool) {
	di._FileManagerShouldRemoveItemAtURL = f
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (di *FileManagerDelegate) FileManagerShouldRemoveItemAtURL(fileManager FileManager, URL URL) bool {
	return di._FileManagerShouldRemoveItemAtURL(fileManager, URL)
}

// A concrete type wrapper for the [PFileManagerDelegate] protocol.
type FileManagerDelegateWrapper struct {
	objc.Object
}

func (f_ FileManagerDelegateWrapper) HasFileManagerShouldRemoveItemAtURL() bool {
	return f_.RespondsToSelector(objc.Sel("fileManager:shouldRemoveItemAtURL:"))
}

// Asks the delegate whether the item at the specified URL should be deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsfilemanagerdelegate/1411918-filemanager?language=objc
func (f_ FileManagerDelegateWrapper) FileManagerShouldRemoveItemAtURL(fileManager IFileManager, URL IURL) bool {
	rv := objc.Call[bool](f_, objc.Sel("fileManager:shouldRemoveItemAtURL:"), objc.Ptr(fileManager), objc.Ptr(URL))
	return rv
}
