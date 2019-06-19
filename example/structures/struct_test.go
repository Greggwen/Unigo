package structTest

import "testing"

// 若struct类型是以大写字母开始命名的，那么它是一个可导出类型， 可以从其他包访问，
// 与之相似的， 其字段命名如果首字母大写也是可以被外部包访问

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func TestStructCreate(t *testing.T) {

	em1 := &Employee{
		firstName: "Rose",
		lastName:  "ping",
		age:       22,
		salary:    500,
	}

	em2 := Employee{"Zy", "ying", 22, 500}

	t.Log(em1, em2)

	t.Log(em1.firstName, (*em1).firstName)

}

func TestAnonymousStruct(t *testing.T) {
	em3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		"Swenly",
		"note",
		20,
		500,
	}

	t.Log(em3)

	var em4 Employee

	t.Log(em4)
}

type Address struct {
	city, state string
}

type Person struct {
	name    string
	age     int
	address Address
}

type Person1 struct {
	name string
	age  int
	Address
}

func TestNestedStruct(t *testing.T) {
	var p Person
	p.name = "Jin"
	p.age = 25
	p.address = Address{
		city:  "BJ",
		state: "Sanlitun",
	}

	t.Logf("p.name is %s", p.name)
	t.Logf("p.age is %d", p.age)
	t.Logf("p.address.city is %s", p.address.city)
	t.Logf("p.address.state is %s", p.address.state)

	var p1 Person1
	p1.name = "Ping"
	p1.age = 25
	p1.Address.city = "Beijing"
	p1.Address.state = "Sanlitun soho"

	t.Logf("p1.name is %s", p1.name)
	t.Logf("p1.age is %d", p1.age)
	t.Logf("p1.city is %s", p1.city)
	t.Logf("p1.state is %s", p1.state)
}
