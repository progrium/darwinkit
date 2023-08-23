// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

// Constants that describe the capture view’s supported controls styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avcaptureviewcontrolsstyle?language=objc
type CaptureViewControlsStyle int

const (
	CaptureViewControlsStyleDefault               CaptureViewControlsStyle = 0
	CaptureViewControlsStyleFloating              CaptureViewControlsStyle = 1
	CaptureViewControlsStyleInline                CaptureViewControlsStyle = 0
	CaptureViewControlsStyleInlineDeviceSelection CaptureViewControlsStyle = 2
)

// Constants that indicate which user interface controls the view displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewcontrolsstyle?language=objc
type PlayerViewControlsStyle int

const (
	PlayerViewControlsStyleDefault  PlayerViewControlsStyle = 1
	PlayerViewControlsStyleFloating PlayerViewControlsStyle = 2
	PlayerViewControlsStyleInline   PlayerViewControlsStyle = 1
	PlayerViewControlsStyleMinimal  PlayerViewControlsStyle = 3
	PlayerViewControlsStyleNone     PlayerViewControlsStyle = 0
)

// Constants that specify an action a user takes when trimming media in a player view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewtrimresult?language=objc
type PlayerViewTrimResult int

const (
	PlayerViewTrimCancelButton PlayerViewTrimResult = 1
	PlayerViewTrimOKButton     PlayerViewTrimResult = 0
)

// Constants that describe the available button states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewbuttonstate?language=objc
type RoutePickerViewButtonState int

const (
	RoutePickerViewButtonStateActive            RoutePickerViewButtonState = 2
	RoutePickerViewButtonStateActiveHighlighted RoutePickerViewButtonState = 3
	RoutePickerViewButtonStateNormal            RoutePickerViewButtonState = 0
	RoutePickerViewButtonStateNormalHighlighted RoutePickerViewButtonState = 1
)
