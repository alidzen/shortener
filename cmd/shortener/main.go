package main

import (
	"log"
	"net/http"
)

var urls = make(map[string]string)

func main() {
	http.HandleFunc("/", HandleRequests)
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт, который будем слушать
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
