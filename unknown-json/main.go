package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var profile map[string]interface{}

	var jsonString string = `{"userId": "U1234", "displayname": "sam.zheng", "friendship": false}`

	if err := json.Unmarshal([]byte(jsonString), &profile); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", profile)
	fmt.Printf("%#v\n", profile["userId"])
}
