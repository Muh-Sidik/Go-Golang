package methodlearn

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age int
}

func jsonCon()  {
	//decode json ke variabel object struct
	var jsonString = `{"Name": "Apria Ningsih", "Age": 19}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.Fullname)
	fmt.Println("age :", data.Age)

	//decode json ke map[string]interface{} dan interface{}
	var data1 map[string]interface{}

	var err1 = json.Unmarshal(jsonData, &data1)
	if err1 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("name: ", data1["Name"])
	fmt.Println("age :", data1["Age"])

	var data2 interface{}

	var err2 = json.Unmarshal(jsonData, &data2)

	if err2 != nil {
		fmt.Println(err.Error())
		return
	}
	var decode = data2.(map[string]interface{})
	fmt.Println("name :", decode["Name"])
	fmt.Println("age : ", decode["Age"])

	var jsonArray = `[
		{"Name": "Muhammad Sidik", "Age": 21},
		{"Name": "Tri Apria Ningsih", "Age": 19}
	]`

	//decode json ke array object
	var jsonCon = []byte(jsonArray)

	var data3 []User

	var err3 = json.Unmarshal(jsonCon, &data3)

	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	fmt.Println("name: ", data3[0].Fullname)
	fmt.Println("name: ", data3[1].Fullname)
	//object ke json
	var object = []User{{ "Tri Apria Ningsih", 19 }, { "Muhammad Sidik", 21 }}

	var jsoData, err4 = json.Marshal(object)

	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}

	var jsoString =string(jsoData)

	fmt.Println(jsoString)


}