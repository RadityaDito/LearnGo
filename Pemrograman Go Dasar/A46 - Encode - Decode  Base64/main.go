package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString2 = string(decoded)
	fmt.Println(decodedString2)

	var data3 = "https://kalipare.com/"

	var encodedString3 = base64.URLEncoding.EncodeToString([]byte(data3))
	fmt.Println(encodedString3)

	var decodedByte3, _ = base64.URLEncoding.DecodeString(encodedString3)
	var decodedString3 = string(decodedByte3)
	fmt.Println(decodedString3)
}
