package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)
	fmt.Println("server started!!")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
