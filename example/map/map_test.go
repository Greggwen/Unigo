package mapt

import "testing"

func TestMapCreate(t *testing.T) {
	// How to create a map
	// personSalary := make(map[string]int)
	// var personSalary map[string]int
	//
	// if personSalary == nil {
	// 	t.Log("map is nil. Going to make one.")
	// 	personSalary = make(map[string]int)
	// }

	// personSalary["rs"] = 100
	// personSalary["steve"] = 12000
	// personSalary["jamie"] = 15000
	// personSalary["mike"] = 9000

	personSalary := map[string]int{
		"steven": 12000,
		"jamie":  15000,
	}

	// v, ok := personSalary["xx"]  接收两个参数，一个是值，另一个是结果
	if rose, ok := personSalary["rose"]; ok {
		t.Log(rose)
	} else {
		t.Log("the key named rose isn't exist")
	}
	personSalary["rose"] = 11000
	personSalary["zy"] = 11200

	for v, k := range personSalary {
		t.Log(v, k)
	}

	// Delete a element
	delete(personSalary, "zy")

	t.Log(personSalary, len(personSalary))

}
