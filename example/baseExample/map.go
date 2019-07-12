package main

import "fmt"

// Go 语言Map： Map是一种无序的键值对的集合。Map最重要的一点是通过key来快速检索数据，key类似索引，指向数据的值
// Map是一种集合， 可以像迭代数组和切片那样来迭代。Map是无序的，我们无法决定它的返回顺序，因为Map是使用Hash表来实现的

func main() {
	var countryCaptialMap map[string]string
	countryCaptialMap = make(map[string]string)

	countryCaptialMap["France"] = "Paris"
	countryCaptialMap["Italy"] = "罗马"
	countryCaptialMap["Japan"] = "东京"
	countryCaptialMap["India"] = "新德里"
	countryCaptialMap["China"] = "北京"

	for country := range countryCaptialMap {
		fmt.Println(country, "首都是", countryCaptialMap[country])
	}

	//for k, country := range countryCaptialMap {
	//	fmt.Println(k, country)
	//}

	captial, ok := countryCaptialMap["美国"]
	if ok {
		fmt.Println("美国的首都是", captial)
	} else {
		fmt.Println("美国的首都不存在！")
	}

	// Delete() 函数
	delete(countryCaptialMap, "France")
	fmt.Println("法国首都被删除")
	for country := range countryCaptialMap {
		fmt.Println(country, "首都是", countryCaptialMap[country])
	}
}
