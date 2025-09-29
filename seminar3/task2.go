package main

import "fmt"

type Shape struct {
	Name   string
	length int
	width  int
	radius float64
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	if s.Name == "Круг" {
		return float64(s.radius * s.radius * 3.14)
	} else {
		return float64(s.length * s.width)
	}
}

func main() {
	circle := Shape{Name: "Круг", radius: 5}
	rectangle := Shape{Name: "Прямоугольник", length: 4, width: 6}
	//
	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())
}
