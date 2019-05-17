package main

import (
	"fmt"
	"encoding/json"
)


func main() {
	var v interface{}


	text := `{"people": [{"craft": "ISS", "name":"Sergey Rizhikov"}, 
	{"craft": "ISS", "name": "AndreyBorisenko"}, 
	{"craft": "ISS", "name": "ShaneKimbrough"}, 
	{"craft": "ISS", "name": "Oleg Novitskiy"},
	{"craft": "ISS", "name": "Thomas Pesquet"}, 
	{"craft":"ISS", "name": "Peggy Whitson"}], 
	"message": "success",
	"number": 6}`


	textBytes := []byte(text)
	err := json.Unmarshal(textBytes, &v)
	data := v.(map[string]interface{})
	if err != nil {
		fmt.Println("error: ",err)
		return
	}
	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}