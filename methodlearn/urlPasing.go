package methodlearn

import (
	"fmt"
	"net/url"
)

func urlParsing()  {
	var urlParsing = "http://sidik.com/hello?name=Apria Ningsih&age=19"
	var url, err = url.Parse(urlParsing)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlParsing)

	fmt.Printf("protocol: %s\n", url.Scheme)
	fmt.Printf("host: %s\n", url.Host)
	fmt.Printf("path: %s\n", url.Path)
	var name = url.Query()["name"][0]
	var age = url.Query()["age"][0]
	fmt.Printf("name: %s,\t age: %s", name, age)
}