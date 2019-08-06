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
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields ", v.NumField())

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Feild: %d Type: %T Value: %v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := order{
		ordID:      1,
		customerID: 100,
	}

	createQuery(o)
}
