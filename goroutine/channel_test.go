package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		time.Sleep(2 * time.Second)
		channel <- "Pratito"
		fmt.Println("selesai mengirim data ke channel")
	}()	

	data := <- channel
	fmt.Println(data)
	defer close(channel)

	time.Sleep(5 * time.Second)
}