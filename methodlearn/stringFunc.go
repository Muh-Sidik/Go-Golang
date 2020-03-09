package methodlearn

import (
	"fmt"
	"strings"
)

func stringFunc()  {
	var apa = strings.Contains("Apria Ningsih", "Ningsih")
	fmt.Println(apa)

	var isPrefix = strings.HasPrefix("Apria Ningsih", "Ap")
	fmt.Println(isPrefix)

	var isPrefix2 = strings.HasPrefix("Apria Ningsih", "Ni")
	fmt.Println(isPrefix2)

	var isSuffix = strings.HasSuffix("Apria Ningsih", "ih")
	fmt.Println(isSuffix) //true

	var isSuffix2 = strings.HasSuffix("Apria Ningsih", "si")
	fmt.Println(isSuffix2) //false

	var Count = strings.Count("Apria Ningsih", "i")
	fmt.Println(Count)

	var indexk = strings.Index("Apria Ningsih", "i")
	fmt.Println(indexk)

	var Split = strings.Split("Apria Ningsih", " ")
	fmt.Println(Split)

	var data = []string{"Tri", "Apria", "Ningsih"}
	var joins = strings.Join(data, "-")
	fmt.Println(joins)


}