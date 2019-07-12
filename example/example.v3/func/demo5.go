package main

import "fmt"

type Human struct {
	name string
}

func (h Human) String() string {
	return fmt.Sprintf("The human named %s", h.name)
}

type Man struct {
	sex  string
	name string
	Human
}

func (m Man) String() string {
	return fmt.Sprintf("The man named %s, sex is %s\n", m.Human.name, m.sex)
}

func main() {
	human := Human{
		name: "小金库so",
	}
	fmt.Printf("The man is %s\n", human)

	man := Man{
		sex:   "男",
		name:  "Roslise",
		Human: human,
	}
	fmt.Printf("The man is %s\n", man)

}
