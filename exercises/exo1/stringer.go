package main

import "fmt"
import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct{
	largeur, longueur float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return r.largeur*r.longueur
}

func (c Circle) Area() float64{
	return math.Pi*math.Pow(c.radius,2)
}

func (r Rectangle) String() string{
	return fmt.Sprintf("Square of Width %f and Length %f", r.largeur, r.longueur)
}

func (c Circle) String() string{
	return fmt.Sprintf("Circle of radius %f", c.radius)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}
}
