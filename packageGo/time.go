package packageGo

import (
	"fmt"
	"time"
)

func Time() {
	timeNow := time.Now()
	fmt.Println(timeNow.Local())

	utc := time.Date(2001, time.June, 27, 0, 0, 0, 0, time.UTC)
	fmt.Println("my birthday", utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error guys", err.Error())
		} else {
		fmt.Println(valueTime)
	}

	duration := time.Second * 100
	duration1 := time.Millisecond * 100
	result := duration - duration1
	fmt.Println(result)
}