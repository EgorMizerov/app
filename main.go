package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Main page")
	})
	http.HandleFunc("/foo", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Response for foo request")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic("Ошибка при запуске сервера: " + err.Error())
	}
}