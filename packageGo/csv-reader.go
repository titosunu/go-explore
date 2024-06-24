package packageGo

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func CsvReader() {
	csvString := "pratito sunu darmalaksana\n" +
		"muhammad razif azzam\n" +
		"timotius gilang adirianto"

	csvReader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"pratito", "sunu", "darmalaksana"})
	_ = writer.Write([]string{"timotius", "gilang"})
	_ = writer.Write([]string{"titus", "rahardian", "azza"})

	writer.Flush()
}