package main

import (
	"database/sql"
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

func (s *server) cadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// O nome e a senha são dadas via um formulário HTML
	r.ParseForm()
	nome := r.PostForm.Get("nome")
	senha := r.PostForm.Get("senha")
	if nome == "" || senha == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// O usuário já existe? Retornamos um código de erro nesse caso
	res := s.db.QueryRow("SELECT nome FROM Usuario WHERE nome = ?", nome)
	if res.Err() != nil {
		// Usuário já existe no banco de dados
		log.Printf("Usuário %s já existe\n", nome)
		w.WriteHeader(http.StatusConflict)
		return
	}
	// Usuário ainda não existe; escrevemos seus dados no banco
	log.Printf("Criando usuário %s\n", nome)
	_, err := s.db.Exec("INSERT INTO Usuario (nome, senha) VALUES (?, ?)", nome, senha)
	if err != nil {
		log.Printf("[!] Algo deu errado no cadastro de %s\n", nome)
		w.WriteHeader(http.StatusInternalServerError)
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
