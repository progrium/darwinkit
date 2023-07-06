package core

type NSNumber struct {
	gen_NSNumber
}

func NSNumber_WithBool(b bool) NSNumber {
	return NSNumber_NumberWithBool(b)
}

func NSNumber_WithInt(n int32) NSNumber {
	return NSNumber_NumberWithInt(n)
}
