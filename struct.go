package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

//Struct Literals
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{5, 6}
	v.X = 4
	fmt.Println(v.X)

	//Pointers to structs
	p := &v
	p.X = 1e9 //(*p).X
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
