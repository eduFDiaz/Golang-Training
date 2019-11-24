// Method demo
package main

import (
	"fmt"
	"log"
)

// Point
type Point struct {
	X int
	Y int
}

func (t *Point) Move(dx int, dy int) {
	t.X += dx
	t.Y += dy
}

// Square
type Square struct {
	Center Point
	Lenght int
}

func (t *Square) Move(dx int, dy int) {
	t.Center.Move(dx, dy)
}

// Value returns the trade value
func (t *Square) Area() int {
	area := t.Lenght * t.Lenght
	return area
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("length must be > 0 (was %f)", length)
	}

	square := &Square{
		Center: Point{x, y},
		Lenght: length,
	}
	return square, nil
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
