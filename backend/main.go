package main

import (
	"log"
	"net/http"
)

func main() {
	var s = initServer()
	log.Println("Servidor foi inicializado com sucesso!")

	http.HandleFunc("/cadastro", s.cadastro)
	http.HandleFunc("/login", s.login)
	http.HandleFunc("/ok", ok)
	http.HandleFunc("/error", error)
	http.HandleFunc("/random", s.random)
	http.HandleFunc("/esperaJogo", s.esperaJogo)
	http.HandleFunc("/ws", handleWebSocket)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
