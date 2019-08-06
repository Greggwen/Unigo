package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID      int
	customerID int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()

	fmt.Println("Type ", t)
	// fmt.Println("Value ", v)
	fmt.Println("Value ", k)
	fmt.Println("value Kind: ", v.Kind())
}

func main() {
	o := order{
		ordID:      1,
		customerID: 100,
	}

	createQuery(o)
}
