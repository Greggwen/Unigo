package main

import "fmt"

// Go语言范围range
// Go 语言range关键字用于for循环中迭代数组（array）、切片（slice）、通道（channel）或集合（map）的元素

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Printf("sum = %d\n", sum)

	// range 处理map键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range处理枚举Unicode字符串

	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
