package packageGo

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (v PersonSlice) Len() int {
	return len(v)
}

func (v PersonSlice) Less(i, j int) bool {
	return v[i].Age < v[j].Age
}

func (v PersonSlice) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func Sort() {
	users := []Person{
		{"pratito", 22},
		{"ivan wijaya", 30},
		{"titus", 24},
		{"bintang", 23},
	}

	sort.Sort(PersonSlice(users))

	fmt.Println(users)
}