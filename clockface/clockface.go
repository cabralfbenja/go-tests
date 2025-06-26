package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

var center = Point{
	X: 150,
	Y: 150,
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale (length of the hand)
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + center.X, p.Y + center.Y}                 // translate (center)

	return p
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func secondHandPoint(t time.Time) Point {
	return getPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return getPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return getPoint(hoursInRadians(t))
}

func getPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
