package sliceFolder

import "testing"

func TestSliceCreate(t *testing.T) {

	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:3]

	t.Log(a, b)

	b[1] = 12
	t.Log(a, b)

	c := []int{11, 12, 13}
	t.Log(c, len(c), cap(c))
}

func TestSliceModify(t *testing.T) {
	darr := [...]int{57, 89, 12, 19, 72, 64, 11, 88}
	dslice := darr[2:5]
	t.Log("arra before is ", darr, dslice)

	for i := range dslice {
		dslice[i]++
	}
	t.Log("arra after is ", darr, dslice)

	for _, v := range dslice {
		v++
	}
	t.Log("arra after is ", darr, dslice)

}

func TestSliceAttr(t *testing.T) {
	// a := make([]int, 1, 1)
	//
	// a = append(a, 1)
	// t.Log(a, len(a), cap(a))
	var a []int

	for i := 0; i < 1281; i++ {
		a = append(a, i)
		t.Log("length and capacity is ", len(a), cap(a))

	}
}

func subtactOne(number []int) {

	for i := range number {
		number[i]--
	}
}

func TestSliceSubtact(t *testing.T) {
	nos := []int{8, 7, 6}
	t.Log(nos)
	subtactOne(nos)
	t.Log(nos)
}

func TestSliceMemoryOpt(t *testing.T) {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)

	t.Log(countriesCpy)
}
