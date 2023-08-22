// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureVideoPreviewLayer] class.
var CaptureVideoPreviewLayerClass = _CaptureVideoPreviewLayerClass{objc.GetClass("AVCaptureVideoPreviewLayer")}

type _CaptureVideoPreviewLayerClass struct {
	objc.Class
}

// An interface definition for the [CaptureVideoPreviewLayer] class.
type ICaptureVideoPreviewLayer interface {
	quartzcore.ILayer
	TransformedMetadataObjectForMetadataObject(metadataObject IMetadataObject) MetadataObject
	MetadataOutputRectOfInterestForRect(rectInLayerCoordinates coregraphics.Rect) coregraphics.Rect
	CaptureDevicePointOfInterestForPoint(pointInLayer coregraphics.Point) coregraphics.Point
	SetSessionWithNoConnection(session ICaptureSession)
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.Rect) coregraphics.Rect
	PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest coregraphics.Point) coregraphics.Point
	VideoGravity() LayerVideoGravity
	SetVideoGravity(value LayerVideoGravity)
	Session() CaptureSession
	SetSession(value ICaptureSession)
	Connection() CaptureConnection
}

// A Core Animation layer that displays video from a camera device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer?language=objc
type CaptureVideoPreviewLayer struct {
	quartzcore.Layer
}

func CaptureVideoPreviewLayerFrom(ptr unsafe.Pointer) CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayer{
		Layer: quartzcore.LayerFrom(ptr),
	}
}

func (c_ CaptureVideoPreviewLayer) InitWithSessionWithNoConnection(session ICaptureSession) CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("initWithSessionWithNoConnection:"), objc.Ptr(session))
	return rv
}

// Creates a layer to preview the visual output of a capture session, without making connections to eligible video inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1387426-initwithsessionwithnoconnection?language=objc
func NewCaptureVideoPreviewLayerWithSessionWithNoConnection(session ICaptureSession) CaptureVideoPreviewLayer {
	instance := CaptureVideoPreviewLayerClass.Alloc().InitWithSessionWithNoConnection(session)
	instance.Autorelease()
	return instance
}

func (cc _CaptureVideoPreviewLayerClass) LayerWithSessionWithNoConnection(session ICaptureSession) CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](cc, objc.Sel("layerWithSessionWithNoConnection:"), objc.Ptr(session))
	return rv
}

// Returns a new layer to preview the visual output of a capture session, without making connections to eligible video inputs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1567197-layerwithsessionwithnoconnection?language=objc
func CaptureVideoPreviewLayer_LayerWithSessionWithNoConnection(session ICaptureSession) CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayerClass.LayerWithSessionWithNoConnection(session)
}

func (cc _CaptureVideoPreviewLayerClass) LayerWithSession(session ICaptureSession) CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](cc, objc.Sel("layerWithSession:"), objc.Ptr(session))
	return rv
}

// Returns a new layer to preview the visual output of a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1567198-layerwithsession?language=objc
func CaptureVideoPreviewLayer_LayerWithSession(session ICaptureSession) CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayerClass.LayerWithSession(session)
}

func (c_ CaptureVideoPreviewLayer) InitWithSession(session ICaptureSession) CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("initWithSession:"), objc.Ptr(session))
	return rv
}

// Creates a layer to preview the visual output of a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1387766-initwithsession?language=objc
func NewCaptureVideoPreviewLayerWithSession(session ICaptureSession) CaptureVideoPreviewLayer {
	instance := CaptureVideoPreviewLayerClass.Alloc().InitWithSession(session)
	instance.Autorelease()
	return instance
}

func (cc _CaptureVideoPreviewLayerClass) Alloc() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](cc, objc.Sel("alloc"))
	return rv
}

func CaptureVideoPreviewLayer_Alloc() CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayerClass.Alloc()
}

func (cc _CaptureVideoPreviewLayerClass) New() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureVideoPreviewLayer() CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayerClass.New()
}

