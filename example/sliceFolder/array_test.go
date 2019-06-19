package sliceFolder

import "testing"

func TestSliceInit(t *testing.T) {

	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	t.Log(a)

	var b = [3]int{4, 5, 6}

	t.Log(b)

	c := [...]int{7, 8, 9}
	t.Log(c)

	country := [...]string{"USA", "China", "India", "Germany", "France"}
	country1 := country

	country1[0] = "Unknow"
	t.Log(country, country1)
}

func TestArrayLoop(t *testing.T) {
	num := [...]float64{10.12, 100, 67.99, 2, 4, 19.29}
	for i := 0; i < len(num); i++ {
		t.Log(num[i])
	}

	t.Log("Using range")
	for i, v := range num {
		t.Log(i, v)
	}
}

func TestArrayFunc(t *testing.T) {
	num := [...]int{5, 6, 7, 8, 9}
	t.Log("Origin array value is ", num)
	changeLocal(num, t)
	t.Log("change local new array is ", num)
}

func changeLocal(num [5]int, t *testing.T) {
	num[0] = 100
	t.Log("changeLocal array value is", num)
}

func TestMultiArray(t *testing.T) {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}

	for _, v := range a {
		for _, v2 := range v {
			t.Logf("%s", v2)
		}
	}

}
