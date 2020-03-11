package methodlearn

import (
	"fmt"
	"os"
)

func ArguMen()  {
	var ArgsRaw = os.Args
	fmt.Printf("-> %v\n", ArgsRaw)

	var args = ArgsRaw[1:]
	fmt.Printf("-> %v\n", args)
}