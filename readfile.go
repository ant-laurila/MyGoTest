package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/tidwall/gjson"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func lookfor(n string,w map[string]interface{}) {

	for k, v := range w {
		switch v := v.(type) {
		case string:
			fmt.Println("key: ",k, "val: ", v, "(string)")
			if(k==n) {
				fmt.Println("\n!!!!This is it !!!!",k," found with value ",v, "\n")
			}
		case float64:
			fmt.Println("key: ",k, "val: ", v, "(float64)")
		case []interface{}:
			fmt.Println("---------",k,"-----------------")
			fmt.Println(v)
			/*for i, u := range v {
				//fmt.Println("   ", i, u)
				//fmt.Printf("\nTYPE: (%T)",u)
				//fmt.Println(" ")
				i=i
				data := u.(map[string]interface{})
				lookfor("metadata_provider_org",data)
			}
			fmt.Println("---------------------------")
*/
		default:
			fmt.Println("key: ",k,"val: ", v, "(unknown)")
		}
	}
}

func main() {
	var v interface{}
	
	dat, err := ioutil.ReadFile("data.json")
	check(err)

	textBytes := []byte(dat)
	err = json.Unmarshal(textBytes, &v)
	check(err)

	data := v.(map[string]interface{})

	json := string(textBytes)
	tval := gjson.Get(json,"title")
	fmt.Println("tval: ",tval)

	result := gjson.Get(json, "results.#.identifier")
	for _, name := range result.Array() {
		println(name.String())
	}

	lookfor("metadata_provider_org",data)
}