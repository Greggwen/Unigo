package main

import (
	"reflect"
	"strconv"
	"fmt"
)

// 反射类型为reflect包提供，它定义了两个重要的类型： Type 和 Value。 Type表示Go语言的一个类型，它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}

}

type Person struct {
	Name string
	Age int
}

func main() {

	ret := Any(10)
	fmt.Println(ret)

	ret1 := Any("stringadsfa")
	fmt.Println(ret1)

	//s1 := []int{1, 2, 3, 4, 5}
	s1 := &Person{
		"xululu",
		21,
	}
	ret2 := Any(s1)
	fmt.Println(ret2)

	//t := reflect.TypeOf(3)
	//fmt.Println(t.String()) // int
	//fmt.Println(t) // int
}
