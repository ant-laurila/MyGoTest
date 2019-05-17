package main

import (
	"io/ioutil"
	"fmt"
	"github.com/tidwall/gjson"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("data.json")
	check(err)
	textBytes := []byte(dat)
	json := string(textBytes)

	var searchfor string = "Mysteeriorganisaatio"

	results := gjson.Get(json, "results")
	for _, name := range results.Array() {
		result := name.String()
		creators := gjson.Get(result, "research_dataset.creator")
		creators.ForEach(func(key, value gjson.Result) bool {
			println(value.String()) 
			matchorganisation := gjson.Get(value.String(), "member_of.name.fi")
			oname := matchorganisation.String()
			fmt.Println(oname)
			if(oname==searchfor) {
				fmt.Println("YEAH! We got a match! Current is " , oname , " and searchfor is ", searchfor, "\n--------- and since match found, abort looping -------------")
				fmt.Println("And the actual item is to be sent to json-parsing2dspace-format:\n",result)
				return false // abort iterating

			}	
			return true // keep iterating
		})

	}

}