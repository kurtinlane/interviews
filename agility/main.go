package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/character/", characterHandler)
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
