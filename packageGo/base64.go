package packageGo

import (
	"encoding/base64"
	"fmt"
)

func Base64() {
	value := "pratito sunu darmalaksana"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error bolo", err.Error())
	}
	fmt.Println(string(decoded))
}