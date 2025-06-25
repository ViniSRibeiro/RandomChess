package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Arquivo separado contendo apenas a lógica de login e de cadastro
// Essa lógica é bem extensa

func (s *Server) cadastro(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, jsonMsg("Use o método POST"), http.StatusMethodNotAllowed)
		return
	}
	// O nome e a senha são dadas via um formulário HTML
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %s", err)
	}

	var userData user

	if err := json.Unmarshal(body, &userData); err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %s", err)
	}
	if userData.Nome == "" || userData.Senha == "" {
		http.Error(w, jsonMsg("Dados inválidos"), http.StatusBadRequest)
		return
	}

	// O usuário já existe? Retornamos um código de erro nesse caso
	rows, _ := s.db.Query("SELECT nome FROM Usuario WHERE nome = ?", userData.Nome)
	if rows.Next() {
		// Usuário já existe no banco de dados (erro: 409)
		msg := fmt.Sprintf("Usuário %s já existe", userData.Nome)
		http.Error(w, jsonMsg(msg), http.StatusConflict)
		log.Println(msg)
		return
	}

	// Usuário ainda não existe; escrevemos seus dados no banco
	log.Printf("Criando usuário %s\n", userData.Nome)
	_, err = s.db.Exec("INSERT INTO Usuario (nome, senha) VALUES (?, ?)", userData.Nome, userData.Senha)
	if err != nil {
		msg := fmt.Sprintf("Algo deu errado no cadastro de %s", userData.Nome)
		http.Error(w, jsonMsg(msg), http.StatusInternalServerError)
		log.Println(msg)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, jsonMsg("Metodo não permitido"), http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Erro ao ler o corpo da resposta: %s", err)
	}

	var userData user

	if err := json.Unmarshal(body, &userData); err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %s", err)
	}
	if userData.Nome == "" || userData.Senha == "" {
		http.Error(w, jsonMsg("Usuário e senha vazios"), http.StatusBadRequest)
		return
	}

	res, err := s.db.Query("SELECT * FROM Usuario WHERE nome = ? AND senha = ?", userData.Nome, userData.Senha)
	if err == sql.ErrNoRows || !res.Next() {
		log.Printf("[!] Usuário ou senha errados")
		http.Error(w, jsonMsg("Usuario ou senha errados"), http.StatusConflict)
		return
	}
	token := uuid.New().String()
	s.sessions[token] = InitSession(userData.Nome)
	// A linha abaixo associa cada nome de usuário a um token, o que permite
	// associá-lo também à sua sessão
	s.userTokens[userData.Nome] = token

	w.Write(jsonToken(token))
	log.Printf("Logou, token: %s\n", token)
}
