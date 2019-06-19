package point

import "testing"

// 什么是指针？指针就是一个存储另外一个变量的内存地址的变量

func change(val *int) {
	*val = 55
}

func TestDeclaringPointer(t *testing.T) {
	b := 255
	var a *int = &b
	t.Logf("Type of a is %T", a)
	t.Log("address of b is ", a)
	t.Log("address of b is ", &b)
	t.Log("address of b is ", &a)

	var c *int
	t.Log("The c is", c)

	d := &b
	t.Log("value of b is", *d)
	t.Log("address of b is", d)

	*d++
	t.Log("value of b is ", b)
	t.Log("value of b is ", *d)

	change(d)
	t.Log(*a, b, c)

}
