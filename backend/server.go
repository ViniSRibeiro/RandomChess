package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/teste", teste)
	http.HandleFunc("/login", s.login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func teste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Olá, mundo!")
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo não permitido", http.StatusMethodNotAllowed)
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Falha no parse do form", http.StatusInternalServerError)
		return
	}

	usuario := r.FormValue("usuario")
	senha := r.FormValue("senha")

	res := s.db.Query
}
