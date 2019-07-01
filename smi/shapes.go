package smi

import "math"

// Shape are things shapes needo to do
type Shape interface {
	Area() float64
}

// Rectangle is a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Triangle is a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Perimiter calculates the distance around a rectangle
func Perimiter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area calculates the area of a Circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Area calculates the area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
