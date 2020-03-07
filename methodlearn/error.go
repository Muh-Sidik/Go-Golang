package methodlearn

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error)  {
	// var input string
	// fmt.Print("masukkan nomor :")
	// fmt.Scanln(&input)

	// var number int
	// var err error

	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, "adalah angka")
	// } else {
	// 	fmt.Println(input, "bukan angka")
	// 	fmt.Println(err.Error())
	// }

	if strings.TrimSpace(input) == "" {
		return false, errors.New("ini Error lah")
	}

	return true, nil


}

func errora()  {
	var name string

	fmt.Print("Ketik nama :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Halo", name)
	} else {
		fmt.Println(err.Error())
	}
}