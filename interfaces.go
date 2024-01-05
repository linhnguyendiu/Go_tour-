package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

//Interfaces are implemented implicitly
type I interface {
	M()
}

type T struct {
	S string
}

// func (t T) M() {
// 	fmt.Println(t.S)
// }

//Interface values
type F float64

func (f F) M() {
	fmt.Println(f)
}

//Interface values with nil underlying values
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// nil interface
type NI interface {
	N()
}

//Type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt(2))
	v := Vertex4{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	fmt.Println(a.Abs())

	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	// // nil interface
	// var nI NI
	// describeNil(nI)
	// nI.N()

	//The empty interface
	var eI interface{}
	describeEmp(eI)
	eI = 42
	describeEmp(eI)

	eI = "hello"
	describeEmp(eI)

	//Type assertions
	s := eI.(string)
	fmt.Println(s)

	s, ok := eI.(string)
	fmt.Println(s, ok)

	g, ok := eI.(float64)
	fmt.Println(g, ok)

	// g = eI.(float64) // panic
	// fmt.Println(g)

	// Type switch
	do(21)
	do("hello")
	do(true)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeNil(ni NI) {
	fmt.Printf("(%v, %T)\n", ni, ni)
}

func describeEmp(eI interface{}) {
	fmt.Printf("(%v, %T)\n", eI, eI)
}
