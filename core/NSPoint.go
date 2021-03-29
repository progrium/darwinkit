package core

import "fmt"

type NSPoint struct {
	X float64
	Y float64
}

func (pt NSPoint) String() string {
	return fmt.Sprintf("(%v, %v)", pt.X, pt.Y)
}
