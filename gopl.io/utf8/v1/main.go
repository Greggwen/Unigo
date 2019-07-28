package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s);  {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))


	s = "abc"
	fmt.Println(s)
	b := []byte(s)
	fmt.Println(b)
	s2 := string(b)
	fmt.Println(s2)

}

func basename (s string) string {
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == '/' {
			s = s[:i]
			break
		}
	}
	return s
}

