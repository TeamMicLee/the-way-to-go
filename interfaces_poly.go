package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Reactangle struct {
	length, width float32
}

func (r Reactangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Reactangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
