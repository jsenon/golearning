package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// ... is a variadic function, is allowed to be called with multiple value

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Perim) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.ra
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

type Circle struct {
	x, y, ra float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.ra * c.ra
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Shape interface {
	area() float64
}

type Perim interface {
	perimeter() float64
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	r2 := Rectangle{0, 0, 5, 5}
	fmt.Println("Rectangle Area is:", r.area())
	fmt.Println("Rectangle2 Area is:", r2.area())
	fmt.Println("Circle Area is:", c.area())
	fmt.Println("Total Area is:", totalArea(&c, &r, &r2))
	fmt.Println("Rectangle Perimeter is:", r.perimeter())
	fmt.Println("Rectangle2 Perimeter is:", r2.perimeter())
	fmt.Println("Circle Perimeter is:", c.perimeter())
	fmt.Println("Total Perimeter is:", totalPerimeter(&c, &r, &r2))
}
