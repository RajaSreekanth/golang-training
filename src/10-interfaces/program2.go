package main

import (
	"fmt"
	"math"
)

type IHaveArea interface {
	Area() float32
}

type IHavePerimeter interface {
	Perimeter() float32
}

type Dimension interface {
	IHaveArea
	IHavePerimeter
}

type Rect struct {
	length, width float32
}

func (r Rect) Area() float32 {
	return r.length * r.width
}

func (r Rect) Perimeter() float32 {
	return 2 * (r.length + r.width)
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (c Rect) String() string {
	return fmt.Sprintf("Rectangle {length = %v, width = %v, area = %v, perimeter = %v}\n", c.length, c.width, c.Area(), c.Perimeter())
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle {radius = %v, area = %v, perimeter = %v}\n", c.radius, c.Area(), c.Perimeter())
}
func main() {
	rect := Rect{10, 20}
	circle := Circle{15}
	/*
		printArea(rect)
		printArea(circle)
		printPerimeter(rect)
		printPerimeter(circle)
	*/
	printDimension(rect)
	printDimension(circle)

	fmt.Println(rect)
	fmt.Println(circle)
}

func printArea(o IHaveArea) {
	fmt.Println("Area = ", o.Area())
}

func printPerimeter(o IHavePerimeter) {
	fmt.Println("Perimeter = ", o.Perimeter())
}

func printDimension(d Dimension) {
	fmt.Println("Area = ", d.Area())
	fmt.Println("Perimeter=", d.Perimeter())
}
