package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//note F in front of println
	//note w, in front of "oh"
	fmt.Fprintln(w, "oh")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
