// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"github.com/progrium/macdriver/objc"
)

// A delegate protocol that defines the methods to implement to respond to asset-writing events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate?language=objc
type PAssetWriterDelegate interface {
	// optional
	AssetWriterDidOutputSegmentDataSegmentType(writer AssetWriter, segmentData []byte, segmentType AssetSegmentType)
	HasAssetWriterDidOutputSegmentDataSegmentType() bool
}

// A delegate implementation builder for the [PAssetWriterDelegate] protocol.
type AssetWriterDelegate struct {
	_AssetWriterDidOutputSegmentDataSegmentType func(writer AssetWriter, segmentData []byte, segmentType AssetSegmentType)
}

func (di *AssetWriterDelegate) HasAssetWriterDidOutputSegmentDataSegmentType() bool {
	return di._AssetWriterDidOutputSegmentDataSegmentType != nil
}

// Tells the delegate that the asset writer output segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate/3546592-assetwriter?language=objc
func (di *AssetWriterDelegate) SetAssetWriterDidOutputSegmentDataSegmentType(f func(writer AssetWriter, segmentData []byte, segmentType AssetSegmentType)) {
	di._AssetWriterDidOutputSegmentDataSegmentType = f
}

// Tells the delegate that the asset writer output segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate/3546592-assetwriter?language=objc
func (di *AssetWriterDelegate) AssetWriterDidOutputSegmentDataSegmentType(writer AssetWriter, segmentData []byte, segmentType AssetSegmentType) {
	di._AssetWriterDidOutputSegmentDataSegmentType(writer, segmentData, segmentType)
}

// A concrete type wrapper for the [PAssetWriterDelegate] protocol.
type AssetWriterDelegateWrapper struct {
	objc.Object
}

func (a_ AssetWriterDelegateWrapper) HasAssetWriterDidOutputSegmentDataSegmentType() bool {
	return a_.RespondsToSelector(objc.Sel("assetWriter:didOutputSegmentData:segmentType:"))
}

// Tells the delegate that the asset writer output segment data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterdelegate/3546592-assetwriter?language=objc
func (a_ AssetWriterDelegateWrapper) AssetWriterDidOutputSegmentDataSegmentType(writer IAssetWriter, segmentData []byte, segmentType AssetSegmentType) {
	objc.Call[objc.Void](a_, objc.Sel("assetWriter:didOutputSegmentData:segmentType:"), objc.Ptr(writer), segmentData, segmentType)
}
