package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const Port int = 8795

type RequiredData struct {
	Time string `json:"time"`
}

type message string

func (m message) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resp, m)
}
func main() {

	data := RequiredData{
		Time: time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Serialization error:", err)
		return
	}

	res := message(jsonData)
	fmt.Println("Server is listening on localhost:" + strconv.Itoa(Port))
	http.HandleFunc("/time", res.ServeHTTP)
	http.ListenAndServe("localhost:"+strconv.Itoa(Port), nil)
}
