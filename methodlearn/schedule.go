package methodlearn

import (
	"fmt"
	"time"
)

func schedule() {
	done := make(chan bool)

	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <- done:
			ticker.Stop()
			return
		case t := <- ticker.C:
			fmt.Println("Hello", t)
		}
	}

}