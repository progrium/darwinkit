package core

import "fmt"

type NSMutableString struct {
	gen_NSMutableString
}

// Needed because it uses printf style variable arguments
func (s NSMutableString) AppendFormat(format string, args ... interface{}) {
	theString := fmt.Sprintf(format, args...)
	s.AppendString(theString)
}
