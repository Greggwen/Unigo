package main

import (
	"fmt"
	"time"
)

type TestStruct struct {
	a int
	b string
}

func (T *TestStruct) setA(i int) {
	T.a = i
}

func NewStructV(value int, s string) TestStruct {
	T := TestStruct{a: value, b: s}
	p := &T
	fmt.Println(p, "__4__")  // &{a:1, b : hello}
	fmt.Println(T, "__5__")  // {a:1, b : hello}
	p.setA(2)
	fmt.Println(p, "__6__")
	fmt.Println(T, "__7__")
	T.setA(3)
	fmt.Println(p, "__8__")
	fmt.Println(T, "__9__")
	defer fmt.Println(T, "__10__")
	defer fmt.Println(*p, "__11__")
	defer fmt.Println(p, "__12__")
	defer fmt.Println(&T, "__13__")
	defer time.Sleep(2 * time.Second)
	defer p.setA(4)
	fmt.Printf("%p__14__\n", p)
	return T
}

func main() {
	q := NewStructV(1, "hello")
	k := &q
	fmt.Printf("%p__1__\n", k)
	fmt.Println(q, "__2__")
	fmt.Println(k, "__3__")
}
