package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/teste", teste)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, mundo!")
}
