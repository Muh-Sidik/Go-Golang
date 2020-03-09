package methodlearn

import (
	"encoding/base64"
	"fmt"
)

func encoding1()  {
	var data = "Tri Apria Ningsih"

	var encodingString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodingString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodingString)
	var decodeString = string(decodeByte)
	fmt.Println(decodeString)

	change := &data
	*change = "http://MuhammadSidik.com"
	var encodingUrlstring = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodingUrlstring)

	var decodeUrlByte, _ = base64.URLEncoding.DecodeString(encodingUrlstring)
	var decode = string(decodeUrlByte)
	fmt.Println(decode)
	
}