package main

import (
	"fmt"
)

type MyInt1 int
type MyInt2 = int

func (i MyInt1) m1() {
	fmt.Println("MyInt1.m1")
}

// func (i MyInt2) m2() {
// 	fmt.Println("MyInt2.m2")
// }

func main() {
	var i1 MyInt1 = MyInt1(2)
	// var i2 MyInt2
	i1.m1()

	fmt.Println(i1)
	// i2.m2()
}

// type myInt int

// func (i *myInt) add(another int) myInt {

// 	fmt.Println(*i, another)
// 	*i = *i + myInt(another)
// 	fmt.Println(*i, another)

// 	return *i
// }

// func main() {
// 	// var ipv4 = [...]uint{192, 168, 0, 4} // ... 表示需由Go编译器计算该值的元素数量并以此获得其长度

// 	// // append(ipv4, 6)

// 	// fmt.Println(ipv4)
// 	//

// 	i1 := myInt(1)
// 	// i2 := myInt(2)
// 	i2 := i1.add(2)

// 	fmt.Println(i1, i2)

// }
