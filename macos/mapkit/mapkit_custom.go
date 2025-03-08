package mapkit

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/corelocation"
)

// MapView represents a view that displays a map interface
type MapView struct {
	objc.Object
}

// MapViewClass is the class instance for MapView
var MapViewClass = objc.GetClass("MKMapView")

// NewMapView creates a new MKMapView instance
func NewMapView() MapView {
	return MapView{objc.Call[objc.Object](MapViewClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetShowsUserLocation sets whether to show the user's location on the map
func (m MapView) SetShowsUserLocation(showsUserLocation bool) {
	m.Send(objc.Sel("setShowsUserLocation:"), showsUserLocation)
}

// ShowsUserLocation returns whether the map shows the user's location
func (m MapView) ShowsUserLocation() bool {
	return objc.Call[bool](m, objc.Sel("showsUserLocation"))
}

// SetDelegate sets the delegate for the map view
func (m MapView) SetDelegate(delegate objc.Object) {
	m.Send(objc.Sel("setDelegate:"), delegate)
}

// Annotation represents a map annotation
type Annotation struct {
	objc.Object
}

// AnnotationClass is the class instance for Annotation
var AnnotationClass = objc.GetClass("MKAnnotation")

// PointAnnotation represents a point annotation on a map
type PointAnnotation struct {
	objc.Object
}

// PointAnnotationClass is the class instance for PointAnnotation
var PointAnnotationClass = objc.GetClass("MKPointAnnotation")

// NewPointAnnotation creates a new MKPointAnnotation instance
func NewPointAnnotation() PointAnnotation {
	return PointAnnotation{objc.Call[objc.Object](PointAnnotationClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetCoordinate sets the coordinate for the point annotation
func (p PointAnnotation) SetCoordinate(coordinate corelocation.Coordinate) {
	p.Send(objc.Sel("setCoordinate:"), coordinate)
}

// SetTitle sets the title for the annotation
func (p PointAnnotation) SetTitle(title string) {
	p.Send(objc.Sel("setTitle:"), foundation.String_StringWithString(title))
}

// SetSubtitle sets the subtitle for the annotation
func (p PointAnnotation) SetSubtitle(subtitle string) {
	p.Send(objc.Sel("setSubtitle:"), foundation.String_StringWithString(subtitle))
}

// AddAnnotation adds an annotation to the map view
func (m MapView) AddAnnotation(annotation objc.Object) {
	m.Send(objc.Sel("addAnnotation:"), annotation)
}

// Region represents a map region
type MapRegion struct {
	Center corelocation.Coordinate
	Span   MapSpan
}

// MapSpan represents the span of a map region
type MapSpan struct {
	LatitudeDelta  float64
	LongitudeDelta float64
}

// SetRegion sets the map's region
func (m MapView) SetRegion(region MapRegion) {
	m.Send(objc.Sel("setRegion:"), region)
}

// MapType constants define the type of map to display
const (
	MapTypeStandard    = 0
	MapTypeSatellite   = 1
	MapTypeHybrid      = 2
	MapTypeSatelliteFlyover = 3
	MapTypeHybridFlyover = 4
	MapTypeMutedStandard = 5
)

// SetMapType sets the type of map to display
func (m MapView) SetMapType(mapType int) {
	m.Send(objc.Sel("setMapType:"), mapType)
}
