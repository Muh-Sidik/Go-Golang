package methodlearn

import (
	"fmt"
	"runtime"
)

func channel()  {
	runtime.GOMAXPROCS(2)

	var massages = make(chan string)

	var SayToHello = func (who string)  {
		var data = fmt.Sprintf("Hello %s", who)

		massages <- data
	}

	go SayToHello("Sidik")
	go SayToHello("Ani")
	go SayToHello("Azzam")

	massages1 := <-massages
	fmt.Println(massages1)

	massages2 := <-massages
	fmt.Println(massages2)

	massages3 := <-massages
	fmt.Println(massages3)
}