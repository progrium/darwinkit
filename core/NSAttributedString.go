package core

import (
	"github.com/progrium/macdriver/objc"
)

// Wrapper for NSAttributedString
// https://developer.apple.com/documentation/foundation/nsattributedstring?language=objc
type NSAttributedString struct {
	gen_NSAttributedString
}

// NSAttributedString_FromString returns an initialized NSAttributedString
// https://developer.apple.com/documentation/foundation/nsattributedstring/1407481-initwithstring?language=objc
func NSAttributedString_FromString(str string) NSAttributedString {
	nsstr := NSString_FromString(str)
	return NSAttributedString_Alloc().InitWithString(nsstr)
}

func NSAttributedString_FromObject(obj objc.Object) NSAttributedString {
	return NSAttributedString_FromRef(obj)
}
