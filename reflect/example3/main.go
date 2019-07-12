package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, ", my name is", u.Name)
}

func main() {
	u := User{
		Id:   1,
		Name: "Rollay",
		Age:  20,
	}

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("Corwien")}

	mv.Call(args)
}
