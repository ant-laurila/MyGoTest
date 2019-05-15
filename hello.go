package main

import (
	"fmt"
	"encoding/json"
)

type people struct {
	Number int `json:"number"`
	Msg string `json:"message"`
	Peo string `json:"people"`
}
/*
type dude struct {
	Name string `json:"name"`
}
*/
func main() {
	text := `{"people": [{"craft": "ISS", "name":"Sergey Rizhikov"}, {"craft": "ISS", "name": "AndreyBorisenko"}, {"craft": "ISS", "name": "ShaneKimbrough"}, {"craft": "ISS", "name": "Oleg Novitskiy"},{"craft": "ISS", "name": "Thomas Pesquet"}, {"craft":"ISS", "name": "Peggy Whitson"}], "message": "success","number": 6}`

	textBytes := []byte(text)
	people1 := people{}
	err := json.Unmarshal(textBytes, &people1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(people1.Number)
	fmt.Println("spacer")
	fmt.Println(people1.Msg)
fmt.Println(people1.Peo)
}