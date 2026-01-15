package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("starting sever")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("'/status' was called")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
