package main

import (
	"fmt"
	"math"
)

type Vertex3 struct {
	X, Y float64
}

type Myfloat float64

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Methods are functions
func AbsFunc(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Methods continued
func (f Myfloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Pointer receivers
func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Pointers and functions
func ScaleFunc(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	fmt.Println(v.Abs())

	fmt.Println(AbsFunc(v))

	f := Myfloat(-math.Sqrt(2))
	fmt.Println(f.Abs3())

	v.Scale(10)
	fmt.Println(v.Abs())

	ScaleFunc(&v, 10)
	fmt.Println(v.Abs())

	//Methods and pointer indirection
	h := Vertex3{3, 4}
	h.Scale(2)
	ScaleFunc(&h, 10)

	p := &Vertex3{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(h, p)

	k := Vertex3{3, 4}
	fmt.Println(k.Abs())
	fmt.Println(AbsFunc(k))

	n := &Vertex3{4, 3}
	fmt.Println(n.Abs())
	fmt.Println(AbsFunc(*n))

	//Choosing a value or pointer receiver
	v2 := &Vertex3{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs())
	v2.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs())

}
