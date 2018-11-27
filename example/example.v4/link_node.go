package main

import "fmt"

type Data interface{}

type Nodes struct {
	data Data
	next *Nodes
}

type LinkList struct {
	size uint64
	head *Nodes
	tail *Nodes
}

// 初始化链表
func (list *LinkList) Init() {
	(*list).size = 0
	(*list).head = nil
	(*list).tail = nil
}

func (list *LinkList) Append(node *Nodes) bool {

	if node == nil {
		return false
	}

	(*node).next = nil


	if (*list).size == 0 {
		(*list).head = node
	} else {
		originTail := (*list).tail
		(*originTail).next = node
	}

	(*list).tail = node
	(*list).size++

	return true
}

func main() {
	n := &Nodes{
		data: 2222,
	}

	list := LinkList{}
	list.Init()

	list.Append(n)

	n1 := &Nodes{
		data: 4444,
	}
	list.Append(n1)

	//fmt.Println(list.size)
	item := list.head
	fmt.Println(item.data)

	var size uint64 = list.size
	var j uint64 = 0
	for j = 0; j < size - 1; j ++ {
		item = (*item).next
		fmt.Println(item.data)
	}

}