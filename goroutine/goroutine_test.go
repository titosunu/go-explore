package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Halo Selamat datang")
}

func TestCreateGoroutine(t *testing.T) {
	go HelloWorld()
	fmt.Println("ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
} 

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}