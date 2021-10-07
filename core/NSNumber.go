package core

type NSNumber struct {
	gen_NSNumber
}

func NSNumber_WithBool(b bool) NSNumber {
	return NSNumber_numberWithBool_(b)
}
