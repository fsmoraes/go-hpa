package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func greeting() string {

	x := 0.0001
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "Code.education Rocks!!!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", greeting())
}

func main() {

	http.HandleFunc("/", handler)

	port := "8000"
	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
