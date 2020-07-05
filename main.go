package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start the server on localhost port 8080 and log any errors
	log.Println("http server started on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
