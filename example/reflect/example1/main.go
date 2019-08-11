package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345

	fmt.Println("The Type of Num is ", reflect.TypeOf(num))   // reflect.TypeOf() 直接给到了我们想要的Type的类型
	fmt.Println("The Value of Num is ", reflect.ValueOf(num)) // reflect.ValueOf：直接给到了我们想要的具体的值

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	fmt.Println("Pointer :", pointer)
	fmt.Println("value : ", value)

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println("convertValue :", convertPointer)
	fmt.Println("convertValue : ", convertValue)

}
