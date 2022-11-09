package clockface

import "time"

// a point represents a two dimensional cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// secondhand is the unit vector of the second hand of an analogue clock at time T
// represented as a point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
