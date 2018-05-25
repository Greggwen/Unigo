package main

import "fmt"

// func main() {
// 	i, j := 42, 2701

// 	p := &i
// 	fmt.Println(*p)

// 	*p = 21
// 	fmt.Println(i)

// 	p = &j
// 	*p = *p / 37
// 	fmt.Println(j)
// }

type Vertex struct {
	X int
	Y int
}

// func main() {
// 	// fmt.Println(Vertex{1, 2})
// 	v := Vertex{1, 2}
// 	v.X = 4
// 	fmt.Println(v.X)
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 169
// 	fmt.Println(v)
// }

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
