package main

import "fmt"

func main() {
	fmt.Println("hola tutorial")
	c := Circle{5}
	r := Rectangle{4, 3}
	PrintShapeProperties(c)
	PrintShapeProperties(r)
}

func PrintShapeProperties(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}