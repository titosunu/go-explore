package packageGo

import (
	"fmt"
	"slices"
)

func Slices() {
	names := []string{"pratito", "sunu", "darmalaksana"}
	numebers := []int{20, 30, 40, 10, 90, 70, 60}

	fmt.Println(slices.Min(numebers))
	fmt.Println(slices.Max(numebers))
	fmt.Println(slices.Contains(names, "pratito"))
	fmt.Println(slices.Index(names, "darmalaksana"))
}