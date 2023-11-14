package appkit

// Variables of type NSModalSession point to information used by the system between NSApplication’s [appkit/nsapplication/beginmodalsessionforwindow] and [appkit/nsapplication/endmodalsession] messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmodalsession?language=objc
type ModalSession uintptr

// The inset distances for views, taking the user interface layout direction into account. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nsdirectionaledgeinsets?language=objc
type DirectionalEdgeInsets struct {
	Top      float64
	Leading  float64
	Bottom   float64
	Trailing float64
}
