package main

/*
	By convention if any method method has a pointer receiver, then all methods should have a pointer receiver.
	Here the rule are broken for showing both kinds of method.
	Method are not allowed in named types that are themselves pointer types.
*/

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point
type P *int

// Distance returns between two points
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Distance returns the distance travelled along the path
func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

// compiler error: invalid receiver type
/*func (P) Sum(value int)  {

}*/

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	// In this case the compiler performs an implicit &p
	p.ScaleBy(2)
	fmt.Println(p)
}
