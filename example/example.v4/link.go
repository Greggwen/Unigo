package main

import "fmt"

type MyLinkedList struct {
	val int
	next *MyLinkedList
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		val: 0,
		next: nil,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	p := this.next
	for j := 0; j < index; j ++ {
		if p == nil {
			return -1
		}
		p = p.next
	}

	return p.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	firstNode := &MyLinkedList{val:val}
	this.next = firstNode
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	point := this.next
	lastNode := &MyLinkedList{val: val, next: nil}
	for point.next != nil {
		point = point.next
	}
	point.next = lastNode
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	// index == 0 则直接调用 AddAtIndex 添加到首节点
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	point := this.next

	for i := 0; i < index; i ++ {
		point = point.next
	}
	//fmt.Println(point.next.val)
	// 如果point.next == nil， 则直接调用 AddAtTail 添加到尾节点
	if point.next == nil {
		this.AddAtTail(val)
		return
	}

	s := &MyLinkedList{val: val}
	s.next = point.next
	point.next = s
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	point := this.next

	j := 1
	for (point != nil && j < index) {
		point = point.next
		j ++
	}

	if point == nil ||  j > index {

	} else {
		point.next = point.next.next
	}

}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(10)
	obj.AddAtIndex(1, 4)
	obj.DeleteAtIndex(2);

	first := obj.next
	for first != nil {
		fmt.Println(first.val)
		first = first.next
	}
}