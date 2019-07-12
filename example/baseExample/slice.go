package main

import "fmt"

// slice Go 语言切片： Go语言切片是对数组的抽象
// Go 语言数组的长度是不可改变的，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片（动态数组）
// 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大

// 定义切片： 切片不需要说明长度。 var identifier []type
// 使用 make()函数来创建切片： var slice1 []type = make([]type, len) , 也可以简写为： slice1 := make([]type, len)

// 切片初始化： s := []int{1, 2, 3}

//func main() {
//var arr = [5]int{1, 2, 4, 5, 6}
//
//slice1 := arr[:]
//fmt.Println(slice1) // [1 2 4 5 6]
//
//slice2 := arr[:4]
//fmt.Println(slice2) // [1 2 4 5]
//
//slice3 := arr[1:]
//fmt.Println(slice3) // [2 4 5 6]
//
//// len() 与 cap()
//var numbers = make([]int , 3)
//printSlice(numbers)

//	var numbers []int
//	printSlice(numbers)
//
//	if numbers == nil {
//		fmt.Printf("切片是空的")
//	}
//}
//
func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// append() 和 copy() 函数

func main() {
	var numbers []int
	printSlice(numbers) // len=0 cap=0 slice=[]
	// 向空切片里添加元素
	numbers = append(numbers, 0)
	printSlice(numbers) // len=1 cap=1 slice=[0]

	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers) // len=2 cap=2 slice=[0 1]

	numbers = append(numbers, 2)
	printSlice(numbers) // len=3 cap=4 slice=[0 1 2]

	numbers = append(numbers, 3)
	printSlice(numbers) // len=4 cap=4 slice=[0 1 2 3]

	numbers = append(numbers, 4)
	printSlice(numbers) // len=5 cap=8 slice=[0 1 2 3 4]

	numbers = append(numbers, 5)
	printSlice(numbers) // len=6 cap=8 slice=[0 1 2 3 4 5]

	numbers = append(numbers, 6)
	printSlice(numbers) // len=7 cap=8 slice=[0 1 2 3 4 5 6]

	numbers = append(numbers, 7)
	printSlice(numbers) // len=8 cap=8 slice=[0 1 2 3 4 5 6 7]

	// 向切片添加多个元素
	numbers = append(numbers, 8, 9, 10)
	printSlice(numbers) // len=11 cap=16 slice=[0 1 2 3 4 5 6 7 8 9 10]
	// 说明： 按Slice扩容机制， 当容量不够时cap(s) 会按当前的len翻倍申请。1 => 2 => 4 => 8 => 16 ... 这也是len与cap不一致的地方
	// 具体参考： https://www.zhihu.com/question/27161493

	// 创建切片numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	printSlice(numbers1)

	copy(numbers1, numbers)
	printSlice(numbers1)
}
