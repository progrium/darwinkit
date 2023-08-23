// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The IKImageEditPanelDataSource protocol describes the methods that an IKImageEditPanel object uses to access the contents of its data source object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpaneldatasource?language=objc
type PImageEditPanelDataSource interface {
	// optional
	SetImageImageProperties(image coregraphics.ImageRef, metaData foundation.Dictionary)
	HasSetImageImageProperties() bool

	// optional
	ThumbnailWithMaximumSize(size foundation.Size) coregraphics.ImageRef
	HasThumbnailWithMaximumSize() bool
}

// A concrete type wrapper for the [PImageEditPanelDataSource] protocol.
type ImageEditPanelDataSourceWrapper struct {
	objc.Object
}

func (i_ ImageEditPanelDataSourceWrapper) HasSetImageImageProperties() bool {
	return i_.RespondsToSelector(objc.Sel("setImage:imageProperties:"))
}

// Sets an image with the specified properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpaneldatasource/1503517-setimage?language=objc
func (i_ ImageEditPanelDataSourceWrapper) SetImageImageProperties(image coregraphics.ImageRef, metaData foundation.Dictionary) {
	objc.Call[objc.Void](i_, objc.Sel("setImage:imageProperties:"), image, metaData)
}

func (i_ ImageEditPanelDataSourceWrapper) HasThumbnailWithMaximumSize() bool {
	return i_.RespondsToSelector(objc.Sel("thumbnailWithMaximumSize:"))
}

// Returns a thumbnail image whose size is no larger than the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikimageeditpaneldatasource/1505271-thumbnailwithmaximumsize?language=objc
func (i_ ImageEditPanelDataSourceWrapper) ThumbnailWithMaximumSize(size foundation.Size) coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](i_, objc.Sel("thumbnailWithMaximumSize:"), size)
	return rv
}
