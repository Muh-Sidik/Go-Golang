package methodlearn

import (
	"fmt"
	"runtime"
)

func getAverage(number []int, ch chan float64)  {
	var sum = 0

	for _, e := range number {
		sum += e
	}

	ch <- float64(sum) / float64(len(number))
}

func getMax(number []int, ch chan int)  {
	var max = number[0]

	for _, e := range number {
		if max < e {
			max = e
		}
	}

	ch <- max
}

func chanSelect()  {
	runtime.GOMAXPROCS(2)

	var number = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}

	fmt.Println("angka :", number)

	var ch1 = make(chan float64)
	go getAverage(number, ch1)

	var ch2 = make(chan int)
	go getMax(number, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <- ch1:
			fmt.Printf("rata- rata \t : %.2f \n", avg)
		case max := <- ch2:
			fmt.Printf("nilai tertinggi \t : %d \n", max)
		}
	}

}

