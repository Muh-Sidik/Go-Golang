package methodlearn

import (
	"fmt"
	"time"
)

func timerr()  {
	fmt.Println("Mulai")
	time.Sleep(time.Second * 5)
	fmt.Println("after 5 second")

	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("mulai")
	<-timer.C
	fmt.Println("selesai")

	// for true {
	// 	fmt.Println("Hai :3")
	// 	time.Sleep(2 * time.Second)
	// }
}