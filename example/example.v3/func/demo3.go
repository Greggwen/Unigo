package main

import "fmt"

func main() {
	// 示例1
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)

	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2) // [a x c]
	fmt.Printf("The original array: %v\n", array1) // [a b c]

	// 示例2
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)

	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2) // [x i z]
	fmt.Printf("The original slice: %v\n", slice1) // [x i z]

	// 示例3
	complexArray1 := [3][]string {
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"J", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2) // [[d e f] [g s i] [o p q]]
	fmt.Printf("The original complex array: %v\n", complexArray1) // [[d e f] [g s i] [J k l]]
}

func modifyArray (arr [3]string) [3]string {
	arr[1] = "x"
	return arr
}

func modifySlice(arr []string) []string {
	arr[1] = "i"
	return arr
}

func modifyComplexArray(arr [3][]string) [3][]string {
	arr[1][1] = "s"
	arr[2] = []string{"o", "p", "q"}
	return arr
}

