package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Server on port 80")
	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}