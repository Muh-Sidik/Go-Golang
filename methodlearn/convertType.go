package methodlearn

import (
	"fmt"
	"reflect"
	"strconv"
)

func convert()  {
	var str = "200"

	var num, err = strconv.Atoi(str)
	var Value = reflect.ValueOf(num)

	if err == nil {
		fmt.Println(num)
		fmt.Println("typenya", Value.Type())
	}

	var inn = 200
	var con = strconv.Itoa(inn)
	var Valuee = reflect.ValueOf(con)

	fmt.Println("hasilnya :", con,"", "typenya:", Valuee.Type())

	var data1 = "Tri"
	var b1 = []byte(data1)

	fmt.Printf("%d, %d, %d\n", b1[0], b1[1],b1[2])

	var data2 = "Apria"
	var b2 = []byte(data2)

	fmt.Printf("%d, %d, %d, %d, %d\n", b2[0], b2[1],b2[2],b2[3],b2[4])

	var data = "Ningsih"
	var b = []byte(data)

	fmt.Printf("%d, %d, %d, %d, %d,%d,%d\n", b[0], b[1],b[2],b[3],b[4],b[5],b[6])

	var con1 = []byte{84, 114, 105}
	var a = string(con1)

	fmt.Printf("%s\n", a)

	var con2 = []byte{65, 112, 114, 105, 97}
	var a2 = string(con2)

	fmt.Printf("%s\n", a2)

	var con3 = []byte{78, 105, 110, 103, 115,105,104}
	var a3 = string(con3)

	fmt.Printf("%s\n", a3)

	var c int64 = int64('h')
	fmt.Println(c)
}