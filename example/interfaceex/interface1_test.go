package interfaceTest

import (
	"fmt"
	"testing"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Contract struct {
	empID    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) int {
	expense := 0

	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	return expense
}

func TestCalculator(t *testing.T) {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}

	expense := totalExpense(employees)

	t.Log(expense)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func TestEmtpyInterface(t *testing.T) {
	s := "Hello World"
	describe(s)
	i := 55
	describe(i)

	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)
}

func assert(i interface{}) {
	s, ok := i.(int)
	fmt.Println(s, ok)
}

func TestAssert(t *testing.T) {
	var s interface{} = 56
	assert(s)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func TestTypeSwitch(t *testing.T) {
	findType("Naveen")
	findType(77)
	findType(89.98)
}
