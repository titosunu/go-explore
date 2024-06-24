package packageGo

import (
	"container/list"
	"fmt"
)

func PackageList() {
	var data *list.List = list.New()

	data.PushBack("pratito")
	data.PushBack("sunu")
	data.PushBack("darmalaksana")

	for v := data.Front(); v != nil; v = v.Next() {
		fmt.Println(v.Value)
	}
}