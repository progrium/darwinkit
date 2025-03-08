package pdfkit

// View display modes
const (
	DisplayModeSinglePage       = 0
	DisplayModeSinglePageContinuous = 1
	DisplayModeTwoUp             = 2
	DisplayModeTwoUpContinuous   = 3
)

// PDF auto scaling options
const (
	AutoScaleNone              = 0
	AutoScaleToFit             = 1
	AutoScaleToWidth           = 2
)

// PDF annotation key constants
const (
	AnnotationTextIconKey      = "IconType"
	AnnotationColorKey         = "color"
	AnnotationContentsKey      = "contents"
	AnnotationBorderKey        = "border"
	AnnotationRectKey          = "rect"
	AnnotationDateKey          = "date"
)

// PDF annotation text icons
const (
	TextAnnotationIconComment       = "Comment"
	TextAnnotationIconKey           = "Key"
	TextAnnotationIconNote          = "Note"
	TextAnnotationIconHelp          = "Help"
	TextAnnotationIconNewParagraph  = "NewParagraph"
	TextAnnotationIconParagraph     = "Paragraph"
	TextAnnotationIconInsert        = "Insert"
)

// PDF action types
const (
	ActionNamedGoTo     = "GoTo"
	ActionGoToRemote    = "GoToR"
	ActionGoToEmbedded  = "GoToE"
	ActionLaunch        = "Launch"
	ActionThread        = "Thread"
	ActionURI           = "URI"
	ActionSound         = "Sound"
	ActionMovie         = "Movie"
	ActionHide          = "Hide"
	ActionNamed         = "Named"
	ActionSubmitForm    = "SubmitForm"
	ActionResetForm     = "ResetForm"
	ActionJavaScript    = "JavaScript"
)