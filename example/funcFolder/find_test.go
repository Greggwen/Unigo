package find_test

import (
	"fmt"
	"testing"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func TestFind(t *testing.T) {
	find(10, 11, 12, 14, 21, 11, 10)
}

func change(s ...string) {
	s[0] = "Golang"
	s = append(s, "rose")
	s[1] = "Swenly"
	fmt.Println(s)
}

func TestChangeEle(t *testing.T) {
	welcome := []string{"hello", "world", "rose", "swenly", "flower"}
	change(welcome...)
	t.Log(welcome)
}
