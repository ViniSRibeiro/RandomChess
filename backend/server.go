package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	db *sql.DB
}

func initServer() server {
	var s server
	s.db = initDB()
	return s
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Println("Não foi possível abrir o banco app.db")
		log.Fatal(err)
		return nil
	}
	// Inicializamos as tabelas por meio da execução do arquivo banco.sql
	contents, err := os.ReadFile("./banco.sql")
	if err != nil {
		log.Println("Não possível executar o script banco.sql")
		log.Fatal(err)
		return nil
	}
	sql := string(contents)
	if _, err := db.Exec(sql); err != nil {
		log.Println("Não possível executar o script banco.sql")
		log.Fatal(err)
		return nil
	}
	return db
}

// -----------------------------------------------------------------------------

func jsonMsg(msg string) string {
	data := map[string]string{"mensagem": msg}
	res, _ := json.Marshal(data)
	return string(res)
}

func (s *server) cadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, jsonMsg("Use o método POST"), http.StatusMethodNotAllowed)
		return
	}
	// O nome e a senha são dadas via um formulário HTML
	r.ParseForm()
	nome := r.PostForm.Get("nome")
	senha := r.PostForm.Get("senha")
	if nome == "" || senha == "" {
		http.Error(w, jsonMsg("Dados inválidos"), http.StatusBadRequest)
		return
	}

	// O usuário já existe? Retornamos um código de erro nesse caso
	rows, _ := s.db.Query("SELECT nome FROM Usuario WHERE nome = ?", nome)
	if rows.Next() {
		// Usuário já existe no banco de dados (erro: 409)
		msg := fmt.Sprintf("Usuário %s já existe", nome)
		http.Error(w, jsonMsg(msg), http.StatusConflict)
		log.Println(msg)
		return
	}

	// Usuário ainda não existe; escrevemos seus dados no banco
	log.Printf("Criando usuário %s\n", nome)
	_, err := s.db.Exec("INSERT INTO Usuario (nome, senha) VALUES (?, ?)", nome, senha)
	if err != nil {
		msg := fmt.Sprintf("Algo deu errado no cadastro de %s", nome)
		http.Error(w, jsonMsg(msg), http.StatusInternalServerError)
		log.Println(msg)
		return
	}
	w.WriteHeader(http.StatusOK)
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

	nome := r.PostForm.Get("nome")
	senha := r.PostForm.Get("senha")
	if nome == "" || senha == "" {
		http.Error(w, "Usuário e senha vazios", http.StatusBadRequest)
		return
	}

	res, err := s.db.Query("SELECT * FROM Usuario WHERE nome = ? AND senha = ?", nome, senha)
	if err == sql.ErrNoRows || res.Next() {
		log.Printf("[!] Usuário ou senha errados")
		http.Error(w, "Usuario ou senha errados", http.StatusConflict)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	var s = initServer()
	log.Println("Servidor foi inicializado com sucesso!")

	http.HandleFunc("/cadastro", s.cadastro)
	http.HandleFunc("/login", s.login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
