package goroutine

import (
	"fmt"
	"strconv"
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

func GiveResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Pratito Sunu Darmalaksana"
}

func TestChannelParameter(t *testing.T) {
	channel := make(chan string)

	go GiveResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}

func OnlyIn(channel chan <- string) {
	time.Sleep(2 * time.Second)
	channel <- "Pratito Sunu Darmalaksana"
}

func OnlyOut(channel <- chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Pratito"
		channel <- "Sunu"
		channel <- "Darmalaksana"
	}()

	go func(){
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()

	fmt.Println("Success")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveResponse(channel1)
	go GiveResponse(channel2)

	counter := 0
	for {
		select {

		case data := <- channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <- channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}

	}

}