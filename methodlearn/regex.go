package methodlearn

import (
	"fmt"
	"regexp"
	"strings"
)

func regex()  {
	var text = strings.ToLower("Sidik Ani Lia")
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var step1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v\n", step1)

	var step2 = regex.FindAllString(text, 3)
	fmt.Printf("%#v\n", step2)

	text = "Sidik Ani"

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	var findString = regex.FindString(text)
	fmt.Println(findString)

}