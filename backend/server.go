package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type randomness = uint8

const (
	RD_bitcoin randomness = iota
	RD_standart
)

type server struct {
	db         *sql.DB
	sessions   map[string]string
	randomness randomness
}

type ticker struct {
	Last    float64 `json:"last"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Volume  float64 `json:"vol"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int64   `json:"updated"`
}

type user struct {
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
}

func initServer() server {
	var s server
	s.db = initDB()
	s.sessions = make(map[string]string)
	s.randomness = RD_bitcoin
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func jsonMsg(msg string) string {
	data := map[string]string{"mensagem": msg}
	res, _ := json.Marshal(data)
	return string(res)
}

func jsonToken(msg string) []byte {
	data := map[string]string{"token": msg}
	res, _ := json.Marshal(data)
	return res
}

func jsonRandom(num float64) []byte {
	data := map[string]float64{"token": num}
	res, _ := json.Marshal(data)
	return res
}

func (s *server) cadastro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *server) login(w http.ResponseWriter, r *http.Request) {
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
	s.sessions[token] = userData.Nome

	w.Write(jsonToken(token))
	log.Printf("Logou, token: %s\n", token)
}

func (s *server) random(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, jsonMsg("Metodo não permitido"), http.StatusMethodNotAllowed)
	}
	var valor float64
	switch s.randomness {
	case RD_bitcoin:
		res, err := http.Get("http://cointradermonitor.com/api/pbb/v1/ticker")
		if err != nil {
			log.Fatalf("[!] Erro ao fazer a requisição HTTP: %s", err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Erro ao ler o corpo da resposta: %s", err)
		}

		var tickerData ticker

		if err := json.Unmarshal(body, &tickerData); err != nil {
			log.Fatalf("Erro ao decodificar o JSON: %s", err)
		}
		valor = tickerData.Last
	case RD_standart:
		valor = rand.Float64()
	}
	w.Write(jsonRandom(valor))
}

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func error(w http.ResponseWriter, r *http.Request) {
	http.Error(w, jsonMsg("Zorza Cabeça de melão"), http.StatusBadRequest)
}

func main() {
	var s = initServer()
	log.Println("Servidor foi inicializado com sucesso!")

	http.HandleFunc("/cadastro", s.cadastro)
	http.HandleFunc("/login", s.login)
	http.HandleFunc("/ok", ok)
	http.HandleFunc("/error", error)
	http.HandleFunc("/random", s.random)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
