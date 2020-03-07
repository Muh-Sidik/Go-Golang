package methodlearn

import (
	"errors"
	"fmt"
	"strings"
)

func validated(input string) (bool, error)  {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Jangan Pake Spasi")
	}
	return true, nil
}

func paniced()  {
	var name string

	fmt.Print("Ketik Nama: ")
	fmt.Scanln(&name)

	if valid, err := validated(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}