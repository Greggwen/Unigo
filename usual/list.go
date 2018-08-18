package main

import (
	"container/list"
	"fmt"
)

func main() {
	link := list.New()

	link.PushBack(101)
	link.PushBack(102)
	link.PushBack(103)

	for e := link.Front(); e != nil; e = e.Next() {

		fmt.Println(e.Value)
	}
}
