package methodlearn

import (
	"fmt"
	"runtime"
)

func chanBuff()  {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("data diterima", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("data dikirim", i)
		messages <- i
	}
}