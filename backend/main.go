package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(
		"/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				log.Println("Hello")
			},
		),
	)

	log.Println("Now server is running on port 3000")
	_ = http.ListenAndServe(":3000", nil)
}
