package methodlearn

import "fmt"


type saya struct {
	ani float64
}

var data = saya{
	ani: 2705,
}

func dataaa()  {
	fmt.Printf("%E\n", data.ani)
}