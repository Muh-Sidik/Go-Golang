package methodlearn

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string)  {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string)  {
	for message := range ch {
		fmt.Println(message)
	}
}

func chanClose()  {
	runtime.GOMAXPROCS(2)

	var message = make(chan string)

	go sendMessage(message)
	printMessage(message)
}