package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Handleroot(w http.ResponseWriter, r *http.Request) {
	log.Println("'/root' was called")

	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Printf("starting sever")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("'/status' was called")
		w.Header().Set("content-type", "application/json")

		status := map[string]string{"status": "ok"}
		_ = json.NewEncoder(w).Encode(status)
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