func (c_ CaptureVideoPreviewLayer) Init() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureVideoPreviewLayerClass) Layer() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](cc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func CaptureVideoPreviewLayer_Layer() CaptureVideoPreviewLayer {
	return CaptureVideoPreviewLayerClass.Layer()
}

func (c_ CaptureVideoPreviewLayer) InitWithLayer(layer objc.IObject) CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewCaptureVideoPreviewLayerWithLayer(layer objc.IObject) CaptureVideoPreviewLayer {
	instance := CaptureVideoPreviewLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (c_ CaptureVideoPreviewLayer) ModelLayer() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func CaptureVideoPreviewLayer_ModelLayer() CaptureVideoPreviewLayer {
	instance := CaptureVideoPreviewLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (c_ CaptureVideoPreviewLayer) PresentationLayer() CaptureVideoPreviewLayer {
	rv := objc.Call[CaptureVideoPreviewLayer](c_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func CaptureVideoPreviewLayer_PresentationLayer() CaptureVideoPreviewLayer {
	instance := CaptureVideoPreviewLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// Converts a metadata object’s visual properties to layer coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1623501-transformedmetadataobjectformeta?language=objc
func (c_ CaptureVideoPreviewLayer) TransformedMetadataObjectForMetadataObject(metadataObject IMetadataObject) MetadataObject {
	rv := objc.Call[MetadataObject](c_, objc.Sel("transformedMetadataObjectForMetadataObject:"), objc.Ptr(metadataObject))
	return rv
}

// Converts a rectangle from layer coordinates to the coordinate space of the metadata output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1623495-metadataoutputrectofinterestforr?language=objc
func (c_ CaptureVideoPreviewLayer) MetadataOutputRectOfInterestForRect(rectInLayerCoordinates coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInLayerCoordinates)
	return rv
}

// Converts a point from layer coordinates to the coordinate space of the capture device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1623497-capturedevicepointofinterestforp?language=objc
func (c_ CaptureVideoPreviewLayer) CaptureDevicePointOfInterestForPoint(pointInLayer coregraphics.Point) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("captureDevicePointOfInterestForPoint:"), pointInLayer)
	return rv
}

// Associates a session with the layer without automatically forming a connection to an eligible input port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1390387-setsessionwithnoconnection?language=objc
func (c_ CaptureVideoPreviewLayer) SetSessionWithNoConnection(session ICaptureSession) {
	objc.Call[objc.Void](c_, objc.Sel("setSessionWithNoConnection:"), objc.Ptr(session))
}

// Converts a rectangle from metadata output coordinates to the coordinate space of the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1623498-rectformetadataoutputrectofinter?language=objc
func (c_ CaptureVideoPreviewLayer) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](c_, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return rv
}

// Converts a point from the coordinate space of the capture device to the coordinate space of the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1623502-pointforcapturedevicepointofinte?language=objc
func (c_ CaptureVideoPreviewLayer) PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest coregraphics.Point) coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("pointForCaptureDevicePointOfInterest:"), captureDevicePointOfInterest)
	return rv
}

// A value that indicates how the layer displays video content within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1386708-videogravity?language=objc
func (c_ CaptureVideoPreviewLayer) VideoGravity() LayerVideoGravity {
	rv := objc.Call[LayerVideoGravity](c_, objc.Sel("videoGravity"))
	return rv
}

// A value that indicates how the layer displays video content within its bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1386708-videogravity?language=objc
func (c_ CaptureVideoPreviewLayer) SetVideoGravity(value LayerVideoGravity) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoGravity:"), value)
}

// A capture session with visual output to preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1386157-session?language=objc
func (c_ CaptureVideoPreviewLayer) Session() CaptureSession {
	rv := objc.Call[CaptureSession](c_, objc.Sel("session"))
	return rv
}

// A capture session with visual output to preview. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1386157-session?language=objc
func (c_ CaptureVideoPreviewLayer) SetSession(value ICaptureSession) {
	objc.Call[objc.Void](c_, objc.Sel("setSession:"), objc.Ptr(value))
}

// An object that describes the connection from the layer to a particular input port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideopreviewlayer/1390893-connection?language=objc
func (c_ CaptureVideoPreviewLayer) Connection() CaptureConnection {
	rv := objc.Call[CaptureConnection](c_, objc.Sel("connection"))
	return rv
}
