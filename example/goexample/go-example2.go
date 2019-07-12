package main

import "fmt"

type Dog struct {
	name string
}

func (d *Dog) SetName(name string) {
	d.name = name
}

func New(name string) Dog {
	return Dog{
		name: name,
	}
}

func main() {
	//dog := Dog{
	//	name: "Kit",
	//}
	//dog.SetName("Little Bit")
	//
	//fmt.Println(dog)
	// 对于一个Dog类型的变量dog来说，调用表达式dog.SetName("Little Bit")会被自动地转译为(&dog).SetName("Little Bit").即先取dog的指针值，再在该指针上调用 SetName方法
	dog := New("Kit")
	dog.SetName("Little Bit")

	fmt.Println(dog)

}
