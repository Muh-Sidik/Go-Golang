package methodlearn

import (
	"flag"
	"fmt"
)

func testFlag() {
	var name = flag.String("name", "Ningsih", "ketik nama")
	var age = flag.Int64("age", 19, "ketik umur")

	flag.Parse()
	fmt.Printf("name: %s\n", *name)
	fmt.Printf("age: %d\n", *age)
}