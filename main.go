package main

import (
	"fmt"
	"net/http"
	"rest-api/handlers"
)

func main() {
	http.HandleFunc("/users", handlers.UsersRouter)
	http.HandleFunc("/users/", handlers.UsersRouter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:2223", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
