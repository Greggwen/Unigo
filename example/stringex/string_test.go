package stringTt

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestStringPrint(t *testing.T) {
	str := "Hello World!"

	for i := 0; i < len(str); i++ {
		t.Logf("The bytes is %x\tThe char is %c\n", str[i], str[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\n", runes[i])
	}
}

// rune the alias name of the int32
func TestRunePrint(t *testing.T) {
	// name := "Señor"
	// printChars(name)

	name := "Hello World"
	t.Logf("%T", name)

	runes := []rune(name)
	t.Logf("%T", runes)
	t.Log(runes)
	t.Logf("%x", runes)
	t.Logf("%x", name)
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func TestPrintBytes(t *testing.T) {
	name := "Señor"
	printCharsAndBytes(name)
}

func TestByteSlice(t *testing.T) {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	t.Log(byteSlice)
	str := string(byteSlice)
	t.Log(str)

	byteSlice1 := []byte{67, 97, 102, 195, 169}
	t.Log(byteSlice1)
	str1 := string(byteSlice1)
	t.Log(str1)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str2 := string(runeSlice)
	fmt.Println(str2)
}

func TestStringLength(t *testing.T) {
	name := "Hello World"
	t.Logf("length of `%s` is %d\n", name, utf8.RuneCountInString(name))
	str := "Señor"
	t.Logf("length of `%s` is %d\n", str, utf8.RuneCountInString(str))

}

// string 是不可变的 （Strings are immutable）
// func mutate(s string) string {
// 	// s[0] = "v"
// 	return s
// }

func mutate(s []rune) string {
	s[0] = 'x'
	return string(s)
}

func TestMutate(t *testing.T) {
	h := "hello"
	t.Log(mutate([]rune(h)))
}
