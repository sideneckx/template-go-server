package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JSONExample struct {
	body string
}

func main() {
	//restfull api
	reps, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
	}
	defer reps.Body.Close()
	var jsonBody *[]map[string]interface{} = &[]map[string]interface{}{}
	json.NewDecoder(reps.Body).Decode(jsonBody)
	fmt.Println((*jsonBody)[0])
	// fmt.Printf(err2.Error())
	// fmt.2Println(err2)
}
