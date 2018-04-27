package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln("I don't know if this will work")
}

func main() {
	http.HandleFunc("/main", mainPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
