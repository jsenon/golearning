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
		fmt.Println("Im in total")

	}
	return area
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

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	r2 := Rectangle{0, 0, 5, 5}
	fmt.Println("Rectangle Area is:", r.area())
	fmt.Println("Rectangle2 Area is:", r2.area())
	fmt.Println("Circle Area is:", c.area())
	fmt.Println("Total Area is:", totalArea(&c, &r, &r2))

}
