package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名字段
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

// Human 实现sayHi方法
func (h Human) SayHi () {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Human 实现Sing方法
func (h Human) Sing (lyrics string) {
	fmt.Println("La la la la ...", lyrics)
}

// Employee 重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-111-xxx"}, "Harard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-xxx"}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", 42, "444-333-xxx"}, "Things Ltd", 5000}

	var i Men

	i = mike
	i.SayHi()
	i.Sing("I love you")

	i = tom
	i.SayHi()
	i.Sing("I love work.")

	x := make([]Men, 3)

	x[0], x[1], x[2] = mike, paul, sam

	for _, value := range x {
		value.SayHi()
	}
}