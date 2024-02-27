package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	ip   string = "127.0.0.1"
	port int    = 8795
)

type Dictionary map[string]interface{}

// ToJSON is Dictionary method that converts it to a JSON string
func (dict Dictionary) ToJSON() (string, error) {
	jsonData, err := json.MarshalIndent(dict, "", "  ")
	if err != nil {
		return "", fmt.Errorf("serialization error (%v)", err)
	}
	return string(jsonData), nil
}

// GetCurrentTime returns current server time
func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

// Home handles `GET /` HTTP request
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Лабораторна робота 1")
	if err != nil {
		fmt.Println("Cannot write the response:", err)
		return
	}
}

// Time handles `GET /time` HTTP request
func Time(w http.ResponseWriter, r *http.Request) {
	jsonTime, err := Dictionary{"time": GetCurrentTime()}.ToJSON()
	if err != nil {
		fmt.Println("Cannot convert the time to JSON format:", err)
		return
	}
	_, err = fmt.Fprint(w, jsonTime)
	if err != nil {
		fmt.Println("Cannot write the response:", err)
		return
	}
}

func main() {
	// Basic routing
	http.HandleFunc("/", Home)
	http.HandleFunc("/time", Time)

	// Starting the server
	address := ip + ":" + strconv.Itoa(port)
	fmt.Println("Server is listening on", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("Cannot start the server:", err)
	}
}
