package packageGo

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
	Adress string `required:"true" max:"10"`
	Hobby string `required:"true" max:"10"`
	Zodiac string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("type name", valueType.Name())
	
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "type =", valueField.Type)
		fmt.Println(valueField.Tag.Get("required"))
	}
}

func IsValid(value any) (result bool) {
	result = true
	typeOf := reflect.TypeOf(value)
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if !result {
				return result
			}
		}
	}
	return true
}

func Reflect() {
	readField(Sample{"pratito", "", "", ""})
	readField(Person{}) // from sort.go
	sample := Sample{
		Name: "Pratito",
		Adress: "Lamongan`",
		Hobby: "suka coding",
		Zodiac: "Cancer",
	}
	fmt.Println(IsValid(sample))
}