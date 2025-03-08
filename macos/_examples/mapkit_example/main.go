package main

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/corelocation"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/mapkit"
	"github.com/progrium/darwinkit/objc"
)

func main() {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// Create a window
		frame := foundation.Rect{Size: foundation.Size{800, 600}}
		window := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask|appkit.ResizableWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&window)
		window.SetTitle(foundation.String_StringWithString("MapKit Example"))
		window.Center()

		// Create a map view
		mapView := mapkit.NewMapView()
		objc.Retain(&mapView)

		// Set map properties
		mapView.SetMapType(mapkit.MapTypeStandard)
		mapView.SetShowsUserLocation(true)

		// Create a location coordinate for San Francisco
		coordinate := corelocation.Coordinate{
			Latitude:  37.7749,
			Longitude: -122.4194,
		}

		// Create a region with a 5-mile radius
		region := mapkit.MapRegion{
			Center: coordinate,
			Span: mapkit.MapSpan{
				LatitudeDelta:  0.1,
				LongitudeDelta: 0.1,
			},
		}

		// Set the map's region
		mapView.SetRegion(region)

		// Create an annotation
		annotation := mapkit.NewPointAnnotation()
		annotation.SetCoordinate(coordinate)
		annotation.SetTitle(foundation.String_StringWithString("San Francisco"))
		annotation.SetSubtitle(foundation.String_StringWithString("California"))

		// Add the annotation to the map
		mapView.AddAnnotation(annotation)

		// Set the map view as the content view
		window.SetContentView(mapView)
		window.MakeKeyAndOrderFront(window)

		// Close app when window is closed
		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}