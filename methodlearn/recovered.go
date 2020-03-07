package methodlearn

import (
	"errors"
	"fmt"
	"strings"
)

func catch()  {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validated2(input string) (bool, error)  {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Jangan Pake Spasi")
	}
	return true, nil
}

func reco()  {
	defer catch()

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