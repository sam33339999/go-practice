package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	To         interface{} `json:"to"`
	ReplyToken string      `json:"replyToken,omitempty"`
}

func (myJson *MyJson) GetType() string {

	if myJson.ReplyToken != "" {
		fmt.Printf("reply\n\n")
		return "reply"
	}

	switch myJson.To.(type) {
	case string:
		fmt.Printf("single\n\n")
		return "single"
	case []interface{}:
		fmt.Printf("multicast\n\n")
		return "multicast"
	}
	fmt.Printf("unMatch\n\n")
	return ""
}

func main() {

	var jsonSingleBlob = []byte(`{"to": "U1234", "messages": [{"type": "text", "text": "Hello world"}]}`)
	var jsonMultiBlob = []byte(`{"to": ["U1234", "U4567"], "messages": [{"type": "text", "text": "Hello world"}]}`)
	var jsonReplyBlob = []byte(`{"replyToken": "12345", "messages": [{"type": "text", "text": "Hello world"}]}`)

	var M1 MyJson
	var M2 MyJson
	var M3 MyJson

	err := json.Unmarshal(jsonSingleBlob, &M1)
	if err != nil {
		fmt.Println("ERROR: ", err)
		panic("shutdown")
	}

	if err = json.Unmarshal(jsonMultiBlob, &M2); err != nil {
		fmt.Println("ERROR: ", err)
		panic("shutdown")

	}

	if err = json.Unmarshal(jsonReplyBlob, &M3); err != nil {
		fmt.Println("ERROR: ", err)
		panic("shutdown")
	}

	fmt.Printf("jsonSingleBlob: %+v\n\n", M1)
	fmt.Printf("jsonMultiBlob: %+v\n\n", M2)
	fmt.Printf("jsonReplyBlob: %+v\n\n", M3)

	M1.GetType()
	M2.GetType()
	M3.GetType()

}
