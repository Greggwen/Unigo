package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	speciality string
}

func main() {
	jane := Student{Human: Human{"小金库so", 18, 110}, speciality: "Biology"}

	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)

	// 修改Skill技能字段
	jane.Skills = []string{"Basketball"}

	fmt.Println("Her skills are ", jane.Skills)

}

// type Student struct {
// 	Human
// 	speciality string
// }

// func main() {
// 	mark := Student{Human{"小金库so", 18, 56}, "Computer Science"}
// 	fmt.Println("His name is ", mark.name)
// }

// type person struct {
// 	name string
// 	age  int
// }

// func Older(p1, p2 person) (person, int) {
// 	if p1.age > p2.age {
// 		return p1, p1.age - p2.age
// 	}

// 	return p2, p2.age - p1.age
// }

// func main() {

// 	var tom person

// 	tom.name, tom.age = "Tom", 18

// 	bob := person{name: "Bob", age: 25}

// 	// paul := person{"Paul", 32}

// 	tb_Older, tb_diff := Older(tom, bob)

// 	fmt.Printf("Of %s and %s, %s is Older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)

// }
