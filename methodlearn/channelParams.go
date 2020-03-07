package methodlearn

import (
	"fmt"
	"runtime"
)

func printMassage(data chan string)  {
	fmt.Println(<-data)
}

func chanParams()  {
	runtime.GOMAXPROCS(2)

	var massages = make(chan string)

	for _, each := range []string{"sidik", "ani", "lia", "azzam"} {
		go func(siapa string) {
			var data = fmt.Sprintf("hello %s", siapa)
			massages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMassage(massages)
	}
}